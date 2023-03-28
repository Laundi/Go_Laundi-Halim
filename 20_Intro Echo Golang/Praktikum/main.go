package main

  import (
    "fmt"
    "net/http"
  
    "github.com/labstack/echo/v4"
  )
  
type User struct {
  Id       int    `json:"id" form:"id"`
  Name     string `json:"name" form:"name"`
  Email    string	`json:"email" form:"email"`
  Password string	`json:"password" form:"password"`
}

var users []User

// -------------------- controller --------------------

// get all users
func GetUsersController(c echo.Context) error {
  return c.JSON(http.StatusOK, map[string]interface{}{
    "messages": "success get all users",
    "users":    users,
  })
}

// get user by id
func GetUserController(c echo.Context) error {
	id := c.Param("id")
	user := User{}
	for _, u := range users {
		if fmt.Sprintf("%d", u.Id) == id {
			user = u
		}
	}
	if user.Id == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{
			"message": "user not found",
		})
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message": fmt.Sprintf("sukses mengambil data user id %s", id),
    "user":    user,
	})
}

// delete user by id
func DeleteUserController(c echo.Context) error {
  id := c.Param("id")
  index := 0
  for i, u := range users {
    if fmt.Sprintf("%d", u.Id) == id {
      index = i
    }
  }
  users = append(users[:index], users[index+1:]...)
  return c.JSON(http.StatusOK, echo.Map{
    "message": fmt.Sprintf("sukses menghapus data user id %s", id),
  })
}

// update user by id
func UpdateUserController(c echo.Context) error {
  id := c.Param("id")
  userParam := new(User)
  if err := c.Bind(userParam); err != nil {
    return c.JSON(http.StatusBadRequest, echo.Map{
      "message": "failed to binding data",
    })
  }
  for i := range users {
    if fmt.Sprintf("%d", users[i]. Id) == id {
      users[i].Name = userParam.Name
      users[i].Email = userParam.Email
      users[i].Password = userParam.Password
    }
  }
  return c.JSON(http.StatusOK, echo.Map{
    "message": fmt.Sprintf("sukses mengubah data user id %s", id),
  })
}

// create new user
func CreateUserController(c echo.Context) error {
  // binding data
  user := User{}
  c.Bind(&user)

  if len(users) == 0 {
    user.Id = 1
  } else {
    newId := users[len(users)-1].Id + 1
    user.Id = newId
  }
  users = append(users, user)
  return c.JSON(http.StatusOK, map[string]interface{}{
    "messages": "success create user",
    "user":     user,
  })
}

// ---------------------------------------------------
func main() {
  e := echo.New()
  // routing with query parameter
  e.GET("/users", GetUsersController)
  e.POST("/users", CreateUserController)
  e.GET("/user/:id", GetUserController)
  e.PUT("/users/:id", UpdateUserController)
  e. DELETE("/users/:id", DeleteUserController)

  // start the server, and log if it fails
  e.Logger.Fatal(e.Start(":8000"))
}