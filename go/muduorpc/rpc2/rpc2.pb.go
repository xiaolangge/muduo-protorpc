// Code generated by protoc-gen-go.
// source: muduo/protorpc2/rpc2.proto
// DO NOT EDIT!

/*
Package rpc2 is a generated protocol buffer package.

It is generated from these files:
	muduo/protorpc2/rpc2.proto

It has these top-level messages:
	Empty
*/
package rpc2

import proto "code.google.com/p/goprotobuf/proto"
import json "encoding/json"
import math "math"
import google_protobuf "code.google.com/p/goprotobuf/protoc-gen-go/descriptor"

// Reference proto, json, and math imports to suppress error if they are not otherwise used.
var _ = proto.Marshal
var _ = &json.SyntaxError{}
var _ = math.Inf

type Empty struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}

var E_Idempotent = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf.MethodOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         1111,
	Name:          "rpc2.idempotent",
	Tag:           "varint,1111,opt,name=idempotent",
}

var E_Noreturn = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf.MethodOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         1112,
	Name:          "rpc2.noreturn",
	Tag:           "varint,1112,opt,name=noreturn",
}

func init() {
	proto.RegisterExtension(E_Idempotent)
	proto.RegisterExtension(E_Noreturn)
}
