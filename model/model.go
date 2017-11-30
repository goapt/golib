package model

import (
	"github.com/verystar/golib/db"
	"github.com/verystar/golib/db/upper"
	"fmt"
	"database/sql"
	"strings"
	"reflect"
)

type ModelRefect struct {
	Fields []string
}

var modelRefectCache  = make(map[string] *ModelRefect, 0)

type Query struct {
	db.IBuilder
	db.IModel
}

type IQuery interface {
	db.IBuilder
	GetByWhere(i interface{}, query string, arg interface{}, args ...interface{}) (bool, error)
	GetById(i interface{}, id interface{}) (bool, error)
	Replace(data map[string]interface{}) (sql.Result , error)
	BatchReplace(keys []string , data [][]interface{}) (sql.Result , error)
	BatchInsert(keys []string , data [][]interface{}) (sql.Result , error)
	GetModelRefect() *ModelRefect
	GetModel() db.IModel
}

func Builder(m db.IModel) db.IBuilder {
	return upper.NewBuilder(m.DbName()).Table(m.TableName())
}

func NewQuery(m db.IModel) IQuery {
	return &Query{
		IBuilder: Builder(m),
		IModel:m,
	}
}

func (q *Query ) GetModel() db.IModel {
	return q.IModel
}
func (q *Query)GetModelRefect() *ModelRefect {
	t := reflect.TypeOf(q.IModel).Elem()
	key := q.DbName() + ":" + q.TableName()
	ref , ok := modelRefectCache[key]
	if ok {
		return  ref
	}

	ref = &ModelRefect{}
	for i := 0 ; i < t.NumField() ; i++ {
		f := t.Field(i).Tag.Get("db")
		if f == "-" {
			continue
		}
		ref.Fields = append(ref.Fields , f)
	}
	modelRefectCache[key] = ref
	return ref
}

func (q *Query) GetByWhere(i interface{}, query string, arg interface{}, args ...interface{}) (bool, error) {
	w := []interface{}{query, arg}
	w = append(w, args...)
	return q.Where(w...).Get(i)
}

func (q *Query) GetById(bean interface{}, id interface{}) (bool, error) {
	return q.Where(id).Get(bean)
}

func (q *Query)Replace(data map[string]interface{}) (sql.Result , error) {
	keys := ""
	binds := ""
	vals := make([]interface{} , 0)
	if len(data) == 0 {
		panic("IQuery::Replace data can not empty")
	}
	for key, value := range data {
		keys += key + ","
		binds += "?" + ","
		vals = append(vals , value)
	}
	keys = keys[:len(keys) - 1]
	binds = binds[:len(binds) -1]

	msql := fmt.Sprintf("REPLACE INTO %s (%s) values (%s)", q.TableName(), keys , binds)
	return q.Exec(msql,vals...)
}

func (q *Query)BatchReplace(keys []string , data [][]interface{}) (sql.Result , error) {
	all_binds := ""
	vals := make([]interface{} , 0)
	if len(keys) == 0 {
		panic("IQuery::BatchReplace keys can not empty")
	}
	for _, item := range data {
		if len(keys) != len(item) {
			panic("len keys not eq item")
		}
		binds := ""
		for _, value := range item {
			binds += "?" + ","
			vals = append(vals , value)
		}
		binds = binds[:len(binds) -1]
		all_binds += "("+binds+"),"
	}
	all_binds = all_binds[:len(all_binds) -1]
	msql := fmt.Sprintf("REPLACE INTO %s (%s) values %s", q.TableName(), strings.Join(keys,",") , all_binds)
	return q.Exec(msql,vals...)
}

func (q *Query)BatchInsert(keys []string , data [][]interface{}) (sql.Result , error) {
	all_binds := ""
	vals := make([]interface{} , 0)
	if len(keys) == 0 {
		panic("IQuery::BatchInsert keys can not empty")
	}
	for _, item := range data {
		if len(keys) != len(item) {
			panic("len keys not eq item")
		}
		binds := ""
		for _, value := range item {
			binds += "?" + ","
			vals = append(vals , value)
		}
		binds = binds[:len(binds) -1]
		all_binds += "("+binds+"),"
	}
	all_binds = all_binds[:len(all_binds) -1]
	msql := fmt.Sprintf("INSERT INTO %s (%s) values %s", q.TableName(), strings.Join(keys,",") , all_binds)
	return q.Exec(msql,vals...)
}