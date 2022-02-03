package tests

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"reflect"
	"testing"
	"unicode/utf8"

	"github.com/stretchr/testify/require"
	"github.com/yu31/protoc-plugin/xgo/pkg/protovalidator"
	"github.com/yu31/protoc-plugin/xgo/tests/govalidatortest"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
)

type CaseDesc struct {
	Name       string
	FieldDesc  string
	TagInfo    *protovalidator.TagInfo
	FieldValue func() interface{}
	BeforeFunc func()
	AfterFunc  func()
	UseError2  bool
}

func valueToString(value interface{}) string {
	v := value

	refVal := reflect.ValueOf(value)
	if refVal.Kind() == reflect.Ptr {
		refVal = refVal.Elem()
	}

	switch refVal.Kind() {
	case reflect.Bool:
		v = refVal.Bool()
	case reflect.Float32:
		v = float32(refVal.Float())
	case reflect.Float64:
		v = refVal.Float()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		v = refVal.Int()
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		v = refVal.Uint()
	case reflect.String:
		v = refVal.String()
	}

	return fmt.Sprintf("%+v", v)
}

func runCases(t *testing.T, data protovalidator.Validator, msgName string, cases []*CaseDesc) {
	for _, c := range cases {
		source := fmt.Sprintf("[%s]-[%s]-[%s]", msgName, c.FieldDesc, c.Name)

		c.BeforeFunc()

		err := data.Validate()
		require.NotNil(t, err, source)
		//fmt.Println(err.Error())

		if c.FieldDesc != "" {
			reason := protovalidator.BuildErrorReason(c.TagInfo, c.FieldDesc)
			var expectedError error
			if c.UseError2 {
				expectedError = protovalidator.FieldError2(msgName, reason)
			} else {
				value := valueToString(c.FieldValue())
				expectedError = protovalidator.FieldError1(msgName, reason, value)
			}

			require.Equal(t, expectedError.Error(), err.Error(), source)
		}

		c.AfterFunc()
	}
}

func Test_GoValidator_ValidOneOfTags1(t *testing.T) {
	data := &govalidatortest.ValidOneOfTags1{}

	cases := []*CaseDesc{
		{
			Name:       fmt.Sprintf("test tag <%s>", protovalidator.TagOneOfNotNull),
			FieldDesc:  "field 'oneof_type1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagOneOfNotNull, Value: nil},
			FieldValue: func() interface{} { return data.OneofType1 },
			BeforeFunc: func() {},
			AfterFunc:  func() { data.OneofType1 = &govalidatortest.ValidOneOfTags1_Oneof1String1{Oneof1String1: "s1"} },
			UseError2:  true,
		},
	}

	msgName := "ValidOneOfTags1"
	runCases(t, data, msgName, cases)

	err := data.Validate()
	require.Nil(t, err)
}

func Test_GoValidator_ValidFloatTagsGeneral1(t *testing.T) {
	data := &govalidatortest.ValidFloatTagsGeneral1{}

	cases := []*CaseDesc{
		// type for float.
		{
			Name:       fmt.Sprintf("test tag <%s>", protovalidator.TagFloatEq),
			FieldDesc:  "field 't_float_eq1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagFloatEq, Value: 1.1},
			FieldValue: func() interface{} { return data.TFloatEq1 },
			BeforeFunc: func() {},
			AfterFunc:  func() { data.TFloatEq1 = 1.1 },
		},
		{
			Name:       fmt.Sprintf("test tag <%s>", protovalidator.TagFloatNe),
			FieldDesc:  "field 't_float_ne1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagFloatNe, Value: 2.1},
			FieldValue: func() interface{} { return data.TFloatNe1 },
			BeforeFunc: func() { data.TFloatNe1 = 2.1 },
			AfterFunc:  func() { data.TFloatNe1 = 1.1 },
		},
		{
			Name:       fmt.Sprintf("test tag <%s>", protovalidator.TagFloatLt),
			FieldDesc:  "field 't_float_lt1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagFloatLt, Value: 3.1},
			FieldValue: func() interface{} { return data.TFloatLt1 },
			BeforeFunc: func() { data.TFloatLt1 = 4 },
			AfterFunc:  func() { data.TFloatLt1 = 1.1 },
		},
		{
			Name:       fmt.Sprintf("test tag <%s>", protovalidator.TagFloatGt),
			FieldDesc:  "field 't_float_gt1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagFloatGt, Value: 4.1},
			FieldValue: func() interface{} { return data.TFloatGt1 },
			BeforeFunc: func() {},
			AfterFunc:  func() { data.TFloatGt1 = 5.1 },
		},
		{
			Name:       fmt.Sprintf("test tag <%s>", protovalidator.TagFloatLte),
			FieldDesc:  "field 't_float_lte1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagFloatLte, Value: 5.1},
			FieldValue: func() interface{} { return data.TFloatLte1 },
			BeforeFunc: func() { data.TFloatLte1 = 5.2 },
			AfterFunc:  func() { data.TFloatLte1 = 5.1 },
		},
		{
			Name:       fmt.Sprintf("test tag <%s>", protovalidator.TagFloatGte),
			FieldDesc:  "field 't_float_gte1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagFloatGte, Value: 6.1},
			FieldValue: func() interface{} { return data.TFloatGte1 },
			BeforeFunc: func() {},
			AfterFunc:  func() { data.TFloatGte1 = 6.1 },
		},
		{
			Name:       fmt.Sprintf("test tag <%s>", protovalidator.TagFloatIn),
			FieldDesc:  "field 't_float_in1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagFloatIn, Value: []float32{1.1, 1.2, 1.3}},
			FieldValue: func() interface{} { return data.TFloatIn1 },
			BeforeFunc: func() {},
			AfterFunc:  func() { data.TFloatIn1 = 1.1 },
		},
		{
			Name:       fmt.Sprintf("test tag <%s>", protovalidator.TagFloatNotIn),
			FieldDesc:  "field 't_float_not_in1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagFloatNotIn, Value: []float32{2.1, 2.2, 2.3}},
			FieldValue: func() interface{} { return data.TFloatNotIn1 },
			BeforeFunc: func() { data.TFloatNotIn1 = 2.1 },
			AfterFunc:  func() { data.TFloatNotIn1 = 2 },
		},

		// type for double.
		{
			Name:       fmt.Sprintf("test tag <%s>", protovalidator.TagFloatEq),
			FieldDesc:  "field 't_double_eq1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagFloatEq, Value: 1.1},
			FieldValue: func() interface{} { return data.TDoubleEq1 },
			BeforeFunc: func() {},
			AfterFunc:  func() { data.TDoubleEq1 = 1.1 },
		},
		{
			Name:       fmt.Sprintf("test tag <%s>", protovalidator.TagFloatNe),
			FieldDesc:  "field 't_double_ne1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagFloatNe, Value: 2.1},
			FieldValue: func() interface{} { return data.TDoubleNe1 },
			BeforeFunc: func() { data.TDoubleNe1 = 2.1 },
			AfterFunc:  func() { data.TDoubleNe1 = 1.1 },
		},
		{
			Name:       fmt.Sprintf("test tag <%s>", protovalidator.TagFloatLt),
			FieldDesc:  "field 't_double_lt1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagFloatLt, Value: 3.1},
			FieldValue: func() interface{} { return data.TDoubleLt1 },
			BeforeFunc: func() { data.TDoubleLt1 = 4 },
			AfterFunc:  func() { data.TDoubleLt1 = 1.1 },
		},
		{
			Name:       fmt.Sprintf("test tag <%s>", protovalidator.TagFloatGt),
			FieldDesc:  "field 't_double_gt1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagFloatGt, Value: 4.1},
			FieldValue: func() interface{} { return data.TDoubleGt1 },
			BeforeFunc: func() {},
			AfterFunc:  func() { data.TDoubleGt1 = 5.1 },
		},
		{
			Name:       fmt.Sprintf("test tag <%s>", protovalidator.TagFloatLte),
			FieldDesc:  "field 't_double_lte1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagFloatLte, Value: 5.1},
			FieldValue: func() interface{} { return data.TDoubleLte1 },
			BeforeFunc: func() { data.TDoubleLte1 = 5.2 },
			AfterFunc:  func() { data.TDoubleLte1 = 5.1 },
		},
		{
			Name:       fmt.Sprintf("test tag <%s>", protovalidator.TagFloatGte),
			FieldDesc:  "field 't_double_gte1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagFloatGte, Value: 6.1},
			FieldValue: func() interface{} { return data.TDoubleGte1 },
			BeforeFunc: func() {},
			AfterFunc:  func() { data.TDoubleGte1 = 6.1 },
		},
		{
			Name:       fmt.Sprintf("test tag <%s>", protovalidator.TagFloatIn),
			FieldDesc:  "field 't_double_in1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagFloatIn, Value: []float32{1.1, 1.2, 1.3}},
			FieldValue: func() interface{} { return data.TDoubleIn1 },
			BeforeFunc: func() {},
			AfterFunc:  func() { data.TDoubleIn1 = 1.1 },
		},
		{
			Name:       fmt.Sprintf("test tag <%s>", protovalidator.TagFloatNotIn),
			FieldDesc:  "field 't_double_not_in1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagFloatNotIn, Value: []float32{2.1, 2.2, 2.3}},
			FieldValue: func() interface{} { return data.TDoubleNotIn1 },
			BeforeFunc: func() { data.TDoubleNotIn1 = 2.1 },
			AfterFunc:  func() { data.TDoubleNotIn1 = 2 },
		},
	}

	msgName := "ValidFloatTagsGeneral1"
	runCases(t, data, msgName, cases)

	err := data.Validate()
	require.Nil(t, err)
}

func Test_GoValidator_ValidFloatTagsOptional1(t *testing.T) {
	data := &govalidatortest.ValidFloatTagsOptional1{}
	cases := []*CaseDesc{
		// type for float.
		{
			Name:       fmt.Sprintf("test tag <%s>", protovalidator.TagFloatEq),
			FieldDesc:  "field 't_float_eq1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagFloatEq, Value: 1.1},
			FieldValue: func() interface{} { return nil },
			BeforeFunc: func() {},
			AfterFunc:  func() { x := float32(1.1); data.TFloatEq1 = &x },
		},
		{
			Name:       fmt.Sprintf("test tag <%s>", protovalidator.TagFloatNe),
			FieldDesc:  "field 't_float_ne1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagFloatNe, Value: 2.1},
			FieldValue: func() interface{} { return *data.TFloatNe1 },
			BeforeFunc: func() { x := float32(2.1); data.TFloatNe1 = &x },
			AfterFunc:  func() { x := float32(2.2); data.TFloatNe1 = &x },
		},
		{
			Name:       fmt.Sprintf("test tag <%s>", protovalidator.TagFloatLt),
			FieldDesc:  "field 't_float_lt1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagFloatLt, Value: 3.1},
			FieldValue: func() interface{} { return *data.TFloatLt1 },
			BeforeFunc: func() { x := float32(4); data.TFloatLt1 = &x },
			AfterFunc:  func() { x := float32(3.0); data.TFloatLt1 = &x },
		},
		{
			Name:       fmt.Sprintf("test tag <%s>", protovalidator.TagFloatGt),
			FieldDesc:  "field 't_float_gt1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagFloatGt, Value: 4.1},
			FieldValue: func() interface{} { return nil },
			BeforeFunc: func() {},
			AfterFunc:  func() { x := float32(4.2); data.TFloatGt1 = &x },
		},
		{
			Name:       fmt.Sprintf("test tag <%s>", protovalidator.TagFloatLte),
			FieldDesc:  "field 't_float_lte1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagFloatLte, Value: 5.1},
			FieldValue: func() interface{} { return *data.TFloatLte1 },
			BeforeFunc: func() { x := float32(5.2); data.TFloatLte1 = &x },
			AfterFunc:  func() { x := float32(5.1); data.TFloatLte1 = &x },
		},
		{
			Name:       fmt.Sprintf("test tag <%s>", protovalidator.TagFloatGte),
			FieldDesc:  "field 't_float_gte1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagFloatGte, Value: 6.1},
			FieldValue: func() interface{} { return nil },
			BeforeFunc: func() {},
			AfterFunc:  func() { x := float32(6.1); data.TFloatGte1 = &x },
		},
		{
			Name:       fmt.Sprintf("test tag <%s>", protovalidator.TagFloatIn),
			FieldDesc:  "field 't_float_in1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagFloatIn, Value: []float32{1.1, 1.2, 1.3}},
			FieldValue: func() interface{} { return nil },
			BeforeFunc: func() {},
			AfterFunc:  func() { x := float32(1.1); data.TFloatIn1 = &x },
		},
		{
			Name:       fmt.Sprintf("test tag <%s>", protovalidator.TagFloatNotIn),
			FieldDesc:  "field 't_float_not_in1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagFloatNotIn, Value: []float32{2.1, 2.2, 2.3}},
			FieldValue: func() interface{} { return *data.TFloatNotIn1 },
			BeforeFunc: func() { x := float32(2.1); data.TFloatNotIn1 = &x },
			AfterFunc:  func() { x := float32(1.1); data.TFloatNotIn1 = &x },
		},
		// type for double.
		{
			Name:       fmt.Sprintf("test tag <%s>", protovalidator.TagFloatEq),
			FieldDesc:  "field 't_double_eq1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagFloatEq, Value: 1.1},
			FieldValue: func() interface{} { return nil },
			BeforeFunc: func() {},
			AfterFunc:  func() { x := float64(1.1); data.TDoubleEq1 = &x },
		},
		{
			Name:       fmt.Sprintf("test tag <%s>", protovalidator.TagFloatNe),
			FieldDesc:  "field 't_double_ne1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagFloatNe, Value: 2.1},
			FieldValue: func() interface{} { return *data.TDoubleNe1 },
			BeforeFunc: func() { x := 2.1; data.TDoubleNe1 = &x },
			AfterFunc:  func() { x := float64(1.1); data.TDoubleNe1 = &x },
		},
		{
			Name:       fmt.Sprintf("test tag <%s>", protovalidator.TagFloatLt),
			FieldDesc:  "field 't_double_lt1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagFloatLt, Value: 3.1},
			FieldValue: func() interface{} { return *data.TDoubleLt1 },
			BeforeFunc: func() { x := 4.1; data.TDoubleLt1 = &x },
			AfterFunc:  func() { x := float64(1.1); data.TDoubleLt1 = &x },
		},
		{
			Name:       fmt.Sprintf("test tag <%s>", protovalidator.TagFloatGt),
			FieldDesc:  "field 't_double_gt1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagFloatGt, Value: 4.1},
			FieldValue: func() interface{} { return nil },
			BeforeFunc: func() {},
			AfterFunc:  func() { x := float64(5); data.TDoubleGt1 = &x },
		},
		{
			Name:       fmt.Sprintf("test tag <%s>", protovalidator.TagFloatLte),
			FieldDesc:  "field 't_double_lte1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagFloatLte, Value: 5.1},
			FieldValue: func() interface{} { return *data.TDoubleLte1 },
			BeforeFunc: func() { x := 5.2; data.TDoubleLte1 = &x },
			AfterFunc:  func() { x := float64(1.1); data.TDoubleLte1 = &x },
		},
		{
			Name:       fmt.Sprintf("test tag <%s>", protovalidator.TagFloatGte),
			FieldDesc:  "field 't_double_gte1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagFloatGte, Value: 6.1},
			FieldValue: func() interface{} { return nil },
			BeforeFunc: func() {},
			AfterFunc:  func() { x := float64(6.2); data.TDoubleGte1 = &x },
		},
		{
			Name:       fmt.Sprintf("test tag <%s>", protovalidator.TagFloatIn),
			FieldDesc:  "field 't_double_in1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagFloatIn, Value: []float32{1.1, 1.2, 1.3}},
			FieldValue: func() interface{} { return nil },
			BeforeFunc: func() {},
			AfterFunc:  func() { x := float64(1.1); data.TDoubleIn1 = &x },
		},
		{
			Name:       fmt.Sprintf("test tag <%s>", protovalidator.TagFloatNotIn),
			FieldDesc:  "field 't_double_not_in1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagFloatNotIn, Value: []float32{2.1, 2.2, 2.3}},
			FieldValue: func() interface{} { return *data.TDoubleNotIn1 },
			BeforeFunc: func() { x := 2.1; data.TDoubleNotIn1 = &x },
			AfterFunc:  func() { x := float64(1.1); data.TDoubleNotIn1 = &x },
		},
	}

	msgName := "ValidFloatTagsOptional1"
	runCases(t, data, msgName, cases)

	err := data.Validate()
	require.Nil(t, err)
}

func Test_GoValidator_ValidFloatTagsOneOf1(t *testing.T) {
	data := &govalidatortest.ValidFloatTagsOneOf1{OneTyp1: nil}

	{
		err := data.Validate()
		require.Nil(t, err)
	}

	cases := []*CaseDesc{
		// type for float.
		{
			Name:       fmt.Sprintf("test tag <%s>", protovalidator.TagFloatEq),
			FieldDesc:  "field 't_float_eq1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagFloatEq, Value: 1.1},
			FieldValue: func() interface{} { return data.OneTyp1.(*govalidatortest.ValidFloatTagsOneOf1_TFloatEq1).TFloatEq1 },
			BeforeFunc: func() { data.OneTyp1 = &govalidatortest.ValidFloatTagsOneOf1_TFloatEq1{TFloatEq1: 0} },
			AfterFunc:  func() { data.OneTyp1 = &govalidatortest.ValidFloatTagsOneOf1_TFloatEq1{TFloatEq1: 1.1} },
		},
		{
			Name:       fmt.Sprintf("test tag <%s>", protovalidator.TagFloatNe),
			FieldDesc:  "field 't_float_ne1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagFloatNe, Value: 2.1},
			FieldValue: func() interface{} { return data.OneTyp1.(*govalidatortest.ValidFloatTagsOneOf1_TFloatNe1).TFloatNe1 },
			BeforeFunc: func() { data.OneTyp1 = &govalidatortest.ValidFloatTagsOneOf1_TFloatNe1{TFloatNe1: 2.1} },
			AfterFunc:  func() { data.OneTyp1 = &govalidatortest.ValidFloatTagsOneOf1_TFloatNe1{TFloatNe1: 1.1} },
		},
		{
			Name:       fmt.Sprintf("test tag <%s>", protovalidator.TagFloatLt),
			FieldDesc:  "field 't_float_lt1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagFloatLt, Value: 3.1},
			FieldValue: func() interface{} { return data.OneTyp1.(*govalidatortest.ValidFloatTagsOneOf1_TFloatLt1).TFloatLt1 },
			BeforeFunc: func() { data.OneTyp1 = &govalidatortest.ValidFloatTagsOneOf1_TFloatLt1{TFloatLt1: 4} },
			AfterFunc:  func() { data.OneTyp1 = &govalidatortest.ValidFloatTagsOneOf1_TFloatLt1{TFloatLt1: 1.1} },
		},
		{
			Name:       fmt.Sprintf("test tag <%s>", protovalidator.TagFloatGt),
			FieldDesc:  "field 't_float_gt1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagFloatGt, Value: 4.1},
			FieldValue: func() interface{} { return data.OneTyp1.(*govalidatortest.ValidFloatTagsOneOf1_TFloatGt1).TFloatGt1 },
			BeforeFunc: func() { data.OneTyp1 = &govalidatortest.ValidFloatTagsOneOf1_TFloatGt1{TFloatGt1: 0} },
			AfterFunc:  func() { data.OneTyp1 = &govalidatortest.ValidFloatTagsOneOf1_TFloatGt1{TFloatGt1: 5.1} },
		},
		{
			Name:       fmt.Sprintf("test tag <%s>", protovalidator.TagFloatLte),
			FieldDesc:  "field 't_float_lte1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagFloatLte, Value: 5.1},
			FieldValue: func() interface{} { return data.OneTyp1.(*govalidatortest.ValidFloatTagsOneOf1_TFloatLte1).TFloatLte1 },
			BeforeFunc: func() { data.OneTyp1 = &govalidatortest.ValidFloatTagsOneOf1_TFloatLte1{TFloatLte1: 5.2} },
			AfterFunc:  func() { data.OneTyp1 = &govalidatortest.ValidFloatTagsOneOf1_TFloatLte1{TFloatLte1: 5.1} },
		},
		{
			Name:       fmt.Sprintf("test tag <%s>", protovalidator.TagFloatGte),
			FieldDesc:  "field 't_float_gte1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagFloatGte, Value: 6.1},
			FieldValue: func() interface{} { return data.OneTyp1.(*govalidatortest.ValidFloatTagsOneOf1_TFloatGte1).TFloatGte1 },
			BeforeFunc: func() { data.OneTyp1 = &govalidatortest.ValidFloatTagsOneOf1_TFloatGte1{TFloatGte1: 0} },
			AfterFunc:  func() { data.OneTyp1 = &govalidatortest.ValidFloatTagsOneOf1_TFloatGte1{TFloatGte1: 6.1} },
		},
		{
			Name:       fmt.Sprintf("test tag <%s>", protovalidator.TagFloatIn),
			FieldDesc:  "field 't_float_in1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagFloatIn, Value: []float32{1.1, 1.2, 1.3}},
			FieldValue: func() interface{} { return data.OneTyp1.(*govalidatortest.ValidFloatTagsOneOf1_TFloatIn1).TFloatIn1 },
			BeforeFunc: func() { data.OneTyp1 = &govalidatortest.ValidFloatTagsOneOf1_TFloatIn1{TFloatIn1: 0} },
			AfterFunc:  func() { data.OneTyp1 = &govalidatortest.ValidFloatTagsOneOf1_TFloatIn1{TFloatIn1: 1.1} },
		},
		{
			Name:      fmt.Sprintf("test tag <%s>", protovalidator.TagFloatNotIn),
			FieldDesc: "field 't_float_not_in1'",
			TagInfo:   &protovalidator.TagInfo{Tag: protovalidator.TagFloatNotIn, Value: []float32{2.1, 2.2, 2.3}},
			FieldValue: func() interface{} {
				return data.OneTyp1.(*govalidatortest.ValidFloatTagsOneOf1_TFloatNotIn1).TFloatNotIn1
			},
			BeforeFunc: func() { data.OneTyp1 = &govalidatortest.ValidFloatTagsOneOf1_TFloatNotIn1{TFloatNotIn1: 2.2} },
			AfterFunc:  func() { data.OneTyp1 = &govalidatortest.ValidFloatTagsOneOf1_TFloatNotIn1{TFloatNotIn1: 2} },
		},

		//
		//// type for double.
		{
			Name:       fmt.Sprintf("test tag <%s>", protovalidator.TagFloatEq),
			FieldDesc:  "field 't_double_eq1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagFloatEq, Value: 1.1},
			FieldValue: func() interface{} { return data.OneTyp1.(*govalidatortest.ValidFloatTagsOneOf1_TDoubleEq1).TDoubleEq1 },
			BeforeFunc: func() { data.OneTyp1 = &govalidatortest.ValidFloatTagsOneOf1_TDoubleEq1{TDoubleEq1: 0} },
			AfterFunc:  func() { data.OneTyp1 = &govalidatortest.ValidFloatTagsOneOf1_TDoubleEq1{TDoubleEq1: 1.1} },
		},
		{
			Name:       fmt.Sprintf("test tag <%s>", protovalidator.TagFloatNe),
			FieldDesc:  "field 't_double_ne1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagFloatNe, Value: 2.1},
			FieldValue: func() interface{} { return data.OneTyp1.(*govalidatortest.ValidFloatTagsOneOf1_TDoubleNe1).TDoubleNe1 },
			BeforeFunc: func() { data.OneTyp1 = &govalidatortest.ValidFloatTagsOneOf1_TDoubleNe1{TDoubleNe1: 2.1} },
			AfterFunc:  func() { data.OneTyp1 = &govalidatortest.ValidFloatTagsOneOf1_TDoubleNe1{TDoubleNe1: 1.1} },
		},
		{
			Name:       fmt.Sprintf("test tag <%s>", protovalidator.TagFloatLt),
			FieldDesc:  "field 't_double_lt1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagFloatLt, Value: 3.1},
			FieldValue: func() interface{} { return data.OneTyp1.(*govalidatortest.ValidFloatTagsOneOf1_TDoubleLt1).TDoubleLt1 },
			BeforeFunc: func() { data.OneTyp1 = &govalidatortest.ValidFloatTagsOneOf1_TDoubleLt1{TDoubleLt1: 4} },
			AfterFunc:  func() { data.OneTyp1 = &govalidatortest.ValidFloatTagsOneOf1_TDoubleLt1{TDoubleLt1: 1.1} },
		},
		{
			Name:       fmt.Sprintf("test tag <%s>", protovalidator.TagFloatEq),
			FieldDesc:  "field 't_double_gt1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagFloatGt, Value: 4.1},
			FieldValue: func() interface{} { return data.OneTyp1.(*govalidatortest.ValidFloatTagsOneOf1_TDoubleGt1).TDoubleGt1 },
			BeforeFunc: func() { data.OneTyp1 = &govalidatortest.ValidFloatTagsOneOf1_TDoubleGt1{TDoubleGt1: 0} },
			AfterFunc:  func() { data.OneTyp1 = &govalidatortest.ValidFloatTagsOneOf1_TDoubleGt1{TDoubleGt1: 5.1} },
		},
		{
			Name:      fmt.Sprintf("test tag <%s>", protovalidator.TagFloatLte),
			FieldDesc: "field 't_double_lte1'",
			TagInfo:   &protovalidator.TagInfo{Tag: protovalidator.TagFloatLte, Value: 5.1},
			FieldValue: func() interface{} {
				return data.OneTyp1.(*govalidatortest.ValidFloatTagsOneOf1_TDoubleLte1).TDoubleLte1
			},
			BeforeFunc: func() { data.OneTyp1 = &govalidatortest.ValidFloatTagsOneOf1_TDoubleLte1{TDoubleLte1: 5.2} },
			AfterFunc:  func() { data.OneTyp1 = &govalidatortest.ValidFloatTagsOneOf1_TDoubleLte1{TDoubleLte1: 5.1} },
		},
		{
			Name:      fmt.Sprintf("test tag <%s>", protovalidator.TagFloatGte),
			FieldDesc: "field 't_double_gte1'",
			TagInfo:   &protovalidator.TagInfo{Tag: protovalidator.TagFloatGte, Value: 6.1},
			FieldValue: func() interface{} {
				return data.OneTyp1.(*govalidatortest.ValidFloatTagsOneOf1_TDoubleGte1).TDoubleGte1
			},
			BeforeFunc: func() { data.OneTyp1 = &govalidatortest.ValidFloatTagsOneOf1_TDoubleGte1{TDoubleGte1: 0} },
			AfterFunc:  func() { data.OneTyp1 = &govalidatortest.ValidFloatTagsOneOf1_TDoubleGte1{TDoubleGte1: 6.1} },
		},
		{
			Name:       fmt.Sprintf("test tag <%s>", protovalidator.TagFloatIn),
			FieldDesc:  "field 't_double_in1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagFloatIn, Value: []float64{1.1, 1.2, 1.3}},
			FieldValue: func() interface{} { return data.OneTyp1.(*govalidatortest.ValidFloatTagsOneOf1_TDoubleIn1).TDoubleIn1 },
			BeforeFunc: func() { data.OneTyp1 = &govalidatortest.ValidFloatTagsOneOf1_TDoubleIn1{TDoubleIn1: 0} },
			AfterFunc:  func() { data.OneTyp1 = &govalidatortest.ValidFloatTagsOneOf1_TDoubleIn1{TDoubleIn1: 1.1} },
		},
		{
			Name:      fmt.Sprintf("test tag <%s>", protovalidator.TagFloatNotIn),
			FieldDesc: "field 't_double_not_in1'",
			TagInfo:   &protovalidator.TagInfo{Tag: protovalidator.TagFloatNotIn, Value: []float64{2.1, 2.2, 2.3}},
			FieldValue: func() interface{} {
				return data.OneTyp1.(*govalidatortest.ValidFloatTagsOneOf1_TDoubleNotIn1).TDoubleNotIn1
			},
			BeforeFunc: func() { data.OneTyp1 = &govalidatortest.ValidFloatTagsOneOf1_TDoubleNotIn1{TDoubleNotIn1: 2.2} },
			AfterFunc:  func() { data.OneTyp1 = &govalidatortest.ValidFloatTagsOneOf1_TDoubleNotIn1{TDoubleNotIn1: 2} },
		},
	}

	msgName := "ValidFloatTagsOneOf1"
	runCases(t, data, msgName, cases)

	{
		err := data.Validate()
		require.Nil(t, err)
	}
}

func Test_GoValidator_ValidIntTagsGeneral1(t *testing.T) {
	data := &govalidatortest.ValidIntTagsGeneral1{}
	{
		err := data.Validate()
		require.NotNil(t, err)
	}

	casesInt32 := []*CaseDesc{
		{
			Name:       fmt.Sprintf("test tag <%s>", protovalidator.TagIntEq),
			FieldDesc:  "field 't_int32_eq1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagIntEq, Value: 1},
			FieldValue: func() interface{} { return data.TInt32Eq1 },
			BeforeFunc: func() {},
			AfterFunc:  func() { data.TInt32Eq1 = 1 },
		},
		{
			Name:       fmt.Sprintf("test tag <%s>", protovalidator.TagIntNe),
			FieldDesc:  "field 't_int32_ne1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagIntNe, Value: 2},
			FieldValue: func() interface{} { return data.TInt32Ne1 },
			BeforeFunc: func() { data.TInt32Ne1 = 2 },
			AfterFunc:  func() { data.TInt32Ne1 = 1 },
		},
		{
			Name:       fmt.Sprintf("test tag <%s>", protovalidator.TagIntLt),
			FieldDesc:  "field 't_int32_lt1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagIntLt, Value: 3},
			FieldValue: func() interface{} { return data.TInt32Lt1 },
			BeforeFunc: func() { data.TInt32Lt1 = 3 },
			AfterFunc:  func() { data.TInt32Lt1 = 1 },
		},
		{
			Name:       fmt.Sprintf("test tag <%s>", protovalidator.TagIntGt),
			FieldDesc:  "field 't_int32_gt1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagIntGt, Value: 4},
			FieldValue: func() interface{} { return data.TInt32Gt1 },
			BeforeFunc: func() {},
			AfterFunc:  func() { data.TInt32Gt1 = 5 },
		},
		{
			Name:       fmt.Sprintf("test tag <%s>", protovalidator.TagIntLte),
			FieldDesc:  "field 't_int32_lte1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagIntLte, Value: 5},
			FieldValue: func() interface{} { return data.TInt32Lte1 },
			BeforeFunc: func() { data.TInt32Lte1 = 6 },
			AfterFunc:  func() { data.TInt32Lte1 = 5 },
		},
		{
			Name:       fmt.Sprintf("test tag <%s>", protovalidator.TagIntGte),
			FieldDesc:  "field 't_int32_gte1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagIntGte, Value: 6},
			FieldValue: func() interface{} { return data.TInt32Gte1 },
			BeforeFunc: func() {},
			AfterFunc:  func() { data.TInt32Gte1 = 6 },
		},
		{
			Name:       fmt.Sprintf("test tag <%s>", protovalidator.TagIntIn),
			FieldDesc:  "field 't_int32_in1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagIntIn, Value: []int32{1, 2, 3}},
			FieldValue: func() interface{} { return data.TInt32In1 },
			BeforeFunc: func() { data.TInt32In1 = 4 },
			AfterFunc:  func() { data.TInt32In1 = 1 },
		},
		{
			Name:       fmt.Sprintf("test tag <%s>", protovalidator.TagIntNotIn),
			FieldDesc:  "field 't_int32_not_in1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagIntNotIn, Value: []int32{1, 2, 3}},
			FieldValue: func() interface{} { return data.TInt32NotIn1 },
			BeforeFunc: func() { data.TInt32NotIn1 = 1 },
			AfterFunc:  func() { data.TInt32NotIn1 = 5 },
		},
	}
	casesInt64 := []*CaseDesc{
		{
			Name:       fmt.Sprintf("test tag <%s>", protovalidator.TagIntEq),
			FieldDesc:  "field 't_int64_eq1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagIntEq, Value: 1},
			FieldValue: func() interface{} { return data.TInt64Eq1 },
			BeforeFunc: func() {},
			AfterFunc:  func() { data.TInt64Eq1 = 1 },
		},
		{
			Name:       fmt.Sprintf("test tag <%s>", protovalidator.TagIntNe),
			FieldDesc:  "field 't_int64_ne1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagIntNe, Value: 2},
			FieldValue: func() interface{} { return data.TInt64Ne1 },
			BeforeFunc: func() { data.TInt64Ne1 = 2 },
			AfterFunc:  func() { data.TInt64Ne1 = 1 },
		},
		{
			Name:       fmt.Sprintf("test tag <%s>", protovalidator.TagIntLt),
			FieldDesc:  "field 't_int64_lt1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagIntLt, Value: 3},
			FieldValue: func() interface{} { return data.TInt64Lt1 },
			BeforeFunc: func() { data.TInt64Lt1 = 3 },
			AfterFunc:  func() { data.TInt64Lt1 = 1 },
		},
		{
			Name:       fmt.Sprintf("test tag <%s>", protovalidator.TagIntGt),
			FieldDesc:  "field 't_int64_gt1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagIntGt, Value: 4},
			FieldValue: func() interface{} { return data.TInt64Gt1 },
			BeforeFunc: func() {},
			AfterFunc:  func() { data.TInt64Gt1 = 5 },
		},
		{
			Name:       fmt.Sprintf("test tag <%s>", protovalidator.TagIntLte),
			FieldDesc:  "field 't_int64_lte1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagIntLte, Value: 5},
			FieldValue: func() interface{} { return data.TInt64Lte1 },
			BeforeFunc: func() { data.TInt64Lte1 = 6 },
			AfterFunc:  func() { data.TInt64Lte1 = 5 },
		},
		{
			Name:       fmt.Sprintf("test tag <%s>", protovalidator.TagIntGte),
			FieldDesc:  "field 't_int64_gte1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagIntGte, Value: 6},
			FieldValue: func() interface{} { return data.TInt64Gte1 },
			BeforeFunc: func() {},
			AfterFunc:  func() { data.TInt64Gte1 = 6 },
		},
		{
			Name:       fmt.Sprintf("test tag <%s>", protovalidator.TagIntIn),
			FieldDesc:  "field 't_int64_in1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagIntIn, Value: []int32{1, 2, 3}},
			FieldValue: func() interface{} { return data.TInt64In1 },
			BeforeFunc: func() { data.TInt64In1 = 4 },
			AfterFunc:  func() { data.TInt64In1 = 1 },
		},
		{
			Name:       fmt.Sprintf("test tag <%s>", protovalidator.TagIntNotIn),
			FieldDesc:  "field 't_int64_not_in1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagIntNotIn, Value: []int32{1, 2, 3}},
			FieldValue: func() interface{} { return data.TInt64NotIn1 },
			BeforeFunc: func() { data.TInt64NotIn1 = 1 },
			AfterFunc:  func() { data.TInt64NotIn1 = 5 },
		},
	}
	casesSint32 := []*CaseDesc{
		{
			Name:       fmt.Sprintf("test tag <%s>", protovalidator.TagIntEq),
			FieldDesc:  "field 't_sint32_eq1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagIntEq, Value: 1},
			FieldValue: func() interface{} { return data.TSint32Eq1 },
			BeforeFunc: func() {},
			AfterFunc:  func() { data.TSint32Eq1 = 1 },
		},
		{
			Name:       fmt.Sprintf("test tag <%s>", protovalidator.TagIntNe),
			FieldDesc:  "field 't_sint32_ne1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagIntNe, Value: 2},
			FieldValue: func() interface{} { return data.TSint32Ne1 },
			BeforeFunc: func() { data.TSint32Ne1 = 2 },
			AfterFunc:  func() { data.TSint32Ne1 = 1 },
		},
		{
			Name:       fmt.Sprintf("test tag <%s>", protovalidator.TagIntLt),
			FieldDesc:  "field 't_sint32_lt1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagIntLt, Value: 3},
			FieldValue: func() interface{} { return data.TSint32Lt1 },
			BeforeFunc: func() { data.TSint32Lt1 = 3 },
			AfterFunc:  func() { data.TSint32Lt1 = 1 },
		},
		{
			Name:       fmt.Sprintf("test tag <%s>", protovalidator.TagIntGt),
			FieldDesc:  "field 't_sint32_gt1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagIntGt, Value: 4},
			FieldValue: func() interface{} { return data.TSint32Gt1 },
			BeforeFunc: func() {},
			AfterFunc:  func() { data.TSint32Gt1 = 5 },
		},
		{
			Name:       fmt.Sprintf("test tag <%s>", protovalidator.TagIntLte),
			FieldDesc:  "field 't_sint32_lte1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagIntLte, Value: 5},
			FieldValue: func() interface{} { return data.TSint32Lte1 },
			BeforeFunc: func() { data.TSint32Lte1 = 6 },
			AfterFunc:  func() { data.TSint32Lte1 = 5 },
		},
		{
			Name:       fmt.Sprintf("test tag <%s>", protovalidator.TagIntGte),
			FieldDesc:  "field 't_sint32_gte1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagIntGte, Value: 6},
			FieldValue: func() interface{} { return data.TSint32Gte1 },
			BeforeFunc: func() {},
			AfterFunc:  func() { data.TSint32Gte1 = 6 },
		},
		{
			Name:       fmt.Sprintf("test tag <%s>", protovalidator.TagIntIn),
			FieldDesc:  "field 't_sint32_in1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagIntIn, Value: []int32{1, 2, 3}},
			FieldValue: func() interface{} { return data.TSint32In1 },
			BeforeFunc: func() { data.TSint32In1 = 4 },
			AfterFunc:  func() { data.TSint32In1 = 1 },
		},
		{
			Name:       fmt.Sprintf("test tag <%s>", protovalidator.TagIntNotIn),
			FieldDesc:  "field 't_sint32_not_in1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagIntNotIn, Value: []int32{1, 2, 3}},
			FieldValue: func() interface{} { return data.TSint32NotIn1 },
			BeforeFunc: func() { data.TSint32NotIn1 = 1 },
			AfterFunc:  func() { data.TSint32NotIn1 = 5 },
		},
	}
	casesSint64 := []*CaseDesc{
		{
			Name:       fmt.Sprintf("test tag <%s>", protovalidator.TagIntEq),
			FieldDesc:  "field 't_sint64_eq1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagIntEq, Value: 1},
			FieldValue: func() interface{} { return data.TSint64Eq1 },
			BeforeFunc: func() {},
			AfterFunc:  func() { data.TSint64Eq1 = 1 },
		},
		{
			Name:       fmt.Sprintf("test tag <%s>", protovalidator.TagIntNe),
			FieldDesc:  "field 't_sint64_ne1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagIntNe, Value: 2},
			FieldValue: func() interface{} { return data.TSint64Ne1 },
			BeforeFunc: func() { data.TSint64Ne1 = 2 },
			AfterFunc:  func() { data.TSint64Ne1 = 1 },
		},
		{
			Name:       fmt.Sprintf("test tag <%s>", protovalidator.TagIntLt),
			FieldDesc:  "field 't_sint64_lt1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagIntLt, Value: 3},
			FieldValue: func() interface{} { return data.TSint64Lt1 },
			BeforeFunc: func() { data.TSint64Lt1 = 3 },
			AfterFunc:  func() { data.TSint64Lt1 = 1 },
		},
		{
			Name:       fmt.Sprintf("test tag <%s>", protovalidator.TagIntGt),
			FieldDesc:  "field 't_sint64_gt1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagIntGt, Value: 4},
			FieldValue: func() interface{} { return data.TSint64Gt1 },
			BeforeFunc: func() {},
			AfterFunc:  func() { data.TSint64Gt1 = 5 },
		},
		{
			Name:       fmt.Sprintf("test tag <%s>", protovalidator.TagIntLte),
			FieldDesc:  "field 't_sint64_lte1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagIntLte, Value: 5},
			FieldValue: func() interface{} { return data.TSint64Lte1 },
			BeforeFunc: func() { data.TSint64Lte1 = 6 },
			AfterFunc:  func() { data.TSint64Lte1 = 5 },
		},
		{
			Name:       fmt.Sprintf("test tag <%s>", protovalidator.TagIntGte),
			FieldDesc:  "field 't_sint64_gte1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagIntGte, Value: 6},
			FieldValue: func() interface{} { return data.TSint64Gte1 },
			BeforeFunc: func() {},
			AfterFunc:  func() { data.TSint64Gte1 = 6 },
		},
		{
			Name:       fmt.Sprintf("test tag <%s>", protovalidator.TagIntIn),
			FieldDesc:  "field 't_sint64_in1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagIntIn, Value: []int32{1, 2, 3}},
			FieldValue: func() interface{} { return data.TSint64In1 },
			BeforeFunc: func() { data.TSint64In1 = 4 },
			AfterFunc:  func() { data.TSint64In1 = 1 },
		},
		{
			Name:       fmt.Sprintf("test tag <%s>", protovalidator.TagIntNotIn),
			FieldDesc:  "field 't_sint64_not_in1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagIntNotIn, Value: []int32{1, 2, 3}},
			FieldValue: func() interface{} { return data.TSint64NotIn1 },
			BeforeFunc: func() { data.TSint64NotIn1 = 1 },
			AfterFunc:  func() { data.TSint64NotIn1 = 5 },
		},
	}
	casesSfixed32 := []*CaseDesc{
		{
			Name:       fmt.Sprintf("test tag <%s>", protovalidator.TagIntEq),
			FieldDesc:  "field 't_sfixed32_eq1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagIntEq, Value: 1},
			FieldValue: func() interface{} { return data.TSfixed32Eq1 },
			BeforeFunc: func() {},
			AfterFunc:  func() { data.TSfixed32Eq1 = 1 },
		},
		{
			Name:       fmt.Sprintf("test tag <%s>", protovalidator.TagIntNe),
			FieldDesc:  "field 't_sfixed32_ne1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagIntNe, Value: 2},
			FieldValue: func() interface{} { return data.TSfixed32Ne1 },
			BeforeFunc: func() { data.TSfixed32Ne1 = 2 },
			AfterFunc:  func() { data.TSfixed32Ne1 = 1 },
		},
		{
			Name:       fmt.Sprintf("test tag <%s>", protovalidator.TagIntLt),
			FieldDesc:  "field 't_sfixed32_lt1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagIntLt, Value: 3},
			FieldValue: func() interface{} { return data.TSfixed32Lt1 },
			BeforeFunc: func() { data.TSfixed32Lt1 = 3 },
			AfterFunc:  func() { data.TSfixed32Lt1 = 1 },
		},
		{
			Name:       fmt.Sprintf("test tag <%s>", protovalidator.TagIntGt),
			FieldDesc:  "field 't_sfixed32_gt1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagIntGt, Value: 4},
			FieldValue: func() interface{} { return data.TSfixed32Gt1 },
			BeforeFunc: func() {},
			AfterFunc:  func() { data.TSfixed32Gt1 = 5 },
		},
		{
			Name:       fmt.Sprintf("test tag <%s>", protovalidator.TagIntLte),
			FieldDesc:  "field 't_sfixed32_lte1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagIntLte, Value: 5},
			FieldValue: func() interface{} { return data.TSfixed32Lte1 },
			BeforeFunc: func() { data.TSfixed32Lte1 = 6 },
			AfterFunc:  func() { data.TSfixed32Lte1 = 5 },
		},
		{
			Name:       fmt.Sprintf("test tag <%s>", protovalidator.TagIntGte),
			FieldDesc:  "field 't_sfixed32_gte1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagIntGte, Value: 6},
			FieldValue: func() interface{} { return data.TSfixed32Gte1 },
			BeforeFunc: func() {},
			AfterFunc:  func() { data.TSfixed32Gte1 = 6 },
		},
		{
			Name:       fmt.Sprintf("test tag <%s>", protovalidator.TagIntIn),
			FieldDesc:  "field 't_sfixed32_in1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagIntIn, Value: []int32{1, 2, 3}},
			FieldValue: func() interface{} { return data.TSfixed32In1 },
			BeforeFunc: func() { data.TSfixed32In1 = 4 },
			AfterFunc:  func() { data.TSfixed32In1 = 1 },
		},
		{
			Name:       fmt.Sprintf("test tag <%s>", protovalidator.TagIntNotIn),
			FieldDesc:  "field 't_sfixed32_not_in1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagIntNotIn, Value: []int32{1, 2, 3}},
			FieldValue: func() interface{} { return data.TSfixed32NotIn1 },
			BeforeFunc: func() { data.TSfixed32NotIn1 = 1 },
			AfterFunc:  func() { data.TSfixed32NotIn1 = 5 },
		},
	}
	casesSfixed64 := []*CaseDesc{
		{
			Name:       fmt.Sprintf("test tag <%s>", protovalidator.TagIntEq),
			FieldDesc:  "field 't_sfixed64_eq1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagIntEq, Value: 1},
			FieldValue: func() interface{} { return data.TSfixed64Eq1 },
			BeforeFunc: func() {},
			AfterFunc:  func() { data.TSfixed64Eq1 = 1 },
		},
		{
			Name:       fmt.Sprintf("test tag <%s>", protovalidator.TagIntNe),
			FieldDesc:  "field 't_sfixed64_ne1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagIntNe, Value: 2},
			FieldValue: func() interface{} { return data.TSfixed64Ne1 },
			BeforeFunc: func() { data.TSfixed64Ne1 = 2 },
			AfterFunc:  func() { data.TSfixed64Ne1 = 1 },
		},
		{
			Name:       fmt.Sprintf("test tag <%s>", protovalidator.TagIntLt),
			FieldDesc:  "field 't_sfixed64_lt1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagIntLt, Value: 3},
			FieldValue: func() interface{} { return data.TSfixed64Lt1 },
			BeforeFunc: func() { data.TSfixed64Lt1 = 3 },
			AfterFunc:  func() { data.TSfixed64Lt1 = 1 },
		},
		{
			Name:       fmt.Sprintf("test tag <%s>", protovalidator.TagIntGt),
			FieldDesc:  "field 't_sfixed64_gt1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagIntGt, Value: 4},
			FieldValue: func() interface{} { return data.TSfixed64Gt1 },
			BeforeFunc: func() {},
			AfterFunc:  func() { data.TSfixed64Gt1 = 5 },
		},
		{
			Name:       fmt.Sprintf("test tag <%s>", protovalidator.TagIntLte),
			FieldDesc:  "field 't_sfixed64_lte1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagIntLte, Value: 5},
			FieldValue: func() interface{} { return data.TSfixed64Lte1 },
			BeforeFunc: func() { data.TSfixed64Lte1 = 6 },
			AfterFunc:  func() { data.TSfixed64Lte1 = 5 },
		},
		{
			Name:       fmt.Sprintf("test tag <%s>", protovalidator.TagIntGte),
			FieldDesc:  "field 't_sfixed64_gte1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagIntGte, Value: 6},
			FieldValue: func() interface{} { return data.TSfixed64Gte1 },
			BeforeFunc: func() {},
			AfterFunc:  func() { data.TSfixed64Gte1 = 6 },
		},
		{
			Name:       fmt.Sprintf("test tag <%s>", protovalidator.TagIntIn),
			FieldDesc:  "field 't_sfixed64_in1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagIntIn, Value: []int32{1, 2, 3}},
			FieldValue: func() interface{} { return data.TSfixed64In1 },
			BeforeFunc: func() { data.TSfixed64In1 = 4 },
			AfterFunc:  func() { data.TSfixed64In1 = 1 },
		},
		{
			Name:       fmt.Sprintf("test tag <%s>", protovalidator.TagIntNotIn),
			FieldDesc:  "field 't_sfixed64_not_in1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagIntNotIn, Value: []int32{1, 2, 3}},
			FieldValue: func() interface{} { return data.TSfixed64NotIn1 },
			BeforeFunc: func() { data.TSfixed64NotIn1 = 1 },
			AfterFunc:  func() { data.TSfixed64NotIn1 = 5 },
		},
	}

	var cases []*CaseDesc
	cases = append(cases, casesInt32...)
	cases = append(cases, casesInt64...)
	cases = append(cases, casesSint32...)
	cases = append(cases, casesSint64...)
	cases = append(cases, casesSfixed32...)
	cases = append(cases, casesSfixed64...)

	msgName := "ValidIntTagsGeneral1"
	runCases(t, data, msgName, cases)

	{
		err := data.Validate()
		require.Nil(t, err)
	}
}

func Test_GoValidator_ValidIntTagsOptional1(t *testing.T) {
	data := &govalidatortest.ValidIntTagsOptional1{}
	{
		err := data.Validate()
		require.NotNil(t, err)
	}

	casesInt32 := []*CaseDesc{
		{
			Name:       fmt.Sprintf("test tag <%s>", protovalidator.TagIntEq),
			FieldDesc:  "field 't_int32_eq1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagIntEq, Value: 1},
			FieldValue: func() interface{} { return nil },
			BeforeFunc: func() {},
			AfterFunc:  func() { x := int32(1); data.TInt32Eq1 = &x },
		},
		{
			Name:       fmt.Sprintf("test tag <%s>", protovalidator.TagIntNe),
			FieldDesc:  "field 't_int32_ne1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagIntNe, Value: 2},
			FieldValue: func() interface{} { return data.TInt32Ne1 },
			BeforeFunc: func() { x := int32(2); data.TInt32Ne1 = &x },
			AfterFunc:  func() { x := int32(1); data.TInt32Ne1 = &x },
		},
		{
			Name:       fmt.Sprintf("test tag <%s>", protovalidator.TagIntLt),
			FieldDesc:  "field 't_int32_lt1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagIntLt, Value: 3},
			FieldValue: func() interface{} { return data.TInt32Lt1 },
			BeforeFunc: func() { x := int32(3); data.TInt32Lt1 = &x },
			AfterFunc:  func() { x := int32(1); data.TInt32Lt1 = &x },
		},
		{
			Name:       fmt.Sprintf("test tag <%s>", protovalidator.TagIntGt),
			FieldDesc:  "field 't_int32_gt1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagIntGt, Value: 4},
			FieldValue: func() interface{} { return nil },
			BeforeFunc: func() {},
			AfterFunc:  func() { x := int32(5); data.TInt32Gt1 = &x },
		},
		{
			Name:       fmt.Sprintf("test tag <%s>", protovalidator.TagIntLte),
			FieldDesc:  "field 't_int32_lte1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagIntLte, Value: 5},
			FieldValue: func() interface{} { return data.TInt32Lte1 },
			BeforeFunc: func() { x := int32(6); data.TInt32Lte1 = &x },
			AfterFunc:  func() { x := int32(5); data.TInt32Lte1 = &x },
		},
		{
			Name:       fmt.Sprintf("test tag <%s>", protovalidator.TagIntGte),
			FieldDesc:  "field 't_int32_gte1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagIntGte, Value: 6},
			FieldValue: func() interface{} { return nil },
			BeforeFunc: func() {},
			AfterFunc:  func() { x := int32(6); data.TInt32Gte1 = &x },
		},
		{
			Name:       fmt.Sprintf("test tag <%s>", protovalidator.TagIntIn),
			FieldDesc:  "field 't_int32_in1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagIntIn, Value: []int32{1, 2, 3}},
			FieldValue: func() interface{} { return data.TInt32In1 },
			BeforeFunc: func() { x := int32(4); data.TInt32In1 = &x },
			AfterFunc:  func() { x := int32(1); data.TInt32In1 = &x },
		},
		{
			Name:       fmt.Sprintf("test tag <%s>", protovalidator.TagIntNotIn),
			FieldDesc:  "field 't_int32_not_in1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagIntNotIn, Value: []int32{1, 2, 3}},
			FieldValue: func() interface{} { return data.TInt32NotIn1 },
			BeforeFunc: func() { x := int32(1); data.TInt32NotIn1 = &x },
			AfterFunc:  func() { x := int32(5); data.TInt32NotIn1 = &x },
		},
	}
	casesInt64 := []*CaseDesc{
		{
			Name:       fmt.Sprintf("test tag <%s>", protovalidator.TagIntEq),
			FieldDesc:  "field 't_int64_eq1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagIntEq, Value: 1},
			FieldValue: func() interface{} { return nil },
			BeforeFunc: func() {},
			AfterFunc:  func() { x := int64(1); data.TInt64Eq1 = &x },
		},
		{
			Name:       fmt.Sprintf("test tag <%s>", protovalidator.TagIntNe),
			FieldDesc:  "field 't_int64_ne1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagIntNe, Value: 2},
			FieldValue: func() interface{} { return data.TInt64Ne1 },
			BeforeFunc: func() { x := int64(2); data.TInt64Ne1 = &x },
			AfterFunc:  func() { x := int64(1); data.TInt64Ne1 = &x },
		},
		{
			Name:       fmt.Sprintf("test tag <%s>", protovalidator.TagIntLt),
			FieldDesc:  "field 't_int64_lt1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagIntLt, Value: 3},
			FieldValue: func() interface{} { return data.TInt64Lt1 },
			BeforeFunc: func() { x := int64(3); data.TInt64Lt1 = &x },
			AfterFunc:  func() { x := int64(1); data.TInt64Lt1 = &x },
		},
		{
			Name:       fmt.Sprintf("test tag <%s>", protovalidator.TagIntGt),
			FieldDesc:  "field 't_int64_gt1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagIntGt, Value: 4},
			FieldValue: func() interface{} { return nil },
			BeforeFunc: func() {},
			AfterFunc:  func() { x := int64(5); data.TInt64Gt1 = &x },
		},
		{
			Name:       fmt.Sprintf("test tag <%s>", protovalidator.TagIntLte),
			FieldDesc:  "field 't_int64_lte1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagIntLte, Value: 5},
			FieldValue: func() interface{} { return data.TInt64Lte1 },
			BeforeFunc: func() { x := int64(6); data.TInt64Lte1 = &x },
			AfterFunc:  func() { x := int64(5); data.TInt64Lte1 = &x },
		},
		{
			Name:       fmt.Sprintf("test tag <%s>", protovalidator.TagIntGte),
			FieldDesc:  "field 't_int64_gte1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagIntGte, Value: 6},
			FieldValue: func() interface{} { return nil },
			BeforeFunc: func() {},
			AfterFunc:  func() { x := int64(6); data.TInt64Gte1 = &x },
		},
		{
			Name:       fmt.Sprintf("test tag <%s>", protovalidator.TagIntIn),
			FieldDesc:  "field 't_int64_in1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagIntIn, Value: []int32{1, 2, 3}},
			FieldValue: func() interface{} { return data.TInt64In1 },
			BeforeFunc: func() { x := int64(4); data.TInt64In1 = &x },
			AfterFunc:  func() { x := int64(1); data.TInt64In1 = &x },
		},
		{
			Name:       fmt.Sprintf("test tag <%s>", protovalidator.TagIntNotIn),
			FieldDesc:  "field 't_int64_not_in1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagIntNotIn, Value: []int32{1, 2, 3}},
			FieldValue: func() interface{} { return data.TInt64NotIn1 },
			BeforeFunc: func() { x := int64(1); data.TInt64NotIn1 = &x },
			AfterFunc:  func() { x := int64(5); data.TInt64NotIn1 = &x },
		},
	}
	casesSint32 := []*CaseDesc{
		{
			Name:       fmt.Sprintf("test tag <%s>", protovalidator.TagIntEq),
			FieldDesc:  "field 't_sint32_eq1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagIntEq, Value: 1},
			FieldValue: func() interface{} { return nil },
			BeforeFunc: func() {},
			AfterFunc:  func() { x := int32(1); data.TSint32Eq1 = &x },
		},
		{
			Name:       fmt.Sprintf("test tag <%s>", protovalidator.TagIntNe),
			FieldDesc:  "field 't_sint32_ne1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagIntNe, Value: 2},
			FieldValue: func() interface{} { return data.TSint32Ne1 },
			BeforeFunc: func() { x := int32(2); data.TSint32Ne1 = &x },
			AfterFunc:  func() { x := int32(1); data.TSint32Ne1 = &x },
		},
		{
			Name:       fmt.Sprintf("test tag <%s>", protovalidator.TagIntLt),
			FieldDesc:  "field 't_sint32_lt1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagIntLt, Value: 3},
			FieldValue: func() interface{} { return data.TSint32Lt1 },
			BeforeFunc: func() { x := int32(3); data.TSint32Lt1 = &x },
			AfterFunc:  func() { x := int32(1); data.TSint32Lt1 = &x },
		},
		{
			Name:       fmt.Sprintf("test tag <%s>", protovalidator.TagIntGt),
			FieldDesc:  "field 't_sint32_gt1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagIntGt, Value: 4},
			FieldValue: func() interface{} { return nil },
			BeforeFunc: func() {},
			AfterFunc:  func() { x := int32(5); data.TSint32Gt1 = &x },
		},
		{
			Name:       fmt.Sprintf("test tag <%s>", protovalidator.TagIntLte),
			FieldDesc:  "field 't_sint32_lte1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagIntLte, Value: 5},
			FieldValue: func() interface{} { return data.TSint32Lte1 },
			BeforeFunc: func() { x := int32(6); data.TSint32Lte1 = &x },
			AfterFunc:  func() { x := int32(5); data.TSint32Lte1 = &x },
		},
		{
			Name:       fmt.Sprintf("test tag <%s>", protovalidator.TagIntGte),
			FieldDesc:  "field 't_sint32_gte1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagIntGte, Value: 6},
			FieldValue: func() interface{} { return nil },
			BeforeFunc: func() {},
			AfterFunc:  func() { x := int32(6); data.TSint32Gte1 = &x },
		},
		{
			Name:       fmt.Sprintf("test tag <%s>", protovalidator.TagIntIn),
			FieldDesc:  "field 't_sint32_in1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagIntIn, Value: []int32{1, 2, 3}},
			FieldValue: func() interface{} { return data.TSint32In1 },
			BeforeFunc: func() { x := int32(4); data.TSint32In1 = &x },
			AfterFunc:  func() { x := int32(1); data.TSint32In1 = &x },
		},
		{
			Name:       fmt.Sprintf("test tag <%s>", protovalidator.TagIntNotIn),
			FieldDesc:  "field 't_sint32_not_in1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagIntNotIn, Value: []int32{1, 2, 3}},
			FieldValue: func() interface{} { return data.TSint32NotIn1 },
			BeforeFunc: func() { x := int32(1); data.TSint32NotIn1 = &x },
			AfterFunc:  func() { x := int32(5); data.TSint32NotIn1 = &x },
		},
	}
	casesSint64 := []*CaseDesc{
		{
			Name:       fmt.Sprintf("test tag <%s>", protovalidator.TagIntEq),
			FieldDesc:  "field 't_sint64_eq1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagIntEq, Value: 1},
			FieldValue: func() interface{} { return nil },
			BeforeFunc: func() {},
			AfterFunc:  func() { x := int64(1); data.TSint64Eq1 = &x },
		},
		{
			Name:       fmt.Sprintf("test tag <%s>", protovalidator.TagIntNe),
			FieldDesc:  "field 't_sint64_ne1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagIntNe, Value: 2},
			FieldValue: func() interface{} { return data.TSint64Ne1 },
			BeforeFunc: func() { x := int64(2); data.TSint64Ne1 = &x },
			AfterFunc:  func() { x := int64(1); data.TSint64Ne1 = &x },
		},
		{
			Name:       fmt.Sprintf("test tag <%s>", protovalidator.TagIntLt),
			FieldDesc:  "field 't_sint64_lt1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagIntLt, Value: 3},
			FieldValue: func() interface{} { return data.TSint64Lt1 },
			BeforeFunc: func() { x := int64(3); data.TSint64Lt1 = &x },
			AfterFunc:  func() { x := int64(1); data.TSint64Lt1 = &x },
		},
		{
			Name:       fmt.Sprintf("test tag <%s>", protovalidator.TagIntGt),
			FieldDesc:  "field 't_sint64_gt1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagIntGt, Value: 4},
			FieldValue: func() interface{} { return nil },
			BeforeFunc: func() {},
			AfterFunc:  func() { x := int64(5); data.TSint64Gt1 = &x },
		},
		{
			Name:       fmt.Sprintf("test tag <%s>", protovalidator.TagIntLte),
			FieldDesc:  "field 't_sint64_lte1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagIntLte, Value: 5},
			FieldValue: func() interface{} { return data.TSint64Lte1 },
			BeforeFunc: func() { x := int64(6); data.TSint64Lte1 = &x },
			AfterFunc:  func() { x := int64(5); data.TSint64Lte1 = &x },
		},
		{
			Name:       fmt.Sprintf("test tag <%s>", protovalidator.TagIntGte),
			FieldDesc:  "field 't_sint64_gte1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagIntGte, Value: 6},
			FieldValue: func() interface{} { return nil },
			BeforeFunc: func() {},
			AfterFunc:  func() { x := int64(6); data.TSint64Gte1 = &x },
		},
		{
			Name:       fmt.Sprintf("test tag <%s>", protovalidator.TagIntIn),
			FieldDesc:  "field 't_sint64_in1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagIntIn, Value: []int32{1, 2, 3}},
			FieldValue: func() interface{} { return data.TSint64In1 },
			BeforeFunc: func() { x := int64(4); data.TSint64In1 = &x },
			AfterFunc:  func() { x := int64(1); data.TSint64In1 = &x },
		},
		{
			Name:       fmt.Sprintf("test tag <%s>", protovalidator.TagIntNotIn),
			FieldDesc:  "field 't_sint64_not_in1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagIntNotIn, Value: []int32{1, 2, 3}},
			FieldValue: func() interface{} { return data.TSint64NotIn1 },
			BeforeFunc: func() { x := int64(1); data.TSint64NotIn1 = &x },
			AfterFunc:  func() { x := int64(5); data.TSint64NotIn1 = &x },
		},
	}
	casesSfixed32 := []*CaseDesc{
		{
			Name:       fmt.Sprintf("test tag <%s>", protovalidator.TagIntEq),
			FieldDesc:  "field 't_sfixed32_eq1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagIntEq, Value: 1},
			FieldValue: func() interface{} { return nil },
			BeforeFunc: func() {},
			AfterFunc:  func() { x := int32(1); data.TSfixed32Eq1 = &x },
		},
		{
			Name:       fmt.Sprintf("test tag <%s>", protovalidator.TagIntNe),
			FieldDesc:  "field 't_sfixed32_ne1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagIntNe, Value: 2},
			FieldValue: func() interface{} { return data.TSfixed32Ne1 },
			BeforeFunc: func() { x := int32(2); data.TSfixed32Ne1 = &x },
			AfterFunc:  func() { x := int32(1); data.TSfixed32Ne1 = &x },
		},
		{
			Name:       fmt.Sprintf("test tag <%s>", protovalidator.TagIntLt),
			FieldDesc:  "field 't_sfixed32_lt1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagIntLt, Value: 3},
			FieldValue: func() interface{} { return data.TSfixed32Lt1 },
			BeforeFunc: func() { x := int32(3); data.TSfixed32Lt1 = &x },
			AfterFunc:  func() { x := int32(1); data.TSfixed32Lt1 = &x },
		},
		{
			Name:       fmt.Sprintf("test tag <%s>", protovalidator.TagIntGt),
			FieldDesc:  "field 't_sfixed32_gt1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagIntGt, Value: 4},
			FieldValue: func() interface{} { return nil },
			BeforeFunc: func() {},
			AfterFunc:  func() { x := int32(5); data.TSfixed32Gt1 = &x },
		},
		{
			Name:       fmt.Sprintf("test tag <%s>", protovalidator.TagIntLte),
			FieldDesc:  "field 't_sfixed32_lte1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagIntLte, Value: 5},
			FieldValue: func() interface{} { return data.TSfixed32Lte1 },
			BeforeFunc: func() { x := int32(6); data.TSfixed32Lte1 = &x },
			AfterFunc:  func() { x := int32(5); data.TSfixed32Lte1 = &x },
		},
		{
			Name:       fmt.Sprintf("test tag <%s>", protovalidator.TagIntGte),
			FieldDesc:  "field 't_sfixed32_gte1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagIntGte, Value: 6},
			FieldValue: func() interface{} { return data.TSfixed32Gte1 },
			BeforeFunc: func() {},
			AfterFunc:  func() { x := int32(6); data.TSfixed32Gte1 = &x },
		},
		{
			Name:       fmt.Sprintf("test tag <%s>", protovalidator.TagIntIn),
			FieldDesc:  "field 't_sfixed32_in1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagIntIn, Value: []int32{1, 2, 3}},
			FieldValue: func() interface{} { return data.TSfixed32In1 },
			BeforeFunc: func() { x := int32(4); data.TSfixed32In1 = &x },
			AfterFunc:  func() { x := int32(1); data.TSfixed32In1 = &x },
		},
		{
			Name:       fmt.Sprintf("test tag <%s>", protovalidator.TagIntNotIn),
			FieldDesc:  "field 't_sfixed32_not_in1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagIntNotIn, Value: []int32{1, 2, 3}},
			FieldValue: func() interface{} { return data.TSfixed32NotIn1 },
			BeforeFunc: func() { x := int32(1); data.TSfixed32NotIn1 = &x },
			AfterFunc:  func() { x := int32(5); data.TSfixed32NotIn1 = &x },
		},
	}
	casesSfixed64 := []*CaseDesc{
		{
			Name:       fmt.Sprintf("test tag <%s>", protovalidator.TagIntEq),
			FieldDesc:  "field 't_sfixed64_eq1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagIntEq, Value: 1},
			FieldValue: func() interface{} { return nil },
			BeforeFunc: func() {},
			AfterFunc:  func() { x := int64(1); data.TSfixed64Eq1 = &x },
		},
		{
			Name:       fmt.Sprintf("test tag <%s>", protovalidator.TagIntNe),
			FieldDesc:  "field 't_sfixed64_ne1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagIntNe, Value: 2},
			FieldValue: func() interface{} { return data.TSfixed64Ne1 },
			BeforeFunc: func() { x := int64(2); data.TSfixed64Ne1 = &x },
			AfterFunc:  func() { x := int64(1); data.TSfixed64Ne1 = &x },
		},
		{
			Name:       fmt.Sprintf("test tag <%s>", protovalidator.TagIntLt),
			FieldDesc:  "field 't_sfixed64_lt1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagIntLt, Value: 3},
			FieldValue: func() interface{} { return data.TSfixed64Lt1 },
			BeforeFunc: func() { x := int64(3); data.TSfixed64Lt1 = &x },
			AfterFunc:  func() { x := int64(1); data.TSfixed64Lt1 = &x },
		},
		{
			Name:       fmt.Sprintf("test tag <%s>", protovalidator.TagIntGt),
			FieldDesc:  "field 't_sfixed64_gt1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagIntGt, Value: 4},
			FieldValue: func() interface{} { return nil },
			BeforeFunc: func() {},
			AfterFunc:  func() { x := int64(5); data.TSfixed64Gt1 = &x },
		},
		{
			Name:       fmt.Sprintf("test tag <%s>", protovalidator.TagIntLte),
			FieldDesc:  "field 't_sfixed64_lte1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagIntLte, Value: 5},
			FieldValue: func() interface{} { return data.TSfixed64Lte1 },
			BeforeFunc: func() { x := int64(6); data.TSfixed64Lte1 = &x },
			AfterFunc:  func() { x := int64(5); data.TSfixed64Lte1 = &x },
		},
		{
			Name:       fmt.Sprintf("test tag <%s>", protovalidator.TagIntGte),
			FieldDesc:  "field 't_sfixed64_gte1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagIntGte, Value: 6},
			FieldValue: func() interface{} { return nil },
			BeforeFunc: func() {},
			AfterFunc:  func() { x := int64(6); data.TSfixed64Gte1 = &x },
		},
		{
			Name:       fmt.Sprintf("test tag <%s>", protovalidator.TagIntIn),
			FieldDesc:  "field 't_sfixed64_in1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagIntIn, Value: []int32{1, 2, 3}},
			FieldValue: func() interface{} { return data.TSfixed64In1 },
			BeforeFunc: func() { x := int64(4); data.TSfixed64In1 = &x },
			AfterFunc:  func() { x := int64(1); data.TSfixed64In1 = &x },
		},
		{
			Name:       fmt.Sprintf("test tag <%s>", protovalidator.TagIntNotIn),
			FieldDesc:  "field 't_sfixed64_not_in1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagIntNotIn, Value: []int32{1, 2, 3}},
			FieldValue: func() interface{} { return data.TSfixed64NotIn1 },
			BeforeFunc: func() { x := int64(1); data.TSfixed64NotIn1 = &x },
			AfterFunc:  func() { x := int64(5); data.TSfixed64NotIn1 = &x },
		},
	}

	var cases []*CaseDesc
	cases = append(cases, casesInt32...)
	cases = append(cases, casesInt64...)
	cases = append(cases, casesSint32...)
	cases = append(cases, casesSint64...)
	cases = append(cases, casesSfixed32...)
	cases = append(cases, casesSfixed64...)

	msgName := "ValidIntTagsOptional1"
	runCases(t, data, msgName, cases)

	{
		err := data.Validate()
		require.Nil(t, err)
	}
}

func Test_GoValidator_ValidIntTagsOneOf1(t *testing.T) {
	data := &govalidatortest.ValidIntTagsOneOf1{}
	{
		err := data.Validate()
		require.Nil(t, err)
	}

	casesInt32 := []*CaseDesc{
		{
			Name:       fmt.Sprintf("test tag <%s>", protovalidator.TagIntEq),
			FieldDesc:  "field 't_int32_eq1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagIntEq, Value: 1},
			FieldValue: func() interface{} { return data.OneTyp1.(*govalidatortest.ValidIntTagsOneOf1_TInt32Eq1).TInt32Eq1 },
			BeforeFunc: func() { data.OneTyp1 = &govalidatortest.ValidIntTagsOneOf1_TInt32Eq1{} },
			AfterFunc:  func() { data.OneTyp1 = &govalidatortest.ValidIntTagsOneOf1_TInt32Eq1{TInt32Eq1: 1} },
		},
		{
			Name:       fmt.Sprintf("test tag <%s>", protovalidator.TagIntNe),
			FieldDesc:  "field 't_int32_ne1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagIntNe, Value: 2},
			FieldValue: func() interface{} { return data.OneTyp1.(*govalidatortest.ValidIntTagsOneOf1_TInt32Ne1).TInt32Ne1 },
			BeforeFunc: func() { data.OneTyp1 = &govalidatortest.ValidIntTagsOneOf1_TInt32Ne1{TInt32Ne1: 2} },
			AfterFunc:  func() { data.OneTyp1 = &govalidatortest.ValidIntTagsOneOf1_TInt32Ne1{TInt32Ne1: 1} },
		},
		{
			Name:       fmt.Sprintf("test tag <%s>", protovalidator.TagIntLt),
			FieldDesc:  "field 't_int32_lt1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagIntLt, Value: 3},
			FieldValue: func() interface{} { return data.OneTyp1.(*govalidatortest.ValidIntTagsOneOf1_TInt32Lt1).TInt32Lt1 },
			BeforeFunc: func() { data.OneTyp1 = &govalidatortest.ValidIntTagsOneOf1_TInt32Lt1{TInt32Lt1: 3} },
			AfterFunc:  func() { data.OneTyp1 = &govalidatortest.ValidIntTagsOneOf1_TInt32Lt1{TInt32Lt1: 1} },
		},
		{
			Name:       fmt.Sprintf("test tag <%s>", protovalidator.TagIntGt),
			FieldDesc:  "field 't_int32_gt1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagIntGt, Value: 4},
			FieldValue: func() interface{} { return data.OneTyp1.(*govalidatortest.ValidIntTagsOneOf1_TInt32Gt1).TInt32Gt1 },
			BeforeFunc: func() { data.OneTyp1 = &govalidatortest.ValidIntTagsOneOf1_TInt32Gt1{} },
			AfterFunc:  func() { data.OneTyp1 = &govalidatortest.ValidIntTagsOneOf1_TInt32Gt1{TInt32Gt1: 5} },
		},
		{
			Name:       fmt.Sprintf("test tag <%s>", protovalidator.TagIntLte),
			FieldDesc:  "field 't_int32_lte1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagIntLte, Value: 5},
			FieldValue: func() interface{} { return data.OneTyp1.(*govalidatortest.ValidIntTagsOneOf1_TInt32Lte1).TInt32Lte1 },
			BeforeFunc: func() { data.OneTyp1 = &govalidatortest.ValidIntTagsOneOf1_TInt32Lte1{TInt32Lte1: 6} },
			AfterFunc:  func() { data.OneTyp1 = &govalidatortest.ValidIntTagsOneOf1_TInt32Lte1{TInt32Lte1: 5} },
		},
		{
			Name:       fmt.Sprintf("test tag <%s>", protovalidator.TagIntGte),
			FieldDesc:  "field 't_int32_gte1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagIntGte, Value: 6},
			FieldValue: func() interface{} { return data.OneTyp1.(*govalidatortest.ValidIntTagsOneOf1_TInt32Gte1).TInt32Gte1 },
			BeforeFunc: func() { data.OneTyp1 = &govalidatortest.ValidIntTagsOneOf1_TInt32Gte1{} },
			AfterFunc:  func() { data.OneTyp1 = &govalidatortest.ValidIntTagsOneOf1_TInt32Gte1{TInt32Gte1: 6} },
		},
		{
			Name:       fmt.Sprintf("test tag <%s>", protovalidator.TagIntIn),
			FieldDesc:  "field 't_int32_in1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagIntIn, Value: []int32{1, 2, 3}},
			FieldValue: func() interface{} { return data.OneTyp1.(*govalidatortest.ValidIntTagsOneOf1_TInt32In1).TInt32In1 },
			BeforeFunc: func() { data.OneTyp1 = &govalidatortest.ValidIntTagsOneOf1_TInt32In1{TInt32In1: 4} },
			AfterFunc:  func() { data.OneTyp1 = &govalidatortest.ValidIntTagsOneOf1_TInt32In1{TInt32In1: 1} },
		},
		{
			Name:      fmt.Sprintf("test tag <%s>", protovalidator.TagIntNotIn),
			FieldDesc: "field 't_int32_not_in1'",
			TagInfo:   &protovalidator.TagInfo{Tag: protovalidator.TagIntNotIn, Value: []int32{1, 2, 3}},
			FieldValue: func() interface{} {
				return data.OneTyp1.(*govalidatortest.ValidIntTagsOneOf1_TInt32NotIn1).TInt32NotIn1
			},
			BeforeFunc: func() { data.OneTyp1 = &govalidatortest.ValidIntTagsOneOf1_TInt32NotIn1{TInt32NotIn1: 1} },
			AfterFunc:  func() { data.OneTyp1 = &govalidatortest.ValidIntTagsOneOf1_TInt32NotIn1{TInt32NotIn1: 5} },
		},
	}
	casesInt64 := []*CaseDesc{
		{
			Name:       fmt.Sprintf("test tag <%s>", protovalidator.TagIntEq),
			FieldDesc:  "field 't_int64_eq1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagIntEq, Value: 1},
			FieldValue: func() interface{} { return data.OneTyp1.(*govalidatortest.ValidIntTagsOneOf1_TInt64Eq1).TInt64Eq1 },
			BeforeFunc: func() { data.OneTyp1 = &govalidatortest.ValidIntTagsOneOf1_TInt64Eq1{} },
			AfterFunc:  func() { data.OneTyp1 = &govalidatortest.ValidIntTagsOneOf1_TInt64Eq1{TInt64Eq1: 1} },
		},
		{
			Name:       fmt.Sprintf("test tag <%s>", protovalidator.TagIntNe),
			FieldDesc:  "field 't_int64_ne1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagIntNe, Value: 2},
			FieldValue: func() interface{} { return data.OneTyp1.(*govalidatortest.ValidIntTagsOneOf1_TInt64Ne1).TInt64Ne1 },
			BeforeFunc: func() { data.OneTyp1 = &govalidatortest.ValidIntTagsOneOf1_TInt64Ne1{TInt64Ne1: 2} },
			AfterFunc:  func() { data.OneTyp1 = &govalidatortest.ValidIntTagsOneOf1_TInt64Ne1{TInt64Ne1: 1} },
		},
		{
			Name:       fmt.Sprintf("test tag <%s>", protovalidator.TagIntLt),
			FieldDesc:  "field 't_int64_lt1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagIntLt, Value: 3},
			FieldValue: func() interface{} { return data.OneTyp1.(*govalidatortest.ValidIntTagsOneOf1_TInt64Lt1).TInt64Lt1 },
			BeforeFunc: func() { data.OneTyp1 = &govalidatortest.ValidIntTagsOneOf1_TInt64Lt1{TInt64Lt1: 3} },
			AfterFunc:  func() { data.OneTyp1 = &govalidatortest.ValidIntTagsOneOf1_TInt64Lt1{TInt64Lt1: 1} },
		},
		{
			Name:       fmt.Sprintf("test tag <%s>", protovalidator.TagIntGt),
			FieldDesc:  "field 't_int64_gt1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagIntGt, Value: 4},
			FieldValue: func() interface{} { return data.OneTyp1.(*govalidatortest.ValidIntTagsOneOf1_TInt64Gt1).TInt64Gt1 },
			BeforeFunc: func() { data.OneTyp1 = &govalidatortest.ValidIntTagsOneOf1_TInt64Gt1{} },
			AfterFunc:  func() { data.OneTyp1 = &govalidatortest.ValidIntTagsOneOf1_TInt64Gt1{TInt64Gt1: 5} },
		},
		{
			Name:       fmt.Sprintf("test tag <%s>", protovalidator.TagIntLte),
			FieldDesc:  "field 't_int64_lte1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagIntLte, Value: 5},
			FieldValue: func() interface{} { return data.OneTyp1.(*govalidatortest.ValidIntTagsOneOf1_TInt64Lte1).TInt64Lte1 },
			BeforeFunc: func() { data.OneTyp1 = &govalidatortest.ValidIntTagsOneOf1_TInt64Lte1{TInt64Lte1: 6} },
			AfterFunc:  func() { data.OneTyp1 = &govalidatortest.ValidIntTagsOneOf1_TInt64Lte1{TInt64Lte1: 5} },
		},
		{
			Name:       fmt.Sprintf("test tag <%s>", protovalidator.TagIntGte),
			FieldDesc:  "field 't_int64_gte1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagIntGte, Value: 6},
			FieldValue: func() interface{} { return data.OneTyp1.(*govalidatortest.ValidIntTagsOneOf1_TInt64Gte1).TInt64Gte1 },
			BeforeFunc: func() { data.OneTyp1 = &govalidatortest.ValidIntTagsOneOf1_TInt64Gte1{} },
			AfterFunc:  func() { data.OneTyp1 = &govalidatortest.ValidIntTagsOneOf1_TInt64Gte1{TInt64Gte1: 6} },
		},
		{
			Name:       fmt.Sprintf("test tag <%s>", protovalidator.TagIntIn),
			FieldDesc:  "field 't_int64_in1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagIntIn, Value: []int32{1, 2, 3}},
			FieldValue: func() interface{} { return data.OneTyp1.(*govalidatortest.ValidIntTagsOneOf1_TInt64In1).TInt64In1 },
			BeforeFunc: func() { data.OneTyp1 = &govalidatortest.ValidIntTagsOneOf1_TInt64In1{TInt64In1: 4} },
			AfterFunc:  func() { data.OneTyp1 = &govalidatortest.ValidIntTagsOneOf1_TInt64In1{TInt64In1: 1} },
		},
		{
			Name:      fmt.Sprintf("test tag <%s>", protovalidator.TagIntNotIn),
			FieldDesc: "field 't_int64_not_in1'",
			TagInfo:   &protovalidator.TagInfo{Tag: protovalidator.TagIntNotIn, Value: []int32{1, 2, 3}},
			FieldValue: func() interface{} {
				return data.OneTyp1.(*govalidatortest.ValidIntTagsOneOf1_TInt64NotIn1).TInt64NotIn1
			},
			BeforeFunc: func() { data.OneTyp1 = &govalidatortest.ValidIntTagsOneOf1_TInt64NotIn1{TInt64NotIn1: 1} },
			AfterFunc:  func() { data.OneTyp1 = &govalidatortest.ValidIntTagsOneOf1_TInt64NotIn1{TInt64NotIn1: 5} },
		},
	}
	casesSint32 := []*CaseDesc{
		{
			Name:       fmt.Sprintf("test tag <%s>", protovalidator.TagIntEq),
			FieldDesc:  "field 't_sint32_eq1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagIntEq, Value: 1},
			FieldValue: func() interface{} { return data.OneTyp1.(*govalidatortest.ValidIntTagsOneOf1_TSint32Eq1).TSint32Eq1 },
			BeforeFunc: func() { data.OneTyp1 = &govalidatortest.ValidIntTagsOneOf1_TSint32Eq1{} },
			AfterFunc:  func() { data.OneTyp1 = &govalidatortest.ValidIntTagsOneOf1_TSint32Eq1{TSint32Eq1: 1} },
		},
		{
			Name:       fmt.Sprintf("test tag <%s>", protovalidator.TagIntNe),
			FieldDesc:  "field 't_sint32_ne1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagIntNe, Value: 2},
			FieldValue: func() interface{} { return data.OneTyp1.(*govalidatortest.ValidIntTagsOneOf1_TSint32Ne1).TSint32Ne1 },
			BeforeFunc: func() { data.OneTyp1 = &govalidatortest.ValidIntTagsOneOf1_TSint32Ne1{TSint32Ne1: 2} },
			AfterFunc:  func() { data.OneTyp1 = &govalidatortest.ValidIntTagsOneOf1_TSint32Ne1{TSint32Ne1: 1} },
		},
		{
			Name:       fmt.Sprintf("test tag <%s>", protovalidator.TagIntLt),
			FieldDesc:  "field 't_sint32_lt1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagIntLt, Value: 3},
			FieldValue: func() interface{} { return data.OneTyp1.(*govalidatortest.ValidIntTagsOneOf1_TSint32Lt1).TSint32Lt1 },
			BeforeFunc: func() { data.OneTyp1 = &govalidatortest.ValidIntTagsOneOf1_TSint32Lt1{TSint32Lt1: 3} },
			AfterFunc:  func() { data.OneTyp1 = &govalidatortest.ValidIntTagsOneOf1_TSint32Lt1{TSint32Lt1: 1} },
		},
		{
			Name:       fmt.Sprintf("test tag <%s>", protovalidator.TagIntGt),
			FieldDesc:  "field 't_sint32_gt1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagIntGt, Value: 4},
			FieldValue: func() interface{} { return data.OneTyp1.(*govalidatortest.ValidIntTagsOneOf1_TSint32Gt1).TSint32Gt1 },
			BeforeFunc: func() { data.OneTyp1 = &govalidatortest.ValidIntTagsOneOf1_TSint32Gt1{} },
			AfterFunc:  func() { data.OneTyp1 = &govalidatortest.ValidIntTagsOneOf1_TSint32Gt1{TSint32Gt1: 5} },
		},
		{
			Name:       fmt.Sprintf("test tag <%s>", protovalidator.TagIntLte),
			FieldDesc:  "field 't_sint32_lte1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagIntLte, Value: 5},
			FieldValue: func() interface{} { return data.OneTyp1.(*govalidatortest.ValidIntTagsOneOf1_TSint32Lte1).TSint32Lte1 },
			BeforeFunc: func() { data.OneTyp1 = &govalidatortest.ValidIntTagsOneOf1_TSint32Lte1{TSint32Lte1: 6} },
			AfterFunc:  func() { data.OneTyp1 = &govalidatortest.ValidIntTagsOneOf1_TSint32Lte1{TSint32Lte1: 5} },
		},
		{
			Name:       fmt.Sprintf("test tag <%s>", protovalidator.TagIntGte),
			FieldDesc:  "field 't_sint32_gte1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagIntGte, Value: 6},
			FieldValue: func() interface{} { return data.OneTyp1.(*govalidatortest.ValidIntTagsOneOf1_TSint32Gte1).TSint32Gte1 },
			BeforeFunc: func() { data.OneTyp1 = &govalidatortest.ValidIntTagsOneOf1_TSint32Gte1{} },
			AfterFunc:  func() { data.OneTyp1 = &govalidatortest.ValidIntTagsOneOf1_TSint32Gte1{TSint32Gte1: 6} },
		},
		{
			Name:       fmt.Sprintf("test tag <%s>", protovalidator.TagIntIn),
			FieldDesc:  "field 't_sint32_in1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagIntIn, Value: []int32{1, 2, 3}},
			FieldValue: func() interface{} { return data.OneTyp1.(*govalidatortest.ValidIntTagsOneOf1_TSint32In1).TSint32In1 },
			BeforeFunc: func() { data.OneTyp1 = &govalidatortest.ValidIntTagsOneOf1_TSint32In1{TSint32In1: 4} },
			AfterFunc:  func() { data.OneTyp1 = &govalidatortest.ValidIntTagsOneOf1_TSint32In1{TSint32In1: 1} },
		},
		{
			Name:      fmt.Sprintf("test tag <%s>", protovalidator.TagIntNotIn),
			FieldDesc: "field 't_sint32_not_in1'",
			TagInfo:   &protovalidator.TagInfo{Tag: protovalidator.TagIntNotIn, Value: []int32{1, 2, 3}},
			FieldValue: func() interface{} {
				return data.OneTyp1.(*govalidatortest.ValidIntTagsOneOf1_TSint32NotIn1).TSint32NotIn1
			},
			BeforeFunc: func() { data.OneTyp1 = &govalidatortest.ValidIntTagsOneOf1_TSint32NotIn1{TSint32NotIn1: 1} },
			AfterFunc:  func() { data.OneTyp1 = &govalidatortest.ValidIntTagsOneOf1_TSint32NotIn1{TSint32NotIn1: 5} },
		},
	}
	casesSint64 := []*CaseDesc{
		{
			Name:       fmt.Sprintf("test tag <%s>", protovalidator.TagIntEq),
			FieldDesc:  "field 't_sint64_eq1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagIntEq, Value: 1},
			FieldValue: func() interface{} { return data.OneTyp1.(*govalidatortest.ValidIntTagsOneOf1_TSint64Eq1).TSint64Eq1 },
			BeforeFunc: func() { data.OneTyp1 = &govalidatortest.ValidIntTagsOneOf1_TSint64Eq1{} },
			AfterFunc:  func() { data.OneTyp1 = &govalidatortest.ValidIntTagsOneOf1_TSint64Eq1{TSint64Eq1: 1} },
		},
		{
			Name:       fmt.Sprintf("test tag <%s>", protovalidator.TagIntNe),
			FieldDesc:  "field 't_sint64_ne1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagIntNe, Value: 2},
			FieldValue: func() interface{} { return data.OneTyp1.(*govalidatortest.ValidIntTagsOneOf1_TSint64Ne1).TSint64Ne1 },
			BeforeFunc: func() { data.OneTyp1 = &govalidatortest.ValidIntTagsOneOf1_TSint64Ne1{TSint64Ne1: 2} },
			AfterFunc:  func() { data.OneTyp1 = &govalidatortest.ValidIntTagsOneOf1_TSint64Ne1{TSint64Ne1: 1} },
		},
		{
			Name:       fmt.Sprintf("test tag <%s>", protovalidator.TagIntLt),
			FieldDesc:  "field 't_sint64_lt1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagIntLt, Value: 3},
			FieldValue: func() interface{} { return data.OneTyp1.(*govalidatortest.ValidIntTagsOneOf1_TSint64Lt1).TSint64Lt1 },
			BeforeFunc: func() { data.OneTyp1 = &govalidatortest.ValidIntTagsOneOf1_TSint64Lt1{TSint64Lt1: 3} },
			AfterFunc:  func() { data.OneTyp1 = &govalidatortest.ValidIntTagsOneOf1_TSint64Lt1{TSint64Lt1: 1} },
		},
		{
			Name:       fmt.Sprintf("test tag <%s>", protovalidator.TagIntGt),
			FieldDesc:  "field 't_sint64_gt1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagIntGt, Value: 4},
			FieldValue: func() interface{} { return data.OneTyp1.(*govalidatortest.ValidIntTagsOneOf1_TSint64Gt1).TSint64Gt1 },
			BeforeFunc: func() { data.OneTyp1 = &govalidatortest.ValidIntTagsOneOf1_TSint64Gt1{} },
			AfterFunc:  func() { data.OneTyp1 = &govalidatortest.ValidIntTagsOneOf1_TSint64Gt1{TSint64Gt1: 5} },
		},
		{
			Name:       fmt.Sprintf("test tag <%s>", protovalidator.TagIntLte),
			FieldDesc:  "field 't_sint64_lte1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagIntLte, Value: 5},
			FieldValue: func() interface{} { return data.OneTyp1.(*govalidatortest.ValidIntTagsOneOf1_TSint64Lte1).TSint64Lte1 },
			BeforeFunc: func() { data.OneTyp1 = &govalidatortest.ValidIntTagsOneOf1_TSint64Lte1{TSint64Lte1: 6} },
			AfterFunc:  func() { data.OneTyp1 = &govalidatortest.ValidIntTagsOneOf1_TSint64Lte1{TSint64Lte1: 5} },
		},
		{
			Name:       fmt.Sprintf("test tag <%s>", protovalidator.TagIntGte),
			FieldDesc:  "field 't_sint64_gte1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagIntGte, Value: 6},
			FieldValue: func() interface{} { return data.OneTyp1.(*govalidatortest.ValidIntTagsOneOf1_TSint64Gte1).TSint64Gte1 },
			BeforeFunc: func() { data.OneTyp1 = &govalidatortest.ValidIntTagsOneOf1_TSint64Gte1{} },
			AfterFunc:  func() { data.OneTyp1 = &govalidatortest.ValidIntTagsOneOf1_TSint64Gte1{TSint64Gte1: 6} },
		},
		{
			Name:       fmt.Sprintf("test tag <%s>", protovalidator.TagIntIn),
			FieldDesc:  "field 't_sint64_in1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagIntIn, Value: []int32{1, 2, 3}},
			FieldValue: func() interface{} { return data.OneTyp1.(*govalidatortest.ValidIntTagsOneOf1_TSint64In1).TSint64In1 },
			BeforeFunc: func() { data.OneTyp1 = &govalidatortest.ValidIntTagsOneOf1_TSint64In1{TSint64In1: 4} },
			AfterFunc:  func() { data.OneTyp1 = &govalidatortest.ValidIntTagsOneOf1_TSint64In1{TSint64In1: 1} },
		},
		{
			Name:      fmt.Sprintf("test tag <%s>", protovalidator.TagIntNotIn),
			FieldDesc: "field 't_sint64_not_in1'",
			TagInfo:   &protovalidator.TagInfo{Tag: protovalidator.TagIntNotIn, Value: []int32{1, 2, 3}},
			FieldValue: func() interface{} {
				return data.OneTyp1.(*govalidatortest.ValidIntTagsOneOf1_TSint64NotIn1).TSint64NotIn1
			},
			BeforeFunc: func() { data.OneTyp1 = &govalidatortest.ValidIntTagsOneOf1_TSint64NotIn1{TSint64NotIn1: 1} },
			AfterFunc:  func() { data.OneTyp1 = &govalidatortest.ValidIntTagsOneOf1_TSint64NotIn1{TSint64NotIn1: 5} },
		},
	}
	casesSfixed32 := []*CaseDesc{
		{
			Name:      fmt.Sprintf("test tag <%s>", protovalidator.TagIntEq),
			FieldDesc: "field 't_sfixed32_eq1'",
			TagInfo:   &protovalidator.TagInfo{Tag: protovalidator.TagIntEq, Value: 1},
			FieldValue: func() interface{} {
				return data.OneTyp1.(*govalidatortest.ValidIntTagsOneOf1_TSfixed32Eq1).TSfixed32Eq1
			},
			BeforeFunc: func() { data.OneTyp1 = &govalidatortest.ValidIntTagsOneOf1_TSfixed32Eq1{} },
			AfterFunc:  func() { data.OneTyp1 = &govalidatortest.ValidIntTagsOneOf1_TSfixed32Eq1{TSfixed32Eq1: 1} },
		},
		{
			Name:      fmt.Sprintf("test tag <%s>", protovalidator.TagIntNe),
			FieldDesc: "field 't_sfixed32_ne1'",
			TagInfo:   &protovalidator.TagInfo{Tag: protovalidator.TagIntNe, Value: 2},
			FieldValue: func() interface{} {
				return data.OneTyp1.(*govalidatortest.ValidIntTagsOneOf1_TSfixed32Ne1).TSfixed32Ne1
			},
			BeforeFunc: func() { data.OneTyp1 = &govalidatortest.ValidIntTagsOneOf1_TSfixed32Ne1{TSfixed32Ne1: 2} },
			AfterFunc:  func() { data.OneTyp1 = &govalidatortest.ValidIntTagsOneOf1_TSfixed32Ne1{TSfixed32Ne1: 1} },
		},
		{
			Name:      fmt.Sprintf("test tag <%s>", protovalidator.TagIntLt),
			FieldDesc: "field 't_sfixed32_lt1'",
			TagInfo:   &protovalidator.TagInfo{Tag: protovalidator.TagIntLt, Value: 3},
			FieldValue: func() interface{} {
				return data.OneTyp1.(*govalidatortest.ValidIntTagsOneOf1_TSfixed32Lt1).TSfixed32Lt1
			},
			BeforeFunc: func() { data.OneTyp1 = &govalidatortest.ValidIntTagsOneOf1_TSfixed32Lt1{TSfixed32Lt1: 3} },
			AfterFunc:  func() { data.OneTyp1 = &govalidatortest.ValidIntTagsOneOf1_TSfixed32Lt1{TSfixed32Lt1: 1} },
		},
		{
			Name:      fmt.Sprintf("test tag <%s>", protovalidator.TagIntGt),
			FieldDesc: "field 't_sfixed32_gt1'",
			TagInfo:   &protovalidator.TagInfo{Tag: protovalidator.TagIntGt, Value: 4},
			FieldValue: func() interface{} {
				return data.OneTyp1.(*govalidatortest.ValidIntTagsOneOf1_TSfixed32Gt1).TSfixed32Gt1
			},
			BeforeFunc: func() { data.OneTyp1 = &govalidatortest.ValidIntTagsOneOf1_TSfixed32Gt1{} },
			AfterFunc:  func() { data.OneTyp1 = &govalidatortest.ValidIntTagsOneOf1_TSfixed32Gt1{TSfixed32Gt1: 5} },
		},
		{
			Name:      fmt.Sprintf("test tag <%s>", protovalidator.TagIntLte),
			FieldDesc: "field 't_sfixed32_lte1'",
			TagInfo:   &protovalidator.TagInfo{Tag: protovalidator.TagIntLte, Value: 5},
			FieldValue: func() interface{} {
				return data.OneTyp1.(*govalidatortest.ValidIntTagsOneOf1_TSfixed32Lte1).TSfixed32Lte1
			},
			BeforeFunc: func() { data.OneTyp1 = &govalidatortest.ValidIntTagsOneOf1_TSfixed32Lte1{TSfixed32Lte1: 6} },
			AfterFunc:  func() { data.OneTyp1 = &govalidatortest.ValidIntTagsOneOf1_TSfixed32Lte1{TSfixed32Lte1: 5} },
		},
		{
			Name:      fmt.Sprintf("test tag <%s>", protovalidator.TagIntGte),
			FieldDesc: "field 't_sfixed32_gte1'",
			TagInfo:   &protovalidator.TagInfo{Tag: protovalidator.TagIntGte, Value: 6},
			FieldValue: func() interface{} {
				return data.OneTyp1.(*govalidatortest.ValidIntTagsOneOf1_TSfixed32Gte1).TSfixed32Gte1
			},
			BeforeFunc: func() { data.OneTyp1 = &govalidatortest.ValidIntTagsOneOf1_TSfixed32Gte1{} },
			AfterFunc:  func() { data.OneTyp1 = &govalidatortest.ValidIntTagsOneOf1_TSfixed32Gte1{TSfixed32Gte1: 6} },
		},
		{
			Name:      fmt.Sprintf("test tag <%s>", protovalidator.TagIntIn),
			FieldDesc: "field 't_sfixed32_in1'",
			TagInfo:   &protovalidator.TagInfo{Tag: protovalidator.TagIntIn, Value: []int32{1, 2, 3}},
			FieldValue: func() interface{} {
				return data.OneTyp1.(*govalidatortest.ValidIntTagsOneOf1_TSfixed32In1).TSfixed32In1
			},
			BeforeFunc: func() { data.OneTyp1 = &govalidatortest.ValidIntTagsOneOf1_TSfixed32In1{TSfixed32In1: 4} },
			AfterFunc:  func() { data.OneTyp1 = &govalidatortest.ValidIntTagsOneOf1_TSfixed32In1{TSfixed32In1: 1} },
		},
		{
			Name:      fmt.Sprintf("test tag <%s>", protovalidator.TagIntNotIn),
			FieldDesc: "field 't_sfixed32_not_in1'",
			TagInfo:   &protovalidator.TagInfo{Tag: protovalidator.TagIntNotIn, Value: []int32{1, 2, 3}},
			FieldValue: func() interface{} {
				return data.OneTyp1.(*govalidatortest.ValidIntTagsOneOf1_TSfixed32NotIn1).TSfixed32NotIn1
			},
			BeforeFunc: func() { data.OneTyp1 = &govalidatortest.ValidIntTagsOneOf1_TSfixed32NotIn1{TSfixed32NotIn1: 1} },
			AfterFunc:  func() { data.OneTyp1 = &govalidatortest.ValidIntTagsOneOf1_TSfixed32NotIn1{TSfixed32NotIn1: 5} },
		},
	}
	casesSfixed64 := []*CaseDesc{
		{
			Name:      fmt.Sprintf("test tag <%s>", protovalidator.TagIntEq),
			FieldDesc: "field 't_sfixed64_eq1'",
			TagInfo:   &protovalidator.TagInfo{Tag: protovalidator.TagIntEq, Value: 1},
			FieldValue: func() interface{} {
				return data.OneTyp1.(*govalidatortest.ValidIntTagsOneOf1_TSfixed64Eq1).TSfixed64Eq1
			},
			BeforeFunc: func() { data.OneTyp1 = &govalidatortest.ValidIntTagsOneOf1_TSfixed64Eq1{} },
			AfterFunc:  func() { data.OneTyp1 = &govalidatortest.ValidIntTagsOneOf1_TSfixed64Eq1{TSfixed64Eq1: 1} },
		},
		{
			Name:      fmt.Sprintf("test tag <%s>", protovalidator.TagIntNe),
			FieldDesc: "field 't_sfixed64_ne1'",
			TagInfo:   &protovalidator.TagInfo{Tag: protovalidator.TagIntNe, Value: 2},
			FieldValue: func() interface{} {
				return data.OneTyp1.(*govalidatortest.ValidIntTagsOneOf1_TSfixed64Ne1).TSfixed64Ne1
			},
			BeforeFunc: func() { data.OneTyp1 = &govalidatortest.ValidIntTagsOneOf1_TSfixed64Ne1{TSfixed64Ne1: 2} },
			AfterFunc:  func() { data.OneTyp1 = &govalidatortest.ValidIntTagsOneOf1_TSfixed64Ne1{TSfixed64Ne1: 1} },
		},
		{
			Name:      fmt.Sprintf("test tag <%s>", protovalidator.TagIntLt),
			FieldDesc: "field 't_sfixed64_lt1'",
			TagInfo:   &protovalidator.TagInfo{Tag: protovalidator.TagIntLt, Value: 3},
			FieldValue: func() interface{} {
				return data.OneTyp1.(*govalidatortest.ValidIntTagsOneOf1_TSfixed64Lt1).TSfixed64Lt1
			},
			BeforeFunc: func() { data.OneTyp1 = &govalidatortest.ValidIntTagsOneOf1_TSfixed64Lt1{TSfixed64Lt1: 3} },
			AfterFunc:  func() { data.OneTyp1 = &govalidatortest.ValidIntTagsOneOf1_TSfixed64Lt1{TSfixed64Lt1: 1} },
		},
		{
			Name:      fmt.Sprintf("test tag <%s>", protovalidator.TagIntGt),
			FieldDesc: "field 't_sfixed64_gt1'",
			TagInfo:   &protovalidator.TagInfo{Tag: protovalidator.TagIntGt, Value: 4},
			FieldValue: func() interface{} {
				return data.OneTyp1.(*govalidatortest.ValidIntTagsOneOf1_TSfixed64Gt1).TSfixed64Gt1
			},
			BeforeFunc: func() { data.OneTyp1 = &govalidatortest.ValidIntTagsOneOf1_TSfixed64Gt1{} },
			AfterFunc:  func() { data.OneTyp1 = &govalidatortest.ValidIntTagsOneOf1_TSfixed64Gt1{TSfixed64Gt1: 5} },
		},
		{
			Name:      fmt.Sprintf("test tag <%s>", protovalidator.TagIntLte),
			FieldDesc: "field 't_sfixed64_lte1'",
			TagInfo:   &protovalidator.TagInfo{Tag: protovalidator.TagIntLte, Value: 5},
			FieldValue: func() interface{} {
				return data.OneTyp1.(*govalidatortest.ValidIntTagsOneOf1_TSfixed64Lte1).TSfixed64Lte1
			},
			BeforeFunc: func() { data.OneTyp1 = &govalidatortest.ValidIntTagsOneOf1_TSfixed64Lte1{TSfixed64Lte1: 6} },
			AfterFunc:  func() { data.OneTyp1 = &govalidatortest.ValidIntTagsOneOf1_TSfixed64Lte1{TSfixed64Lte1: 5} },
		},
		{
			Name:      fmt.Sprintf("test tag <%s>", protovalidator.TagIntGte),
			FieldDesc: "field 't_sfixed64_gte1'",
			TagInfo:   &protovalidator.TagInfo{Tag: protovalidator.TagIntGte, Value: 6},
			FieldValue: func() interface{} {
				return data.OneTyp1.(*govalidatortest.ValidIntTagsOneOf1_TSfixed64Gte1).TSfixed64Gte1
			},
			BeforeFunc: func() { data.OneTyp1 = &govalidatortest.ValidIntTagsOneOf1_TSfixed64Gte1{} },
			AfterFunc:  func() { data.OneTyp1 = &govalidatortest.ValidIntTagsOneOf1_TSfixed64Gte1{TSfixed64Gte1: 6} },
		},
		{
			Name:      fmt.Sprintf("test tag <%s>", protovalidator.TagIntIn),
			FieldDesc: "field 't_sfixed64_in1'",
			TagInfo:   &protovalidator.TagInfo{Tag: protovalidator.TagIntIn, Value: []int32{1, 2, 3}},
			FieldValue: func() interface{} {
				return data.OneTyp1.(*govalidatortest.ValidIntTagsOneOf1_TSfixed64In1).TSfixed64In1
			},
			BeforeFunc: func() { data.OneTyp1 = &govalidatortest.ValidIntTagsOneOf1_TSfixed64In1{TSfixed64In1: 4} },
			AfterFunc:  func() { data.OneTyp1 = &govalidatortest.ValidIntTagsOneOf1_TSfixed64In1{TSfixed64In1: 1} },
		},
		{
			Name:      fmt.Sprintf("test tag <%s>", protovalidator.TagIntNotIn),
			FieldDesc: "field 't_sfixed64_not_in1'",
			TagInfo:   &protovalidator.TagInfo{Tag: protovalidator.TagIntNotIn, Value: []int32{1, 2, 3}},
			FieldValue: func() interface{} {
				return data.OneTyp1.(*govalidatortest.ValidIntTagsOneOf1_TSfixed64NotIn1).TSfixed64NotIn1
			},
			BeforeFunc: func() { data.OneTyp1 = &govalidatortest.ValidIntTagsOneOf1_TSfixed64NotIn1{TSfixed64NotIn1: 1} },
			AfterFunc:  func() { data.OneTyp1 = &govalidatortest.ValidIntTagsOneOf1_TSfixed64NotIn1{TSfixed64NotIn1: 5} },
		},
	}

	var cases []*CaseDesc
	cases = append(cases, casesInt32...)
	cases = append(cases, casesInt64...)
	cases = append(cases, casesSint32...)
	cases = append(cases, casesSint64...)
	cases = append(cases, casesSfixed32...)
	cases = append(cases, casesSfixed64...)

	msgName := "ValidIntTagsOneOf1"
	runCases(t, data, msgName, cases)

	{
		err := data.Validate()
		require.Nil(t, err)
	}
}

func Test_GoValidator_ValidUintTagsGeneral1(t *testing.T) {
	data := &govalidatortest.ValidUintTagsGeneral1{}
	{
		err := data.Validate()
		require.NotNil(t, err)
	}
	casesUint32 := []*CaseDesc{
		{
			Name:       fmt.Sprintf("test tag <%s>", protovalidator.TagUintEq),
			FieldDesc:  "field 't_uint32_eq1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagUintEq, Value: 1},
			FieldValue: func() interface{} { return data.TUint32Eq1 },
			BeforeFunc: func() {},
			AfterFunc:  func() { data.TUint32Eq1 = 1 },
		},
		{
			Name:       fmt.Sprintf("test tag <%s>", protovalidator.TagUintNe),
			FieldDesc:  "field 't_uint32_ne1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagUintNe, Value: 2},
			FieldValue: func() interface{} { return data.TUint32Ne1 },
			BeforeFunc: func() { data.TUint32Ne1 = 2 },
			AfterFunc:  func() { data.TUint32Ne1 = 1 },
		},
		{
			Name:       fmt.Sprintf("test tag <%s>", protovalidator.TagUintLt),
			FieldDesc:  "field 't_uint32_lt1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagUintLt, Value: 3},
			FieldValue: func() interface{} { return data.TUint32Lt1 },
			BeforeFunc: func() { data.TUint32Lt1 = 3 },
			AfterFunc:  func() { data.TUint32Lt1 = 1 },
		},
		{
			Name:       fmt.Sprintf("test tag <%s>", protovalidator.TagUintGt),
			FieldDesc:  "field 't_uint32_gt1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagUintGt, Value: 4},
			FieldValue: func() interface{} { return data.TUint32Gt1 },
			BeforeFunc: func() {},
			AfterFunc:  func() { data.TUint32Gt1 = 5 },
		},
		{
			Name:       fmt.Sprintf("test tag <%s>", protovalidator.TagUintLte),
			FieldDesc:  "field 't_uint32_lte1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagUintLte, Value: 5},
			FieldValue: func() interface{} { return data.TUint32Lte1 },
			BeforeFunc: func() { data.TUint32Lte1 = 6 },
			AfterFunc:  func() { data.TUint32Lte1 = 5 },
		},
		{
			Name:       fmt.Sprintf("test tag <%s>", protovalidator.TagUintGte),
			FieldDesc:  "field 't_uint32_gte1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagUintGte, Value: 6},
			FieldValue: func() interface{} { return data.TUint32Gte1 },
			BeforeFunc: func() {},
			AfterFunc:  func() { data.TUint32Gte1 = 6 },
		},
		{
			Name:       fmt.Sprintf("test tag <%s>", protovalidator.TagUintIn),
			FieldDesc:  "field 't_uint32_in1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagUintIn, Value: []int32{1, 2, 3}},
			FieldValue: func() interface{} { return data.TUint32In1 },
			BeforeFunc: func() { data.TUint32In1 = 4 },
			AfterFunc:  func() { data.TUint32In1 = 1 },
		},
		{
			Name:       fmt.Sprintf("test tag <%s>", protovalidator.TagUintNotIn),
			FieldDesc:  "field 't_uint32_not_in1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagUintNotIn, Value: []int32{1, 2, 3}},
			FieldValue: func() interface{} { return data.TUint32NotIn1 },
			BeforeFunc: func() { data.TUint32NotIn1 = 1 },
			AfterFunc:  func() { data.TUint32NotIn1 = 5 },
		},
	}
	casesUint64 := []*CaseDesc{
		{
			Name:       fmt.Sprintf("test tag <%s>", protovalidator.TagUintEq),
			FieldDesc:  "field 't_uint64_eq1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagUintEq, Value: 1},
			FieldValue: func() interface{} { return data.TUint64Eq1 },
			BeforeFunc: func() {},
			AfterFunc:  func() { data.TUint64Eq1 = 1 },
		},
		{
			Name:       fmt.Sprintf("test tag <%s>", protovalidator.TagUintNe),
			FieldDesc:  "field 't_uint64_ne1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagUintNe, Value: 2},
			FieldValue: func() interface{} { return data.TUint64Ne1 },
			BeforeFunc: func() { data.TUint64Ne1 = 2 },
			AfterFunc:  func() { data.TUint64Ne1 = 1 },
		},
		{
			Name:       fmt.Sprintf("test tag <%s>", protovalidator.TagUintLt),
			FieldDesc:  "field 't_uint64_lt1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagUintLt, Value: 3},
			FieldValue: func() interface{} { return data.TUint64Lt1 },
			BeforeFunc: func() { data.TUint64Lt1 = 3 },
			AfterFunc:  func() { data.TUint64Lt1 = 1 },
		},
		{
			Name:       fmt.Sprintf("test tag <%s>", protovalidator.TagUintGt),
			FieldDesc:  "field 't_uint64_gt1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagUintGt, Value: 4},
			FieldValue: func() interface{} { return data.TUint64Gt1 },
			BeforeFunc: func() {},
			AfterFunc:  func() { data.TUint64Gt1 = 5 },
		},
		{
			Name:       fmt.Sprintf("test tag <%s>", protovalidator.TagUintLte),
			FieldDesc:  "field 't_uint64_lte1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagUintLte, Value: 5},
			FieldValue: func() interface{} { return data.TUint64Lte1 },
			BeforeFunc: func() { data.TUint64Lte1 = 6 },
			AfterFunc:  func() { data.TUint64Lte1 = 5 },
		},
		{
			Name:       fmt.Sprintf("test tag <%s>", protovalidator.TagUintGte),
			FieldDesc:  "field 't_uint64_gte1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagUintGte, Value: 6},
			FieldValue: func() interface{} { return data.TUint64Gte1 },
			BeforeFunc: func() {},
			AfterFunc:  func() { data.TUint64Gte1 = 6 },
		},
		{
			Name:       fmt.Sprintf("test tag <%s>", protovalidator.TagUintIn),
			FieldDesc:  "field 't_uint64_in1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagUintIn, Value: []int32{1, 2, 3}},
			FieldValue: func() interface{} { return data.TUint64In1 },
			BeforeFunc: func() { data.TUint64In1 = 4 },
			AfterFunc:  func() { data.TUint64In1 = 1 },
		},
		{
			Name:       fmt.Sprintf("test tag <%s>", protovalidator.TagUintNotIn),
			FieldDesc:  "field 't_uint64_not_in1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagUintNotIn, Value: []int32{1, 2, 3}},
			FieldValue: func() interface{} { return data.TUint64NotIn1 },
			BeforeFunc: func() { data.TUint64NotIn1 = 1 },
			AfterFunc:  func() { data.TUint64NotIn1 = 5 },
		},
	}
	casesFixed32 := []*CaseDesc{
		{
			Name:       fmt.Sprintf("test tag <%s>", protovalidator.TagUintEq),
			FieldDesc:  "field 't_fixed32_eq1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagUintEq, Value: 1},
			FieldValue: func() interface{} { return data.TFixed32Eq1 },
			BeforeFunc: func() {},
			AfterFunc:  func() { data.TFixed32Eq1 = 1 },
		},
		{
			Name:       fmt.Sprintf("test tag <%s>", protovalidator.TagUintNe),
			FieldDesc:  "field 't_fixed32_ne1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagUintNe, Value: 2},
			FieldValue: func() interface{} { return data.TFixed32Ne1 },
			BeforeFunc: func() { data.TFixed32Ne1 = 2 },
			AfterFunc:  func() { data.TFixed32Ne1 = 1 },
		},
		{
			Name:       fmt.Sprintf("test tag <%s>", protovalidator.TagUintLt),
			FieldDesc:  "field 't_fixed32_lt1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagUintLt, Value: 3},
			FieldValue: func() interface{} { return data.TFixed32Lt1 },
			BeforeFunc: func() { data.TFixed32Lt1 = 3 },
			AfterFunc:  func() { data.TFixed32Lt1 = 1 },
		},
		{
			Name:       fmt.Sprintf("test tag <%s>", protovalidator.TagUintGt),
			FieldDesc:  "field 't_fixed32_gt1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagUintGt, Value: 4},
			FieldValue: func() interface{} { return data.TFixed32Gt1 },
			BeforeFunc: func() {},
			AfterFunc:  func() { data.TFixed32Gt1 = 5 },
		},
		{
			Name:       fmt.Sprintf("test tag <%s>", protovalidator.TagUintLte),
			FieldDesc:  "field 't_fixed32_lte1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagUintLte, Value: 5},
			FieldValue: func() interface{} { return data.TFixed32Lte1 },
			BeforeFunc: func() { data.TFixed32Lte1 = 6 },
			AfterFunc:  func() { data.TFixed32Lte1 = 5 },
		},
		{
			Name:       fmt.Sprintf("test tag <%s>", protovalidator.TagUintGte),
			FieldDesc:  "field 't_fixed32_gte1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagUintGte, Value: 6},
			FieldValue: func() interface{} { return data.TFixed32Gte1 },
			BeforeFunc: func() {},
			AfterFunc:  func() { data.TFixed32Gte1 = 6 },
		},
		{
			Name:       fmt.Sprintf("test tag <%s>", protovalidator.TagUintIn),
			FieldDesc:  "field 't_fixed32_in1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagUintIn, Value: []int32{1, 2, 3}},
			FieldValue: func() interface{} { return data.TFixed32In1 },
			BeforeFunc: func() { data.TFixed32In1 = 4 },
			AfterFunc:  func() { data.TFixed32In1 = 1 },
		},
		{
			Name:       fmt.Sprintf("test tag <%s>", protovalidator.TagUintNotIn),
			FieldDesc:  "field 't_fixed32_not_in1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagUintNotIn, Value: []int32{1, 2, 3}},
			FieldValue: func() interface{} { return data.TFixed32NotIn1 },
			BeforeFunc: func() { data.TFixed32NotIn1 = 1 },
			AfterFunc:  func() { data.TFixed32NotIn1 = 5 },
		},
	}
	casesFixed64 := []*CaseDesc{
		{
			Name:       fmt.Sprintf("test tag <%s>", protovalidator.TagUintEq),
			FieldDesc:  "field 't_fixed64_eq1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagUintEq, Value: 1},
			FieldValue: func() interface{} { return data.TFixed64Eq1 },
			BeforeFunc: func() {},
			AfterFunc:  func() { data.TFixed64Eq1 = 1 },
		},
		{
			Name:       fmt.Sprintf("test tag <%s>", protovalidator.TagUintNe),
			FieldDesc:  "field 't_fixed64_ne1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagUintNe, Value: 2},
			FieldValue: func() interface{} { return data.TFixed64Ne1 },
			BeforeFunc: func() { data.TFixed64Ne1 = 2 },
			AfterFunc:  func() { data.TFixed64Ne1 = 1 },
		},
		{
			Name:       fmt.Sprintf("test tag <%s>", protovalidator.TagUintLt),
			FieldDesc:  "field 't_fixed64_lt1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagUintLt, Value: 3},
			FieldValue: func() interface{} { return data.TFixed64Lt1 },
			BeforeFunc: func() { data.TFixed64Lt1 = 3 },
			AfterFunc:  func() { data.TFixed64Lt1 = 1 },
		},
		{
			Name:       fmt.Sprintf("test tag <%s>", protovalidator.TagUintGt),
			FieldDesc:  "field 't_fixed64_gt1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagUintGt, Value: 4},
			FieldValue: func() interface{} { return data.TFixed64Gt1 },
			BeforeFunc: func() {},
			AfterFunc:  func() { data.TFixed64Gt1 = 5 },
		},
		{
			Name:       fmt.Sprintf("test tag <%s>", protovalidator.TagUintLte),
			FieldDesc:  "field 't_fixed64_lte1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagUintLte, Value: 5},
			FieldValue: func() interface{} { return data.TFixed64Lte1 },
			BeforeFunc: func() { data.TFixed64Lte1 = 6 },
			AfterFunc:  func() { data.TFixed64Lte1 = 5 },
		},
		{
			Name:       fmt.Sprintf("test tag <%s>", protovalidator.TagUintGte),
			FieldDesc:  "field 't_fixed64_gte1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagUintGte, Value: 6},
			FieldValue: func() interface{} { return data.TFixed64Gte1 },
			BeforeFunc: func() {},
			AfterFunc:  func() { data.TFixed64Gte1 = 6 },
		},
		{
			Name:       fmt.Sprintf("test tag <%s>", protovalidator.TagUintIn),
			FieldDesc:  "field 't_fixed64_in1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagUintIn, Value: []int32{1, 2, 3}},
			FieldValue: func() interface{} { return data.TFixed64In1 },
			BeforeFunc: func() { data.TFixed64In1 = 4 },
			AfterFunc:  func() { data.TFixed64In1 = 1 },
		},
		{
			Name:       fmt.Sprintf("test tag <%s>", protovalidator.TagUintNotIn),
			FieldDesc:  "field 't_fixed64_not_in1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagUintNotIn, Value: []int32{1, 2, 3}},
			FieldValue: func() interface{} { return data.TFixed64NotIn1 },
			BeforeFunc: func() { data.TFixed64NotIn1 = 1 },
			AfterFunc:  func() { data.TFixed64NotIn1 = 5 },
		},
	}

	var cases []*CaseDesc
	cases = append(cases, casesUint32...)
	cases = append(cases, casesUint64...)
	cases = append(cases, casesFixed32...)
	cases = append(cases, casesFixed64...)

	msgName := "ValidUintTagsGeneral1"
	runCases(t, data, msgName, cases)

	{
		err := data.Validate()
		require.Nil(t, err)
	}
}

func Test_GoValidator_ValidUintTagsOptional1(t *testing.T) {
	data := &govalidatortest.ValidUintTagsOptional1{}
	{
		err := data.Validate()
		require.NotNil(t, err)
	}
	casesUint32 := []*CaseDesc{
		{
			Name:       fmt.Sprintf("test tag <%s>", protovalidator.TagUintEq),
			FieldDesc:  "field 't_uint32_eq1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagUintEq, Value: 1},
			FieldValue: func() interface{} { return nil },
			BeforeFunc: func() {},
			AfterFunc:  func() { x := uint32(1); data.TUint32Eq1 = &x },
		},
		{
			Name:       fmt.Sprintf("test tag <%s>", protovalidator.TagUintNe),
			FieldDesc:  "field 't_uint32_ne1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagUintNe, Value: 2},
			FieldValue: func() interface{} { return data.TUint32Ne1 },
			BeforeFunc: func() { x := uint32(2); data.TUint32Ne1 = &x },
			AfterFunc:  func() { x := uint32(1); data.TUint32Ne1 = &x },
		},
		{
			Name:       fmt.Sprintf("test tag <%s>", protovalidator.TagUintLt),
			FieldDesc:  "field 't_uint32_lt1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagUintLt, Value: 3},
			FieldValue: func() interface{} { return data.TUint32Lt1 },
			BeforeFunc: func() { x := uint32(3); data.TUint32Lt1 = &x },
			AfterFunc:  func() { x := uint32(1); data.TUint32Lt1 = &x },
		},
		{
			Name:       fmt.Sprintf("test tag <%s>", protovalidator.TagUintGt),
			FieldDesc:  "field 't_uint32_gt1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagUintGt, Value: 4},
			FieldValue: func() interface{} { return nil },
			BeforeFunc: func() {},
			AfterFunc:  func() { x := uint32(5); data.TUint32Gt1 = &x },
		},
		{
			Name:       fmt.Sprintf("test tag <%s>", protovalidator.TagUintLte),
			FieldDesc:  "field 't_uint32_lte1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagUintLte, Value: 5},
			FieldValue: func() interface{} { return data.TUint32Lte1 },
			BeforeFunc: func() { x := uint32(6); data.TUint32Lte1 = &x },
			AfterFunc:  func() { x := uint32(5); data.TUint32Lte1 = &x },
		},
		{
			Name:       fmt.Sprintf("test tag <%s>", protovalidator.TagUintGte),
			FieldDesc:  "field 't_uint32_gte1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagUintGte, Value: 6},
			FieldValue: func() interface{} { return nil },
			BeforeFunc: func() {},
			AfterFunc:  func() { x := uint32(6); data.TUint32Gte1 = &x },
		},
		{
			Name:       fmt.Sprintf("test tag <%s>", protovalidator.TagUintIn),
			FieldDesc:  "field 't_uint32_in1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagUintIn, Value: []int32{1, 2, 3}},
			FieldValue: func() interface{} { return data.TUint32In1 },
			BeforeFunc: func() { x := uint32(4); data.TUint32In1 = &x },
			AfterFunc:  func() { x := uint32(1); data.TUint32In1 = &x },
		},
		{
			Name:       fmt.Sprintf("test tag <%s>", protovalidator.TagUintNotIn),
			FieldDesc:  "field 't_uint32_not_in1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagUintNotIn, Value: []int32{1, 2, 3}},
			FieldValue: func() interface{} { return data.TUint32NotIn1 },
			BeforeFunc: func() { x := uint32(1); data.TUint32NotIn1 = &x },
			AfterFunc:  func() { x := uint32(5); data.TUint32NotIn1 = &x },
		},
	}
	casesUint64 := []*CaseDesc{
		{
			Name:       fmt.Sprintf("test tag <%s>", protovalidator.TagUintEq),
			FieldDesc:  "field 't_uint64_eq1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagUintEq, Value: 1},
			FieldValue: func() interface{} { return nil },
			BeforeFunc: func() {},
			AfterFunc:  func() { x := uint64(1); data.TUint64Eq1 = &x },
		},
		{
			Name:       fmt.Sprintf("test tag <%s>", protovalidator.TagUintNe),
			FieldDesc:  "field 't_uint64_ne1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagUintNe, Value: 2},
			FieldValue: func() interface{} { return data.TUint64Ne1 },
			BeforeFunc: func() { x := uint64(2); data.TUint64Ne1 = &x },
			AfterFunc:  func() { x := uint64(1); data.TUint64Ne1 = &x },
		},
		{
			Name:       fmt.Sprintf("test tag <%s>", protovalidator.TagUintLt),
			FieldDesc:  "field 't_uint64_lt1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagUintLt, Value: 3},
			FieldValue: func() interface{} { return data.TUint64Lt1 },
			BeforeFunc: func() { x := uint64(3); data.TUint64Lt1 = &x },
			AfterFunc:  func() { x := uint64(1); data.TUint64Lt1 = &x },
		},
		{
			Name:       fmt.Sprintf("test tag <%s>", protovalidator.TagUintGt),
			FieldDesc:  "field 't_uint64_gt1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagUintGt, Value: 4},
			FieldValue: func() interface{} { return nil },
			BeforeFunc: func() {},
			AfterFunc:  func() { x := uint64(5); data.TUint64Gt1 = &x },
		},
		{
			Name:       fmt.Sprintf("test tag <%s>", protovalidator.TagUintLte),
			FieldDesc:  "field 't_uint64_lte1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagUintLte, Value: 5},
			FieldValue: func() interface{} { return data.TUint64Lte1 },
			BeforeFunc: func() { x := uint64(6); data.TUint64Lte1 = &x },
			AfterFunc:  func() { x := uint64(5); data.TUint64Lte1 = &x },
		},
		{
			Name:       fmt.Sprintf("test tag <%s>", protovalidator.TagUintGte),
			FieldDesc:  "field 't_uint64_gte1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagUintGte, Value: 6},
			FieldValue: func() interface{} { return nil },
			BeforeFunc: func() {},
			AfterFunc:  func() { x := uint64(6); data.TUint64Gte1 = &x },
		},
		{
			Name:       fmt.Sprintf("test tag <%s>", protovalidator.TagUintIn),
			FieldDesc:  "field 't_uint64_in1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagUintIn, Value: []int32{1, 2, 3}},
			FieldValue: func() interface{} { return data.TUint64In1 },
			BeforeFunc: func() { x := uint64(4); data.TUint64In1 = &x },
			AfterFunc:  func() { x := uint64(1); data.TUint64In1 = &x },
		},
		{
			Name:       fmt.Sprintf("test tag <%s>", protovalidator.TagUintNotIn),
			FieldDesc:  "field 't_uint64_not_in1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagUintNotIn, Value: []int32{1, 2, 3}},
			FieldValue: func() interface{} { return data.TUint64NotIn1 },
			BeforeFunc: func() { x := uint64(1); data.TUint64NotIn1 = &x },
			AfterFunc:  func() { x := uint64(5); data.TUint64NotIn1 = &x },
		},
	}
	casesFixed32 := []*CaseDesc{
		{
			Name:       fmt.Sprintf("test tag <%s>", protovalidator.TagUintEq),
			FieldDesc:  "field 't_fixed32_eq1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagUintEq, Value: 1},
			FieldValue: func() interface{} { return nil },
			BeforeFunc: func() {},
			AfterFunc:  func() { x := uint32(1); data.TFixed32Eq1 = &x },
		},
		{
			Name:       fmt.Sprintf("test tag <%s>", protovalidator.TagUintNe),
			FieldDesc:  "field 't_fixed32_ne1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagUintNe, Value: 2},
			FieldValue: func() interface{} { return data.TFixed32Ne1 },
			BeforeFunc: func() { x := uint32(2); data.TFixed32Ne1 = &x },
			AfterFunc:  func() { x := uint32(1); data.TFixed32Ne1 = &x },
		},
		{
			Name:       fmt.Sprintf("test tag <%s>", protovalidator.TagUintLt),
			FieldDesc:  "field 't_fixed32_lt1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagUintLt, Value: 3},
			FieldValue: func() interface{} { return data.TFixed32Lt1 },
			BeforeFunc: func() { x := uint32(3); data.TFixed32Lt1 = &x },
			AfterFunc:  func() { x := uint32(1); data.TFixed32Lt1 = &x },
		},
		{
			Name:       fmt.Sprintf("test tag <%s>", protovalidator.TagUintGt),
			FieldDesc:  "field 't_fixed32_gt1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagUintGt, Value: 4},
			FieldValue: func() interface{} { return nil },
			BeforeFunc: func() {},
			AfterFunc:  func() { x := uint32(5); data.TFixed32Gt1 = &x },
		},
		{
			Name:       fmt.Sprintf("test tag <%s>", protovalidator.TagUintLte),
			FieldDesc:  "field 't_fixed32_lte1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagUintLte, Value: 5},
			FieldValue: func() interface{} { return data.TFixed32Lte1 },
			BeforeFunc: func() { x := uint32(6); data.TFixed32Lte1 = &x },
			AfterFunc:  func() { x := uint32(5); data.TFixed32Lte1 = &x },
		},
		{
			Name:       fmt.Sprintf("test tag <%s>", protovalidator.TagUintGte),
			FieldDesc:  "field 't_fixed32_gte1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagUintGte, Value: 6},
			FieldValue: func() interface{} { return nil },
			BeforeFunc: func() {},
			AfterFunc:  func() { x := uint32(6); data.TFixed32Gte1 = &x },
		},
		{
			Name:       fmt.Sprintf("test tag <%s>", protovalidator.TagUintIn),
			FieldDesc:  "field 't_fixed32_in1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagUintIn, Value: []int32{1, 2, 3}},
			FieldValue: func() interface{} { return data.TFixed32In1 },
			BeforeFunc: func() { x := uint32(4); data.TFixed32In1 = &x },
			AfterFunc:  func() { x := uint32(1); data.TFixed32In1 = &x },
		},
		{
			Name:       fmt.Sprintf("test tag <%s>", protovalidator.TagUintNotIn),
			FieldDesc:  "field 't_fixed32_not_in1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagUintNotIn, Value: []int32{1, 2, 3}},
			FieldValue: func() interface{} { return data.TFixed32NotIn1 },
			BeforeFunc: func() { x := uint32(1); data.TFixed32NotIn1 = &x },
			AfterFunc:  func() { x := uint32(5); data.TFixed32NotIn1 = &x },
		},
	}
	casesFixed64 := []*CaseDesc{
		{
			Name:       fmt.Sprintf("test tag <%s>", protovalidator.TagUintEq),
			FieldDesc:  "field 't_fixed64_eq1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagUintEq, Value: 1},
			FieldValue: func() interface{} { return nil },
			BeforeFunc: func() {},
			AfterFunc:  func() { x := uint64(1); data.TFixed64Eq1 = &x },
		},
		{
			Name:       fmt.Sprintf("test tag <%s>", protovalidator.TagUintNe),
			FieldDesc:  "field 't_fixed64_ne1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagUintNe, Value: 2},
			FieldValue: func() interface{} { return data.TFixed64Ne1 },
			BeforeFunc: func() { x := uint64(2); data.TFixed64Ne1 = &x },
			AfterFunc:  func() { x := uint64(1); data.TFixed64Ne1 = &x },
		},
		{
			Name:       fmt.Sprintf("test tag <%s>", protovalidator.TagUintLt),
			FieldDesc:  "field 't_fixed64_lt1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagUintLt, Value: 3},
			FieldValue: func() interface{} { return data.TFixed64Lt1 },
			BeforeFunc: func() { x := uint64(3); data.TFixed64Lt1 = &x },
			AfterFunc:  func() { x := uint64(1); data.TFixed64Lt1 = &x },
		},
		{
			Name:       fmt.Sprintf("test tag <%s>", protovalidator.TagUintGt),
			FieldDesc:  "field 't_fixed64_gt1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagUintGt, Value: 4},
			FieldValue: func() interface{} { return nil },
			BeforeFunc: func() {},
			AfterFunc:  func() { x := uint64(5); data.TFixed64Gt1 = &x },
		},
		{
			Name:       fmt.Sprintf("test tag <%s>", protovalidator.TagUintLte),
			FieldDesc:  "field 't_fixed64_lte1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagUintLte, Value: 5},
			FieldValue: func() interface{} { return data.TFixed64Lte1 },
			BeforeFunc: func() { x := uint64(6); data.TFixed64Lte1 = &x },
			AfterFunc:  func() { x := uint64(5); data.TFixed64Lte1 = &x },
		},
		{
			Name:       fmt.Sprintf("test tag <%s>", protovalidator.TagUintGte),
			FieldDesc:  "field 't_fixed64_gte1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagUintGte, Value: 6},
			FieldValue: func() interface{} { return nil },
			BeforeFunc: func() {},
			AfterFunc:  func() { x := uint64(6); data.TFixed64Gte1 = &x },
		},
		{
			Name:       fmt.Sprintf("test tag <%s>", protovalidator.TagUintIn),
			FieldDesc:  "field 't_fixed64_in1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagUintIn, Value: []int32{1, 2, 3}},
			FieldValue: func() interface{} { return data.TFixed64In1 },
			BeforeFunc: func() { x := uint64(4); data.TFixed64In1 = &x },
			AfterFunc:  func() { x := uint64(1); data.TFixed64In1 = &x },
		},
		{
			Name:       fmt.Sprintf("test tag <%s>", protovalidator.TagUintNotIn),
			FieldDesc:  "field 't_fixed64_not_in1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagUintNotIn, Value: []int32{1, 2, 3}},
			FieldValue: func() interface{} { return data.TFixed64NotIn1 },
			BeforeFunc: func() { x := uint64(1); data.TFixed64NotIn1 = &x },
			AfterFunc:  func() { x := uint64(5); data.TFixed64NotIn1 = &x },
		},
	}

	var cases []*CaseDesc
	cases = append(cases, casesUint32...)
	cases = append(cases, casesUint64...)
	cases = append(cases, casesFixed32...)
	cases = append(cases, casesFixed64...)

	msgName := "ValidUintTagsOptional1"
	runCases(t, data, msgName, cases)

	{
		err := data.Validate()
		require.Nil(t, err)
	}
}

func Test_GoValidator_ValidUintTagsOneOf1(t *testing.T) {
	data := &govalidatortest.ValidUintTagsOneOf1{}
	{
		err := data.Validate()
		require.Nil(t, err)
	}
	casesUint32 := []*CaseDesc{
		{
			Name:       fmt.Sprintf("test tag <%s>", protovalidator.TagUintEq),
			FieldDesc:  "field 't_uint32_eq1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagUintEq, Value: 1},
			FieldValue: func() interface{} { return data.OneTyp1.(*govalidatortest.ValidUintTagsOneOf1_TUint32Eq1).TUint32Eq1 },
			BeforeFunc: func() { data.OneTyp1 = &govalidatortest.ValidUintTagsOneOf1_TUint32Eq1{} },
			AfterFunc:  func() { data.OneTyp1 = &govalidatortest.ValidUintTagsOneOf1_TUint32Eq1{TUint32Eq1: 1} },
		},
		{
			Name:       fmt.Sprintf("test tag <%s>", protovalidator.TagUintNe),
			FieldDesc:  "field 't_uint32_ne1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagUintNe, Value: 2},
			FieldValue: func() interface{} { return data.OneTyp1.(*govalidatortest.ValidUintTagsOneOf1_TUint32Ne1).TUint32Ne1 },
			BeforeFunc: func() { data.OneTyp1 = &govalidatortest.ValidUintTagsOneOf1_TUint32Ne1{TUint32Ne1: 2} },
			AfterFunc:  func() { data.OneTyp1 = &govalidatortest.ValidUintTagsOneOf1_TUint32Ne1{TUint32Ne1: 1} },
		},
		{
			Name:       fmt.Sprintf("test tag <%s>", protovalidator.TagUintLt),
			FieldDesc:  "field 't_uint32_lt1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagUintLt, Value: 3},
			FieldValue: func() interface{} { return data.OneTyp1.(*govalidatortest.ValidUintTagsOneOf1_TUint32Lt1).TUint32Lt1 },
			BeforeFunc: func() { data.OneTyp1 = &govalidatortest.ValidUintTagsOneOf1_TUint32Lt1{TUint32Lt1: 3} },
			AfterFunc:  func() { data.OneTyp1 = &govalidatortest.ValidUintTagsOneOf1_TUint32Lt1{TUint32Lt1: 1} },
		},
		{
			Name:       fmt.Sprintf("test tag <%s>", protovalidator.TagUintGt),
			FieldDesc:  "field 't_uint32_gt1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagUintGt, Value: 4},
			FieldValue: func() interface{} { return data.OneTyp1.(*govalidatortest.ValidUintTagsOneOf1_TUint32Gt1).TUint32Gt1 },
			BeforeFunc: func() { data.OneTyp1 = &govalidatortest.ValidUintTagsOneOf1_TUint32Gt1{} },
			AfterFunc:  func() { data.OneTyp1 = &govalidatortest.ValidUintTagsOneOf1_TUint32Gt1{TUint32Gt1: 5} },
		},
		{
			Name:      fmt.Sprintf("test tag <%s>", protovalidator.TagUintLte),
			FieldDesc: "field 't_uint32_lte1'",
			TagInfo:   &protovalidator.TagInfo{Tag: protovalidator.TagUintLte, Value: 5},
			FieldValue: func() interface{} {
				return data.OneTyp1.(*govalidatortest.ValidUintTagsOneOf1_TUint32Lte1).TUint32Lte1
			},
			BeforeFunc: func() { data.OneTyp1 = &govalidatortest.ValidUintTagsOneOf1_TUint32Lte1{TUint32Lte1: 6} },
			AfterFunc:  func() { data.OneTyp1 = &govalidatortest.ValidUintTagsOneOf1_TUint32Lte1{TUint32Lte1: 5} },
		},
		{
			Name:      fmt.Sprintf("test tag <%s>", protovalidator.TagUintGte),
			FieldDesc: "field 't_uint32_gte1'",
			TagInfo:   &protovalidator.TagInfo{Tag: protovalidator.TagUintGte, Value: 6},
			FieldValue: func() interface{} {
				return data.OneTyp1.(*govalidatortest.ValidUintTagsOneOf1_TUint32Gte1).TUint32Gte1
			},
			BeforeFunc: func() { data.OneTyp1 = &govalidatortest.ValidUintTagsOneOf1_TUint32Gte1{} },
			AfterFunc:  func() { data.OneTyp1 = &govalidatortest.ValidUintTagsOneOf1_TUint32Gte1{TUint32Gte1: 6} },
		},
		{
			Name:       fmt.Sprintf("test tag <%s>", protovalidator.TagUintIn),
			FieldDesc:  "field 't_uint32_in1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagUintIn, Value: []int32{1, 2, 3}},
			FieldValue: func() interface{} { return data.OneTyp1.(*govalidatortest.ValidUintTagsOneOf1_TUint32In1).TUint32In1 },
			BeforeFunc: func() { data.OneTyp1 = &govalidatortest.ValidUintTagsOneOf1_TUint32In1{TUint32In1: 4} },
			AfterFunc:  func() { data.OneTyp1 = &govalidatortest.ValidUintTagsOneOf1_TUint32In1{TUint32In1: 1} },
		},
		{
			Name:      fmt.Sprintf("test tag <%s>", protovalidator.TagUintNotIn),
			FieldDesc: "field 't_uint32_not_in1'",
			TagInfo:   &protovalidator.TagInfo{Tag: protovalidator.TagUintNotIn, Value: []int32{1, 2, 3}},
			FieldValue: func() interface{} {
				return data.OneTyp1.(*govalidatortest.ValidUintTagsOneOf1_TUint32NotIn1).TUint32NotIn1
			},
			BeforeFunc: func() { data.OneTyp1 = &govalidatortest.ValidUintTagsOneOf1_TUint32NotIn1{TUint32NotIn1: 1} },
			AfterFunc:  func() { data.OneTyp1 = &govalidatortest.ValidUintTagsOneOf1_TUint32NotIn1{TUint32NotIn1: 5} },
		},
	}
	casesUint64 := []*CaseDesc{
		{
			Name:       fmt.Sprintf("test tag <%s>", protovalidator.TagUintEq),
			FieldDesc:  "field 't_uint64_eq1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagUintEq, Value: 1},
			FieldValue: func() interface{} { return data.OneTyp1.(*govalidatortest.ValidUintTagsOneOf1_TUint64Eq1).TUint64Eq1 },
			BeforeFunc: func() { data.OneTyp1 = &govalidatortest.ValidUintTagsOneOf1_TUint64Eq1{} },
			AfterFunc:  func() { data.OneTyp1 = &govalidatortest.ValidUintTagsOneOf1_TUint64Eq1{TUint64Eq1: 1} },
		},
		{
			Name:       fmt.Sprintf("test tag <%s>", protovalidator.TagUintNe),
			FieldDesc:  "field 't_uint64_ne1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagUintNe, Value: 2},
			FieldValue: func() interface{} { return data.OneTyp1.(*govalidatortest.ValidUintTagsOneOf1_TUint64Ne1).TUint64Ne1 },
			BeforeFunc: func() { data.OneTyp1 = &govalidatortest.ValidUintTagsOneOf1_TUint64Ne1{TUint64Ne1: 2} },
			AfterFunc:  func() { data.OneTyp1 = &govalidatortest.ValidUintTagsOneOf1_TUint64Ne1{TUint64Ne1: 1} },
		},
		{
			Name:       fmt.Sprintf("test tag <%s>", protovalidator.TagUintLt),
			FieldDesc:  "field 't_uint64_lt1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagUintLt, Value: 3},
			FieldValue: func() interface{} { return data.OneTyp1.(*govalidatortest.ValidUintTagsOneOf1_TUint64Lt1).TUint64Lt1 },
			BeforeFunc: func() { data.OneTyp1 = &govalidatortest.ValidUintTagsOneOf1_TUint64Lt1{TUint64Lt1: 3} },
			AfterFunc:  func() { data.OneTyp1 = &govalidatortest.ValidUintTagsOneOf1_TUint64Lt1{TUint64Lt1: 1} },
		},
		{
			Name:       fmt.Sprintf("test tag <%s>", protovalidator.TagUintGt),
			FieldDesc:  "field 't_uint64_gt1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagUintGt, Value: 4},
			FieldValue: func() interface{} { return data.OneTyp1.(*govalidatortest.ValidUintTagsOneOf1_TUint64Gt1).TUint64Gt1 },
			BeforeFunc: func() { data.OneTyp1 = &govalidatortest.ValidUintTagsOneOf1_TUint64Gt1{} },
			AfterFunc:  func() { data.OneTyp1 = &govalidatortest.ValidUintTagsOneOf1_TUint64Gt1{TUint64Gt1: 5} },
		},
		{
			Name:      fmt.Sprintf("test tag <%s>", protovalidator.TagUintLte),
			FieldDesc: "field 't_uint64_lte1'",
			TagInfo:   &protovalidator.TagInfo{Tag: protovalidator.TagUintLte, Value: 5},
			FieldValue: func() interface{} {
				return data.OneTyp1.(*govalidatortest.ValidUintTagsOneOf1_TUint64Lte1).TUint64Lte1
			},
			BeforeFunc: func() { data.OneTyp1 = &govalidatortest.ValidUintTagsOneOf1_TUint64Lte1{TUint64Lte1: 6} },
			AfterFunc:  func() { data.OneTyp1 = &govalidatortest.ValidUintTagsOneOf1_TUint64Lte1{TUint64Lte1: 5} },
		},
		{
			Name:      fmt.Sprintf("test tag <%s>", protovalidator.TagUintGte),
			FieldDesc: "field 't_uint64_gte1'",
			TagInfo:   &protovalidator.TagInfo{Tag: protovalidator.TagUintGte, Value: 6},
			FieldValue: func() interface{} {
				return data.OneTyp1.(*govalidatortest.ValidUintTagsOneOf1_TUint64Gte1).TUint64Gte1
			},
			BeforeFunc: func() { data.OneTyp1 = &govalidatortest.ValidUintTagsOneOf1_TUint64Gte1{} },
			AfterFunc:  func() { data.OneTyp1 = &govalidatortest.ValidUintTagsOneOf1_TUint64Gte1{TUint64Gte1: 6} },
		},
		{
			Name:       fmt.Sprintf("test tag <%s>", protovalidator.TagUintIn),
			FieldDesc:  "field 't_uint64_in1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagUintIn, Value: []int32{1, 2, 3}},
			FieldValue: func() interface{} { return data.OneTyp1.(*govalidatortest.ValidUintTagsOneOf1_TUint64In1).TUint64In1 },
			BeforeFunc: func() { data.OneTyp1 = &govalidatortest.ValidUintTagsOneOf1_TUint64In1{TUint64In1: 4} },
			AfterFunc:  func() { data.OneTyp1 = &govalidatortest.ValidUintTagsOneOf1_TUint64In1{TUint64In1: 1} },
		},
		{
			Name:      fmt.Sprintf("test tag <%s>", protovalidator.TagUintNotIn),
			FieldDesc: "field 't_uint64_not_in1'",
			TagInfo:   &protovalidator.TagInfo{Tag: protovalidator.TagUintNotIn, Value: []int32{1, 2, 3}},
			FieldValue: func() interface{} {
				return data.OneTyp1.(*govalidatortest.ValidUintTagsOneOf1_TUint64NotIn1).TUint64NotIn1
			},
			BeforeFunc: func() { data.OneTyp1 = &govalidatortest.ValidUintTagsOneOf1_TUint64NotIn1{TUint64NotIn1: 1} },
			AfterFunc:  func() { data.OneTyp1 = &govalidatortest.ValidUintTagsOneOf1_TUint64NotIn1{TUint64NotIn1: 5} },
		},
	}
	casesFixed32 := []*CaseDesc{
		{
			Name:      fmt.Sprintf("test tag <%s>", protovalidator.TagUintEq),
			FieldDesc: "field 't_fixed32_eq1'",
			TagInfo:   &protovalidator.TagInfo{Tag: protovalidator.TagUintEq, Value: 1},
			FieldValue: func() interface{} {
				return data.OneTyp1.(*govalidatortest.ValidUintTagsOneOf1_TFixed32Eq1).TFixed32Eq1
			},
			BeforeFunc: func() { data.OneTyp1 = &govalidatortest.ValidUintTagsOneOf1_TFixed32Eq1{} },
			AfterFunc:  func() { data.OneTyp1 = &govalidatortest.ValidUintTagsOneOf1_TFixed32Eq1{TFixed32Eq1: 1} },
		},
		{
			Name:      fmt.Sprintf("test tag <%s>", protovalidator.TagUintNe),
			FieldDesc: "field 't_fixed32_ne1'",
			TagInfo:   &protovalidator.TagInfo{Tag: protovalidator.TagUintNe, Value: 2},
			FieldValue: func() interface{} {
				return data.OneTyp1.(*govalidatortest.ValidUintTagsOneOf1_TFixed32Ne1).TFixed32Ne1
			},
			BeforeFunc: func() { data.OneTyp1 = &govalidatortest.ValidUintTagsOneOf1_TFixed32Ne1{TFixed32Ne1: 2} },
			AfterFunc:  func() { data.OneTyp1 = &govalidatortest.ValidUintTagsOneOf1_TFixed32Ne1{TFixed32Ne1: 1} },
		},
		{
			Name:      fmt.Sprintf("test tag <%s>", protovalidator.TagUintLt),
			FieldDesc: "field 't_fixed32_lt1'",
			TagInfo:   &protovalidator.TagInfo{Tag: protovalidator.TagUintLt, Value: 3},
			FieldValue: func() interface{} {
				return data.OneTyp1.(*govalidatortest.ValidUintTagsOneOf1_TFixed32Lt1).TFixed32Lt1
			},
			BeforeFunc: func() { data.OneTyp1 = &govalidatortest.ValidUintTagsOneOf1_TFixed32Lt1{TFixed32Lt1: 3} },
			AfterFunc:  func() { data.OneTyp1 = &govalidatortest.ValidUintTagsOneOf1_TFixed32Lt1{TFixed32Lt1: 1} },
		},
		{
			Name:      fmt.Sprintf("test tag <%s>", protovalidator.TagUintGt),
			FieldDesc: "field 't_fixed32_gt1'",
			TagInfo:   &protovalidator.TagInfo{Tag: protovalidator.TagUintGt, Value: 4},
			FieldValue: func() interface{} {
				return data.OneTyp1.(*govalidatortest.ValidUintTagsOneOf1_TFixed32Gt1).TFixed32Gt1
			},
			BeforeFunc: func() { data.OneTyp1 = &govalidatortest.ValidUintTagsOneOf1_TFixed32Gt1{} },
			AfterFunc:  func() { data.OneTyp1 = &govalidatortest.ValidUintTagsOneOf1_TFixed32Gt1{TFixed32Gt1: 5} },
		},
		{
			Name:      fmt.Sprintf("test tag <%s>", protovalidator.TagUintLte),
			FieldDesc: "field 't_fixed32_lte1'",
			TagInfo:   &protovalidator.TagInfo{Tag: protovalidator.TagUintLte, Value: 5},
			FieldValue: func() interface{} {
				return data.OneTyp1.(*govalidatortest.ValidUintTagsOneOf1_TFixed32Lte1).TFixed32Lte1
			},
			BeforeFunc: func() { data.OneTyp1 = &govalidatortest.ValidUintTagsOneOf1_TFixed32Lte1{TFixed32Lte1: 6} },
			AfterFunc:  func() { data.OneTyp1 = &govalidatortest.ValidUintTagsOneOf1_TFixed32Lte1{TFixed32Lte1: 5} },
		},
		{
			Name:      fmt.Sprintf("test tag <%s>", protovalidator.TagUintGte),
			FieldDesc: "field 't_fixed32_gte1'",
			TagInfo:   &protovalidator.TagInfo{Tag: protovalidator.TagUintGte, Value: 6},
			FieldValue: func() interface{} {
				return data.OneTyp1.(*govalidatortest.ValidUintTagsOneOf1_TFixed32Gte1).TFixed32Gte1
			},
			BeforeFunc: func() { data.OneTyp1 = &govalidatortest.ValidUintTagsOneOf1_TFixed32Gte1{} },
			AfterFunc:  func() { data.OneTyp1 = &govalidatortest.ValidUintTagsOneOf1_TFixed32Gte1{TFixed32Gte1: 6} },
		},
		{
			Name:      fmt.Sprintf("test tag <%s>", protovalidator.TagUintIn),
			FieldDesc: "field 't_fixed32_in1'",
			TagInfo:   &protovalidator.TagInfo{Tag: protovalidator.TagUintIn, Value: []int32{1, 2, 3}},
			FieldValue: func() interface{} {
				return data.OneTyp1.(*govalidatortest.ValidUintTagsOneOf1_TFixed32In1).TFixed32In1
			},
			BeforeFunc: func() { data.OneTyp1 = &govalidatortest.ValidUintTagsOneOf1_TFixed32In1{TFixed32In1: 4} },
			AfterFunc:  func() { data.OneTyp1 = &govalidatortest.ValidUintTagsOneOf1_TFixed32In1{TFixed32In1: 1} },
		},
		{
			Name:      fmt.Sprintf("test tag <%s>", protovalidator.TagUintNotIn),
			FieldDesc: "field 't_fixed32_not_in1'",
			TagInfo:   &protovalidator.TagInfo{Tag: protovalidator.TagUintNotIn, Value: []int32{1, 2, 3}},
			FieldValue: func() interface{} {
				return data.OneTyp1.(*govalidatortest.ValidUintTagsOneOf1_TFixed32NotIn1).TFixed32NotIn1
			},
			BeforeFunc: func() { data.OneTyp1 = &govalidatortest.ValidUintTagsOneOf1_TFixed32NotIn1{TFixed32NotIn1: 1} },
			AfterFunc:  func() { data.OneTyp1 = &govalidatortest.ValidUintTagsOneOf1_TFixed32NotIn1{TFixed32NotIn1: 5} },
		},
	}
	casesFixed64 := []*CaseDesc{
		{
			Name:      fmt.Sprintf("test tag <%s>", protovalidator.TagIntEq),
			FieldDesc: "field 't_fixed64_eq1'",
			TagInfo:   &protovalidator.TagInfo{Tag: protovalidator.TagIntEq, Value: 1},
			FieldValue: func() interface{} {
				return data.OneTyp1.(*govalidatortest.ValidUintTagsOneOf1_TFixed64Eq1).TFixed64Eq1
			},
			BeforeFunc: func() { data.OneTyp1 = &govalidatortest.ValidUintTagsOneOf1_TFixed64Eq1{} },
			AfterFunc:  func() { data.OneTyp1 = &govalidatortest.ValidUintTagsOneOf1_TFixed64Eq1{TFixed64Eq1: 1} },
		},
		{
			Name:      fmt.Sprintf("test tag <%s>", protovalidator.TagUintNe),
			FieldDesc: "field 't_fixed64_ne1'",
			TagInfo:   &protovalidator.TagInfo{Tag: protovalidator.TagUintNe, Value: 2},
			FieldValue: func() interface{} {
				return data.OneTyp1.(*govalidatortest.ValidUintTagsOneOf1_TFixed64Ne1).TFixed64Ne1
			},
			BeforeFunc: func() { data.OneTyp1 = &govalidatortest.ValidUintTagsOneOf1_TFixed64Ne1{TFixed64Ne1: 2} },
			AfterFunc:  func() { data.OneTyp1 = &govalidatortest.ValidUintTagsOneOf1_TFixed64Ne1{TFixed64Ne1: 1} },
		},
		{
			Name:      fmt.Sprintf("test tag <%s>", protovalidator.TagUintLt),
			FieldDesc: "field 't_fixed64_lt1'",
			TagInfo:   &protovalidator.TagInfo{Tag: protovalidator.TagUintLt, Value: 3},
			FieldValue: func() interface{} {
				return data.OneTyp1.(*govalidatortest.ValidUintTagsOneOf1_TFixed64Lt1).TFixed64Lt1
			},
			BeforeFunc: func() { data.OneTyp1 = &govalidatortest.ValidUintTagsOneOf1_TFixed64Lt1{TFixed64Lt1: 3} },
			AfterFunc:  func() { data.OneTyp1 = &govalidatortest.ValidUintTagsOneOf1_TFixed64Lt1{TFixed64Lt1: 1} },
		},
		{
			Name:      fmt.Sprintf("test tag <%s>", protovalidator.TagUintGt),
			FieldDesc: "field 't_fixed64_gt1'",
			TagInfo:   &protovalidator.TagInfo{Tag: protovalidator.TagUintGt, Value: 4},
			FieldValue: func() interface{} {
				return data.OneTyp1.(*govalidatortest.ValidUintTagsOneOf1_TFixed64Gt1).TFixed64Gt1
			},
			BeforeFunc: func() { data.OneTyp1 = &govalidatortest.ValidUintTagsOneOf1_TFixed64Gt1{} },
			AfterFunc:  func() { data.OneTyp1 = &govalidatortest.ValidUintTagsOneOf1_TFixed64Gt1{TFixed64Gt1: 5} },
		},
		{
			Name:      fmt.Sprintf("test tag <%s>", protovalidator.TagUintLte),
			FieldDesc: "field 't_fixed64_lte1'",
			TagInfo:   &protovalidator.TagInfo{Tag: protovalidator.TagUintLte, Value: 5},
			FieldValue: func() interface{} {
				return data.OneTyp1.(*govalidatortest.ValidUintTagsOneOf1_TFixed64Lte1).TFixed64Lte1
			},
			BeforeFunc: func() { data.OneTyp1 = &govalidatortest.ValidUintTagsOneOf1_TFixed64Lte1{TFixed64Lte1: 6} },
			AfterFunc:  func() { data.OneTyp1 = &govalidatortest.ValidUintTagsOneOf1_TFixed64Lte1{TFixed64Lte1: 5} },
		},
		{
			Name:      fmt.Sprintf("test tag <%s>", protovalidator.TagUintGte),
			FieldDesc: "field 't_fixed64_gte1'",
			TagInfo:   &protovalidator.TagInfo{Tag: protovalidator.TagUintGte, Value: 6},
			FieldValue: func() interface{} {
				return data.OneTyp1.(*govalidatortest.ValidUintTagsOneOf1_TFixed64Gte1).TFixed64Gte1
			},
			BeforeFunc: func() { data.OneTyp1 = &govalidatortest.ValidUintTagsOneOf1_TFixed64Gte1{} },
			AfterFunc:  func() { data.OneTyp1 = &govalidatortest.ValidUintTagsOneOf1_TFixed64Gte1{TFixed64Gte1: 6} },
		},
		{
			Name:      fmt.Sprintf("test tag <%s>", protovalidator.TagUintIn),
			FieldDesc: "field 't_fixed64_in1'",
			TagInfo:   &protovalidator.TagInfo{Tag: protovalidator.TagUintIn, Value: []int32{1, 2, 3}},
			FieldValue: func() interface{} {
				return data.OneTyp1.(*govalidatortest.ValidUintTagsOneOf1_TFixed64In1).TFixed64In1
			},
			BeforeFunc: func() { data.OneTyp1 = &govalidatortest.ValidUintTagsOneOf1_TFixed64In1{TFixed64In1: 4} },
			AfterFunc:  func() { data.OneTyp1 = &govalidatortest.ValidUintTagsOneOf1_TFixed64In1{TFixed64In1: 1} },
		},
		{
			Name:      fmt.Sprintf("test tag <%s>", protovalidator.TagUintNotIn),
			FieldDesc: "field 't_fixed64_not_in1'",
			TagInfo:   &protovalidator.TagInfo{Tag: protovalidator.TagUintNotIn, Value: []int32{1, 2, 3}},
			FieldValue: func() interface{} {
				return data.OneTyp1.(*govalidatortest.ValidUintTagsOneOf1_TFixed64NotIn1).TFixed64NotIn1
			},
			BeforeFunc: func() { data.OneTyp1 = &govalidatortest.ValidUintTagsOneOf1_TFixed64NotIn1{TFixed64NotIn1: 1} },
			AfterFunc:  func() { data.OneTyp1 = &govalidatortest.ValidUintTagsOneOf1_TFixed64NotIn1{TFixed64NotIn1: 5} },
		},
	}

	var cases []*CaseDesc
	cases = append(cases, casesUint32...)
	cases = append(cases, casesUint64...)
	cases = append(cases, casesFixed32...)
	cases = append(cases, casesFixed64...)

	msgName := "ValidUintTagsOneOf1"
	runCases(t, data, msgName, cases)

	{
		err := data.Validate()
		require.Nil(t, err)
	}
}

func Test_GoValidator_ValidBoolTags1(t *testing.T) {
	data := &govalidatortest.ValidBoolTags1{}
	{
		err := data.Validate()
		require.NotNil(t, err)
	}

	cases := []*CaseDesc{
		{
			Name:       fmt.Sprintf("test tag <%s> general", protovalidator.TagBoolEq),
			FieldDesc:  "field 't_bool_general_eq1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagBoolEq, Value: true},
			FieldValue: func() interface{} { return data.TBoolGeneralEq1 },
			BeforeFunc: func() {},
			AfterFunc:  func() { data.TBoolGeneralEq1 = true },
		},
		{
			Name:       fmt.Sprintf("test tag <%s> optional", protovalidator.TagBoolEq),
			FieldDesc:  "field 't_bool_optional_eq1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagBoolEq, Value: true},
			FieldValue: func() interface{} { return nil },
			BeforeFunc: func() {},
			AfterFunc:  func() { x := true; data.TBoolOptionalEq1 = &x },
		},
		{
			Name:       fmt.Sprintf("test tag <%s> oneof", protovalidator.TagBoolEq),
			FieldDesc:  "field 't_bool_oneof_eq1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagBoolEq, Value: true},
			FieldValue: func() interface{} { return false },
			BeforeFunc: func() { data.OneTyp1 = &govalidatortest.ValidBoolTags1_TBoolOneofEq1{TBoolOneofEq1: false} },
			AfterFunc:  func() { data.OneTyp1 = &govalidatortest.ValidBoolTags1_TBoolOneofEq1{TBoolOneofEq1: true} },
		},
	}

	msgName := "ValidBoolTags1"
	runCases(t, data, msgName, cases)

	{
		err := data.Validate()
		require.Nil(t, err)
	}
}

func Test_GoValidator_ValidMessageTags(t *testing.T) {
	data := &govalidatortest.ValidMessageTags{}
	{
		err := data.Validate()
		require.NotNil(t, err)
	}

	cases := []*CaseDesc{
		{
			Name:       fmt.Sprintf("test tag <%s> in general 1", protovalidator.TagMessageNotNull),
			FieldDesc:  "field 't_message_general_not_null1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagMessageNotNull, Value: nil},
			FieldValue: func() interface{} { return data.TMessageGeneralNotNull1 },
			BeforeFunc: func() {},
			AfterFunc:  func() { data.TMessageGeneralNotNull1 = &govalidatortest.Config{Ip: "127.0.0.1", Port: 8080} },
			UseError2:  true,
		},
		{
			Name:       "test message skip in general 1",
			FieldDesc:  "",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagMessageSkip, Value: nil},
			FieldValue: func() interface{} { return data.TMessageGeneralSkip1 },
			BeforeFunc: func() { data.TMessageGeneralSkip1 = &govalidatortest.Config{} },
			AfterFunc:  func() { data.TMessageGeneralSkip1 = &govalidatortest.Config{Ip: "127.0.0.1", Port: 8080} },
		},
		{
			Name:       fmt.Sprintf("test tag <%s> in oneof 1", protovalidator.TagMessageNotNull),
			FieldDesc:  "field 't_message_oneof_not_null1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagMessageNotNull, Value: nil},
			FieldValue: func() interface{} { return nil },
			BeforeFunc: func() {
				data.OneTyp1 = &govalidatortest.ValidMessageTags_TMessageOneofNotNull1{TMessageOneofNotNull1: nil}
			},
			AfterFunc: func() {
				data.OneTyp1 = &govalidatortest.ValidMessageTags_TMessageOneofNotNull1{TMessageOneofNotNull1: &govalidatortest.Config{Ip: "127.0.0.1", Port: 8080}}
			},
			UseError2: true,
		},
		{
			Name:       "test message skip in oneof 1",
			FieldDesc:  "",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagMessageSkip, Value: nil},
			FieldValue: func() interface{} { return nil },
			BeforeFunc: func() {
				data.OneTyp1 = &govalidatortest.ValidMessageTags_TMessageOneofSkip1{TMessageOneofSkip1: &govalidatortest.Config{}}
			},
			AfterFunc: func() {
				data.OneTyp1 = &govalidatortest.ValidMessageTags_TMessageOneofSkip1{TMessageOneofSkip1: &govalidatortest.Config{Ip: "127.0.0.1", Port: 8080}}
			},
		},
	}

	msgName := "ValidMessageTags"
	runCases(t, data, msgName, cases)

	{
		// Test message skip
		data.TMessageGeneralSkip2 = &govalidatortest.Config{}
		data.OneTyp1 = &govalidatortest.ValidMessageTags_TMessageOneofSkip2{TMessageOneofSkip2: &govalidatortest.Config{}}

		err := data.Validate()
		require.Nil(t, err)
	}
}

func Test_GoValidator_ValidEnumTagsGeneral1(t *testing.T) {
	data := &govalidatortest.ValidEnumTagsGeneral1{}
	{
		err := data.Validate()
		require.NotNil(t, err)
	}

	cases := []*CaseDesc{
		{
			Name:       fmt.Sprintf("test tag <%s>", protovalidator.TagEnumEq),
			FieldDesc:  "field 't_enum_eq1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagEnumEq, Value: int32(govalidatortest.Enum1_January)},
			FieldValue: func() interface{} { return data.TEnumEq1 },
			BeforeFunc: func() { data.TEnumEq1 = govalidatortest.Enum1_April },
			AfterFunc:  func() { data.TEnumEq1 = govalidatortest.Enum1_January },
		},
		{
			Name:       fmt.Sprintf("test tag <%s>", protovalidator.TagEnumNe),
			FieldDesc:  "field 't_enum_ne1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagEnumNe, Value: int32(govalidatortest.Enum1_February)},
			FieldValue: func() interface{} { return data.TEnumNe1 },
			BeforeFunc: func() { data.TEnumNe1 = govalidatortest.Enum1_February },
			AfterFunc:  func() { data.TEnumNe1 = govalidatortest.Enum1_January },
		},
		{
			Name:       fmt.Sprintf("test tag <%s>", protovalidator.TagEnumLt),
			FieldDesc:  "field 't_enum_lt1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagEnumLt, Value: int32(govalidatortest.Enum1_March)},
			FieldValue: func() interface{} { return data.TEnumLt1 },
			BeforeFunc: func() { data.TEnumLt1 = govalidatortest.Enum1_April },
			AfterFunc:  func() { data.TEnumLt1 = govalidatortest.Enum1_January },
		},
		{
			Name:       fmt.Sprintf("test tag <%s>", protovalidator.TagEnumGt),
			FieldDesc:  "field 't_enum_gt1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagEnumGt, Value: int32(govalidatortest.Enum1_April)},
			FieldValue: func() interface{} { return data.TEnumGt1 },
			BeforeFunc: func() { data.TEnumGt1 = govalidatortest.Enum1_April },
			AfterFunc:  func() { data.TEnumGt1 = govalidatortest.Enum1_May },
		},
		{
			Name:       fmt.Sprintf("test tag <%s>", protovalidator.TagEnumLte),
			FieldDesc:  "field 't_enum_lte1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagEnumLte, Value: int32(govalidatortest.Enum1_May)},
			FieldValue: func() interface{} { return data.TEnumLte1 },
			BeforeFunc: func() { data.TEnumLte1 = govalidatortest.Enum1_June },
			AfterFunc:  func() { data.TEnumLte1 = govalidatortest.Enum1_May },
		},
		{
			Name:       fmt.Sprintf("test tag <%s>", protovalidator.TagEnumGte),
			FieldDesc:  "field 't_enum_gte1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagEnumGte, Value: int32(govalidatortest.Enum1_June)},
			FieldValue: func() interface{} { return data.TEnumGte1 },
			BeforeFunc: func() { data.TEnumGte1 = govalidatortest.Enum1_April },
			AfterFunc:  func() { data.TEnumGte1 = govalidatortest.Enum1_June },
		},
		{
			Name:       fmt.Sprintf("test tag <%s>", protovalidator.TagEnumIn),
			FieldDesc:  "field 't_enum_in1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagEnumIn, Value: []int32{int32(govalidatortest.Enum1_January), int32(govalidatortest.Enum1_February), int32(govalidatortest.Enum1_March)}},
			FieldValue: func() interface{} { return data.TEnumIn1 },
			BeforeFunc: func() { data.TEnumIn1 = govalidatortest.Enum1_June },
			AfterFunc:  func() { data.TEnumIn1 = govalidatortest.Enum1_February },
		},
		{
			Name:       fmt.Sprintf("test tag <%s>", protovalidator.TagEnumNotIn),
			FieldDesc:  "field 't_enum_not_in1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagEnumNotIn, Value: []int32{int32(govalidatortest.Enum1_April), int32(govalidatortest.Enum1_May), int32(govalidatortest.Enum1_June)}},
			FieldValue: func() interface{} { return data.TEnumNotIn1 },
			BeforeFunc: func() { data.TEnumNotIn1 = govalidatortest.Enum1_June },
			AfterFunc:  func() { data.TEnumNotIn1 = govalidatortest.Enum1_March },
		},
		{
			Name:       fmt.Sprintf("test tag <%s>", protovalidator.TagEnumInEnums),
			FieldDesc:  "field 't_enum_in_enums'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagEnumInEnums, Value: []int32{0, 1, 2, 3, 4, 8}},
			FieldValue: func() interface{} { return data.TEnumInEnums },
			BeforeFunc: func() { data.TEnumInEnums = govalidatortest.Enum1(10) },
			AfterFunc:  func() { data.TEnumInEnums = govalidatortest.Enum1_January },
		},
	}

	msgName := "ValidEnumTagsGeneral1"
	runCases(t, data, msgName, cases)

	{
		err := data.Validate()
		require.Nil(t, err)
	}
}

func Test_GoValidator_ValidEnumTagsOptional1(t *testing.T) {
	data := &govalidatortest.ValidEnumTagsOptional1{}
	{
		err := data.Validate()
		require.NotNil(t, err)
	}

	cases := []*CaseDesc{
		{
			Name:       fmt.Sprintf("test tag <%s>", protovalidator.TagEnumEq),
			FieldDesc:  "field 't_enum_eq1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagEnumEq, Value: int32(govalidatortest.Enum1_January)},
			FieldValue: func() interface{} { return nil },
			BeforeFunc: func() {},
			AfterFunc:  func() { x := govalidatortest.Enum1_January; data.TEnumEq1 = &x },
		},
		{
			Name:       fmt.Sprintf("test tag <%s>", protovalidator.TagEnumNe),
			FieldDesc:  "field 't_enum_ne1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagEnumNe, Value: int32(govalidatortest.Enum1_February)},
			FieldValue: func() interface{} { return data.TEnumNe1 },
			BeforeFunc: func() { x := govalidatortest.Enum1_February; data.TEnumNe1 = &x },
			AfterFunc:  func() { x := govalidatortest.Enum1_April; data.TEnumNe1 = &x },
		},
		{
			Name:       fmt.Sprintf("test tag <%s>", protovalidator.TagEnumLt),
			FieldDesc:  "field 't_enum_lt1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagEnumLt, Value: int32(govalidatortest.Enum1_March)},
			FieldValue: func() interface{} { return data.TEnumLt1 },
			BeforeFunc: func() { x := govalidatortest.Enum1_June; data.TEnumLt1 = &x },
			AfterFunc:  func() { x := govalidatortest.Enum1_February; data.TEnumLt1 = &x },
		},
		{
			Name:       fmt.Sprintf("test tag <%s>", protovalidator.TagEnumGt),
			FieldDesc:  "field 't_enum_gt1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagEnumGt, Value: int32(govalidatortest.Enum1_April)},
			FieldValue: func() interface{} { return data.TEnumGt1 },
			BeforeFunc: func() { x := govalidatortest.Enum1_February; data.TEnumGt1 = &x },
			AfterFunc:  func() { x := govalidatortest.Enum1_May; data.TEnumGt1 = &x },
		},
		{
			Name:       fmt.Sprintf("test tag <%s>", protovalidator.TagEnumLte),
			FieldDesc:  "field 't_enum_lte1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagEnumLte, Value: int32(govalidatortest.Enum1_May)},
			FieldValue: func() interface{} { return data.TEnumLte1 },
			BeforeFunc: func() { x := govalidatortest.Enum1_June; data.TEnumLte1 = &x },
			AfterFunc:  func() { x := govalidatortest.Enum1_May; data.TEnumLte1 = &x },
		},
		{
			Name:       fmt.Sprintf("test tag <%s>", protovalidator.TagEnumGte),
			FieldDesc:  "field 't_enum_gte1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagEnumGte, Value: int32(govalidatortest.Enum1_June)},
			FieldValue: func() interface{} { return nil },
			BeforeFunc: func() {},
			AfterFunc:  func() { x := govalidatortest.Enum1_June; data.TEnumGte1 = &x },
		},
		{
			Name:       fmt.Sprintf("test tag <%s>", protovalidator.TagEnumIn),
			FieldDesc:  "field 't_enum_in1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagEnumIn, Value: []int32{int32(govalidatortest.Enum1_January), int32(govalidatortest.Enum1_February), int32(govalidatortest.Enum1_March)}},
			FieldValue: func() interface{} { return data.TEnumIn1 },
			BeforeFunc: func() { x := govalidatortest.Enum1_April; data.TEnumIn1 = &x },
			AfterFunc:  func() { x := govalidatortest.Enum1_January; data.TEnumIn1 = &x },
		},
		{
			Name:       fmt.Sprintf("test tag <%s>", protovalidator.TagEnumNotIn),
			FieldDesc:  "field 't_enum_not_in1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagEnumNotIn, Value: []int32{int32(govalidatortest.Enum1_April), int32(govalidatortest.Enum1_May), int32(govalidatortest.Enum1_June)}},
			FieldValue: func() interface{} { return data.TEnumNotIn1 },
			BeforeFunc: func() { x := govalidatortest.Enum1_April; data.TEnumNotIn1 = &x },
			AfterFunc:  func() { x := govalidatortest.Enum1_March; data.TEnumNotIn1 = &x },
		},
		{
			Name:       fmt.Sprintf("test tag <%s>", protovalidator.TagEnumInEnums),
			FieldDesc:  "field 't_enum_in_enums'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagEnumInEnums, Value: []int32{0, 1, 2, 3, 4, 8}},
			FieldValue: func() interface{} { return data.TEnumInEnums },
			BeforeFunc: func() { x := govalidatortest.Enum1(10); data.TEnumInEnums = &x },
			AfterFunc:  func() { x := govalidatortest.Enum1_April; data.TEnumInEnums = &x },
		},
	}

	msgName := "ValidEnumTagsOptional1"
	runCases(t, data, msgName, cases)

	{
		err := data.Validate()
		require.Nil(t, err)
	}
}

func Test_GoValidator_ValidEnumTagsOneOf1(t *testing.T) {
	data := &govalidatortest.ValidEnumTagsOneOf1{}
	{
		err := data.Validate()
		require.Nil(t, err)
	}

	cases := []*CaseDesc{
		{
			Name:       fmt.Sprintf("test tag <%s>", protovalidator.TagEnumEq),
			FieldDesc:  "field 't_enum_eq1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagEnumEq, Value: int32(govalidatortest.Enum1_January)},
			FieldValue: func() interface{} { return data.OneTyp1.(*govalidatortest.ValidEnumTagsOneOf1_TEnumEq1).TEnumEq1 },
			BeforeFunc: func() {
				data.OneTyp1 = &govalidatortest.ValidEnumTagsOneOf1_TEnumEq1{TEnumEq1: govalidatortest.Enum1_April}
			},
			AfterFunc: func() {
				data.OneTyp1 = &govalidatortest.ValidEnumTagsOneOf1_TEnumEq1{TEnumEq1: govalidatortest.Enum1_January}
			},
		},
		{
			Name:       fmt.Sprintf("test tag <%s>", protovalidator.TagEnumNe),
			FieldDesc:  "field 't_enum_ne1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagEnumNe, Value: int32(govalidatortest.Enum1_February)},
			FieldValue: func() interface{} { return data.OneTyp1.(*govalidatortest.ValidEnumTagsOneOf1_TEnumNe1).TEnumNe1 },
			BeforeFunc: func() {
				data.OneTyp1 = &govalidatortest.ValidEnumTagsOneOf1_TEnumNe1{TEnumNe1: govalidatortest.Enum1_February}
			},
			AfterFunc: func() {
				data.OneTyp1 = &govalidatortest.ValidEnumTagsOneOf1_TEnumNe1{TEnumNe1: govalidatortest.Enum1_January}
			},
		},
		{
			Name:       fmt.Sprintf("test tag <%s>", protovalidator.TagEnumLt),
			FieldDesc:  "field 't_enum_lt1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagEnumLt, Value: int32(govalidatortest.Enum1_March)},
			FieldValue: func() interface{} { return data.OneTyp1.(*govalidatortest.ValidEnumTagsOneOf1_TEnumLt1).TEnumLt1 },
			BeforeFunc: func() {
				data.OneTyp1 = &govalidatortest.ValidEnumTagsOneOf1_TEnumLt1{TEnumLt1: govalidatortest.Enum1_March}
			},
			AfterFunc: func() {
				data.OneTyp1 = &govalidatortest.ValidEnumTagsOneOf1_TEnumLt1{TEnumLt1: govalidatortest.Enum1_January}
			},
		},
		{
			Name:       fmt.Sprintf("test tag <%s>", protovalidator.TagEnumGt),
			FieldDesc:  "field 't_enum_gt1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagEnumGt, Value: int32(govalidatortest.Enum1_April)},
			FieldValue: func() interface{} { return data.OneTyp1.(*govalidatortest.ValidEnumTagsOneOf1_TEnumGt1).TEnumGt1 },
			BeforeFunc: func() {
				data.OneTyp1 = &govalidatortest.ValidEnumTagsOneOf1_TEnumGt1{TEnumGt1: govalidatortest.Enum1_February}
			},
			AfterFunc: func() {
				data.OneTyp1 = &govalidatortest.ValidEnumTagsOneOf1_TEnumGt1{TEnumGt1: govalidatortest.Enum1_May}
			},
		},
		{
			Name:       fmt.Sprintf("test tag <%s>", protovalidator.TagEnumLte),
			FieldDesc:  "field 't_enum_lte1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagEnumLte, Value: int32(govalidatortest.Enum1_May)},
			FieldValue: func() interface{} { return data.OneTyp1.(*govalidatortest.ValidEnumTagsOneOf1_TEnumLte1).TEnumLte1 },
			BeforeFunc: func() {
				data.OneTyp1 = &govalidatortest.ValidEnumTagsOneOf1_TEnumLte1{TEnumLte1: govalidatortest.Enum1_June}
			},
			AfterFunc: func() {
				data.OneTyp1 = &govalidatortest.ValidEnumTagsOneOf1_TEnumLte1{TEnumLte1: govalidatortest.Enum1_May}
			},
		},
		{
			Name:       fmt.Sprintf("test tag <%s>", protovalidator.TagEnumGte),
			FieldDesc:  "field 't_enum_gte1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagEnumGte, Value: int32(govalidatortest.Enum1_June)},
			FieldValue: func() interface{} { return data.OneTyp1.(*govalidatortest.ValidEnumTagsOneOf1_TEnumGte1).TEnumGte1 },
			BeforeFunc: func() {
				data.OneTyp1 = &govalidatortest.ValidEnumTagsOneOf1_TEnumGte1{TEnumGte1: govalidatortest.Enum1_January}
			},
			AfterFunc: func() {
				data.OneTyp1 = &govalidatortest.ValidEnumTagsOneOf1_TEnumGte1{TEnumGte1: govalidatortest.Enum1_June}
			},
		},
		{
			Name:       fmt.Sprintf("test tag <%s>", protovalidator.TagEnumIn),
			FieldDesc:  "field 't_enum_in1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagEnumIn, Value: []int32{int32(govalidatortest.Enum1_January), int32(govalidatortest.Enum1_February), int32(govalidatortest.Enum1_March)}},
			FieldValue: func() interface{} { return data.OneTyp1.(*govalidatortest.ValidEnumTagsOneOf1_TEnumIn1).TEnumIn1 },
			BeforeFunc: func() {
				data.OneTyp1 = &govalidatortest.ValidEnumTagsOneOf1_TEnumIn1{TEnumIn1: govalidatortest.Enum1_April}
			},
			AfterFunc: func() {
				data.OneTyp1 = &govalidatortest.ValidEnumTagsOneOf1_TEnumIn1{TEnumIn1: govalidatortest.Enum1_January}
			},
		},
		{
			Name:      fmt.Sprintf("test tag <%s>", protovalidator.TagEnumNotIn),
			FieldDesc: "field 't_enum_not_in1'",
			TagInfo:   &protovalidator.TagInfo{Tag: protovalidator.TagEnumNotIn, Value: []int32{int32(govalidatortest.Enum1_April), int32(govalidatortest.Enum1_May), int32(govalidatortest.Enum1_June)}},
			FieldValue: func() interface{} {
				return data.OneTyp1.(*govalidatortest.ValidEnumTagsOneOf1_TEnumNotIn1).TEnumNotIn1
			},
			BeforeFunc: func() {
				data.OneTyp1 = &govalidatortest.ValidEnumTagsOneOf1_TEnumNotIn1{TEnumNotIn1: govalidatortest.Enum1_April}
			},
			AfterFunc: func() {
				data.OneTyp1 = &govalidatortest.ValidEnumTagsOneOf1_TEnumNotIn1{TEnumNotIn1: govalidatortest.Enum1_January}
			},
		},
		{
			Name:      fmt.Sprintf("test tag <%s>", protovalidator.TagEnumInEnums),
			FieldDesc: "field 't_enum_in_enums'",
			TagInfo:   &protovalidator.TagInfo{Tag: protovalidator.TagEnumInEnums, Value: []int32{0, 1, 2, 3, 4, 8}},
			FieldValue: func() interface{} {
				return data.OneTyp1.(*govalidatortest.ValidEnumTagsOneOf1_TEnumInEnums).TEnumInEnums
			},
			BeforeFunc: func() {
				data.OneTyp1 = &govalidatortest.ValidEnumTagsOneOf1_TEnumInEnums{TEnumInEnums: govalidatortest.Enum1(10)}
			},
			AfterFunc: func() {
				data.OneTyp1 = &govalidatortest.ValidEnumTagsOneOf1_TEnumInEnums{TEnumInEnums: govalidatortest.Enum1_January}
			},
		},
	}

	msgName := "ValidEnumTagsOneOf1"
	runCases(t, data, msgName, cases)

	{
		err := data.Validate()
		require.Nil(t, err)
	}
}

func Test_GoValidator_ValidBytesTags1(t *testing.T) {
	data := &govalidatortest.ValidBytesTags1{}
	{
		err := data.Validate()
		require.NotNil(t, err)
	}

	cases := []*CaseDesc{
		{
			Name:       fmt.Sprintf("test tag <%s>", protovalidator.TagBytesLenEq),
			FieldDesc:  "field 't_bytes_len_eq1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagBytesLenEq, Value: 1},
			FieldValue: func() interface{} { return len(data.TBytesLenEq1) },
			BeforeFunc: func() {},
			AfterFunc:  func() { data.TBytesLenEq1 = []byte("h") },
		},
		{
			Name:       fmt.Sprintf("test tag <%s>", protovalidator.TagBytesLenNe),
			FieldDesc:  "field 't_bytes_len_ne1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagBytesLenNe, Value: 2},
			FieldValue: func() interface{} { return len(data.TBytesLenNe1) },
			BeforeFunc: func() { data.TBytesLenNe1 = []byte("he") },
			AfterFunc:  func() { data.TBytesLenNe1 = []byte("hello") },
		},
		{
			Name:       fmt.Sprintf("test tag <%s>", protovalidator.TagBytesLenLt),
			FieldDesc:  "field 't_bytes_len_lt1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagBytesLenLt, Value: 3},
			FieldValue: func() interface{} { return len(data.TBytesLenLt1) },
			BeforeFunc: func() { data.TBytesLenLt1 = []byte("hello") },
			AfterFunc:  func() { data.TBytesLenLt1 = []byte("he") },
		},
		{
			Name:       fmt.Sprintf("test tag <%s>", protovalidator.TagBytesLenGt),
			FieldDesc:  "field 't_bytes_len_gt1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagBytesLenGt, Value: 4},
			FieldValue: func() interface{} { return len(data.TBytesLenGt1) },
			BeforeFunc: func() { data.TBytesLenGt1 = []byte("he") },
			AfterFunc:  func() { data.TBytesLenGt1 = []byte("hello") },
		},
		{
			Name:       fmt.Sprintf("test tag <%s>", protovalidator.TagBytesLenLte),
			FieldDesc:  "field 't_bytes_len_lte1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagBytesLenLte, Value: 5},
			FieldValue: func() interface{} { return len(data.TBytesLenLte1) },
			BeforeFunc: func() { data.TBytesLenLte1 = []byte("hello world") },
			AfterFunc:  func() { data.TBytesLenLte1 = []byte("hello") },
		},
		{
			Name:       fmt.Sprintf("test tag <%s>", protovalidator.TagBytesLenGte),
			FieldDesc:  "field 't_bytes_len_gte1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagBytesLenGte, Value: 6},
			FieldValue: func() interface{} { return len(data.TBytesLenGte1) },
			BeforeFunc: func() {},
			AfterFunc:  func() { data.TBytesLenGte1 = []byte("hello w") },
		},
	}

	msgName := "ValidBytesTags1"
	runCases(t, data, msgName, cases)

	{
		err := data.Validate()
		require.Nil(t, err)
	}
}

func Test_GoValidator_ValidRepeatedTagsGeneral1(t *testing.T) {
	data := &govalidatortest.ValidRepeatedTagsGeneral1{}
	{
		err := data.Validate()
		require.NotNil(t, err)
	}

	cases := []*CaseDesc{
		{
			Name:       fmt.Sprintf("test tag <%s>", protovalidator.TagRepeatedNotNull),
			FieldDesc:  "field 't_list_not_null1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagRepeatedNotNull, Value: nil},
			FieldValue: func() interface{} { return nil },
			BeforeFunc: func() {},
			AfterFunc:  func() { data.TListNotNull1 = []string{"s1"} },
			UseError2:  true,
		},
		{
			Name:       fmt.Sprintf("test tag <%s>", protovalidator.TagRepeatedLenEq),
			FieldDesc:  "field 't_list_len_eq1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagRepeatedLenEq, Value: 1},
			FieldValue: func() interface{} { return len(data.TListLenEq1) },
			BeforeFunc: func() {},
			AfterFunc:  func() { data.TListLenEq1 = []string{"s1"} },
		},
		{
			Name:       fmt.Sprintf("test tag <%s>", protovalidator.TagRepeatedLenNe),
			FieldDesc:  "field 't_list_len_ne1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagRepeatedLenNe, Value: 2},
			FieldValue: func() interface{} { return len(data.TListLenNe1) },
			BeforeFunc: func() { data.TListLenNe1 = []string{"s1", "s2"} },
			AfterFunc:  func() { data.TListLenNe1 = []string{"s1"} },
		},
		{
			Name:       fmt.Sprintf("test tag <%s>", protovalidator.TagRepeatedLenLt),
			FieldDesc:  "field 't_list_len_lt1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagRepeatedLenLt, Value: 3},
			FieldValue: func() interface{} { return len(data.TListLenLt1) },
			BeforeFunc: func() { data.TListLenLt1 = []string{"s1", "s2", "s3", "s4"} },
			AfterFunc:  func() { data.TListLenLt1 = []string{"s1"} },
		},
		{
			Name:       fmt.Sprintf("test tag <%s>", protovalidator.TagRepeatedLenGt),
			FieldDesc:  "field 't_list_len_gt1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagRepeatedLenGt, Value: 4},
			FieldValue: func() interface{} { return len(data.TListLenGt1) },
			BeforeFunc: func() { data.TListLenGt1 = []string{"s1"} },
			AfterFunc:  func() { data.TListLenGt1 = []string{"s1", "s2", "s3", "s4", "s5"} },
		},
		{
			Name:       fmt.Sprintf("test tag <%s>", protovalidator.TagRepeatedLenLte),
			FieldDesc:  "field 't_list_len_lte1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagRepeatedLenLte, Value: 5},
			FieldValue: func() interface{} { return len(data.TListLenLte1) },
			BeforeFunc: func() { data.TListLenLte1 = []string{"s1", "s2", "s3", "s4", "s5", "s6"} },
			AfterFunc:  func() { data.TListLenLte1 = []string{"s1", "s2", "s3", "s4", "s5"} },
		},
		{
			Name:       fmt.Sprintf("test tag <%s>", protovalidator.TagRepeatedLenGte),
			FieldDesc:  "field 't_list_len_gte1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagRepeatedLenGte, Value: 6},
			FieldValue: func() interface{} { return len(data.TListLenGte1) },
			BeforeFunc: func() {},
			AfterFunc:  func() { data.TListLenGte1 = []string{"s1", "s2", "s3", "s4", "s5", "s6"} },
		},
		{
			Name:       fmt.Sprintf("test tag <%s> with string", protovalidator.TagRepeatedUnique),
			FieldDesc:  "field 't_list_unique_string'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagRepeatedUnique, Value: nil},
			FieldValue: func() interface{} { return data.TListUniqueString },
			BeforeFunc: func() { data.TListUniqueString = []string{"s1", "s2", "s1"} },
			AfterFunc:  func() { data.TListUniqueString = []string{"s1", "s2", "s3"} },
			UseError2:  true,
		},
		{
			Name:       fmt.Sprintf("test tag <%s> with double", protovalidator.TagRepeatedUnique),
			FieldDesc:  "field 't_list_unique_double'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagRepeatedUnique, Value: nil},
			FieldValue: func() interface{} { return data.TListUniqueDouble },
			BeforeFunc: func() { data.TListUniqueDouble = []float64{1.1, 1.2, 1.1} },
			AfterFunc:  func() { data.TListUniqueDouble = []float64{1.1, 1.2, 1.3} },
			UseError2:  true,
		},
		{
			Name:       fmt.Sprintf("test tag <%s> with float", protovalidator.TagRepeatedUnique),
			FieldDesc:  "field 't_list_unique_float'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagRepeatedUnique, Value: nil},
			FieldValue: func() interface{} { return data.TListUniqueFloat },
			BeforeFunc: func() { data.TListUniqueFloat = []float32{1.1, 1.2, 1.1} },
			AfterFunc:  func() { data.TListUniqueFloat = []float32{1.1, 1.2, 1.3} },
			UseError2:  true,
		},
		{
			Name:       fmt.Sprintf("test tag <%s> with int32", protovalidator.TagRepeatedUnique),
			FieldDesc:  "field 't_list_unique_int32'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagRepeatedUnique, Value: nil},
			FieldValue: func() interface{} { return data.TListUniqueInt32 },
			BeforeFunc: func() { data.TListUniqueInt32 = []int32{1, 2, 3, 1} },
			AfterFunc:  func() { data.TListUniqueInt32 = []int32{1, 2, 3, 4} },
			UseError2:  true,
		},
		{
			Name:       fmt.Sprintf("test tag <%s> with int64", protovalidator.TagRepeatedUnique),
			FieldDesc:  "field 't_list_unique_int64'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagRepeatedUnique, Value: nil},
			FieldValue: func() interface{} { return data.TListUniqueInt64 },
			BeforeFunc: func() { data.TListUniqueInt64 = []int64{1, 2, 3, 1, 4} },
			AfterFunc:  func() { data.TListUniqueInt64 = []int64{1, 2, 3, 5, 4} },
			UseError2:  true,
		},
		{
			Name:       fmt.Sprintf("test tag <%s> with uint32", protovalidator.TagRepeatedUnique),
			FieldDesc:  "field 't_list_unique_uint32'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagRepeatedUnique, Value: nil},
			FieldValue: func() interface{} { return data.TListUniqueUint32 },
			BeforeFunc: func() { data.TListUniqueUint32 = []uint32{1, 1, 1, 1} },
			AfterFunc:  func() { data.TListUniqueUint32 = []uint32{1, 2, 3, 4} },
			UseError2:  true,
		},
		{
			Name:       fmt.Sprintf("test tag <%s> with uint64", protovalidator.TagRepeatedUnique),
			FieldDesc:  "field 't_list_unique_uint64'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagRepeatedUnique, Value: nil},
			FieldValue: func() interface{} { return data.TListUniqueUint64 },
			BeforeFunc: func() { data.TListUniqueUint64 = []uint64{1, 2, 2, 3} },
			AfterFunc:  func() { data.TListUniqueUint64 = []uint64{1, 3, 4, 5} },
			UseError2:  true,
		},
		{
			Name:       fmt.Sprintf("test tag <%s> with sint32", protovalidator.TagRepeatedUnique),
			FieldDesc:  "field 't_list_unique_sint32'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagRepeatedUnique, Value: nil},
			FieldValue: func() interface{} { return data.TListUniqueSint32 },
			BeforeFunc: func() { data.TListUniqueSint32 = []int32{1, 2, 3, 1} },
			AfterFunc:  func() { data.TListUniqueSint32 = []int32{4, 2, 3, 1} },
			UseError2:  true,
		},
		{
			Name:       fmt.Sprintf("test tag <%s> with sint64", protovalidator.TagRepeatedUnique),
			FieldDesc:  "field 't_list_unique_sint64'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagRepeatedUnique, Value: nil},
			FieldValue: func() interface{} { return data.TListUniqueSint64 },
			BeforeFunc: func() { data.TListUniqueSint64 = []int64{1, 2, 3, 1, 4} },
			AfterFunc:  func() { data.TListUniqueSint64 = []int64{1, 2, 3, 7, 4} },
			UseError2:  true,
		},
		{
			Name:       fmt.Sprintf("test tag <%s> with sfixed32", protovalidator.TagRepeatedUnique),
			FieldDesc:  "field 't_list_unique_sfixed32'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagRepeatedUnique, Value: nil},
			FieldValue: func() interface{} { return data.TListUniqueSfixed32 },
			BeforeFunc: func() { data.TListUniqueSfixed32 = []int32{1, 2, 3, 1} },
			AfterFunc:  func() { data.TListUniqueSfixed32 = []int32{4, 2, 3, 1} },
			UseError2:  true,
		},
		{
			Name:       fmt.Sprintf("test tag <%s> with sfixed64", protovalidator.TagRepeatedUnique),
			FieldDesc:  "field 't_list_unique_sfixed64'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagRepeatedUnique, Value: nil},
			FieldValue: func() interface{} { return data.TListUniqueSfixed64 },
			BeforeFunc: func() { data.TListUniqueSfixed64 = []int64{1, 2, 3, 1, 4} },
			AfterFunc:  func() { data.TListUniqueSfixed64 = []int64{1, 2, 3, 7, 4} },
			UseError2:  true,
		},
		{
			Name:       fmt.Sprintf("test tag <%s> with fixed32", protovalidator.TagRepeatedUnique),
			FieldDesc:  "field 't_list_unique_fixed32'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagRepeatedUnique, Value: nil},
			FieldValue: func() interface{} { return data.TListUniqueFixed32 },
			BeforeFunc: func() { data.TListUniqueFixed32 = []uint32{1, 1, 1, 1} },
			AfterFunc:  func() { data.TListUniqueFixed32 = []uint32{1, 2, 3, 4} },
			UseError2:  true,
		},
		{
			Name:       fmt.Sprintf("test tag <%s> with fixed64", protovalidator.TagRepeatedUnique),
			FieldDesc:  "field 't_list_unique_fixed64'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagRepeatedUnique, Value: nil},
			FieldValue: func() interface{} { return data.TListUniqueFixed64 },
			BeforeFunc: func() { data.TListUniqueFixed64 = []uint64{1, 2, 2, 3} },
			AfterFunc:  func() { data.TListUniqueFixed64 = []uint64{1, 3, 4, 5} },
			UseError2:  true,
		},
		{
			Name:       fmt.Sprintf("test tag <%s> with bool", protovalidator.TagRepeatedUnique),
			FieldDesc:  "field 't_list_unique_bool'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagRepeatedUnique, Value: nil},
			FieldValue: func() interface{} { return data.TListUniqueBool },
			BeforeFunc: func() { data.TListUniqueBool = []bool{true, true, false, true} },
			AfterFunc:  func() { data.TListUniqueBool = []bool{false, true} },
			UseError2:  true,
		},
		{
			Name:       fmt.Sprintf("test tag <%s> with enum", protovalidator.TagRepeatedUnique),
			FieldDesc:  "field 't_list_unique_enum'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagRepeatedUnique, Value: nil},
			FieldValue: func() interface{} { return []int32{1, 8, 1} },
			BeforeFunc: func() {
				data.TListUniqueEnum = []govalidatortest.Enum1{govalidatortest.Enum1_February, govalidatortest.Enum1_June, govalidatortest.Enum1_February}
			},
			AfterFunc: func() { data.TListUniqueEnum = nil },
			UseError2: true,
		},
		{
			Name:       fmt.Sprintf("test tag <%s> with message", protovalidator.TagRepeatedUnique),
			FieldDesc:  "field 't_list_unique_message'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagRepeatedUnique, Value: nil},
			FieldValue: func() interface{} { return nil },
			BeforeFunc: func() {
				m := &govalidatortest.Config{}
				data.TListUniqueMessage = []*govalidatortest.Config{m, m}
			},
			AfterFunc: func() {
				m1 := &govalidatortest.Config{}
				m2 := &govalidatortest.Config{}
				data.TListUniqueMessage = []*govalidatortest.Config{m1, m2}
			},
			UseError2: true,
		},
		{
			Name:       fmt.Sprintf("test tag <%s> with bytes", protovalidator.TagRepeatedUnique),
			FieldDesc:  "field 't_list_unique_bytes'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagRepeatedUnique, Value: nil},
			FieldValue: func() interface{} { return nil },
			BeforeFunc: func() {
				m := []byte{1, 2, 3}
				data.TListUniqueBytes = [][]byte{m, m}
			},
			AfterFunc: func() {
				m1 := []byte{1, 2, 3}
				m2 := []byte{4, 2, 3}
				data.TListUniqueBytes = [][]byte{m1, m2}
			},
			UseError2: true,
		},
	}

	msgName := "ValidRepeatedTagsGeneral1"
	runCases(t, data, msgName, cases)

	{
		err := data.Validate()
		require.Nil(t, err)
	}
}

func Test_GoValidator_ValidRepeatedTagsItem1(t *testing.T) {
	data := &govalidatortest.ValidRepeatedTagsItem1{}
	{
		err := data.Validate()
		require.NotNil(t, err)
	}

	cases := []*CaseDesc{
		{
			Name:       fmt.Sprintf("test tag <%s> with string 1", protovalidator.TagRepeatedItem),
			FieldDesc:  "array item where in field 't_list_item_string'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagStringPrefix, Value: "id-"},
			FieldValue: func() interface{} { return "k1" },
			BeforeFunc: func() { data.TListItemString = []string{"k1", "id-k2"} },
			AfterFunc:  func() { data.TListItemString = []string{"id-k1", "id-k2"} },
		},

		{
			Name:       fmt.Sprintf("test tag <%s> with double 1", protovalidator.TagRepeatedItem),
			FieldDesc:  "array item where in field 't_list_item_double'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagFloatGte, Value: 1.1},
			FieldValue: func() interface{} { return 0.8 },
			BeforeFunc: func() { data.TListItemDouble = []float64{0.8, 5.5} },
			AfterFunc:  func() { data.TListItemDouble = []float64{1.1, 10} },
		},
		{
			Name:       fmt.Sprintf("test tag <%s> with double 2", protovalidator.TagRepeatedItem),
			FieldDesc:  "array item where in field 't_list_item_double'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagFloatLte, Value: 11.1},
			FieldValue: func() interface{} { return 12 },
			BeforeFunc: func() { data.TListItemDouble = []float64{1.1, 12} },
			AfterFunc:  func() { data.TListItemDouble = []float64{1.1, 10} },
		},

		{
			Name:       fmt.Sprintf("test tag <%s> with float 1", protovalidator.TagRepeatedItem),
			FieldDesc:  "array item where in field 't_list_item_float'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagFloatGte, Value: 1.1},
			FieldValue: func() interface{} { return float32(0.8) },
			BeforeFunc: func() { data.TListItemFloat = []float32{0.8, 5.5} },
			AfterFunc:  func() { data.TListItemFloat = []float32{1.1, 10} },
		},
		{
			Name:       fmt.Sprintf("test tag <%s> with float 2", protovalidator.TagRepeatedItem),
			FieldDesc:  "array item where in field 't_list_item_float'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagFloatLte, Value: 11.1},
			FieldValue: func() interface{} { return float32(12) },
			BeforeFunc: func() { data.TListItemFloat = []float32{1.1, 12} },
			AfterFunc:  func() { data.TListItemFloat = []float32{1.1, 10} },
		},

		{
			Name:       fmt.Sprintf("test tag <%s> with int32 1", protovalidator.TagRepeatedItem),
			FieldDesc:  "array item where in field 't_list_item_int32'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagIntGte, Value: 2},
			FieldValue: func() interface{} { return 1 },
			BeforeFunc: func() { data.TListItemInt32 = []int32{1, 5} },
			AfterFunc:  func() { data.TListItemInt32 = []int32{2, 10} },
		},
		{
			Name:       fmt.Sprintf("test tag <%s> with int32 2", protovalidator.TagRepeatedItem),
			FieldDesc:  "array item where in field 't_list_item_int32'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagIntLte, Value: 10},
			FieldValue: func() interface{} { return 12 },
			BeforeFunc: func() { data.TListItemInt32 = []int32{2, 12} },
			AfterFunc:  func() { data.TListItemInt32 = []int32{2, 10} },
		},

		{
			Name:       fmt.Sprintf("test tag <%s> with int64 1", protovalidator.TagRepeatedItem),
			FieldDesc:  "array item where in field 't_list_item_int64'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagIntGte, Value: 2},
			FieldValue: func() interface{} { return 1 },
			BeforeFunc: func() { data.TListItemInt64 = []int64{1, 5} },
			AfterFunc:  func() { data.TListItemInt64 = []int64{2, 10} },
		},
		{
			Name:       fmt.Sprintf("test tag <%s> with int64 2", protovalidator.TagRepeatedItem),
			FieldDesc:  "array item where in field 't_list_item_int64'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagIntLte, Value: 10},
			FieldValue: func() interface{} { return 12 },
			BeforeFunc: func() { data.TListItemInt64 = []int64{2, 12} },
			AfterFunc:  func() { data.TListItemInt64 = []int64{2, 10} },
		},

		{
			Name:       fmt.Sprintf("test tag <%s> with sint32 1", protovalidator.TagRepeatedItem),
			FieldDesc:  "array item where in field 't_list_item_sint32'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagIntGte, Value: 2},
			FieldValue: func() interface{} { return 1 },
			BeforeFunc: func() { data.TListItemSint32 = []int32{1, 5} },
			AfterFunc:  func() { data.TListItemSint32 = []int32{2, 10} },
		},
		{
			Name:       fmt.Sprintf("test tag <%s> with sint32 2", protovalidator.TagRepeatedItem),
			FieldDesc:  "array item where in field 't_list_item_sint32'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagIntLte, Value: 10},
			FieldValue: func() interface{} { return 12 },
			BeforeFunc: func() { data.TListItemSint32 = []int32{2, 12} },
			AfterFunc:  func() { data.TListItemSint32 = []int32{2, 10} },
		},

		{
			Name:       fmt.Sprintf("test tag <%s> with sint64 1", protovalidator.TagRepeatedItem),
			FieldDesc:  "array item where in field 't_list_item_sint64'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagIntGte, Value: 2},
			FieldValue: func() interface{} { return 1 },
			BeforeFunc: func() { data.TListItemSint64 = []int64{1, 5} },
			AfterFunc:  func() { data.TListItemSint64 = []int64{2, 10} },
		},
		{
			Name:       fmt.Sprintf("test tag <%s> with sint64 2", protovalidator.TagRepeatedItem),
			FieldDesc:  "array item where in field 't_list_item_sint64'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagIntLte, Value: 10},
			FieldValue: func() interface{} { return 12 },
			BeforeFunc: func() { data.TListItemSint64 = []int64{2, 12} },
			AfterFunc:  func() { data.TListItemSint64 = []int64{2, 10} },
		},

		{
			Name:       fmt.Sprintf("test tag <%s> with sfixed32 1", protovalidator.TagRepeatedItem),
			FieldDesc:  "array item where in field 't_list_item_sfixed32'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagIntGte, Value: 2},
			FieldValue: func() interface{} { return 1 },
			BeforeFunc: func() { data.TListItemSfixed32 = []int32{1, 5} },
			AfterFunc:  func() { data.TListItemSfixed32 = []int32{2, 10} },
		},
		{
			Name:       fmt.Sprintf("test tag <%s> with sfixed32 2", protovalidator.TagRepeatedItem),
			FieldDesc:  "array item where in field 't_list_item_sfixed32'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagIntLte, Value: 10},
			FieldValue: func() interface{} { return 12 },
			BeforeFunc: func() { data.TListItemSfixed32 = []int32{2, 12} },
			AfterFunc:  func() { data.TListItemSfixed32 = []int32{2, 10} },
		},

		{
			Name:       fmt.Sprintf("test tag <%s> with sfixed64 1", protovalidator.TagRepeatedItem),
			FieldDesc:  "array item where in field 't_list_item_sfixed64'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagIntGte, Value: 2},
			FieldValue: func() interface{} { return 1 },
			BeforeFunc: func() { data.TListItemSfixed64 = []int64{1, 5} },
			AfterFunc:  func() { data.TListItemSfixed64 = []int64{2, 10} },
		},
		{
			Name:       fmt.Sprintf("test tag <%s> with sfixed64 2", protovalidator.TagRepeatedItem),
			FieldDesc:  "array item where in field 't_list_item_sfixed64'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagIntLte, Value: 10},
			FieldValue: func() interface{} { return 12 },
			BeforeFunc: func() { data.TListItemSfixed64 = []int64{2, 12} },
			AfterFunc:  func() { data.TListItemSfixed64 = []int64{2, 10} },
		},

		{
			Name:       fmt.Sprintf("test tag <%s> with uint32 1", protovalidator.TagRepeatedItem),
			FieldDesc:  "array item where in field 't_list_item_uint32'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagIntGte, Value: 2},
			FieldValue: func() interface{} { return 1 },
			BeforeFunc: func() { data.TListItemUint32 = []uint32{1, 5} },
			AfterFunc:  func() { data.TListItemUint32 = []uint32{2, 10} },
		},
		{
			Name:       fmt.Sprintf("test tag <%s> with uint32 2", protovalidator.TagRepeatedItem),
			FieldDesc:  "array item where in field 't_list_item_uint32'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagIntLte, Value: 10},
			FieldValue: func() interface{} { return 12 },
			BeforeFunc: func() { data.TListItemUint32 = []uint32{2, 12} },
			AfterFunc:  func() { data.TListItemUint32 = []uint32{2, 10} },
		},

		{
			Name:       fmt.Sprintf("test tag <%s> with uint64 1", protovalidator.TagRepeatedItem),
			FieldDesc:  "array item where in field 't_list_item_uint64'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagIntGte, Value: 2},
			FieldValue: func() interface{} { return 1 },
			BeforeFunc: func() { data.TListItemUint64 = []uint64{1, 5} },
			AfterFunc:  func() { data.TListItemUint64 = []uint64{2, 10} },
		},
		{
			Name:       fmt.Sprintf("test tag <%s> with uint64 2", protovalidator.TagRepeatedItem),
			FieldDesc:  "array item where in field 't_list_item_uint64'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagIntLte, Value: 10},
			FieldValue: func() interface{} { return 12 },
			BeforeFunc: func() { data.TListItemUint64 = []uint64{2, 12} },
			AfterFunc:  func() { data.TListItemUint64 = []uint64{2, 10} },
		},

		{
			Name:       fmt.Sprintf("test tag <%s> with fixed32 1", protovalidator.TagRepeatedItem),
			FieldDesc:  "array item where in field 't_list_item_fixed32'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagUintGte, Value: 2},
			FieldValue: func() interface{} { return 1 },
			BeforeFunc: func() { data.TListItemFixed32 = []uint32{1, 5} },
			AfterFunc:  func() { data.TListItemFixed32 = []uint32{2, 10} },
		},
		{
			Name:       fmt.Sprintf("test tag <%s> with fixed32 2", protovalidator.TagRepeatedItem),
			FieldDesc:  "array item where in field 't_list_item_fixed32'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagUintLte, Value: 10},
			FieldValue: func() interface{} { return 12 },
			BeforeFunc: func() { data.TListItemFixed32 = []uint32{2, 12} },
			AfterFunc:  func() { data.TListItemFixed32 = []uint32{2, 10} },
		},

		{
			Name:       fmt.Sprintf("test tag <%s> with fixed64 1", protovalidator.TagRepeatedItem),
			FieldDesc:  "array item where in field 't_list_item_fixed64'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagUintGte, Value: 2},
			FieldValue: func() interface{} { return 1 },
			BeforeFunc: func() { data.TListItemFixed64 = []uint64{1, 5} },
			AfterFunc:  func() { data.TListItemFixed64 = []uint64{2, 10} },
		},
		{
			Name:       fmt.Sprintf("test tag <%s> with fixed64 2", protovalidator.TagRepeatedItem),
			FieldDesc:  "array item where in field 't_list_item_fixed64'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagUintLte, Value: 10},
			FieldValue: func() interface{} { return 12 },
			BeforeFunc: func() { data.TListItemFixed64 = []uint64{2, 12} },
			AfterFunc:  func() { data.TListItemFixed64 = []uint64{2, 10} },
		},

		{
			Name:       fmt.Sprintf("test tag <%s> with bool 2", protovalidator.TagRepeatedItem),
			FieldDesc:  "array item where in field 't_list_item_bool'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagBoolEq, Value: true},
			FieldValue: func() interface{} { return false },
			BeforeFunc: func() { data.TListItemBool = []bool{false, true} },
			AfterFunc:  func() { data.TListItemBool = []bool{true, true} },
		},

		{
			Name:       fmt.Sprintf("test tag <%s> with enum 1", protovalidator.TagRepeatedItem),
			FieldDesc:  "array item where in field 't_list_item_enum'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagEnumInEnums, Value: []int32{0, 1, 2, 3, 4, 8}},
			FieldValue: func() interface{} { return 11 },
			BeforeFunc: func() {
				data.TListItemEnum = []govalidatortest.Enum1{govalidatortest.Enum1_February, govalidatortest.Enum1(11)}
			},
			AfterFunc: func() {
				data.TListItemEnum = []govalidatortest.Enum1{govalidatortest.Enum1_February, govalidatortest.Enum1_June}
			},
		},

		{
			Name:       fmt.Sprintf("test tag <%s> with bytes 1", protovalidator.TagRepeatedItem),
			FieldDesc:  "array item where in field 't_list_item_bytes'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagBytesLenEq, Value: 3},
			FieldValue: func() interface{} { return len(data.TListItemBytes) },
			BeforeFunc: func() { data.TListItemBytes = [][]byte{[]byte("s1"), []byte("s2")} },
			AfterFunc:  func() { data.TListItemBytes = [][]byte{[]byte("ss1"), []byte("ss2")} },
		},

		{
			Name:       fmt.Sprintf("test tag <%s> with message 1", protovalidator.TagRepeatedItem),
			FieldDesc:  "array item where in field 't_list_item_message'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagMessageNotNull, Value: nil},
			FieldValue: func() interface{} { return nil }, // FIXME
			BeforeFunc: func() { data.TListItemMessage = []*govalidatortest.Config{nil, nil} },
			AfterFunc: func() {
				data.TListItemMessage = []*govalidatortest.Config{{Ip: "127.0.0.1", Port: 8080}, {Ip: "127.0.0.1", Port: 8080}}
			},
			UseError2: true,
		},
	}

	msgName := "ValidRepeatedTagsItem1"
	runCases(t, data, msgName, cases)

	{
		err := data.Validate()
		require.Nil(t, err)
	}
}

func Test_GoValidator_ValidMapTagsGeneral1(t *testing.T) {
	data := &govalidatortest.ValidMapTagsGeneral1{}
	{
		err := data.Validate()
		require.NotNil(t, err)
	}

	cases := []*CaseDesc{
		{
			Name:       fmt.Sprintf("test tag <%s>", protovalidator.TagMapNotNull),
			FieldDesc:  "field 't_map_not_null1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagMapNotNull, Value: nil},
			FieldValue: func() interface{} { return data.TMapNotNull1 },
			BeforeFunc: func() {},
			AfterFunc:  func() { data.TMapNotNull1 = map[string]string{"k1": "v1"} },
			UseError2:  true,
		},
		{
			Name:       fmt.Sprintf("test tag <%s>", protovalidator.TagMapLenEq),
			FieldDesc:  "field 't_map_len_eq1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagMapLenEq, Value: 1},
			FieldValue: func() interface{} { return len(data.TMapLenEq1) },
			BeforeFunc: func() {},
			AfterFunc:  func() { data.TMapLenEq1 = map[string]string{"k1": "v1"} },
		},
		{
			Name:       fmt.Sprintf("test tag <%s>", protovalidator.TagMapLenNe),
			FieldDesc:  "field 't_map_len_ne1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagMapLenNe, Value: 2},
			FieldValue: func() interface{} { return len(data.TMapLenNe1) },
			BeforeFunc: func() { data.TMapLenNe1 = map[string]string{"k1": "v1", "k2": "v2"} },
			AfterFunc:  func() { data.TMapLenNe1 = map[string]string{"k1": "v1"} },
		},
		{
			Name:       fmt.Sprintf("test tag <%s>", protovalidator.TagMapLenLt),
			FieldDesc:  "field 't_map_len_lt1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagMapLenLt, Value: 3},
			FieldValue: func() interface{} { return len(data.TMapLenLt1) },
			BeforeFunc: func() { data.TMapLenLt1 = map[string]string{"k1": "v1", "k2": "v2", "k3": "v3", "k4": "v4"} },
			AfterFunc:  func() { data.TMapLenLt1 = map[string]string{"k1": "v1"} },
		},
		{
			Name:       fmt.Sprintf("test tag <%s>", protovalidator.TagMapLenGt),
			FieldDesc:  "field 't_map_len_gt1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagMapLenGt, Value: 4},
			FieldValue: func() interface{} { return len(data.TMapLenGt1) },
			BeforeFunc: func() { data.TMapLenGt1 = map[string]string{"k1": "v1"} },
			AfterFunc: func() {
				data.TMapLenGt1 = map[string]string{"k1": "v1", "k2": "v2", "k3": "v3", "k4": "v4", "k5": "v5"}
			},
		},
		{
			Name:       fmt.Sprintf("test tag <%s>", protovalidator.TagMapLenLte),
			FieldDesc:  "field 't_map_len_lte1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagMapLenLte, Value: 5},
			FieldValue: func() interface{} { return len(data.TMapLenLte1) },
			BeforeFunc: func() {
				data.TMapLenLte1 = map[string]string{"k1": "v1", "k2": "v2", "k3": "v3", "k4": "v4", "k5": "v5", "k6": "v6"}
			},
			AfterFunc: func() {
				data.TMapLenLte1 = map[string]string{"k1": "v1", "k2": "v2", "k3": "v3", "k4": "v4", "k5": "v5"}
			},
		},
		{
			Name:       fmt.Sprintf("test tag <%s>", protovalidator.TagMapLenGte),
			FieldDesc:  "field 't_map_len_gte1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagMapLenGte, Value: 6},
			FieldValue: func() interface{} { return len(data.TMapLenGte1) },
			BeforeFunc: func() {},
			AfterFunc: func() {
				data.TMapLenGte1 = map[string]string{"k1": "v1", "k2": "v2", "k3": "v3", "k4": "v4", "k5": "v5", "k6": "v6"}
			},
		},
	}

	msgName := "ValidMapTagsGeneral1"
	runCases(t, data, msgName, cases)

	{
		err := data.Validate()
		require.Nil(t, err)
	}
}

func Test_GoValidator_ValidMapTagsKey1(t *testing.T) {
	data := &govalidatortest.ValidMapTagsKey1{}
	{
		err := data.Validate()
		require.NotNil(t, err)
	}

	cases := []*CaseDesc{
		{
			Name:       fmt.Sprintf("test tag <%s> with key string 1", protovalidator.TagMapKey),
			FieldDesc:  "map key where in field 't_map_key_string'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagStringPrefix, Value: "id-"},
			FieldValue: func() interface{} { return "ik1" },
			BeforeFunc: func() { data.TMapKeyString = map[string]int32{"ik1": 1, "id-k2": 1} },
			AfterFunc:  func() { data.TMapKeyString = map[string]int32{"id-k1": 1, "id-k2": 1} },
		},

		{
			Name:       fmt.Sprintf("test tag <%s> with key int32 1", protovalidator.TagMapKey),
			FieldDesc:  "map key where in field 't_map_key_int32'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagIntGte, Value: 2},
			FieldValue: func() interface{} { return 1 },
			BeforeFunc: func() { data.TMapKeyInt32 = map[int32]int32{1: 1, 5: 1} },
			AfterFunc:  func() { data.TMapKeyInt32 = map[int32]int32{2: 1, 10: 1} },
		},
		{
			Name:       fmt.Sprintf("test tag <%s> with int32 2", protovalidator.TagMapKey),
			FieldDesc:  "map key where in field 't_map_key_int32'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagIntLte, Value: 10},
			FieldValue: func() interface{} { return 12 },
			BeforeFunc: func() { data.TMapKeyInt32 = map[int32]int32{2: 1, 12: 1} },
			AfterFunc:  func() { data.TMapKeyInt32 = map[int32]int32{2: 1, 10: 1} },
		},

		{
			Name:       fmt.Sprintf("test tag <%s> with key int64 1", protovalidator.TagMapKey),
			FieldDesc:  "map key where in field 't_map_key_int64'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagIntGte, Value: 2},
			FieldValue: func() interface{} { return 1 },
			BeforeFunc: func() { data.TMapKeyInt64 = map[int64]int32{1: 1, 5: 1} },
			AfterFunc:  func() { data.TMapKeyInt64 = map[int64]int32{2: 1, 10: 1} },
		},
		{
			Name:       fmt.Sprintf("test tag <%s> with int64 2", protovalidator.TagMapKey),
			FieldDesc:  "map key where in field 't_map_key_int64'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagIntLte, Value: 10},
			FieldValue: func() interface{} { return 12 },
			BeforeFunc: func() { data.TMapKeyInt64 = map[int64]int32{2: 1, 12: 1} },
			AfterFunc:  func() { data.TMapKeyInt64 = map[int64]int32{2: 1, 10: 1} },
		},

		{
			Name:       fmt.Sprintf("test tag <%s> with key sint32 1", protovalidator.TagMapKey),
			FieldDesc:  "map key where in field 't_map_key_sint32'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagIntGte, Value: 2},
			FieldValue: func() interface{} { return 1 },
			BeforeFunc: func() { data.TMapKeySint32 = map[int32]int32{1: 1, 5: 1} },
			AfterFunc:  func() { data.TMapKeySint32 = map[int32]int32{2: 1, 10: 1} },
		},
		{
			Name:       fmt.Sprintf("test tag <%s> with sint32 2", protovalidator.TagMapKey),
			FieldDesc:  "map key where in field 't_map_key_sint32'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagIntLte, Value: 10},
			FieldValue: func() interface{} { return 12 },
			BeforeFunc: func() { data.TMapKeySint32 = map[int32]int32{2: 1, 12: 1} },
			AfterFunc:  func() { data.TMapKeySint32 = map[int32]int32{2: 1, 10: 1} },
		},

		{
			Name:       fmt.Sprintf("test tag <%s> with key sint64 1", protovalidator.TagMapKey),
			FieldDesc:  "map key where in field 't_map_key_sint64'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagIntGte, Value: 2},
			FieldValue: func() interface{} { return 1 },
			BeforeFunc: func() { data.TMapKeySint64 = map[int64]int32{1: 1, 5: 1} },
			AfterFunc:  func() { data.TMapKeySint64 = map[int64]int32{2: 1, 10: 1} },
		},
		{
			Name:       fmt.Sprintf("test tag <%s> with sint64 2", protovalidator.TagMapKey),
			FieldDesc:  "map key where in field 't_map_key_sint64'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagIntLte, Value: 10},
			FieldValue: func() interface{} { return 12 },
			BeforeFunc: func() { data.TMapKeySint64 = map[int64]int32{2: 1, 12: 1} },
			AfterFunc:  func() { data.TMapKeySint64 = map[int64]int32{2: 1, 10: 1} },
		},

		{
			Name:       fmt.Sprintf("test tag <%s> with key sfixed32 1", protovalidator.TagMapKey),
			FieldDesc:  "map key where in field 't_map_key_sfixed32'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagIntGte, Value: 2},
			FieldValue: func() interface{} { return 1 },
			BeforeFunc: func() { data.TMapKeySfixed32 = map[int32]int32{1: 1, 5: 1} },
			AfterFunc:  func() { data.TMapKeySfixed32 = map[int32]int32{2: 1, 10: 1} },
		},
		{
			Name:       fmt.Sprintf("test tag <%s> with sfixed32 2", protovalidator.TagMapKey),
			FieldDesc:  "map key where in field 't_map_key_sfixed32'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagIntLte, Value: 10},
			FieldValue: func() interface{} { return 12 },
			BeforeFunc: func() { data.TMapKeySfixed32 = map[int32]int32{2: 1, 12: 1} },
			AfterFunc:  func() { data.TMapKeySfixed32 = map[int32]int32{2: 1, 10: 1} },
		},

		{
			Name:       fmt.Sprintf("test tag <%s> with key sfixed64 1", protovalidator.TagMapKey),
			FieldDesc:  "map key where in field 't_map_key_sfixed64'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagIntGte, Value: 2},
			FieldValue: func() interface{} { return 1 },
			BeforeFunc: func() { data.TMapKeySfixed64 = map[int64]int32{1: 1, 5: 1} },
			AfterFunc:  func() { data.TMapKeySfixed64 = map[int64]int32{2: 1, 10: 1} },
		},
		{
			Name:       fmt.Sprintf("test tag <%s> with sfixed64 2", protovalidator.TagMapKey),
			FieldDesc:  "map key where in field 't_map_key_sfixed64'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagIntLte, Value: 10},
			FieldValue: func() interface{} { return 12 },
			BeforeFunc: func() { data.TMapKeySfixed64 = map[int64]int32{2: 1, 12: 1} },
			AfterFunc:  func() { data.TMapKeySfixed64 = map[int64]int32{2: 1, 10: 1} },
		},

		{
			Name:       fmt.Sprintf("test tag <%s> with key uint32 1", protovalidator.TagMapKey),
			FieldDesc:  "map key where in field 't_map_key_uint32'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagUintGte, Value: 2},
			FieldValue: func() interface{} { return 1 },
			BeforeFunc: func() { data.TMapKeyUint32 = map[uint32]int32{1: 1, 5: 1} },
			AfterFunc:  func() { data.TMapKeyUint32 = map[uint32]int32{2: 1, 10: 1} },
		},
		{
			Name:       fmt.Sprintf("test tag <%s> with uint32 2", protovalidator.TagMapKey),
			FieldDesc:  "map key where in field 't_map_key_uint32'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagUintLte, Value: 10},
			FieldValue: func() interface{} { return 12 },
			BeforeFunc: func() { data.TMapKeyUint32 = map[uint32]int32{2: 1, 12: 1} },
			AfterFunc:  func() { data.TMapKeyUint32 = map[uint32]int32{2: 1, 10: 1} },
		},

		{
			Name:       fmt.Sprintf("test tag <%s> with key uint64 1", protovalidator.TagMapKey),
			FieldDesc:  "map key where in field 't_map_key_uint64'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagUintGte, Value: 2},
			FieldValue: func() interface{} { return 1 },
			BeforeFunc: func() { data.TMapKeyUint64 = map[uint64]int32{1: 1, 5: 1} },
			AfterFunc:  func() { data.TMapKeyUint64 = map[uint64]int32{2: 1, 10: 1} },
		},
		{
			Name:       fmt.Sprintf("test tag <%s> with uint64 2", protovalidator.TagMapKey),
			FieldDesc:  "map key where in field 't_map_key_uint64'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagUintLte, Value: 10},
			FieldValue: func() interface{} { return 12 },
			BeforeFunc: func() { data.TMapKeyUint64 = map[uint64]int32{2: 1, 12: 1} },
			AfterFunc:  func() { data.TMapKeyUint64 = map[uint64]int32{2: 1, 10: 1} },
		},

		{
			Name:       fmt.Sprintf("test tag <%s> with key fixed32 1", protovalidator.TagMapKey),
			FieldDesc:  "map key where in field 't_map_key_fixed32'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagUintGte, Value: 2},
			FieldValue: func() interface{} { return 1 },
			BeforeFunc: func() { data.TMapKeyFixed32 = map[uint32]int32{1: 1, 5: 1} },
			AfterFunc:  func() { data.TMapKeyFixed32 = map[uint32]int32{2: 1, 10: 1} },
		},
		{
			Name:       fmt.Sprintf("test tag <%s> with fixed32 2", protovalidator.TagMapKey),
			FieldDesc:  "map key where in field 't_map_key_fixed32'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagUintLte, Value: 10},
			FieldValue: func() interface{} { return 12 },
			BeforeFunc: func() { data.TMapKeyFixed32 = map[uint32]int32{2: 1, 12: 1} },
			AfterFunc:  func() { data.TMapKeyFixed32 = map[uint32]int32{2: 1, 10: 1} },
		},

		{
			Name:       fmt.Sprintf("test tag <%s> with key fixed64 1", protovalidator.TagMapKey),
			FieldDesc:  "map key where in field 't_map_key_fixed64'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagUintGte, Value: 2},
			FieldValue: func() interface{} { return 1 },
			BeforeFunc: func() { data.TMapKeyFixed64 = map[uint64]int32{1: 1, 5: 1} },
			AfterFunc:  func() { data.TMapKeyFixed64 = map[uint64]int32{2: 1, 10: 1} },
		},
		{
			Name:       fmt.Sprintf("test tag <%s> with fixed64 2", protovalidator.TagMapKey),
			FieldDesc:  "map key where in field 't_map_key_fixed64'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagUintLte, Value: 10},
			FieldValue: func() interface{} { return 12 },
			BeforeFunc: func() { data.TMapKeyFixed64 = map[uint64]int32{2: 1, 12: 1} },
			AfterFunc:  func() { data.TMapKeyFixed64 = map[uint64]int32{2: 1, 10: 1} },
		},
	}

	msgName := "ValidMapTagsKey1"
	runCases(t, data, msgName, cases)

	{
		err := data.Validate()
		require.Nil(t, err)
	}
}

func Test_GoValidator_ValidMapTagsValue1(t *testing.T) {
	data := &govalidatortest.ValidMapTagsValue1{}
	{
		err := data.Validate()
		require.NotNil(t, err)
	}

	cases := []*CaseDesc{
		{
			Name:       fmt.Sprintf("test tag <%s> with value string 1", protovalidator.TagMapValue),
			FieldDesc:  "map value where in field 't_map_value_string'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagStringPrefix, Value: "id-"},
			FieldValue: func() interface{} { return "v1" },
			BeforeFunc: func() { data.TMapValueString = map[string]string{"k1": "v1", "k2": "id-v2"} },
			AfterFunc:  func() { data.TMapValueString = map[string]string{"k1": "id-v1", "k2": "id-v2"} },
		},

		{
			Name:       fmt.Sprintf("test tag <%s> with value double 1", protovalidator.TagMapValue),
			FieldDesc:  "map value where in field 't_map_value_double'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagFloatGte, Value: 1.1},
			FieldValue: func() interface{} { return float64(0.5) },
			BeforeFunc: func() { data.TMapValueDouble = map[string]float64{"k1": 0.5, "k2": 5.5} },
			AfterFunc:  func() { data.TMapValueDouble = map[string]float64{"k1": 1.1, "k2": 10} },
		},
		{
			Name:       fmt.Sprintf("test tag <%s> with value double 2", protovalidator.TagMapValue),
			FieldDesc:  "map value where in field 't_map_value_double'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagFloatLte, Value: 11.1},
			FieldValue: func() interface{} { return float64(12) },
			BeforeFunc: func() { data.TMapValueDouble = map[string]float64{"k1": 1.1, "k2": 12} },
			AfterFunc:  func() { data.TMapValueDouble = map[string]float64{"k1": 1.1, "k2": 10} },
		},

		{
			Name:       fmt.Sprintf("test tag <%s> with value float 1", protovalidator.TagMapValue),
			FieldDesc:  "map value where in field 't_map_value_float'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagFloatGte, Value: 1.1},
			FieldValue: func() interface{} { return float32(0.8) },
			BeforeFunc: func() { data.TMapValueFloat = map[string]float32{"k1": 0.8, "k2": 5.5} },
			AfterFunc:  func() { data.TMapValueFloat = map[string]float32{"k1": 1.1, "k2": 10} },
		},
		{
			Name:       fmt.Sprintf("test tag <%s> with value float 2", protovalidator.TagMapValue),
			FieldDesc:  "map value where in field 't_map_value_float'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagFloatLte, Value: 11.1},
			FieldValue: func() interface{} { return float32(12) },
			BeforeFunc: func() { data.TMapValueFloat = map[string]float32{"k1": 1.1, "k2": 12} },
			AfterFunc:  func() { data.TMapValueFloat = map[string]float32{"k1": 1.1, "k2": 10} },
		},

		{
			Name:       fmt.Sprintf("test tag <%s> with value int32 1", protovalidator.TagMapValue),
			FieldDesc:  "map value where in field 't_map_value_int32'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagIntGte, Value: 2},
			FieldValue: func() interface{} { return 1 },
			BeforeFunc: func() { data.TMapValueInt32 = map[string]int32{"k1": 1, "k2": 10} },
			AfterFunc:  func() { data.TMapValueInt32 = map[string]int32{"k1": 2, "k2": 10} },
		},
		{
			Name:       fmt.Sprintf("test tag <%s> with value int32 2", protovalidator.TagMapValue),
			FieldDesc:  "map value where in field 't_map_value_int32'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagIntLte, Value: 10},
			FieldValue: func() interface{} { return 12 },
			BeforeFunc: func() { data.TMapValueInt32 = map[string]int32{"k1": 2, "k2": 12} },
			AfterFunc:  func() { data.TMapValueInt32 = map[string]int32{"k1": 2, "k2": 10} },
		},

		{
			Name:       fmt.Sprintf("test tag <%s> with value int64 1", protovalidator.TagMapValue),
			FieldDesc:  "map value where in field 't_map_value_int64'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagIntGte, Value: 2},
			FieldValue: func() interface{} { return 1 },
			BeforeFunc: func() { data.TMapValueInt64 = map[string]int64{"k1": 1, "k2": 10} },
			AfterFunc:  func() { data.TMapValueInt64 = map[string]int64{"k1": 2, "k2": 10} },
		},
		{
			Name:       fmt.Sprintf("test tag <%s> with value int64 2", protovalidator.TagMapValue),
			FieldDesc:  "map value where in field 't_map_value_int64'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagIntLte, Value: 10},
			FieldValue: func() interface{} { return 12 },
			BeforeFunc: func() { data.TMapValueInt64 = map[string]int64{"k1": 2, "k2": 12} },
			AfterFunc:  func() { data.TMapValueInt64 = map[string]int64{"k1": 2, "k2": 10} },
		},

		{
			Name:       fmt.Sprintf("test tag <%s> with value sint32 1", protovalidator.TagMapValue),
			FieldDesc:  "map value where in field 't_map_value_sint32'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagIntGte, Value: 2},
			FieldValue: func() interface{} { return 1 },
			BeforeFunc: func() { data.TMapValueSint32 = map[string]int32{"k1": 1, "k2": 10} },
			AfterFunc:  func() { data.TMapValueSint32 = map[string]int32{"k1": 2, "k2": 10} },
		},
		{
			Name:       fmt.Sprintf("test tag <%s> with value sint32 2", protovalidator.TagMapValue),
			FieldDesc:  "map value where in field 't_map_value_sint32'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagIntLte, Value: 10},
			FieldValue: func() interface{} { return 12 },
			BeforeFunc: func() { data.TMapValueSint32 = map[string]int32{"k1": 2, "k2": 12} },
			AfterFunc:  func() { data.TMapValueSint32 = map[string]int32{"k1": 2, "k2": 10} },
		},

		{
			Name:       fmt.Sprintf("test tag <%s> with value sint64 1", protovalidator.TagMapValue),
			FieldDesc:  "map value where in field 't_map_value_sint64'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagIntGte, Value: 2},
			FieldValue: func() interface{} { return 1 },
			BeforeFunc: func() { data.TMapValueSint64 = map[string]int64{"k1": 1, "k2": 10} },
			AfterFunc:  func() { data.TMapValueSint64 = map[string]int64{"k1": 2, "k2": 10} },
		},
		{
			Name:       fmt.Sprintf("test tag <%s> with value sint64 2", protovalidator.TagMapValue),
			FieldDesc:  "map value where in field 't_map_value_sint64'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagIntLte, Value: 10},
			FieldValue: func() interface{} { return 12 },
			BeforeFunc: func() { data.TMapValueSint64 = map[string]int64{"k1": 2, "k2": 12} },
			AfterFunc:  func() { data.TMapValueSint64 = map[string]int64{"k1": 2, "k2": 10} },
		},

		{
			Name:       fmt.Sprintf("test tag <%s> with value sfixed32 1", protovalidator.TagMapValue),
			FieldDesc:  "map value where in field 't_map_value_sfixed32'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagIntGte, Value: 2},
			FieldValue: func() interface{} { return 1 },
			BeforeFunc: func() { data.TMapValueSfixed32 = map[string]int32{"k1": 1, "k2": 10} },
			AfterFunc:  func() { data.TMapValueSfixed32 = map[string]int32{"k1": 2, "k2": 10} },
		},
		{
			Name:       fmt.Sprintf("test tag <%s> with value sfixed32 2", protovalidator.TagMapValue),
			FieldDesc:  "map value where in field 't_map_value_sfixed32'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagIntLte, Value: 10},
			FieldValue: func() interface{} { return 12 },
			BeforeFunc: func() { data.TMapValueSfixed32 = map[string]int32{"k1": 2, "k2": 12} },
			AfterFunc:  func() { data.TMapValueSfixed32 = map[string]int32{"k1": 2, "k2": 10} },
		},

		{
			Name:       fmt.Sprintf("test tag <%s> with value sfixed64 1", protovalidator.TagMapValue),
			FieldDesc:  "map value where in field 't_map_value_sfixed64'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagIntGte, Value: 2},
			FieldValue: func() interface{} { return 1 },
			BeforeFunc: func() { data.TMapValueSfixed64 = map[string]int64{"k1": 1, "k2": 10} },
			AfterFunc:  func() { data.TMapValueSfixed64 = map[string]int64{"k1": 2, "k2": 10} },
		},
		{
			Name:       fmt.Sprintf("test tag <%s> with value sfixed64 2", protovalidator.TagMapValue),
			FieldDesc:  "map value where in field 't_map_value_sfixed64'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagIntLte, Value: 10},
			FieldValue: func() interface{} { return 12 },
			BeforeFunc: func() { data.TMapValueSfixed64 = map[string]int64{"k1": 2, "k2": 12} },
			AfterFunc:  func() { data.TMapValueSfixed64 = map[string]int64{"k1": 2, "k2": 10} },
		},

		{
			Name:       fmt.Sprintf("test tag <%s> with value uint32 1", protovalidator.TagMapValue),
			FieldDesc:  "map value where in field 't_map_value_uint32'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagUintGte, Value: 2},
			FieldValue: func() interface{} { return 1 },
			BeforeFunc: func() { data.TMapValueUint32 = map[string]uint32{"k1": 1, "k2": 10} },
			AfterFunc:  func() { data.TMapValueUint32 = map[string]uint32{"k1": 2, "k2": 10} },
		},
		{
			Name:       fmt.Sprintf("test tag <%s> with value uint32 2", protovalidator.TagMapValue),
			FieldDesc:  "map value where in field 't_map_value_uint32'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagUintLte, Value: 10},
			FieldValue: func() interface{} { return 12 },
			BeforeFunc: func() { data.TMapValueUint32 = map[string]uint32{"k1": 2, "k2": 12} },
			AfterFunc:  func() { data.TMapValueUint32 = map[string]uint32{"k1": 2, "k2": 10} },
		},

		{
			Name:       fmt.Sprintf("test tag <%s> with value uint64 1", protovalidator.TagMapValue),
			FieldDesc:  "map value where in field 't_map_value_uint64'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagUintGte, Value: 2},
			FieldValue: func() interface{} { return 1 },
			BeforeFunc: func() { data.TMapValueUint64 = map[string]uint64{"k1": 1, "k2": 10} },
			AfterFunc:  func() { data.TMapValueUint64 = map[string]uint64{"k1": 2, "k2": 10} },
		},
		{
			Name:       fmt.Sprintf("test tag <%s> with value uint64 2", protovalidator.TagMapValue),
			FieldDesc:  "map value where in field 't_map_value_uint64'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagUintLte, Value: 10},
			FieldValue: func() interface{} { return 12 },
			BeforeFunc: func() { data.TMapValueUint64 = map[string]uint64{"k1": 2, "k2": 12} },
			AfterFunc:  func() { data.TMapValueUint64 = map[string]uint64{"k1": 2, "k2": 10} },
		},

		{
			Name:       fmt.Sprintf("test tag <%s> with value fixed32 1", protovalidator.TagMapValue),
			FieldDesc:  "map value where in field 't_map_value_fixed32'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagUintGte, Value: 2},
			FieldValue: func() interface{} { return 1 },
			BeforeFunc: func() { data.TMapValueFixed32 = map[string]uint32{"k1": 1, "k2": 5} },
			AfterFunc:  func() { data.TMapValueFixed32 = map[string]uint32{"k1": 2, "k2": 10} },
		},
		{
			Name:       fmt.Sprintf("test tag <%s> with value fixed32 2", protovalidator.TagMapValue),
			FieldDesc:  "map value where in field 't_map_value_fixed32'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagUintLte, Value: 10},
			FieldValue: func() interface{} { return 12 },
			BeforeFunc: func() { data.TMapValueFixed32 = map[string]uint32{"k1": 2, "k2": 12} },
			AfterFunc:  func() { data.TMapValueFixed32 = map[string]uint32{"k1": 2, "k2": 10} },
		},

		{
			Name:       fmt.Sprintf("test tag <%s> with value fixed64 1", protovalidator.TagMapValue),
			FieldDesc:  "map value where in field 't_map_value_fixed64'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagUintGte, Value: 2},
			FieldValue: func() interface{} { return 1 },
			BeforeFunc: func() { data.TMapValueFixed64 = map[string]uint64{"k1": 1, "k2": 10} },
			AfterFunc:  func() { data.TMapValueFixed64 = map[string]uint64{"k1": 2, "k2": 10} },
		},
		{
			Name:       fmt.Sprintf("test tag <%s> with value fixed64 2", protovalidator.TagMapValue),
			FieldDesc:  "map value where in field 't_map_value_fixed64'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagUintLte, Value: 10},
			FieldValue: func() interface{} { return 12 },
			BeforeFunc: func() { data.TMapValueFixed64 = map[string]uint64{"k1": 2, "k2": 12} },
			AfterFunc:  func() { data.TMapValueFixed64 = map[string]uint64{"k1": 2, "k2": 10} },
		},

		{
			Name:       fmt.Sprintf("test tag <%s> with value bool 1", protovalidator.TagMapValue),
			FieldDesc:  "map value where in field 't_map_value_bool'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagBoolEq, Value: true},
			FieldValue: func() interface{} { return false },
			BeforeFunc: func() { data.TMapValueBool = map[string]bool{"k1": false, "k2": true} },
			AfterFunc:  func() { data.TMapValueBool = map[string]bool{"k1": true, "k2": true} },
		},

		{
			Name:       fmt.Sprintf("test tag <%s> with value enum 1", protovalidator.TagMapValue),
			FieldDesc:  "map value where in field 't_map_value_enum'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagEnumInEnums, Value: []int32{0, 1, 2, 3, 4, 8}},
			FieldValue: func() interface{} { return 11 },
			BeforeFunc: func() {
				data.TMapValueEnum = map[string]govalidatortest.Enum1{"k1": govalidatortest.Enum1_February, "k2": govalidatortest.Enum1(11)}
			},
			AfterFunc: func() {
				data.TMapValueEnum = map[string]govalidatortest.Enum1{"k1": govalidatortest.Enum1_February, "k2": govalidatortest.Enum1_June}
			},
		},

		{
			Name:       fmt.Sprintf("test tag <%s> with value bytes 1", protovalidator.TagMapValue),
			FieldDesc:  "map value where in field 't_map_value_bytes'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagBytesLenEq, Value: 3},
			FieldValue: func() interface{} { return len(data.TMapValueBytes) },
			BeforeFunc: func() { data.TMapValueBytes = map[string][]byte{"k1": []byte("s1"), "k2": []byte("s2")} },
			AfterFunc:  func() { data.TMapValueBytes = map[string][]byte{"k1": []byte("ss1"), "k2": []byte("ss2")} },
		},

		{
			Name:       fmt.Sprintf("test tag <%s> with value message 1", protovalidator.TagMapValue),
			FieldDesc:  "map value where in field 't_map_value_message'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagMessageNotNull, Value: nil},
			FieldValue: func() interface{} { return nil }, // FIXME
			BeforeFunc: func() { data.TMapValueMessage = map[string]*govalidatortest.Config{"k1": nil, "k2": nil} },
			AfterFunc: func() {
				data.TMapValueMessage = map[string]*govalidatortest.Config{"k1": {Ip: "127.0.0.1", Port: 8080}, "k2": {Ip: "127.0.0.1", Port: 8080}}
			},
			UseError2: true,
		},
	}

	msgName := "ValidMapTagsValue1"
	runCases(t, data, msgName, cases)

	{
		err := data.Validate()
		require.Nil(t, err)
	}
}

func Test_GoValidator_ValidStringTagsGeneral1(t *testing.T) {
	data := &govalidatortest.ValidStringTagsGeneral1{}
	{
		err := data.Validate()
		require.NotNil(t, err)
	}

	caseGeneral1 := []*CaseDesc{
		{
			Name:       fmt.Sprintf("test tag <%s> 1", protovalidator.TagStringEq),
			FieldDesc:  "field 't_string_eq1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagStringEq, Value: "b"},
			FieldValue: func() interface{} { return data.TStringEq1 },
			BeforeFunc: func() { data.TStringEq1 = "" },
			AfterFunc:  func() { data.TStringEq1 = "b" },
		},
		{
			Name:       fmt.Sprintf("test tag <%s> 1", protovalidator.TagStringNe),
			FieldDesc:  "field 't_string_ne1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagStringNe, Value: "b"},
			FieldValue: func() interface{} { return data.TStringNe1 },
			BeforeFunc: func() { data.TStringNe1 = "b" },
			AfterFunc:  func() { data.TStringNe1 = "a" },
		},
		{
			Name:       fmt.Sprintf("test tag <%s> 1", protovalidator.TagStringLt),
			FieldDesc:  "field 't_string_lt1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagStringLt, Value: "c"},
			FieldValue: func() interface{} { return data.TStringLt1 },
			BeforeFunc: func() { data.TStringLt1 = "c" },
			AfterFunc:  func() { data.TStringLt1 = "b" },
		},
		{
			Name:       fmt.Sprintf("test tag <%s> 1", protovalidator.TagStringGt),
			FieldDesc:  "field 't_string_gt1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagStringGt, Value: "c"},
			FieldValue: func() interface{} { return data.TStringGt1 },
			BeforeFunc: func() { data.TStringGt1 = "c" },
			AfterFunc:  func() { data.TStringGt1 = "d" },
		},
		{
			Name:       fmt.Sprintf("test tag <%s> 1", protovalidator.TagStringLte),
			FieldDesc:  "field 't_string_lte1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagStringLte, Value: "c"},
			FieldValue: func() interface{} { return data.TStringLte1 },
			BeforeFunc: func() { data.TStringLte1 = "d" },
			AfterFunc:  func() { data.TStringLte1 = "c" },
		},
		{
			Name:       fmt.Sprintf("test tag <%s> 1", protovalidator.TagStringGte),
			FieldDesc:  "field 't_string_gte1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagStringGte, Value: "c"},
			FieldValue: func() interface{} { return data.TStringGte1 },
			BeforeFunc: func() { data.TStringGte1 = "b" },
			AfterFunc:  func() { data.TStringGte1 = "c" },
		},
		{
			Name:       fmt.Sprintf("test tag <%s> 1", protovalidator.TagStringIn),
			FieldDesc:  "field 't_string_in1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagStringIn, Value: []string{"a", "b", "c"}},
			FieldValue: func() interface{} { return data.TStringIn1 },
			BeforeFunc: func() { data.TStringIn1 = "d" },
			AfterFunc:  func() { data.TStringIn1 = "b" },
		},
		{
			Name:       fmt.Sprintf("test tag <%s> 1", protovalidator.TagStringNotIn),
			FieldDesc:  "field 't_string_not_in1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagStringNotIn, Value: []string{"x", "y", "z"}},
			FieldValue: func() interface{} { return data.TStringNotIn1 },
			BeforeFunc: func() { data.TStringNotIn1 = "x" },
			AfterFunc:  func() { data.TStringNotIn1 = "v" },
		},
		{
			Name:       fmt.Sprintf("test tag <%s> 1", protovalidator.TagStringCharLenEq),
			FieldDesc:  "field 't_string_char_len_eq1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagStringCharLenEq, Value: 3},
			FieldValue: func() interface{} { return utf8.RuneCountInString(data.TStringCharLenEq1) },
			BeforeFunc: func() { data.TStringCharLenEq1 = "x" },
			AfterFunc:  func() { data.TStringCharLenEq1 = "a中b" },
		},
		{
			Name:       fmt.Sprintf("test tag <%s> 1", protovalidator.TagStringCharLenNe),
			FieldDesc:  "field 't_string_char_len_ne1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagStringCharLenNe, Value: 3},
			FieldValue: func() interface{} { return utf8.RuneCountInString(data.TStringCharLenNe1) },
			BeforeFunc: func() { data.TStringCharLenNe1 = "a中b" },
			AfterFunc:  func() { data.TStringCharLenNe1 = "中" },
		},
		{
			Name:       fmt.Sprintf("test tag <%s> 1", protovalidator.TagStringCharLenGt),
			FieldDesc:  "field 't_string_char_len_gt1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagStringCharLenGt, Value: 3},
			FieldValue: func() interface{} { return utf8.RuneCountInString(data.TStringCharLenGt1) },
			BeforeFunc: func() { data.TStringCharLenGt1 = "a中b" },
			AfterFunc:  func() { data.TStringCharLenGt1 = "a中b文" },
		},
		{
			Name:       fmt.Sprintf("test tag <%s> 1", protovalidator.TagStringCharLenLt),
			FieldDesc:  "field 't_string_char_len_lt1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagStringCharLenLt, Value: 3},
			FieldValue: func() interface{} { return utf8.RuneCountInString(data.TStringCharLenLt1) },
			BeforeFunc: func() { data.TStringCharLenLt1 = "a中b" },
			AfterFunc:  func() { data.TStringCharLenLt1 = "中" },
		},
		{
			Name:       fmt.Sprintf("test tag <%s> 1", protovalidator.TagStringCharLenGte),
			FieldDesc:  "field 't_string_char_len_gte1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagStringCharLenGte, Value: 3},
			FieldValue: func() interface{} { return utf8.RuneCountInString(data.TStringCharLenGte1) },
			BeforeFunc: func() { data.TStringCharLenGte1 = "a中" },
			AfterFunc:  func() { data.TStringCharLenGte1 = "a中b文" },
		},
		{
			Name:       fmt.Sprintf("test tag <%s> 1", protovalidator.TagStringCharLenLte),
			FieldDesc:  "field 't_string_char_len_lte1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagStringCharLenLte, Value: 3},
			FieldValue: func() interface{} { return utf8.RuneCountInString(data.TStringCharLenLte1) },
			BeforeFunc: func() { data.TStringCharLenLte1 = "a中b文" },
			AfterFunc:  func() { data.TStringCharLenLte1 = "a中b" },
		},

		{
			Name:       fmt.Sprintf("test tag <%s> 1", protovalidator.TagStringByteLenEq),
			FieldDesc:  "field 't_string_byte_len_eq1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagStringByteLenEq, Value: 5},
			FieldValue: func() interface{} { return len(data.TStringByteLenEq1) },
			BeforeFunc: func() { data.TStringByteLenEq1 = "a中b文C" },
			AfterFunc:  func() { data.TStringByteLenEq1 = "a中b" },
		},
		{
			Name:       fmt.Sprintf("test tag <%s> 1", protovalidator.TagStringByteLenNe),
			FieldDesc:  "field 't_string_byte_len_ne1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagStringByteLenNe, Value: 5},
			FieldValue: func() interface{} { return len(data.TStringByteLenNe1) },
			BeforeFunc: func() { data.TStringByteLenNe1 = "a中b" },
			AfterFunc:  func() { data.TStringByteLenNe1 = "a" },
		},
		{
			Name:       fmt.Sprintf("test tag <%s> 1", protovalidator.TagStringByteLenGt),
			FieldDesc:  "field 't_string_byte_len_gt1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagStringByteLenGt, Value: 5},
			FieldValue: func() interface{} { return len(data.TStringByteLenGt1) },
			BeforeFunc: func() { data.TStringByteLenGt1 = "a中b" },
			AfterFunc:  func() { data.TStringByteLenGt1 = "a中bc" },
		},
		{
			Name:       fmt.Sprintf("test tag <%s> 1", protovalidator.TagStringByteLenLt),
			FieldDesc:  "field 't_string_byte_len_lt1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagStringByteLenLt, Value: 5},
			FieldValue: func() interface{} { return len(data.TStringByteLenLt1) },
			BeforeFunc: func() { data.TStringByteLenLt1 = "a中b文" },
			AfterFunc:  func() { data.TStringByteLenLt1 = "a中" },
		},
		{
			Name:       fmt.Sprintf("test tag <%s> 1", protovalidator.TagStringByteLenGte),
			FieldDesc:  "field 't_string_byte_len_gte1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagStringByteLenGte, Value: 5},
			FieldValue: func() interface{} { return len(data.TStringByteLenGte1) },
			BeforeFunc: func() { data.TStringByteLenGte1 = "a中" },
			AfterFunc:  func() { data.TStringByteLenGte1 = "a中b" },
		},
		{
			Name:       fmt.Sprintf("test tag <%s> 1", protovalidator.TagStringByteLenLte),
			FieldDesc:  "field 't_string_byte_len_lte1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagStringByteLenLte, Value: 5},
			FieldValue: func() interface{} { return len(data.TStringByteLenLte1) },
			BeforeFunc: func() { data.TStringByteLenLte1 = "a中b文" },
			AfterFunc:  func() { data.TStringByteLenLte1 = "a中b" },
		},
	}

	caseGeneral2 := []*CaseDesc{
		{
			Name:       fmt.Sprintf("test tag <%s> 1", protovalidator.TagStringRegex),
			FieldDesc:  "field 't_string_regex1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagStringRegex, Value: `^[0-9]+.jar$`},
			FieldValue: func() interface{} { return data.TStringRegex1 },
			BeforeFunc: func() { data.TStringRegex1 = "" },
			AfterFunc:  func() { data.TStringRegex1 = "0001.jar" },
		},

		{
			Name:       fmt.Sprintf("test tag <%s> 1", protovalidator.TagStringPrefix),
			FieldDesc:  "field 't_string_prefix1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagStringPrefix, Value: "prefix"},
			FieldValue: func() interface{} { return data.TStringPrefix1 },
			BeforeFunc: func() { data.TStringPrefix1 = "" },
			AfterFunc:  func() { data.TStringPrefix1 = "prefix-xxx1" },
		},
		{
			Name:       fmt.Sprintf("test tag <%s> 1", protovalidator.TagStringNoPrefix),
			FieldDesc:  "field 't_string_no_prefix1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagStringNoPrefix, Value: "prefix"},
			FieldValue: func() interface{} { return data.TStringNoPrefix1 },
			BeforeFunc: func() { data.TStringNoPrefix1 = "prefix-xxx1" },
			AfterFunc:  func() { data.TStringNoPrefix1 = "xxx1" },
		},

		{
			Name:       fmt.Sprintf("test tag <%s> 1", protovalidator.TagStringSuffix),
			FieldDesc:  "field 't_string_suffix1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagStringSuffix, Value: "suffix"},
			FieldValue: func() interface{} { return data.TStringSuffix1 },
			BeforeFunc: func() { data.TStringSuffix1 = "" },
			AfterFunc:  func() { data.TStringSuffix1 = "xxx-suffix" },
		},
		{
			Name:       fmt.Sprintf("test tag <%s> 1", protovalidator.TagStringNoSuffix),
			FieldDesc:  "field 't_string_no_suffix1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagStringNoSuffix, Value: "suffix"},
			FieldValue: func() interface{} { return data.TStringNoSuffix1 },
			BeforeFunc: func() { data.TStringNoSuffix1 = "xxx-suffix" },
			AfterFunc:  func() { data.TStringNoSuffix1 = "xxx" },
		},

		{
			Name:       fmt.Sprintf("test tag <%s> 1", protovalidator.TagStringContains),
			FieldDesc:  "field 't_string_contains1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagStringContains, Value: "abc"},
			FieldValue: func() interface{} { return data.TStringContains1 },
			BeforeFunc: func() { data.TStringContains1 = "1a2b3c" },
			AfterFunc:  func() { data.TStringContains1 = "1abc23" },
		},
		{
			Name:       fmt.Sprintf("test tag <%s> 1", protovalidator.TagStringNotContains),
			FieldDesc:  "field 't_string_no_contains1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagStringNotContains, Value: "abc"},
			FieldValue: func() interface{} { return data.TStringNoContains1 },
			BeforeFunc: func() { data.TStringNoContains1 = "1abc23" },
			AfterFunc:  func() { data.TStringNoContains1 = "1a2b3c" },
		},

		{
			Name:       fmt.Sprintf("test tag <%s> 1", protovalidator.TagStringContainsAny),
			FieldDesc:  "field 't_string_contains_any1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagStringContainsAny, Value: "xyz"},
			FieldValue: func() interface{} { return data.TStringContainsAny1 },
			BeforeFunc: func() { data.TStringContainsAny1 = "mn" },
			AfterFunc:  func() { data.TStringContainsAny1 = "x121" },
		},
		{
			Name:       fmt.Sprintf("test tag <%s> 1", protovalidator.TagStringNotContainsAny),
			FieldDesc:  "field 't_string_not_contains_any1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagStringNotContainsAny, Value: "xyz"},
			FieldValue: func() interface{} { return data.TStringNotContainsAny1 },
			BeforeFunc: func() { data.TStringNotContainsAny1 = "1y1" },
			AfterFunc:  func() { data.TStringNotContainsAny1 = "123" },
		},

		{
			Name:       fmt.Sprintf("test tag <%s> 1", protovalidator.TagStringUTF8),
			FieldDesc:  "field 't_string_utf8'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagStringUTF8, Value: nil},
			FieldValue: func() interface{} { return data.TStringUtf8 },
			BeforeFunc: func() {
				src := bytes.NewReader([]byte("中文abc"))
				encoder := simplifiedchinese.GBK.NewEncoder()
				tr := transform.NewReader(src, encoder)
				b, err := ioutil.ReadAll(tr)
				if err != nil {
					panic(err)
				}
				data.TStringUtf8 = string(b)
			},
			AfterFunc: func() { data.TStringUtf8 = "中文abc" },
		},

		{
			Name:       fmt.Sprintf("test tag <%s> 1", protovalidator.TagStringAscii),
			FieldDesc:  "field 't_string_ascii'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagStringAscii, Value: nil},
			FieldValue: func() interface{} { return data.TStringAscii },
			BeforeFunc: func() {
				src := bytes.NewReader([]byte("中文abc"))
				encoder := simplifiedchinese.GBK.NewEncoder()
				tr := transform.NewReader(src, encoder)
				b, err := ioutil.ReadAll(tr)
				if err != nil {
					panic(err)
				}
				data.TStringAscii = string(b)
			},
			AfterFunc: func() { data.TStringAscii = "abc" },
		},
		{
			Name:       fmt.Sprintf("test tag <%s> 1", protovalidator.TagStringPrintAscii),
			FieldDesc:  "field 't_string_print_ascii'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagStringPrintAscii, Value: nil},
			FieldValue: func() interface{} { return data.TStringPrintAscii },
			BeforeFunc: func() { data.TStringPrintAscii = string([]byte{3, 2, 1}) },
			AfterFunc:  func() { data.TStringPrintAscii = "abc" },
		},
		{
			Name:       fmt.Sprintf("test tag <%s> 1", protovalidator.TagStringBoolean),
			FieldDesc:  "field 't_string_boolean'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagStringBoolean, Value: nil},
			FieldValue: func() interface{} { return data.TStringBoolean },
			BeforeFunc: func() { data.TStringBoolean = "bool" },
			AfterFunc:  func() { data.TStringBoolean = "true" },
		},
		{
			Name:       fmt.Sprintf("test tag <%s> 1", protovalidator.TagStringLowercase),
			FieldDesc:  "field 't_string_lowercase'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagStringLowercase, Value: nil},
			FieldValue: func() interface{} { return data.TStringLowercase },
			BeforeFunc: func() { data.TStringLowercase = "AbC" },
			AfterFunc:  func() { data.TStringLowercase = "abc" },
		},
		{
			Name:       fmt.Sprintf("test tag <%s> 1", protovalidator.TagStringUppercase),
			FieldDesc:  "field 't_string_uppercase'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagStringUppercase, Value: nil},
			FieldValue: func() interface{} { return data.TStringUppercase },
			BeforeFunc: func() { data.TStringUppercase = "aBC" },
			AfterFunc:  func() { data.TStringUppercase = "ABC" },
		},
		{
			Name:       fmt.Sprintf("test tag <%s> 1", protovalidator.TagStringAlpha),
			FieldDesc:  "field 't_string_alpha'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagStringAlpha, Value: nil},
			FieldValue: func() interface{} { return data.TStringAlpha },
			BeforeFunc: func() { data.TStringAlpha = "123" },
			AfterFunc:  func() { data.TStringAlpha = "abc" },
		},
		{
			Name:       fmt.Sprintf("test tag <%s> 1", protovalidator.TagStringNumber),
			FieldDesc:  "field 't_string_number'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagStringNumber, Value: nil},
			FieldValue: func() interface{} { return data.TStringNumber },
			BeforeFunc: func() { data.TStringNumber = "abc" },
			AfterFunc:  func() { data.TStringNumber = "123" },
		},
		{
			Name:       fmt.Sprintf("test tag <%s> 1", protovalidator.TagStringAlphaNumber),
			FieldDesc:  "field 't_string_alpha_number'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagStringAlphaNumber, Value: nil},
			FieldValue: func() interface{} { return data.TStringAlphaNumber },
			BeforeFunc: func() { data.TStringAlphaNumber = "中文" },
			AfterFunc:  func() { data.TStringAlphaNumber = "abc123" },
		},
	}

	caseNetwork1 := []*CaseDesc{
		{
			Name:       fmt.Sprintf("test tag <%s> 1", protovalidator.TagStringIp),
			FieldDesc:  "field 't_string_ip'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagStringIp, Value: nil},
			FieldValue: func() interface{} { return data.TStringIp },
			BeforeFunc: func() { data.TStringIp = "127.0." },
			AfterFunc:  func() { data.TStringIp = "127.0.0.1" },
		},
		{
			Name:       fmt.Sprintf("test tag <%s> 2", protovalidator.TagStringIp),
			FieldDesc:  "field 't_string_ip'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagStringIp, Value: nil},
			FieldValue: func() interface{} { return data.TStringIp },
			BeforeFunc: func() { data.TStringIp = "AD80:0000:0000:0000:ABAA:0000:00C2:0002:xxxxx" },
			AfterFunc:  func() { data.TStringIp = "AD80:0000:0000:0000:ABAA:0000:00C2:0002" },
		},
		{
			Name:       fmt.Sprintf("test tag <%s> 1", protovalidator.TagStringIpv4),
			FieldDesc:  "field 't_string_ipv4'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagStringIpv4, Value: nil},
			FieldValue: func() interface{} { return data.TStringIpv4 },
			BeforeFunc: func() { data.TStringIpv4 = "127.0.0.1.1" },
			AfterFunc:  func() { data.TStringIpv4 = "127.0.0.1" },
		},
		{
			Name:       fmt.Sprintf("test tag <%s> 1", protovalidator.TagStringIpv6),
			FieldDesc:  "field 't_string_ipv6'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagStringIpv6, Value: nil},
			FieldValue: func() interface{} { return data.TStringIpv6 },
			BeforeFunc: func() { data.TStringIpv6 = "127.0.0.1" },
			AfterFunc:  func() { data.TStringIpv6 = "AD80:0000:0000:0000:ABAA:0000:00C2:0002" },
		},
		{
			Name:       fmt.Sprintf("test tag <%s> 1", protovalidator.TagStringIpAddr),
			FieldDesc:  "field 't_string_ip_addr'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagStringIpAddr, Value: nil},
			FieldValue: func() interface{} { return data.TStringIpAddr },
			BeforeFunc: func() { data.TStringIpAddr = "127.0.0.1.x" },
			AfterFunc:  func() { data.TStringIpAddr = "127.0.0.1" },
		},
		{
			Name:       fmt.Sprintf("test tag <%s> 1", protovalidator.TagStringIp4Addr),
			FieldDesc:  "field 't_string_ip4_addr'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagStringIp4Addr, Value: nil},
			FieldValue: func() interface{} { return data.TStringIp4Addr },
			BeforeFunc: func() { data.TStringIp4Addr = "127.0.0.1.x" },
			AfterFunc:  func() { data.TStringIp4Addr = "127.0.0.1" },
		},
		{
			Name:       fmt.Sprintf("test tag <%s> 1", protovalidator.TagStringIp6Addr),
			FieldDesc:  "field 't_string_ip6_addr'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagStringIp6Addr, Value: nil},
			FieldValue: func() interface{} { return data.TStringIp6Addr },
			BeforeFunc: func() { data.TStringIp6Addr = "127.0.0.1" },
			AfterFunc:  func() { data.TStringIp6Addr = "AD80:0000:0000:0000:ABAA:0000:00C2:0002" },
		},
		{
			Name:       fmt.Sprintf("test tag <%s> 1", protovalidator.TagStringCidr),
			FieldDesc:  "field 't_string_cidr'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagStringCidr, Value: nil},
			FieldValue: func() interface{} { return data.TStringCidr },
			BeforeFunc: func() { data.TStringCidr = "192.0.2.0/128" },
			AfterFunc:  func() { data.TStringCidr = "192.0.2.0/24" },
		},
		{
			Name:       fmt.Sprintf("test tag <%s> 1", protovalidator.TagStringCidrv4),
			FieldDesc:  "field 't_string_cidrv4'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagStringCidrv4, Value: nil},
			FieldValue: func() interface{} { return data.TStringCidrv4 },
			BeforeFunc: func() { data.TStringCidrv4 = "192.0.2.0/243" },
			AfterFunc:  func() { data.TStringCidrv4 = "192.0.2.0/24" },
		},
		{
			Name:       fmt.Sprintf("test tag <%s> 1", protovalidator.TagStringCidrv6),
			FieldDesc:  "field 't_string_cidrv6'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagStringCidrv6, Value: nil},
			FieldValue: func() interface{} { return data.TStringCidrv6 },
			BeforeFunc: func() { data.TStringCidrv6 = "2001:db8::/129" },
			AfterFunc:  func() { data.TStringCidrv6 = "2001:db8::/32" },
		},
		{
			Name:       fmt.Sprintf("test tag <%s> 1", protovalidator.TagStringMac),
			FieldDesc:  "field 't_string_mac'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagStringMac, Value: nil},
			FieldValue: func() interface{} { return data.TStringMac },
			BeforeFunc: func() { data.TStringMac = "00:00:5e:00:53:01:dsafadsfd" },
			AfterFunc:  func() { data.TStringMac = "00:00:5e:00:53:01" },
		},
		{
			Name:       fmt.Sprintf("test tag <%s> 1", protovalidator.TagStringTcpAddr),
			FieldDesc:  "field 't_string_tcp_addr'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagStringTcpAddr, Value: nil},
			FieldValue: func() interface{} { return data.TStringTcpAddr },
			BeforeFunc: func() { data.TStringTcpAddr = "" },
			AfterFunc:  func() { data.TStringTcpAddr = "127.0.0.1:80" },
		},
		{
			Name:       fmt.Sprintf("test tag <%s> 1", protovalidator.TagStringTcp4Addr),
			FieldDesc:  "field 't_string_tcp4_addr'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagStringTcp4Addr, Value: nil},
			FieldValue: func() interface{} { return data.TStringTcp4Addr },
			BeforeFunc: func() { data.TStringTcp4Addr = "" },
			AfterFunc:  func() { data.TStringTcp4Addr = "127.0.0.1:80" },
		},
		{
			Name:       fmt.Sprintf("test tag <%s> 1", protovalidator.TagStringTcp6Addr),
			FieldDesc:  "field 't_string_tcp6_addr'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagStringTcp6Addr, Value: nil},
			FieldValue: func() interface{} { return data.TStringTcp6Addr },
			BeforeFunc: func() { data.TStringTcp6Addr = "" },
			AfterFunc:  func() { data.TStringTcp6Addr = "[2001:db8::1]:http" },
		},
		{
			Name:       fmt.Sprintf("test tag <%s> 1", protovalidator.TagStringUdpAddr),
			FieldDesc:  "field 't_string_udp_addr'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagStringUdpAddr, Value: nil},
			FieldValue: func() interface{} { return data.TStringUdpAddr },
			BeforeFunc: func() { data.TStringUdpAddr = "" },
			AfterFunc:  func() { data.TStringUdpAddr = "127.0.0.1:domain" },
		},
		{
			Name:       fmt.Sprintf("test tag <%s> 1", protovalidator.TagStringUdp4Addr),
			FieldDesc:  "field 't_string_udp4_addr'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagStringUdp4Addr, Value: nil},
			FieldValue: func() interface{} { return data.TStringUdp4Addr },
			BeforeFunc: func() { data.TStringUdp4Addr = "" },
			AfterFunc:  func() { data.TStringUdp4Addr = "127.0.0.1:domain" },
		},
		{
			Name:       fmt.Sprintf("test tag <%s> 1", protovalidator.TagStringUdp6Addr),
			FieldDesc:  "field 't_string_udp6_addr'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagStringUdp6Addr, Value: nil},
			FieldValue: func() interface{} { return data.TStringUdp6Addr },
			BeforeFunc: func() { data.TStringUdp6Addr = "" },
			AfterFunc:  func() { data.TStringUdp6Addr = "[2001:db8::1]:domain" },
		},
		{
			Name:       fmt.Sprintf("test tag <%s> 1", protovalidator.TagStringUnixAddr),
			FieldDesc:  "field 't_string_unix_addr'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagStringUnixAddr, Value: nil},
			FieldValue: func() interface{} { return data.TStringUnixAddr },
			BeforeFunc: func() { data.TStringUnixAddr = "" },
			AfterFunc:  func() { data.TStringUnixAddr = "unix:///mysql/mysql.sock" },
		},
		{
			Name:       fmt.Sprintf("test tag <%s> 1", protovalidator.TagStringHostname),
			FieldDesc:  "field 't_string_hostname'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagStringHostname, Value: nil},
			FieldValue: func() interface{} { return data.TStringHostname },
			BeforeFunc: func() { data.TStringHostname = "" },
			AfterFunc:  func() { data.TStringHostname = "www.google.com" },
		},
		{
			Name:       fmt.Sprintf("test tag <%s> 1", protovalidator.TagStringHostnameRfc1123),
			FieldDesc:  "field 't_string_hostname_rfc1123'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagStringHostnameRfc1123, Value: nil},
			FieldValue: func() interface{} { return data.TStringHostnameRfc1123 },
			BeforeFunc: func() { data.TStringHostnameRfc1123 = "" },
			AfterFunc:  func() { data.TStringHostnameRfc1123 = "www.google.com" },
		},
		{
			Name:       fmt.Sprintf("test tag <%s> 1", protovalidator.TagStringHostnamePort),
			FieldDesc:  "field 't_string_hostname_port'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagStringHostnamePort, Value: nil},
			FieldValue: func() interface{} { return data.TStringHostnamePort },
			BeforeFunc: func() { data.TStringHostnamePort = "" },
			AfterFunc:  func() { data.TStringHostnamePort = "www.google.com:80" },
		},
		{
			Name:       fmt.Sprintf("test tag <%s> 1", protovalidator.TagStringDataURI),
			FieldDesc:  "field 't_string_data_uri'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagStringDataURI, Value: nil},
			FieldValue: func() interface{} { return data.TStringDataUri },
			BeforeFunc: func() { data.TStringDataUri = "" },
			AfterFunc:  func() { data.TStringDataUri = "data:text/html,WkhPTkdHVU9uaWhhbzEyMw==" },
		},
		{
			Name:       fmt.Sprintf("test tag <%s> 1", protovalidator.TagStringFQDN),
			FieldDesc:  "field 't_string_fqdn'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagStringFQDN, Value: nil},
			FieldValue: func() interface{} { return data.TStringFqdn },
			BeforeFunc: func() { data.TStringFqdn = "" },
			AfterFunc:  func() { data.TStringFqdn = "MacBook-Pro.local" },
		},
		{
			Name:       fmt.Sprintf("test tag <%s> 1", protovalidator.TagStringURI),
			FieldDesc:  "field 't_string_uri'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagStringURI, Value: nil},
			FieldValue: func() interface{} { return data.TStringUri },
			BeforeFunc: func() { data.TStringUri = "" },
			AfterFunc:  func() { data.TStringUri = "/v1/api" },
		},
		{
			Name:       fmt.Sprintf("test tag <%s> 1", protovalidator.TagStringURL),
			FieldDesc:  "field 't_string_url'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagStringURL, Value: nil},
			FieldValue: func() interface{} { return data.TStringUrl },
			BeforeFunc: func() { data.TStringUrl = "" },
			AfterFunc:  func() { data.TStringUrl = "https://www.google.com/v1/api" },
		},
		{
			Name:       fmt.Sprintf("test tag <%s> 1", protovalidator.TagStringURLEncoded),
			FieldDesc:  "field 't_string_url_encoded'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagStringURLEncoded, Value: nil},
			FieldValue: func() interface{} { return data.TStringUrlEncoded },
			BeforeFunc: func() { data.TStringUrlEncoded = "" },
			AfterFunc:  func() { data.TStringUrlEncoded = "https://www.google.com%2Fv1%2Fapi" },
		},
	}

	caseFormat1 := []*CaseDesc{
		{
			Name:       fmt.Sprintf("test tag <%s> 1", protovalidator.TagStringUnixCron),
			FieldDesc:  "field 't_string_unix_cron'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagStringUnixCron, Value: nil},
			FieldValue: func() interface{} { return data.TStringUnixCron },
			BeforeFunc: func() { data.TStringUnixCron = "" },
			AfterFunc:  func() { data.TStringUnixCron = "* * * * *" },
		},
		{
			Name:       fmt.Sprintf("test tag <%s> 1", protovalidator.TagStringEmail),
			FieldDesc:  "field 't_string_email'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagStringEmail, Value: nil},
			FieldValue: func() interface{} { return data.TStringEmail },
			BeforeFunc: func() { data.TStringEmail = "" },
			AfterFunc:  func() { data.TStringEmail = "xxx@gmail.com" },
		},
		{
			Name:       fmt.Sprintf("test tag <%s> 1", protovalidator.TagStringJSON),
			FieldDesc:  "field 't_string_json'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagStringJSON, Value: nil},
			FieldValue: func() interface{} { return data.TStringJson },
			BeforeFunc: func() { data.TStringJson = "" },
			AfterFunc:  func() { data.TStringJson = `{"k1":"v1"}` },
		},
		{
			Name:       fmt.Sprintf("test tag <%s> 1", protovalidator.TagStringJWT),
			FieldDesc:  "field 't_string_jwt'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagStringJWT, Value: nil},
			FieldValue: func() interface{} { return data.TStringJwt },
			BeforeFunc: func() { data.TStringJwt = "" },
			AfterFunc: func() {
				data.TStringJwt = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiYWRtaW4iOnRydWV9.TJVA95OrM7E2cBab30RMHrHDcEfxjoYZgeFONFh7HgQ"
			},
		},
		{
			Name:       fmt.Sprintf("test tag <%s> 1", protovalidator.TagStringHTML),
			FieldDesc:  "field 't_string_html'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagStringHTML, Value: nil},
			FieldValue: func() interface{} { return data.TStringHtml },
			BeforeFunc: func() { data.TStringHtml = "" },
			AfterFunc:  func() { data.TStringHtml = "<html></html>" },
		},
		{
			Name:       fmt.Sprintf("test tag <%s> 1", protovalidator.TagStringHTMLEncoded),
			FieldDesc:  "field 't_string_html_encoded'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagStringHTMLEncoded, Value: nil},
			FieldValue: func() interface{} { return data.TStringHtmlEncoded },
			BeforeFunc: func() { data.TStringHtmlEncoded = "" },
			AfterFunc:  func() { data.TStringHtmlEncoded = "&lt;html&gt;&lt;/html&gt;" },
		},
		{
			Name:       fmt.Sprintf("test tag <%s> 1", protovalidator.TagStringBase64),
			FieldDesc:  "field 't_string_base64'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagStringBase64, Value: nil},
			FieldValue: func() interface{} { return data.TStringBase64 },
			BeforeFunc: func() { data.TStringBase64 = "" },
			AfterFunc:  func() { data.TStringBase64 = "WkhPTkdHVU9uaWhhbzEyMw==" },
		},
		{
			Name:       fmt.Sprintf("test tag <%s> 1", protovalidator.TagStringBase64URL),
			FieldDesc:  "field 't_string_base64_url'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagStringBase64URL, Value: nil},
			FieldValue: func() interface{} { return data.TStringBase64Url },
			BeforeFunc: func() { data.TStringBase64Url = "" },
			AfterFunc:  func() { data.TStringBase64Url = "YWRmYWRzZg==" },
		},
		{
			Name:       fmt.Sprintf("test tag <%s> 1", protovalidator.TagStringHexadecimal),
			FieldDesc:  "field 't_string_hexadecimal'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagStringHexadecimal, Value: nil},
			FieldValue: func() interface{} { return data.TStringHexadecimal },
			BeforeFunc: func() { data.TStringHexadecimal = "" },
			AfterFunc:  func() { data.TStringHexadecimal = "6461666164736661647366736461" },
		},
		{
			Name:       fmt.Sprintf("test tag <%s> 1", protovalidator.TagStringDatetime),
			FieldDesc:  "field 't_string_datetime'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagStringDatetime, Value: "2006-01-02 15:04:05"},
			FieldValue: func() interface{} { return data.TStringDatetime },
			BeforeFunc: func() { data.TStringDatetime = "" },
			AfterFunc:  func() { data.TStringDatetime = "2020-11-02 15:04:05" },
		},
		{
			Name:       fmt.Sprintf("test tag <%s> 1", protovalidator.TagStringTimezone),
			FieldDesc:  "field 't_string_timezone'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagStringTimezone, Value: nil},
			FieldValue: func() interface{} { return data.TStringTimezone },
			BeforeFunc: func() { data.TStringTimezone = "" },
			AfterFunc:  func() { data.TStringTimezone = "UTC" },
		},
		{
			Name:       fmt.Sprintf("test tag <%s> 1", protovalidator.TagStringUUID),
			FieldDesc:  "field 't_string_uuid'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagStringUUID, Value: nil},
			FieldValue: func() interface{} { return data.TStringUuid },
			BeforeFunc: func() { data.TStringUuid = "" },
			AfterFunc:  func() { data.TStringUuid = "80363646-7361-11ec-b4f0-020017000b7b" },
		},
		{
			Name:       fmt.Sprintf("test tag <%s> 1", protovalidator.TagStringUUID1),
			FieldDesc:  "field 't_string_uuid1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagStringUUID1, Value: nil},
			FieldValue: func() interface{} { return data.TStringUuid1 },
			BeforeFunc: func() { data.TStringUuid1 = "" },
			AfterFunc:  func() { data.TStringUuid1 = "80363646-7361-11ec-b4f0-020017000b7b" },
		},
		{
			Name:       fmt.Sprintf("test tag <%s> 1", protovalidator.TagStringUUID3),
			FieldDesc:  "field 't_string_uuid3'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagStringUUID3, Value: nil},
			FieldValue: func() interface{} { return data.TStringUuid3 },
			BeforeFunc: func() { data.TStringUuid3 = "" },
			AfterFunc:  func() { data.TStringUuid3 = "fa1f0699-6eac-3c47-b0f1-9d67ce5d0b76" },
		},
		{
			Name:       fmt.Sprintf("test tag <%s> 1", protovalidator.TagStringUUID4),
			FieldDesc:  "field 't_string_uuid4'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagStringUUID4, Value: nil},
			FieldValue: func() interface{} { return data.TStringUuid4 },
			BeforeFunc: func() { data.TStringUuid4 = "" },
			AfterFunc:  func() { data.TStringUuid4 = "3af73fea-62c7-4a00-8314-942a1841288a" },
		},
		{
			Name:       fmt.Sprintf("test tag <%s> 1", protovalidator.TagStringUUID5),
			FieldDesc:  "field 't_string_uuid5'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagStringUUID5, Value: nil},
			FieldValue: func() interface{} { return data.TStringUuid5 },
			BeforeFunc: func() { data.TStringUuid5 = "" },
			AfterFunc:  func() { data.TStringUuid5 = "78381e24-2d01-5a28-b2bf-9eedab6605ed" },
		},
	}

	var cases []*CaseDesc
	cases = append(cases, caseGeneral1...)
	cases = append(cases, caseGeneral2...)
	cases = append(cases, caseNetwork1...)
	cases = append(cases, caseFormat1...)

	msgName := "ValidStringTagsGeneral1"
	runCases(t, data, msgName, cases)

	{
		err := data.Validate()
		require.Nil(t, err)
	}
}

func Test_GoValidator_ValidStringTagsOptional1(t *testing.T) {
	data := &govalidatortest.ValidStringTagsOptional1{}
	{
		err := data.Validate()
		require.NotNil(t, err)
	}

	caseGeneral1 := []*CaseDesc{
		{
			Name:       fmt.Sprintf("test tag <%s> 1", protovalidator.TagStringEq),
			FieldDesc:  "field 't_string_eq1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagStringEq, Value: "b"},
			FieldValue: func() interface{} { return data.TStringEq1 },
			BeforeFunc: func() {},
			AfterFunc:  func() { x := "b"; data.TStringEq1 = &x },
		},
		{
			Name:       fmt.Sprintf("test tag <%s> 1", protovalidator.TagStringNe),
			FieldDesc:  "field 't_string_ne1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagStringNe, Value: "b"},
			FieldValue: func() interface{} { return data.TStringNe1 },
			BeforeFunc: func() { x := "b"; data.TStringNe1 = &x },
			AfterFunc:  func() { x := "a"; data.TStringNe1 = &x },
		},
		{
			Name:       fmt.Sprintf("test tag <%s> 1", protovalidator.TagStringLt),
			FieldDesc:  "field 't_string_lt1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagStringLt, Value: "c"},
			FieldValue: func() interface{} { return data.TStringLt1 },
			BeforeFunc: func() { x := "c"; data.TStringLt1 = &x },
			AfterFunc:  func() { x := "b"; data.TStringLt1 = &x },
		},
		{
			Name:       fmt.Sprintf("test tag <%s> 1", protovalidator.TagStringGt),
			FieldDesc:  "field 't_string_gt1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagStringGt, Value: "c"},
			FieldValue: func() interface{} { return data.TStringGt1 },
			BeforeFunc: func() { x := "c"; data.TStringGt1 = &x },
			AfterFunc:  func() { x := "d"; data.TStringGt1 = &x },
		},
		{
			Name:       fmt.Sprintf("test tag <%s> 1", protovalidator.TagStringLte),
			FieldDesc:  "field 't_string_lte1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagStringLte, Value: "c"},
			FieldValue: func() interface{} { return data.TStringLte1 },
			BeforeFunc: func() { x := "d"; data.TStringLte1 = &x },
			AfterFunc:  func() { x := "c"; data.TStringLte1 = &x },
		},
		{
			Name:       fmt.Sprintf("test tag <%s> 1", protovalidator.TagStringGte),
			FieldDesc:  "field 't_string_gte1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagStringGte, Value: "c"},
			FieldValue: func() interface{} { return data.TStringGte1 },
			BeforeFunc: func() { x := "b"; data.TStringGte1 = &x },
			AfterFunc:  func() { x := "c"; data.TStringGte1 = &x },
		},
		{
			Name:       fmt.Sprintf("test tag <%s> 1", protovalidator.TagStringIn),
			FieldDesc:  "field 't_string_in1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagStringIn, Value: []string{"a", "b", "c"}},
			FieldValue: func() interface{} { return data.TStringIn1 },
			BeforeFunc: func() { x := "d"; data.TStringIn1 = &x },
			AfterFunc:  func() { x := "b"; data.TStringIn1 = &x },
		},
		{
			Name:       fmt.Sprintf("test tag <%s> 1", protovalidator.TagStringNotIn),
			FieldDesc:  "field 't_string_not_in1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagStringNotIn, Value: []string{"x", "y", "z"}},
			FieldValue: func() interface{} { return data.TStringNotIn1 },
			BeforeFunc: func() { x := "x"; data.TStringNotIn1 = &x },
			AfterFunc:  func() { x := "v"; data.TStringNotIn1 = &x },
		},
		{
			Name:       fmt.Sprintf("test tag <%s> 1", protovalidator.TagStringCharLenEq),
			FieldDesc:  "field 't_string_char_len_eq1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagStringCharLenEq, Value: 3},
			FieldValue: func() interface{} { return utf8.RuneCountInString(*data.TStringCharLenEq1) },
			BeforeFunc: func() { x := "x"; data.TStringCharLenEq1 = &x },
			AfterFunc:  func() { x := "a中b"; data.TStringCharLenEq1 = &x },
		},
		{
			Name:       fmt.Sprintf("test tag <%s> 1", protovalidator.TagStringCharLenNe),
			FieldDesc:  "field 't_string_char_len_ne1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagStringCharLenNe, Value: 3},
			FieldValue: func() interface{} { return utf8.RuneCountInString(*data.TStringCharLenNe1) },
			BeforeFunc: func() { x := "a中b"; data.TStringCharLenNe1 = &x },
			AfterFunc:  func() { x := "中"; data.TStringCharLenNe1 = &x },
		},
		{
			Name:       fmt.Sprintf("test tag <%s> 1", protovalidator.TagStringCharLenGt),
			FieldDesc:  "field 't_string_char_len_gt1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagStringCharLenGt, Value: 3},
			FieldValue: func() interface{} { return utf8.RuneCountInString(*data.TStringCharLenGt1) },
			BeforeFunc: func() { x := "a中b"; data.TStringCharLenGt1 = &x },
			AfterFunc:  func() { x := "a中b文"; data.TStringCharLenGt1 = &x },
		},
		{
			Name:       fmt.Sprintf("test tag <%s> 1", protovalidator.TagStringCharLenLt),
			FieldDesc:  "field 't_string_char_len_lt1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagStringCharLenLt, Value: 3},
			FieldValue: func() interface{} { return utf8.RuneCountInString(*data.TStringCharLenLt1) },
			BeforeFunc: func() { x := "a中b"; data.TStringCharLenLt1 = &x },
			AfterFunc:  func() { x := "中"; data.TStringCharLenLt1 = &x },
		},
		{
			Name:       fmt.Sprintf("test tag <%s> 1", protovalidator.TagStringCharLenGte),
			FieldDesc:  "field 't_string_char_len_gte1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagStringCharLenGte, Value: 3},
			FieldValue: func() interface{} { return utf8.RuneCountInString(*data.TStringCharLenGte1) },
			BeforeFunc: func() { x := "a中"; data.TStringCharLenGte1 = &x },
			AfterFunc:  func() { x := "a中b文"; data.TStringCharLenGte1 = &x },
		},
		{
			Name:       fmt.Sprintf("test tag <%s> 1", protovalidator.TagStringCharLenLte),
			FieldDesc:  "field 't_string_char_len_lte1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagStringCharLenLte, Value: 3},
			FieldValue: func() interface{} { return utf8.RuneCountInString(*data.TStringCharLenLte1) },
			BeforeFunc: func() { x := "a中b文"; data.TStringCharLenLte1 = &x },
			AfterFunc:  func() { x := "a中b"; data.TStringCharLenLte1 = &x },
		},

		{
			Name:       fmt.Sprintf("test tag <%s> 1", protovalidator.TagStringByteLenEq),
			FieldDesc:  "field 't_string_byte_len_eq1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagStringByteLenEq, Value: 5},
			FieldValue: func() interface{} { return len(*data.TStringByteLenEq1) },
			BeforeFunc: func() { x := "a中b文C"; data.TStringByteLenEq1 = &x },
			AfterFunc:  func() { x := "a中b"; data.TStringByteLenEq1 = &x },
		},
		{
			Name:       fmt.Sprintf("test tag <%s> 1", protovalidator.TagStringByteLenNe),
			FieldDesc:  "field 't_string_byte_len_ne1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagStringByteLenNe, Value: 5},
			FieldValue: func() interface{} { return len(*data.TStringByteLenNe1) },
			BeforeFunc: func() { x := "a中b"; data.TStringByteLenNe1 = &x },
			AfterFunc:  func() { x := "a"; data.TStringByteLenNe1 = &x },
		},
		{
			Name:       fmt.Sprintf("test tag <%s> 1", protovalidator.TagStringByteLenGt),
			FieldDesc:  "field 't_string_byte_len_gt1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagStringByteLenGt, Value: 5},
			FieldValue: func() interface{} { return len(*data.TStringByteLenGt1) },
			BeforeFunc: func() { x := "a中b"; data.TStringByteLenGt1 = &x },
			AfterFunc:  func() { x := "a中bc"; data.TStringByteLenGt1 = &x },
		},
		{
			Name:       fmt.Sprintf("test tag <%s> 1", protovalidator.TagStringByteLenLt),
			FieldDesc:  "field 't_string_byte_len_lt1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagStringByteLenLt, Value: 5},
			FieldValue: func() interface{} { return len(*data.TStringByteLenLt1) },
			BeforeFunc: func() { x := "a中b文"; data.TStringByteLenLt1 = &x },
			AfterFunc:  func() { x := "a中"; data.TStringByteLenLt1 = &x },
		},
		{
			Name:       fmt.Sprintf("test tag <%s> 1", protovalidator.TagStringByteLenGte),
			FieldDesc:  "field 't_string_byte_len_gte1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagStringByteLenGte, Value: 5},
			FieldValue: func() interface{} { return len(*data.TStringByteLenGte1) },
			BeforeFunc: func() { x := "a中"; data.TStringByteLenGte1 = &x },
			AfterFunc:  func() { x := "a中b"; data.TStringByteLenGte1 = &x },
		},
		{
			Name:       fmt.Sprintf("test tag <%s> 1", protovalidator.TagStringByteLenLte),
			FieldDesc:  "field 't_string_byte_len_lte1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagStringByteLenLte, Value: 5},
			FieldValue: func() interface{} { return len(*data.TStringByteLenLte1) },
			BeforeFunc: func() { x := "a中b文"; data.TStringByteLenLte1 = &x },
			AfterFunc:  func() { x := "a中b"; data.TStringByteLenLte1 = &x },
		},
	}

	caseGeneral1NilPoint := []*CaseDesc{
		{
			Name:       fmt.Sprintf("test tag <%s> 2", protovalidator.TagStringEq),
			FieldDesc:  "field 't_string_eq1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagStringEq, Value: "b"},
			FieldValue: func() interface{} { return data.TStringEq1 },
			BeforeFunc: func() { data.TStringEq1 = nil },
			AfterFunc:  func() { x := "b"; data.TStringEq1 = &x },
		},
		{
			Name:       fmt.Sprintf("test tag <%s> 2", protovalidator.TagStringNe),
			FieldDesc:  "field 't_string_ne1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagStringNe, Value: "b"},
			FieldValue: func() interface{} { return data.TStringNe1 },
			BeforeFunc: func() { data.TStringNe1 = nil },
			AfterFunc:  func() { x := "a"; data.TStringNe1 = &x },
		},
		{
			Name:       fmt.Sprintf("test tag <%s> 2", protovalidator.TagStringLt),
			FieldDesc:  "field 't_string_lt1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagStringLt, Value: "c"},
			FieldValue: func() interface{} { return data.TStringLt1 },
			BeforeFunc: func() { data.TStringLt1 = nil },
			AfterFunc:  func() { x := "b"; data.TStringLt1 = &x },
		},
		{
			Name:       fmt.Sprintf("test tag <%s> 2", protovalidator.TagStringGt),
			FieldDesc:  "field 't_string_gt1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagStringGt, Value: "c"},
			FieldValue: func() interface{} { return data.TStringGt1 },
			BeforeFunc: func() { data.TStringGt1 = nil },
			AfterFunc:  func() { x := "d"; data.TStringGt1 = &x },
		},
		{
			Name:       fmt.Sprintf("test tag <%s> 2", protovalidator.TagStringLte),
			FieldDesc:  "field 't_string_lte1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagStringLte, Value: "c"},
			FieldValue: func() interface{} { return data.TStringLte1 },
			BeforeFunc: func() { data.TStringLte1 = nil },
			AfterFunc:  func() { x := "c"; data.TStringLte1 = &x },
		},
		{
			Name:       fmt.Sprintf("test tag <%s> 2", protovalidator.TagStringGte),
			FieldDesc:  "field 't_string_gte1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagStringGte, Value: "c"},
			FieldValue: func() interface{} { return data.TStringGte1 },
			BeforeFunc: func() { data.TStringGte1 = nil },
			AfterFunc:  func() { x := "c"; data.TStringGte1 = &x },
		},
		{
			Name:       fmt.Sprintf("test tag <%s> 2", protovalidator.TagStringIn),
			FieldDesc:  "field 't_string_in1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagStringIn, Value: []string{"a", "b", "c"}},
			FieldValue: func() interface{} { return data.TStringIn1 },
			BeforeFunc: func() { data.TStringIn1 = nil },
			AfterFunc:  func() { x := "b"; data.TStringIn1 = &x },
		},
		{
			Name:       fmt.Sprintf("test tag <%s> 2", protovalidator.TagStringNotIn),
			FieldDesc:  "field 't_string_not_in1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagStringNotIn, Value: []string{"x", "y", "z"}},
			FieldValue: func() interface{} { return data.TStringNotIn1 },
			BeforeFunc: func() { data.TStringNotIn1 = nil },
			AfterFunc:  func() { x := "v"; data.TStringNotIn1 = &x },
		},
		{
			Name:       fmt.Sprintf("test tag <%s> 1", protovalidator.TagStringCharLenEq),
			FieldDesc:  "field 't_string_char_len_eq1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagStringCharLenEq, Value: 3},
			FieldValue: func() interface{} { return nil },
			BeforeFunc: func() { data.TStringCharLenEq1 = nil },
			AfterFunc:  func() { x := "a中b"; data.TStringCharLenEq1 = &x },
		},
		{
			Name:       fmt.Sprintf("test tag <%s> 2", protovalidator.TagStringCharLenNe),
			FieldDesc:  "field 't_string_char_len_ne1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagStringCharLenNe, Value: 3},
			FieldValue: func() interface{} { return nil },
			BeforeFunc: func() { data.TStringCharLenNe1 = nil },
			AfterFunc:  func() { x := "中"; data.TStringCharLenNe1 = &x },
		},
		{
			Name:       fmt.Sprintf("test tag <%s> 2", protovalidator.TagStringCharLenGt),
			FieldDesc:  "field 't_string_char_len_gt1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagStringCharLenGt, Value: 3},
			FieldValue: func() interface{} { return nil },
			BeforeFunc: func() { data.TStringCharLenGt1 = nil },
			AfterFunc:  func() { x := "a中b文"; data.TStringCharLenGt1 = &x },
		},
		{
			Name:       fmt.Sprintf("test tag <%s> 2", protovalidator.TagStringCharLenLt),
			FieldDesc:  "field 't_string_char_len_lt1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagStringCharLenLt, Value: 3},
			FieldValue: func() interface{} { return nil },
			BeforeFunc: func() { data.TStringCharLenLt1 = nil },
			AfterFunc:  func() { x := "中"; data.TStringCharLenLt1 = &x },
		},
		{
			Name:       fmt.Sprintf("test tag <%s> 2", protovalidator.TagStringCharLenGte),
			FieldDesc:  "field 't_string_char_len_gte1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagStringCharLenGte, Value: 3},
			FieldValue: func() interface{} { return nil },
			BeforeFunc: func() { data.TStringCharLenGte1 = nil },
			AfterFunc:  func() { x := "a中b文"; data.TStringCharLenGte1 = &x },
		},
		{
			Name:       fmt.Sprintf("test tag <%s> 2", protovalidator.TagStringCharLenLte),
			FieldDesc:  "field 't_string_char_len_lte1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagStringCharLenLte, Value: 3},
			FieldValue: func() interface{} { return nil },
			BeforeFunc: func() { data.TStringCharLenLte1 = nil },
			AfterFunc:  func() { x := "a中b"; data.TStringCharLenLte1 = &x },
		},

		{
			Name:       fmt.Sprintf("test tag <%s> 2", protovalidator.TagStringByteLenEq),
			FieldDesc:  "field 't_string_byte_len_eq1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagStringByteLenEq, Value: 5},
			FieldValue: func() interface{} { return nil },
			BeforeFunc: func() { data.TStringByteLenEq1 = nil },
			AfterFunc:  func() { x := "a中b"; data.TStringByteLenEq1 = &x },
		},
		{
			Name:       fmt.Sprintf("test tag <%s> 2", protovalidator.TagStringByteLenNe),
			FieldDesc:  "field 't_string_byte_len_ne1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagStringByteLenNe, Value: 5},
			FieldValue: func() interface{} { return nil },
			BeforeFunc: func() { data.TStringByteLenNe1 = nil },
			AfterFunc:  func() { x := "a"; data.TStringByteLenNe1 = &x },
		},
		{
			Name:       fmt.Sprintf("test tag <%s> 2", protovalidator.TagStringByteLenGt),
			FieldDesc:  "field 't_string_byte_len_gt1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagStringByteLenGt, Value: 5},
			FieldValue: func() interface{} { return nil },
			BeforeFunc: func() { data.TStringByteLenGt1 = nil },
			AfterFunc:  func() { x := "a中bc"; data.TStringByteLenGt1 = &x },
		},
		{
			Name:       fmt.Sprintf("test tag <%s> 2", protovalidator.TagStringByteLenLt),
			FieldDesc:  "field 't_string_byte_len_lt1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagStringByteLenLt, Value: 5},
			FieldValue: func() interface{} { return nil },
			BeforeFunc: func() { data.TStringByteLenLt1 = nil },
			AfterFunc:  func() { x := "a中"; data.TStringByteLenLt1 = &x },
		},
		{
			Name:       fmt.Sprintf("test tag <%s> 2", protovalidator.TagStringByteLenGte),
			FieldDesc:  "field 't_string_byte_len_gte1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagStringByteLenGte, Value: 5},
			FieldValue: func() interface{} { return nil },
			BeforeFunc: func() { data.TStringByteLenGte1 = nil },
			AfterFunc:  func() { x := "a中b"; data.TStringByteLenGte1 = &x },
		},
		{
			Name:       fmt.Sprintf("test tag <%s> 2", protovalidator.TagStringByteLenLte),
			FieldDesc:  "field 't_string_byte_len_lte1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagStringByteLenLte, Value: 5},
			FieldValue: func() interface{} { return nil },
			BeforeFunc: func() { data.TStringByteLenLte1 = nil },
			AfterFunc:  func() { x := "a中b"; data.TStringByteLenLte1 = &x },
		},
	}

	caseGeneral2 := []*CaseDesc{
		{
			Name:       fmt.Sprintf("test tag <%s> 1", protovalidator.TagStringRegex),
			FieldDesc:  "field 't_string_regex1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagStringRegex, Value: `^[0-9]+.jar$`},
			FieldValue: func() interface{} { return data.TStringRegex1 },
			BeforeFunc: func() { x := "xx"; data.TStringRegex1 = &x },
			AfterFunc:  func() { x := "0001.jar"; data.TStringRegex1 = &x },
		},
		{
			Name:       fmt.Sprintf("test tag <%s> 1", protovalidator.TagStringPrefix),
			FieldDesc:  "field 't_string_prefix1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagStringPrefix, Value: "prefix"},
			FieldValue: func() interface{} { return data.TStringPrefix1 },
			BeforeFunc: func() { x := "1"; data.TStringPrefix1 = &x },
			AfterFunc:  func() { x := "prefix-xxx1"; data.TStringPrefix1 = &x },
		},
		{
			Name:       fmt.Sprintf("test tag <%s> 1", protovalidator.TagStringNoPrefix),
			FieldDesc:  "field 't_string_no_prefix1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagStringNoPrefix, Value: "prefix"},
			FieldValue: func() interface{} { return data.TStringNoPrefix1 },
			BeforeFunc: func() { x := "prefix-xxx1"; data.TStringNoPrefix1 = &x },
			AfterFunc:  func() { x := "xxx1"; data.TStringNoPrefix1 = &x },
		},

		{
			Name:       fmt.Sprintf("test tag <%s> 1", protovalidator.TagStringSuffix),
			FieldDesc:  "field 't_string_suffix1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagStringSuffix, Value: "suffix"},
			FieldValue: func() interface{} { return data.TStringSuffix1 },
			BeforeFunc: func() { x := "xxx"; data.TStringSuffix1 = &x },
			AfterFunc:  func() { x := "xxx-suffix"; data.TStringSuffix1 = &x },
		},
		{
			Name:       fmt.Sprintf("test tag <%s> 1", protovalidator.TagStringNoSuffix),
			FieldDesc:  "field 't_string_no_suffix1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagStringNoSuffix, Value: "suffix"},
			FieldValue: func() interface{} { return data.TStringNoSuffix1 },
			BeforeFunc: func() { x := "xxx-suffix"; data.TStringNoSuffix1 = &x },
			AfterFunc:  func() { x := "xxx"; data.TStringNoSuffix1 = &x },
		},

		{
			Name:       fmt.Sprintf("test tag <%s> 1", protovalidator.TagStringContains),
			FieldDesc:  "field 't_string_contains1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagStringContains, Value: "abc"},
			FieldValue: func() interface{} { return data.TStringContains1 },
			BeforeFunc: func() { x := "1a2b3c"; data.TStringContains1 = &x },
			AfterFunc:  func() { x := "1abc23"; data.TStringContains1 = &x },
		},
		{
			Name:       fmt.Sprintf("test tag <%s> 1", protovalidator.TagStringNotContains),
			FieldDesc:  "field 't_string_no_contains1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagStringNotContains, Value: "abc"},
			FieldValue: func() interface{} { return data.TStringNoContains1 },
			BeforeFunc: func() { x := "1abc23"; data.TStringNoContains1 = &x },
			AfterFunc:  func() { x := "1a2b3c"; data.TStringNoContains1 = &x },
		},

		{
			Name:       fmt.Sprintf("test tag <%s> 1", protovalidator.TagStringContainsAny),
			FieldDesc:  "field 't_string_contains_any1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagStringContainsAny, Value: "xyz"},
			FieldValue: func() interface{} { return data.TStringContainsAny1 },
			BeforeFunc: func() { x := "mn"; data.TStringContainsAny1 = &x },
			AfterFunc:  func() { x := "x121"; data.TStringContainsAny1 = &x },
		},
		{
			Name:       fmt.Sprintf("test tag <%s> 1", protovalidator.TagStringNotContainsAny),
			FieldDesc:  "field 't_string_not_contains_any1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagStringNotContainsAny, Value: "xyz"},
			FieldValue: func() interface{} { return data.TStringNotContainsAny1 },
			BeforeFunc: func() { x := "1y1"; data.TStringNotContainsAny1 = &x },
			AfterFunc:  func() { x := "123"; data.TStringNotContainsAny1 = &x },
		},

		{
			Name:       fmt.Sprintf("test tag <%s> 1", protovalidator.TagStringUTF8),
			FieldDesc:  "field 't_string_utf8'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagStringUTF8, Value: nil},
			FieldValue: func() interface{} { return data.TStringUtf8 },
			BeforeFunc: func() {
				src := bytes.NewReader([]byte("中文abc"))
				encoder := simplifiedchinese.GBK.NewEncoder()
				tr := transform.NewReader(src, encoder)
				b, err := ioutil.ReadAll(tr)
				if err != nil {
					panic(err)
				}
				x := string(b)
				data.TStringUtf8 = &x
			},
			AfterFunc: func() { x := "中文abc"; data.TStringUtf8 = &x },
		},

		{
			Name:       fmt.Sprintf("test tag <%s> 1", protovalidator.TagStringAscii),
			FieldDesc:  "field 't_string_ascii'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagStringAscii, Value: nil},
			FieldValue: func() interface{} { return data.TStringAscii },
			BeforeFunc: func() {
				src := bytes.NewReader([]byte("中文abc"))
				encoder := simplifiedchinese.GBK.NewEncoder()
				tr := transform.NewReader(src, encoder)
				b, err := ioutil.ReadAll(tr)
				if err != nil {
					panic(err)
				}
				x := string(b)
				data.TStringAscii = &x
			},
			AfterFunc: func() { x := "abc"; data.TStringAscii = &x },
		},
		{
			Name:       fmt.Sprintf("test tag <%s> 1", protovalidator.TagStringPrintAscii),
			FieldDesc:  "field 't_string_print_ascii'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagStringPrintAscii, Value: nil},
			FieldValue: func() interface{} { return data.TStringPrintAscii },
			BeforeFunc: func() { x := string([]byte{3, 2, 1}); data.TStringPrintAscii = &x },
			AfterFunc:  func() { x := "abc"; data.TStringPrintAscii = &x },
		},
		{
			Name:       fmt.Sprintf("test tag <%s> 1", protovalidator.TagStringBoolean),
			FieldDesc:  "field 't_string_boolean'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagStringBoolean, Value: nil},
			FieldValue: func() interface{} { return data.TStringBoolean },
			BeforeFunc: func() { x := "bool"; data.TStringBoolean = &x },
			AfterFunc:  func() { x := "true"; data.TStringBoolean = &x },
		},
		{
			Name:       fmt.Sprintf("test tag <%s> 1", protovalidator.TagStringLowercase),
			FieldDesc:  "field 't_string_lowercase'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagStringLowercase, Value: nil},
			FieldValue: func() interface{} { return data.TStringLowercase },
			BeforeFunc: func() { x := "AbC"; data.TStringLowercase = &x },
			AfterFunc:  func() { x := "abc"; data.TStringLowercase = &x },
		},
		{
			Name:       fmt.Sprintf("test tag <%s> 1", protovalidator.TagStringUppercase),
			FieldDesc:  "field 't_string_uppercase'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagStringUppercase, Value: nil},
			FieldValue: func() interface{} { return data.TStringUppercase },
			BeforeFunc: func() { x := "aBC"; data.TStringUppercase = &x },
			AfterFunc:  func() { x := "ABC"; data.TStringUppercase = &x },
		},
		{
			Name:       fmt.Sprintf("test tag <%s> 1", protovalidator.TagStringAlpha),
			FieldDesc:  "field 't_string_alpha'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagStringAlpha, Value: nil},
			FieldValue: func() interface{} { return data.TStringAlpha },
			BeforeFunc: func() { x := "123"; data.TStringAlpha = &x },
			AfterFunc:  func() { x := "abc"; data.TStringAlpha = &x },
		},
		{
			Name:       fmt.Sprintf("test tag <%s> 1", protovalidator.TagStringNumber),
			FieldDesc:  "field 't_string_number'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagStringNumber, Value: nil},
			FieldValue: func() interface{} { return data.TStringNumber },
			BeforeFunc: func() { x := "abc"; data.TStringNumber = &x },
			AfterFunc:  func() { x := "123"; data.TStringNumber = &x },
		},
		{
			Name:       fmt.Sprintf("test tag <%s> 1", protovalidator.TagStringAlphaNumber),
			FieldDesc:  "field 't_string_alpha_number'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagStringAlphaNumber, Value: nil},
			FieldValue: func() interface{} { return data.TStringAlphaNumber },
			BeforeFunc: func() { x := "中文"; data.TStringAlphaNumber = &x },
			AfterFunc:  func() { x := "abc123"; data.TStringAlphaNumber = &x },
		},
	}

	caseGeneral2NilPoint := []*CaseDesc{
		{
			Name:       fmt.Sprintf("test tag <%s> 2", protovalidator.TagStringRegex),
			FieldDesc:  "field 't_string_regex1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagStringRegex, Value: `^[0-9]+.jar$`},
			FieldValue: func() interface{} { return data.TStringRegex1 },
			BeforeFunc: func() { data.TStringRegex1 = nil },
			AfterFunc:  func() { x := "0001.jar"; data.TStringRegex1 = &x },
		},
		{
			Name:       fmt.Sprintf("test tag <%s> 2", protovalidator.TagStringPrefix),
			FieldDesc:  "field 't_string_prefix1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagStringPrefix, Value: "prefix"},
			FieldValue: func() interface{} { return data.TStringPrefix1 },
			BeforeFunc: func() { data.TStringPrefix1 = nil },
			AfterFunc:  func() { x := "prefix-xxx1"; data.TStringPrefix1 = &x },
		},
		{
			Name:       fmt.Sprintf("test tag <%s> 2", protovalidator.TagStringNoPrefix),
			FieldDesc:  "field 't_string_no_prefix1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagStringNoPrefix, Value: "prefix"},
			FieldValue: func() interface{} { return data.TStringNoPrefix1 },
			BeforeFunc: func() { data.TStringNoPrefix1 = nil },
			AfterFunc:  func() { x := "xxx1"; data.TStringNoPrefix1 = &x },
		},

		{
			Name:       fmt.Sprintf("test tag <%s> 2", protovalidator.TagStringSuffix),
			FieldDesc:  "field 't_string_suffix1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagStringSuffix, Value: "suffix"},
			FieldValue: func() interface{} { return data.TStringSuffix1 },
			BeforeFunc: func() { data.TStringSuffix1 = nil },
			AfterFunc:  func() { x := "xxx-suffix"; data.TStringSuffix1 = &x },
		},
		{
			Name:       fmt.Sprintf("test tag <%s> 2", protovalidator.TagStringNoSuffix),
			FieldDesc:  "field 't_string_no_suffix1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagStringNoSuffix, Value: "suffix"},
			FieldValue: func() interface{} { return data.TStringNoSuffix1 },
			BeforeFunc: func() { data.TStringNoSuffix1 = nil },
			AfterFunc:  func() { x := "xxx"; data.TStringNoSuffix1 = &x },
		},

		{
			Name:       fmt.Sprintf("test tag <%s> 2", protovalidator.TagStringContains),
			FieldDesc:  "field 't_string_contains1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagStringContains, Value: "abc"},
			FieldValue: func() interface{} { return data.TStringContains1 },
			BeforeFunc: func() { data.TStringContains1 = nil },
			AfterFunc:  func() { x := "1abc23"; data.TStringContains1 = &x },
		},
		{
			Name:       fmt.Sprintf("test tag <%s> 2", protovalidator.TagStringNotContains),
			FieldDesc:  "field 't_string_no_contains1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagStringNotContains, Value: "abc"},
			FieldValue: func() interface{} { return data.TStringNoContains1 },
			BeforeFunc: func() { data.TStringNoContains1 = nil },
			AfterFunc:  func() { x := "1a2b3c"; data.TStringNoContains1 = &x },
		},

		{
			Name:       fmt.Sprintf("test tag <%s> 2", protovalidator.TagStringContainsAny),
			FieldDesc:  "field 't_string_contains_any1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagStringContainsAny, Value: "xyz"},
			FieldValue: func() interface{} { return data.TStringContainsAny1 },
			BeforeFunc: func() { data.TStringContainsAny1 = nil },
			AfterFunc:  func() { x := "x121"; data.TStringContainsAny1 = &x },
		},
		{
			Name:       fmt.Sprintf("test tag <%s> 2", protovalidator.TagStringNotContainsAny),
			FieldDesc:  "field 't_string_not_contains_any1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagStringNotContainsAny, Value: "xyz"},
			FieldValue: func() interface{} { return data.TStringNotContainsAny1 },
			BeforeFunc: func() { data.TStringNotContainsAny1 = nil },
			AfterFunc:  func() { x := "123"; data.TStringNotContainsAny1 = &x },
		},

		{
			Name:       fmt.Sprintf("test tag <%s> 2", protovalidator.TagStringUTF8),
			FieldDesc:  "field 't_string_utf8'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagStringUTF8, Value: nil},
			FieldValue: func() interface{} { return data.TStringUtf8 },
			BeforeFunc: func() { data.TStringUtf8 = nil },
			AfterFunc:  func() { x := "中文abc"; data.TStringUtf8 = &x },
		},

		{
			Name:       fmt.Sprintf("test tag <%s> 2", protovalidator.TagStringAscii),
			FieldDesc:  "field 't_string_ascii'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagStringAscii, Value: nil},
			FieldValue: func() interface{} { return data.TStringAscii },
			BeforeFunc: func() { data.TStringAscii = nil },
			AfterFunc:  func() { x := "abc"; data.TStringAscii = &x },
		},
		{
			Name:       fmt.Sprintf("test tag <%s> 2", protovalidator.TagStringPrintAscii),
			FieldDesc:  "field 't_string_print_ascii'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagStringPrintAscii, Value: nil},
			FieldValue: func() interface{} { return data.TStringPrintAscii },
			BeforeFunc: func() { data.TStringPrintAscii = nil },
			AfterFunc:  func() { x := "abc"; data.TStringPrintAscii = &x },
		},
		{
			Name:       fmt.Sprintf("test tag <%s> 2", protovalidator.TagStringBoolean),
			FieldDesc:  "field 't_string_boolean'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagStringBoolean, Value: nil},
			FieldValue: func() interface{} { return data.TStringBoolean },
			BeforeFunc: func() { data.TStringBoolean = nil },
			AfterFunc:  func() { x := "true"; data.TStringBoolean = &x },
		},
		{
			Name:       fmt.Sprintf("test tag <%s> 2", protovalidator.TagStringLowercase),
			FieldDesc:  "field 't_string_lowercase'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagStringLowercase, Value: nil},
			FieldValue: func() interface{} { return data.TStringLowercase },
			BeforeFunc: func() { data.TStringLowercase = nil },
			AfterFunc:  func() { x := "abc"; data.TStringLowercase = &x },
		},
		{
			Name:       fmt.Sprintf("test tag <%s> 2", protovalidator.TagStringUppercase),
			FieldDesc:  "field 't_string_uppercase'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagStringUppercase, Value: nil},
			FieldValue: func() interface{} { return data.TStringUppercase },
			BeforeFunc: func() { data.TStringUppercase = nil },
			AfterFunc:  func() { x := "ABC"; data.TStringUppercase = &x },
		},
		{
			Name:       fmt.Sprintf("test tag <%s> 2", protovalidator.TagStringAlpha),
			FieldDesc:  "field 't_string_alpha'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagStringAlpha, Value: nil},
			FieldValue: func() interface{} { return data.TStringAlpha },
			BeforeFunc: func() { data.TStringAlpha = nil },
			AfterFunc:  func() { x := "abc"; data.TStringAlpha = &x },
		},
		{
			Name:       fmt.Sprintf("test tag <%s> 2", protovalidator.TagStringNumber),
			FieldDesc:  "field 't_string_number'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagStringNumber, Value: nil},
			FieldValue: func() interface{} { return data.TStringNumber },
			BeforeFunc: func() { data.TStringNumber = nil },
			AfterFunc:  func() { x := "123"; data.TStringNumber = &x },
		},
		{
			Name:       fmt.Sprintf("test tag <%s> 2", protovalidator.TagStringAlphaNumber),
			FieldDesc:  "field 't_string_alpha_number'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagStringAlphaNumber, Value: nil},
			FieldValue: func() interface{} { return data.TStringAlphaNumber },
			BeforeFunc: func() { data.TStringAlphaNumber = nil },
			AfterFunc:  func() { x := "abc123"; data.TStringAlphaNumber = &x },
		},
	}

	caseNetwork1 := []*CaseDesc{
		{
			Name:       fmt.Sprintf("test tag <%s> 1", protovalidator.TagStringIp),
			FieldDesc:  "field 't_string_ip'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagStringIp, Value: nil},
			FieldValue: func() interface{} { return data.TStringIp },
			BeforeFunc: func() { x := "127.0."; data.TStringIp = &x },
			AfterFunc:  func() { x := "127.0.0.1"; data.TStringIp = &x },
		},
		{
			Name:       fmt.Sprintf("test tag <%s> 2", protovalidator.TagStringIp),
			FieldDesc:  "field 't_string_ip'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagStringIp, Value: nil},
			FieldValue: func() interface{} { return data.TStringIp },
			BeforeFunc: func() { x := "AD80:0000:0000:0000:ABAA:0000:00C2:0002:xxxxx"; data.TStringIp = &x },
			AfterFunc:  func() { x := "AD80:0000:0000:0000:ABAA:0000:00C2:0002"; data.TStringIp = &x },
		},
		{
			Name:       fmt.Sprintf("test tag <%s> 1", protovalidator.TagStringIpv4),
			FieldDesc:  "field 't_string_ipv4'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagStringIpv4, Value: nil},
			FieldValue: func() interface{} { return data.TStringIpv4 },
			BeforeFunc: func() { x := "127.0.0.1.1"; data.TStringIpv4 = &x },
			AfterFunc:  func() { x := "127.0.0.1"; data.TStringIpv4 = &x },
		},
		{
			Name:       fmt.Sprintf("test tag <%s> 1", protovalidator.TagStringIpv6),
			FieldDesc:  "field 't_string_ipv6'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagStringIpv6, Value: nil},
			FieldValue: func() interface{} { return data.TStringIpv6 },
			BeforeFunc: func() { x := "127.0.0.1"; data.TStringIpv6 = &x },
			AfterFunc:  func() { x := "AD80:0000:0000:0000:ABAA:0000:00C2:0002"; data.TStringIpv6 = &x },
		},
		{
			Name:       fmt.Sprintf("test tag <%s> 1", protovalidator.TagStringIpAddr),
			FieldDesc:  "field 't_string_ip_addr'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagStringIpAddr, Value: nil},
			FieldValue: func() interface{} { return data.TStringIpAddr },
			BeforeFunc: func() { x := "127.0.0.1.x"; data.TStringIpAddr = &x },
			AfterFunc:  func() { x := "127.0.0.1"; data.TStringIpAddr = &x },
		},
		{
			Name:       fmt.Sprintf("test tag <%s> 1", protovalidator.TagStringIp4Addr),
			FieldDesc:  "field 't_string_ip4_addr'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagStringIp4Addr, Value: nil},
			FieldValue: func() interface{} { return data.TStringIp4Addr },
			BeforeFunc: func() { x := "127.0.0.1.x"; data.TStringIp4Addr = &x },
			AfterFunc:  func() { x := "127.0.0.1"; data.TStringIp4Addr = &x },
		},
		{
			Name:       fmt.Sprintf("test tag <%s> 1", protovalidator.TagStringIp6Addr),
			FieldDesc:  "field 't_string_ip6_addr'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagStringIp6Addr, Value: nil},
			FieldValue: func() interface{} { return data.TStringIp6Addr },
			BeforeFunc: func() { x := "127.0.0.1"; data.TStringIp6Addr = &x },
			AfterFunc:  func() { x := "AD80:0000:0000:0000:ABAA:0000:00C2:0002"; data.TStringIp6Addr = &x },
		},
		{
			Name:       fmt.Sprintf("test tag <%s> 1", protovalidator.TagStringCidr),
			FieldDesc:  "field 't_string_cidr'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagStringCidr, Value: nil},
			FieldValue: func() interface{} { return data.TStringCidr },
			BeforeFunc: func() { x := "192.0.2.0/128"; data.TStringCidr = &x },
			AfterFunc:  func() { x := "192.0.2.0/24"; data.TStringCidr = &x },
		},
		{
			Name:       fmt.Sprintf("test tag <%s> 1", protovalidator.TagStringCidrv4),
			FieldDesc:  "field 't_string_cidrv4'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagStringCidrv4, Value: nil},
			FieldValue: func() interface{} { return data.TStringCidrv4 },
			BeforeFunc: func() { x := "192.0.2.0/243"; data.TStringCidrv4 = &x },
			AfterFunc:  func() { x := "192.0.2.0/24"; data.TStringCidrv4 = &x },
		},
		{
			Name:       fmt.Sprintf("test tag <%s> 1", protovalidator.TagStringCidrv6),
			FieldDesc:  "field 't_string_cidrv6'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagStringCidrv6, Value: nil},
			FieldValue: func() interface{} { return data.TStringCidrv6 },
			BeforeFunc: func() { x := "2001:db8::/129"; data.TStringCidrv6 = &x },
			AfterFunc:  func() { x := "2001:db8::/32"; data.TStringCidrv6 = &x },
		},
		{
			Name:       fmt.Sprintf("test tag <%s> 1", protovalidator.TagStringMac),
			FieldDesc:  "field 't_string_mac'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagStringMac, Value: nil},
			FieldValue: func() interface{} { return data.TStringMac },
			BeforeFunc: func() { x := "00:00:5e:00:53:01:dsafadsfd"; data.TStringMac = &x },
			AfterFunc:  func() { x := "00:00:5e:00:53:01"; data.TStringMac = &x },
		},
		{
			Name:       fmt.Sprintf("test tag <%s> 1", protovalidator.TagStringTcpAddr),
			FieldDesc:  "field 't_string_tcp_addr'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagStringTcpAddr, Value: nil},
			FieldValue: func() interface{} { return data.TStringTcpAddr },
			BeforeFunc: func() { x := "127.0."; data.TStringTcpAddr = &x },
			AfterFunc:  func() { x := "127.0.0.1:80"; data.TStringTcpAddr = &x },
		},
		{
			Name:       fmt.Sprintf("test tag <%s> 1", protovalidator.TagStringTcp4Addr),
			FieldDesc:  "field 't_string_tcp4_addr'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagStringTcp4Addr, Value: nil},
			FieldValue: func() interface{} { return data.TStringTcp4Addr },
			BeforeFunc: func() { x := ""; data.TStringTcp4Addr = &x },
			AfterFunc:  func() { x := "127.0.0.1:80"; data.TStringTcp4Addr = &x },
		},
		{
			Name:       fmt.Sprintf("test tag <%s> 1", protovalidator.TagStringTcp6Addr),
			FieldDesc:  "field 't_string_tcp6_addr'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagStringTcp6Addr, Value: nil},
			FieldValue: func() interface{} { return data.TStringTcp6Addr },
			BeforeFunc: func() { x := ""; data.TStringTcp6Addr = &x },
			AfterFunc:  func() { x := "[2001:db8::1]:http"; data.TStringTcp6Addr = &x },
		},
		{
			Name:       fmt.Sprintf("test tag <%s> 1", protovalidator.TagStringUdpAddr),
			FieldDesc:  "field 't_string_udp_addr'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagStringUdpAddr, Value: nil},
			FieldValue: func() interface{} { return data.TStringUdpAddr },
			BeforeFunc: func() { x := ""; data.TStringUdpAddr = &x },
			AfterFunc:  func() { x := "127.0.0.1:domain"; data.TStringUdpAddr = &x },
		},
		{
			Name:       fmt.Sprintf("test tag <%s> 1", protovalidator.TagStringUdp4Addr),
			FieldDesc:  "field 't_string_udp4_addr'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagStringUdp4Addr, Value: nil},
			FieldValue: func() interface{} { return data.TStringUdp4Addr },
			BeforeFunc: func() { x := ""; data.TStringUdp4Addr = &x },
			AfterFunc:  func() { x := "127.0.0.1:domain"; data.TStringUdp4Addr = &x },
		},
		{
			Name:       fmt.Sprintf("test tag <%s> 1", protovalidator.TagStringUdp6Addr),
			FieldDesc:  "field 't_string_udp6_addr'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagStringUdp6Addr, Value: nil},
			FieldValue: func() interface{} { return data.TStringUdp6Addr },
			BeforeFunc: func() { x := ""; data.TStringUdp6Addr = &x },
			AfterFunc:  func() { x := "[2001:db8::1]:domain"; data.TStringUdp6Addr = &x },
		},
		{
			Name:       fmt.Sprintf("test tag <%s> 1", protovalidator.TagStringUnixAddr),
			FieldDesc:  "field 't_string_unix_addr'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagStringUnixAddr, Value: nil},
			FieldValue: func() interface{} { return data.TStringUnixAddr },
			BeforeFunc: func() { x := ""; data.TStringUnixAddr = &x },
			AfterFunc:  func() { x := "unix:///mysql/mysql.sock"; data.TStringUnixAddr = &x },
		},
		{
			Name:       fmt.Sprintf("test tag <%s> 1", protovalidator.TagStringHostname),
			FieldDesc:  "field 't_string_hostname'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagStringHostname, Value: nil},
			FieldValue: func() interface{} { return data.TStringHostname },
			BeforeFunc: func() { x := ""; data.TStringHostname = &x },
			AfterFunc:  func() { x := "www.google.com"; data.TStringHostname = &x },
		},
		{
			Name:       fmt.Sprintf("test tag <%s> 1", protovalidator.TagStringHostnameRfc1123),
			FieldDesc:  "field 't_string_hostname_rfc1123'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagStringHostnameRfc1123, Value: nil},
			FieldValue: func() interface{} { return data.TStringHostnameRfc1123 },
			BeforeFunc: func() { x := ""; data.TStringHostnameRfc1123 = &x },
			AfterFunc:  func() { x := "www.google.com"; data.TStringHostnameRfc1123 = &x },
		},
		{
			Name:       fmt.Sprintf("test tag <%s> 1", protovalidator.TagStringHostnamePort),
			FieldDesc:  "field 't_string_hostname_port'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagStringHostnamePort, Value: nil},
			FieldValue: func() interface{} { return data.TStringHostnamePort },
			BeforeFunc: func() { x := ""; data.TStringHostnamePort = &x },
			AfterFunc:  func() { x := "www.google.com:80"; data.TStringHostnamePort = &x },
		},
		{
			Name:       fmt.Sprintf("test tag <%s> 1", protovalidator.TagStringDataURI),
			FieldDesc:  "field 't_string_data_uri'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagStringDataURI, Value: nil},
			FieldValue: func() interface{} { return data.TStringDataUri },
			BeforeFunc: func() { x := ""; data.TStringDataUri = &x },
			AfterFunc:  func() { x := "data:text/html,WkhPTkdHVU9uaWhhbzEyMw=="; data.TStringDataUri = &x },
		},
		{
			Name:       fmt.Sprintf("test tag <%s> 1", protovalidator.TagStringFQDN),
			FieldDesc:  "field 't_string_fqdn'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagStringFQDN, Value: nil},
			FieldValue: func() interface{} { return data.TStringFqdn },
			BeforeFunc: func() { x := ""; data.TStringFqdn = &x },
			AfterFunc:  func() { x := "MacBook-Pro.local"; data.TStringFqdn = &x },
		},
		{
			Name:       fmt.Sprintf("test tag <%s> 1", protovalidator.TagStringURI),
			FieldDesc:  "field 't_string_uri'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagStringURI, Value: nil},
			FieldValue: func() interface{} { return data.TStringUri },
			BeforeFunc: func() { x := ""; data.TStringUri = &x },
			AfterFunc:  func() { x := "/v1/api"; data.TStringUri = &x },
		},
		{
			Name:       fmt.Sprintf("test tag <%s> 1", protovalidator.TagStringURL),
			FieldDesc:  "field 't_string_url'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagStringURL, Value: nil},
			FieldValue: func() interface{} { return data.TStringUrl },
			BeforeFunc: func() { x := ""; data.TStringUrl = &x },
			AfterFunc:  func() { x := "https://www.google.com/v1/api"; data.TStringUrl = &x },
		},
		{
			Name:       fmt.Sprintf("test tag <%s> 1", protovalidator.TagStringURLEncoded),
			FieldDesc:  "field 't_string_url_encoded'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagStringURLEncoded, Value: nil},
			FieldValue: func() interface{} { return data.TStringUrlEncoded },
			BeforeFunc: func() { x := ""; data.TStringUrlEncoded = &x },
			AfterFunc:  func() { x := "https://www.google.com%2Fv1%2Fapi"; data.TStringUrlEncoded = &x },
		},
	}

	caseNetwork1NilPoint := []*CaseDesc{
		{
			Name:       fmt.Sprintf("test tag <%s> 2", protovalidator.TagStringIp),
			FieldDesc:  "field 't_string_ip'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagStringIp, Value: nil},
			FieldValue: func() interface{} { return data.TStringIp },
			BeforeFunc: func() { data.TStringIp = nil },
			AfterFunc:  func() { x := "127.0.0.1"; data.TStringIp = &x },
		},
		{
			Name:       fmt.Sprintf("test tag <%s> 2", protovalidator.TagStringIp),
			FieldDesc:  "field 't_string_ip'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagStringIp, Value: nil},
			FieldValue: func() interface{} { return data.TStringIp },
			BeforeFunc: func() { data.TStringIp = nil },
			AfterFunc:  func() { x := "AD80:0000:0000:0000:ABAA:0000:00C2:0002"; data.TStringIp = &x },
		},
		{
			Name:       fmt.Sprintf("test tag <%s> 2", protovalidator.TagStringIpv4),
			FieldDesc:  "field 't_string_ipv4'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagStringIpv4, Value: nil},
			FieldValue: func() interface{} { return data.TStringIpv4 },
			BeforeFunc: func() { data.TStringIpv4 = nil },
			AfterFunc:  func() { x := "127.0.0.1"; data.TStringIpv4 = &x },
		},
		{
			Name:       fmt.Sprintf("test tag <%s> 2", protovalidator.TagStringIpv6),
			FieldDesc:  "field 't_string_ipv6'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagStringIpv6, Value: nil},
			FieldValue: func() interface{} { return data.TStringIpv6 },
			BeforeFunc: func() { data.TStringIpv6 = nil },
			AfterFunc:  func() { x := "AD80:0000:0000:0000:ABAA:0000:00C2:0002"; data.TStringIpv6 = &x },
		},
		{
			Name:       fmt.Sprintf("test tag <%s> 2", protovalidator.TagStringIpAddr),
			FieldDesc:  "field 't_string_ip_addr'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagStringIpAddr, Value: nil},
			FieldValue: func() interface{} { return data.TStringIpAddr },
			BeforeFunc: func() { data.TStringIpAddr = nil },
			AfterFunc:  func() { x := "127.0.0.1"; data.TStringIpAddr = &x },
		},
		{
			Name:       fmt.Sprintf("test tag <%s> 2", protovalidator.TagStringIp4Addr),
			FieldDesc:  "field 't_string_ip4_addr'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagStringIp4Addr, Value: nil},
			FieldValue: func() interface{} { return data.TStringIp4Addr },
			BeforeFunc: func() { data.TStringIp4Addr = nil },
			AfterFunc:  func() { x := "127.0.0.1"; data.TStringIp4Addr = &x },
		},
		{
			Name:       fmt.Sprintf("test tag <%s> 2", protovalidator.TagStringIp6Addr),
			FieldDesc:  "field 't_string_ip6_addr'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagStringIp6Addr, Value: nil},
			FieldValue: func() interface{} { return data.TStringIp6Addr },
			BeforeFunc: func() { data.TStringIp6Addr = nil },
			AfterFunc:  func() { x := "AD80:0000:0000:0000:ABAA:0000:00C2:0002"; data.TStringIp6Addr = &x },
		},
		{
			Name:       fmt.Sprintf("test tag <%s> 2", protovalidator.TagStringCidr),
			FieldDesc:  "field 't_string_cidr'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagStringCidr, Value: nil},
			FieldValue: func() interface{} { return data.TStringCidr },
			BeforeFunc: func() { data.TStringCidr = nil },
			AfterFunc:  func() { x := "192.0.2.0/24"; data.TStringCidr = &x },
		},
		{
			Name:       fmt.Sprintf("test tag <%s> 2", protovalidator.TagStringCidrv4),
			FieldDesc:  "field 't_string_cidrv4'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagStringCidrv4, Value: nil},
			FieldValue: func() interface{} { return data.TStringCidrv4 },
			BeforeFunc: func() { data.TStringCidrv4 = nil },
			AfterFunc:  func() { x := "192.0.2.0/24"; data.TStringCidrv4 = &x },
		},
		{
			Name:       fmt.Sprintf("test tag <%s> 2", protovalidator.TagStringCidrv6),
			FieldDesc:  "field 't_string_cidrv6'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagStringCidrv6, Value: nil},
			FieldValue: func() interface{} { return data.TStringCidrv6 },
			BeforeFunc: func() { data.TStringCidrv6 = nil },
			AfterFunc:  func() { x := "2001:db8::/32"; data.TStringCidrv6 = &x },
		},
		{
			Name:       fmt.Sprintf("test tag <%s> 2", protovalidator.TagStringMac),
			FieldDesc:  "field 't_string_mac'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagStringMac, Value: nil},
			FieldValue: func() interface{} { return data.TStringMac },
			BeforeFunc: func() { data.TStringMac = nil },
			AfterFunc:  func() { x := "00:00:5e:00:53:01"; data.TStringMac = &x },
		},
		{
			Name:       fmt.Sprintf("test tag <%s> 2", protovalidator.TagStringTcpAddr),
			FieldDesc:  "field 't_string_tcp_addr'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagStringTcpAddr, Value: nil},
			FieldValue: func() interface{} { return data.TStringTcpAddr },
			BeforeFunc: func() { data.TStringTcpAddr = nil },
			AfterFunc:  func() { x := "127.0.0.1:80"; data.TStringTcpAddr = &x },
		},
		{
			Name:       fmt.Sprintf("test tag <%s> 2", protovalidator.TagStringTcp4Addr),
			FieldDesc:  "field 't_string_tcp4_addr'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagStringTcp4Addr, Value: nil},
			FieldValue: func() interface{} { return data.TStringTcp4Addr },
			BeforeFunc: func() { data.TStringTcp4Addr = nil },
			AfterFunc:  func() { x := "127.0.0.1:80"; data.TStringTcp4Addr = &x },
		},
		{
			Name:       fmt.Sprintf("test tag <%s> 2", protovalidator.TagStringTcp6Addr),
			FieldDesc:  "field 't_string_tcp6_addr'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagStringTcp6Addr, Value: nil},
			FieldValue: func() interface{} { return data.TStringTcp6Addr },
			BeforeFunc: func() { data.TStringTcp6Addr = nil },
			AfterFunc:  func() { x := "[2001:db8::1]:http"; data.TStringTcp6Addr = &x },
		},
		{
			Name:       fmt.Sprintf("test tag <%s> 2", protovalidator.TagStringUdpAddr),
			FieldDesc:  "field 't_string_udp_addr'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagStringUdpAddr, Value: nil},
			FieldValue: func() interface{} { return data.TStringUdpAddr },
			BeforeFunc: func() { data.TStringUdpAddr = nil },
			AfterFunc:  func() { x := "127.0.0.1:domain"; data.TStringUdpAddr = &x },
		},
		{
			Name:       fmt.Sprintf("test tag <%s> 2", protovalidator.TagStringUdp4Addr),
			FieldDesc:  "field 't_string_udp4_addr'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagStringUdp4Addr, Value: nil},
			FieldValue: func() interface{} { return data.TStringUdp4Addr },
			BeforeFunc: func() { data.TStringUdp4Addr = nil },
			AfterFunc:  func() { x := "127.0.0.1:domain"; data.TStringUdp4Addr = &x },
		},
		{
			Name:       fmt.Sprintf("test tag <%s> 2", protovalidator.TagStringUdp6Addr),
			FieldDesc:  "field 't_string_udp6_addr'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagStringUdp6Addr, Value: nil},
			FieldValue: func() interface{} { return data.TStringUdp6Addr },
			BeforeFunc: func() { data.TStringUdp6Addr = nil },
			AfterFunc:  func() { x := "[2001:db8::1]:domain"; data.TStringUdp6Addr = &x },
		},
		{
			Name:       fmt.Sprintf("test tag <%s> 2", protovalidator.TagStringUnixAddr),
			FieldDesc:  "field 't_string_unix_addr'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagStringUnixAddr, Value: nil},
			FieldValue: func() interface{} { return data.TStringUnixAddr },
			BeforeFunc: func() { data.TStringUnixAddr = nil },
			AfterFunc:  func() { x := "unix:///mysql/mysql.sock"; data.TStringUnixAddr = &x },
		},
		{
			Name:       fmt.Sprintf("test tag <%s> 2", protovalidator.TagStringHostname),
			FieldDesc:  "field 't_string_hostname'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagStringHostname, Value: nil},
			FieldValue: func() interface{} { return data.TStringHostname },
			BeforeFunc: func() { data.TStringHostname = nil },
			AfterFunc:  func() { x := "www.google.com"; data.TStringHostname = &x },
		},
		{
			Name:       fmt.Sprintf("test tag <%s> 2", protovalidator.TagStringHostnameRfc1123),
			FieldDesc:  "field 't_string_hostname_rfc1123'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagStringHostnameRfc1123, Value: nil},
			FieldValue: func() interface{} { return data.TStringHostnameRfc1123 },
			BeforeFunc: func() { data.TStringHostnameRfc1123 = nil },
			AfterFunc:  func() { x := "www.google.com"; data.TStringHostnameRfc1123 = &x },
		},
		{
			Name:       fmt.Sprintf("test tag <%s> 2", protovalidator.TagStringHostnamePort),
			FieldDesc:  "field 't_string_hostname_port'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagStringHostnamePort, Value: nil},
			FieldValue: func() interface{} { return data.TStringHostnamePort },
			BeforeFunc: func() { data.TStringHostnamePort = nil },
			AfterFunc:  func() { x := "www.google.com:80"; data.TStringHostnamePort = &x },
		},
		{
			Name:       fmt.Sprintf("test tag <%s> 2", protovalidator.TagStringDataURI),
			FieldDesc:  "field 't_string_data_uri'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagStringDataURI, Value: nil},
			FieldValue: func() interface{} { return data.TStringDataUri },
			BeforeFunc: func() { data.TStringDataUri = nil },
			AfterFunc:  func() { x := "data:text/html,WkhPTkdHVU9uaWhhbzEyMw=="; data.TStringDataUri = &x },
		},
		{
			Name:       fmt.Sprintf("test tag <%s> 2", protovalidator.TagStringFQDN),
			FieldDesc:  "field 't_string_fqdn'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagStringFQDN, Value: nil},
			FieldValue: func() interface{} { return data.TStringFqdn },
			BeforeFunc: func() { data.TStringFqdn = nil },
			AfterFunc:  func() { x := "MacBook-Pro.local"; data.TStringFqdn = &x },
		},
		{
			Name:       fmt.Sprintf("test tag <%s> 2", protovalidator.TagStringURI),
			FieldDesc:  "field 't_string_uri'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagStringURI, Value: nil},
			FieldValue: func() interface{} { return data.TStringUri },
			BeforeFunc: func() { data.TStringUri = nil },
			AfterFunc:  func() { x := "/v1/api"; data.TStringUri = &x },
		},
		{
			Name:       fmt.Sprintf("test tag <%s> 2", protovalidator.TagStringURL),
			FieldDesc:  "field 't_string_url'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagStringURL, Value: nil},
			FieldValue: func() interface{} { return data.TStringUrl },
			BeforeFunc: func() { data.TStringUrl = nil },
			AfterFunc:  func() { x := "https://www.google.com/v1/api"; data.TStringUrl = &x },
		},
		{
			Name:       fmt.Sprintf("test tag <%s> 2", protovalidator.TagStringURLEncoded),
			FieldDesc:  "field 't_string_url_encoded'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagStringURLEncoded, Value: nil},
			FieldValue: func() interface{} { return data.TStringUrlEncoded },
			BeforeFunc: func() { data.TStringUrlEncoded = nil },
			AfterFunc:  func() { x := "https://www.google.com%2Fv1%2Fapi"; data.TStringUrlEncoded = &x },
		},
	}

	caseFormat1 := []*CaseDesc{
		{
			Name:       fmt.Sprintf("test tag <%s> 1", protovalidator.TagStringUnixCron),
			FieldDesc:  "field 't_string_unix_cron'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagStringUnixCron, Value: nil},
			FieldValue: func() interface{} { return data.TStringUnixCron },
			BeforeFunc: func() { x := ""; data.TStringUnixCron = &x },
			AfterFunc:  func() { x := "* * * * *"; data.TStringUnixCron = &x },
		},
		{
			Name:       fmt.Sprintf("test tag <%s> 1", protovalidator.TagStringEmail),
			FieldDesc:  "field 't_string_email'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagStringEmail, Value: nil},
			FieldValue: func() interface{} { return data.TStringEmail },
			BeforeFunc: func() { x := ""; data.TStringEmail = &x },
			AfterFunc:  func() { x := "xxx@gmail.com"; data.TStringEmail = &x },
		},
		{
			Name:       fmt.Sprintf("test tag <%s> 1", protovalidator.TagStringJSON),
			FieldDesc:  "field 't_string_json'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagStringJSON, Value: nil},
			FieldValue: func() interface{} { return data.TStringJson },
			BeforeFunc: func() { x := ""; data.TStringJson = &x },
			AfterFunc:  func() { x := `{"k1":"v1"}`; data.TStringJson = &x },
		},
		{
			Name:       fmt.Sprintf("test tag <%s> 1", protovalidator.TagStringJWT),
			FieldDesc:  "field 't_string_jwt'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagStringJWT, Value: nil},
			FieldValue: func() interface{} { return data.TStringJwt },
			BeforeFunc: func() { x := ""; data.TStringJwt = &x },
			AfterFunc: func() {
				x := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiYWRtaW4iOnRydWV9.TJVA95OrM7E2cBab30RMHrHDcEfxjoYZgeFONFh7HgQ"
				data.TStringJwt = &x
			},
		},
		{
			Name:       fmt.Sprintf("test tag <%s> 1", protovalidator.TagStringHTML),
			FieldDesc:  "field 't_string_html'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagStringHTML, Value: nil},
			FieldValue: func() interface{} { return data.TStringHtml },
			BeforeFunc: func() { x := ""; data.TStringHtml = &x },
			AfterFunc:  func() { x := "<html></html>"; data.TStringHtml = &x },
		},
		{
			Name:       fmt.Sprintf("test tag <%s> 1", protovalidator.TagStringHTMLEncoded),
			FieldDesc:  "field 't_string_html_encoded'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagStringHTMLEncoded, Value: nil},
			FieldValue: func() interface{} { return data.TStringHtmlEncoded },
			BeforeFunc: func() { x := ""; data.TStringHtmlEncoded = &x },
			AfterFunc:  func() { x := "&lt;html&gt;&lt;/html&gt;"; data.TStringHtmlEncoded = &x },
		},
		{
			Name:       fmt.Sprintf("test tag <%s> 1", protovalidator.TagStringBase64),
			FieldDesc:  "field 't_string_base64'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagStringBase64, Value: nil},
			FieldValue: func() interface{} { return data.TStringBase64 },
			BeforeFunc: func() { x := ""; data.TStringBase64 = &x },
			AfterFunc:  func() { x := "WkhPTkdHVU9uaWhhbzEyMw=="; data.TStringBase64 = &x },
		},
		{
			Name:       fmt.Sprintf("test tag <%s> 1", protovalidator.TagStringBase64URL),
			FieldDesc:  "field 't_string_base64_url'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagStringBase64URL, Value: nil},
			FieldValue: func() interface{} { return data.TStringBase64Url },
			BeforeFunc: func() { x := ""; data.TStringBase64Url = &x },
			AfterFunc:  func() { x := "YWRmYWRzZg=="; data.TStringBase64Url = &x },
		},
		{
			Name:       fmt.Sprintf("test tag <%s> 1", protovalidator.TagStringHexadecimal),
			FieldDesc:  "field 't_string_hexadecimal'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagStringHexadecimal, Value: nil},
			FieldValue: func() interface{} { return data.TStringHexadecimal },
			BeforeFunc: func() { x := ""; data.TStringHexadecimal = &x },
			AfterFunc:  func() { x := "6461666164736661647366736461"; data.TStringHexadecimal = &x },
		},
		{
			Name:       fmt.Sprintf("test tag <%s> 1", protovalidator.TagStringDatetime),
			FieldDesc:  "field 't_string_datetime'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagStringDatetime, Value: "2006-01-02 15:04:05"},
			FieldValue: func() interface{} { return data.TStringDatetime },
			BeforeFunc: func() { x := ""; data.TStringDatetime = &x },
			AfterFunc:  func() { x := "2020-11-02 15:04:05"; data.TStringDatetime = &x },
		},
		{
			Name:       fmt.Sprintf("test tag <%s> 1", protovalidator.TagStringTimezone),
			FieldDesc:  "field 't_string_timezone'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagStringTimezone, Value: nil},
			FieldValue: func() interface{} { return data.TStringTimezone },
			BeforeFunc: func() { x := ""; data.TStringTimezone = &x },
			AfterFunc:  func() { x := "UTC"; data.TStringTimezone = &x },
		},
		{
			Name:       fmt.Sprintf("test tag <%s> 1", protovalidator.TagStringUUID),
			FieldDesc:  "field 't_string_uuid'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagStringUUID, Value: nil},
			FieldValue: func() interface{} { return data.TStringUuid },
			BeforeFunc: func() { x := ""; data.TStringUuid = &x },
			AfterFunc:  func() { x := "80363646-7361-11ec-b4f0-020017000b7b"; data.TStringUuid = &x },
		},
		{
			Name:       fmt.Sprintf("test tag <%s> 1", protovalidator.TagStringUUID1),
			FieldDesc:  "field 't_string_uuid1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagStringUUID1, Value: nil},
			FieldValue: func() interface{} { return data.TStringUuid1 },
			BeforeFunc: func() { x := ""; data.TStringUuid1 = &x },
			AfterFunc:  func() { x := "80363646-7361-11ec-b4f0-020017000b7b"; data.TStringUuid1 = &x },
		},
		{
			Name:       fmt.Sprintf("test tag <%s> 1", protovalidator.TagStringUUID3),
			FieldDesc:  "field 't_string_uuid3'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagStringUUID3, Value: nil},
			FieldValue: func() interface{} { return data.TStringUuid3 },
			BeforeFunc: func() { x := ""; data.TStringUuid3 = &x },
			AfterFunc:  func() { x := "fa1f0699-6eac-3c47-b0f1-9d67ce5d0b76"; data.TStringUuid3 = &x },
		},
		{
			Name:       fmt.Sprintf("test tag <%s> 1", protovalidator.TagStringUUID4),
			FieldDesc:  "field 't_string_uuid4'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagStringUUID4, Value: nil},
			FieldValue: func() interface{} { return data.TStringUuid4 },
			BeforeFunc: func() { x := ""; data.TStringUuid4 = &x },
			AfterFunc:  func() { x := "3af73fea-62c7-4a00-8314-942a1841288a"; data.TStringUuid4 = &x },
		},
		{
			Name:       fmt.Sprintf("test tag <%s> 1", protovalidator.TagStringUUID5),
			FieldDesc:  "field 't_string_uuid5'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagStringUUID5, Value: nil},
			FieldValue: func() interface{} { return data.TStringUuid5 },
			BeforeFunc: func() { x := ""; data.TStringUuid5 = &x },
			AfterFunc:  func() { x := "78381e24-2d01-5a28-b2bf-9eedab6605ed"; data.TStringUuid5 = &x },
		},
	}
	caseFormat1NilPoint := []*CaseDesc{
		{
			Name:       fmt.Sprintf("test tag <%s> 2", protovalidator.TagStringUnixCron),
			FieldDesc:  "field 't_string_unix_cron'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagStringUnixCron, Value: nil},
			FieldValue: func() interface{} { return data.TStringUnixCron },
			BeforeFunc: func() { data.TStringUnixCron = nil },
			AfterFunc:  func() { x := "* * * * *"; data.TStringUnixCron = &x },
		},
		{
			Name:       fmt.Sprintf("test tag <%s> 2", protovalidator.TagStringEmail),
			FieldDesc:  "field 't_string_email'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagStringEmail, Value: nil},
			FieldValue: func() interface{} { return data.TStringEmail },
			BeforeFunc: func() { data.TStringEmail = nil },
			AfterFunc:  func() { x := "xxx@gmail.com"; data.TStringEmail = &x },
		},
		{
			Name:       fmt.Sprintf("test tag <%s> 2", protovalidator.TagStringJSON),
			FieldDesc:  "field 't_string_json'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagStringJSON, Value: nil},
			FieldValue: func() interface{} { return data.TStringJson },
			BeforeFunc: func() { data.TStringJson = nil },
			AfterFunc:  func() { x := `{"k1":"v1"}`; data.TStringJson = &x },
		},
		{
			Name:       fmt.Sprintf("test tag <%s> 2", protovalidator.TagStringJWT),
			FieldDesc:  "field 't_string_jwt'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagStringJWT, Value: nil},
			FieldValue: func() interface{} { return data.TStringJwt },
			BeforeFunc: func() { data.TStringJwt = nil },
			AfterFunc: func() {
				x := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiYWRtaW4iOnRydWV9.TJVA95OrM7E2cBab30RMHrHDcEfxjoYZgeFONFh7HgQ"
				data.TStringJwt = &x
			},
		},
		{
			Name:       fmt.Sprintf("test tag <%s> 2", protovalidator.TagStringHTML),
			FieldDesc:  "field 't_string_html'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagStringHTML, Value: nil},
			FieldValue: func() interface{} { return data.TStringHtml },
			BeforeFunc: func() { data.TStringHtml = nil },
			AfterFunc:  func() { x := "<html></html>"; data.TStringHtml = &x },
		},
		{
			Name:       fmt.Sprintf("test tag <%s> 2", protovalidator.TagStringHTMLEncoded),
			FieldDesc:  "field 't_string_html_encoded'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagStringHTMLEncoded, Value: nil},
			FieldValue: func() interface{} { return data.TStringHtmlEncoded },
			BeforeFunc: func() { data.TStringHtmlEncoded = nil },
			AfterFunc:  func() { x := "&lt;html&gt;&lt;/html&gt;"; data.TStringHtmlEncoded = &x },
		},
		{
			Name:       fmt.Sprintf("test tag <%s> 2", protovalidator.TagStringBase64),
			FieldDesc:  "field 't_string_base64'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagStringBase64, Value: nil},
			FieldValue: func() interface{} { return data.TStringBase64 },
			BeforeFunc: func() { data.TStringBase64 = nil },
			AfterFunc:  func() { x := "WkhPTkdHVU9uaWhhbzEyMw=="; data.TStringBase64 = &x },
		},
		{
			Name:       fmt.Sprintf("test tag <%s> 2", protovalidator.TagStringBase64URL),
			FieldDesc:  "field 't_string_base64_url'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagStringBase64URL, Value: nil},
			FieldValue: func() interface{} { return data.TStringBase64Url },
			BeforeFunc: func() { data.TStringBase64Url = nil },
			AfterFunc:  func() { x := "YWRmYWRzZg=="; data.TStringBase64Url = &x },
		},
		{
			Name:       fmt.Sprintf("test tag <%s> 2", protovalidator.TagStringHexadecimal),
			FieldDesc:  "field 't_string_hexadecimal'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagStringHexadecimal, Value: nil},
			FieldValue: func() interface{} { return data.TStringHexadecimal },
			BeforeFunc: func() { data.TStringHexadecimal = nil },
			AfterFunc:  func() { x := "6461666164736661647366736461"; data.TStringHexadecimal = &x },
		},
		{
			Name:       fmt.Sprintf("test tag <%s> 2", protovalidator.TagStringDatetime),
			FieldDesc:  "field 't_string_datetime'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagStringDatetime, Value: "2006-01-02 15:04:05"},
			FieldValue: func() interface{} { return data.TStringDatetime },
			BeforeFunc: func() { data.TStringDatetime = nil },
			AfterFunc:  func() { x := "2020-11-02 15:04:05"; data.TStringDatetime = &x },
		},
		{
			Name:       fmt.Sprintf("test tag <%s> 2", protovalidator.TagStringTimezone),
			FieldDesc:  "field 't_string_timezone'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagStringTimezone, Value: nil},
			FieldValue: func() interface{} { return data.TStringTimezone },
			BeforeFunc: func() { data.TStringTimezone = nil },
			AfterFunc:  func() { x := "UTC"; data.TStringTimezone = &x },
		},
		{
			Name:       fmt.Sprintf("test tag <%s> 2", protovalidator.TagStringUUID),
			FieldDesc:  "field 't_string_uuid'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagStringUUID, Value: nil},
			FieldValue: func() interface{} { return data.TStringUuid },
			BeforeFunc: func() { data.TStringUuid = nil },
			AfterFunc:  func() { x := "80363646-7361-11ec-b4f0-020017000b7b"; data.TStringUuid = &x },
		},
		{
			Name:       fmt.Sprintf("test tag <%s> 2", protovalidator.TagStringUUID1),
			FieldDesc:  "field 't_string_uuid1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagStringUUID1, Value: nil},
			FieldValue: func() interface{} { return data.TStringUuid1 },
			BeforeFunc: func() { data.TStringUuid1 = nil },
			AfterFunc:  func() { x := "80363646-7361-11ec-b4f0-020017000b7b"; data.TStringUuid1 = &x },
		},
		{
			Name:       fmt.Sprintf("test tag <%s> 2", protovalidator.TagStringUUID3),
			FieldDesc:  "field 't_string_uuid3'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagStringUUID3, Value: nil},
			FieldValue: func() interface{} { return data.TStringUuid3 },
			BeforeFunc: func() { data.TStringUuid3 = nil },
			AfterFunc:  func() { x := "fa1f0699-6eac-3c47-b0f1-9d67ce5d0b76"; data.TStringUuid3 = &x },
		},
		{
			Name:       fmt.Sprintf("test tag <%s> 2", protovalidator.TagStringUUID4),
			FieldDesc:  "field 't_string_uuid4'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagStringUUID4, Value: nil},
			FieldValue: func() interface{} { return data.TStringUuid4 },
			BeforeFunc: func() { data.TStringUuid4 = nil },
			AfterFunc:  func() { x := "3af73fea-62c7-4a00-8314-942a1841288a"; data.TStringUuid4 = &x },
		},
		{
			Name:       fmt.Sprintf("test tag <%s> 2", protovalidator.TagStringUUID5),
			FieldDesc:  "field 't_string_uuid5'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagStringUUID5, Value: nil},
			FieldValue: func() interface{} { return data.TStringUuid5 },
			BeforeFunc: func() { data.TStringUuid5 = nil },
			AfterFunc:  func() { x := "78381e24-2d01-5a28-b2bf-9eedab6605ed"; data.TStringUuid5 = &x },
		},
	}

	var cases []*CaseDesc
	cases = append(cases, caseGeneral1...)
	cases = append(cases, caseGeneral1NilPoint...)
	cases = append(cases, caseGeneral2...)
	cases = append(cases, caseGeneral2NilPoint...)
	cases = append(cases, caseNetwork1...)
	cases = append(cases, caseNetwork1NilPoint...)
	cases = append(cases, caseFormat1...)
	cases = append(cases, caseFormat1NilPoint...)

	msgName := "ValidStringTagsOptional1"
	runCases(t, data, msgName, cases)

	{
		err := data.Validate()
		require.Nil(t, err)
	}
}

func Test_GoValidator_ValidStringTagsOneOf1(t *testing.T) {
	data := &govalidatortest.ValidStringTagsOneOf1{}
	{
		err := data.Validate()
		require.Nil(t, err)
	}

	caseGeneral1 := []*CaseDesc{
		{
			Name:      fmt.Sprintf("test tag <%s> 1", protovalidator.TagStringEq),
			FieldDesc: "field 't_string_eq1'",
			TagInfo:   &protovalidator.TagInfo{Tag: protovalidator.TagStringEq, Value: "b"},
			FieldValue: func() interface{} {
				return data.OneTyp1.(*govalidatortest.ValidStringTagsOneOf1_TStringEq1).TStringEq1
			},
			BeforeFunc: func() { x := ""; data.OneTyp1 = &govalidatortest.ValidStringTagsOneOf1_TStringEq1{TStringEq1: x} },
			AfterFunc:  func() { x := "b"; data.OneTyp1 = &govalidatortest.ValidStringTagsOneOf1_TStringEq1{TStringEq1: x} },
		},
		{
			Name:      fmt.Sprintf("test tag <%s> 1", protovalidator.TagStringNe),
			FieldDesc: "field 't_string_ne1'",
			TagInfo:   &protovalidator.TagInfo{Tag: protovalidator.TagStringNe, Value: "b"},
			FieldValue: func() interface{} {
				return data.OneTyp1.(*govalidatortest.ValidStringTagsOneOf1_TStringNe1).TStringNe1
			},
			BeforeFunc: func() { x := "b"; data.OneTyp1 = &govalidatortest.ValidStringTagsOneOf1_TStringNe1{TStringNe1: x} },
			AfterFunc:  func() { x := "a"; data.OneTyp1 = &govalidatortest.ValidStringTagsOneOf1_TStringNe1{TStringNe1: x} },
		},
		{
			Name:      fmt.Sprintf("test tag <%s> 1", protovalidator.TagStringLt),
			FieldDesc: "field 't_string_lt1'",
			TagInfo:   &protovalidator.TagInfo{Tag: protovalidator.TagStringLt, Value: "c"},
			FieldValue: func() interface{} {
				return data.OneTyp1.(*govalidatortest.ValidStringTagsOneOf1_TStringLt1).TStringLt1
			},
			BeforeFunc: func() { x := "c"; data.OneTyp1 = &govalidatortest.ValidStringTagsOneOf1_TStringLt1{TStringLt1: x} },
			AfterFunc:  func() { x := "b"; data.OneTyp1 = &govalidatortest.ValidStringTagsOneOf1_TStringLt1{TStringLt1: x} },
		},
		{
			Name:      fmt.Sprintf("test tag <%s> 1", protovalidator.TagStringGt),
			FieldDesc: "field 't_string_gt1'",
			TagInfo:   &protovalidator.TagInfo{Tag: protovalidator.TagStringGt, Value: "c"},
			FieldValue: func() interface{} {
				return data.OneTyp1.(*govalidatortest.ValidStringTagsOneOf1_TStringGt1).TStringGt1
			},
			BeforeFunc: func() { x := "c"; data.OneTyp1 = &govalidatortest.ValidStringTagsOneOf1_TStringGt1{TStringGt1: x} },
			AfterFunc:  func() { x := "d"; data.OneTyp1 = &govalidatortest.ValidStringTagsOneOf1_TStringGt1{TStringGt1: x} },
		},
		{
			Name:      fmt.Sprintf("test tag <%s> 1", protovalidator.TagStringLte),
			FieldDesc: "field 't_string_lte1'",
			TagInfo:   &protovalidator.TagInfo{Tag: protovalidator.TagStringLte, Value: "c"},
			FieldValue: func() interface{} {
				return data.OneTyp1.(*govalidatortest.ValidStringTagsOneOf1_TStringLte1).TStringLte1
			},
			BeforeFunc: func() { x := "d"; data.OneTyp1 = &govalidatortest.ValidStringTagsOneOf1_TStringLte1{TStringLte1: x} },
			AfterFunc:  func() { x := "c"; data.OneTyp1 = &govalidatortest.ValidStringTagsOneOf1_TStringLte1{TStringLte1: x} },
		},
		{
			Name:      fmt.Sprintf("test tag <%s> 1", protovalidator.TagStringGte),
			FieldDesc: "field 't_string_gte1'",
			TagInfo:   &protovalidator.TagInfo{Tag: protovalidator.TagStringGte, Value: "c"},
			FieldValue: func() interface{} {
				return data.OneTyp1.(*govalidatortest.ValidStringTagsOneOf1_TStringGte1).TStringGte1
			},
			BeforeFunc: func() { x := "b"; data.OneTyp1 = &govalidatortest.ValidStringTagsOneOf1_TStringGte1{TStringGte1: x} },
			AfterFunc:  func() { x := "c"; data.OneTyp1 = &govalidatortest.ValidStringTagsOneOf1_TStringGte1{TStringGte1: x} },
		},
		{
			Name:      fmt.Sprintf("test tag <%s> 1", protovalidator.TagStringIn),
			FieldDesc: "field 't_string_in1'",
			TagInfo:   &protovalidator.TagInfo{Tag: protovalidator.TagStringIn, Value: []string{"a", "b", "c"}},
			FieldValue: func() interface{} {
				return data.OneTyp1.(*govalidatortest.ValidStringTagsOneOf1_TStringIn1).TStringIn1
			},
			BeforeFunc: func() { x := "d"; data.OneTyp1 = &govalidatortest.ValidStringTagsOneOf1_TStringIn1{TStringIn1: x} },
			AfterFunc:  func() { x := "b"; data.OneTyp1 = &govalidatortest.ValidStringTagsOneOf1_TStringIn1{TStringIn1: x} },
		},
		{
			Name:      fmt.Sprintf("test tag <%s> 1", protovalidator.TagStringNotIn),
			FieldDesc: "field 't_string_not_in1'",
			TagInfo:   &protovalidator.TagInfo{Tag: protovalidator.TagStringNotIn, Value: []string{"x", "y", "z"}},
			FieldValue: func() interface{} {
				return data.OneTyp1.(*govalidatortest.ValidStringTagsOneOf1_TStringNotIn1).TStringNotIn1
			},
			BeforeFunc: func() {
				x := "x"
				data.OneTyp1 = &govalidatortest.ValidStringTagsOneOf1_TStringNotIn1{TStringNotIn1: x}
			},
			AfterFunc: func() {
				x := "v"
				data.OneTyp1 = &govalidatortest.ValidStringTagsOneOf1_TStringNotIn1{TStringNotIn1: x}
			},
		},
		{
			Name:      fmt.Sprintf("test tag <%s> 1", protovalidator.TagStringCharLenEq),
			FieldDesc: "field 't_string_char_len_eq1'",
			TagInfo:   &protovalidator.TagInfo{Tag: protovalidator.TagStringCharLenEq, Value: 3},
			FieldValue: func() interface{} {
				x := data.OneTyp1.(*govalidatortest.ValidStringTagsOneOf1_TStringCharLenEq1).TStringCharLenEq1
				return utf8.RuneCountInString(x)
			},
			BeforeFunc: func() {
				x := "x"
				data.OneTyp1 = &govalidatortest.ValidStringTagsOneOf1_TStringCharLenEq1{TStringCharLenEq1: x}
			},
			AfterFunc: func() {
				x := "a中b"
				data.OneTyp1 = &govalidatortest.ValidStringTagsOneOf1_TStringCharLenEq1{TStringCharLenEq1: x}
			},
		},
		{
			Name:      fmt.Sprintf("test tag <%s> 1", protovalidator.TagStringCharLenNe),
			FieldDesc: "field 't_string_char_len_ne1'",
			TagInfo:   &protovalidator.TagInfo{Tag: protovalidator.TagStringCharLenNe, Value: 3},
			FieldValue: func() interface{} {
				x := data.OneTyp1.(*govalidatortest.ValidStringTagsOneOf1_TStringCharLenNe1).TStringCharLenNe1
				return utf8.RuneCountInString(x)
			},
			BeforeFunc: func() {
				x := "a中b"
				data.OneTyp1 = &govalidatortest.ValidStringTagsOneOf1_TStringCharLenNe1{TStringCharLenNe1: x}
			},
			AfterFunc: func() {
				x := "中"
				data.OneTyp1 = &govalidatortest.ValidStringTagsOneOf1_TStringCharLenNe1{TStringCharLenNe1: x}
			},
		},
		{
			Name:      fmt.Sprintf("test tag <%s> 1", protovalidator.TagStringCharLenGt),
			FieldDesc: "field 't_string_char_len_gt1'",
			TagInfo:   &protovalidator.TagInfo{Tag: protovalidator.TagStringCharLenGt, Value: 3},
			FieldValue: func() interface{} {
				x := data.OneTyp1.(*govalidatortest.ValidStringTagsOneOf1_TStringCharLenGt1).TStringCharLenGt1
				return utf8.RuneCountInString(x)
			},
			BeforeFunc: func() {
				x := "a中b"
				data.OneTyp1 = &govalidatortest.ValidStringTagsOneOf1_TStringCharLenGt1{TStringCharLenGt1: x}
			},
			AfterFunc: func() {
				x := "a中b文"
				data.OneTyp1 = &govalidatortest.ValidStringTagsOneOf1_TStringCharLenGt1{TStringCharLenGt1: x}
			},
		},
		{
			Name:      fmt.Sprintf("test tag <%s> 1", protovalidator.TagStringCharLenLt),
			FieldDesc: "field 't_string_char_len_lt1'",
			TagInfo:   &protovalidator.TagInfo{Tag: protovalidator.TagStringCharLenLt, Value: 3},
			FieldValue: func() interface{} {
				x := data.OneTyp1.(*govalidatortest.ValidStringTagsOneOf1_TStringCharLenLt1).TStringCharLenLt1
				return utf8.RuneCountInString(x)
			},
			BeforeFunc: func() {
				x := "a中b"
				data.OneTyp1 = &govalidatortest.ValidStringTagsOneOf1_TStringCharLenLt1{TStringCharLenLt1: x}
			},
			AfterFunc: func() {
				x := "中"
				data.OneTyp1 = &govalidatortest.ValidStringTagsOneOf1_TStringCharLenLt1{TStringCharLenLt1: x}
			},
		},
		{
			Name:      fmt.Sprintf("test tag <%s> 1", protovalidator.TagStringCharLenGte),
			FieldDesc: "field 't_string_char_len_gte1'",
			TagInfo:   &protovalidator.TagInfo{Tag: protovalidator.TagStringCharLenGte, Value: 3},
			FieldValue: func() interface{} {
				x := data.OneTyp1.(*govalidatortest.ValidStringTagsOneOf1_TStringCharLenGte1).TStringCharLenGte1
				return utf8.RuneCountInString(x)
			},
			BeforeFunc: func() {
				x := "a中"
				data.OneTyp1 = &govalidatortest.ValidStringTagsOneOf1_TStringCharLenGte1{TStringCharLenGte1: x}
			},
			AfterFunc: func() {
				x := "a中b文"
				data.OneTyp1 = &govalidatortest.ValidStringTagsOneOf1_TStringCharLenGte1{TStringCharLenGte1: x}
			},
		},
		{
			Name:      fmt.Sprintf("test tag <%s> 1", protovalidator.TagStringCharLenLte),
			FieldDesc: "field 't_string_char_len_lte1'",
			TagInfo:   &protovalidator.TagInfo{Tag: protovalidator.TagStringCharLenLte, Value: 3},
			FieldValue: func() interface{} {
				x := data.OneTyp1.(*govalidatortest.ValidStringTagsOneOf1_TStringCharLenLte1).TStringCharLenLte1
				return utf8.RuneCountInString(x)
			},
			BeforeFunc: func() {
				x := "a中b文"
				data.OneTyp1 = &govalidatortest.ValidStringTagsOneOf1_TStringCharLenLte1{TStringCharLenLte1: x}
			},
			AfterFunc: func() {
				x := "a中b"
				data.OneTyp1 = &govalidatortest.ValidStringTagsOneOf1_TStringCharLenLte1{TStringCharLenLte1: x}
			},
		},

		{
			Name:      fmt.Sprintf("test tag <%s> 1", protovalidator.TagStringByteLenEq),
			FieldDesc: "field 't_string_byte_len_eq1'",
			TagInfo:   &protovalidator.TagInfo{Tag: protovalidator.TagStringByteLenEq, Value: 5},
			FieldValue: func() interface{} {
				x := data.OneTyp1.(*govalidatortest.ValidStringTagsOneOf1_TStringByteLenEq1).TStringByteLenEq1
				return len(x)
			},
			BeforeFunc: func() {
				x := "a中b文C"
				data.OneTyp1 = &govalidatortest.ValidStringTagsOneOf1_TStringByteLenEq1{TStringByteLenEq1: x}
			},
			AfterFunc: func() {
				x := "a中b"
				data.OneTyp1 = &govalidatortest.ValidStringTagsOneOf1_TStringByteLenEq1{TStringByteLenEq1: x}
			},
		},
		{
			Name:      fmt.Sprintf("test tag <%s> 1", protovalidator.TagStringByteLenNe),
			FieldDesc: "field 't_string_byte_len_ne1'",
			TagInfo:   &protovalidator.TagInfo{Tag: protovalidator.TagStringByteLenNe, Value: 5},
			FieldValue: func() interface{} {
				x := data.OneTyp1.(*govalidatortest.ValidStringTagsOneOf1_TStringByteLenNe1).TStringByteLenNe1
				return len(x)
			},
			BeforeFunc: func() {
				x := "a中b"
				data.OneTyp1 = &govalidatortest.ValidStringTagsOneOf1_TStringByteLenNe1{TStringByteLenNe1: x}
			},
			AfterFunc: func() {
				x := "a"
				data.OneTyp1 = &govalidatortest.ValidStringTagsOneOf1_TStringByteLenNe1{TStringByteLenNe1: x}
			},
		},
		{
			Name:      fmt.Sprintf("test tag <%s> 1", protovalidator.TagStringByteLenGt),
			FieldDesc: "field 't_string_byte_len_gt1'",
			TagInfo:   &protovalidator.TagInfo{Tag: protovalidator.TagStringByteLenGt, Value: 5},
			FieldValue: func() interface{} {
				x := data.OneTyp1.(*govalidatortest.ValidStringTagsOneOf1_TStringByteLenGt1).TStringByteLenGt1
				return len(x)
			},
			BeforeFunc: func() {
				x := "a中b"
				data.OneTyp1 = &govalidatortest.ValidStringTagsOneOf1_TStringByteLenGt1{TStringByteLenGt1: x}
			},
			AfterFunc: func() {
				x := "a中bc"
				data.OneTyp1 = &govalidatortest.ValidStringTagsOneOf1_TStringByteLenGt1{TStringByteLenGt1: x}
			},
		},
		{
			Name:      fmt.Sprintf("test tag <%s> 1", protovalidator.TagStringByteLenLt),
			FieldDesc: "field 't_string_byte_len_lt1'",
			TagInfo:   &protovalidator.TagInfo{Tag: protovalidator.TagStringByteLenLt, Value: 5},
			FieldValue: func() interface{} {
				x := data.OneTyp1.(*govalidatortest.ValidStringTagsOneOf1_TStringByteLenLt1).TStringByteLenLt1
				return len(x)
			},
			BeforeFunc: func() {
				x := "a中b文"
				data.OneTyp1 = &govalidatortest.ValidStringTagsOneOf1_TStringByteLenLt1{TStringByteLenLt1: x}
			},
			AfterFunc: func() {
				x := "a中"
				data.OneTyp1 = &govalidatortest.ValidStringTagsOneOf1_TStringByteLenLt1{TStringByteLenLt1: x}
			},
		},
		{
			Name:      fmt.Sprintf("test tag <%s> 1", protovalidator.TagStringByteLenGte),
			FieldDesc: "field 't_string_byte_len_gte1'",
			TagInfo:   &protovalidator.TagInfo{Tag: protovalidator.TagStringByteLenGte, Value: 5},
			FieldValue: func() interface{} {
				x := data.OneTyp1.(*govalidatortest.ValidStringTagsOneOf1_TStringByteLenGte1).TStringByteLenGte1
				return len(x)
			},
			BeforeFunc: func() {
				x := "a中"
				data.OneTyp1 = &govalidatortest.ValidStringTagsOneOf1_TStringByteLenGte1{TStringByteLenGte1: x}
			},
			AfterFunc: func() {
				x := "a中b"
				data.OneTyp1 = &govalidatortest.ValidStringTagsOneOf1_TStringByteLenGte1{TStringByteLenGte1: x}
			},
		},
		{
			Name:      fmt.Sprintf("test tag <%s> 1", protovalidator.TagStringByteLenLte),
			FieldDesc: "field 't_string_byte_len_lte1'",
			TagInfo:   &protovalidator.TagInfo{Tag: protovalidator.TagStringByteLenLte, Value: 5},
			FieldValue: func() interface{} {
				x := data.OneTyp1.(*govalidatortest.ValidStringTagsOneOf1_TStringByteLenLte1).TStringByteLenLte1
				return len(x)
			},
			BeforeFunc: func() {
				x := "a中b文"
				data.OneTyp1 = &govalidatortest.ValidStringTagsOneOf1_TStringByteLenLte1{TStringByteLenLte1: x}
			},
			AfterFunc: func() {
				x := "a中b"
				data.OneTyp1 = &govalidatortest.ValidStringTagsOneOf1_TStringByteLenLte1{TStringByteLenLte1: x}
			},
		},
	}

	caseGeneral2 := []*CaseDesc{
		{
			Name:      fmt.Sprintf("test tag <%s> 1", protovalidator.TagStringRegex),
			FieldDesc: "field 't_string_regex1'",
			TagInfo:   &protovalidator.TagInfo{Tag: protovalidator.TagStringRegex, Value: `^[0-9]+.jar$`},
			FieldValue: func() interface{} {
				return data.OneTyp1.(*govalidatortest.ValidStringTagsOneOf1_TStringRegex1).TStringRegex1
			},
			BeforeFunc: func() {
				x := ""
				data.OneTyp1 = &govalidatortest.ValidStringTagsOneOf1_TStringRegex1{TStringRegex1: x}
			},
			AfterFunc: func() {
				x := "0001.jar"
				data.OneTyp1 = &govalidatortest.ValidStringTagsOneOf1_TStringRegex1{TStringRegex1: x}
			},
		},

		{
			Name:      fmt.Sprintf("test tag <%s> 1", protovalidator.TagStringPrefix),
			FieldDesc: "field 't_string_prefix1'",
			TagInfo:   &protovalidator.TagInfo{Tag: protovalidator.TagStringPrefix, Value: "prefix"},
			FieldValue: func() interface{} {
				return data.OneTyp1.(*govalidatortest.ValidStringTagsOneOf1_TStringPrefix1).TStringPrefix1
			},
			BeforeFunc: func() {
				x := ""
				data.OneTyp1 = &govalidatortest.ValidStringTagsOneOf1_TStringPrefix1{TStringPrefix1: x}
			},
			AfterFunc: func() {
				x := "prefix-xxx1"
				data.OneTyp1 = &govalidatortest.ValidStringTagsOneOf1_TStringPrefix1{TStringPrefix1: x}
			},
		},
		{
			Name:      fmt.Sprintf("test tag <%s> 1", protovalidator.TagStringNoPrefix),
			FieldDesc: "field 't_string_no_prefix1'",
			TagInfo:   &protovalidator.TagInfo{Tag: protovalidator.TagStringNoPrefix, Value: "prefix"},
			FieldValue: func() interface{} {
				return data.OneTyp1.(*govalidatortest.ValidStringTagsOneOf1_TStringNoPrefix1).TStringNoPrefix1
			},
			BeforeFunc: func() {
				x := "prefix-xxx1"
				data.OneTyp1 = &govalidatortest.ValidStringTagsOneOf1_TStringNoPrefix1{TStringNoPrefix1: x}
			},
			AfterFunc: func() {
				x := "xxx1"
				data.OneTyp1 = &govalidatortest.ValidStringTagsOneOf1_TStringNoPrefix1{TStringNoPrefix1: x}
			},
		},

		{
			Name:      fmt.Sprintf("test tag <%s> 1", protovalidator.TagStringSuffix),
			FieldDesc: "field 't_string_suffix1'",
			TagInfo:   &protovalidator.TagInfo{Tag: protovalidator.TagStringSuffix, Value: "suffix"},
			FieldValue: func() interface{} {
				return data.OneTyp1.(*govalidatortest.ValidStringTagsOneOf1_TStringSuffix1).TStringSuffix1
			},
			BeforeFunc: func() {
				x := ""
				data.OneTyp1 = &govalidatortest.ValidStringTagsOneOf1_TStringSuffix1{TStringSuffix1: x}
			},
			AfterFunc: func() {
				x := "xxx-suffix"
				data.OneTyp1 = &govalidatortest.ValidStringTagsOneOf1_TStringSuffix1{TStringSuffix1: x}
			},
		},
		{
			Name:      fmt.Sprintf("test tag <%s> 1", protovalidator.TagStringNoSuffix),
			FieldDesc: "field 't_string_no_suffix1'",
			TagInfo:   &protovalidator.TagInfo{Tag: protovalidator.TagStringNoSuffix, Value: "suffix"},
			FieldValue: func() interface{} {
				return data.OneTyp1.(*govalidatortest.ValidStringTagsOneOf1_TStringNoSuffix1).TStringNoSuffix1
			},
			BeforeFunc: func() {
				x := "xxx-suffix"
				data.OneTyp1 = &govalidatortest.ValidStringTagsOneOf1_TStringNoSuffix1{TStringNoSuffix1: x}
			},
			AfterFunc: func() {
				x := "xxx"
				data.OneTyp1 = &govalidatortest.ValidStringTagsOneOf1_TStringNoSuffix1{TStringNoSuffix1: x}
			},
		},

		{
			Name:      fmt.Sprintf("test tag <%s> 1", protovalidator.TagStringContains),
			FieldDesc: "field 't_string_contains1'",
			TagInfo:   &protovalidator.TagInfo{Tag: protovalidator.TagStringContains, Value: "abc"},
			FieldValue: func() interface{} {
				return data.OneTyp1.(*govalidatortest.ValidStringTagsOneOf1_TStringContains1).TStringContains1
			},
			BeforeFunc: func() {
				x := "1a2b3c"
				data.OneTyp1 = &govalidatortest.ValidStringTagsOneOf1_TStringContains1{TStringContains1: x}
			},
			AfterFunc: func() {
				x := "1abc23"
				data.OneTyp1 = &govalidatortest.ValidStringTagsOneOf1_TStringContains1{TStringContains1: x}
			},
		},
		{
			Name:      fmt.Sprintf("test tag <%s> 1", protovalidator.TagStringNotContains),
			FieldDesc: "field 't_string_no_contains1'",
			TagInfo:   &protovalidator.TagInfo{Tag: protovalidator.TagStringNotContains, Value: "abc"},
			FieldValue: func() interface{} {
				return data.OneTyp1.(*govalidatortest.ValidStringTagsOneOf1_TStringNoContains1).TStringNoContains1
			},
			BeforeFunc: func() {
				x := "1abc23"
				data.OneTyp1 = &govalidatortest.ValidStringTagsOneOf1_TStringNoContains1{TStringNoContains1: x}
			},
			AfterFunc: func() {
				x := "1a2b3c"
				data.OneTyp1 = &govalidatortest.ValidStringTagsOneOf1_TStringNoContains1{TStringNoContains1: x}
			},
		},

		{
			Name:      fmt.Sprintf("test tag <%s> 1", protovalidator.TagStringContainsAny),
			FieldDesc: "field 't_string_contains_any1'",
			TagInfo:   &protovalidator.TagInfo{Tag: protovalidator.TagStringContainsAny, Value: "xyz"},
			FieldValue: func() interface{} {
				return data.OneTyp1.(*govalidatortest.ValidStringTagsOneOf1_TStringContainsAny1).TStringContainsAny1
			},
			BeforeFunc: func() {
				x := "mn"
				data.OneTyp1 = &govalidatortest.ValidStringTagsOneOf1_TStringContainsAny1{TStringContainsAny1: x}
			},
			AfterFunc: func() {
				x := "x121"
				data.OneTyp1 = &govalidatortest.ValidStringTagsOneOf1_TStringContainsAny1{TStringContainsAny1: x}
			},
		},
		{
			Name:      fmt.Sprintf("test tag <%s> 1", protovalidator.TagStringNotContainsAny),
			FieldDesc: "field 't_string_not_contains_any1'",
			TagInfo:   &protovalidator.TagInfo{Tag: protovalidator.TagStringNotContainsAny, Value: "xyz"},
			FieldValue: func() interface{} {
				return data.OneTyp1.(*govalidatortest.ValidStringTagsOneOf1_TStringNotContainsAny1).TStringNotContainsAny1
			},
			BeforeFunc: func() {
				x := "1y1"
				data.OneTyp1 = &govalidatortest.ValidStringTagsOneOf1_TStringNotContainsAny1{TStringNotContainsAny1: x}
			},
			AfterFunc: func() {
				x := "123"
				data.OneTyp1 = &govalidatortest.ValidStringTagsOneOf1_TStringNotContainsAny1{TStringNotContainsAny1: x}
			},
		},

		{
			Name:      fmt.Sprintf("test tag <%s> 1", protovalidator.TagStringUTF8),
			FieldDesc: "field 't_string_utf8'",
			TagInfo:   &protovalidator.TagInfo{Tag: protovalidator.TagStringUTF8, Value: nil},
			FieldValue: func() interface{} {
				return data.OneTyp1.(*govalidatortest.ValidStringTagsOneOf1_TStringUtf8).TStringUtf8
			},
			BeforeFunc: func() {
				src := bytes.NewReader([]byte("中文abc"))
				encoder := simplifiedchinese.GBK.NewEncoder()
				tr := transform.NewReader(src, encoder)
				b, err := ioutil.ReadAll(tr)
				if err != nil {
					panic(err)
				}
				x := string(b)
				data.OneTyp1 = &govalidatortest.ValidStringTagsOneOf1_TStringUtf8{TStringUtf8: x}
			},
			AfterFunc: func() {
				x := "中文abc"
				data.OneTyp1 = &govalidatortest.ValidStringTagsOneOf1_TStringUtf8{TStringUtf8: x}
			},
		},

		{
			Name:      fmt.Sprintf("test tag <%s> 1", protovalidator.TagStringAscii),
			FieldDesc: "field 't_string_ascii'",
			TagInfo:   &protovalidator.TagInfo{Tag: protovalidator.TagStringAscii, Value: nil},
			FieldValue: func() interface{} {
				return data.OneTyp1.(*govalidatortest.ValidStringTagsOneOf1_TStringAscii).TStringAscii
			},
			BeforeFunc: func() {
				src := bytes.NewReader([]byte("中文abc"))
				encoder := simplifiedchinese.GBK.NewEncoder()
				tr := transform.NewReader(src, encoder)
				b, err := ioutil.ReadAll(tr)
				if err != nil {
					panic(err)
				}
				x := string(b)
				data.OneTyp1 = &govalidatortest.ValidStringTagsOneOf1_TStringAscii{TStringAscii: x}
			},
			AfterFunc: func() {
				x := "abc"
				data.OneTyp1 = &govalidatortest.ValidStringTagsOneOf1_TStringAscii{TStringAscii: x}
			},
		},
		{
			Name:      fmt.Sprintf("test tag <%s> 1", protovalidator.TagStringPrintAscii),
			FieldDesc: "field 't_string_print_ascii'",
			TagInfo:   &protovalidator.TagInfo{Tag: protovalidator.TagStringPrintAscii, Value: nil},
			FieldValue: func() interface{} {
				return data.OneTyp1.(*govalidatortest.ValidStringTagsOneOf1_TStringPrintAscii).TStringPrintAscii
			},
			BeforeFunc: func() {
				x := string([]byte{3, 2, 1})
				data.OneTyp1 = &govalidatortest.ValidStringTagsOneOf1_TStringPrintAscii{TStringPrintAscii: x}
			},
			AfterFunc: func() {
				x := "abc"
				data.OneTyp1 = &govalidatortest.ValidStringTagsOneOf1_TStringPrintAscii{TStringPrintAscii: x}
			},
		},
		{
			Name:      fmt.Sprintf("test tag <%s> 1", protovalidator.TagStringBoolean),
			FieldDesc: "field 't_string_boolean'",
			TagInfo:   &protovalidator.TagInfo{Tag: protovalidator.TagStringBoolean, Value: nil},
			FieldValue: func() interface{} {
				return data.OneTyp1.(*govalidatortest.ValidStringTagsOneOf1_TStringBoolean).TStringBoolean
			},
			BeforeFunc: func() {
				x := "bool"
				data.OneTyp1 = &govalidatortest.ValidStringTagsOneOf1_TStringBoolean{TStringBoolean: x}
			},
			AfterFunc: func() {
				x := "true"
				data.OneTyp1 = &govalidatortest.ValidStringTagsOneOf1_TStringBoolean{TStringBoolean: x}
			},
		},
		{
			Name:      fmt.Sprintf("test tag <%s> 1", protovalidator.TagStringLowercase),
			FieldDesc: "field 't_string_lowercase'",
			TagInfo:   &protovalidator.TagInfo{Tag: protovalidator.TagStringLowercase, Value: nil},
			FieldValue: func() interface{} {
				return data.OneTyp1.(*govalidatortest.ValidStringTagsOneOf1_TStringLowercase).TStringLowercase
			},
			BeforeFunc: func() {
				x := "AbC"
				data.OneTyp1 = &govalidatortest.ValidStringTagsOneOf1_TStringLowercase{TStringLowercase: x}
			},
			AfterFunc: func() {
				x := "abc"
				data.OneTyp1 = &govalidatortest.ValidStringTagsOneOf1_TStringLowercase{TStringLowercase: x}
			},
		},
		{
			Name:      fmt.Sprintf("test tag <%s> 1", protovalidator.TagStringUppercase),
			FieldDesc: "field 't_string_uppercase'",
			TagInfo:   &protovalidator.TagInfo{Tag: protovalidator.TagStringUppercase, Value: nil},
			FieldValue: func() interface{} {
				return data.OneTyp1.(*govalidatortest.ValidStringTagsOneOf1_TStringUppercase).TStringUppercase
			},
			BeforeFunc: func() {
				x := "aBC"
				data.OneTyp1 = &govalidatortest.ValidStringTagsOneOf1_TStringUppercase{TStringUppercase: x}
			},
			AfterFunc: func() {
				x := "ABC"
				data.OneTyp1 = &govalidatortest.ValidStringTagsOneOf1_TStringUppercase{TStringUppercase: x}
			},
		},
		{
			Name:      fmt.Sprintf("test tag <%s> 1", protovalidator.TagStringAlpha),
			FieldDesc: "field 't_string_alpha'",
			TagInfo:   &protovalidator.TagInfo{Tag: protovalidator.TagStringAlpha, Value: nil},
			FieldValue: func() interface{} {
				return data.OneTyp1.(*govalidatortest.ValidStringTagsOneOf1_TStringAlpha).TStringAlpha
			},
			BeforeFunc: func() {
				x := "123"
				data.OneTyp1 = &govalidatortest.ValidStringTagsOneOf1_TStringAlpha{TStringAlpha: x}
			},
			AfterFunc: func() {
				x := "abc"
				data.OneTyp1 = &govalidatortest.ValidStringTagsOneOf1_TStringAlpha{TStringAlpha: x}
			},
		},
		{
			Name:      fmt.Sprintf("test tag <%s> 1", protovalidator.TagStringNumber),
			FieldDesc: "field 't_string_number'",
			TagInfo:   &protovalidator.TagInfo{Tag: protovalidator.TagStringNumber, Value: nil},
			FieldValue: func() interface{} {
				return data.OneTyp1.(*govalidatortest.ValidStringTagsOneOf1_TStringNumber).TStringNumber
			},
			BeforeFunc: func() {
				x := "abc"
				data.OneTyp1 = &govalidatortest.ValidStringTagsOneOf1_TStringNumber{TStringNumber: x}
			},
			AfterFunc: func() {
				x := "123"
				data.OneTyp1 = &govalidatortest.ValidStringTagsOneOf1_TStringNumber{TStringNumber: x}
			},
		},
		{
			Name:      fmt.Sprintf("test tag <%s> 1", protovalidator.TagStringAlphaNumber),
			FieldDesc: "field 't_string_alpha_number'",
			TagInfo:   &protovalidator.TagInfo{Tag: protovalidator.TagStringAlphaNumber, Value: nil},
			FieldValue: func() interface{} {
				return data.OneTyp1.(*govalidatortest.ValidStringTagsOneOf1_TStringAlphaNumber).TStringAlphaNumber
			},
			BeforeFunc: func() {
				x := "中文"
				data.OneTyp1 = &govalidatortest.ValidStringTagsOneOf1_TStringAlphaNumber{TStringAlphaNumber: x}
			},
			AfterFunc: func() {
				x := "abc123"
				data.OneTyp1 = &govalidatortest.ValidStringTagsOneOf1_TStringAlphaNumber{TStringAlphaNumber: x}
			},
		},
	}

	caseNetwork1 := []*CaseDesc{
		{
			Name:       fmt.Sprintf("test tag <%s> 1", protovalidator.TagStringIp),
			FieldDesc:  "field 't_string_ip'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagStringIp, Value: nil},
			FieldValue: func() interface{} { return data.OneTyp1.(*govalidatortest.ValidStringTagsOneOf1_TStringIp).TStringIp },
			BeforeFunc: func() { x := "127.0."; data.OneTyp1 = &govalidatortest.ValidStringTagsOneOf1_TStringIp{TStringIp: x} },
			AfterFunc: func() {
				x := "127.0.0.1"
				data.OneTyp1 = &govalidatortest.ValidStringTagsOneOf1_TStringIp{TStringIp: x}
			},
		},
		{
			Name:       fmt.Sprintf("test tag <%s> 2", protovalidator.TagStringIp),
			FieldDesc:  "field 't_string_ip'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagStringIp, Value: nil},
			FieldValue: func() interface{} { return data.OneTyp1.(*govalidatortest.ValidStringTagsOneOf1_TStringIp).TStringIp },
			BeforeFunc: func() {
				x := "AD80:0000:0000:0000:ABAA:0000:00C2:0002:xxxxx"
				data.OneTyp1 = &govalidatortest.ValidStringTagsOneOf1_TStringIp{TStringIp: x}
			},
			AfterFunc: func() {
				x := "AD80:0000:0000:0000:ABAA:0000:00C2:0002"
				data.OneTyp1 = &govalidatortest.ValidStringTagsOneOf1_TStringIp{TStringIp: x}
			},
		},
		{
			Name:      fmt.Sprintf("test tag <%s> 1", protovalidator.TagStringIpv4),
			FieldDesc: "field 't_string_ipv4'",
			TagInfo:   &protovalidator.TagInfo{Tag: protovalidator.TagStringIpv4, Value: nil},
			FieldValue: func() interface{} {
				return data.OneTyp1.(*govalidatortest.ValidStringTagsOneOf1_TStringIpv4).TStringIpv4
			},
			BeforeFunc: func() {
				x := "127.0.0.1.1"
				data.OneTyp1 = &govalidatortest.ValidStringTagsOneOf1_TStringIpv4{TStringIpv4: x}
			},
			AfterFunc: func() {
				x := "127.0.0.1"
				data.OneTyp1 = &govalidatortest.ValidStringTagsOneOf1_TStringIpv4{TStringIpv4: x}
			},
		},
		{
			Name:      fmt.Sprintf("test tag <%s> 1", protovalidator.TagStringIpv6),
			FieldDesc: "field 't_string_ipv6'",
			TagInfo:   &protovalidator.TagInfo{Tag: protovalidator.TagStringIpv6, Value: nil},
			FieldValue: func() interface{} {
				return data.OneTyp1.(*govalidatortest.ValidStringTagsOneOf1_TStringIpv6).TStringIpv6
			},
			BeforeFunc: func() {
				x := "127.0.0.1"
				data.OneTyp1 = &govalidatortest.ValidStringTagsOneOf1_TStringIpv6{TStringIpv6: x}
			},
			AfterFunc: func() {
				x := "AD80:0000:0000:0000:ABAA:0000:00C2:0002"
				data.OneTyp1 = &govalidatortest.ValidStringTagsOneOf1_TStringIpv6{TStringIpv6: x}
			},
		},
		{
			Name:      fmt.Sprintf("test tag <%s> 1", protovalidator.TagStringIpAddr),
			FieldDesc: "field 't_string_ip_addr'",
			TagInfo:   &protovalidator.TagInfo{Tag: protovalidator.TagStringIpAddr, Value: nil},
			FieldValue: func() interface{} {
				return data.OneTyp1.(*govalidatortest.ValidStringTagsOneOf1_TStringIpAddr).TStringIpAddr
			},
			BeforeFunc: func() {
				x := "127.0.0.1.x"
				data.OneTyp1 = &govalidatortest.ValidStringTagsOneOf1_TStringIpAddr{TStringIpAddr: x}
			},
			AfterFunc: func() {
				x := "127.0.0.1"
				data.OneTyp1 = &govalidatortest.ValidStringTagsOneOf1_TStringIpAddr{TStringIpAddr: x}
			},
		},
		{
			Name:      fmt.Sprintf("test tag <%s> 1", protovalidator.TagStringIp4Addr),
			FieldDesc: "field 't_string_ip4_addr'",
			TagInfo:   &protovalidator.TagInfo{Tag: protovalidator.TagStringIp4Addr, Value: nil},
			FieldValue: func() interface{} {
				return data.OneTyp1.(*govalidatortest.ValidStringTagsOneOf1_TStringIp4Addr).TStringIp4Addr
			},
			BeforeFunc: func() {
				x := "127.0.0.1.x"
				data.OneTyp1 = &govalidatortest.ValidStringTagsOneOf1_TStringIp4Addr{TStringIp4Addr: x}
			},
			AfterFunc: func() {
				x := "127.0.0.1"
				data.OneTyp1 = &govalidatortest.ValidStringTagsOneOf1_TStringIp4Addr{TStringIp4Addr: x}
			},
		},
		{
			Name:      fmt.Sprintf("test tag <%s> 1", protovalidator.TagStringIp6Addr),
			FieldDesc: "field 't_string_ip6_addr'",
			TagInfo:   &protovalidator.TagInfo{Tag: protovalidator.TagStringIp6Addr, Value: nil},
			FieldValue: func() interface{} {
				return data.OneTyp1.(*govalidatortest.ValidStringTagsOneOf1_TStringIp6Addr).TStringIp6Addr
			},
			BeforeFunc: func() {
				x := "127.0.0.1"
				data.OneTyp1 = &govalidatortest.ValidStringTagsOneOf1_TStringIp6Addr{TStringIp6Addr: x}
			},
			AfterFunc: func() {
				x := "AD80:0000:0000:0000:ABAA:0000:00C2:0002"
				data.OneTyp1 = &govalidatortest.ValidStringTagsOneOf1_TStringIp6Addr{TStringIp6Addr: x}
			},
		},
		{
			Name:      fmt.Sprintf("test tag <%s> 1", protovalidator.TagStringCidr),
			FieldDesc: "field 't_string_cidr'",
			TagInfo:   &protovalidator.TagInfo{Tag: protovalidator.TagStringCidr, Value: nil},
			FieldValue: func() interface{} {
				return data.OneTyp1.(*govalidatortest.ValidStringTagsOneOf1_TStringCidr).TStringCidr
			},
			BeforeFunc: func() {
				x := "192.0.2.0/128"
				data.OneTyp1 = &govalidatortest.ValidStringTagsOneOf1_TStringCidr{TStringCidr: x}
			},
			AfterFunc: func() {
				x := "192.0.2.0/24"
				data.OneTyp1 = &govalidatortest.ValidStringTagsOneOf1_TStringCidr{TStringCidr: x}
			},
		},
		{
			Name:      fmt.Sprintf("test tag <%s> 1", protovalidator.TagStringCidrv4),
			FieldDesc: "field 't_string_cidrv4'",
			TagInfo:   &protovalidator.TagInfo{Tag: protovalidator.TagStringCidrv4, Value: nil},
			FieldValue: func() interface{} {
				return data.OneTyp1.(*govalidatortest.ValidStringTagsOneOf1_TStringCidrv4).TStringCidrv4
			},
			BeforeFunc: func() {
				x := "192.0.2.0/243"
				data.OneTyp1 = &govalidatortest.ValidStringTagsOneOf1_TStringCidrv4{TStringCidrv4: x}
			},
			AfterFunc: func() {
				x := "192.0.2.0/24"
				data.OneTyp1 = &govalidatortest.ValidStringTagsOneOf1_TStringCidrv4{TStringCidrv4: x}
			},
		},
		{
			Name:      fmt.Sprintf("test tag <%s> 1", protovalidator.TagStringCidrv6),
			FieldDesc: "field 't_string_cidrv6'",
			TagInfo:   &protovalidator.TagInfo{Tag: protovalidator.TagStringCidrv6, Value: nil},
			FieldValue: func() interface{} {
				return data.OneTyp1.(*govalidatortest.ValidStringTagsOneOf1_TStringCidrv6).TStringCidrv6
			},
			BeforeFunc: func() {
				x := "2001:db8::/129"
				data.OneTyp1 = &govalidatortest.ValidStringTagsOneOf1_TStringCidrv6{TStringCidrv6: x}
			},
			AfterFunc: func() {
				x := "2001:db8::/32"
				data.OneTyp1 = &govalidatortest.ValidStringTagsOneOf1_TStringCidrv6{TStringCidrv6: x}
			},
		},
		{
			Name:      fmt.Sprintf("test tag <%s> 1", protovalidator.TagStringMac),
			FieldDesc: "field 't_string_mac'",
			TagInfo:   &protovalidator.TagInfo{Tag: protovalidator.TagStringMac, Value: nil},
			FieldValue: func() interface{} {
				return data.OneTyp1.(*govalidatortest.ValidStringTagsOneOf1_TStringMac).TStringMac
			},
			BeforeFunc: func() {
				x := "00:00:5e:00:53:01:dsafadsfd"
				data.OneTyp1 = &govalidatortest.ValidStringTagsOneOf1_TStringMac{TStringMac: x}
			},
			AfterFunc: func() {
				x := "00:00:5e:00:53:01"
				data.OneTyp1 = &govalidatortest.ValidStringTagsOneOf1_TStringMac{TStringMac: x}
			},
		},
		{
			Name:      fmt.Sprintf("test tag <%s> 1", protovalidator.TagStringTcpAddr),
			FieldDesc: "field 't_string_tcp_addr'",
			TagInfo:   &protovalidator.TagInfo{Tag: protovalidator.TagStringTcpAddr, Value: nil},
			FieldValue: func() interface{} {
				return data.OneTyp1.(*govalidatortest.ValidStringTagsOneOf1_TStringTcpAddr).TStringTcpAddr
			},
			BeforeFunc: func() {
				x := ""
				data.OneTyp1 = &govalidatortest.ValidStringTagsOneOf1_TStringTcpAddr{TStringTcpAddr: x}
			},
			AfterFunc: func() {
				x := "127.0.0.1:80"
				data.OneTyp1 = &govalidatortest.ValidStringTagsOneOf1_TStringTcpAddr{TStringTcpAddr: x}
			},
		},
		{
			Name:      fmt.Sprintf("test tag <%s> 1", protovalidator.TagStringTcp4Addr),
			FieldDesc: "field 't_string_tcp4_addr'",
			TagInfo:   &protovalidator.TagInfo{Tag: protovalidator.TagStringTcp4Addr, Value: nil},
			FieldValue: func() interface{} {
				return data.OneTyp1.(*govalidatortest.ValidStringTagsOneOf1_TStringTcp4Addr).TStringTcp4Addr
			},
			BeforeFunc: func() {
				x := ""
				data.OneTyp1 = &govalidatortest.ValidStringTagsOneOf1_TStringTcp4Addr{TStringTcp4Addr: x}
			},
			AfterFunc: func() {
				x := "127.0.0.1:80"
				data.OneTyp1 = &govalidatortest.ValidStringTagsOneOf1_TStringTcp4Addr{TStringTcp4Addr: x}
			},
		},
		{
			Name:      fmt.Sprintf("test tag <%s> 1", protovalidator.TagStringTcp6Addr),
			FieldDesc: "field 't_string_tcp6_addr'",
			TagInfo:   &protovalidator.TagInfo{Tag: protovalidator.TagStringTcp6Addr, Value: nil},
			FieldValue: func() interface{} {
				return data.OneTyp1.(*govalidatortest.ValidStringTagsOneOf1_TStringTcp6Addr).TStringTcp6Addr
			},
			BeforeFunc: func() {
				x := ""
				data.OneTyp1 = &govalidatortest.ValidStringTagsOneOf1_TStringTcp6Addr{TStringTcp6Addr: x}
			},
			AfterFunc: func() {
				x := "[2001:db8::1]:http"
				data.OneTyp1 = &govalidatortest.ValidStringTagsOneOf1_TStringTcp6Addr{TStringTcp6Addr: x}
			},
		},
		{
			Name:      fmt.Sprintf("test tag <%s> 1", protovalidator.TagStringUdpAddr),
			FieldDesc: "field 't_string_udp_addr'",
			TagInfo:   &protovalidator.TagInfo{Tag: protovalidator.TagStringUdpAddr, Value: nil},
			FieldValue: func() interface{} {
				return data.OneTyp1.(*govalidatortest.ValidStringTagsOneOf1_TStringUdpAddr).TStringUdpAddr
			},
			BeforeFunc: func() {
				x := ""
				data.OneTyp1 = &govalidatortest.ValidStringTagsOneOf1_TStringUdpAddr{TStringUdpAddr: x}
			},
			AfterFunc: func() {
				x := "127.0.0.1:domain"
				data.OneTyp1 = &govalidatortest.ValidStringTagsOneOf1_TStringUdpAddr{TStringUdpAddr: x}
			},
		},
		{
			Name:      fmt.Sprintf("test tag <%s> 1", protovalidator.TagStringUdp4Addr),
			FieldDesc: "field 't_string_udp4_addr'",
			TagInfo:   &protovalidator.TagInfo{Tag: protovalidator.TagStringUdp4Addr, Value: nil},
			FieldValue: func() interface{} {
				return data.OneTyp1.(*govalidatortest.ValidStringTagsOneOf1_TStringUdp4Addr).TStringUdp4Addr
			},
			BeforeFunc: func() {
				x := ""
				data.OneTyp1 = &govalidatortest.ValidStringTagsOneOf1_TStringUdp4Addr{TStringUdp4Addr: x}
			},
			AfterFunc: func() {
				x := "127.0.0.1:domain"
				data.OneTyp1 = &govalidatortest.ValidStringTagsOneOf1_TStringUdp4Addr{TStringUdp4Addr: x}
			},
		},
		{
			Name:      fmt.Sprintf("test tag <%s> 1", protovalidator.TagStringUdp6Addr),
			FieldDesc: "field 't_string_udp6_addr'",
			TagInfo:   &protovalidator.TagInfo{Tag: protovalidator.TagStringUdp6Addr, Value: nil},
			FieldValue: func() interface{} {
				return data.OneTyp1.(*govalidatortest.ValidStringTagsOneOf1_TStringUdp6Addr).TStringUdp6Addr
			},
			BeforeFunc: func() {
				x := ""
				data.OneTyp1 = &govalidatortest.ValidStringTagsOneOf1_TStringUdp6Addr{TStringUdp6Addr: x}
			},
			AfterFunc: func() {
				x := "[2001:db8::1]:domain"
				data.OneTyp1 = &govalidatortest.ValidStringTagsOneOf1_TStringUdp6Addr{TStringUdp6Addr: x}
			},
		},
		{
			Name:      fmt.Sprintf("test tag <%s> 1", protovalidator.TagStringUnixAddr),
			FieldDesc: "field 't_string_unix_addr'",
			TagInfo:   &protovalidator.TagInfo{Tag: protovalidator.TagStringUnixAddr, Value: nil},
			FieldValue: func() interface{} {
				return data.OneTyp1.(*govalidatortest.ValidStringTagsOneOf1_TStringUnixAddr).TStringUnixAddr
			},
			BeforeFunc: func() {
				x := ""
				data.OneTyp1 = &govalidatortest.ValidStringTagsOneOf1_TStringUnixAddr{TStringUnixAddr: x}
			},
			AfterFunc: func() {
				x := "unix:///mysql/mysql.sock"
				data.OneTyp1 = &govalidatortest.ValidStringTagsOneOf1_TStringUnixAddr{TStringUnixAddr: x}
			},
		},
		{
			Name:      fmt.Sprintf("test tag <%s> 1", protovalidator.TagStringHostname),
			FieldDesc: "field 't_string_hostname'",
			TagInfo:   &protovalidator.TagInfo{Tag: protovalidator.TagStringHostname, Value: nil},
			FieldValue: func() interface{} {
				return data.OneTyp1.(*govalidatortest.ValidStringTagsOneOf1_TStringHostname).TStringHostname
			},
			BeforeFunc: func() {
				x := ""
				data.OneTyp1 = &govalidatortest.ValidStringTagsOneOf1_TStringHostname{TStringHostname: x}
			},
			AfterFunc: func() {
				x := "www.google.com"
				data.OneTyp1 = &govalidatortest.ValidStringTagsOneOf1_TStringHostname{TStringHostname: x}
			},
		},
		{
			Name:      fmt.Sprintf("test tag <%s> 1", protovalidator.TagStringHostnameRfc1123),
			FieldDesc: "field 't_string_hostname_rfc1123'",
			TagInfo:   &protovalidator.TagInfo{Tag: protovalidator.TagStringHostnameRfc1123, Value: nil},
			FieldValue: func() interface{} {
				return data.OneTyp1.(*govalidatortest.ValidStringTagsOneOf1_TStringHostnameRfc1123).TStringHostnameRfc1123
			},
			BeforeFunc: func() {
				x := ""
				data.OneTyp1 = &govalidatortest.ValidStringTagsOneOf1_TStringHostnameRfc1123{TStringHostnameRfc1123: x}
			},
			AfterFunc: func() {
				x := "www.google.com"
				data.OneTyp1 = &govalidatortest.ValidStringTagsOneOf1_TStringHostnameRfc1123{TStringHostnameRfc1123: x}
			},
		},
		{
			Name:      fmt.Sprintf("test tag <%s> 1", protovalidator.TagStringHostnamePort),
			FieldDesc: "field 't_string_hostname_port'",
			TagInfo:   &protovalidator.TagInfo{Tag: protovalidator.TagStringHostnamePort, Value: nil},
			FieldValue: func() interface{} {
				return data.OneTyp1.(*govalidatortest.ValidStringTagsOneOf1_TStringHostnamePort).TStringHostnamePort
			},
			BeforeFunc: func() {
				x := ""
				data.OneTyp1 = &govalidatortest.ValidStringTagsOneOf1_TStringHostnamePort{TStringHostnamePort: x}
			},
			AfterFunc: func() {
				x := "www.google.com:80"
				data.OneTyp1 = &govalidatortest.ValidStringTagsOneOf1_TStringHostnamePort{TStringHostnamePort: x}
			},
		},
		{
			Name:      fmt.Sprintf("test tag <%s> 1", protovalidator.TagStringDataURI),
			FieldDesc: "field 't_string_data_uri'",
			TagInfo:   &protovalidator.TagInfo{Tag: protovalidator.TagStringDataURI, Value: nil},
			FieldValue: func() interface{} {
				return data.OneTyp1.(*govalidatortest.ValidStringTagsOneOf1_TStringDataUri).TStringDataUri
			},
			BeforeFunc: func() {
				x := ""
				data.OneTyp1 = &govalidatortest.ValidStringTagsOneOf1_TStringDataUri{TStringDataUri: x}
			},
			AfterFunc: func() {
				x := "data:text/html,WkhPTkdHVU9uaWhhbzEyMw=="
				data.OneTyp1 = &govalidatortest.ValidStringTagsOneOf1_TStringDataUri{TStringDataUri: x}
			},
		},
		{
			Name:      fmt.Sprintf("test tag <%s> 1", protovalidator.TagStringFQDN),
			FieldDesc: "field 't_string_fqdn'",
			TagInfo:   &protovalidator.TagInfo{Tag: protovalidator.TagStringFQDN, Value: nil},
			FieldValue: func() interface{} {
				return data.OneTyp1.(*govalidatortest.ValidStringTagsOneOf1_TStringFqdn).TStringFqdn
			},
			BeforeFunc: func() { x := ""; data.OneTyp1 = &govalidatortest.ValidStringTagsOneOf1_TStringFqdn{TStringFqdn: x} },
			AfterFunc: func() {
				x := "MacBook-Pro.local"
				data.OneTyp1 = &govalidatortest.ValidStringTagsOneOf1_TStringFqdn{TStringFqdn: x}
			},
		},
		{
			Name:      fmt.Sprintf("test tag <%s> 1", protovalidator.TagStringURI),
			FieldDesc: "field 't_string_uri'",
			TagInfo:   &protovalidator.TagInfo{Tag: protovalidator.TagStringURI, Value: nil},
			FieldValue: func() interface{} {
				return data.OneTyp1.(*govalidatortest.ValidStringTagsOneOf1_TStringUri).TStringUri
			},
			BeforeFunc: func() { x := ""; data.OneTyp1 = &govalidatortest.ValidStringTagsOneOf1_TStringUri{TStringUri: x} },
			AfterFunc: func() {
				x := "/v1/api"
				data.OneTyp1 = &govalidatortest.ValidStringTagsOneOf1_TStringUri{TStringUri: x}
			},
		},
		{
			Name:      fmt.Sprintf("test tag <%s> 1", protovalidator.TagStringURL),
			FieldDesc: "field 't_string_url'",
			TagInfo:   &protovalidator.TagInfo{Tag: protovalidator.TagStringURL, Value: nil},
			FieldValue: func() interface{} {
				return data.OneTyp1.(*govalidatortest.ValidStringTagsOneOf1_TStringUrl).TStringUrl
			},
			BeforeFunc: func() { x := ""; data.OneTyp1 = &govalidatortest.ValidStringTagsOneOf1_TStringUrl{TStringUrl: x} },
			AfterFunc: func() {
				x := "https://www.google.com/v1/api"
				data.OneTyp1 = &govalidatortest.ValidStringTagsOneOf1_TStringUrl{TStringUrl: x}
			},
		},
		{
			Name:      fmt.Sprintf("test tag <%s> 1", protovalidator.TagStringURLEncoded),
			FieldDesc: "field 't_string_url_encoded'",
			TagInfo:   &protovalidator.TagInfo{Tag: protovalidator.TagStringURLEncoded, Value: nil},
			FieldValue: func() interface{} {
				return data.OneTyp1.(*govalidatortest.ValidStringTagsOneOf1_TStringUrlEncoded).TStringUrlEncoded
			},
			BeforeFunc: func() {
				x := ""
				data.OneTyp1 = &govalidatortest.ValidStringTagsOneOf1_TStringUrlEncoded{TStringUrlEncoded: x}
			},
			AfterFunc: func() {
				x := "https://www.google.com%2Fv1%2Fapi"
				data.OneTyp1 = &govalidatortest.ValidStringTagsOneOf1_TStringUrlEncoded{TStringUrlEncoded: x}
			},
		},
	}

	caseFormat1 := []*CaseDesc{
		{
			Name:      fmt.Sprintf("test tag <%s> 1", protovalidator.TagStringUnixCron),
			FieldDesc: "field 't_string_unix_cron'",
			TagInfo:   &protovalidator.TagInfo{Tag: protovalidator.TagStringUnixCron, Value: nil},
			FieldValue: func() interface{} {
				return data.OneTyp1.(*govalidatortest.ValidStringTagsOneOf1_TStringUnixCron).TStringUnixCron
			},
			BeforeFunc: func() {
				x := ""
				data.OneTyp1 = &govalidatortest.ValidStringTagsOneOf1_TStringUnixCron{TStringUnixCron: x}
			},
			AfterFunc: func() {
				x := "* * * * *"
				data.OneTyp1 = &govalidatortest.ValidStringTagsOneOf1_TStringUnixCron{TStringUnixCron: x}
			},
		},
		{
			Name:      fmt.Sprintf("test tag <%s> 1", protovalidator.TagStringEmail),
			FieldDesc: "field 't_string_email'",
			TagInfo:   &protovalidator.TagInfo{Tag: protovalidator.TagStringEmail, Value: nil},
			FieldValue: func() interface{} {
				return data.OneTyp1.(*govalidatortest.ValidStringTagsOneOf1_TStringEmail).TStringEmail
			},
			BeforeFunc: func() { x := ""; data.OneTyp1 = &govalidatortest.ValidStringTagsOneOf1_TStringEmail{TStringEmail: x} },
			AfterFunc: func() {
				x := "xxx@gmail.com"
				data.OneTyp1 = &govalidatortest.ValidStringTagsOneOf1_TStringEmail{TStringEmail: x}
			},
		},
		{
			Name:      fmt.Sprintf("test tag <%s> 1", protovalidator.TagStringJSON),
			FieldDesc: "field 't_string_json'",
			TagInfo:   &protovalidator.TagInfo{Tag: protovalidator.TagStringJSON, Value: nil},
			FieldValue: func() interface{} {
				return data.OneTyp1.(*govalidatortest.ValidStringTagsOneOf1_TStringJson).TStringJson
			},
			BeforeFunc: func() { x := ""; data.OneTyp1 = &govalidatortest.ValidStringTagsOneOf1_TStringJson{TStringJson: x} },
			AfterFunc: func() {
				x := `{"k1":"v1"}`
				data.OneTyp1 = &govalidatortest.ValidStringTagsOneOf1_TStringJson{TStringJson: x}
			},
		},
		{
			Name:      fmt.Sprintf("test tag <%s> 1", protovalidator.TagStringJWT),
			FieldDesc: "field 't_string_jwt'",
			TagInfo:   &protovalidator.TagInfo{Tag: protovalidator.TagStringJWT, Value: nil},
			FieldValue: func() interface{} {
				return data.OneTyp1.(*govalidatortest.ValidStringTagsOneOf1_TStringJwt).TStringJwt
			},
			BeforeFunc: func() { x := ""; data.OneTyp1 = &govalidatortest.ValidStringTagsOneOf1_TStringJwt{TStringJwt: x} },
			AfterFunc: func() {
				x := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiYWRtaW4iOnRydWV9.TJVA95OrM7E2cBab30RMHrHDcEfxjoYZgeFONFh7HgQ"
				data.OneTyp1 = &govalidatortest.ValidStringTagsOneOf1_TStringJwt{TStringJwt: x}
			},
		},
		{
			Name:      fmt.Sprintf("test tag <%s> 1", protovalidator.TagStringHTML),
			FieldDesc: "field 't_string_html'",
			TagInfo:   &protovalidator.TagInfo{Tag: protovalidator.TagStringHTML, Value: nil},
			FieldValue: func() interface{} {
				return data.OneTyp1.(*govalidatortest.ValidStringTagsOneOf1_TStringHtml).TStringHtml
			},
			BeforeFunc: func() { x := ""; data.OneTyp1 = &govalidatortest.ValidStringTagsOneOf1_TStringHtml{TStringHtml: x} },
			AfterFunc: func() {
				x := "<html></html>"
				data.OneTyp1 = &govalidatortest.ValidStringTagsOneOf1_TStringHtml{TStringHtml: x}
			},
		},
		{
			Name:      fmt.Sprintf("test tag <%s> 1", protovalidator.TagStringHTMLEncoded),
			FieldDesc: "field 't_string_html_encoded'",
			TagInfo:   &protovalidator.TagInfo{Tag: protovalidator.TagStringHTMLEncoded, Value: nil},
			FieldValue: func() interface{} {
				return data.OneTyp1.(*govalidatortest.ValidStringTagsOneOf1_TStringHtmlEncoded).TStringHtmlEncoded
			},
			BeforeFunc: func() {
				x := ""
				data.OneTyp1 = &govalidatortest.ValidStringTagsOneOf1_TStringHtmlEncoded{TStringHtmlEncoded: x}
			},
			AfterFunc: func() {
				x := "&lt;html&gt;&lt;/html&gt;"
				data.OneTyp1 = &govalidatortest.ValidStringTagsOneOf1_TStringHtmlEncoded{TStringHtmlEncoded: x}
			},
		},
		{
			Name:      fmt.Sprintf("test tag <%s> 1", protovalidator.TagStringBase64),
			FieldDesc: "field 't_string_base64'",
			TagInfo:   &protovalidator.TagInfo{Tag: protovalidator.TagStringBase64, Value: nil},
			FieldValue: func() interface{} {
				return data.OneTyp1.(*govalidatortest.ValidStringTagsOneOf1_TStringBase64).TStringBase64
			},
			BeforeFunc: func() {
				x := ""
				data.OneTyp1 = &govalidatortest.ValidStringTagsOneOf1_TStringBase64{TStringBase64: x}
			},
			AfterFunc: func() {
				x := "WkhPTkdHVU9uaWhhbzEyMw=="
				data.OneTyp1 = &govalidatortest.ValidStringTagsOneOf1_TStringBase64{TStringBase64: x}
			},
		},
		{
			Name:      fmt.Sprintf("test tag <%s> 1", protovalidator.TagStringBase64URL),
			FieldDesc: "field 't_string_base64_url'",
			TagInfo:   &protovalidator.TagInfo{Tag: protovalidator.TagStringBase64URL, Value: nil},
			FieldValue: func() interface{} {
				return data.OneTyp1.(*govalidatortest.ValidStringTagsOneOf1_TStringBase64Url).TStringBase64Url
			},
			BeforeFunc: func() {
				x := ""
				data.OneTyp1 = &govalidatortest.ValidStringTagsOneOf1_TStringBase64Url{TStringBase64Url: x}
			},
			AfterFunc: func() {
				x := "YWRmYWRzZg=="
				data.OneTyp1 = &govalidatortest.ValidStringTagsOneOf1_TStringBase64Url{TStringBase64Url: x}
			},
		},
		{
			Name:      fmt.Sprintf("test tag <%s> 1", protovalidator.TagStringHexadecimal),
			FieldDesc: "field 't_string_hexadecimal'",
			TagInfo:   &protovalidator.TagInfo{Tag: protovalidator.TagStringHexadecimal, Value: nil},
			FieldValue: func() interface{} {
				return data.OneTyp1.(*govalidatortest.ValidStringTagsOneOf1_TStringHexadecimal).TStringHexadecimal
			},
			BeforeFunc: func() {
				x := ""
				data.OneTyp1 = &govalidatortest.ValidStringTagsOneOf1_TStringHexadecimal{TStringHexadecimal: x}
			},
			AfterFunc: func() {
				x := "6461666164736661647366736461"
				data.OneTyp1 = &govalidatortest.ValidStringTagsOneOf1_TStringHexadecimal{TStringHexadecimal: x}
			},
		},
		{
			Name:      fmt.Sprintf("test tag <%s> 1", protovalidator.TagStringDatetime),
			FieldDesc: "field 't_string_datetime'",
			TagInfo:   &protovalidator.TagInfo{Tag: protovalidator.TagStringDatetime, Value: "2006-01-02 15:04:05"},
			FieldValue: func() interface{} {
				return data.OneTyp1.(*govalidatortest.ValidStringTagsOneOf1_TStringDatetime).TStringDatetime
			},
			BeforeFunc: func() {
				x := ""
				data.OneTyp1 = &govalidatortest.ValidStringTagsOneOf1_TStringDatetime{TStringDatetime: x}
			},
			AfterFunc: func() {
				x := "2020-11-02 15:04:05"
				data.OneTyp1 = &govalidatortest.ValidStringTagsOneOf1_TStringDatetime{TStringDatetime: x}
			},
		},
		{
			Name:      fmt.Sprintf("test tag <%s> 1", protovalidator.TagStringTimezone),
			FieldDesc: "field 't_string_timezone'",
			TagInfo:   &protovalidator.TagInfo{Tag: protovalidator.TagStringTimezone, Value: nil},
			FieldValue: func() interface{} {
				return data.OneTyp1.(*govalidatortest.ValidStringTagsOneOf1_TStringTimezone).TStringTimezone
			},
			BeforeFunc: func() {
				x := ""
				data.OneTyp1 = &govalidatortest.ValidStringTagsOneOf1_TStringTimezone{TStringTimezone: x}
			},
			AfterFunc: func() {
				x := "UTC"
				data.OneTyp1 = &govalidatortest.ValidStringTagsOneOf1_TStringTimezone{TStringTimezone: x}
			},
		},
		{
			Name:      fmt.Sprintf("test tag <%s> 1", protovalidator.TagStringUUID),
			FieldDesc: "field 't_string_uuid'",
			TagInfo:   &protovalidator.TagInfo{Tag: protovalidator.TagStringUUID, Value: nil},
			FieldValue: func() interface{} {
				return data.OneTyp1.(*govalidatortest.ValidStringTagsOneOf1_TStringUuid).TStringUuid
			},
			BeforeFunc: func() { x := ""; data.OneTyp1 = &govalidatortest.ValidStringTagsOneOf1_TStringUuid{TStringUuid: x} },
			AfterFunc: func() {
				x := "80363646-7361-11ec-b4f0-020017000b7b"
				data.OneTyp1 = &govalidatortest.ValidStringTagsOneOf1_TStringUuid{TStringUuid: x}
			},
		},
		{
			Name:      fmt.Sprintf("test tag <%s> 1", protovalidator.TagStringUUID1),
			FieldDesc: "field 't_string_uuid1'",
			TagInfo:   &protovalidator.TagInfo{Tag: protovalidator.TagStringUUID1, Value: nil},
			FieldValue: func() interface{} {
				return data.OneTyp1.(*govalidatortest.ValidStringTagsOneOf1_TStringUuid1).TStringUuid1
			},
			BeforeFunc: func() { x := ""; data.OneTyp1 = &govalidatortest.ValidStringTagsOneOf1_TStringUuid1{TStringUuid1: x} },
			AfterFunc: func() {
				x := "80363646-7361-11ec-b4f0-020017000b7b"
				data.OneTyp1 = &govalidatortest.ValidStringTagsOneOf1_TStringUuid1{TStringUuid1: x}
			},
		},
		{
			Name:      fmt.Sprintf("test tag <%s> 1", protovalidator.TagStringUUID3),
			FieldDesc: "field 't_string_uuid3'",
			TagInfo:   &protovalidator.TagInfo{Tag: protovalidator.TagStringUUID3, Value: nil},
			FieldValue: func() interface{} {
				return data.OneTyp1.(*govalidatortest.ValidStringTagsOneOf1_TStringUuid3).TStringUuid3
			},
			BeforeFunc: func() { x := ""; data.OneTyp1 = &govalidatortest.ValidStringTagsOneOf1_TStringUuid3{TStringUuid3: x} },
			AfterFunc: func() {
				x := "fa1f0699-6eac-3c47-b0f1-9d67ce5d0b76"
				data.OneTyp1 = &govalidatortest.ValidStringTagsOneOf1_TStringUuid3{TStringUuid3: x}
			},
		},
		{
			Name:      fmt.Sprintf("test tag <%s> 1", protovalidator.TagStringUUID4),
			FieldDesc: "field 't_string_uuid4'",
			TagInfo:   &protovalidator.TagInfo{Tag: protovalidator.TagStringUUID4, Value: nil},
			FieldValue: func() interface{} {
				return data.OneTyp1.(*govalidatortest.ValidStringTagsOneOf1_TStringUuid4).TStringUuid4
			},
			BeforeFunc: func() { x := ""; data.OneTyp1 = &govalidatortest.ValidStringTagsOneOf1_TStringUuid4{TStringUuid4: x} },
			AfterFunc: func() {
				x := "3af73fea-62c7-4a00-8314-942a1841288a"
				data.OneTyp1 = &govalidatortest.ValidStringTagsOneOf1_TStringUuid4{TStringUuid4: x}
			},
		},
		{
			Name:      fmt.Sprintf("test tag <%s> 1", protovalidator.TagStringUUID5),
			FieldDesc: "field 't_string_uuid5'",
			TagInfo:   &protovalidator.TagInfo{Tag: protovalidator.TagStringUUID5, Value: nil},
			FieldValue: func() interface{} {
				return data.OneTyp1.(*govalidatortest.ValidStringTagsOneOf1_TStringUuid5).TStringUuid5
			},
			BeforeFunc: func() { x := ""; data.OneTyp1 = &govalidatortest.ValidStringTagsOneOf1_TStringUuid5{TStringUuid5: x} },
			AfterFunc: func() {
				x := "78381e24-2d01-5a28-b2bf-9eedab6605ed"
				data.OneTyp1 = &govalidatortest.ValidStringTagsOneOf1_TStringUuid5{TStringUuid5: x}
			},
		},
	}

	var cases []*CaseDesc
	cases = append(cases, caseGeneral1...)
	cases = append(cases, caseGeneral2...)
	cases = append(cases, caseNetwork1...)
	cases = append(cases, caseFormat1...)

	msgName := "ValidStringTagsOneOf1"
	runCases(t, data, msgName, cases)

	{
		err := data.Validate()
		require.Nil(t, err)
	}
}

func Test_GoValidator_ValidOptionsMultiCond1(t *testing.T) {
	data := &govalidatortest.ValidOptionsMultiCond1{}
	{
		err := data.Validate()
		require.NotNil(t, err)
	}

	cases1 := []*CaseDesc{
		{
			Name:       fmt.Sprintf("test tag <%s> 1", protovalidator.TagStringByteLenGt),
			FieldDesc:  "field 't_basic_string1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagStringByteLenGt, Value: 5},
			FieldValue: func() interface{} { return len(data.TBasicString1) },
			BeforeFunc: func() {},
			AfterFunc:  func() { data.TBasicString1 = "1234567" },
		},
		{
			Name:       fmt.Sprintf("test tag <%s> 1", protovalidator.TagStringByteLenLt),
			FieldDesc:  "field 't_basic_string1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagStringByteLenLt, Value: 7},
			FieldValue: func() interface{} { return len(data.TBasicString1) },
			BeforeFunc: func() {},
			AfterFunc:  func() { data.TBasicString1 = "123456" },
		},
		{
			Name:       fmt.Sprintf("test tag <%s> 1", protovalidator.TagStringPrefix),
			FieldDesc:  "field 't_basic_string1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagStringPrefix, Value: "id-"},
			FieldValue: func() interface{} { return data.TBasicString1 },
			BeforeFunc: func() {},
			AfterFunc:  func() { data.TBasicString1 = "id-123" },
		},

		{
			Name:       fmt.Sprintf("test tag <%s> 2", protovalidator.TagStringByteLenGt),
			FieldDesc:  "field 't_basic_string2'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagStringByteLenGt, Value: 5},
			FieldValue: func() interface{} { return nil },
			BeforeFunc: func() { data.TBasicString2 = nil },
			AfterFunc:  func() { x := "1234567"; data.TBasicString2 = &x },
		},
		{
			Name:       fmt.Sprintf("test tag <%s> 2", protovalidator.TagStringByteLenLt),
			FieldDesc:  "field 't_basic_string2'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagStringByteLenLt, Value: 7},
			FieldValue: func() interface{} { return len(*data.TBasicString2) },
			BeforeFunc: func() {},
			AfterFunc:  func() { x := "123456"; data.TBasicString2 = &x },
		},
		{
			Name:       fmt.Sprintf("test tag <%s> 2", protovalidator.TagStringPrefix),
			FieldDesc:  "field 't_basic_string2'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagStringPrefix, Value: "id-"},
			FieldValue: func() interface{} { return *data.TBasicString2 },
			BeforeFunc: func() {},
			AfterFunc:  func() { x := "id-123"; data.TBasicString2 = &x },
		},

		{
			Name:       fmt.Sprintf("test tag <%s> 1", protovalidator.TagIntGte),
			FieldDesc:  "field 't_basic_int64'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagIntGte, Value: 1},
			FieldValue: func() interface{} { return data.TBasicInt64 },
			BeforeFunc: func() {},
			AfterFunc:  func() { x := int64(11); data.TBasicInt64 = x },
		},
		{
			Name:       fmt.Sprintf("test tag <%s> 1", protovalidator.TagIntLte),
			FieldDesc:  "field 't_basic_int64'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagIntLte, Value: 10},
			FieldValue: func() interface{} { return data.TBasicInt64 },
			BeforeFunc: func() {},
			AfterFunc:  func() { x := int64(10); data.TBasicInt64 = x },
		},

		{
			Name:       fmt.Sprintf("test tag <%s> 1", protovalidator.TagIntGte),
			FieldDesc:  "field 't_basic_int32'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagIntGte, Value: 1},
			FieldValue: func() interface{} { return data.TBasicInt32 },
			BeforeFunc: func() {},
			AfterFunc:  func() { x := int32(11); data.TBasicInt32 = &x },
		},
		{
			Name:       fmt.Sprintf("test tag <%s> 1", protovalidator.TagIntLte),
			FieldDesc:  "field 't_basic_int32'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagIntLte, Value: 10},
			FieldValue: func() interface{} { return data.TBasicInt32 },
			BeforeFunc: func() {},
			AfterFunc:  func() { x := int32(10); data.TBasicInt32 = &x },
		},
	}

	cases2 := []*CaseDesc{
		{
			Name:       fmt.Sprintf("test tag <%s> 1", protovalidator.TagRepeatedLenGte),
			FieldDesc:  "field 't_list_string'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagRepeatedLenGte, Value: 1},
			FieldValue: func() interface{} { return len(data.TListString) },
			BeforeFunc: func() {},
			AfterFunc:  func() { data.TListString = []string{"s1", "s2", "s3", "s4"} },
		},
		{
			Name:       fmt.Sprintf("test tag <%s> 1", protovalidator.TagRepeatedLenLte),
			FieldDesc:  "field 't_list_string'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagRepeatedLenLte, Value: 3},
			FieldValue: func() interface{} { return len(data.TListString) },
			BeforeFunc: func() {},
			AfterFunc:  func() { data.TListString = []string{"1", "2", "3"} },
		},
		{
			Name:       fmt.Sprintf("test tag <%s> 1", protovalidator.TagStringByteLenGt),
			FieldDesc:  "array item where in field 't_list_string'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagStringByteLenGt, Value: 1},
			FieldValue: func() interface{} { return 1 },
			BeforeFunc: func() {},
			AfterFunc:  func() { data.TListString = []string{"12345", "12345", "12345"} },
		},
		{
			Name:       fmt.Sprintf("test tag <%s> 1", protovalidator.TagStringByteLenLt),
			FieldDesc:  "array item where in field 't_list_string'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagStringByteLenLt, Value: 5},
			FieldValue: func() interface{} { return 5 },
			BeforeFunc: func() {},
			AfterFunc:  func() { data.TListString = []string{"12", "12", "13"} },
		},
		{
			Name:       fmt.Sprintf("test tag <%s> 1", protovalidator.TagStringPrefix),
			FieldDesc:  "array item where in field 't_list_string'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagStringPrefix, Value: "id-"},
			FieldValue: func() interface{} { return "12" },
			BeforeFunc: func() {},
			AfterFunc:  func() { data.TListString = []string{"id-1", "id-2", "id-3"} },
		},

		{
			Name:       fmt.Sprintf("test tag <%s> 1", protovalidator.TagRepeatedLenGte),
			FieldDesc:  "field 't_list_int64'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagRepeatedLenGte, Value: 1},
			FieldValue: func() interface{} { return len(data.TListInt64) },
			BeforeFunc: func() {},
			AfterFunc:  func() { data.TListInt64 = []int64{1, 2, 3, 4} },
		},
		{
			Name:       fmt.Sprintf("test tag <%s> 1", protovalidator.TagRepeatedLenLte),
			FieldDesc:  "field 't_list_int64'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagRepeatedLenLte, Value: 3},
			FieldValue: func() interface{} { return len(data.TListInt64) },
			BeforeFunc: func() {},
			AfterFunc:  func() { data.TListInt64 = []int64{1, 2, 3} },
		},
		{
			Name:       fmt.Sprintf("test tag <%s> 1", protovalidator.TagIntGte),
			FieldDesc:  "array item where in field 't_list_int64'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagIntGte, Value: 1},
			FieldValue: func() interface{} { return 0 },
			BeforeFunc: func() { data.TListInt64 = []int64{0} },
			AfterFunc:  func() { data.TListInt64 = []int64{11, 12, 13} },
		},
		{
			Name:       fmt.Sprintf("test tag <%s> 1", protovalidator.TagIntLte),
			FieldDesc:  "array item where in field 't_list_int64'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagIntLte, Value: 10},
			FieldValue: func() interface{} { return 11 },
			BeforeFunc: func() {},
			AfterFunc:  func() { data.TListInt64 = []int64{5, 2, 3} },
		},
	}

	cases3 := []*CaseDesc{
		{
			Name:       fmt.Sprintf("test tag <%s> 1", protovalidator.TagMapLenGte),
			FieldDesc:  "field 't_map_string1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagMapLenGte, Value: 1},
			FieldValue: func() interface{} { return len(data.TMapString1) },
			BeforeFunc: func() {},
			AfterFunc: func() {
				data.TMapString1 = map[string]string{"k1": "v1", "k2": "v2", "k3": "v3", "k4": "v4"}
			},
		},
		{
			Name:       fmt.Sprintf("test tag <%s> 2", protovalidator.TagMapLenLte),
			FieldDesc:  "field 't_map_string1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagMapLenLte, Value: 3},
			FieldValue: func() interface{} { return len(data.TMapString1) },
			BeforeFunc: func() {},
			AfterFunc:  func() { data.TMapString1 = map[string]string{"k1": "v1", "k2": "v2", "k3": "v3"} },
		},

		{
			Name:       fmt.Sprintf("test tag <%s> 3", protovalidator.TagStringByteLenGt),
			FieldDesc:  "map key where in field 't_map_string1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagStringByteLenGt, Value: 5},
			FieldValue: func() interface{} { return 2 },
			BeforeFunc: func() {},
			AfterFunc:  func() { data.TMapString1 = map[string]string{"k1234563": "v1", "k1234561": "v2", "k1234562": "v3"} },
		},
		{
			Name:       fmt.Sprintf("test tag <%s> 4", protovalidator.TagStringByteLenLt),
			FieldDesc:  "map key where in field 't_map_string1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagStringByteLenLt, Value: 7},
			FieldValue: func() interface{} { return 8 },
			BeforeFunc: func() {},
			AfterFunc:  func() { data.TMapString1 = map[string]string{"k12341": "v1", "id-122": "v2", "id-123": "v3"} },
		},
		{
			Name:       fmt.Sprintf("test tag <%s> 5", protovalidator.TagStringPrefix),
			FieldDesc:  "map key where in field 't_map_string1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagStringPrefix, Value: "id-"},
			FieldValue: func() interface{} { return "k12341" },
			BeforeFunc: func() {},
			AfterFunc:  func() { data.TMapString1 = map[string]string{"id-121": "v1", "id-122": "v2", "id-123": "v3"} },
		},

		{
			Name:       fmt.Sprintf("test tag <%s> 6", protovalidator.TagStringByteLenGt),
			FieldDesc:  "map value where in field 't_map_string1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagStringByteLenGt, Value: 5},
			FieldValue: func() interface{} { return 2 },
			BeforeFunc: func() {},
			AfterFunc: func() {
				data.TMapString1 = map[string]string{"id-121": "k1234563", "id-122": "k1234563", "id-123": "k1234563"}
			},
		},
		{
			Name:       fmt.Sprintf("test tag <%s> 7", protovalidator.TagStringByteLenLt),
			FieldDesc:  "map value where in field 't_map_string1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagStringByteLenLt, Value: 7},
			FieldValue: func() interface{} { return 8 },
			BeforeFunc: func() {},
			AfterFunc: func() {
				data.TMapString1 = map[string]string{"id-121": "k12341", "id-122": "id-121", "id-123": "id-121"}
			},
		},
		{
			Name:       fmt.Sprintf("test tag <%s> 8", protovalidator.TagStringPrefix),
			FieldDesc:  "map value where in field 't_map_string1'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagStringPrefix, Value: "id-"},
			FieldValue: func() interface{} { return "k12341" },
			BeforeFunc: func() {},
			AfterFunc: func() {
				data.TMapString1 = map[string]string{"id-121": "id-121", "id-122": "id-121", "id-123": "id-121"}
			},
		},
	}

	cases4 := []*CaseDesc{
		{
			Name:       fmt.Sprintf("test tag <%s> 2", protovalidator.TagMapLenGte),
			FieldDesc:  "field 't_map_int64'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagMapLenGte, Value: 1},
			FieldValue: func() interface{} { return len(data.TMapInt64) },
			BeforeFunc: func() {},
			AfterFunc:  func() { data.TMapInt64 = map[int64]int64{0: 0, 2: 2, 3: 3, 4: 4} },
		},
		{
			Name:       fmt.Sprintf("test tag <%s> 2", protovalidator.TagMapLenLte),
			FieldDesc:  "field 't_map_int64'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagMapLenLte, Value: 3},
			FieldValue: func() interface{} { return len(data.TMapInt64) },
			BeforeFunc: func() {},
			AfterFunc:  func() { data.TMapInt64 = map[int64]int64{1: 0, 6: 2, 7: 3} },
		},

		{
			Name:       fmt.Sprintf("test tag <%s> 3", protovalidator.TagIntGt),
			FieldDesc:  "map key where in field 't_map_int64'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagIntGt, Value: 5},
			FieldValue: func() interface{} { return 1 },
			BeforeFunc: func() {},
			AfterFunc:  func() { data.TMapInt64 = map[int64]int64{6: 0, 7: 2, 13: 3} },
		},
		{
			Name:       fmt.Sprintf("test tag <%s> 4", protovalidator.TagIntLt),
			FieldDesc:  "map key where in field 't_map_int64'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagIntLt, Value: 10},
			FieldValue: func() interface{} { return 13 },
			BeforeFunc: func() {},
			AfterFunc:  func() { data.TMapInt64 = map[int64]int64{6: 3, 7: 7, 8: 8} },
		},

		{
			Name:       fmt.Sprintf("test tag <%s> 6", protovalidator.TagIntGt),
			FieldDesc:  "map value where in field 't_map_int64'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagIntGt, Value: 5},
			FieldValue: func() interface{} { return 3 },
			BeforeFunc: func() {},
			AfterFunc:  func() { data.TMapInt64 = map[int64]int64{6: 10, 7: 6, 8: 8} },
		},
		{
			Name:       fmt.Sprintf("test tag <%s> 7", protovalidator.TagIntLt),
			FieldDesc:  "map value where in field 't_map_int64'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagIntLt, Value: 10},
			FieldValue: func() interface{} { return 10 },
			BeforeFunc: func() {},
			AfterFunc:  func() { data.TMapInt64 = map[int64]int64{6: 6, 7: 7, 8: 8} },
		},
	}

	cases5 := []*CaseDesc{
		{
			Name:      fmt.Sprintf("test tag <%s> 1", protovalidator.TagStringByteLenGt),
			FieldDesc: "field 'oneof_string1'",
			TagInfo:   &protovalidator.TagInfo{Tag: protovalidator.TagStringByteLenGt, Value: 5},
			FieldValue: func() interface{} {
				x := data.OneTyp1.(*govalidatortest.ValidOptionsMultiCond1_OneofString1).OneofString1
				return len(x)
			},
			BeforeFunc: func() {
				data.OneTyp1 = &govalidatortest.ValidOptionsMultiCond1_OneofString1{OneofString1: "1234"}
			},
			AfterFunc: func() {
				data.OneTyp1 = &govalidatortest.ValidOptionsMultiCond1_OneofString1{OneofString1: "1234567"}
			},
		},
		{
			Name:      fmt.Sprintf("test tag <%s> 1", protovalidator.TagStringByteLenLt),
			FieldDesc: "field 'oneof_string1'",
			TagInfo:   &protovalidator.TagInfo{Tag: protovalidator.TagStringByteLenLt, Value: 7},
			FieldValue: func() interface{} {
				x := data.OneTyp1.(*govalidatortest.ValidOptionsMultiCond1_OneofString1).OneofString1
				return len(x)
			},
			BeforeFunc: func() {
			},
			AfterFunc: func() {
				data.OneTyp1 = &govalidatortest.ValidOptionsMultiCond1_OneofString1{OneofString1: "123456"}
			},
		},
		{
			Name:      fmt.Sprintf("test tag <%s> 1", protovalidator.TagStringPrefix),
			FieldDesc: "field 'oneof_string1'",
			TagInfo:   &protovalidator.TagInfo{Tag: protovalidator.TagStringPrefix, Value: "id-"},
			FieldValue: func() interface{} {
				x := data.OneTyp1.(*govalidatortest.ValidOptionsMultiCond1_OneofString1).OneofString1
				return x
			},
			BeforeFunc: func() {},
			AfterFunc: func() {
				data.OneTyp1 = &govalidatortest.ValidOptionsMultiCond1_OneofString1{OneofString1: "id-123"}
			},
		},
	}

	cases6 := []*CaseDesc{
		{
			Name:      fmt.Sprintf("test tag <%s> 1", protovalidator.TagIntGte),
			FieldDesc: "field 'oneof_int64'",
			TagInfo:   &protovalidator.TagInfo{Tag: protovalidator.TagIntGte, Value: 1},
			FieldValue: func() interface{} {
				x := data.OneTyp1.(*govalidatortest.ValidOptionsMultiCond1_OneofInt64).OneofInt64
				return x
			},
			BeforeFunc: func() {
				data.OneTyp1 = &govalidatortest.ValidOptionsMultiCond1_OneofInt64{OneofInt64: 0}
			},
			AfterFunc: func() {
				data.OneTyp1 = &govalidatortest.ValidOptionsMultiCond1_OneofInt64{OneofInt64: 11}
			},
		},
		{
			Name:      fmt.Sprintf("test tag <%s> 1", protovalidator.TagIntLte),
			FieldDesc: "field 'oneof_int64'",
			TagInfo:   &protovalidator.TagInfo{Tag: protovalidator.TagIntLte, Value: 10},
			FieldValue: func() interface{} {
				x := data.OneTyp1.(*govalidatortest.ValidOptionsMultiCond1_OneofInt64).OneofInt64
				return x
			},
			BeforeFunc: func() {
			},
			AfterFunc: func() {
				data.OneTyp1 = &govalidatortest.ValidOptionsMultiCond1_OneofInt64{OneofInt64: 5}
			},
		},
	}

	cases7 := []*CaseDesc{
		{
			Name:       fmt.Sprintf("test tag <%s> 1", protovalidator.TagMapLenGte),
			FieldDesc:  "field 't_map_string2'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagMapLenGte, Value: 1},
			FieldValue: func() interface{} { return len(data.TMapString2) },
			BeforeFunc: func() {},
			AfterFunc: func() {
				data.TMapString2 = map[string]string{"k1": "v1", "k2": "v2", "k3": "v3", "k4": "v4"}
			},
		},
		{
			Name:       fmt.Sprintf("test tag <%s> 2", protovalidator.TagMapLenLte),
			FieldDesc:  "field 't_map_string2'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagMapLenLte, Value: 3},
			FieldValue: func() interface{} { return len(data.TMapString2) },
			BeforeFunc: func() {},
			AfterFunc:  func() { data.TMapString2 = map[string]string{"a1": "v1", "a2": "v2", "a4": "v3"} },
		},
		{
			Name:       fmt.Sprintf("test tag <%s> 3", protovalidator.TagStringIn),
			FieldDesc:  "map key where in field 't_map_string2'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagStringIn, Value: []string{"a1", "a2", "a3"}},
			FieldValue: func() interface{} { return "a4" },
			BeforeFunc: func() {},
			AfterFunc:  func() { data.TMapString2 = map[string]string{"a1": "v1", "a3": "v3"} },
		},
		{
			Name:       fmt.Sprintf("test tag <%s> 4", protovalidator.TagStringNotIn),
			FieldDesc:  "map key where in field 't_map_string2'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagStringNotIn, Value: []string{"a3", "a4", "a5"}},
			FieldValue: func() interface{} { return "a3" },
			BeforeFunc: func() {},
			AfterFunc:  func() { data.TMapString2 = map[string]string{"a1": "a1", "a2": "a4"} },
		},

		{
			Name:       fmt.Sprintf("test tag <%s> 5", protovalidator.TagStringIn),
			FieldDesc:  "map value where in field 't_map_string2'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagStringIn, Value: []string{"a1", "a2", "a3"}},
			FieldValue: func() interface{} { return "a4" },
			BeforeFunc: func() {},
			AfterFunc:  func() { data.TMapString2 = map[string]string{"a1": "a1", "a2": "a3"} },
		},
		{
			Name:       fmt.Sprintf("test tag <%s> 6", protovalidator.TagStringNotIn),
			FieldDesc:  "map value where in field 't_map_string2'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagStringNotIn, Value: []string{"a3", "a4", "a5"}},
			FieldValue: func() interface{} { return "a3" },
			BeforeFunc: func() {},
			AfterFunc:  func() { data.TMapString2 = map[string]string{"a1": "a1", "a2": "a2"} },
		},
	}

	var cases []*CaseDesc
	cases = append(cases, cases1...)
	cases = append(cases, cases2...)
	cases = append(cases, cases3...)
	cases = append(cases, cases4...)
	cases = append(cases, cases5...)
	cases = append(cases, cases6...)
	cases = append(cases, cases7...)

	msgName := "ValidOptionsMultiCond1"
	runCases(t, data, msgName, cases)

	{
		err := data.Validate()
		require.Nil(t, err)
	}
}

func Test_GoValidator_CheckIfOptions1(t *testing.T) {
	data := &govalidatortest.CheckIfOptions1{}
	{
		err := data.Validate()
		require.Nil(t, err)
	}

	cases := []*CaseDesc{
		{
			Name:       fmt.Sprintf("test check_if tag <%s> 1", protovalidator.TagOneOfNotNull),
			FieldDesc:  "field 't_basic_string'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagStringByteLenEq, Value: 3},
			FieldValue: func() interface{} { return len(data.TBasicString) },
			BeforeFunc: func() { data.OneofType1 = &govalidatortest.CheckIfOptions1_Oneof1String1{Oneof1String1: "11"} },
			AfterFunc:  func() { data.OneofType1 = nil; data.TBasicString = "abc" },
		},
		{
			Name:       fmt.Sprintf("test check_if tag <%s> 2", protovalidator.TagOneOfNotNull),
			FieldDesc:  "field 't_list_string'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagRepeatedLenEq, Value: 3},
			FieldValue: func() interface{} { return len(data.TListString) },
			BeforeFunc: func() { data.OneofType1 = &govalidatortest.CheckIfOptions1_Oneof1String1{Oneof1String1: "11"} },
			AfterFunc:  func() { data.OneofType1 = nil; data.TListString = []string{"a", "b", "c"} },
		},
		{
			Name:       fmt.Sprintf("test check_if tag <%s> 3", protovalidator.TagOneOfNotNull),
			FieldDesc:  "field 't_map_string'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagMapLenEq, Value: 3},
			FieldValue: func() interface{} { return len(data.TMapString) },
			BeforeFunc: func() { data.OneofType1 = &govalidatortest.CheckIfOptions1_Oneof1String1{Oneof1String1: "11"} },
			AfterFunc:  func() { data.OneofType1 = nil; data.TMapString = map[string]string{"k1": "v1", "k2": "v2", "k3": "v3"} },
		},
		{
			Name:       fmt.Sprintf("test check_if tag <%s> 4", protovalidator.TagOneOfNotNull),
			FieldDesc:  "field 'oneof_type2'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagOneOfNotNull, Value: nil},
			FieldValue: func() interface{} { return nil },
			BeforeFunc: func() { data.OneofType1 = &govalidatortest.CheckIfOptions1_Oneof1String1{Oneof1String1: "11"} },
			AfterFunc: func() {
				data.OneofType1 = nil
				data.OneofType2 = &govalidatortest.CheckIfOptions1_Oneof2String{Oneof2String: ""}
			},
			UseError2: true,
		},
		{
			Name:      fmt.Sprintf("test check_if tag <%s> 5", protovalidator.TagOneOfNotNull),
			FieldDesc: "field 'oneof2_string'",
			TagInfo:   &protovalidator.TagInfo{Tag: protovalidator.TagStringByteLenEq, Value: 3},
			FieldValue: func() interface{} {
				x := data.OneofType2.(*govalidatortest.CheckIfOptions1_Oneof2String).Oneof2String
				return len(x)
			},
			BeforeFunc: func() { data.OneofType1 = &govalidatortest.CheckIfOptions1_Oneof1String1{Oneof1String1: "11"} },
			AfterFunc: func() {
				data.OneofType1 = nil
				data.OneofType2 = &govalidatortest.CheckIfOptions1_Oneof2String{Oneof2String: "abc"}
			},
		},
	}

	msgName := "CheckIfOptions1"
	runCases(t, data, msgName, cases)

	{
		err := data.Validate()
		require.Nil(t, err)
	}
}

func Test_GoValidator_CheckIfOptions2(t *testing.T) {
	data := &govalidatortest.CheckIfOptions2{}
	{
		err := data.Validate()
		require.Nil(t, err)
	}

	cases := []*CaseDesc{
		{
			Name:       fmt.Sprintf("test check_if tag <%s> 1", protovalidator.TagOneOfNotNull),
			FieldDesc:  "field 't_basic_string'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagStringByteLenEq, Value: 3},
			FieldValue: func() interface{} { return len(data.TBasicString) },
			BeforeFunc: func() { data.OneofType1 = &govalidatortest.CheckIfOptions2_Oneof1String1{Oneof1String1: "check"} },
			AfterFunc:  func() { data.OneofType1 = nil; data.TBasicString = "abc" },
		},
		{
			Name:       fmt.Sprintf("test check_if tag <%s> 2", protovalidator.TagOneOfNotNull),
			FieldDesc:  "field 't_list_string'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagRepeatedLenEq, Value: 3},
			FieldValue: func() interface{} { return len(data.TListString) },
			BeforeFunc: func() { data.OneofType1 = &govalidatortest.CheckIfOptions2_Oneof1String1{Oneof1String1: "check"} },
			AfterFunc:  func() { data.OneofType1 = nil; data.TListString = []string{"a", "b", "c"} },
		},
		{
			Name:       fmt.Sprintf("test check_if tag <%s> 3", protovalidator.TagOneOfNotNull),
			FieldDesc:  "field 't_map_string'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagMapLenEq, Value: 3},
			FieldValue: func() interface{} { return len(data.TMapString) },
			BeforeFunc: func() { data.OneofType1 = &govalidatortest.CheckIfOptions2_Oneof1String1{Oneof1String1: "check"} },
			AfterFunc:  func() { data.OneofType1 = nil; data.TMapString = map[string]string{"k1": "v1", "k2": "v2", "k3": "v3"} },
		},
		{
			Name:       fmt.Sprintf("test check_if tag <%s> 4", protovalidator.TagOneOfNotNull),
			FieldDesc:  "field 'oneof_type2'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagOneOfNotNull, Value: nil},
			FieldValue: func() interface{} { return nil },
			BeforeFunc: func() { data.OneofType1 = &govalidatortest.CheckIfOptions2_Oneof1String1{Oneof1String1: "check"} },
			AfterFunc: func() {
				data.OneofType1 = nil
				data.OneofType2 = &govalidatortest.CheckIfOptions2_Oneof2String{Oneof2String: ""}
			},
			UseError2: true,
		},
		{
			Name:      fmt.Sprintf("test check_if tag <%s> 5", protovalidator.TagOneOfNotNull),
			FieldDesc: "field 'oneof2_string'",
			TagInfo:   &protovalidator.TagInfo{Tag: protovalidator.TagStringByteLenEq, Value: 3},
			FieldValue: func() interface{} {
				x := data.OneofType2.(*govalidatortest.CheckIfOptions2_Oneof2String).Oneof2String
				return len(x)
			},
			BeforeFunc: func() { data.OneofType1 = &govalidatortest.CheckIfOptions2_Oneof1String1{Oneof1String1: "check"} },
			AfterFunc: func() {
				data.OneofType1 = nil
				data.OneofType2 = &govalidatortest.CheckIfOptions2_Oneof2String{Oneof2String: "abc"}
			},
		},
	}

	msgName := "CheckIfOptions2"
	runCases(t, data, msgName, cases)

	{
		err := data.Validate()
		require.Nil(t, err)
	}
}

func Test_GoValidator_CheckIfOptions3(t *testing.T) {
	data := &govalidatortest.CheckIfOptions3{}
	{
		err := data.Validate()
		require.Nil(t, err)
	}

	cases := []*CaseDesc{
		{
			Name:       fmt.Sprintf("test check_if tag <%s> 1", protovalidator.TagOneOfNotNull),
			FieldDesc:  "field 't_basic_string'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagStringByteLenEq, Value: 3},
			FieldValue: func() interface{} { return len(data.TBasicString) },
			BeforeFunc: func() { data.SeedString1 = "check" },
			AfterFunc:  func() { data.TBasicString = "abc" },
		},
		{
			Name:       fmt.Sprintf("test check_if tag <%s> 2", protovalidator.TagOneOfNotNull),
			FieldDesc:  "field 't_list_string'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagRepeatedLenEq, Value: 3},
			FieldValue: func() interface{} { return len(data.TListString) },
			BeforeFunc: func() { data.SeedString1 = "check" },
			AfterFunc:  func() { data.SeedString1 = ""; data.TListString = []string{"a", "b", "c"} },
		},
		{
			Name:       fmt.Sprintf("test check_if tag <%s> 3", protovalidator.TagOneOfNotNull),
			FieldDesc:  "field 't_map_string'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagMapLenEq, Value: 3},
			FieldValue: func() interface{} { return len(data.TMapString) },
			BeforeFunc: func() { data.SeedString1 = "check" },
			AfterFunc:  func() { data.SeedString1 = ""; data.TMapString = map[string]string{"k1": "v1", "k2": "v2", "k3": "v3"} },
		},
		{
			Name:       fmt.Sprintf("test check_if tag <%s> 4", protovalidator.TagOneOfNotNull),
			FieldDesc:  "field 'oneof_type2'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagOneOfNotNull, Value: nil},
			FieldValue: func() interface{} { return nil },
			BeforeFunc: func() { data.SeedString1 = "check" },
			AfterFunc: func() {
				data.SeedString1 = ""
				data.OneofType2 = &govalidatortest.CheckIfOptions3_Oneof2String{Oneof2String: ""}
			},
			UseError2: true,
		},
		{
			Name:      fmt.Sprintf("test check_if tag <%s> 5", protovalidator.TagOneOfNotNull),
			FieldDesc: "field 'oneof2_string'",
			TagInfo:   &protovalidator.TagInfo{Tag: protovalidator.TagStringByteLenEq, Value: 3},
			FieldValue: func() interface{} {
				x := data.OneofType2.(*govalidatortest.CheckIfOptions3_Oneof2String).Oneof2String
				return len(x)
			},
			BeforeFunc: func() { data.SeedString1 = "check" },
			AfterFunc: func() {
				data.SeedString1 = ""
				data.OneofType2 = &govalidatortest.CheckIfOptions3_Oneof2String{Oneof2String: "abc"}
			},
		},
	}

	msgName := "CheckIfOptions3"
	runCases(t, data, msgName, cases)

	{
		err := data.Validate()
		require.Nil(t, err)
	}
}

func Test_GoValidator_CheckIfOptions4(t *testing.T) {
	data := &govalidatortest.CheckIfOptions4{}
	{
		err := data.Validate()
		require.Nil(t, err)
	}

	cases := []*CaseDesc{
		{
			Name:       fmt.Sprintf("test check_if tag <%s> 1", protovalidator.TagOneOfNotNull),
			FieldDesc:  "field 't_basic_string'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagStringByteLenEq, Value: 3},
			FieldValue: func() interface{} { return len(data.TBasicString) },
			BeforeFunc: func() { data.OneofType1 = &govalidatortest.CheckIfOptions4_Oneof1String1{Oneof1String1: "1234"} },
			AfterFunc:  func() { data.OneofType1 = nil; data.TBasicString = "abc" },
		},
		{
			Name:       fmt.Sprintf("test check_if tag <%s> 2", protovalidator.TagOneOfNotNull),
			FieldDesc:  "field 't_list_string'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagRepeatedLenEq, Value: 3},
			FieldValue: func() interface{} { return len(data.TListString) },
			BeforeFunc: func() { data.OneofType1 = &govalidatortest.CheckIfOptions4_Oneof1String1{Oneof1String1: "1234"} },
			AfterFunc:  func() { data.OneofType1 = nil; data.TListString = []string{"a", "b", "c"} },
		},
		{
			Name:       fmt.Sprintf("test check_if tag <%s> 3", protovalidator.TagOneOfNotNull),
			FieldDesc:  "field 't_map_string'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagMapLenEq, Value: 3},
			FieldValue: func() interface{} { return len(data.TMapString) },
			BeforeFunc: func() { data.OneofType1 = &govalidatortest.CheckIfOptions4_Oneof1String1{Oneof1String1: "1234"} },
			AfterFunc:  func() { data.OneofType1 = nil; data.TMapString = map[string]string{"k1": "v1", "k2": "v2", "k3": "v3"} },
		},
		{
			Name:       fmt.Sprintf("test check_if tag <%s> 4", protovalidator.TagOneOfNotNull),
			FieldDesc:  "field 'oneof_type2'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagOneOfNotNull, Value: nil},
			FieldValue: func() interface{} { return nil },
			BeforeFunc: func() { data.OneofType1 = &govalidatortest.CheckIfOptions4_Oneof1String1{Oneof1String1: "1234"} },
			AfterFunc: func() {
				data.OneofType1 = nil
				data.OneofType2 = &govalidatortest.CheckIfOptions4_Oneof2String{Oneof2String: ""}
			},
			UseError2: true,
		},
		{
			Name:      fmt.Sprintf("test check_if tag <%s> 3", protovalidator.TagOneOfNotNull),
			FieldDesc: "field 'oneof2_string'",
			TagInfo:   &protovalidator.TagInfo{Tag: protovalidator.TagStringByteLenEq, Value: 3},
			FieldValue: func() interface{} {
				x := data.OneofType2.(*govalidatortest.CheckIfOptions4_Oneof2String).Oneof2String
				return len(x)
			},
			BeforeFunc: func() { data.OneofType1 = &govalidatortest.CheckIfOptions4_Oneof1String1{Oneof1String1: "1234"} },
			AfterFunc: func() {
				data.OneofType1 = nil
				data.OneofType2 = &govalidatortest.CheckIfOptions4_Oneof2String{Oneof2String: "abc"}
			},
		},
	}

	msgName := "CheckIfOptions4"
	runCases(t, data, msgName, cases)

	{
		err := data.Validate()
		require.Nil(t, err)
	}
}

func Test_GoValidator_CheckIfOptions5(t *testing.T) {
	data := &govalidatortest.CheckIfOptions5{}
	{
		err := data.Validate()
		require.Nil(t, err)
	}

	cases := []*CaseDesc{
		{
			Name:       fmt.Sprintf("test check_if tag <%s> 1", protovalidator.TagOneOfNotNull),
			FieldDesc:  "field 't_basic_string'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagStringByteLenEq, Value: 3},
			FieldValue: func() interface{} { return len(data.TBasicString) },
			BeforeFunc: func() { data.SeedString1 = "1234" },
			AfterFunc:  func() { data.TBasicString = "abc" },
		},
		{
			Name:       fmt.Sprintf("test check_if tag <%s> 2", protovalidator.TagOneOfNotNull),
			FieldDesc:  "field 't_list_string'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagRepeatedLenEq, Value: 3},
			FieldValue: func() interface{} { return len(data.TListString) },
			BeforeFunc: func() { data.SeedString1 = "1234" },
			AfterFunc:  func() { data.SeedString1 = ""; data.TListString = []string{"a", "b", "c"} },
		},
		{
			Name:       fmt.Sprintf("test check_if tag <%s> 3", protovalidator.TagOneOfNotNull),
			FieldDesc:  "field 't_map_string'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagMapLenEq, Value: 3},
			FieldValue: func() interface{} { return len(data.TMapString) },
			BeforeFunc: func() { data.SeedString1 = "1234" },
			AfterFunc:  func() { data.SeedString1 = ""; data.TMapString = map[string]string{"k1": "v1", "k2": "v2", "k3": "v3"} },
		},
		{
			Name:       fmt.Sprintf("test check_if tag <%s> 4", protovalidator.TagOneOfNotNull),
			FieldDesc:  "field 'oneof_type2'",
			TagInfo:    &protovalidator.TagInfo{Tag: protovalidator.TagOneOfNotNull, Value: nil},
			FieldValue: func() interface{} { return nil },
			BeforeFunc: func() { data.SeedString1 = "1234" },
			AfterFunc: func() {
				data.SeedString1 = ""
				data.OneofType2 = &govalidatortest.CheckIfOptions5_Oneof2String{Oneof2String: ""}
			},
			UseError2: true,
		},
		{
			Name:      fmt.Sprintf("test check_if tag <%s> 5", protovalidator.TagOneOfNotNull),
			FieldDesc: "field 'oneof2_string'",
			TagInfo:   &protovalidator.TagInfo{Tag: protovalidator.TagStringByteLenEq, Value: 3},
			FieldValue: func() interface{} {
				x := data.OneofType2.(*govalidatortest.CheckIfOptions5_Oneof2String).Oneof2String
				return len(x)
			},
			BeforeFunc: func() { data.SeedString1 = "1234" },
			AfterFunc: func() {
				data.SeedString1 = ""
				data.OneofType2 = &govalidatortest.CheckIfOptions5_Oneof2String{Oneof2String: "abc"}
			},
		},
	}

	msgName := "CheckIfOptions5"
	runCases(t, data, msgName, cases)

	{
		err := data.Validate()
		require.Nil(t, err)
	}
}
