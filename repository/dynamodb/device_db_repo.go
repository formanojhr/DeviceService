package dynamodb

import (
	"DeviceService/domain/model/device"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	"github.com/google/uuid"
	"log"
)

type DeviceDynamoRepo struct {
	Db dynamodbiface.DynamoDBAPI // a dynamo db client
}

const tname string = "Device"
const deviceIdIndex string = "DeviceId-index"

//constructor
func New() *DeviceDynamoRepo {
	log.Print("Instantiating deviceDynamoRepo.")
	// Initialize a session that the SDK will use to load
	// credentials from the shared credentials file ~/.aws/credentials
	// and region from the shared configuration file ~/.aws/config.
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	// Create DynamoDB client
	svc := dynamodb.New(sess)
	if svc != nil {
		log.Print("Initialized successfully DynamoDb client.")
	}
	// send back a new ref to DeviceModelDBRepository
	dmd := &DeviceDynamoRepo{
		Db: dynamodbiface.DynamoDBAPI(svc),
	}
	return dmd
}

func (r *DeviceDynamoRepo) Get(id string) (*device.Device, error) {
	log.Printf("Fetching all items...from %s table\n", tname)
	// fetch all items
	//params := &dynamodb.ScanInput{
	//	TableName: aws.String(tname),
	//}
	var queryInput = &dynamodb.QueryInput{
		TableName: aws.String(tname),
		IndexName: aws.String(deviceIdIndex),
		KeyConditions: map[string]*dynamodb.Condition{
			"ID": {
				ComparisonOperator: aws.String("EQ"),
				AttributeValueList: []*dynamodb.AttributeValue{
					{
						S: aws.String(id),
					},
				},
			},
		},
	}
	var resp, err = r.Db.Query(queryInput)
	if err != nil {
		fmt.Errorf("failed to make Query API call, %v", err)
		return nil, err
	} else {
		dev := new(device.Device)
		err = dynamodbattribute.UnmarshalListOfMaps(resp.Items, &dev)
		if err != nil {
			fmt.Errorf("failed to unmarshall device, %v", err)
			return nil, err
		} else {
			return dev, nil
		}
	}
}

func (r *DeviceDynamoRepo) GetByMacID(id string) (*device.Device, error) {
	panic("implement me")
}

func (r *DeviceDynamoRepo) Search(query string) ([]*device.Device, error) {
	panic("implement me")
}

func (r *DeviceDynamoRepo) UpdateDevice(b *device.Device) (string, error) {
	panic("implement me")
}

func (r *DeviceDynamoRepo) DeleteDevice(id string) error {
	panic("implement me")
}

func (r *DeviceDynamoRepo) GetAll() ([]*device.Device, error) {
	log.Printf("Fetching all items...from %s table\n", tname)

	// fetch all items
	params := &dynamodb.ScanInput{
		TableName: aws.String(tname),
	}
	allItems, err := r.Db.Scan(params)
	if err != nil {
		log.Fatalf("failed to make Query API call, %v", err)
		return nil, err
	}

	var devs []device.Device

	err = dynamodbattribute.UnmarshalListOfMaps(allItems.Items, &devs)

	if err != nil {
		fmt.Errorf("failed to unmarshal Query result items, %v", err)
		return nil, err
	}

	devices := make([]*device.Device, len(devs))
	i := 0
	for _, dev := range devs {
		devices[i] = device.NewDevice(dev.ID, dev.Name, dev.UserId, dev.DeviceTypeId,
			dev.CreatedAt, dev.UpdatedAt)
		i++
	}
	return devices, nil
}

func (r *DeviceDynamoRepo) RegisterDevice(d *device.Device) (string, error) {
	u := uuid.New()
	d.ID = u.String() // create and update UUID

	av, err := dynamodbattribute.MarshalMap(d)

	if err != nil {
		log.Println("Got error marshalling map:")
		log.Println(err.Error())
		return "", err
	}

	log.Printf("UUID %s generated for device  : %s\n", u, d.ID)
	log.Printf("Saving %s.. \n", d.Name)

	// Create item in table Device Model
	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String(tname),
	}
	_, err = r.Db.PutItem(input)
	if err != nil {
		log.Println("Got error calling DynamoDB PutItem:")
		log.Println(err.Error())
		return "", err
	}
	log.Printf("Success saving device m %s ID %s \n", d.Name, d.ID)

	return d.ID, nil
}
