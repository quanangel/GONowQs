package mysql

// AuthGroup is auth_group table
type AuthGroup struct {
	// id
	ID int `gorm:"column:id;type:int(10);primaryKey;autoIncrement"`
	// name
	Name string `gorm:"column:name;type:varchar(20);not null"`
	// status: 1normal、2disable
	Status int8 `gorm:"column:status;type:tinyint(1);default:1"`
	// rule string
	Rules string `gorm:"column:rules;type:text;"`
	// add time
	AddTime int `gorm:"column:add_time;type:int(10);not null"`
	// update time
	UpdateTime int `gorm:"column:update_time;type:int(10);not null"`
}

// NewAuthGroup is return AuthGroup struct function
func NewAuthGroup() AuthGroup {
	return AuthGroup{}
}

// TODO: NOT FINISHS
