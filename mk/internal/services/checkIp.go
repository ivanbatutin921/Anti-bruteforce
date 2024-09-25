package services

import (
	db "github.com/ivanbatutin921/Anti-bruteforce/mk/internal/database/postgresql"
)

func CheckIp(ip string) bool {
	bool, err := db.CheckIp(&db.PostgreSQLDB{}, ip)
	if err != nil {
		return bool
	}

	return bool
}
