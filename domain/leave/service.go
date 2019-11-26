package leave

func CreateLeaveInfo(l *Leave) uint {
	return CreateLeave(l)
}

func GetLeaveInfo(id uint) *Leave {
	return GetLeave(id)
}
