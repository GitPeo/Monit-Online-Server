package models

type Todo struct {
	Id      int `gorm:"unique"`
	Tid     int `gorm:"unique"`
	Title   string
	Checked bool
	Order   float64
	Version int
	Deleted bool
}

func (Todo) TableName() string {
	return "todo"
}
