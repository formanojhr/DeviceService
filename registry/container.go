package registry

import (
	service2 "DeviceService/domain/service"
	"DeviceService/endpoint/rest/api/device"
	repo "DeviceService/repository/dynamodb"
	"github.com/sarulabs/di"
)

type DeviceContainer struct {
	ctn di.Container
}

//Constructor
func NewContainer() (*DeviceContainer, error) {
	builder, err := di.NewBuilder()

	if err != nil {
		return nil, err
	}

	if err := builder.Add([]di.Def{
		{
			Name:  "device-API",
			Build: buildDeviceAPI,
		},
	}...); err != nil {
		return nil, err
	}

	return &DeviceContainer{
		ctn: builder.Build(),
	}, nil
}

func (c *DeviceContainer) Resolve(name string) interface{} {
	return c.ctn.Get(name)
}

func (c *DeviceContainer) Clean() error {
	return c.ctn.Clean()
}

func buildDeviceAPI(ctn di.Container) (interface{}, error) {
	//repo := inmemory.NewDeviceModelRepository()
	repo := repo.New()
	service := service2.NewDeviceService(repo)
	return device.New(service), nil
}
