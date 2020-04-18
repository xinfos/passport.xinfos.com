package errs

const (
	ErrSuccess      = 200    //ErrSuccess  - 请求成功
	ErrInternal     = 500    //ErrInternal - 内部错误
	ErrDBQuery      = 1001   //ErrDBQuery  - DB 查询异常
	ErrParamInvalid = 100001 //ErrParamInvalid 参数提交错误
	ErrParamVerify  = 100002 //ErrParamVerify 参数校验错误
)

//ErrorMsg - customization error message
var ErrorMsg = map[int]string{
	ErrSuccess:      "请求成功",
	ErrParamInvalid: "抱歉，提交参数错误",
	ErrInternal:     "抱歉，服务内部错误",
	ErrDBQuery:      "抱歉，数据内部错误",
	ErrParamVerify:  "抱歉，提交参数不合法",
}
