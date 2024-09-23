package services

import (
	db "github.com/ivanbatutin921/Anti-bruteforce/internal/database/postgresql"
	"github.com/ivanbatutin921/Anti-bruteforce/internal/models"
)

func CheckIp(ip string) bool {

	var blackList models.BlackList
	db.DB.Where("ip = ?", ip).First(&blackList)
	if blackList.ID != 0 {
		return false
	}
	
	return true
}
