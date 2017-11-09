package model

import (
	"github.com/verystar/golib/db"
	"github.com/verystar/golib/db/upper"
)

type Query struct {
	db.IBuilder
}

type IQuery interface {
	db.IBuilder
	GetByWhere(i interface{}, query string, arg interface{}, args ...interface{}) (bool, error)
	GetById(i interface{}, id interface{}) (bool, error)
}

func Builder(m db.IModel) db.IBuilder {
	return upper.NewBuilder(m.DbName()).Table(m.TableName())
}

func NewQuery(m db.IModel) IQuery {
	return &Query{
		IBuilder: Builder(m),
	}
}

func (q *Query) GetByWhere(i interface{}, query string, arg interface{}, args ...interface{}) (bool, error) {
	w := []interface{}{query, arg}
	w = append(w, args...)
	return q.Where(w...).Get(i)
}

func (q *Query) GetById(bean interface{}, id interface{}) (bool, error) {
	return q.Where(id).Get(bean)
}
