package register

import (
	"DeviceService/endpoint/rest/api/device"
	"DeviceService/endpoint/rest/response"
	"DeviceService/registry"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"log"
	"net/http"
	"os"
)

// a REST http endpoint for handling device register
var errorLogger = log.New(os.Stderr, "ERROR ", log.Llongfile)

func registerDevice(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	fmt.Printf("Invoking regsiter Device Function...")
	api, err := createDeviceModelAPIInstance()
	if err != nil {
		return serverError(err)
	}
	device := new(response.Device)
	err2 := json.Unmarshal([]byte(req.Body), device)
	if err2 != nil {
		return clientError(http.StatusUnprocessableEntity)
	}

	if device.ModelName == "" || device.DeviceType == "" || device.DeviceTypeId == "" || device.UserId == "" {
		log.Fatalf("Some mandatory fields are missing in device payload.")
		return clientError(http.StatusBadRequest)
	}
	deviceExists, err := api.Service.GetByMacId(device.MacAddress)
	if err != nil {
		return serverError(err)
	}
	if deviceExists != nil {
		log.Fatalf("Device with macId %s already exists.", device.MacAddress)
		return clientError(http.StatusConflict)
	}

	devices, err := api.RegisterDevice(device)
	if err != nil {
		//handle error
		return serverError(err)
	}

	response, err := json.Marshal(devices)
	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}
	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
		Body: string(response),
	}, nil
}

// create device  API instance
func createDeviceModelAPIInstance() (*device.APIInstance, error) {
	ctn, err := registry.NewContainer()
	if err != nil {
		return nil, err
	}
	api := ctn.Resolve("device-model-API").(*device.APIInstance)
	return api, nil
}

// Helper to convert server error into APIGateway error response
func serverError(err error) (events.APIGatewayProxyResponse, error) {
	errorLogger.Println(err.Error())

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusInternalServerError,
		Body:       http.StatusText(http.StatusInternalServerError),
	}, nil
}

func clientError(status int) (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{
		StatusCode: status,
		Body:       http.StatusText(status),
	}, nil
}

func main() {
	lambda.Start(registerDevice)
}
