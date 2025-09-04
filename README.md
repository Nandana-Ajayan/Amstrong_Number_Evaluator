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

### 1. Clone the Repository

```bash
git clone https://github.com/your-username/armstrong_number_evaluator.git
cd armstrong_number_evaluator

## 2. Set Up PostgreSQL Database

Open **pgAdmin4** or your preferred PostgreSQL client.

Create a new database:

```sql
CREATE DATABASE armstrong_db;
