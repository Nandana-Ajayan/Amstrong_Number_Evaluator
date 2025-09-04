# Armstrong Number Evaluator

Armstrong Number Evaluator is a full-stack web application built with **Go** and **React** that allows users to register, log in with email, verify Armstrong numbers, and view their personal collection of numbers.

---

## Features

- **User Authentication:** Register or log in using an email. Only registered users can verify numbers.  
- **Number Verification:** Verify if a number is an Armstrong number; valid numbers are saved to the user’s account.  
- **Personal Dashboard:** View all Armstrong numbers saved by the logged-in user.  
- **Input Validation:** Ensures correct email format and positive numbers on both frontend and backend.  
- **API Responses:** Backend provides clear success or error messages with HTTP status codes.  
- **Secure Input Handling:** Emails are sanitized (trimmed, lowercased) before storage.
- **Global Dashboard (Restricted):** A global users endpoint is implemented in the backend.  
  This allows logged-in users to view all registered users (with privacy-safe details) and their Armstrong numbers.  
  It is **not visible to non-logged-in users**


---

## Technical Stack

- **Backend:** Go (Golang)  
- **Database:** PostgreSQL  
- **Frontend:** React  

---

## Setup and Installation

 **1. Clone the Repository**

```bash
git clone https://github.com/your-username/armstrong_number_evaluator.git
cd armstrong_number_evaluator
```

 **2. Set Up PostgreSQL Database** 

Open **pgAdmin4** or your preferred PostgreSQL client.

Create a new database:

```sql
CREATE DATABASE armstrong_db;
```

Create the required tables using the following query:
```sql
-- Users Table
CREATE TABLE users (
    user_id SERIAL PRIMARY KEY,
    email VARCHAR(255) UNIQUE NOT NULL,
    created_at TIMESTAMPTZ DEFAULT NOW()
);
```
```sql
-- Armstrong Numbers Table
CREATE TABLE armstrong_numbers (
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL REFERENCES users(user_id) ON DELETE CASCADE,
    number BIGINT NOT NULL,
    created_at TIMESTAMPTZ DEFAULT NOW()
);

```
**3. Configure the Backend (Go)**


- Navigate to the backend directory:

```bash
cd backend
```
- Open the `db.go` file and update the `dsn` variable with your PostgreSQL credentials:


**3. Configure the Frontend (React)**


- Navigate to the frontend directory:
   ```bash
cd frontend
```
- Install node.js dependencies.

Running the Application

Run the backend and frontend servers in two separate terminals.

 Start the Backend Server

1. In a terminal, navigate to the backend directory.
.Run the following command:go run .
.The backend server will start and be available at:http://localhost:8080

Start the Frontend Server
.In a new terminal, navigate to the frontend directory:
.Run the following command:npm start
.This will automatically open the application in your default web browser at: http://localhost:3000


