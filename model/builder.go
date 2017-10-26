package model

import (
	"fmt"
	goxorm "github.com/go-xorm/xorm"
	"errors"
	"strings"
)

type Builder struct {
	Table     string
	Statement string
	Order     string
	Offset    int
	Limit     int
	Field     string
	Param     [] interface{}
	//DB *goxorm.Engine
	Session *goxorm.Session
}

func (this *Builder) Patch(session *goxorm.Session) {
	if this.Statement != "" {
		session.Where(this.Statement, this.Param...)
	}
	if this.Order != "" {
		session.OrderBy(this.Order)
	}

	offset := []int{this.Offset}
	session.Limit(this.Limit , offset...)

	if this.Field != "" {
		session.Cols( strings.Split(this.Field , ",")... )
	}
}

//Statement Order  Offset Limit field
func (this *Builder) QuerySql() string {
	field := "*"
	if this.Field != "" {
		field = this.Field
	}

	where := ""
	if this.Statement != "" {
		where = " WHERE " + this.Statement
	}

	order := ""
	if this.Order != "" {
		order = " ORDER BY " + this.Order
	}

	var offset int = 0
	if this.Offset != 0 {
		offset = this.Offset
	}

	var limit int = 0
	if this.Limit != 0 {
		limit = this.Limit
	}

	var limitStr = ""
	if limit > 0 {
		limitStr = fmt.Sprintf(" LIMIT  %v,  %v", offset, limit)
	}
	return fmt.Sprintf("SELECT %s FROM %s%s%s%s", field, this.Table, where, order, limitStr)
}

func (this *Builder) Query() ([]map[string]string, error) {
	data, err := this.Session.QueryString(this.QuerySql(), this.Param...)
	return data, err
}

func (this *Builder) One() (map[string]string, error) {

	this.Limit = 1
	data, err := this.Session.QueryString(this.QuerySql(), this.Param...)
	if err != nil {
		return nil, err
	}

	if len(data) == 0 {
		return nil, errors.New("RECORDER DOES NOT EXISTS")
	}
	return data[0], err
}

func CreateByMap(bean IModel, data map[string]interface{}) (int64, error) {
	//INSERT INTO `rpc_test` (`id`, `name`, `age`, `sex`) VALUES (NULL, '1', '2', '3');
	params := make([]interface{}, 0)
	fieldsArr := make([]string, 0)
	binds := "" // ? , ?
	for field, value := range data {
		fieldsArr = append(fieldsArr, field)
		params = append(params, value)
		binds += " ? ,"
	}
	binds = binds[0:len(binds)-1]
	//`id`, `name`, `age`, `sex`
	fields := "`" + strings.Join(fieldsArr, "`,`") + "`"
	sql := fmt.Sprintf("INSERT INTO `%v` (%v) VALUES (%v);", bean.TableName(), fields, binds)
	result, err := bean.GetSession().Exec(sql, params...)
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}
