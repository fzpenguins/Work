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

//分页操作相关设置
var (
	//每页显示的条数
	PageSize int

	//总共的条数
	PageNum int

	//当前页面
	CurrentPage int

	//上一页的最后一条记录
	Start int

	//当前页的最后一条记录
	End int
)
