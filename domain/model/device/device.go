package device

import "time"

//Instance of device
type Device struct {
	ID           string
	MacAddress   string // typical device manufacturer provided ID
	Name         string // name for a device could be empty
	UserId       string // reference to user / owner of device
	DeviceTypeId string // a device type UUID
	LastName     string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

//constructor
func NewDevice(ID string, name string, usId string, devTyID string, creaAt time.Time, upAt time.Time) *Device {
	return &Device{
		ID:           ID,
		Name:         name,
		UserId:       usId,
		DeviceTypeId: devTyID,
		CreatedAt:    creaAt,
		UpdatedAt:    upAt,
	}
}
