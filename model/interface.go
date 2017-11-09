package model

type IModel interface {
	DbName() string
	TableName() string
	PK() string
}