// Code generated by protoc-gen-go.
// source: envelope.proto
// DO NOT EDIT!

package loggregator

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Log_Type int32

const (
	Log_OUT Log_Type = 0
	Log_ERR Log_Type = 1
)

var Log_Type_name = map[int32]string{
	0: "OUT",
	1: "ERR",
}
var Log_Type_value = map[string]int32{
	"OUT": 0,
	"ERR": 1,
}

func (x Log_Type) String() string {
	return proto.EnumName(Log_Type_name, int32(x))
}
func (Log_Type) EnumDescriptor() ([]byte, []int) { return fileDescriptor1, []int{2, 0} }

type Envelope struct {
	Timestamp  int64             `protobuf:"varint,1,opt,name=timestamp" json:"timestamp,omitempty"`
	SourceUuid string            `protobuf:"bytes,2,opt,name=source_uuid,json=sourceUuid" json:"source_uuid,omitempty"`
	Tags       map[string]*Value `protobuf:"bytes,3,rep,name=tags" json:"tags,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	// Types that are valid to be assigned to Message:
	//	*Envelope_Log
	//	*Envelope_Counter
	//	*Envelope_Gauge
	//	*Envelope_CorrelatedGauge
	//	*Envelope_Timer
	Message isEnvelope_Message `protobuf_oneof:"message"`
}

func (m *Envelope) Reset()                    { *m = Envelope{} }
func (m *Envelope) String() string            { return proto.CompactTextString(m) }
func (*Envelope) ProtoMessage()               {}
func (*Envelope) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

type isEnvelope_Message interface {
	isEnvelope_Message()
}

type Envelope_Log struct {
	Log *Log `protobuf:"bytes,4,opt,name=log,oneof"`
}
type Envelope_Counter struct {
	Counter *Counter `protobuf:"bytes,5,opt,name=counter,oneof"`
}
type Envelope_Gauge struct {
	Gauge *Gauge `protobuf:"bytes,6,opt,name=gauge,oneof"`
}
type Envelope_CorrelatedGauge struct {
	CorrelatedGauge *CorrelatedGauge `protobuf:"bytes,7,opt,name=correlated_gauge,json=correlatedGauge,oneof"`
}
type Envelope_Timer struct {
	Timer *Timer `protobuf:"bytes,8,opt,name=timer,oneof"`
}

func (*Envelope_Log) isEnvelope_Message()             {}
func (*Envelope_Counter) isEnvelope_Message()         {}
func (*Envelope_Gauge) isEnvelope_Message()           {}
func (*Envelope_CorrelatedGauge) isEnvelope_Message() {}
func (*Envelope_Timer) isEnvelope_Message()           {}

func (m *Envelope) GetMessage() isEnvelope_Message {
	if m != nil {
		return m.Message
	}
	return nil
}

func (m *Envelope) GetTags() map[string]*Value {
	if m != nil {
		return m.Tags
	}
	return nil
}

func (m *Envelope) GetLog() *Log {
	if x, ok := m.GetMessage().(*Envelope_Log); ok {
		return x.Log
	}
	return nil
}

func (m *Envelope) GetCounter() *Counter {
	if x, ok := m.GetMessage().(*Envelope_Counter); ok {
		return x.Counter
	}
	return nil
}

func (m *Envelope) GetGauge() *Gauge {
	if x, ok := m.GetMessage().(*Envelope_Gauge); ok {
		return x.Gauge
	}
	return nil
}

func (m *Envelope) GetCorrelatedGauge() *CorrelatedGauge {
	if x, ok := m.GetMessage().(*Envelope_CorrelatedGauge); ok {
		return x.CorrelatedGauge
	}
	return nil
}

func (m *Envelope) GetTimer() *Timer {
	if x, ok := m.GetMessage().(*Envelope_Timer); ok {
		return x.Timer
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Envelope) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Envelope_OneofMarshaler, _Envelope_OneofUnmarshaler, _Envelope_OneofSizer, []interface{}{
		(*Envelope_Log)(nil),
		(*Envelope_Counter)(nil),
		(*Envelope_Gauge)(nil),
		(*Envelope_CorrelatedGauge)(nil),
		(*Envelope_Timer)(nil),
	}
}

func _Envelope_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Envelope)
	// message
	switch x := m.Message.(type) {
	case *Envelope_Log:
		b.EncodeVarint(4<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Log); err != nil {
			return err
		}
	case *Envelope_Counter:
		b.EncodeVarint(5<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Counter); err != nil {
			return err
		}
	case *Envelope_Gauge:
		b.EncodeVarint(6<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Gauge); err != nil {
			return err
		}
	case *Envelope_CorrelatedGauge:
		b.EncodeVarint(7<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.CorrelatedGauge); err != nil {
			return err
		}
	case *Envelope_Timer:
		b.EncodeVarint(8<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Timer); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("Envelope.Message has unexpected type %T", x)
	}
	return nil
}

func _Envelope_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Envelope)
	switch tag {
	case 4: // message.log
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Log)
		err := b.DecodeMessage(msg)
		m.Message = &Envelope_Log{msg}
		return true, err
	case 5: // message.counter
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Counter)
		err := b.DecodeMessage(msg)
		m.Message = &Envelope_Counter{msg}
		return true, err
	case 6: // message.gauge
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Gauge)
		err := b.DecodeMessage(msg)
		m.Message = &Envelope_Gauge{msg}
		return true, err
	case 7: // message.correlated_gauge
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(CorrelatedGauge)
		err := b.DecodeMessage(msg)
		m.Message = &Envelope_CorrelatedGauge{msg}
		return true, err
	case 8: // message.timer
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Timer)
		err := b.DecodeMessage(msg)
		m.Message = &Envelope_Timer{msg}
		return true, err
	default:
		return false, nil
	}
}

func _Envelope_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Envelope)
	// message
	switch x := m.Message.(type) {
	case *Envelope_Log:
		s := proto.Size(x.Log)
		n += proto.SizeVarint(4<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Envelope_Counter:
		s := proto.Size(x.Counter)
		n += proto.SizeVarint(5<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Envelope_Gauge:
		s := proto.Size(x.Gauge)
		n += proto.SizeVarint(6<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Envelope_CorrelatedGauge:
		s := proto.Size(x.CorrelatedGauge)
		n += proto.SizeVarint(7<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Envelope_Timer:
		s := proto.Size(x.Timer)
		n += proto.SizeVarint(8<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type Value struct {
	// Types that are valid to be assigned to Data:
	//	*Value_Text
	//	*Value_Integer
	//	*Value_Decimal
	Data isValue_Data `protobuf_oneof:"data"`
}

func (m *Value) Reset()                    { *m = Value{} }
func (m *Value) String() string            { return proto.CompactTextString(m) }
func (*Value) ProtoMessage()               {}
func (*Value) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

type isValue_Data interface {
	isValue_Data()
}

type Value_Text struct {
	Text string `protobuf:"bytes,1,opt,name=text,oneof"`
}
type Value_Integer struct {
	Integer int64 `protobuf:"varint,2,opt,name=integer,oneof"`
}
type Value_Decimal struct {
	Decimal float64 `protobuf:"fixed64,3,opt,name=decimal,oneof"`
}

func (*Value_Text) isValue_Data()    {}
func (*Value_Integer) isValue_Data() {}
func (*Value_Decimal) isValue_Data() {}

func (m *Value) GetData() isValue_Data {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *Value) GetText() string {
	if x, ok := m.GetData().(*Value_Text); ok {
		return x.Text
	}
	return ""
}

func (m *Value) GetInteger() int64 {
	if x, ok := m.GetData().(*Value_Integer); ok {
		return x.Integer
	}
	return 0
}

func (m *Value) GetDecimal() float64 {
	if x, ok := m.GetData().(*Value_Decimal); ok {
		return x.Decimal
	}
	return 0
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Value) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Value_OneofMarshaler, _Value_OneofUnmarshaler, _Value_OneofSizer, []interface{}{
		(*Value_Text)(nil),
		(*Value_Integer)(nil),
		(*Value_Decimal)(nil),
	}
}

func _Value_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Value)
	// data
	switch x := m.Data.(type) {
	case *Value_Text:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.Text)
	case *Value_Integer:
		b.EncodeVarint(2<<3 | proto.WireVarint)
		b.EncodeVarint(uint64(x.Integer))
	case *Value_Decimal:
		b.EncodeVarint(3<<3 | proto.WireFixed64)
		b.EncodeFixed64(math.Float64bits(x.Decimal))
	case nil:
	default:
		return fmt.Errorf("Value.Data has unexpected type %T", x)
	}
	return nil
}

func _Value_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Value)
	switch tag {
	case 1: // data.text
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Data = &Value_Text{x}
		return true, err
	case 2: // data.integer
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.Data = &Value_Integer{int64(x)}
		return true, err
	case 3: // data.decimal
		if wire != proto.WireFixed64 {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeFixed64()
		m.Data = &Value_Decimal{math.Float64frombits(x)}
		return true, err
	default:
		return false, nil
	}
}

func _Value_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Value)
	// data
	switch x := m.Data.(type) {
	case *Value_Text:
		n += proto.SizeVarint(1<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.Text)))
		n += len(x.Text)
	case *Value_Integer:
		n += proto.SizeVarint(2<<3 | proto.WireVarint)
		n += proto.SizeVarint(uint64(x.Integer))
	case *Value_Decimal:
		n += proto.SizeVarint(3<<3 | proto.WireFixed64)
		n += 8
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type Log struct {
	Payload []byte   `protobuf:"bytes,1,opt,name=payload,proto3" json:"payload,omitempty"`
	Type    Log_Type `protobuf:"varint,2,opt,name=type,enum=loggregator.Log_Type" json:"type,omitempty"`
}

func (m *Log) Reset()                    { *m = Log{} }
func (m *Log) String() string            { return proto.CompactTextString(m) }
func (*Log) ProtoMessage()               {}
func (*Log) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{2} }

type Counter struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// Types that are valid to be assigned to Value:
	//	*Counter_Delta
	//	*Counter_Total
	Value isCounter_Value `protobuf_oneof:"value"`
}

func (m *Counter) Reset()                    { *m = Counter{} }
func (m *Counter) String() string            { return proto.CompactTextString(m) }
func (*Counter) ProtoMessage()               {}
func (*Counter) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{3} }

type isCounter_Value interface {
	isCounter_Value()
}

type Counter_Delta struct {
	Delta uint64 `protobuf:"varint,2,opt,name=delta,oneof"`
}
type Counter_Total struct {
	Total uint64 `protobuf:"varint,3,opt,name=total,oneof"`
}

func (*Counter_Delta) isCounter_Value() {}
func (*Counter_Total) isCounter_Value() {}

func (m *Counter) GetValue() isCounter_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *Counter) GetDelta() uint64 {
	if x, ok := m.GetValue().(*Counter_Delta); ok {
		return x.Delta
	}
	return 0
}

func (m *Counter) GetTotal() uint64 {
	if x, ok := m.GetValue().(*Counter_Total); ok {
		return x.Total
	}
	return 0
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Counter) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Counter_OneofMarshaler, _Counter_OneofUnmarshaler, _Counter_OneofSizer, []interface{}{
		(*Counter_Delta)(nil),
		(*Counter_Total)(nil),
	}
}

func _Counter_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Counter)
	// value
	switch x := m.Value.(type) {
	case *Counter_Delta:
		b.EncodeVarint(2<<3 | proto.WireVarint)
		b.EncodeVarint(uint64(x.Delta))
	case *Counter_Total:
		b.EncodeVarint(3<<3 | proto.WireVarint)
		b.EncodeVarint(uint64(x.Total))
	case nil:
	default:
		return fmt.Errorf("Counter.Value has unexpected type %T", x)
	}
	return nil
}

func _Counter_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Counter)
	switch tag {
	case 2: // value.delta
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.Value = &Counter_Delta{x}
		return true, err
	case 3: // value.total
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.Value = &Counter_Total{x}
		return true, err
	default:
		return false, nil
	}
}

func _Counter_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Counter)
	// value
	switch x := m.Value.(type) {
	case *Counter_Delta:
		n += proto.SizeVarint(2<<3 | proto.WireVarint)
		n += proto.SizeVarint(uint64(x.Delta))
	case *Counter_Total:
		n += proto.SizeVarint(3<<3 | proto.WireVarint)
		n += proto.SizeVarint(uint64(x.Total))
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type Gauge struct {
	Name  string  `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Value float64 `protobuf:"fixed64,2,opt,name=value" json:"value,omitempty"`
}

func (m *Gauge) Reset()                    { *m = Gauge{} }
func (m *Gauge) String() string            { return proto.CompactTextString(m) }
func (*Gauge) ProtoMessage()               {}
func (*Gauge) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{4} }

type CorrelatedGauge struct {
	Metrics map[string]float64 `protobuf:"bytes,1,rep,name=metrics" json:"metrics,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"fixed64,2,opt,name=value"`
}

func (m *CorrelatedGauge) Reset()                    { *m = CorrelatedGauge{} }
func (m *CorrelatedGauge) String() string            { return proto.CompactTextString(m) }
func (*CorrelatedGauge) ProtoMessage()               {}
func (*CorrelatedGauge) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{5} }

func (m *CorrelatedGauge) GetMetrics() map[string]float64 {
	if m != nil {
		return m.Metrics
	}
	return nil
}

type Timer struct {
	Name  string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Start uint64 `protobuf:"varint,2,opt,name=start" json:"start,omitempty"`
	Stop  uint64 `protobuf:"varint,3,opt,name=stop" json:"stop,omitempty"`
}

func (m *Timer) Reset()                    { *m = Timer{} }
func (m *Timer) String() string            { return proto.CompactTextString(m) }
func (*Timer) ProtoMessage()               {}
func (*Timer) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{6} }

func init() {
	proto.RegisterType((*Envelope)(nil), "loggregator.Envelope")
	proto.RegisterType((*Value)(nil), "loggregator.Value")
	proto.RegisterType((*Log)(nil), "loggregator.Log")
	proto.RegisterType((*Counter)(nil), "loggregator.Counter")
	proto.RegisterType((*Gauge)(nil), "loggregator.Gauge")
	proto.RegisterType((*CorrelatedGauge)(nil), "loggregator.CorrelatedGauge")
	proto.RegisterType((*Timer)(nil), "loggregator.Timer")
	proto.RegisterEnum("loggregator.Log_Type", Log_Type_name, Log_Type_value)
}

func init() { proto.RegisterFile("envelope.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 539 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x7c, 0x54, 0xc1, 0x6a, 0xdb, 0x40,
	0x10, 0xb5, 0x22, 0xc9, 0x8a, 0xc7, 0x21, 0x31, 0x8b, 0x5b, 0x16, 0x13, 0x88, 0x11, 0x3d, 0x28,
	0x3d, 0x88, 0xd6, 0xb9, 0x94, 0x1c, 0x1d, 0x4c, 0x5d, 0x9a, 0x52, 0x58, 0x9c, 0xdc, 0x4a, 0xd8,
	0x4a, 0xcb, 0x22, 0x2a, 0x7b, 0xc5, 0x6a, 0x15, 0xea, 0xdf, 0xe8, 0x9f, 0xf5, 0x8f, 0xca, 0xce,
	0x4a, 0xad, 0x6d, 0xdc, 0xdc, 0x66, 0xe6, 0xbd, 0x99, 0xb7, 0x7a, 0x33, 0x08, 0xce, 0xc5, 0xe6,
	0x59, 0x94, 0xaa, 0x12, 0x69, 0xa5, 0x95, 0x51, 0x64, 0x58, 0x2a, 0x29, 0xb5, 0x90, 0xdc, 0x28,
	0x1d, 0xff, 0xf6, 0xe1, 0x74, 0xd1, 0xe2, 0xe4, 0x12, 0x06, 0xa6, 0x58, 0x8b, 0xda, 0xf0, 0x75,
	0x45, 0xbd, 0xa9, 0x97, 0xf8, 0xec, 0x5f, 0x81, 0x5c, 0xc1, 0xb0, 0x56, 0x8d, 0xce, 0xc4, 0x53,
	0xd3, 0x14, 0x39, 0x3d, 0x99, 0x7a, 0xc9, 0x80, 0x81, 0x2b, 0x3d, 0x34, 0x45, 0x4e, 0x6e, 0x20,
	0x30, 0x5c, 0xd6, 0xd4, 0x9f, 0xfa, 0xc9, 0x70, 0x76, 0x95, 0xee, 0xe8, 0xa4, 0x9d, 0x46, 0xba,
	0xe2, 0xb2, 0x5e, 0x6c, 0x8c, 0xde, 0x32, 0x24, 0x93, 0x37, 0xe0, 0x97, 0x4a, 0xd2, 0x60, 0xea,
	0x25, 0xc3, 0xd9, 0x68, 0xaf, 0xe7, 0x5e, 0xc9, 0x65, 0x8f, 0x59, 0x98, 0xbc, 0x83, 0x28, 0x53,
	0xcd, 0xc6, 0x08, 0x4d, 0x43, 0x64, 0x8e, 0xf7, 0x98, 0x77, 0x0e, 0x5b, 0xf6, 0x58, 0x47, 0x23,
	0x6f, 0x21, 0x94, 0xbc, 0x91, 0x82, 0xf6, 0x91, 0x4f, 0xf6, 0xf8, 0x1f, 0x2d, 0xb2, 0xec, 0x31,
	0x47, 0x21, 0x9f, 0x60, 0x94, 0x29, 0xad, 0x45, 0xc9, 0x8d, 0xc8, 0x9f, 0x5c, 0x5b, 0x84, 0x6d,
	0x97, 0x07, 0x32, 0x1d, 0xa9, 0x1b, 0x70, 0x91, 0xed, 0x97, 0xac, 0xac, 0x75, 0x4c, 0xd3, 0xd3,
	0x23, 0xb2, 0x2b, 0x8b, 0x58, 0x59, 0xa4, 0x4c, 0x3e, 0xc3, 0xe0, 0xaf, 0x1b, 0x64, 0x04, 0xfe,
	0x0f, 0xb1, 0x45, 0xd7, 0x07, 0xcc, 0x86, 0x24, 0x81, 0xf0, 0x99, 0x97, 0x8d, 0x40, 0xa7, 0x0f,
	0x47, 0x3d, 0x5a, 0x84, 0x39, 0xc2, 0xed, 0xc9, 0x07, 0x6f, 0x3e, 0x80, 0x68, 0x2d, 0xea, 0x9a,
	0x4b, 0x11, 0x7f, 0x83, 0x10, 0x61, 0x32, 0x86, 0xc0, 0x88, 0x9f, 0xc6, 0x0d, 0x5d, 0xf6, 0x18,
	0x66, 0x64, 0x02, 0x51, 0xb1, 0x31, 0x42, 0x0a, 0x8d, 0x93, 0x7d, 0xeb, 0x5a, 0x5b, 0xb0, 0x58,
	0x2e, 0xb2, 0x62, 0xcd, 0x4b, 0xea, 0x4f, 0xbd, 0xc4, 0xb3, 0x58, 0x5b, 0x98, 0xf7, 0x21, 0xc8,
	0xb9, 0xe1, 0x71, 0x0e, 0xfe, 0xbd, 0x92, 0x84, 0x42, 0x54, 0xf1, 0x6d, 0xa9, 0x78, 0x8e, 0xf3,
	0xcf, 0x58, 0x97, 0x92, 0x6b, 0x08, 0xcc, 0xb6, 0x72, 0xef, 0x3e, 0x9f, 0xbd, 0x3a, 0xdc, 0x69,
	0xba, 0xda, 0x56, 0x82, 0x21, 0x25, 0xa6, 0x10, 0xd8, 0x8c, 0x44, 0xe0, 0x7f, 0x7d, 0x58, 0x8d,
	0x7a, 0x36, 0x58, 0x30, 0x36, 0xf2, 0xe2, 0x47, 0x88, 0xda, 0xad, 0x12, 0x02, 0xc1, 0x86, 0xaf,
	0x45, 0xeb, 0x0d, 0xc6, 0xe4, 0x35, 0x84, 0xb9, 0x28, 0x0d, 0x47, 0x91, 0xc0, 0x7a, 0x8a, 0xa9,
	0xad, 0x1b, 0x65, 0xda, 0xe7, 0x63, 0x1d, 0xd3, 0x79, 0xd4, 0x9a, 0x19, 0xbf, 0x87, 0xd0, 0x6d,
	0xea, 0xd8, 0xd4, 0xf1, 0xae, 0xe5, 0x5e, 0x6b, 0x6f, 0xfc, 0xcb, 0x83, 0x8b, 0x83, 0xd5, 0x93,
	0x3b, 0x6b, 0xb7, 0xd1, 0x45, 0x56, 0x53, 0x0f, 0xcf, 0xfd, 0xfa, 0xa5, 0x4b, 0x49, 0xbf, 0x38,
	0xae, 0x3b, 0xfc, 0xae, 0x73, 0x72, 0x0b, 0x67, 0xbb, 0xc0, 0x91, 0x1b, 0x38, 0xfa, 0x20, 0xbb,
	0xef, 0x78, 0x01, 0x21, 0x9e, 0xd3, 0xff, 0xbe, 0xa3, 0x36, 0x5c, 0x1b, 0xe7, 0x0e, 0x73, 0x89,
	0x65, 0xd6, 0x46, 0x55, 0xce, 0x1a, 0x86, 0xf1, 0xf7, 0x3e, 0xfe, 0x13, 0x6e, 0xfe, 0x04, 0x00,
	0x00, 0xff, 0xff, 0xe4, 0xc6, 0x53, 0x4d, 0x25, 0x04, 0x00, 0x00,
}
