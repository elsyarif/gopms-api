package entities

type Group struct {
	Id          string `json:"id,omitempty" db:"id" `
	Name        string `json:"name,omitempty" db:"name" binding:"required"`
	Description string `json:"description,omitempty" db:"description" binding:"required"`
}
