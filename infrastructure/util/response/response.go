package response

import "github.com/gin-gonic/gin"

type Gin struct {
	Ctx *gin.Context
}

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func (g *Gin) Response(code int, message string, data interface{}) {
	g.Ctx.JSON(200, Response{
		Code:    code,
		Message: message,
		Data:    data,
	})
	return
}

func (g *Gin) SuccessResponse(data interface{}) {
	g.Response(0, "success", data)
	return
}

func (g *Gin) ErrorResponse(message string) {
	g.Response(-1, message, nil)
	return
}
