package gateway

import "rest-api-golang/gadgets/smartphones/models"

func (s *SmartphoneCreateGtw) Create(cmd *models.CreateSmartphoneCMD) (*models.Smartphone, error) {
	return s.create(cmd)
}
