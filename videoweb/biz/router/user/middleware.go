// Code generated by hertz generator.

package user

import (
	"github.com/cloudwego/hertz/pkg/app"
	"videoweb/middleware"
)

func rootMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _infoMw() []app.HandlerFunc {
	// your code...
	return []app.HandlerFunc{middleware.AuthMiddleware()}
}

func _loginMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _registerMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _avatarMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _avataruploadMw() []app.HandlerFunc {
	// your code...
	return []app.HandlerFunc{middleware.AuthMiddleware()}
}

func _userMw() []app.HandlerFunc {
	// your code...
	return nil
}