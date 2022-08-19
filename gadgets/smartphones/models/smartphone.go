package models

import "errors"

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

// validate update a smartphone from CMD
func (cmd *CreateSmartphoneCMD) validate() error {
	if len(cmd.Name) < 1 {
		return errors.New("name is required")
	}

	if cmd.Price < 1 {
		return errors.New("price must be greater than 0")
	}

	if len(cmd.CountryOrigin) < 1 {
		return errors.New("CountryOrigin is required")
	}

	if len(cmd.OS) < 1 {
		return errors.New("operative System is required")
	}

	return nil

}
