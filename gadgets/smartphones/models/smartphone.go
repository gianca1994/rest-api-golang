package models

// Smartphone models a smartphones
type Smartphone struct {
	Id            int64
	Name          string
	Price         int
	CountryOrigin string
	OS            string
}

// CreateSmartphoneCMD create a new smartphone from CMD
type CreateSmartphoneCMD struct {
	Name          string `json:"name"`
	Price         int    `json:"price"`
	CountryOrigin string `json:"country_origin"`
	OS            string `json:"os"`
}
