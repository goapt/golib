package model

import (
	goxorm "github.com/go-xorm/xorm"
)

type IModel interface {
	DbName() string
	TableName() string
	PK() string
	Columns() []string
	DB() *goxorm.Engine

	//data format
	ToMapString() map[string]string
	ToMapInerface() map[string]interface{}
	ToStruct(data map[string]string) error

	//fast handle
	GetById(id interface{}) (bool, error)
	GetByWhere(query string, args ...interface{}) (bool, error)
	UpdateByMap(data map[string]interface{}, query interface{}, args ...interface{}) (int64, error)
	CreateByMap(data map[string]interface{}) (int64, error)

	//conditions
	Where(query string, args ...interface{}) IModel
	Order(order string) IModel
	Field(field string) IModel
	Limit(limit int) IModel
	Offset(offset int) IModel

	//operation
	Get() (bool, error)
	All(beans interface{}) error //  必须传递指针
	Update() (int64, error)
	Create() (int64, error)
	Delete() (int64, error)
	Session() (*goxorm.Session)
	GetSession() (*goxorm.Session)              // 事务中传递
	WithSession(session *goxorm.Session) IModel // 绑定session

	Reset() IModel
	GetField(name string) interface{}
	GetBuilder() *Builder
	QueryString() ([]map[string]string, error)
}