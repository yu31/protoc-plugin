// Code generated by protoc-gen-govalidator. DO NOT EDIT.
// versions:
// 		protoc-gen-govalidator 0.0.1
// source: xgo/tests/govalidatorexternal/test_invalid_bool.proto

package govalidatorexternal

import (
	_ "github.com/yu31/protoc-plugin/xgo/pb/pbvalidator"
)

func (this *InvalidMessageBool) _xxx_xxx_Validator_Validate_t_field1() error {
	return nil
}

// Set default value for message godefaultstest.InvalidMessageBool
func (this *InvalidMessageBool) Validate() error {
	if this == nil {
		return nil
	}
	if err := this._xxx_xxx_Validator_Validate_t_field1(); err != nil {
		return err
	}
	return nil
}
