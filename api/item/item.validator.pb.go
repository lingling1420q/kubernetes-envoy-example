// Code generated by protoc-gen-gogo.
// source: item/item.proto
// DO NOT EDIT!

/*
Package item is a generated protocol buffer package.

It is generated from these files:
	item/item.proto

It has these top-level messages:
	Item
	GetItemRequest
	ListItemsRequest
	ListItemsResponse
*/
package item

import github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *Item) Validate() error {
	return nil
}
func (this *GetItemRequest) Validate() error {
	return nil
}
func (this *ListItemsRequest) Validate() error {
	return nil
}
func (this *ListItemsResponse) Validate() error {
	for _, item := range this.Items {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Items", err)
			}
		}
	}
	return nil
}
