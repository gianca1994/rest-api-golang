package gateway

import (
	"rest-api-golang/gadgets/smartphones/models"
	"rest-api-golang/internal/database"
)

type SmartphoneCreateGateway interface {
	Create(cmd *models.CreateSmartphoneCMD) (*models.Smartphone, error)
}

type SmartphoneCreateGtw struct {
	SmartphoneStorageGateway
}

func NewSmartphoneCreateGateway(client *database.MySqlClient) SmartphoneCreateGateway {
	return &SmartphoneCreateGtw{&SmartphoneStorage{client}}
}
