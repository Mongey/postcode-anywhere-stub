package main

type InteractiveResponse struct {
	Table Table
}

type Table struct {
	Columns ColumnWrapper
	Rows    RowWrapper
}
type Column struct {
	Name string
}

type ColumnWrapper struct {
	Column []Column
}
type RowWrapper struct {
	Row interface{}
}
