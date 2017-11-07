package upper

import (
	"database/sql"
	"reflect"
	"time"

	"github.com/verystar/golib/db"
	upperdb "upper.io/db.v3"
	"upper.io/db.v3/lib/reflectx"
	"upper.io/db.v3/lib/sqlbuilder"
)

var _ db.IBuilder = (*Builder)(nil)
var mapper = reflectx.NewMapper("db")

type UpperDatabase interface {
	upperdb.Database
	sqlbuilder.SQLBuilder
}

type Builder struct {
	db         UpperDatabase
	collection upperdb.Collection
	where      upperdb.Result
}

func NewBuilder(db ... string) *Builder {
	link_db := "default"
	if len(db) > 0 {
		link_db = db[0]
	}
	client := MustDB(link_db)
	return &Builder{
		db: client,
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
	b.where.Limit(i)
	return b
}

func (b *Builder) Offset(i int) db.IBuilder {
	b.where.Offset(i)
	return b
}

func (b *Builder) OrderBy(s string) db.IBuilder {
	b.where.OrderBy(s)
	return b
}

func (b *Builder) Get(i interface{}) (bool, error) {
	err := b.where.One(i)
	if err != nil {
		if err == upperdb.ErrNoMoreRows {
			return false, nil
		}
		return false, err
	}

	return true, nil
}

func (b *Builder) Exec(sql string, args ...interface{}) (sql.Result, error) {
	return b.db.Exec(sql, args)
}

func (b *Builder) Query(i interface{}, sql string, param ... interface{}) error {
	rows, err := b.db.Query(sql, param...)
	if err != nil {
		return err
	}

	iter := sqlbuilder.NewIterator(rows)
	return iter.All(i)
}

func (b *Builder) All(i interface{}) error {
	return b.where.All(i)
}

func (b *Builder) Count() (uint64, error) {
	return b.where.Count()
}

func (b *Builder) Create(i interface{}) (int64, error) {
	t := reflect.TypeOf(i).Elem().Kind()

	if t == reflect.Struct {
		fields := mapper.FieldMap(reflect.ValueOf(i))
		autoTime(fields, []string{"create_time", "create_at"})
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

	t := reflect.TypeOf(i).Elem().Kind()

	if t == reflect.Struct {
		fields := mapper.FieldMap(reflect.ValueOf(i))
		autoTime(fields, []string{"update_time", "update_at"})
		i = zeroValueFilter(fields, zv)
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

func autoTime(fields map[string]reflect.Value, f []string) {
	for k, v := range fields {
		v = reflect.Indirect(v)
		if v.IsValid() && inSlice(k, f) && v.Type().Kind() == reflect.String {
			v.SetString(time.Now().Format("2006-01-02 15:04:05"))
		}
	}
}
