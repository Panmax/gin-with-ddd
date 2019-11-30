package facade

import (
	"github.com/Panmax/gin-template/application/app_service"
	"github.com/Panmax/gin-template/infrastructure/util/bind"
	"github.com/Panmax/gin-template/infrastructure/util/response"
	"github.com/Panmax/gin-template/interfaces/assembler"
	"github.com/Panmax/gin-template/interfaces/dto"
	"github.com/gin-gonic/gin"
	"strconv"
)

func LeaveRouterRegister(router *gin.RouterGroup) {
	router.POST("", CreateLeaveInfo)
	router.GET("", GetLeaveInfo)
}

func CreateLeaveInfo(c *gin.Context) {
	utilGin := response.Gin{Ctx: c}

	s, err := bind.Bind(&dto.LeaveAdd{}, c)

	if err != nil {
		utilGin.ErrorResponse(err.Error())
		return
	}

	leaveAdd := s.(*dto.LeaveAdd)
	id := app_service.CreateLeaveInfo(assembler.LeaveAddConvertTo(leaveAdd))
	utilGin.SuccessResponse(id)
}

func GetLeaveInfo(c *gin.Context) {
	utilGin := response.Gin{Ctx: c}

	id, _ := strconv.Atoi(c.Query("id"))

	leave := app_service.GetLeaveInfo(uint(id))
	utilGin.SuccessResponse(assembler.LeaveDTOConvertFrom(leave))
}
