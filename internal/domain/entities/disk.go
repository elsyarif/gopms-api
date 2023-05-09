package entities

type Disk struct {
	Id       string `json:"id,omitempty" db:"id"`
	ServerId string `json:"server_id,omitempty" db:"server_id" binding:"required"`
	Name     string `json:"name,omitempty" db:"name" binding:"required"`
	Total    int    `json:"total,omitempty" db:"total" binding:"required"`
}
