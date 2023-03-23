package controllers

import (
	"database/sql"

	"github.com/revel/revel"
)

type UserController struct {
	*revel.Controller
}

func (c UserController) TestPullRequest() revel.Result {
	return c.RenderText("Hello World")

}

// GetUserById...
func (c UserController) GetUserById(id int) revel.Result {
	db := connect()
	var user User
	err := db.QueryRow("SELECT user_id, name, email FROM users WHERE user_id=?", id).Scan(&user.Id, &user.Name, &user.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			return c.NotFound("User not found")
		}
		return c.RenderJSON(ErrorResponse{Code: 500, Message: "Error retrieving user"})
	}
	return c.RenderJSON(user)
}

// InsertUser...
func (c UserController) InsertUser(name string, email string) revel.Result {
	db := connect()
	result, err := db.Exec("INSERT INTO users(name, email) VALUES(?, ?)", name, email)
	if err != nil {
		return c.RenderJSON(ErrorResponse{Code: 500, Message: "Insert user failed"})
	}
	id, err := result.LastInsertId()
	if err != nil {
		return c.RenderJSON(ErrorResponse{Code: 500, Message: "Insert user failed"})
	}
	user := User{Id: int(id), Name: name, Email: email}
	return c.RenderJSON(user)
}

// UpdateUser...
func (c UserController) UpdateUser(id int, name string, email string) revel.Result {
	db := connect()
	_, err := db.Exec("UPDATE users SET name = ? , email = ? WHERE user_id = ?", name, email, id)
	if err != nil {
		return c.RenderJSON(ErrorResponse{Code: 500, Message: "Update user failed"})
	}
	user := User{Id: id, Name: name, Email: email}
	return c.RenderJSON(user)
}

// DeleteUser...
func (c UserController) DeleteUser(id int) revel.Result {
	db := connect()
	_, err := db.Exec("DELETE FROM users WHERE user_id = ?", id)
	if err != nil {
		return c.RenderJSON(ErrorResponse{Code: 500, Message: "Delete user failed"})
	}
	return c.RenderJSON(SuccessResponse{Code: 200, Message: "Delete user success"})
}

// RenderError...
func (c UserController) RenderError(code int, message string) revel.Result {
	c.Response.Status = code
	return c.RenderJSON(ErrorResponse{Code: code, Message: message})
}

// RenderSuccess...
func (c UserController) RenderSuccess(code int, message string) revel.Result {
	c.Response.Status = code
	return c.RenderJSON(SuccessResponse{Code: code, Message: message})
}
