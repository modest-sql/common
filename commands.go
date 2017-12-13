package common

import "fmt"

type tableModifier interface {
	TableName() string
}

//TableColumnDefiners is an array of TableColumnDefiner
type TableColumnDefiners []TableColumnDefiner

//TableColumnDefiner defines a column of the table to be created.
type TableColumnDefiner interface {
	ColumnName() string
	DefaultValue() interface{}
	Nullable() bool
	Autoincrementable() bool
	PrimaryKey() bool
	ForeignKey() bool
}

type baseTableColumn struct {
	columnName        string
	defaultValue      interface{}
	nullable          bool
	autoincrementable bool
	primaryKey        bool
	foreignKey        bool
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

func (c baseTableColumn) PrimaryKey() bool {
	return c.primaryKey
}

func (c baseTableColumn) ForeignKey() bool {
	return c.foreignKey
}

//IntegerTableColumn represents the definition of an integer table column.
type IntegerTableColumn struct {
	baseTableColumn
}

//NewIntegerTableColumn creates an instance of IntegerTableColumn.
func NewIntegerTableColumn(columnName string, defaultValue interface{}, nullable, autoincrementable bool, primaryKey bool, foreignKey bool) IntegerTableColumn {
	return IntegerTableColumn{baseTableColumn{columnName, defaultValue, nullable, autoincrementable, primaryKey, foreignKey}}
}

//FloatTableColumn represents the definition of an float table column.
type FloatTableColumn struct {
	baseTableColumn
}

//NewFloatTableColumn creates an instance of FloatTableColumn.
func NewFloatTableColumn(columnName string, defaultValue interface{}, nullable, autoincrementable bool, primaryKey bool, foreignKey bool) FloatTableColumn {
	return FloatTableColumn{baseTableColumn{columnName, defaultValue, nullable, autoincrementable, primaryKey, foreignKey}}
}

//BooleanTableColumn represents the definition of a boolean table column.
type BooleanTableColumn struct {
	baseTableColumn
}

//NewBooleanTableColumn creates an instance of BooleanTableColumn.
func NewBooleanTableColumn(columnName string, defaultValue interface{}, nullable, autoincrementable bool, primaryKey bool, foreignKey bool) BooleanTableColumn {
	return BooleanTableColumn{baseTableColumn{columnName, defaultValue, nullable, autoincrementable, primaryKey, foreignKey}}
}

//DatetimeTableColumn repesents the definition of a datetime table column.
type DatetimeTableColumn struct {
	baseTableColumn
}

//NewDatetimeTableColumn creates an instance of DatetimeTableColumn.
func NewDatetimeTableColumn(columnName string, defaultValue interface{}, nullable, autoincrementable bool, primaryKey bool, foreignKey bool) DatetimeTableColumn {
	return DatetimeTableColumn{baseTableColumn{columnName, defaultValue, nullable, autoincrementable, primaryKey, foreignKey}}
}

//CharTableColumn represents the definition of a char table column.
type CharTableColumn struct {
	baseTableColumn
	size uint16
}

//NewCharTableColumn creates an instance of CharTableColumn.
func NewCharTableColumn(columnName string, defaultValue interface{}, nullable, autoincrementable bool, primaryKey bool, foreignKey bool, size uint16) CharTableColumn {
	return CharTableColumn{baseTableColumn{columnName, defaultValue, nullable, autoincrementable, primaryKey, foreignKey}, size}
}

//Size returns the amount of bytes the char table column will consume.
func (c CharTableColumn) Size() uint16 {
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
	isStar bool
	prefix     string
	columnName string
	alias      string
	function interface{}
}

//NewTableColumnSelector creates an instance of a TableColumnSelector.
func NewTableColumnSelector(isStar bool,prefix string, columnName string, alias string,function interface{}) *TableColumnSelector {
	return &TableColumnSelector{isStar,prefix, columnName, alias,function}
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


type GroupBySelect struct{
	table string
	column  string
}

//NewTableColumnSelector creates an instance of a TableColumnSelector.
func NewGroupBySelect(table string,column  string) *GroupBySelect {
	return &GroupBySelect{table ,column}
}
type JoinSelect struct {
	targetTable string
	targetAlias string
	filterCriteria interface{}
}

func NewJoinSelect(targetTable string,targetAlias  string,filterCriteria interface{}) *JoinSelect {
	return &JoinSelect{targetTable ,targetAlias  ,filterCriteria }
}
//SelectTableCommand represents a select from table query.
type SelectTableCommand struct {
	tableName            string
	mainAlias string
	tableColumnSelectors TableColumnSelectors
	joinList []JoinSelect
	whereExpression interface{}
	groupBy  []GroupBySelect
}

//NewSelectTableCommand returns an instance of SelectTableCommand.
func NewSelectTableCommand(tableName string,mainAlias string,tableColumnSelectors TableColumnSelectors,joinList []JoinSelect,whereExpression interface{},groupBy  []GroupBySelect) *SelectTableCommand {
	return &SelectTableCommand{tableName,mainAlias,tableColumnSelectors,joinList,whereExpression,groupBy}
}

//SourceTable returns the sourceTable of the table in which the values will be inserted.
func (s SelectTableCommand) TableName() string {
	return s.tableName
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

//DropCommand represents an drop statement.
type DropCommand struct {
	tableName string
}

//NewDropCommand returns an instance of an DropCommand
func NewDropCommand(tableName string) *DropCommand {
	return &DropCommand{tableName}
}

//TableName returns the name of the table in which the values will be inserted.
func (i DropCommand) TableName() string {
	return i.tableName
}

//AlterCommand represents an alter statement.
type AlterCommand struct {
	table       string
	instruction interface{}
}

//NewAlterCommand returns an instance of an AlterCommand
func NewAlterCommand(tableName string, instruction interface{}) *AlterCommand {
	return &AlterCommand{tableName, instruction}
}

//AlterDropInst represents an alter drop instruction.
type AlterDropInst struct {
	table string
}

//NewAlterDropInst returns an instance of an AlterDropInst
func NewAlterDropInst(tableName string) *AlterDropInst {
	return &AlterDropInst{tableName}
}

//AlterAddInst represents an alter add instruction.
type AlterAddInst struct {
	tableColumnDefiners TableColumnDefiner
}

//NewAlterAddInst returns an instance of an AlterAddInst
func NewAlterAddInst(tableColumnDefiners TableColumnDefiner) *AlterAddInst {
	return &AlterAddInst{tableColumnDefiners}
}

//AlterModifyInst represents an modify add instruction.
type AlterModifyInst struct {
	tableColumnDefiners TableColumnDefiner
}

//NewAlterModifyInst returns an instance of an AlterModifyInst
func NewAlterModifyInst(tableColumnDefiners TableColumnDefiner) *AlterModifyInst {
	return &AlterModifyInst{tableColumnDefiners}
}

type UpdateTableCommand struct {
	tableName string
}

func (c UpdateTableCommand) TableName() string {
	return c.tableName
}

type DeleteCommand struct {
	tableName string
}

func (c DeleteCommand) TableName() string {
	return c.tableName
}

//Instruction executes the command.
type Instruction func()

//InstructionType is used to determine the type of the instruction.
type InstructionType int

//Instruction type constants.
const (
	Create InstructionType = iota
	Select
	Update
	Insert
	Delete
	Drop
	Alter
)

var instructionName = map[InstructionType]string{
	Create: "CREATE",
	Select: "SELECT",
	Update: "UPDATE",
	Insert: "INSERT",
	Delete: "DELETE",
	Drop:   "DROP",
	Alter:  "ALTER",
}

func (i InstructionType) String() string {
	return instructionName[i]
}

//Command contains information about the instruction type and the instruction itself.
type Command struct {
	tableModifier
	InstructionType
	Instruction
}

func NewCommand(tableModifier tableModifier, instructionType InstructionType, instruction Instruction) Command {
	return Command{tableModifier, instructionType, instruction}
}

func (c Command) String() string {
	return fmt.Sprintf("%s %s", c.InstructionType.String(), c.tableModifier.TableName())

}

type IdCommon struct {
	name  string
	alias string
}

func NewIdCommon(tableName string, alias string) *IdCommon {
	return &IdCommon{tableName, alias}
}

type IntCommon struct {
	value int64
}

func NewIntCommon(value int64) *IntCommon {
	return &IntCommon{value}
}

type BoolCommon struct {
	value bool
}

func NewBoolCommon(value bool) *BoolCommon {
	return &BoolCommon{value}
}

type FloatCommon struct {
	value float64
}

func NewFloatCommon(value float64) *FloatCommon {
	return &FloatCommon{value}
}

type StringCommon struct {
	value string
}

func NewStringCommon(value string) *StringCommon {
	return &StringCommon{value}
}

type AssignmentCommon struct {
	value      string
	expression interface{}
}

func NewAssignmentCommon(value string, expression interface{}) *AssignmentCommon {
	return &AssignmentCommon{value, expression}
}

type SumCommon struct {
	rightValue interface{}
	leftValue  interface{}
}

func NewSumCommon(value interface{}, expression interface{}) *SumCommon {
	return &SumCommon{value, expression}
}

type SubCommon struct {
	rightValue interface{}
	leftValue  interface{}
}

func NewSubCommon(value interface{}, expression interface{}) *SubCommon {
	return &SubCommon{value, expression}
}

type MultCommon struct {
	rightValue interface{}
	leftValue  interface{}
}

func NewMultCommon(value interface{}, expression interface{}) *MultCommon {
	return &MultCommon{value, expression}
}

type DivCommon struct {
	rightValue interface{}
	leftValue  interface{}
}

func NewDivCommon(value interface{}, expression interface{}) *DivCommon {
	return &DivCommon{value, expression}
}

type EqCommon struct {
	rightValue interface{}
	leftValue  interface{}
}

func NewEqCommon(value interface{}, expression interface{}) *EqCommon {
	return &EqCommon{value, expression}
}

type NeCommon struct {
	rightValue interface{}
	leftValue  interface{}
}

func NewNeCommon(value interface{}, expression interface{}) *NeCommon {
	return &NeCommon{value, expression}
}

type LtCommon struct {
	rightValue interface{}
	leftValue  interface{}
}

func NewLtCommon(value interface{}, expression interface{}) *LtCommon {
	return &LtCommon{value, expression}
}

type GtCommon struct {
	rightValue interface{}
	leftValue  interface{}
}

func NewGtCommon(value interface{}, expression interface{}) *GtCommon {
	return &GtCommon{value, expression}
}

type LteCommon struct {
	rightValue interface{}
	leftValue  interface{}
}

func NewLteCommon(value interface{}, expression interface{}) *LteCommon {
	return &LteCommon{value, expression}
}

type GteCommon struct {
	rightValue interface{}
	leftValue  interface{}
}

func NewGteCommon(value interface{}, expression interface{}) *GteCommon {
	return &GteCommon{value, expression}
}

type BetweenCommon struct {
	rightValue interface{}
	leftValue  interface{}
}

func NewBetweenCommon(value interface{}, expression interface{}) *BetweenCommon {
	return &BetweenCommon{value, expression}
}

type LikeCommon struct {
	rightValue interface{}
	leftValue  interface{}
}

func NewLikeCommon(value interface{}, expression interface{}) *LikeCommon {
	return &LikeCommon{value, expression}
}

type NotCommon struct {
	not interface{}
}

func NewNotCommon(value interface{}) *NotCommon {
	return &NotCommon{value}
}

type AndCommon struct {
	rightValue interface{}
	leftValue  interface{}
}

func NewAndCommon(value interface{}, expression interface{}) *AndCommon {
	return &AndCommon{value, expression}
}

type OrCommon struct {
	rightValue interface{}
	leftValue  interface{}
}

func NewOrCommon(value interface{}, expression interface{}) *OrCommon {
	return &OrCommon{value, expression}
}

type NullCommon struct {
}

func NewNullCommon() *NullCommon {
	return &NullCommon{}
}

type FalseCommon struct {
}

func NewFalseCommon() *FalseCommon {
	return &FalseCommon{}
}

type TrueCommon struct {
}

func NewTrueCommon() *TrueCommon {
	return &TrueCommon{}
}
