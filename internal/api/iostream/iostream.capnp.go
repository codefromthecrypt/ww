// Code generated by capnpc-go. DO NOT EDIT.

package iostream

import (
	capnp "capnproto.org/go/capnp/v3"
	text "capnproto.org/go/capnp/v3/encoding/text"
	schemas "capnproto.org/go/capnp/v3/schemas"
	server "capnproto.org/go/capnp/v3/server"
	stream "capnproto.org/go/capnp/v3/std/capnp/stream"
	context "context"
	channel "github.com/wetware/ww/internal/api/channel"
)

type Stream struct{ Client capnp.Client }

// Stream_TypeID is the unique identifier for the type Stream.
const Stream_TypeID = 0x800fee1ed6b441e2

func (c Stream) Send(ctx context.Context, params func(channel.Sender_send_Params) error) (stream.StreamResult_Future, capnp.ReleaseFunc) {
	s := capnp.Send{
		Method: capnp.Method{
			InterfaceID:   0xe8bbed1438ea16ee,
			MethodID:      0,
			InterfaceName: "channel.capnp:Sender",
			MethodName:    "send",
		},
	}
	if params != nil {
		s.ArgsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 1}
		s.PlaceArgs = func(s capnp.Struct) error { return params(channel.Sender_send_Params{Struct: s}) }
	}
	ans, release := c.Client.SendCall(ctx, s)
	return stream.StreamResult_Future{Future: ans.Future()}, release
}
func (c Stream) Close(ctx context.Context, params func(channel.Closer_close_Params) error) (channel.Closer_close_Results_Future, capnp.ReleaseFunc) {
	s := capnp.Send{
		Method: capnp.Method{
			InterfaceID:   0xfad0e4b80d3779c3,
			MethodID:      0,
			InterfaceName: "channel.capnp:Closer",
			MethodName:    "close",
		},
	}
	if params != nil {
		s.ArgsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 0}
		s.PlaceArgs = func(s capnp.Struct) error { return params(channel.Closer_close_Params{Struct: s}) }
	}
	ans, release := c.Client.SendCall(ctx, s)
	return channel.Closer_close_Results_Future{Future: ans.Future()}, release
}

func (c Stream) AddRef() Stream {
	return Stream{
		Client: c.Client.AddRef(),
	}
}

func (c Stream) Release() {
	c.Client.Release()
}

// A Stream_Server is a Stream with a local implementation.
type Stream_Server interface {
	Send(context.Context, channel.Sender_send) error

	Close(context.Context, channel.Closer_close) error
}

// Stream_NewServer creates a new Server from an implementation of Stream_Server.
func Stream_NewServer(s Stream_Server, policy *server.Policy) *server.Server {
	c, _ := s.(server.Shutdowner)
	return server.New(Stream_Methods(nil, s), s, c, policy)
}

// Stream_ServerToClient creates a new Client from an implementation of Stream_Server.
// The caller is responsible for calling Release on the returned Client.
func Stream_ServerToClient(s Stream_Server, policy *server.Policy) Stream {
	return Stream{Client: capnp.NewClient(Stream_NewServer(s, policy))}
}

// Stream_Methods appends Methods to a slice that invoke the methods on s.
// This can be used to create a more complicated Server.
func Stream_Methods(methods []server.Method, s Stream_Server) []server.Method {
	if cap(methods) == 0 {
		methods = make([]server.Method, 0, 2)
	}

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xe8bbed1438ea16ee,
			MethodID:      0,
			InterfaceName: "channel.capnp:Sender",
			MethodName:    "send",
		},
		Impl: func(ctx context.Context, call *server.Call) error {
			return s.Send(ctx, channel.Sender_send{call})
		},
	})

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xfad0e4b80d3779c3,
			MethodID:      0,
			InterfaceName: "channel.capnp:Closer",
			MethodName:    "close",
		},
		Impl: func(ctx context.Context, call *server.Call) error {
			return s.Close(ctx, channel.Closer_close{call})
		},
	})

	return methods
}

// Stream_List is a list of Stream.
type Stream_List = capnp.CapList[Stream]

// NewStream creates a new list of Stream.
func NewStream_List(s *capnp.Segment, sz int32) (Stream_List, error) {
	l, err := capnp.NewPointerList(s, sz)
	return capnp.CapList[Stream](l), err
}

type Provider struct{ Client capnp.Client }

// Provider_TypeID is the unique identifier for the type Provider.
const Provider_TypeID = 0xec225e7f00ef3b55

func (c Provider) Provide(ctx context.Context, params func(Provider_provide_Params) error) (Provider_provide_Results_Future, capnp.ReleaseFunc) {
	s := capnp.Send{
		Method: capnp.Method{
			InterfaceID:   0xec225e7f00ef3b55,
			MethodID:      0,
			InterfaceName: "iostream.capnp:Provider",
			MethodName:    "provide",
		},
	}
	if params != nil {
		s.ArgsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 1}
		s.PlaceArgs = func(s capnp.Struct) error { return params(Provider_provide_Params{Struct: s}) }
	}
	ans, release := c.Client.SendCall(ctx, s)
	return Provider_provide_Results_Future{Future: ans.Future()}, release
}

func (c Provider) AddRef() Provider {
	return Provider{
		Client: c.Client.AddRef(),
	}
}

func (c Provider) Release() {
	c.Client.Release()
}

// A Provider_Server is a Provider with a local implementation.
type Provider_Server interface {
	Provide(context.Context, Provider_provide) error
}

// Provider_NewServer creates a new Server from an implementation of Provider_Server.
func Provider_NewServer(s Provider_Server, policy *server.Policy) *server.Server {
	c, _ := s.(server.Shutdowner)
	return server.New(Provider_Methods(nil, s), s, c, policy)
}

// Provider_ServerToClient creates a new Client from an implementation of Provider_Server.
// The caller is responsible for calling Release on the returned Client.
func Provider_ServerToClient(s Provider_Server, policy *server.Policy) Provider {
	return Provider{Client: capnp.NewClient(Provider_NewServer(s, policy))}
}

// Provider_Methods appends Methods to a slice that invoke the methods on s.
// This can be used to create a more complicated Server.
func Provider_Methods(methods []server.Method, s Provider_Server) []server.Method {
	if cap(methods) == 0 {
		methods = make([]server.Method, 0, 1)
	}

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xec225e7f00ef3b55,
			MethodID:      0,
			InterfaceName: "iostream.capnp:Provider",
			MethodName:    "provide",
		},
		Impl: func(ctx context.Context, call *server.Call) error {
			return s.Provide(ctx, Provider_provide{call})
		},
	})

	return methods
}

// Provider_provide holds the state for a server call to Provider.provide.
// See server.Call for documentation.
type Provider_provide struct {
	*server.Call
}

// Args returns the call's arguments.
func (c Provider_provide) Args() Provider_provide_Params {
	return Provider_provide_Params{Struct: c.Call.Args()}
}

// AllocResults allocates the results struct.
func (c Provider_provide) AllocResults() (Provider_provide_Results, error) {
	r, err := c.Call.AllocResults(capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return Provider_provide_Results{Struct: r}, err
}

// Provider_List is a list of Provider.
type Provider_List = capnp.CapList[Provider]

// NewProvider creates a new list of Provider.
func NewProvider_List(s *capnp.Segment, sz int32) (Provider_List, error) {
	l, err := capnp.NewPointerList(s, sz)
	return capnp.CapList[Provider](l), err
}

type Provider_provide_Params struct{ capnp.Struct }

// Provider_provide_Params_TypeID is the unique identifier for the type Provider_provide_Params.
const Provider_provide_Params_TypeID = 0xbc8b7aa049d95800

func NewProvider_provide_Params(s *capnp.Segment) (Provider_provide_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return Provider_provide_Params{st}, err
}

func NewRootProvider_provide_Params(s *capnp.Segment) (Provider_provide_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return Provider_provide_Params{st}, err
}

func ReadRootProvider_provide_Params(msg *capnp.Message) (Provider_provide_Params, error) {
	root, err := msg.Root()
	return Provider_provide_Params{root.Struct()}, err
}

func (s Provider_provide_Params) String() string {
	str, _ := text.Marshal(0xbc8b7aa049d95800, s.Struct)
	return str
}

func (s Provider_provide_Params) Stream() Stream {
	p, _ := s.Struct.Ptr(0)
	return Stream{Client: p.Interface().Client()}
}

func (s Provider_provide_Params) HasStream() bool {
	return s.Struct.HasPtr(0)
}

func (s Provider_provide_Params) SetStream(v Stream) error {
	if !v.Client.IsValid() {
		return s.Struct.SetPtr(0, capnp.Ptr{})
	}
	seg := s.Segment()
	in := capnp.NewInterface(seg, seg.Message().AddCap(v.Client))
	return s.Struct.SetPtr(0, in.ToPtr())
}

// Provider_provide_Params_List is a list of Provider_provide_Params.
type Provider_provide_Params_List = capnp.StructList[Provider_provide_Params]

// NewProvider_provide_Params creates a new list of Provider_provide_Params.
func NewProvider_provide_Params_List(s *capnp.Segment, sz int32) (Provider_provide_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	return capnp.StructList[Provider_provide_Params]{List: l}, err
}

// Provider_provide_Params_Future is a wrapper for a Provider_provide_Params promised by a client call.
type Provider_provide_Params_Future struct{ *capnp.Future }

func (p Provider_provide_Params_Future) Struct() (Provider_provide_Params, error) {
	s, err := p.Future.Struct()
	return Provider_provide_Params{s}, err
}

func (p Provider_provide_Params_Future) Stream() Stream {
	return Stream{Client: p.Future.Field(0, nil).Client()}
}

type Provider_provide_Results struct{ capnp.Struct }

// Provider_provide_Results_TypeID is the unique identifier for the type Provider_provide_Results.
const Provider_provide_Results_TypeID = 0xfcf105dc8b5ae862

func NewProvider_provide_Results(s *capnp.Segment) (Provider_provide_Results, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return Provider_provide_Results{st}, err
}

func NewRootProvider_provide_Results(s *capnp.Segment) (Provider_provide_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return Provider_provide_Results{st}, err
}

func ReadRootProvider_provide_Results(msg *capnp.Message) (Provider_provide_Results, error) {
	root, err := msg.Root()
	return Provider_provide_Results{root.Struct()}, err
}

func (s Provider_provide_Results) String() string {
	str, _ := text.Marshal(0xfcf105dc8b5ae862, s.Struct)
	return str
}

// Provider_provide_Results_List is a list of Provider_provide_Results.
type Provider_provide_Results_List = capnp.StructList[Provider_provide_Results]

// NewProvider_provide_Results creates a new list of Provider_provide_Results.
func NewProvider_provide_Results_List(s *capnp.Segment, sz int32) (Provider_provide_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0}, sz)
	return capnp.StructList[Provider_provide_Results]{List: l}, err
}

// Provider_provide_Results_Future is a wrapper for a Provider_provide_Results promised by a client call.
type Provider_provide_Results_Future struct{ *capnp.Future }

func (p Provider_provide_Results_Future) Struct() (Provider_provide_Results, error) {
	s, err := p.Future.Struct()
	return Provider_provide_Results{s}, err
}

const schema_89c985e63e991441 = "x\xdat\xd0?HBQ\x14\x06\xf0\xef\xbcs\x9f\xaf" +
	"E\xe4r\x0d\x9b\x12\xa25I\xa2\xc5 \xff,\x11D" +
	"\xbc[\x04\xd5\x10\xbc\xcaA\xc8\x94\xf7,(\x10kh" +
	"\x08\x89\xe6\x9aZ\"Z\xa3\xb1\xa9\xad\xb1%\x08\x9a\x1a" +
	"\"!\x8a\xa6\x96\xe8\xc53Q\x11\xda\xce\xbd\x1c~\xf7" +
	"\xbb\xdfh?eD2|\x1c\x82\xa1g\xcc\x90\xff\x9c" +
	"\xbd~\x18|\x8f\xecAF\xd8\xcfFO&_\x0e\xee" +
	"\x0e\x01R#|\xa5\xc6\xd9\x02T\x92-\x95\xe4\x18\xf0" +
	"\xb3\xf88}\xb6[\xbf\x91\x03\x04\x98d\x01c\x9aS" +
	"\x04RK\x9c\x06}-L|\xd4V\x86\xdez\xa5\x1d" +
	"\xbeU\xfbM\xa9\xcaS\xea<\x98\xfc\xd5\xd7\xe5\xfa\x93" +
	"\xf9\xf9\x8d\xa6%\x02\xea\x88s\x84K\xbfP\xf2*n" +
	"\xde)Rb\xcd)o\x96S\xf3\xf1\xe6\xd1&\xb2\xd9" +
	"\xd4\x82\xc8\x9fmD\xab\xa7\xf7\x17\x0ddHR\\\x0b" +
	"\xa3\xeb\x0a\x90\x14\x0b\xb6\x82u\xa20\x8c6\xc8-\xd0" +
	"vK\xdb\x85\xf5\xbc\x9b(\xff\x0d\xc3i\xdbq\x9d\xa2" +
	"\xa7\x05\x0b@\x10 \xc3)@\xf71\xe9\xa8A\xe9V" +
	"\x1c\xd9\xa9\x0aD\x12\xd4\x86\x8d\x1e\x186\x91\x16lv" +
	"\x0a\xeb\xfa\xae\xcc\xc1\x90\xa6Uk=\x9e!\x9b:\x92" +
	"\xf8/\xe2\\\xde\xdb\xda\xa8x\xf8\x0d\x00\x00\xff\xff\xb0" +
	"=\x85\xba"

func init() {
	schemas.Register(schema_89c985e63e991441,
		0x800fee1ed6b441e2,
		0xbc8b7aa049d95800,
		0xec225e7f00ef3b55,
		0xfcf105dc8b5ae862)
}
