package storage

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/shuza/rest-and-go/model"
)

type SqliteConnectin struct {
	conn *gorm.DB
}

func (c *SqliteConnectin) Init() error {
	db, err := gorm.Open("sqlite3", "test.db")
	c.conn = db
	return err
}
func (c *SqliteConnectin) Save(model interface{}) {
	c.conn.AutoMigrate(&model)
	c.conn.Create(&model)
}
func (c *SqliteConnectin) FindUsernameAndPassword(username string, password string) model.User {
	user := model.User{}
	c.conn.Where("username = ? AND password = ?", username, password).First(&user)
	return user
}

func (c *SqliteConnectin) GetUserByName(name string) []model.User {
	users := make([]model.User, 0)
	c.conn.Where("username LIKE ?", name+"%").Find(&users)
	return users
}
