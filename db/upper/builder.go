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
	b.where = b.collection.Find(w)
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
	autoTime(i, []string{"create_time", "created_at"})

	id, err := b.collection.Insert(i)
	if err != nil {
		return 0, err
	}
	return id.(int64), nil
}

func (b *Builder) Update(i interface{}) (int64, error) {
	autoTime(i, []string{"update_time", "updated_at"})
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

func autoTime(i interface{}, f []string) {
	switch i.(type) {
	case struct{}:
		fields := mapper.FieldsByName(reflect.ValueOf(i), f)
		for i := range fields {
			if fields[i].IsValid() {
				t := time.Now()
				fields[i].Set(reflect.ValueOf(t).Convert(fields[i].Type()))
			}
		}
	case map[string]interface{}:
		i := i.(map[string]interface{})
		for _, v := range f {
			_, ok := i[v]
			if !ok {
				i[v] = time.Now().Format("2006-01-02 15:04:05")
			}
		}
	}
}
