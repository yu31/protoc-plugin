// Code generated by protoc-gen-govalidator. DO NOT EDIT.
// versions:
// 		protoc-gen-govalidator 0.0.1
// source: xgo/tests/govalidatorexternal/test_invalid_list.proto

package govalidatorexternal

import (
	_ "github.com/yu31/protoc-plugin/xgo/pb/pbvalidator"
	protovalidator "github.com/yu31/protoc-plugin/xgo/pkg/protovalidator"
)

func (this *InvalidMessageList) _xxx_xxx_Validator_Validate_t_list_string1() error {
	if !(this.TListString1 != nil) {
		return protovalidator.FieldError2("InvalidMessageList", "the value of field 't_list_string1' cannot be null")
	}
	return nil
}

// Set default value for message godefaultstest.InvalidMessageList
func (this *InvalidMessageList) Validate() error {
	if this == nil {
		return nil
	}
	if err := this._xxx_xxx_Validator_Validate_t_list_string1(); err != nil {
		return err
	}
	return nil
}