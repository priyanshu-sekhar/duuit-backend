package request

type UserRequest struct {
	UID string `gorm:"primarykey"`
	Name string
	Bio string
}
