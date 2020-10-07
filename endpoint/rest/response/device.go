package response

import "time"

type Device struct {
	ID           string    `json:"id"`
	MacAddress   string    `json:"mac_address"`
	Name         string    `json:"name"`         // name for a device could be empty
	UserId       string    `json:"userId"`       // reference to user outside
	DeviceTypeId string    `json:"deviceTypeId"` // a device type UUID
	DeviceType   string    `json:"deviceType"`
	ModelName    string    `json:"modelName"`
	Vendor       string    `json:"Vendor"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
