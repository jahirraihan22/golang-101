package main

import (
	"database/sql"
	"net/http"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
)

// User represents a user record in the database
type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	// Password string `json:"password"`
}

var db *sql.DB

func main() {
	// Set up MySQL database connection
	var err error
	db, err = sql.Open("mysql", "root:root@tcp(localhost:3307)/test_go_db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Echo instance
	e := echo.New()

	// Routes
	e.GET("/users", getUsers)
	e.GET("/user/:id", getUser)
	e.POST("/users", createUser)
	// e.PUT("/users/:id", updateUser)
	// e.DELETE("/users/:id", deleteUser)

	// Start server
	e.Logger.Fatal(e.Start(":8081"))
}

// getUsers returns a list of all users in the database
func getUsers(c echo.Context) error {
	// Query database for all users
	rows, err := db.Query("SELECT id,username,email FROM users")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]error{"error": err})
	}
	defer rows.Close()

	// Loop through rows and create a list of users
	users := make([]User, 0)
	for rows.Next() {
		var user User
		err := rows.Scan(&user.ID, &user.Name, &user.Email)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]any{"error": err, "message": "database error"})
		}
		users = append(users, user)
	}

	// Return list of users
	return c.JSON(http.StatusOK, users)
}

// getUser returns a single user record from the database
func getUser(c echo.Context) error {
	// Get user ID from URL parameters
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{"error": err, "message": "invalid user ID"})
	}

	// Query database for user by ID
	row := db.QueryRow("SELECT id, username, email FROM users WHERE id = ?", id)

	// Scan user data into User struct
	var user User
	err = row.Scan(&user.ID, &user.Name, &user.Email)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]any{"message": "user not found", "error": err})
	}

	// Return user data
	return c.JSON(http.StatusOK, user)
}

// createUser adds a new user record to the database
func createUser(c echo.Context) error {
	// Get user data from request body
	var user User
	err := c.Bind(&user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid user data"})
	}

	// Insert new user record into database
	result, err := db.Exec("INSERT INTO users (username, email) VALUES (?, ?)", user.Name, user.Email)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "database error"})
	}

	// Get ID of newly inserted user record
	id, err := result.LastInsertId()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "database error"})
	}
	user.ID = int(id)

	// Return new user data
	return c.JSON(http.StatusCreated, user)
}

//
