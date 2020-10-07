package service

import (
	"DeviceService/domain/model/device"
	"fmt"
	"log"
)

type DeviceService struct {
	repo device.Repository
}

//Constructor for DeviceModelService
func NewDeviceService(repo device.Repository) *DeviceService {
	//return &DeviceModelService{repo: repo}
	dSer := new(DeviceService)
	dSer.repo = repo
	return dSer
}

// Registers a device first time
func (ds *DeviceService) RegisterDevice(d *device.Device) (error, string) {
	dExists, errModelExists := ds.repo.GetByMacID(d.MacAddress)
	// if model does not exist create the model and save
	if dExists != nil {
		log.Fatalf(" Device already with macID %s exists", d.MacAddress)
		return fmt.Errorf(" Device already with macID %s exists", d.MacAddress), ""
	}
	if errModelExists != nil {
		return errModelExists, ""
	}

	// now save the new model
	dSave, errSave := ds.repo.RegisterDevice(d)

	if errSave != nil {
		log.Fatalf("Error saving device %s ", d.MacAddress)
		return fmt.Errorf("Error saving device %s ", d.MacAddress), ""
	}
	return nil, dSave
}

func (ds *DeviceService) GetByMacId(mId string) (*device.Device, error) {
	device, err := ds.repo.GetByMacID(mId)

	if err != nil {
		return nil, err
	}

	if device == nil {
		fmt.Printf(" device not found mac ID  %s \n", mId)
		return nil, fmt.Errorf("%s macId device not found", mId)
	}
	return device, nil
}

func (ds *DeviceService) GetAll() ([]*device.Device, error) {
	devices, err := ds.repo.GetAll()

	if err != nil {
		return nil, err
	}

	if devices == nil {
		fmt.Printf(" Get all error  \n")
		return nil, fmt.Errorf("Get all error \n")
	}
	return devices, nil
}
