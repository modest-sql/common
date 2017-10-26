package common

//TableColumnDefiners is an array of TableColumnDefiner
type TableColumnDefiners []TableColumnDefiner

//TableColumnDefiner defines a column of the table to be created.
type TableColumnDefiner interface {
	ColumnName() string
	DefaultValue() interface{}
	Nullable() bool
	Autoincrementable() bool
}

type baseTableColumn struct {
	columnName        string
	defaultValue      interface{}
	nullable          bool
	autoincrementable bool
}

func (c baseTableColumn) ColumnName() string {
	return c.columnName
}

func (c baseTableColumn) DefaultValue() interface{} {
	return c.defaultValue
}

func (c baseTableColumn) Nullable() bool {
	return c.nullable
}

func (c baseTableColumn) Autoincrementable() bool {
	return c.autoincrementable
}

//IntegerTableColumn represents the definition of an integer table column.
type IntegerTableColumn struct {
	baseTableColumn
}

//NewIntegerTableColumn creates an instance of IntegerTableColumn.
func NewIntegerTableColumn(columnName string, defaultValue interface{}, nullable, autoincrementable bool) *IntegerTableColumn {
	return &IntegerTableColumn{baseTableColumn{columnName, defaultValue, nullable, autoincrementable}}
}

//BooleanTableColumn represents the definition of a boolean table column.
type BooleanTableColumn struct {
	baseTableColumn
}

//NewBooleanTableColumn creates an instance of BooleanTableColumn.
func NewBooleanTableColumn(columnName string, defaultValue interface{}, nullable, autoincrementable bool) *BooleanTableColumn {
	return &BooleanTableColumn{baseTableColumn{columnName, defaultValue, nullable, autoincrementable}}
}

//DatetimeTableColumn repesents the definition of a datetime table column.
type DatetimeTableColumn struct {
	baseTableColumn
}

//NewDatetimeTableColumn creates an instance of DatetimeTableColumn.
func NewDatetimeTableColumn(columnName string, defaultValue interface{}, nullable, autoincrementable bool) *DatetimeTableColumn {
	return &DatetimeTableColumn{baseTableColumn{columnName, defaultValue, nullable, autoincrementable}}
}

//CharTableColumn represents the definition of a char table column.
type CharTableColumn struct {
	baseTableColumn
	size int
}

//NewCharTableColumn creates an instance of CharTableColumn.
func NewCharTableColumn(columnName string, defaultValue interface{}, nullable, autoincrementable bool, size int) *CharTableColumn {
	return &CharTableColumn{baseTableColumn{columnName, defaultValue, nullable, autoincrementable}, size}
}

//Size returns the amount of bytes the char table column will consume.
func (c CharTableColumn) Size() int {
	return c.size
}

//CreateTableCommand represents a table creation statement.
type CreateTableCommand struct {
	tableName           string
	tableColumnDefiners TableColumnDefiners
}

//TableName returns the name of the table to be created.
func (c CreateTableCommand) TableName() string {
	return c.tableName
}

//TableColumnDefiners returns all column definitions of the table.
func (c CreateTableCommand) TableColumnDefiners() TableColumnDefiners {
	return c.tableColumnDefiners
}