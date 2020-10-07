package main

import (
	"DeviceService/endpoint/rest/api/device"
	"DeviceService/registry"
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"log"
	"net/http"
	"os"
)

// a REST http endpoint for handling device get

var errorLogger = log.New(os.Stderr, "ERROR ", log.Llongfile)

// Handles /devices and /devices?queryParams,.,.,
func getDevices(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	log.Printf("Invoking get Device Function...")
	api, err := createDeviceModelAPIInstance()
	if err != nil {
		return serverError(err)
	}
	macId, findByMacId := req.QueryStringParameters["macId"]
	if findByMacId {
		devices, err := api.GetByMacId(macId)
		if err != nil {
			//handle error
			return serverError(err)
		}

		response, err := json.Marshal(devices)
		return events.APIGatewayProxyResponse{
			StatusCode: 200,
			Headers: map[string]string{
				"Content-Type": "application/json",
			},
			Body: string(response),
		}, nil
	} else {
		devices, err := api.GetAll()
		if err != nil {
			//handle error
			return serverError(err)
		}

		response, err := json.Marshal(devices)
		return events.APIGatewayProxyResponse{
			StatusCode: 200,
			Headers: map[string]string{
				"Content-Type": "application/json",
			},
			Body: string(response),
		}, nil
	}

	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
		Body: "",
	}, nil
}

// create device  API instance
func createDeviceModelAPIInstance() (*device.APIInstance, error) {
	ctn, err := registry.NewContainer()
	if err != nil {
		return nil, err
	}
	api := ctn.Resolve("device-API").(*device.APIInstance)
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
	lambda.Start(getDevices)
}
