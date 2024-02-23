package e

import "time"

//response的相关设置
const (
	SuccessCode = 10000
	SuccessMsg  = "success"

	FailureCode = -1
	FailureMsg  = "密码错误"
)

//设有accesstoken和refreshtoken的有效时间
const (
	AccessTokenExpireDuration  = time.Hour * 24
	RefreshTokenExpireDuration = time.Hour * 24 * 7
)
