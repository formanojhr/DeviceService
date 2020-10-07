package device

type API interface {
	Get(id string) (*Device, error)
	GetByMacID(id string) (*Device, error)
	GetAll() ([]*Device, error)
	Search(query string) ([]*Device, error)
	RegisterDevice(b *Device) (string, error)
	UpdateDevice(b *Device) (string, error)
	DeleteDevice(id string) error
}

type Repository interface {
	Get(id string) (*Device, error)
	GetByMacID(id string) (*Device, error)
	GetAll() ([]*Device, error)
	Search(query string) ([]*Device, error)
	RegisterDevice(b *Device) (string, error)
	UpdateDevice(b *Device) (string, error)
	DeleteDevice(id string) error
}
