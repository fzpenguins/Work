// Code generated by hertz generator.

package comment

import (
	"github.com/cloudwego/hertz/pkg/app"
	"videoweb/middleware"
)

func rootMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _commentMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _deleteMw() []app.HandlerFunc {
	// your code...
	return []app.HandlerFunc{middleware.AuthMiddleware()}
}

func _listMw() []app.HandlerFunc {
	// your code...
	return []app.HandlerFunc{middleware.AuthMiddleware()}
}

func _publishMw() []app.HandlerFunc {
	// your code...
	return []app.HandlerFunc{middleware.AuthMiddleware()}
}
