package auth

import (
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

const USER_TABLE = "login-information"

type User struct {
	Username string
	Name     string
	Password string
}

type DynamoClient struct {
	dynamoSvc *dynamodb.DynamoDB
}

// NewDynamoSvc creates a new dynamo client
func NewDynamoSvc() *DynamoClient {
	// Initialize a session that the SDK will use to load
	// credentials from the shared credentials file ~/.aws/credentials
	// and region from the shared configuration file ~/.aws/config.
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		Config: aws.Config{
			Credentials: credentials.NewStaticCredentials(os.Getenv("AWS_ACCESS_KEY_ID"), os.Getenv("AWS_SECRET_ACCESS_KEY"), ""),
			Region:      aws.String("us-east-2"),
		},
		SharedConfigState: session.SharedConfigEnable,
	}))
	return &DynamoClient{dynamodb.New(sess)}
}

// VerifyByUsernameAndPassword determines if a user with a specific login exists or not
func (c *DynamoClient) VerifyByUsernameAndPassword(username string, password string) bool {
	result, err := c.dynamoSvc.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String(USER_TABLE),
		Key: map[string]*dynamodb.AttributeValue{
			"username": {
				S: aws.String(username),
			},
		},
	})
	if err != nil {
		log.Fatalf("Got error calling GetItem: %s", err)
	}

	if result.Item == nil {
		return false
	}

	item := User{}
	err = dynamodbattribute.UnmarshalMap(result.Item, &item)
	if err != nil {
		panic(fmt.Sprintf("Failed to unmarshal Record, %v", err))
	}

	fmt.Println("Found item:")
	fmt.Println("User:  ", item.Name)
	fmt.Println("Username:  ", item.Username)
	fmt.Println("Password:", item.Password)
	return item.Password == password
}
