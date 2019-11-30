package leave

import (
	"log"
	"strconv"
)

type Leave struct {
	Name string
}

func CreateLeave(l *Leave) uint {
	log.Println(l)
	log.Println(l.Name)
	return 10
}

func GetLeave(id uint) *Leave {
	return &Leave{Name: strconv.Itoa(int(id))}
}
