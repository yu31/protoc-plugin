// Code generated by protoc-gen-govalidator. DO NOT EDIT.
// versions:
// 		protoc-gen-govalidator 0.0.1
// source: xgo/tests/govalidatorexternal/test_invalid_map.proto

package govalidatorexternal

import (
	_ "github.com/yu31/protoc-plugin/xgo/pb/pbvalidator"
	protovalidator "github.com/yu31/protoc-plugin/xgo/pkg/protovalidator"
)

func (this *InvalidMessageMap) _xxx_xxx_Validator_Validate_t_map_string() error {
	if !(this.TMapString != nil) {
		return protovalidator.FieldError2("InvalidMessageMap", "the value of field 't_map_string' cannot be null")
	}
	return nil
}

// Set default value for message godefaultstest.InvalidMessageMap
func (this *InvalidMessageMap) Validate() error {
	if this == nil {
		return nil
	}
	if err := this._xxx_xxx_Validator_Validate_t_map_string(); err != nil {
		return err
	}
	return nil
}
