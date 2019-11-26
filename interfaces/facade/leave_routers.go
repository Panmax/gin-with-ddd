package facade

import (
	"github.com/Panmax/gin-template/application/service"
	"github.com/Panmax/gin-template/interfaces/assembler"
	"github.com/Panmax/gin-template/interfaces/dto"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func LeaveRouterRegister(router *gin.RouterGroup) {
	router.POST("", CreateLeaveInfo)
	router.GET("", GetLeaveInfo)
}

func CreateLeaveInfo(c *gin.Context) {
	leaveDTO := dto.LeaveDTO{}
	err := c.ShouldBind(&leaveDTO)
	if err != nil {
		return
	}
	id := service.CreateLeaveInfo(assembler.LeaveDTOConvertToLeave(leaveDTO))
	c.JSON(http.StatusCreated, gin.H{"message": "success", "data": id})
}

func GetLeaveInfo(c *gin.Context) {
	id, _ := strconv.Atoi(c.Query("id"))
	leaveInfo := service.GetLeaveInfo(uint(id))
	c.JSON(http.StatusOK, gin.H{"message": "success", "data": leaveInfo})
}
