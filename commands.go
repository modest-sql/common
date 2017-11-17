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
func NewIntegerTableColumn(columnName string, defaultValue interface{}, nullable, autoincrementable bool) IntegerTableColumn {
	return IntegerTableColumn{baseTableColumn{columnName, defaultValue, nullable, autoincrementable}}
}

//IntegerTableColumn represents the definition of an integer table column.
type FloatTableColumn struct {
	baseTableColumn
}

//NewIntegerTableColumn creates an instance of IntegerTableColumn.
func NewFloatTableColumn(columnName string, defaultValue interface{}, nullable, autoincrementable bool) FloatTableColumn {
	return FloatTableColumn{baseTableColumn{columnName, defaultValue, nullable, autoincrementable}}
}

//BooleanTableColumn represents the definition of a boolean table column.
type BooleanTableColumn struct {
	baseTableColumn
}

//NewBooleanTableColumn creates an instance of BooleanTableColumn.
func NewBooleanTableColumn(columnName string, defaultValue interface{}, nullable, autoincrementable bool) BooleanTableColumn {
	return BooleanTableColumn{baseTableColumn{columnName, defaultValue, nullable, autoincrementable}}
}

//DatetimeTableColumn repesents the definition of a datetime table column.
type DatetimeTableColumn struct {
	baseTableColumn
}

//NewDatetimeTableColumn creates an instance of DatetimeTableColumn.
func NewDatetimeTableColumn(columnName string, defaultValue interface{}, nullable, autoincrementable bool) DatetimeTableColumn {
	return DatetimeTableColumn{baseTableColumn{columnName, defaultValue, nullable, autoincrementable}}
}

//CharTableColumn represents the definition of a char table column.
type CharTableColumn struct {
	baseTableColumn
	size int
}

//NewCharTableColumn creates an instance of CharTableColumn.
func NewCharTableColumn(columnName string, defaultValue interface{}, nullable, autoincrementable bool, size int) CharTableColumn {
	return CharTableColumn{baseTableColumn{columnName, defaultValue, nullable, autoincrementable}, size}
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

//NewCreateTableCommand creates an instance of CreateTableCommand.
func NewCreateTableCommand(tableName string, tableColumnDefiners TableColumnDefiners) *CreateTableCommand {
	return &CreateTableCommand{tableName, tableColumnDefiners}
}

//TableName returns the name of the table to be created.
func (c CreateTableCommand) TableName() string {
	return c.tableName
}

//TableColumnDefiners returns all column definitions of the table.
func (c CreateTableCommand) TableColumnDefiners() TableColumnDefiners {
	return c.tableColumnDefiners
}

//TableColumnSelectors is an array of TableColumnSelector and TableColumnStarSelector.
type TableColumnSelectors []interface{}

//TableColumnSelector represents a selected column in a select query.
type TableColumnSelector struct {
	prefix     string
	columnName string
	alias      string
}

//NewTableColumnSelector creates an instance of a TableColumnSelector.
func NewTableColumnSelector(prefix, columnName, alias string) *TableColumnSelector {
	return &TableColumnSelector{prefix, columnName, alias}
}

//Prefix returns the column prefix and returns true if it isn't empty.
func (s TableColumnSelector) Prefix() (string, bool) {
	return s.prefix, len(s.prefix) > 0
}

//ColumnName returns the name of the selected column in a select query.
func (s TableColumnSelector) ColumnName() string {
	return s.columnName
}

//Alias returns the column alias and returns true if it isn't empty.
func (s TableColumnSelector) Alias() (string, bool) {
	return s.alias, len(s.alias) > 0
}

//TableColumnStarSelector represents a star selector in a select query.
type TableColumnStarSelector struct {
}

//NewTableColumnStarSelector creates an instance of TableColumnStarSelector.
func NewTableColumnStarSelector() *TableColumnStarSelector {
	return &TableColumnStarSelector{}
}

//SelectTableCommand represents a select from table query.
type SelectTableCommand struct {
	sourceTable          string
	tableColumnSelectors TableColumnSelectors
}

//NewSelectTableCommand returns an instance of SelectTableCommand.
func NewSelectTableCommand(sourceTable string) *SelectTableCommand {
	return &SelectTableCommand{sourceTable: sourceTable}
}

func (s SelectTableCommand) SourceTable() string {
	return s.sourceTable
}

//InsertCommand represents an insert statement.
type InsertCommand struct {
	tableName string
	values    map[string]interface{}
}

//NewInsertCommand returns an instance of an InsertCommand.
func NewInsertCommand(tableName string, values map[string]interface{}) *InsertCommand {
	return &InsertCommand{tableName, values}
}

//TableName returns the name of the table in which the values will be inserted.
func (i InsertCommand) TableName() string {
	return i.tableName
}

//Values returns a map in which the keys are the columns in which the values will be inserted.
func (i InsertCommand) Values() map[string]interface{} {
	return i.values
}
