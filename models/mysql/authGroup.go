package mysql

import (
	"strconv"
	"strings"

	"gorm.io/gorm"
)

// AuthGroup is auth_group table
type AuthGroup struct {
	// id
	ID int `gorm:"column:id;type:int(10);primaryKey;autoIncrement;comment:id"`
	// name
	Name string `gorm:"column:name;type:varchar(20);not null;comment:name"`
	// status: 1normal、2disable
	Status int8 `gorm:"column:status;type:tinyint(1);default:1;comment:status:1normal、2disable"`
	// rule string
	Rules string `gorm:"column:rules;type:text;comment:rule list"`
	// add time
	AddTime int `gorm:"column:add_time;type:int(10);not null;comment:add time"`
	// update time
	UpdateTime int `gorm:"column:update_time;type:int(10);not null;comment:update time"`
}

// NewAuthGroup is return AuthGroup struct function
func NewAuthGroup() AuthGroup {
	return AuthGroup{}
}

// Add is add the new group function
func (m *AuthGroup) Add(name string, status int8, rules string) int {
	m.Name = name
	m.Status = status
	m.Rules = rules
	db := GetDb()
	result := db.Create(m)
	if result.RowsAffected > 0 {
		return m.ID
	}
	return 0
}

// Edit is edit the group message function
func (m *AuthGroup) Edit(search map[string]interface{}, data map[string]interface{}) int64 {
	db := GetDb()
	result := db.Where(search).Updates(data)
	if result.RowsAffected > 0 {
		return result.RowsAffected
	}
	return 0
}

// Del is batch delete auth group and group access message function
func (m *AuthGroup) Del(search map[string]interface{}) bool {
	db := GetDb()
	resultErr := db.Transaction(func(tx *gorm.DB) error {
		authGroupList := make(map[int]AuthGroup)
		err := tx.Where(search).Find(authGroupList).Error
		if err != nil {
			return err
		}

		err = tx.Where(search).Delete(authGroupList).Error
		if err != nil {
			return err
		}

		authGroupId := make(map[int]int)
		for key, value := range authGroupList {
			authGroupId[key] = value.ID
		}

		err = tx.Where("group_id IN ?", authGroupId).Delete(AuthGroupAccess{}).Error
		if err != nil {
			return err
		}
		return nil
	})
	if resultErr != nil {
		return false
	}
	return true
}

// CheckUser is check the user is it have permission
func (m *AuthGroup) CheckUser(userID int64, url string, condition string) bool {
	db := GetDb()

	ruleIdMap := m.GetRules(userID)

	err := db.Where("id IN ? AND url = ? AND condition=? AND status = 1", ruleIdMap, url, condition).First(m).Error
	if nil != err {
		return false
	}

	return true
}

// GetRules is get rules map
func (m *AuthGroup) GetRules(userID int64) (rules map[int]int) {
	db := GetDb()
	accessKV := make(map[int]AuthGroupAccess)
	err := db.Where("user_id = ?", userID).Find(accessKV).Error
	if nil != err {
		return rules
	}

	groupIdMap := make(map[int]int)
	for key, value := range accessKV {
		groupIdMap[key] = value.GroupID
	}

	groupKV := make(map[int]AuthGroup)
	err = db.Where("id IN ?", groupIdMap).Find(groupKV).Error

	for _, value := range groupKV {
		if len(value.Rules) > 0 {
			tmp := strings.Split(value.Rules, ",")
			for _, v := range tmp {
				vInt, _ := strconv.Atoi(v)
				rules[vInt] = vInt
			}
		}
	}
	return rules
}
