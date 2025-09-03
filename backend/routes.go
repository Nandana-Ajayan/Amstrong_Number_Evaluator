package main

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"log"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type apiError struct {
	Message string `json:"message"`
}

func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(v)
}

func RegisterRoutes(r *mux.Router, db *gorm.DB) {

	log.Println("Routes registered successfully")

	// Health check
	r.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Health route was called")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	}).Methods("GET")

	// Register user
	// Register or login user
	r.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		var req struct{ Email string }
		_ = json.NewDecoder(r.Body).Decode(&req)
		if req.Email == "" {
			writeJSON(w, http.StatusBadRequest, apiError{"email required"})
			return
		}

		var user User
		err := db.Where("email = ?", req.Email).First(&user).Error
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				// Create new user
				user = User{Email: req.Email}
				if err := db.Create(&user).Error; err != nil {
					writeJSON(w, http.StatusInternalServerError, apiError{"failed to create user"})
					return
				}
				writeJSON(w, http.StatusCreated, user)
				return
			}
			writeJSON(w, http.StatusInternalServerError, apiError{"db error"})
			return
		}

		// User exists → return user_id
		writeJSON(w, http.StatusOK, user)
	}).Methods("POST")

	// Get Armstrong numbers for a user
	r.HandleFunc("/users/{id}/numbers", func(w http.ResponseWriter, r *http.Request) {
		idStr := mux.Vars(r)["id"]
		id, _ := strconv.Atoi(idStr)

		log.Printf(" Fetch Armstrong numbers for UserID=%d\n", id)

		var nums []ArmstrongNumber
		if err := db.Where("user_id = ?", id).Find(&nums).Error; err != nil {
			writeJSON(w, http.StatusInternalServerError, apiError{"failed"})
			return
		}
		writeJSON(w, http.StatusOK, nums)
	}).Methods("GET")

	// Verify Armstrong number
	r.HandleFunc("/verify", func(w http.ResponseWriter, r *http.Request) {

		log.Println(" Verify number request received")
		var req struct {
			UserID uint  `json:"user_id"`
			Number int64 `json:"number"`
		}
		_ = json.NewDecoder(r.Body).Decode(&req)

		if req.Number <= 0 {
			writeJSON(w, http.StatusBadRequest, apiError{"positive integer required"})
			return
		}

		// Check if user exists
		var user User
		if err := db.First(&user, req.UserID).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				writeJSON(w, http.StatusNotFound, apiError{"user not found"})
			} else {
				writeJSON(w, http.StatusInternalServerError, apiError{"db error"})
			}
			return
		}

		// Armstrong check
		if isArmstrong(req.Number) {
			an := ArmstrongNumber{UserID: req.UserID, Number: req.Number}
			db.Create(&an)
			writeJSON(w, http.StatusOK, map[string]any{"is_armstrong": true, "saved": true})
		} else {
			writeJSON(w, http.StatusOK, map[string]any{"is_armstrong": false, "saved": false})
		}
	}).Methods("POST")

	// Get all users with their Armstrong numbers (global view)
	r.HandleFunc("/users/all", func(w http.ResponseWriter, r *http.Request) {
		// Optional pagination params

		log.Println(" Fetch all users with their Armstrong numbers")
		pageStr := r.URL.Query().Get("page")
		sizeStr := r.URL.Query().Get("size")

		page, _ := strconv.Atoi(pageStr)
		size, _ := strconv.Atoi(sizeStr)
		if page < 1 {
			page = 1
		}
		if size < 1 {
			size = 10
		}

		var users []User
		offset := (page - 1) * size

		if err := db.Offset(offset).Limit(size).Find(&users).Error; err != nil {
			writeJSON(w, http.StatusInternalServerError, apiError{"failed to fetch users"})
			return
		}

		// Attach Armstrong numbers for each user
		type UserWithNumbers struct {
			UserID    uint              `json:"user_id"`
			Email     string            `json:"email"`
			Numbers   []ArmstrongNumber `json:"armstrong_numbers"`
			CreatedAt string            `json:"created_at"`
		}

		var result []UserWithNumbers
		for _, u := range users {
			var nums []ArmstrongNumber
			db.Where("user_id = ?", u.UserID).Find(&nums)

			result = append(result, UserWithNumbers{
				UserID:    u.UserID,
				Email:     u.Email,
				Numbers:   nums,
				CreatedAt: u.CreatedAt.Format("2006-01-02 15:04:05"),
			})
		}

		writeJSON(w, http.StatusOK, map[string]any{
			"page":  page,
			"size":  size,
			"users": result,
		})
	}).Methods("GET")

}

// Armstrong logic
func isArmstrong(n int64) bool {
	if n <= 0 {
		return false
	}
	tmp := n
	digits := 0
	for x := n; x > 0; x /= 10 {
		digits++
	}
	var sum int64
	for tmp > 0 {
		d := tmp % 10
		pow := int64(1)
		for i := 0; i < digits; i++ {
			pow *= d
		}
		sum += pow
		tmp /= 10
	}
	return sum == n
}
