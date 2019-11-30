package dto

import "github.com/Panmax/gin-template/domain/leave"

type LeaveAdd struct {
	Name string `form:"name" json:"name" binding:"required"`
}

func (l *LeaveAdd) ConvertTo() *leave.Leave {
	return &leave.Leave{Name: l.Name}
}

type LeaveDTO struct {
	Name string `json:"name"`
}

func (l *LeaveDTO) ConvertFrom(leave *leave.Leave) *LeaveDTO {
	return &LeaveDTO{Name: leave.Name}
}
