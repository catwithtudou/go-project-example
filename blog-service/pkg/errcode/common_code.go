package errcode

/**
 *@Author tudou
 *@Date 2020/7/26
 **/

var(
	Success                   = NewError(0, "success")
	ServerError               = NewError(10000000, "the server error")
	InvalidParams             = NewError(10000001, "the param error")
	NotFound                  = NewError(10000002, "not found")
	UnauthorizedAuthNotExist  = NewError(10000003, "failed in authorization by AppKey and AppSecret error")
	UnauthorizedTokenError    = NewError(10000004, "failed in authorization by wrong token")
	UnauthorizedTokenTimeout  = NewError(10000005, "failed in authorization by token timeout")
	UnauthorizedTokenGenerate = NewError(10000006, "failed in authorization by token generating")
	TooManyRequests           = NewError(10000007, "too many requests")
)