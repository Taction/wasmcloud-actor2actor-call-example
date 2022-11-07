package uppercase

import (
	"github.com/wasmcloud/actor-tinygo"           //nolint
	cbor "github.com/wasmcloud/tinygo-cbor"       //nolint
	msgpack "github.com/wasmcloud/tinygo-msgpack" //nolint
)

// Parameters sent for Uppercase
type UppercaseRequest struct {
	// method to be Uppercased
	Data string
}

// MEncode serializes a UppercaseRequest using msgpack
func (o *UppercaseRequest) MEncode(encoder msgpack.Writer) error {
	encoder.WriteMapSize(1)
	encoder.WriteString("data")
	encoder.WriteString(o.Data)

	return encoder.CheckError()
}

// MDecodeUppercaseRequest deserializes a UppercaseRequest using msgpack
func MDecodeUppercaseRequest(d *msgpack.Decoder) (UppercaseRequest, error) {
	var val UppercaseRequest
	isNil, err := d.IsNextNil()
	if err != nil || isNil {
		return val, err
	}
	size, err := d.ReadMapSize()
	if err != nil {
		return val, err
	}
	for i := uint32(0); i < size; i++ {
		field, err := d.ReadString()
		if err != nil {
			return val, err
		}
		switch field {
		case "data":
			val.Data, err = d.ReadString()
		default:
			err = d.Skip()
		}
		if err != nil {
			return val, err
		}
	}
	return val, nil
}

// CEncode serializes a UppercaseRequest using cbor
func (o *UppercaseRequest) CEncode(encoder cbor.Writer) error {
	encoder.WriteMapSize(1)
	encoder.WriteString("data")
	encoder.WriteString(o.Data)

	return encoder.CheckError()
}

// CDecodeUppercaseRequest deserializes a UppercaseRequest using cbor
func CDecodeUppercaseRequest(d *cbor.Decoder) (UppercaseRequest, error) {
	var val UppercaseRequest
	isNil, err := d.IsNextNil()
	if err != nil || isNil {
		return val, err
	}
	size, indef, err := d.ReadMapSize()
	if err != nil && indef {
		err = cbor.NewReadError("indefinite maps not supported")
	}
	if err != nil {
		return val, err
	}
	for i := uint32(0); i < size; i++ {
		field, err := d.ReadString()
		if err != nil {
			return val, err
		}
		switch field {
		case "data":
			val.Data, err = d.ReadString()
		default:
			err = d.Skip()
		}
		if err != nil {
			return val, err
		}
	}
	return val, nil
}

// Response to Uppercase
type UppercaseResponse struct {
	// response data
	Data string
	// Indicates a successful authorization
	Success bool
}

// MEncode serializes a UppercaseResponse using msgpack
func (o *UppercaseResponse) MEncode(encoder msgpack.Writer) error {
	encoder.WriteMapSize(2)
	encoder.WriteString("data")
	encoder.WriteString(o.Data)
	encoder.WriteString("success")
	encoder.WriteBool(o.Success)

	return encoder.CheckError()
}

// MDecodeUppercaseResponse deserializes a UppercaseResponse using msgpack
func MDecodeUppercaseResponse(d *msgpack.Decoder) (UppercaseResponse, error) {
	var val UppercaseResponse
	isNil, err := d.IsNextNil()
	if err != nil || isNil {
		return val, err
	}
	size, err := d.ReadMapSize()
	if err != nil {
		return val, err
	}
	for i := uint32(0); i < size; i++ {
		field, err := d.ReadString()
		if err != nil {
			return val, err
		}
		switch field {
		case "data":
			val.Data, err = d.ReadString()
		case "success":
			val.Success, err = d.ReadBool()
		default:
			err = d.Skip()
		}
		if err != nil {
			return val, err
		}
	}
	return val, nil
}

// CEncode serializes a UppercaseResponse using cbor
func (o *UppercaseResponse) CEncode(encoder cbor.Writer) error {
	encoder.WriteMapSize(2)
	encoder.WriteString("data")
	encoder.WriteString(o.Data)
	encoder.WriteString("success")
	encoder.WriteBool(o.Success)

	return encoder.CheckError()
}

// CDecodeUppercaseResponse deserializes a UppercaseResponse using cbor
func CDecodeUppercaseResponse(d *cbor.Decoder) (UppercaseResponse, error) {
	var val UppercaseResponse
	isNil, err := d.IsNextNil()
	if err != nil || isNil {
		return val, err
	}
	size, indef, err := d.ReadMapSize()
	if err != nil && indef {
		err = cbor.NewReadError("indefinite maps not supported")
	}
	if err != nil {
		return val, err
	}
	for i := uint32(0); i < size; i++ {
		field, err := d.ReadString()
		if err != nil {
			return val, err
		}
		switch field {
		case "data":
			val.Data, err = d.ReadString()
		case "success":
			val.Success, err = d.ReadBool()
		default:
			err = d.Skip()
		}
		if err != nil {
			return val, err
		}
	}
	return val, nil
}

// Description of Uppercase service
type Uppercase interface {
	// Uppercase - Execute transaction
	Upper(ctx *actor.Context, arg UppercaseRequest) (*UppercaseResponse, error)
}

// UppercaseHandler is called by an actor during `main` to generate a dispatch handler
// The output of this call should be passed into `actor.RegisterHandlers`
func UppercaseHandler(actor_ Uppercase) actor.Handler {
	return actor.NewHandler("Uppercase", &UppercaseReceiver{}, actor_)
}

// UppercaseReceiver receives messages defined in the Uppercase service interface
// Description of Uppercase service
type UppercaseReceiver struct{}

func (r *UppercaseReceiver) Dispatch(ctx *actor.Context, svc interface{}, message *actor.Message) (*actor.Message, error) {
	svc_, _ := svc.(Uppercase)
	switch message.Method {

	case "Upper":
		{

			d := msgpack.NewDecoder(message.Arg)
			value, err_ := MDecodeUppercaseRequest(&d)
			if err_ != nil {
				return nil, err_
			}

			resp, err := svc_.Upper(ctx, value)
			if err != nil {
				return nil, err
			}

			var sizer msgpack.Sizer
			size_enc := &sizer
			resp.MEncode(size_enc)
			buf := make([]byte, sizer.Len())
			encoder := msgpack.NewEncoder(buf)
			enc := &encoder
			resp.MEncode(enc)
			return &actor.Message{Method: "Uppercase.Upper", Arg: buf}, nil
		}
	default:
		return nil, actor.NewRpcError("MethodNotHandled", "Uppercase."+message.Method)
	}
}

// UppercaseSender sends messages to a Uppercase service
// Description of Uppercase service
type UppercaseSender struct{ transport actor.Transport }

// NewActorSender constructs a client for actor-to-actor messaging
// using the recipient actor's public key
func NewActorUppercaseSender(actor_id string) *UppercaseSender {
	transport := actor.ToActor(actor_id)
	return &UppercaseSender{transport: transport}
}

// Uppercase - Execute transaction
func (s *UppercaseSender) Upper(ctx *actor.Context, arg UppercaseRequest) (*UppercaseResponse, error) {

	var sizer msgpack.Sizer
	size_enc := &sizer
	arg.MEncode(size_enc)
	buf := make([]byte, sizer.Len())

	var encoder = msgpack.NewEncoder(buf)
	enc := &encoder
	arg.MEncode(enc)

	out_buf, _ := s.transport.Send(ctx, actor.Message{Method: "Uppercase.Upper", Arg: buf})
	d := msgpack.NewDecoder(out_buf)
	resp, err_ := MDecodeUppercaseResponse(&d)
	if err_ != nil {
		return nil, err_
	}
	return &resp, nil
}

// This file is generated automatically using wasmcloud/weld-codegen 0.4.6
