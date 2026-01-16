package homework03

import (
	"errors"
	"fmt"
	"time"

	"github.com/jheader/go-homework1-template/util"
)

type User struct {
	ID        uint      `gorm:"primaryKey"`
	Name      string    `gorm:"size:64;not null"`
	Email     string    `gorm:"size:128;uniqueIndex;not null"`
	Age       uint8     `gorm:"not null"`
	Status    string    `gorm:"size:16;default:active;index"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

func (u *User) Insert(user User) error {
	//这里的user是拷贝，不会填充到外面，如果值指针的话，会填充到外面去
	db, err := util.NewUtilDB()
	if err != nil {
		return errors.New("连接数据库失败: " + err.Error())
	}

	if err := db.Create(&user).Error; err != nil {
		return errors.New("创建数据失败: " + err.Error())
	}
	fmt.Println(user)
	return nil
}

func (u *User) DeleteById(user User) error {

	db, err := util.NewUtilDB()
	if err != nil {
		return errors.New("连接数据库失败: " + err.Error())
	}
	db.Delete(&user)
	return nil
}

func (u *User) SelectPageByMap(page, size int, mp map[string]any) ([]User, error) {

	var users []User

	db, err := util.NewUtilDB()
	if err != nil {
		return nil, errors.New("连接数据库失败: " + err.Error())
	}
	db.Scopes(util.Paginate(page, size)).Where(mp).Find(&users)

	return users, nil

}
