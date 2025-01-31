package grpc

import (

	models "github.com/ivanbatutin921/Anti-bruteforce/mk/service/models"
	"gorm.io/gorm"
)

func (s *Server) CheckLogin(user *models.Auth) (*models.Auth, error) {
	var auth models.Auth
	s.db.Where("login = ?", user.Login).First(&auth)
	if auth.ID != 0 {
		return &auth, nil
	}
	return nil, nil // no error if user does not exist
}

func (s *Server) CheckIp(ip string) bool {
	err := s.db.Where("ip = ?", ip).First(&models.BlackList{}).Error
	if err == gorm.ErrRecordNotFound {
		s.logger.Infof("ip %s is not in blacklist", ip)
		return true
	}
	return false
}

func (s *Server) CreateUser(user *models.Auth) error {
	err := s.db.Create(user).Error
	if err != nil {
		return err
	}
	return nil
}

func (s *Server) DeleteBlackList(ip string) error {
	err := s.db.Where("ip = ?", ip).Delete(&models.BlackList{}).Error
	if err != nil {
		return err
	}
	return nil
}

func (s *Server) CreateBlackList(bl *models.BlackList) error {
	blackList := models.BlackList{
		Ip: bl.Ip,
	}
	err := s.db.Create(&blackList).Error
	if err != nil {
		return err
	}
	return nil
}

func (s *Server) CreateWhiteList(wl *models.WhiteList) error {
	whiteList := models.WhiteList{
		Ip: wl.Ip,
	}
	err := s.db.Create(&whiteList).Error
	if err != nil {
		return err
	}
	return nil
}

func (s *Server) DeleteWhiteList(ip string) error {
	err := s.db.Where("ip = ?", ip).Delete(&models.WhiteList{}).Error
	if err != nil {
		return err
	}
	return nil
}
