package tools

// 当rpc调用错误的时候的响应
type BaseResp interface {
	SetStatusCode(int32)
	SetStatusMsg(string)
	GetStatusCode() int32
	GetStatusMsg() string
}

func BuildBaseResp(err error, errno int32, resp BaseResp) {
	resp.SetStatusCode(errno)
	resp.SetStatusMsg(err.Error())
	return
}

type InvalidActionTypeResp struct {
}
func (InvalidActionTypeResp)Error() string {
	return "Invalid action type"
}
func BuildInvalidActionTypeResp(errno int32, resp BaseResp, err error) {
	resp.SetStatusCode(errno)
	resp.SetStatusMsg("Invalid action type")
	err = InvalidActionTypeResp{}
}
