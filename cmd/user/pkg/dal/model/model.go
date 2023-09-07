package model

import (
	"douyin/shared/rpc/kitex_gen/rpc"
)

type UserInfo struct {
	rpc.UserInfo
	Password string
}

func (UserInfo)TableName() string {
	return "users"
}

func (UserInfo) PrimaryKey() string {
  return "user_id"
}
