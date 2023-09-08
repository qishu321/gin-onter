package models_lottery

type Prize struct {
	ID             int    `json:"id" db:"id" form:"id"`
	Prizename string `json:"prizename" db:"prizename" form:"prizename"`
	Whatprize	string	`json:"whatprize" db:"whatprize" form:"whatprize"`
	Whatprizes []Lottery  `gorm:"FOREIGNKEY:Whatprize;ASSOCIATION_FOREIGNKEY:Whatprize"`

	Number int64 `json:"number" db:"number" form:"number"`

	Label          string `json:"label" db:"label" form:"label"`
}
