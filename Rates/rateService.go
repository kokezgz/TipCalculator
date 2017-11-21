package Rates

import (
	"../Services"
	"../Utils"
)

type IRateService interface {
	GetRates() ([]Rate, error)
	injectRates()
}

type RateService struct {
	client Services.IclientService
	backup Utils.IBackup
	config Utils.Config
}

func (r *RateService) GetRates() ([]Rate, error) {
	r.injectRates()
	r.backup.SetFileSection(r.config.Backup.Dir, r.config.Backup.File)

	response, err := r.client.CallService(r.config.Service.Route, r.config.Service.EndPoint, r.config.Service.ContentType)

	if err != nil {
		response, err = r.backup.ReadBackup()

		if err != nil {
			return nil, err
		}
	} else {
		err = r.backup.WriteBackup(response)
		if err != nil {
			println(err.Error())
		}
	}

	rates := parseRates(response)
	return rates, err
}

func (r *RateService) injectRates() {
	//Injects
	var inj Services.ClientService
	var inj2 Utils.Backup
	r.client = &inj
	r.backup = &inj2

	//New Config
	inj3 := Utils.NewSettings("Rates")
	r.config = inj3
}
