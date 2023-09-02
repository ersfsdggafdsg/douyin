package errno

import (
	"fmt"
	"log"
)

// 当rpc调用错误的时候的响应
type BaseResp interface {
	SetStatusCode(int32)
	SetStatusMsg(string)
	GetStatusCode() int32
	GetStatusMsg() string
}

func BuildBaseResp(errno ErrCode, resp BaseResp) {
	log.Output(2, fmt.Sprintf("\033[41m[debug]%s\033[0m", errno.String()))
	resp.SetStatusCode(int32(errno))
	resp.SetStatusMsg(errno.String())
	return
}

type Error struct {
	errno ErrCode
}

func (e Error)Error() string {
	return e.errno.String()
}

func NewError(errno ErrCode) error {
	fmt.Println("Create errno", errno)
	return Error{errno}
}
