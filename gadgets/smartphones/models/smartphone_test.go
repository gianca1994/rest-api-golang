package models

import "testing"

func NewSmartphone(name string, price int, countryOrigin string, os string) *CreateSmartphoneCMD {
	return &CreateSmartphoneCMD{Name: name, Price: price, CountryOrigin: countryOrigin, OS: os}
}

func TestCreateSmartphoneValidateCorrect(t *testing.T) {
	r := NewSmartphone("Iphone", 100, "USA", "IOS")
	err := r.validate()
	if err != nil {
		t.Error(err)
		t.Fail()
	}
}

func TestCreateSmartphoneValidateWithoutName(t *testing.T) {
	r := NewSmartphone("", 100, "USA", "IOS")
	err := r.validate()
	if err == nil {
		t.Error(err)
		t.Fail()
	}
}

func TestCreateSmartphoneValidateWithoutPrice(t *testing.T) {
	r := NewSmartphone("Iphone", 0, "USA", "IOS")
	err := r.validate()
	if err == nil {
		t.Error(err)
		t.Fail()
	}
}

func TestCreateSmartphoneValidateWithoutCountryOrigin(t *testing.T) {
	r := NewSmartphone("Iphone", 100, "", "IOS")
	err := r.validate()
	if err == nil {
		t.Error(err)
		t.Fail()
	}
}

func TestCreateSmartphoneValidateWithoutOS(t *testing.T) {
	r := NewSmartphone("Iphone", 100, "USA", "")
	err := r.validate()
	if err == nil {
		t.Error(err)
		t.Fail()
	}
}