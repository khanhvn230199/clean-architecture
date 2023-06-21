package entities

type Entity interface {
	GetTableName() string
	GetNamesAndFields() [][]interface{}
}
