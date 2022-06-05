package errorx

const (
	SuccessCode    = iota
	ServiceErrCode = iota + 10000
	ParamErrCode
	LoginErrCode
	UserNotExistErrCode
	UserAlreadyExistErrCode
	AuthCheckTokenTimeOutErrCode
	AuthErrCode
	AuthCheckTokenErrCode
)

var (
	Success                  = NewError(SuccessCode, "Success")
	ServiceErr               = NewError(ServiceErrCode, "Service is unable to start successfully")
	ParamErr                 = NewError(ParamErrCode, "Wrong Parameter has been given")
	LoginErr                 = NewError(LoginErrCode, "Wrong username or password")
	UserNotExistErr          = NewError(UserNotExistErrCode, "User does not exists")
	UserAlreadyExistErr      = NewError(UserAlreadyExistErrCode, "User already exists")
	AuthCheckTokenTimeOutErr = NewError(AuthCheckTokenTimeOutErrCode, "token has expired")
	AuthErr                  = NewError(AuthErrCode, "wrong token")
	AuthCheckTokenErr        = NewError(AuthCheckTokenErrCode, "check token fail")
)
