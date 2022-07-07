package common

const (
	SUCCESS = 0
)

/**************************************
 * error code
 **************************************/

const (
	// 系统错误 10001 ～ 10100
	ErrSys        = 10001
	ErrSysReqData = 10002
	ErrSysDbExec  = 10003

	// 参数校验错误 10101 ～ 10200
	ErrValidReq = 10101
	ErrAuth     = 10205

	// 全局统一错误 10201 ～ 10300
	ErrInsert   = 10201
	ErrDelete   = 10202
	ErrUpdate   = 10203
	ErrNotExist = 10204

	// 三方平台错误 10401 ～ 10500
	ErrWx = 10401

	// 用户模块 20001 ～ 21000
	ErrUserRegister     = 20001
	ErrUserLogin        = 20002
	ErrUserLogout       = 20003
	ErrUserPassword     = 20004
	ErrUserTokenInvalid = 20005

	// 活动模块 30001 ~ 31000
	ErrActiveLotteryAward = 30001
	ErrActiveJoinCount = 30002
)

func GetMsg(code int) string {
	switch code {
	case SUCCESS:
		return "成功"

	case ErrSys:
		return "系统繁忙 请稍后重试"
	case ErrSysReqData:
		return "请求数据解析失败 请稍后重试"
	case ErrSysDbExec:
		return "数据错误"

	case ErrValidReq:
		return "参数错误"
	case ErrAuth:
		return "鉴权失败"

	case ErrInsert:
		return "添加失败"
	case ErrDelete:
		return "删除失败"
	case ErrUpdate:
		return "内容无变更"
	case ErrNotExist:
		return "不存在的记录"

	case ErrUserRegister:
		return "注册失败"
	case ErrUserLogin:
		return "登录失败"
	case ErrUserLogout:
		return "退出登录失败"
	case ErrUserPassword:
		return "密码错误"
	case ErrUserTokenInvalid:
		return "登录信息失效，请重新登录"

	case ErrActiveLotteryAward:
		return "活动奖品失效"
	case ErrActiveJoinCount:
		return "活动参与次数到达上限"
	default:
		return "未知错误"
	}
}

type RespErr struct {
	Code int
	Msg  string
}

func (err RespErr) Error() string {
	return err.Msg
}
func (err RespErr) NotNil() bool {
	return err.Code != 0
}

func NewRespErr(code int, msg string) RespErr {
	if msg == "" {
		msg = GetMsg(code)
	}
	return RespErr{
		Code: code,
		Msg:  msg,
	}
}

/**************************************
 * 其它全局 panic 错误类型
 **************************************/

// 系统错误
type SysErr struct {
	Msg string
}

func (err SysErr) Error() string {
	return err.Msg
}

func NewSysErr(err error) (sysErr SysErr) {
	sysErr = SysErr{Msg: err.Error()}
	return
}

// 校验错误
type ValidErr struct {
	Msg string
}

func (err ValidErr) Error() string {
	return err.Msg
}

func NewValidErr(err error) (validErr ValidErr) {
	validErr = ValidErr{Msg: err.Error()}
	return
}

// 数据库错误
type DbErr struct {
	Msg string
}

func (err DbErr) Error() string {
	return err.Msg
}

func NewDbErr(err error) (dbErr DbErr) {
	dbErr = DbErr{Msg: err.Error()}
	return
}
