package db

import (
	"database/sql"
)

type IModel interface {
	TableName() string
	DbName() string
}

type IBuilder interface {
	//表
	Table(t string) IBuilder
	//条件
	Where(w ...interface{}) IBuilder
	//分页
	Limit(i int) IBuilder
	//分页
	Offset(i int) IBuilder
	//排序
	OrderBy(s string) IBuilder
	//查询一条
	Get(i interface{}) (bool, error)
	//原生执行
	Exec(sql string,args ...interface{}) (sql.Result, error)
	//原生查询
	Query(i interface{}, sql string,args... interface{}) error
	//原生查询一条
	QueryRow(i interface{}, sql string,args... interface{}) error
	//查询多条
	All(i interface{}) error
	//统计
	Count() (uint64, error)
	//创建
	Create(i interface{}) (int64, error)
	//更新
	Update(i interface{},s ...[]string) (int64, error)
	//删除
	Delete() (int64, error)
	//设置事务上下文
	WithContext(i interface{}) IBuilder
	////事务开始
	//Begin() IBuilder
	////事务提交
	//Commit() IBuilder
	////事务回滚
	//Rollback() IBuilder
}