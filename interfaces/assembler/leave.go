package assembler

import (
	"github.com/Panmax/gin-template/domain/leave"
	"github.com/Panmax/gin-template/interfaces/dto"
)

func LeaveAddConvertTo(leaveAdd *dto.LeaveAdd) *leave.Leave {
	return &leave.Leave{Name: leaveAdd.Name}
}

func LeaveDTOConvertFrom(leave *leave.Leave) *dto.LeaveDTO {
	return &dto.LeaveDTO{Name: leave.Name}
}
