// Code generated by protoc-gen-godefaults. DO NOT EDIT.
// versions:
// 		protoc-gen-godefaults 0.0.2
// source: xgo/tests/godefaultstest/godefaults_test_external3.proto

package godefaultstest

import (
	_ "github.com/yu31/protoc-plugin/xgo/pb/pbdefaults"
	godefaultsexternal "github.com/yu31/protoc-plugin/xgo/tests/godefaultsexternal"
)

// Set default value for message godefaultstest.MessageExternal3
func (this *MessageExternal3) SetDefaults() {
	if this == nil {
		return
	}
	if this.Status == nil {
		this.Status = []godefaultsexternal.ExternalMessage1_EmbedEnum1{1, 2, 0}
	}
	return
}
