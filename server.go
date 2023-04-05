package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)
type User struct {
	// ID    string    `json:"id" xml:"id" form:"id" query:"id"`
	Name  string `json:"name" xml:"name" form:"name" query:"name"`
	Email string `json:"email" xml:"email" form:"email" query:"email"`
}
type allUsers []User

var users = allUsers{
	{
		// ID:          "1",
		Name:       "John Doe",
		Email: 		"someone@example.com",
	},
}
func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.POST("/users", saveUser)
	e.GET("/users/:id", getUser)
	// e.PUT("/users/:id", updateUser)
	// e.DELETE("/users/:id", deleteUser)


	e.Logger.Fatal(e.Start(":8081"))
}

// e.GET("/users/:id", getUser)
func getUser(c echo.Context) error {
  	// User ID from path `users/:id`
  	id := c.Param("id")

	return c.String(http.StatusOK, id)
}

// e.POST("/save", save)
func saveUser(c echo.Context) error {
	// // Get name and email
	name := c.FormValue("name")
	email := c.FormValue("email")
	// // fmt.Printf(name,email)
	// return c.String(http.StatusOK, "name:" + name + ", email:" + email) 

	// u := new(User)
	// if err := c.Bind(u); err != nil {
	// 	return err
	// }
	user := []allUsers {
			Name: name, 
			Email: email,
		} 
	append(user, users)
	// return c.JSON(http.StatusCreated, u)
	// or
	return c.XML(http.StatusCreated, user)
}
