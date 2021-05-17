package mysql

import "time"

// Uploads is uploads table struct
type Uploads struct {
	// id
	ID int64 `gorm:"column:id;type:bigint(20) auto_increment;primaryKey;comment:id"`
	// ClassifyName
	ClassifyName `gorm:"column:classify_name;type:varchar(20);not null;index:classify_name;comment:classify name"`
	// file name
	FileName string `gorm:"column:file_name;type:varchar(255);not null;comment:file name"`
	// file type
	FileType string `gorm:"column:file_type;type:varchar(20);not null;comment:file type"`
	// file address
	FileAdd string `gorm:"column:file_add;type:text;not null;comment:file address"`
	// file md5
	FileMd5 string `gorm:"column:file_md5;type:varchar(100);index;not null;comment:file md5"`
	// upload user id
	UploadUserID int64 `gorm:"column:upload_user_id:bigint(20);index;default:0;comment:upload user id"`
	// file status: 0delete、1normal、2show for upload user
	Status int8 `gorm:"column:status:tinyint(1);default:1;comment:file status:0delete、1normal、2only show for upload user"`
	// add time
	AddTime int `gorm:"column:add_time;type:int(10);not null;comment:add time"`
}

// NewImages is return Images struct
func NewImages() Uploads {
	return Uploads{}
}

// Add is add action function
func (m *Uploads) Add(classifyName string, fileName string, fileType string, fileAdd string, fileMd5 string, uploadUserID int64, status int8) int64 {
	m.ClassifyName = classifyName
	m.FileName = fileName
	m.FileType = fileType
	m.FileAdd = fileAdd
	m.FileMd5 = fileMd5
	m.UploadUserID = uploadUserID
	m.Status = status
	m.AddTime = int(time.Now().Unix())
	db := GetDb()
	result := db.Create(m)
	if result.RowsAffected > 0 {
		return m.ID
	}
	return 0
}

// GetOne is get one message by FileMd5
func (m *Uploads) GetOne(fileMd5 string) *Uploads {
	db := GetDb()
	db.Where("file_md5 = ?", fileMd5).First(m)
	return m
}
