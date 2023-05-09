package entities

import "time"

type User struct {
	Id        string    `db:"id"   json:"id,omitempty" form:"id"`
	Name      string    `db:"name" json:"name,omitempty" binding:"required" form:"name"`
	Username  string    `db:"username" json:"username,omitempty" binding:"required" form:"username"`
	Password  string    `db:"password" json:"password,omitempty" binding:"required" form:"password"`
	IsActive  bool      `db:"is_active" json:"is_active,omitempty"`
	CreatedAt time.Time `db:"created_at" json:"-"`
	UpdatedAt time.Time `db:"updated_at" json:"-"`
}

type UserResponse struct {
	Id        string    `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Username  string    `json:"username,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func UserToResponse(user *User) UserResponse {
	return UserResponse{
		Id:        user.Id,
		Name:      user.Name,
		Username:  user.Username,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}
