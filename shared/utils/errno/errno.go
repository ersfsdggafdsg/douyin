package errno

type ErrCode int32
const (
	SuccessCode                ErrCode = 0     // Success
	ServiceErrCode             ErrCode = 10001 // Service not avaliable
	ParamErrCode               ErrCode = 10002 // Wrong parameter
	AuthorizationFailedErrCode ErrCode = 10003 // Authorization failed
	UserAlreadyExistErrCode    ErrCode = 10004 // User existed
	UserNotExistErrCode        ErrCode = 10005 // User not exists
	InvalidActionTypeErrCode   ErrCode = 10006 // Invalid action type
	VideoNotFoundErrCode       ErrCode = 10007 // Video not found
	RecodeNotFoundCode         ErrCode = 10008 // Recode not found
	NotMotifiedCode            ErrCode = 10009 // Not motified
	WrongPasswordCode          ErrCode = 10010 // Wrong password
	UnknowUserNameCode         ErrCode = 10011 // Not registed
	BadRequestCode             ErrCode = 10012 // Not registed
	InvalidOperationCode       ErrCode = 10013 // Invalid Operation
	AlreadyLikedCode           ErrCode = 10014 // Already liked
	AlreadyFollowedCode        ErrCode = 10015 // Already followed
)

var (
	Success                = NewError(ErrCode(0    ))
	ServiceErr             = NewError(ErrCode(10001))
	ParamErr               = NewError(ErrCode(10002))
	AuthorizationFailedErr = NewError(ErrCode(10003))
	UserAlreadyExistErr    = NewError(ErrCode(10004))
	UserNotExistErr        = NewError(ErrCode(10005))
	InvalidActionTypeErr   = NewError(ErrCode(10006))
	VideoNotFoundErr       = NewError(ErrCode(10007))
	RecodeNotFound         = NewError(ErrCode(10008))
	NotMotified            = NewError(ErrCode(10009))
	WrongPassword          = NewError(ErrCode(10010))
	UnknowUserName         = NewError(ErrCode(10011))
	BadRequest             = NewError(ErrCode(10012))
	InvalidOperation       = NewError(ErrCode(10013))
	AlreadyLiked           = NewError(ErrCode(10014))
	AlreadyFollowed        = NewError(ErrCode(10015))
)
