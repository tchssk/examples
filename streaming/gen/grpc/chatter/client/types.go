// Code generated by goa v3.4.2, DO NOT EDIT.
//
// chatter gRPC client types
//
// Command:
// $ goa gen goa.design/examples/streaming/design -o
// $(GOPATH)/src/goa.design/examples/streaming

package client

import (
	chatter "goa.design/examples/streaming/gen/chatter"
	chatterviews "goa.design/examples/streaming/gen/chatter/views"
	chatterpb "goa.design/examples/streaming/gen/grpc/chatter/pb"
	goa "goa.design/goa/v3/pkg"
)

// NewLoginRequest builds the gRPC request type from the payload of the "login"
// endpoint of the "chatter" service.
func NewLoginRequest() *chatterpb.LoginRequest {
	message := &chatterpb.LoginRequest{}
	return message
}

// NewLoginResult builds the result type of the "login" endpoint of the
// "chatter" service from the gRPC response type.
func NewLoginResult(message *chatterpb.LoginResponse) string {
	result := message.Field
	return result
}

func NewEchoerResponse(v *chatterpb.EchoerResponse) string {
	result := v.Field
	return result
}

func NewEchoerStreamingRequest(spayload string) *chatterpb.EchoerStreamingRequest {
	v := &chatterpb.EchoerStreamingRequest{}
	v.Field = spayload
	return v
}

func NewListenerStreamingRequest(spayload string) *chatterpb.ListenerStreamingRequest {
	v := &chatterpb.ListenerStreamingRequest{}
	v.Field = spayload
	return v
}

func NewChatSummaryCollection(v *chatterpb.ChatSummaryCollection) chatterviews.ChatSummaryCollectionView {
	vresult := make([]*chatterviews.ChatSummaryView, len(v.Field))
	for i, val := range v.Field {
		vresult[i] = &chatterviews.ChatSummaryView{
			Message: &val.Message_,
			SentAt:  &val.SentAt,
		}
		if val.Length != 0 {
			lengthptr := int(val.Length)
			vresult[i].Length = &lengthptr
		}
	}
	return vresult
}

func NewSummaryStreamingRequest(spayload string) *chatterpb.SummaryStreamingRequest {
	v := &chatterpb.SummaryStreamingRequest{}
	v.Field = spayload
	return v
}

// NewSubscribeRequest builds the gRPC request type from the payload of the
// "subscribe" endpoint of the "chatter" service.
func NewSubscribeRequest() *chatterpb.SubscribeRequest {
	message := &chatterpb.SubscribeRequest{}
	return message
}

func NewEvent(v *chatterpb.SubscribeResponse) *chatter.Event {
	result := &chatter.Event{
		Message: v.Message_,
		Action:  v.Action,
		AddedAt: v.AddedAt,
	}
	return result
}

// NewHistoryRequest builds the gRPC request type from the payload of the
// "history" endpoint of the "chatter" service.
func NewHistoryRequest() *chatterpb.HistoryRequest {
	message := &chatterpb.HistoryRequest{}
	return message
}

func NewChatSummaryView(v *chatterpb.HistoryResponse) *chatterviews.ChatSummaryView {
	vresult := &chatterviews.ChatSummaryView{
		Message: &v.Message_,
		SentAt:  &v.SentAt,
	}
	if v.Length != 0 {
		lengthptr := int(v.Length)
		vresult.Length = &lengthptr
	}
	return vresult
}

// ValidateChatSummaryCollection runs the validations defined on
// ChatSummaryCollection.
func ValidateChatSummaryCollection(message *chatterpb.ChatSummaryCollection) (err error) {
	for _, e := range message.Field {
		if e != nil {
			if err2 := ValidateChatSummary(e); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// ValidateChatSummary runs the validations defined on ChatSummary.
func ValidateChatSummary(message *chatterpb.ChatSummary) (err error) {
	err = goa.MergeErrors(err, goa.ValidateFormat("message.sent_at", message.SentAt, goa.FormatDateTime))

	return
}

// ValidateSubscribeResponse runs the validations defined on SubscribeResponse.
func ValidateSubscribeResponse(message *chatterpb.SubscribeResponse) (err error) {
	if !(message.Action == "added") {
		err = goa.MergeErrors(err, goa.InvalidEnumValueError("message.action", message.Action, []interface{}{"added"}))
	}
	err = goa.MergeErrors(err, goa.ValidateFormat("message.added_at", message.AddedAt, goa.FormatDateTime))

	return
}

// ValidateHistoryResponse runs the validations defined on HistoryResponse.
func ValidateHistoryResponse(message *chatterpb.HistoryResponse) (err error) {
	err = goa.MergeErrors(err, goa.ValidateFormat("message.sent_at", message.SentAt, goa.FormatDateTime))

	return
}
