package manager

import (
	"douyin/cmd/user/pkg/dal/mysql"
)

type Manager struct {
	Db mysql.UserManager
}
