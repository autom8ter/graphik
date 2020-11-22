// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: api/graphik.proto

package apipb

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/golang/protobuf/ptypes/any"
	_ "github.com/golang/protobuf/ptypes/empty"
	_ "github.com/golang/protobuf/ptypes/struct"
	_ "github.com/golang/protobuf/ptypes/timestamp"
	_ "github.com/mwitkow/go-proto-validators"
	github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
	math "math"
	regexp "regexp"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

var _regex_Path_Gtype = regexp.MustCompile(`^.{1,225}$`)

func (this *Path) Validate() error {
	if !_regex_Path_Gtype.MatchString(this.Gtype) {
		return github_com_mwitkow_go_proto_validators.FieldError("Gtype", fmt.Errorf(`value '%v' must be a string conforming to regex "^.{1,225}$"`, this.Gtype))
	}
	return nil
}
func (this *Metadata) Validate() error {
	if this.UpdatedBy != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.UpdatedBy); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("UpdatedBy", err)
		}
	}
	return nil
}
func (this *Paths) Validate() error {
	for _, item := range this.Paths {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Paths", err)
			}
		}
	}
	return nil
}
func (this *Node) Validate() error {
	if this.Path != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Path); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Path", err)
		}
	}
	if this.Attributes != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Attributes); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Attributes", err)
		}
	}
	if this.Metadata != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Metadata); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Metadata", err)
		}
	}
	return nil
}
func (this *NodeConstructor) Validate() error {
	if nil == this.Path {
		return github_com_mwitkow_go_proto_validators.FieldError("Path", fmt.Errorf("message must exist"))
	}
	if this.Path != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Path); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Path", err)
		}
	}
	if this.Attributes != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Attributes); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Attributes", err)
		}
	}
	return nil
}
func (this *NodeConstructors) Validate() error {
	for _, item := range this.Nodes {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Nodes", err)
			}
		}
	}
	return nil
}
func (this *Nodes) Validate() error {
	for _, item := range this.Nodes {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Nodes", err)
			}
		}
	}
	return nil
}
func (this *NodeDetail) Validate() error {
	if this.Path != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Path); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Path", err)
		}
	}
	if this.Attributes != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Attributes); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Attributes", err)
		}
	}
	if this.EdgesFrom != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.EdgesFrom); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("EdgesFrom", err)
		}
	}
	if this.EdgesTo != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.EdgesTo); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("EdgesTo", err)
		}
	}
	if this.Metadata != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Metadata); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Metadata", err)
		}
	}
	return nil
}
func (this *NodeDetails) Validate() error {
	for _, item := range this.NodeDetails {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("NodeDetails", err)
			}
		}
	}
	return nil
}
func (this *NodeDetailFilter) Validate() error {
	if this.Path != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Path); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Path", err)
		}
	}
	if this.EdgesFrom != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.EdgesFrom); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("EdgesFrom", err)
		}
	}
	if this.EdgesTo != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.EdgesTo); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("EdgesTo", err)
		}
	}
	return nil
}
func (this *Edge) Validate() error {
	if this.Path != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Path); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Path", err)
		}
	}
	if this.Attributes != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Attributes); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Attributes", err)
		}
	}
	if this.From != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.From); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("From", err)
		}
	}
	if this.To != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.To); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("To", err)
		}
	}
	if this.Metadata != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Metadata); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Metadata", err)
		}
	}
	return nil
}
func (this *EdgeConstructor) Validate() error {
	if nil == this.Path {
		return github_com_mwitkow_go_proto_validators.FieldError("Path", fmt.Errorf("message must exist"))
	}
	if this.Path != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Path); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Path", err)
		}
	}
	if this.Attributes != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Attributes); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Attributes", err)
		}
	}
	if nil == this.From {
		return github_com_mwitkow_go_proto_validators.FieldError("From", fmt.Errorf("message must exist"))
	}
	if this.From != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.From); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("From", err)
		}
	}
	if nil == this.To {
		return github_com_mwitkow_go_proto_validators.FieldError("To", fmt.Errorf("message must exist"))
	}
	if this.To != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.To); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("To", err)
		}
	}
	return nil
}
func (this *EdgeConstructors) Validate() error {
	for _, item := range this.Edges {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Edges", err)
			}
		}
	}
	return nil
}
func (this *Edges) Validate() error {
	for _, item := range this.Edges {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Edges", err)
			}
		}
	}
	return nil
}
func (this *EdgeDetail) Validate() error {
	if this.Path != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Path); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Path", err)
		}
	}
	if this.Attributes != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Attributes); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Attributes", err)
		}
	}
	if this.From != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.From); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("From", err)
		}
	}
	if this.To != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.To); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("To", err)
		}
	}
	if this.Metadata != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Metadata); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Metadata", err)
		}
	}
	return nil
}
func (this *EdgeDetails) Validate() error {
	for _, item := range this.Edges {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Edges", err)
			}
		}
	}
	return nil
}

var _regex_EdgeFilter_Gtype = regexp.MustCompile(`^.{1,225}$`)

func (this *EdgeFilter) Validate() error {
	if nil == this.NodePath {
		return github_com_mwitkow_go_proto_validators.FieldError("NodePath", fmt.Errorf("message must exist"))
	}
	if this.NodePath != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.NodePath); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("NodePath", err)
		}
	}
	if !_regex_EdgeFilter_Gtype.MatchString(this.Gtype) {
		return github_com_mwitkow_go_proto_validators.FieldError("Gtype", fmt.Errorf(`value '%v' must be a string conforming to regex "^.{1,225}$"`, this.Gtype))
	}
	if !(this.Limit > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("Limit", fmt.Errorf(`value '%v' must be greater than '0'`, this.Limit))
	}
	return nil
}

var _regex_Filter_Gtype = regexp.MustCompile(`^.{1,225}$`)

func (this *Filter) Validate() error {
	if !_regex_Filter_Gtype.MatchString(this.Gtype) {
		return github_com_mwitkow_go_proto_validators.FieldError("Gtype", fmt.Errorf(`value '%v' must be a string conforming to regex "^.{1,225}$"`, this.Gtype))
	}
	if !(this.Limit > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("Limit", fmt.Errorf(`value '%v' must be greater than '0'`, this.Limit))
	}
	return nil
}
func (this *MeFilter) Validate() error {
	if this.EdgesFrom != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.EdgesFrom); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("EdgesFrom", err)
		}
	}
	if this.EdgesTo != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.EdgesTo); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("EdgesTo", err)
		}
	}
	return nil
}

var _regex_ChannelFilter_Channel = regexp.MustCompile(`^.{1,225}$`)

func (this *ChannelFilter) Validate() error {
	if !_regex_ChannelFilter_Channel.MatchString(this.Channel) {
		return github_com_mwitkow_go_proto_validators.FieldError("Channel", fmt.Errorf(`value '%v' must be a string conforming to regex "^.{1,225}$"`, this.Channel))
	}
	return nil
}
func (this *SubGraphFilter) Validate() error {
	if nil == this.Nodes {
		return github_com_mwitkow_go_proto_validators.FieldError("Nodes", fmt.Errorf("message must exist"))
	}
	if this.Nodes != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Nodes); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Nodes", err)
		}
	}
	if nil == this.Edges {
		return github_com_mwitkow_go_proto_validators.FieldError("Edges", fmt.Errorf("message must exist"))
	}
	if this.Edges != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Edges); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Edges", err)
		}
	}
	return nil
}
func (this *Graph) Validate() error {
	if this.Nodes != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Nodes); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Nodes", err)
		}
	}
	if this.Edges != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Edges); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Edges", err)
		}
	}
	return nil
}
func (this *Patch) Validate() error {
	if nil == this.Path {
		return github_com_mwitkow_go_proto_validators.FieldError("Path", fmt.Errorf("message must exist"))
	}
	if this.Path != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Path); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Path", err)
		}
	}
	if this.Attributes != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Attributes); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Attributes", err)
		}
	}
	return nil
}
func (this *Patches) Validate() error {
	for _, item := range this.Patches {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Patches", err)
			}
		}
	}
	return nil
}
func (this *Pong) Validate() error {
	return nil
}

var _regex_OutboundMessage_Channel = regexp.MustCompile(`^.{1,225}$`)

func (this *OutboundMessage) Validate() error {
	if !_regex_OutboundMessage_Channel.MatchString(this.Channel) {
		return github_com_mwitkow_go_proto_validators.FieldError("Channel", fmt.Errorf(`value '%v' must be a string conforming to regex "^.{1,225}$"`, this.Channel))
	}
	if nil == this.Data {
		return github_com_mwitkow_go_proto_validators.FieldError("Data", fmt.Errorf("message must exist"))
	}
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}

var _regex_Message_Channel = regexp.MustCompile(`^.{1,225}$`)

func (this *Message) Validate() error {
	if !_regex_Message_Channel.MatchString(this.Channel) {
		return github_com_mwitkow_go_proto_validators.FieldError("Channel", fmt.Errorf(`value '%v' must be a string conforming to regex "^.{1,225}$"`, this.Channel))
	}
	if nil == this.Data {
		return github_com_mwitkow_go_proto_validators.FieldError("Data", fmt.Errorf("message must exist"))
	}
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	if this.Sender != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Sender); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Sender", err)
		}
	}
	return nil
}
func (this *Schema) Validate() error {
	return nil
}

var _regex_Interception_Method = regexp.MustCompile(`^.{1,225}$`)

func (this *Interception) Validate() error {
	if !_regex_Interception_Method.MatchString(this.Method) {
		return github_com_mwitkow_go_proto_validators.FieldError("Method", fmt.Errorf(`value '%v' must be a string conforming to regex "^.{1,225}$"`, this.Method))
	}
	if this.Identity != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Identity); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Identity", err)
		}
	}
	if this.Request != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Request); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Request", err)
		}
	}
	return nil
}
func (this *TriggerFilter) Validate() error {
	return nil
}
