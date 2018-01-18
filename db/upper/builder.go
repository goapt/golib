package upper

import (
	"database/sql"
	"reflect"
	"time"

	upperdb "github.com/verystar/db"
	"github.com/verystar/db/lib/reflectx"
	"github.com/verystar/db/lib/sqlbuilder"
	"github.com/verystar/golib/cache"
	"github.com/verystar/golib/db"
)

var _ db.IBuilder = (*Builder)(nil)
var mapper = reflectx.NewMapper("db")

var AUTO_CREATE_TIME_FIELDS = []string{
	"create_time",
	"create_at",
	"created_at",
	"update_time",
	"update_at",
	"updated_at",
}

var AUTO_UPDATE_TIME_FIELDS = []string{
	"update_time",
	"update_at",
	"updated_at",
}

type UpperDatabase interface {
	upperdb.Database
	sqlbuilder.SQLBuilder
}

type Builder struct {
	db           UpperDatabase
	collection   upperdb.Collection
	where        upperdb.Result
	cacheColumns *cache.Cache
}

func NewBuilder(db ... string) *Builder {
	link_db := "default"
	if len(db) > 0 {
		link_db = db[0]
	}
	client := MustDB(link_db)
	return &Builder{
		db:           client.Link,
		cacheColumns: client.CachedColumns,
	}
}

func (b *Builder) Table(t string) db.IBuilder {
	b.collection = b.db.Collection(t)
	return b
}

func (b *Builder) Where(w ...interface{}) db.IBuilder {
	b.where = b.collection.Find(w...)
	return b
}

func (b *Builder) Limit(i int) db.IBuilder {
	b.where = b.where.Limit(i)
	return b
}

func (b *Builder) Offset(i int) db.IBuilder {
	b.where = b.where.Offset(i)
	return b
}

func (b *Builder) OrderBy(s string) db.IBuilder {
	b.where = b.where.OrderBy(s)
	return b
}

func (b *Builder) Get(i interface{}) (bool, error) {
	err := b.where.Limit(1).One(i)
	if err != nil {
		if err == upperdb.ErrNoMoreRows {
			return false, nil
		}
		return false, err
	}

	return true, nil
}

func (b *Builder) Exec(sql string, args ...interface{}) (sql.Result, error) {
	return b.db.Exec(sql, args...)
}

func (b *Builder) Query(i interface{}, sql string, param ... interface{}) error {
	iter := b.db.Iterator(sql, param...)
	if err := iter.Err(); err != nil {
		return err
	}
	defer iter.Close()
	return iter.All(i)
}

func (b *Builder) QueryRow(i interface{}, sql string, param ... interface{}) error {
	iter := b.db.Iterator(sql, param...)
	if err := iter.Err(); err != nil {
		return err
	}
	defer iter.Close()
	return iter.One(i)
}

func (b *Builder) All(i interface{}) error {
	return b.where.All(i)
}

func (b *Builder) Count() (uint64, error) {
	return b.where.Count()
}

func (b *Builder) Create(i interface{}) (int64, error) {

	itemV := reflect.ValueOf(i)
	itemV = reflect.Indirect(itemV)
	itemT := itemV.Type()

	switch itemT.Kind() {
	case reflect.Struct:
		fields := mapper.FieldMap(itemV)
		structAutoTime(fields, AUTO_CREATE_TIME_FIELDS)
	case reflect.Map:
		cols, err := b.Columns()
		if err == nil {
			i = mapAutoTime(i, cols, AUTO_CREATE_TIME_FIELDS)
		}
	}

	id, err := b.collection.Insert(i)
	if err != nil {
		return 0, err
	}
	return id.(int64), nil
}

func inSlice(k string, s []string) bool {
	for _, v := range s {
		if k == v {
			return true
		}
	}
	return false
}

func (b *Builder) Update(i interface{}, zeroValues ...[]string) (int64, error) {
	zv := make([]string, 0)

	if len(zeroValues) > 0 {
		zv = zeroValues[0]
	}

	itemV := reflect.ValueOf(i)
	itemV = reflect.Indirect(itemV)
	itemT := itemV.Type()

	switch itemT.Kind() {
	case reflect.Struct:
		fields := mapper.FieldMap(itemV)
		structAutoTime(fields, AUTO_UPDATE_TIME_FIELDS)
		i = zeroValueFilter(fields, zv)
	case reflect.Map:
		cols, err := b.Columns()
		if err == nil {
			i = mapAutoTime(i, cols, AUTO_UPDATE_TIME_FIELDS)
		}
	}

	err := b.where.Update(i)
	return 0, err
}

func (b *Builder) Delete() (int64, error) {
	err := b.where.Delete()
	return 0, err
}

func (b *Builder) WithContext(i interface{}) db.IBuilder {
	tx, _ := i.(sqlbuilder.Tx)
	b.db = tx
	return b
}

func (b *Builder) Columns() (clms []string, err error) {

	h := cache.String("columns_" + b.db.Name() + "_" + b.collection.Name())

	ccol, ok := b.cacheColumns.ReadRaw(h)
	if ok {
		return ccol.([]string), nil
	}

	q := b.db.Select("column_name").
		From("information_schema.columns").
		Where("table_schema = '" + b.db.Name() + "' AND table_name = '" + b.collection.Name() + "'")

	iter := q.Iterator()
	defer iter.Close()

	for iter.Next() {
		var columnName string
		if err := iter.Scan(&columnName); err != nil {
			return nil, err
		}
		clms = append(clms, columnName)
	}
	b.cacheColumns.Write(h, clms)

	return clms, nil
}

func zeroValueFilter(fields map[string]reflect.Value, zv []string) map[string]interface{} {
	m := make(map[string]interface{})

	for k, v := range fields {
		v = reflect.Indirect(v)
		if v.IsValid() && !inSlice(k, zv) {
			switch v.Interface().(type) {
			case int, int8, int16, int32, int64:
				c := v.Int()
				if c != 0 {
					m[k] = c
				}
			case uint, uint8, uint16, uint32, uint64:
				c := v.Uint()
				if c != 0 {
					m[k] = c
				}
			case float32, float64:
				c := v.Float()
				if c != 0.0 {
					m[k] = c
				}
			case bool:
				c := v.Bool()
				if c != false {
					m[k] = c
				}
			case string:
				c := v.String()
				if c != "" {
					m[k] = c
				}
			default:
				m[k] = v.Interface()
			}
		} else {
			m[k] = v.Interface()
		}
	}

	return m
}

func structAutoTime(fields map[string]reflect.Value, f []string) {
	for k, v := range fields {
		v = reflect.Indirect(v)
		if v.IsValid() && inSlice(k, f) && v.Type().Kind() == reflect.String {
			v.SetString(time.Now().Format("2006-01-02 15:04:05"))
		}
	}
}

func mapAutoTime(fields interface{}, cols []string, f []string) interface{} {

	switch ff := fields.(type) {
	case map[string]interface{}:
		for _, v := range cols {
			if inSlice(v, f) {
				ff[v] = time.Now().Format("2006-01-02 15:04:05")
			}
		}
		return ff
	case map[string]string:
		for _, v := range cols {
			if inSlice(v, f) {
				ff[v] = time.Now().Format("2006-01-02 15:04:05")
			}
		}
		return ff
	}
	return fields
}
