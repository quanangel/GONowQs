package mysql

type AccountingBooksContent struct {
	// ID
	ID int64 `gorm:"column:id;type:bigint(20) auto_increment;primaryKey;comment:id"`
	// UserID
	UserID int64 `gorm:"column:user_id;type:bigint(20);default:0;index:user_id;comment:user id"`
	// ClassifyID
	ClassifyID int64 `gorm:"column:classify_id;type:bigint(20);default:0;index:classify_id;comment:classify id"`
	// Name
	Name string `gorm:"column:name;type:varchar(200);not null;comment:name"`
	// Images
	Images string `gorm:"column:images;type:text;comment:images commas to separate characters"`
	// Status 0deleted 1normal 2disable
	Status int8 `gorm:"column:status;type:tinyint(1);default(1);comment:status:0deleted/1normal/2disable"`
	// BaseTimeModel
	BaseTimeModel
}

// TODO:
