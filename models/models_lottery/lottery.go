package models_lottery

type Lottery struct {
	ID             int    `json:"id" db:"id" form:"id"`
	Userid   int64  `gorm:"type:bigint;not null" json:"userid"`
	Username string `json:"username" db:"username" form:"username"`
	Status int64 `json:"status" db:"status" form:"status"`
	Whatprize	string	`json:"whatprize" db:"whatprize" form:"whatprize"`
	Label          string `json:"label" db:"label" form:"label"`
}

