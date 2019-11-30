package app_service

import "github.com/Panmax/gin-template/domain/leave"

func CreateLeaveInfo(l *leave.Leave) uint {
	return leave.CreateLeaveInfo(l)
}

func GetLeaveInfo(id uint) *leave.Leave {
	return leave.GetLeaveInfo(id)
}
