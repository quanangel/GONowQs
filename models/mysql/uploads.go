package mysql

import "time"

// Uploads is uploads table struct
type Uploads struct {
	// id
	ID int64 `gorm:"column:id;type:bigint(20);primaryKey;autoIncrement"`
	// 文件名
	FileName string `gorm:"column:file_name;type:varchar(255);not null"`
	// 文件類型
	FileType string `gorm:"column:file_type;type:varchar(20);"`
	// 文件地址
	FileAdd string `gorm:"column:file_add;type:text;not null"`
	// 文件标识MD5
	FileMd5 string `gorm:"column:file_md5;type:varchar(100);index;not null"`
	// 上传文件的用户ID
	UploadUserID int64 `gorm:"column:upload_user_id:bigint(20);index;default:0"`
	// 文件状态：0已删除、1正常、2仅上传用户查看
	Status int8 `gorm:"column:status:tinyint(1);default:1"`
	// 添加时间
	AddTime int `gorm:"column:add_time;type:int(10);not null"`
}

// NewImages is return Images struct
func NewImages() Uploads {
	return Uploads{}
}

// Add is add action function
func (m *Uploads) Add(fileName string, fileType string, fileAdd string, fileMd5 string, uploadUserID int64, status int8) int64 {
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
