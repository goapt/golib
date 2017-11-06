package model

import (
	"github.com/verystar/golib/xorm"
	goxorm "github.com/go-xorm/xorm"
	"reflect"
	"encoding/json"
	"github.com/tidwall/gjson"
)

var _ IModel = (*BaseModel)(nil)

type BaseModel struct {
	Model              IModel
	dbName             string
	Builder            *Builder               `xorm:"-" json:"-"`
	TransactionSession *goxorm.Session        `xorm:"-" json:"-"`
}

//兼容性修改
func (this *BaseModel)GetModel() IModel {
	return this.Model
}

func (this *BaseModel) GetSession() *goxorm.Session {
	if this.TransactionSession != nil {
		return this.TransactionSession
	}
	return this.DB().Table(this.Model.TableName())
}

func (this *BaseModel) WithSession(session *goxorm.Session) IModel {
	this.TransactionSession = session
	return this
}

func (this *BaseModel) TableName() string {
	return this.Model.TableName()
}

func (this *BaseModel) PK() string {
	return this.Model.PK()
}

func (this *BaseModel) Columns() []string {
	s := reflect.TypeOf(this.Model).Elem() //通过反射获取type定义
	var col []string
	for i := 0; i < s.NumField(); i++ {
		val := s.Field(i).Tag.Get("json")

		if val != "-" {
			col = append(col, val)
		}
	}
	return col
}

func (this *BaseModel) ToMapString() map[string]string {
	str, _ := json.Marshal(this.Model)
	maps := gjson.ParseBytes(str).Map()
	m := make(map[string]string)

	for k, v := range maps {
		m[k] = v.String()
	}

	return m
}

func (this *BaseModel) ToMapInerface() map[string]interface{} {
	m := this.ToMapString()
	d := make(map[string]interface{})
	for k, v := range m {
		d[k] = v
	}

	return d
}

func (this *BaseModel) ToJSON() []byte {
	str, _ := json.Marshal(this.Model)
	return str
}

func (this *BaseModel) ToStruct(data map[string]string) error {
	return ScanStruct(data, this.Model)
}

func (this *BaseModel) GetField(name string) interface{} {
	s := reflect.TypeOf(this.Model).Elem()  //通过反射获取type定义
	v := reflect.ValueOf(this.Model).Elem() //通过反射出值

	for i := 0; i < s.NumField(); i++ {
		val := s.Field(i).Tag.Get("json")
		if val == name {
			return v.Field(i)
		}
	}
	return nil
}

func (this *BaseModel) DbName() string {
	return "default"
}

func (this *BaseModel) DB() *goxorm.Engine {
	db := this.Model.DbName()
	if db == "" {
		db = this.DbName()
	}
	return xorm.MustDB(db)
}

func (this *BaseModel) GetById(id interface{}) (bool, error) {
	return this.GetSession().ID(id).Get(this.Model)
}

func (this *BaseModel) GetByWhere(query string, args ...interface{}) (bool, error) {
	return this.Where(query, args...).Get()
}

func (this *BaseModel) Reset() IModel {
	c := this.Model
	p := reflect.ValueOf(&c).Elem()
	p.Set(reflect.Zero(p.Type()))
	//fmt.Printf("%+v",p)
	return this
}

func (this *BaseModel) Create() (int64, error) {
	return this.DB().Insert(this.Model)
}

func (this *BaseModel) CreateByMap(data map[string]interface{}) (int64, error) {
	return CreateByMap(this.Model, data)
}

func (this *BaseModel) Where(query string, args ...interface{}) IModel {
	if this.Builder == nil {
		this.Builder = &Builder{}
	}
	this.Builder.Statement = query
	this.Builder.Param = args
	return this.Model
}

func (this *BaseModel) Order(order string) IModel {
	if this.Builder == nil {
		this.Builder = &Builder{}
	}
	this.Builder.Order = order
	return this
}

func (this *BaseModel) Field(field string) IModel {
	if this.Builder == nil {
		this.Builder = &Builder{}
	}
	this.Builder.Field = field
	return this
}

func (this *BaseModel) Limit(limit int) IModel {
	if this.Builder == nil {
		this.Builder = &Builder{}
	}
	this.Builder.Limit = limit
	return this
}

func (this *BaseModel) Offset(offset int) IModel {
	if this.Builder == nil {
		this.Builder = &Builder{}
	}
	this.Builder.Offset = offset
	return this
}

func (this *BaseModel) GetBuilder() *Builder {
	return this.Builder
}

func (this *BaseModel) Session() (*goxorm.Session) {
	session := this.DB().Table(this.Model.TableName())
	if this.Builder != nil {
		this.Builder.Patch(session)
	}
	return session
}

func (this *BaseModel) Get() (bool, error) {
	session := this.GetSession()
	if this.Builder != nil {
		this.Builder.Patch(session)
	}
	return session.NoAutoCondition().Get(this.Model)
}

func (this *BaseModel) All(beans interface{}) error {
	session := this.GetSession()
	if this.Builder != nil {
		this.Builder.Patch(session)
	}
	return session.Find(beans)
}

func (this *BaseModel) Delete() (int64, error) {
	session := this.GetSession()
	if this.Builder != nil {
		this.Builder.Patch(session)
	}
	return session.Delete(this.Model)
}

func (this *BaseModel) QueryString() ([]map[string]string, error) {
	if this.Builder == nil {
		this.Builder = &Builder{}
	}
	this.Builder.Table = this.Model.TableName()
	if this.Builder.Session == nil {
		this.Builder.Session = this.GetSession()
	}
	return this.GetBuilder().Query()
}

func (this *BaseModel) QueryBySql(sql string, args ... interface{}) ([]map[string]string, error) {
	data, err := this.GetSession().QueryString(sql, args...)
	return data, err
}

func (this *BaseModel) Update() (int64, error) {
	model := this.Model
	session := this.GetSession()
	if this.Builder != nil {
		this.Builder.Patch(session)
	}
	return session.Update(model)
}

func (this *BaseModel) UpdateByMap(data map[string]interface{}, query interface{}, args ...interface{}) (int64, error) {
	return this.GetSession().Where(query, args...).Update(data)
}