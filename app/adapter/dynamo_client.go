package adapter

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/guregu/dynamo"
	"github.com/pkg/errors"
)

type DynamoClient struct {
	Client *dynamo.DB
	Config *aws.Config
}

func NewClient(config *aws.Config) *DynamoClient {
	return &DynamoClient{Config: config}
}

func (c *DynamoClient) Connect() (*dynamo.DB, error) {
	if c.Client == nil {
		sess, err := session.NewSession(c.Config)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		c.Client = dynamo.New(sess)
	}
	return c.Client, nil
}

func (c *DynamoClient) ConnectTable(tableName string) (*dynamo.Table, error) {
	db, err := c.Connect()
	if err != nil {
		return nil, errors.WithStack(err)
	}

	table := db.Table(tableName)

	return &table, nil
}

func (c *DynamoClient) CreateTableForTest(tableName string, table interface{}) error {
	db, err := c.Connect()
	if err != nil {
		return errors.WithStack(err)
	}

	err = db.CreateTable(tableName, table).Run()
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func (c *DynamoClient) DropTable(tableName string) error {
	db, err := c.Connect()
	if err != nil {
		return errors.WithStack(err)
	}

	_, err = db.Client().DeleteTable(&dynamodb.DeleteTableInput{
		TableName: aws.String(tableName),
	})
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}
