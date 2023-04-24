package gorepo

type CrudModel interface {
	TableName() string
	Index() []string
}
