package device

import (
	"DeviceService/domain/model/device"
	"DeviceService/domain/service"
	"DeviceService/endpoint/rest/response"
	"log"
)

// API Use cases for interactions with this microservice. This API will depend on the service layer
// and repositories
type API interface {
	Get(id string) (*response.Device, error)
	GetByMacId(id string) (*response.Device, error)
	GetAll() ([]*response.Device, error)
	Search(query string) ([]*response.Device, error)
	RegisterDevice(b *response.Device) (string, error)
	UpdateDevice(b *response.Device) (string, error)
	DeleteDevice(id string) error
}

type APIInstance struct {
	Service *service.DeviceService
}

func New(service *service.DeviceService) *APIInstance {
	return &APIInstance{Service: service}
}

func (s *APIInstance) RegisterDevice(d *response.Device) (string, error) {
	err, deviceId := s.Service.RegisterDevice(toModelDeviceType(d))
	if err != nil {
		return "", err
	}
	return deviceId, nil
}

func (s *APIInstance) GetByMacId(mId string) (*response.Device, error) {
	log.Printf("Get by mad if %s.", mId)
	device, err := s.Service.GetByMacId(mId)
	if err != nil {
		return nil, err
	}
	return toAPIResType(device), nil
}

func (s *APIInstance) GetAll() ([]*response.Device, error) {
	log.Printf("Device API GET ALL")
	devices, err := s.Service.GetAll()
	if err != nil {
		return nil, err
	}
	i := 0
	if len(devices) == 0 {
		log.Println("Device API GET ALL. No devices found for query")
	}
	// iterate the model domain type to rpc response type
	res := make([]*response.Device, len(devices))
	for _, device := range devices {
		res[i] = toAPIResType(device)
		i++
	}

	return res, nil
}

//convert from domain model type to rpc protocol serialization type
func toAPIResType(d *device.Device) *response.Device {
	res := response.Device{
		ID:           d.ID,
		Name:         d.Name,
		UserId:       d.UserId,
		DeviceTypeId: d.DeviceTypeId,
		DeviceType:   "", // TODO CALL OTHER API to fill this
		ModelName:    "", //TODO CALL OTHER API to fill this
		Vendor:       "", // Fill this as well
		CreatedAt:    d.CreatedAt,
		UpdatedAt:    d.UpdatedAt,
	}

	return &res
}

func toModelDeviceType(d *response.Device) *device.Device {
	res := device.Device{
		ID:           d.ID,
		Name:         d.Name,
		UserId:       d.UserId,
		DeviceTypeId: d.DeviceTypeId,
		CreatedAt:    d.CreatedAt,
		UpdatedAt:    d.UpdatedAt,
	}
	return &res
}
