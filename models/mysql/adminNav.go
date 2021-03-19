package mysql

// AdminNav is table admin_nav table
type AdminNav struct {
	// id
	ID int `gorm:"column:id;type:int(10);primaryKey;autoIncrement"`
	// name
	Name string `gorm:"column:name;type:varchar(10);not null;comment:name"`
	// url
	Url string `gorm:"column:name;type:varchar(100);comment:url"`
	// status
	Status int8 `gorm:"column:status;type:tinyint(1);default:1;comment:status:1normal、2disable"`
	// add time
	AddTime int `gorm:"column:add_time;type:int(10);not null;comment:add time"`
	// udpate time
	UpdateTime int `gorm:"column:update_time;type:int(10);not null;comment:update time"`
}

// NewAdminNav is return AdminNav struct
func NewAdminNav() AdminNav {
	return AdminNav{}
}

// TODO:not finish
