package common

import (
	"fmt"
	"reflect"
)

type Expression interface {
	Evaluate(symbols map[string]interface{}) interface{}
}

type IdCommon struct {
	name  string
	alias string
}

func NewIdCommon(tableName string, alias string) *IdCommon {
	return &IdCommon{tableName, alias}
}

func (id IdCommon) Evaluate(symbols map[string]interface{}) interface{} {
	if id.alias == "" {
		return symbols[id.name]
	}
	return symbols[fmt.Sprintf("%s.%s", id.name, id.alias)]
}

type IntCommon struct {
	value int64
}

func (i IntCommon) Evaluate(symbols map[string]interface{}) interface{} {
	return i.value
}

func NewIntCommon(value int64) *IntCommon {
	return &IntCommon{value}
}

type BoolCommon struct {
	value bool
}

func (b BoolCommon) Evaluate(symbols map[string]interface{}) interface{} {
	return b.value
}

func NewBoolCommon(value bool) *BoolCommon {
	return &BoolCommon{value}
}

type FloatCommon struct {
	value float64
}

func (f FloatCommon) Evaluate(symbols map[string]interface{}) interface{} {
	return f.value
}

func NewFloatCommon(value float64) *FloatCommon {
	return &FloatCommon{value}
}

type StringCommon struct {
	value string
}

func (s StringCommon) Evaluate(symbols map[string]interface{}) interface{} {
	return s.value
}

func NewStringCommon(value string) *StringCommon {
	return &StringCommon{value}
}

type AssignmentCommon struct {
	value      string
	expression Expression
}

func NewAssignmentCommon(value string, expression Expression) *AssignmentCommon {
	return &AssignmentCommon{value, expression}
}

type SumCommon struct {
	rightValue Expression
	leftValue  Expression
}

func (s SumCommon) Evaluate(symbols map[string]interface{}) interface{} {
	left := s.leftValue.Evaluate(symbols)
	right := s.rightValue.Evaluate(symbols)

	switch l := left.(type) {
	case int64:
		switch r := right.(type) {
		case int64:
			return l + r
		case float64:
			return l + int64(r)
		default:
			panic(fmt.Sprintf("Undefined + operator for types %v and %v", reflect.TypeOf(left), reflect.TypeOf(right)))
		}
	case float64:
		switch r := right.(type) {
		case int64:
			return int64(l) + r
		case float64:
			return l + r
		default:
			panic(fmt.Sprintf("Undefined + operator for types %v and %v", reflect.TypeOf(left), reflect.TypeOf(right)))
		}
	}

	panic(fmt.Sprintf("Undefined + operator for types %v and %v", reflect.TypeOf(left), reflect.TypeOf(right)))
}

func NewSumCommon(value Expression, expression Expression) *SumCommon {
	return &SumCommon{value, expression}
}

type SubCommon struct {
	rightValue Expression
	leftValue  Expression
}

func (s SubCommon) Evaluate(symbols map[string]interface{}) interface{} {
	left := s.leftValue.Evaluate(symbols)
	right := s.rightValue.Evaluate(symbols)

	switch l := left.(type) {
	case int64:
		switch r := right.(type) {
		case int64:
			return l - r
		case float64:
			return l - int64(r)
		default:
			panic(fmt.Sprintf("Undefined - operator for types %v and %v", reflect.TypeOf(left), reflect.TypeOf(right)))
		}
	case float64:
		switch r := right.(type) {
		case int64:
			return int64(l) - r
		case float64:
			return l - r
		default:
			panic(fmt.Sprintf("Undefined - operator for types %v and %v", reflect.TypeOf(left), reflect.TypeOf(right)))
		}
	}

	panic(fmt.Sprintf("Undefined - operator for types %v and %v", reflect.TypeOf(left), reflect.TypeOf(right)))
}

func NewSubCommon(value Expression, expression Expression) *SubCommon {
	return &SubCommon{value, expression}
}

type MultCommon struct {
	rightValue Expression
	leftValue  Expression
}

func (m MultCommon) Evaluate(symbols map[string]interface{}) interface{} {
	left := m.leftValue.Evaluate(symbols)
	right := m.rightValue.Evaluate(symbols)

	switch l := left.(type) {
	case int64:
		switch r := right.(type) {
		case int64:
			return l * r
		case float64:
			return l * int64(r)
		default:
			panic(fmt.Sprintf("Undefined * operator for types %v and %v", reflect.TypeOf(left), reflect.TypeOf(right)))
		}
	case float64:
		switch r := right.(type) {
		case int64:
			return int64(l) * r
		case float64:
			return l * r
		default:
			panic(fmt.Sprintf("Undefined * operator for types %v and %v", reflect.TypeOf(left), reflect.TypeOf(right)))
		}
	}

	panic(fmt.Sprintf("Undefined * operator for types %v and %v", reflect.TypeOf(left), reflect.TypeOf(right)))
}

func NewMultCommon(value Expression, expression Expression) *MultCommon {
	return &MultCommon{value, expression}
}

type DivCommon struct {
	rightValue Expression
	leftValue  Expression
}

func NewDivCommon(value Expression, expression Expression) *DivCommon {
	return &DivCommon{value, expression}
}

func (d DivCommon) Evaluate(symbols map[string]interface{}) interface{} {
	left := d.leftValue.Evaluate(symbols)
	right := d.rightValue.Evaluate(symbols)

	switch l := left.(type) {
	case int64:
		switch r := right.(type) {
		case int64:
			return l / r
		case float64:
			return l / int64(r)
		default:
			panic(fmt.Sprintf("Undefined / operator for types %v and %v", reflect.TypeOf(left), reflect.TypeOf(right)))
		}
	case float64:
		switch r := right.(type) {
		case int64:
			return int64(l) / r
		case float64:
			return l / r
		default:
			panic(fmt.Sprintf("Undefined / operator for types %v and %v", reflect.TypeOf(left), reflect.TypeOf(right)))
		}
	}

	panic(fmt.Sprintf("Undefined / operator for types %v and %v", reflect.TypeOf(left), reflect.TypeOf(right)))
}

type EqCommon struct {
	rightValue Expression
	leftValue  Expression
}

func (e EqCommon) Evaluate(symbols map[string]interface{}) interface{} {
	return e.leftValue.Evaluate(symbols) == e.rightValue.Evaluate(symbols)
}

func NewEqCommon(value Expression, expression Expression) *EqCommon {
	return &EqCommon{value, expression}
}

type NeCommon struct {
	rightValue Expression
	leftValue  Expression
}

func (ne NeCommon) Evaluate(symbols map[string]interface{}) interface{} {
	return ne.leftValue.Evaluate(symbols) != ne.rightValue.Evaluate(symbols)
}

func NewNeCommon(value Expression, expression Expression) *NeCommon {
	return &NeCommon{value, expression}
}

type LtCommon struct {
	rightValue Expression
	leftValue  Expression
}

func (lt LtCommon) Evaluate(symbols map[string]interface{}) interface{} {
	left := lt.leftValue.Evaluate(symbols)
	right := lt.rightValue.Evaluate(symbols)

	switch l := left.(type) {
	case int64:
		switch r := right.(type) {
		case int64:
			return l < r
		case float64:
			return l < int64(r)
		default:
			panic(fmt.Sprintf("Undefined < operator for types %v and %v", reflect.TypeOf(left), reflect.TypeOf(right)))
		}
	case float64:
		switch r := right.(type) {
		case int64:
			return int64(l) < r
		case float64:
			return l < r
		default:
			panic(fmt.Sprintf("Undefined < operator for types %v and %v", reflect.TypeOf(left), reflect.TypeOf(right)))
		}
	}

	panic(fmt.Sprintf("Undefined < operator for types %v and %v", reflect.TypeOf(left), reflect.TypeOf(right)))
}

func NewLtCommon(value Expression, expression Expression) *LtCommon {
	return &LtCommon{value, expression}
}

type GtCommon struct {
	rightValue Expression
	leftValue  Expression
}

func (gt GtCommon) Evaluate(symbols map[string]interface{}) interface{} {
	left := gt.leftValue.Evaluate(symbols)
	right := gt.rightValue.Evaluate(symbols)

	switch l := left.(type) {
	case int64:
		switch r := right.(type) {
		case int64:
			return l > r
		case float64:
			return l > int64(r)
		default:
			panic(fmt.Sprintf("Undefined > operator for types %v and %v", reflect.TypeOf(left), reflect.TypeOf(right)))
		}
	case float64:
		switch r := right.(type) {
		case int64:
			return int64(l) > r
		case float64:
			return l > r
		default:
			panic(fmt.Sprintf("Undefined > operator for types %v and %v", reflect.TypeOf(left), reflect.TypeOf(right)))
		}
	}

	panic(fmt.Sprintf("Undefined > operator for types %v and %v", reflect.TypeOf(left), reflect.TypeOf(right)))
}

func NewGtCommon(value Expression, expression Expression) *GtCommon {
	return &GtCommon{value, expression}
}

type LteCommon struct {
	rightValue Expression
	leftValue  Expression
}

func NewLteCommon(value Expression, expression Expression) *LteCommon {
	return &LteCommon{value, expression}
}

func (lte LteCommon) Evaluate(symbols map[string]interface{}) interface{} {
	left := lte.leftValue.Evaluate(symbols)
	right := lte.rightValue.Evaluate(symbols)

	switch l := left.(type) {
	case int64:
		switch r := right.(type) {
		case int64:
			return l <= r
		case float64:
			return l <= int64(r)
		default:
			panic(fmt.Sprintf("Undefined <= operator for types %v and %v", reflect.TypeOf(left), reflect.TypeOf(right)))
		}
	case float64:
		switch r := right.(type) {
		case int64:
			return int64(l) <= r
		case float64:
			return l <= r
		default:
			panic(fmt.Sprintf("Undefined <= operator for types %v and %v", reflect.TypeOf(left), reflect.TypeOf(right)))
		}
	}

	panic(fmt.Sprintf("Undefined <= operator for types %v and %v", reflect.TypeOf(left), reflect.TypeOf(right)))
}

type GteCommon struct {
	rightValue Expression
	leftValue  Expression
}

func (gte GteCommon) Evaluate(symbols map[string]interface{}) interface{} {
	left := gte.leftValue.Evaluate(symbols)
	right := gte.rightValue.Evaluate(symbols)

	switch l := left.(type) {
	case int64:
		switch r := right.(type) {
		case int64:
			return l >= r
		case float64:
			return l >= int64(r)
		default:
			panic(fmt.Sprintf("Undefined >= operator for types %v and %v", reflect.TypeOf(left), reflect.TypeOf(right)))
		}
	case float64:
		switch r := right.(type) {
		case int64:
			return int64(l) >= r
		case float64:
			return l >= r
		default:
			panic(fmt.Sprintf("Undefined >= operator for types %v and %v", reflect.TypeOf(left), reflect.TypeOf(right)))
		}
	}

	panic(fmt.Sprintf("Undefined >= operator for types %v and %v", reflect.TypeOf(left), reflect.TypeOf(right)))
}

func NewGteCommon(value Expression, expression Expression) *GteCommon {
	return &GteCommon{value, expression}
}

type BetweenCommon struct {
	rightValue Expression
	leftValue  Expression
}

func (b BetweenCommon) Evaluate(symbols map[string]interface{}) interface{} {
	panic("BETWEEN operator not implemented")
}

func NewBetweenCommon(value Expression, expression Expression) *BetweenCommon {
	return &BetweenCommon{value, expression}
}

type LikeCommon struct {
	rightValue Expression
	leftValue  Expression
}

func (l LikeCommon) Evaluate(symbols map[string]interface{}) interface{} {
	panic("BETWEEN operator not implemented")
}

func NewLikeCommon(value Expression, expression Expression) *LikeCommon {
	return &LikeCommon{value, expression}
}

type NotCommon struct {
	not Expression
}

func (n NotCommon) Evaluate(symbols map[string]interface{}) interface{} {
	operand := n.not.Evaluate(symbols)
	if b, ok := operand.(bool); ok {
		return !b
	}
	panic(fmt.Sprintf("Undefined NOT operator for type %d", reflect.TypeOf(operand)))
}

func NewNotCommon(value Expression) *NotCommon {
	return &NotCommon{value}
}

type AndCommon struct {
	rightValue Expression
	leftValue  Expression
}

func (a AndCommon) Evaluate(symbols map[string]interface{}) interface{} {
	left, ok1 := a.leftValue.Evaluate(symbols).(bool)
	right, ok2 := a.rightValue.Evaluate(symbols).(bool)

	if !(ok1 && ok2) {
		panic(fmt.Sprintf("Undefined AND operator for types %v and %v", reflect.TypeOf(left), reflect.TypeOf(right)))
	}

	return left && right
}

func NewAndCommon(value Expression, expression Expression) *AndCommon {
	return &AndCommon{value, expression}
}

type OrCommon struct {
	rightValue Expression
	leftValue  Expression
}

func (o OrCommon) Evaluate(symbols map[string]interface{}) interface{} {
	left, ok1 := o.leftValue.Evaluate(symbols).(bool)
	right, ok2 := o.rightValue.Evaluate(symbols).(bool)

	if !(ok1 && ok2) {
		panic(fmt.Sprintf("Undefined OR operator for types %v and %v", reflect.TypeOf(left), reflect.TypeOf(right)))
	}

	return left || right
}

func NewOrCommon(value Expression, expression Expression) *OrCommon {
	return &OrCommon{value, expression}
}

type NullCommon struct {
}

func (n NullCommon) Evaluate(symbols map[string]interface{}) interface{} {
	return nil
}

func NewNullCommon() *NullCommon {
	return &NullCommon{}
}

type FalseCommon struct {
}

func (f FalseCommon) Evaluate(symbols map[string]interface{}) interface{} {
	return false
}

func NewFalseCommon() *FalseCommon {
	return &FalseCommon{}
}

type TrueCommon struct {
}

func (t TrueCommon) Evaluate(symbols map[string]interface{}) interface{} {
	return true
}

func NewTrueCommon() *TrueCommon {
	return &TrueCommon{}
}
