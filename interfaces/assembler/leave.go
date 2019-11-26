package assembler

import (
	"github.com/Panmax/gin-template/domain/leave"
	"github.com/Panmax/gin-template/interfaces/dto"
)

func LeaveDTOConvertToLeave(leaveDTO dto.LeaveDTO) *leave.Leave {
	return &leave.Leave{Name: leaveDTO.Name}
}
