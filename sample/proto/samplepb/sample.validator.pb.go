// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: sample.proto

package sample

import (
	fmt "fmt"
	math "math"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/mwitkow/go-proto-validators"
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"
	_ "github.com/golang/protobuf/ptypes/any"
	_ "github.com/golang/protobuf/ptypes/struct"
	_ "github.com/golang/protobuf/ptypes/wrappers"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *Request) Validate() error {
	if !(this.Age > 20) {
		return github_com_mwitkow_go_proto_validators.FieldError("Age", fmt.Errorf(`value '%v' must be greater than '20'`, this.Age))
	}
	if !(this.Age < 27) {
		return github_com_mwitkow_go_proto_validators.FieldError("Age", fmt.Errorf(`value '%v' must be less than '27'`, this.Age))
	}
	// Validation of proto3 map<> fields is unsupported.
	return nil
}
func (this *Reply) Validate() error {
	return nil
}
func (this *PingRequest) Validate() error {
	if this.StMetaData != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.StMetaData); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("StMetaData", err)
		}
	}
	if this.StCustomData != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.StCustomData); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("StCustomData", err)
		}
	}
	if this.DoubleValue != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.DoubleValue); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("DoubleValue", err)
		}
	}
	if this.StringValue != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.StringValue); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("StringValue", err)
		}
	}
	if this.BoolValue != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.BoolValue); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("BoolValue", err)
		}
	}
	if this.Int64Value != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Int64Value); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Int64Value", err)
		}
	}
	return nil
}
func (this *PongReply) Validate() error {
	if this.StMetaData != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.StMetaData); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("StMetaData", err)
		}
	}
	if this.StCustomData != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.StCustomData); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("StCustomData", err)
		}
	}
	if this.AnyData != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.AnyData); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("AnyData", err)
		}
	}
	return nil
}
func (this *UploadReq) Validate() error {
	return nil
}
func (this *UploadResp) Validate() error {
	return nil
}
func (this *GetMsgReq) Validate() error {
	return nil
}
func (this *GetMsgResp) Validate() error {
	return nil
}
