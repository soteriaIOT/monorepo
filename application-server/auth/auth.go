package auth

import (
	"fmt"
	"log"
	"os"
	"time"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"

	"github.com/arora-aditya/monorepo/application-server/graph/model"
)

const USER_TABLE = "login-information"

type AWSUser struct {
	Username string `json:"username"`
	Name     string `json:"name"`
	Password string `json:"password"`
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
func (c *DynamoClient) VerifyByUsernameAndPassword(username string, password string) (*model.Token, error){
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
		return nil, fmt.Errorf("User not found")
	}

	item := AWSUser{}
	err = dynamodbattribute.UnmarshalMap(result.Item, &item)
	if err != nil {
		panic(fmt.Sprintf("Failed to unmarshal Record, %v", err))
	}

	if !CheckPasswordHash("thisisatestpassword", item.Password) {
		return nil, fmt.Errorf("Invalid password")
	}

	expiredAt := int(time.Now().Add(time.Hour * 1).Unix())

	return &model.Token{
		Token: GenerateJwt(item.Username, int64(expiredAt)),
		ExpiredAt: expiredAt,
	}, nil
}

func (c *DynamoClient) CreateUser(name string, username string, password string) (*model.Token, error) {
	hashedPassword, err := HashPassword(password)
	if err != nil {
		return nil, err
	}
	item := AWSUser{
		Username: username,
		Name: name,
		Password: hashedPassword,
	}

	av, err := dynamodbattribute.MarshalMap(item)
	if err != nil {
		panic(fmt.Sprintf("Failed to marshal Record, %v", err))
	}

	_, err = c.dynamoSvc.PutItem(&dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String(USER_TABLE),
	})
	if err != nil {
		panic(fmt.Sprintf("Failed to put item, %v", err))
	}

	expiredAt := int(time.Now().Add(time.Hour * 1).Unix())

	return &model.Token{
		Token: GenerateJwt(item.Username, int64(expiredAt)),
		ExpiredAt: expiredAt,
	}, nil
}

