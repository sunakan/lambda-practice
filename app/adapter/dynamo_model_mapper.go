package adapter

import (
	"fmt"
	"github.com/guregu/dynamo"
	"time"
	"github.com/pkg/errors"
	"reflect"
)

type DynamoResource interface {
	EntityName() string
	PK() string
	SetPK()
	SK() string
	SetSK()
	SetID(id uint64)
	ID() uint64
	SetVersion(v int)
	Version() int
	CreatedAt() time.Time
	SetCreatedAt(t time.Time)
	UpdatedAt() time.Time
	SetUpdatedAt(t time.Time)
}

type DynamoModelMapper struct {
	Client    *ResourceTableOperator
	TableName string
	PKName    string
	SKName    string
}

func (d *DynamoModelMapper) GetEntityNameFromStruct(s interface{}) string {
	r := reflect.TypeOf(s)
	return r.Name()
}

func (d *DynamoModelMapper) GetEntityByID(id uint64, resource DynamoResource, ret interface{}) (interface{}, error) {
	table, err := d.Client.ConnectTable()
	if err != nil {
		return nil, errors.WithStack(err)
	}

	resource.SetID(id)
	err = table.
		Get(d.PKName, resource.PK()).
		Range(d.SKName, dynamo.Equal, resource.SK()).
		One(ret)

	if err != nil {
		return nil, errors.WithStack(err)
	}

	return ret, nil
}

func (d *DynamoModelMapper) GetPK(resource DynamoResource) string {
	return fmt.Sprintf("%s-%011d", resource.EntityName(), resource.ID())
}

func (d *DynamoModelMapper) GetSK(resource DynamoResource) string {
	return fmt.Sprintf("%011d", resource.ID())
}
