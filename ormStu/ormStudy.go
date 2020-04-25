package ormStu

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type User struct {
	Id   int64  `gorm:"primary_key;AUTO_INCREMENT;" json:"id"`
	Name string `gorm:"type:varchar(60);UNIQUE;NOT NULL;" json:"name"`
	Age  int    `gorm:"type:int(20);DEFAULT:18;" json:"age"`
}

var orm *gorm.DB

func init() {
	orm, _ = gorm.Open("mysql", "root:liangqinghai@tcp(127.0.0.1:3306)/test")
	orm.LogMode(true).Set("gorm:table_options", "ENGINE=InnoDB CHARSET=utf8mb4")
	orm.AutoMigrate(new(User))
}

func (User) TableName() string {
	return "user"
}

func (u *User) Insert(user *User) {
	orm.Table(u.TableName()).Create(&user)
}

// getAll
// @return User, error
func (u *User) GetAll() ([]User, error) {
	var users []User

	table := orm.Table(u.TableName())

	if err := table.Find(&users).Error; err != nil {
		return users, err
	}

	return users, nil
}
