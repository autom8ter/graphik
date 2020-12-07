// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: graphik.proto

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

var _regex_Ref_Gtype = regexp.MustCompile(`^.{1,225}$`)
var _regex_Ref_Gid = regexp.MustCompile(`^.{1,225}$`)

func (this *Ref) Validate() error {
	if !_regex_Ref_Gtype.MatchString(this.Gtype) {
		return github_com_mwitkow_go_proto_validators.FieldError("Gtype", fmt.Errorf(`value '%v' must be a string conforming to regex "^.{1,225}$"`, this.Gtype))
	}
	if !_regex_Ref_Gid.MatchString(this.Gid) {
		return github_com_mwitkow_go_proto_validators.FieldError("Gid", fmt.Errorf(`value '%v' must be a string conforming to regex "^.{1,225}$"`, this.Gid))
	}
	return nil
}

var _regex_RefConstructor_Gtype = regexp.MustCompile(`^.{1,225}$`)

func (this *RefConstructor) Validate() error {
	if !_regex_RefConstructor_Gtype.MatchString(this.Gtype) {
		return github_com_mwitkow_go_proto_validators.FieldError("Gtype", fmt.Errorf(`value '%v' must be a string conforming to regex "^.{1,225}$"`, this.Gtype))
	}
	return nil
}
func (this *Refs) Validate() error {
	for _, item := range this.Refs {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Refs", err)
			}
		}
	}
	return nil
}
func (this *Doc) Validate() error {
	if nil == this.Ref {
		return github_com_mwitkow_go_proto_validators.FieldError("Ref", fmt.Errorf("message must exist"))
	}
	if this.Ref != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Ref); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Ref", err)
		}
	}
	if this.Attributes != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Attributes); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Attributes", err)
		}
	}
	return nil
}
func (this *DocConstructor) Validate() error {
	if nil == this.Ref {
		return github_com_mwitkow_go_proto_validators.FieldError("Ref", fmt.Errorf("message must exist"))
	}
	if this.Ref != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Ref); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Ref", err)
		}
	}
	if this.Attributes != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Attributes); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Attributes", err)
		}
	}
	return nil
}
func (this *DocConstructors) Validate() error {
	for _, item := range this.Docs {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Docs", err)
			}
		}
	}
	return nil
}
func (this *Docs) Validate() error {
	for _, item := range this.Docs {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Docs", err)
			}
		}
	}
	return nil
}
func (this *Connection) Validate() error {
	if nil == this.Ref {
		return github_com_mwitkow_go_proto_validators.FieldError("Ref", fmt.Errorf("message must exist"))
	}
	if this.Ref != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Ref); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Ref", err)
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
func (this *ConnectionConstructor) Validate() error {
	if nil == this.Ref {
		return github_com_mwitkow_go_proto_validators.FieldError("Ref", fmt.Errorf("message must exist"))
	}
	if this.Ref != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Ref); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Ref", err)
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
func (this *SConnectFilter) Validate() error {
	if this.Filter != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Filter); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Filter", err)
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
	return nil
}
func (this *ConnectionConstructors) Validate() error {
	for _, item := range this.Connections {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Connections", err)
			}
		}
	}
	return nil
}
func (this *Connections) Validate() error {
	for _, item := range this.Connections {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Connections", err)
			}
		}
	}
	return nil
}

var _regex_CFilter_Gtype = regexp.MustCompile(`^.{1,225}$`)
var _regex_CFilter_Sort = regexp.MustCompile(`((^|, )(|ref.gid|ref.gtype|^attributes.(.*)))+$`)

func (this *CFilter) Validate() error {
	if nil == this.DocRef {
		return github_com_mwitkow_go_proto_validators.FieldError("DocRef", fmt.Errorf("message must exist"))
	}
	if this.DocRef != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.DocRef); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("DocRef", err)
		}
	}
	if !_regex_CFilter_Gtype.MatchString(this.Gtype) {
		return github_com_mwitkow_go_proto_validators.FieldError("Gtype", fmt.Errorf(`value '%v' must be a string conforming to regex "^.{1,225}$"`, this.Gtype))
	}
	if !(this.Limit > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("Limit", fmt.Errorf(`value '%v' must be greater than '0'`, this.Limit))
	}
	if !_regex_CFilter_Sort.MatchString(this.Sort) {
		return github_com_mwitkow_go_proto_validators.FieldError("Sort", fmt.Errorf(`value '%v' must be a string conforming to regex "((^|, )(|ref.gid|ref.gtype|^attributes.(.*)))+$"`, this.Sort))
	}
	return nil
}

var _regex_Filter_Gtype = regexp.MustCompile(`^.{1,225}$`)
var _regex_Filter_Sort = regexp.MustCompile(`((^|, )(|ref.gid|ref.gtype|^attributes.(.*)))+$`)

func (this *Filter) Validate() error {
	if !_regex_Filter_Gtype.MatchString(this.Gtype) {
		return github_com_mwitkow_go_proto_validators.FieldError("Gtype", fmt.Errorf(`value '%v' must be a string conforming to regex "^.{1,225}$"`, this.Gtype))
	}
	if !(this.Limit > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("Limit", fmt.Errorf(`value '%v' must be greater than '0'`, this.Limit))
	}
	if !_regex_Filter_Sort.MatchString(this.Sort) {
		return github_com_mwitkow_go_proto_validators.FieldError("Sort", fmt.Errorf(`value '%v' must be a string conforming to regex "((^|, )(|ref.gid|ref.gtype|^attributes.(.*)))+$"`, this.Sort))
	}
	return nil
}

var _regex_AggFilter_Aggregate = regexp.MustCompile(`((^|, )(sum|count|max|min|avg|prod))+$`)
var _regex_AggFilter_Field = regexp.MustCompile(`((^|, )(|^attributes.(.*)))+$`)

func (this *AggFilter) Validate() error {
	if nil == this.Filter {
		return github_com_mwitkow_go_proto_validators.FieldError("Filter", fmt.Errorf("message must exist"))
	}
	if this.Filter != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Filter); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Filter", err)
		}
	}
	if !_regex_AggFilter_Aggregate.MatchString(this.Aggregate) {
		return github_com_mwitkow_go_proto_validators.FieldError("Aggregate", fmt.Errorf(`value '%v' must be a string conforming to regex "((^|, )(sum|count|max|min|avg|prod))+$"`, this.Aggregate))
	}
	if !_regex_AggFilter_Field.MatchString(this.Field) {
		return github_com_mwitkow_go_proto_validators.FieldError("Field", fmt.Errorf(`value '%v' must be a string conforming to regex "((^|, )(|^attributes.(.*)))+$"`, this.Field))
	}
	return nil
}

var _regex_TFilter_Sort = regexp.MustCompile(`((^|, )(|ref.gid|ref.gtype|^attributes.(.*)))+$`)

func (this *TFilter) Validate() error {
	if nil == this.Root {
		return github_com_mwitkow_go_proto_validators.FieldError("Root", fmt.Errorf("message must exist"))
	}
	if this.Root != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Root); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Root", err)
		}
	}
	if !(this.Limit > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("Limit", fmt.Errorf(`value '%v' must be greater than '0'`, this.Limit))
	}
	if !_regex_TFilter_Sort.MatchString(this.Sort) {
		return github_com_mwitkow_go_proto_validators.FieldError("Sort", fmt.Errorf(`value '%v' must be a string conforming to regex "((^|, )(|ref.gid|ref.gtype|^attributes.(.*)))+$"`, this.Sort))
	}
	return nil
}

var _regex_IndexConstructor_Name = regexp.MustCompile(`^.{1,225}$`)
var _regex_IndexConstructor_Gtype = regexp.MustCompile(`^.{1,225}$`)
var _regex_IndexConstructor_Expression = regexp.MustCompile(`^.{1,225}$`)

func (this *IndexConstructor) Validate() error {
	if !_regex_IndexConstructor_Name.MatchString(this.Name) {
		return github_com_mwitkow_go_proto_validators.FieldError("Name", fmt.Errorf(`value '%v' must be a string conforming to regex "^.{1,225}$"`, this.Name))
	}
	if !_regex_IndexConstructor_Gtype.MatchString(this.Gtype) {
		return github_com_mwitkow_go_proto_validators.FieldError("Gtype", fmt.Errorf(`value '%v' must be a string conforming to regex "^.{1,225}$"`, this.Gtype))
	}
	if !_regex_IndexConstructor_Expression.MatchString(this.Expression) {
		return github_com_mwitkow_go_proto_validators.FieldError("Expression", fmt.Errorf(`value '%v' must be a string conforming to regex "^.{1,225}$"`, this.Expression))
	}
	return nil
}

var _regex_Authorizer_Name = regexp.MustCompile(`^.{1,225}$`)
var _regex_Authorizer_Expression = regexp.MustCompile(`^.{1,225}$`)

func (this *Authorizer) Validate() error {
	if !_regex_Authorizer_Name.MatchString(this.Name) {
		return github_com_mwitkow_go_proto_validators.FieldError("Name", fmt.Errorf(`value '%v' must be a string conforming to regex "^.{1,225}$"`, this.Name))
	}
	if !_regex_Authorizer_Expression.MatchString(this.Expression) {
		return github_com_mwitkow_go_proto_validators.FieldError("Expression", fmt.Errorf(`value '%v' must be a string conforming to regex "^.{1,225}$"`, this.Expression))
	}
	return nil
}
func (this *Authorizers) Validate() error {
	for _, item := range this.Authorizers {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Authorizers", err)
			}
		}
	}
	return nil
}

var _regex_TypeValidator_Name = regexp.MustCompile(`^.{1,225}$`)
var _regex_TypeValidator_Gtype = regexp.MustCompile(`^.{1,225}$`)
var _regex_TypeValidator_Expression = regexp.MustCompile(`^.{1,225}$`)

func (this *TypeValidator) Validate() error {
	if !_regex_TypeValidator_Name.MatchString(this.Name) {
		return github_com_mwitkow_go_proto_validators.FieldError("Name", fmt.Errorf(`value '%v' must be a string conforming to regex "^.{1,225}$"`, this.Name))
	}
	if !_regex_TypeValidator_Gtype.MatchString(this.Gtype) {
		return github_com_mwitkow_go_proto_validators.FieldError("Gtype", fmt.Errorf(`value '%v' must be a string conforming to regex "^.{1,225}$"`, this.Gtype))
	}
	if !_regex_TypeValidator_Expression.MatchString(this.Expression) {
		return github_com_mwitkow_go_proto_validators.FieldError("Expression", fmt.Errorf(`value '%v' must be a string conforming to regex "^.{1,225}$"`, this.Expression))
	}
	return nil
}
func (this *TypeValidators) Validate() error {
	for _, item := range this.Validators {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Validators", err)
			}
		}
	}
	return nil
}

var _regex_Index_Name = regexp.MustCompile(`^.{1,225}$`)
var _regex_Index_Gtype = regexp.MustCompile(`^.{1,225}$`)
var _regex_Index_Expression = regexp.MustCompile(`^.{1,225}$`)

func (this *Index) Validate() error {
	if !_regex_Index_Name.MatchString(this.Name) {
		return github_com_mwitkow_go_proto_validators.FieldError("Name", fmt.Errorf(`value '%v' must be a string conforming to regex "^.{1,225}$"`, this.Name))
	}
	if !_regex_Index_Gtype.MatchString(this.Gtype) {
		return github_com_mwitkow_go_proto_validators.FieldError("Gtype", fmt.Errorf(`value '%v' must be a string conforming to regex "^.{1,225}$"`, this.Gtype))
	}
	if !_regex_Index_Expression.MatchString(this.Expression) {
		return github_com_mwitkow_go_proto_validators.FieldError("Expression", fmt.Errorf(`value '%v' must be a string conforming to regex "^.{1,225}$"`, this.Expression))
	}
	return nil
}
func (this *Indexes) Validate() error {
	for _, item := range this.Indexes {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Indexes", err)
			}
		}
	}
	return nil
}

var _regex_ChanFilter_Channel = regexp.MustCompile(`^.{1,225}$`)

func (this *ChanFilter) Validate() error {
	if !_regex_ChanFilter_Channel.MatchString(this.Channel) {
		return github_com_mwitkow_go_proto_validators.FieldError("Channel", fmt.Errorf(`value '%v' must be a string conforming to regex "^.{1,225}$"`, this.Channel))
	}
	return nil
}
func (this *Graph) Validate() error {
	if this.Docs != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Docs); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Docs", err)
		}
	}
	if this.Connections != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Connections); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Connections", err)
		}
	}
	return nil
}
func (this *Flags) Validate() error {
	return nil
}
func (this *Edit) Validate() error {
	if nil == this.Ref {
		return github_com_mwitkow_go_proto_validators.FieldError("Ref", fmt.Errorf("message must exist"))
	}
	if this.Ref != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Ref); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Ref", err)
		}
	}
	if this.Attributes != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Attributes); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Attributes", err)
		}
	}
	return nil
}
func (this *EFilter) Validate() error {
	if this.Filter != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Filter); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Filter", err)
		}
	}
	if this.Attributes != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Attributes); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Attributes", err)
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
	if nil == this.Sender {
		return github_com_mwitkow_go_proto_validators.FieldError("Sender", fmt.Errorf("message must exist"))
	}
	if this.Sender != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Sender); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Sender", err)
		}
	}
	if nil == this.Timestamp {
		return github_com_mwitkow_go_proto_validators.FieldError("Timestamp", fmt.Errorf("message must exist"))
	}
	if this.Timestamp != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Timestamp); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Timestamp", err)
		}
	}
	return nil
}
func (this *Schema) Validate() error {
	if this.Authorizers != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Authorizers); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Authorizers", err)
		}
	}
	if this.Validators != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Validators); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Validators", err)
		}
	}
	if this.Indexes != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Indexes); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Indexes", err)
		}
	}
	return nil
}
func (this *ExprFilter) Validate() error {
	return nil
}

var _regex_Request_Method = regexp.MustCompile(`^.{1,225}$`)

func (this *Request) Validate() error {
	if !_regex_Request_Method.MatchString(this.Method) {
		return github_com_mwitkow_go_proto_validators.FieldError("Method", fmt.Errorf(`value '%v' must be a string conforming to regex "^.{1,225}$"`, this.Method))
	}
	if nil == this.Identity {
		return github_com_mwitkow_go_proto_validators.FieldError("Identity", fmt.Errorf("message must exist"))
	}
	if this.Identity != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Identity); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Identity", err)
		}
	}
	if nil == this.Timestamp {
		return github_com_mwitkow_go_proto_validators.FieldError("Timestamp", fmt.Errorf("message must exist"))
	}
	if this.Timestamp != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Timestamp); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Timestamp", err)
		}
	}
	if this.Request != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Request); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Request", err)
		}
	}
	return nil
}
