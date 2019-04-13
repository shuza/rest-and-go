package storage

import "github.com/shuza/rest-and-go/model"

type IDbClient interface {
	Init() error
	Save(model interface{})
	FindUsernameAndPassword(username string, password string) model.User
	GetUserByName(name string) []model.User
}

var DbClient IDbClient
