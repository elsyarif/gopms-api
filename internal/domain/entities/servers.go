package entities

type Server struct {
	Id         string `json:"id,omitempty" db:"id" form:"id"`
	GroupId    string `json:"group_id,omitempty" db:"group_id" form:"group_id" binding:"required"`
	ServerName string `json:"server_name,omitempty" db:"server_name" form:"server_name" binding:"required"`
	Location   string `json:"location,omitempty" db:"location" form:"location" binding:"required"`
	Status     string `json:"status,omitempty" db:"status" form:"status" binding:"required"`
	Memory     int    `json:"memory,omitempty" db:"memory" form:"memory" binding:"required"`
	Ip         string `json:"ip,omitempty" db:"ip" form:"ip" binding:"required"`
}
