package upper

import (
	"github.com/verystar/golib/db"
	"database/sql"
	"upper.io/db.v3/lib/sqlbuilder"
	upperdb "upper.io/db.v3"
)

var _ db.IBuilder = (*Builder)(nil)

type Builder struct {
	db         sqlbuilder.Database
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
	id, err := b.collection.Insert(i)
	if err != nil {
		return 0, err
	}
	return id.(int64), nil
}

func (b *Builder) Update(i interface{}) (int64, error) {
	err := b.where.Update(i)
	return 0, err
}

func (b *Builder) Delete() (int64, error) {
	err := b.where.Delete()
	return 0, err
}
