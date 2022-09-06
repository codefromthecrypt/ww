package pubsub

import (
	"context"

	capnp "capnproto.org/go/capnp/v3"

	chan_api "github.com/wetware/ww/internal/api/channel"
	api "github.com/wetware/ww/internal/api/pubsub"
)

type PubSub api.PubSub

func (ps PubSub) Join(ctx context.Context, topic string) (FutureTopic, capnp.ReleaseFunc) {
	f, release := (api.PubSub)(ps).Join(ctx, func(ps api.PubSub_join_Params) error {
		return ps.SetName(topic)
	})

	return FutureTopic(f), release
}

func (ps PubSub) AddRef() PubSub {
	return PubSub(capnp.Client(ps).AddRef())
}

func (ps PubSub) Release() {
	capnp.Client(ps).Release()
}

type FutureTopic api.PubSub_join_Results_Future

func (ft FutureTopic) Topic() Topic {
	return Topic(api.PubSub_join_Results_Future(ft).Topic())
}

func (ft FutureTopic) Struct() (Topic, error) {
	res, err := (api.PubSub_join_Results_Future)(ft).Struct()
	if err != nil {
		return Topic{}, err
	}

	return Topic(res.Topic()), nil
}

type Topic api.Topic

func (t Topic) Name(ctx context.Context) (string, error) {
	f, release := (api.Topic)(t).Name(ctx, nil)
	defer release()

	res, err := f.Struct()
	if err != nil {
		return "", err
	}

	return res.Name()
}

func (t Topic) Publish(ctx context.Context, b []byte) error {
	f, release := (api.Topic)(t).Publish(ctx, func(ps api.Topic_publish_Params) error {
		return ps.SetMsg(b)
	})
	defer release()

	_, err := f.Struct()
	return err
}

func (t Topic) Subscribe(ctx context.Context, ch chan<- []byte) (release capnp.ReleaseFunc, err error) {
	h := chan_api.Sender_ServerToClient(handler{
		ms:      ch,
		release: t.AddRef().Release,
	})

	f, release := api.Topic(t).Subscribe(ctx, sender(h, cap(ch)))
	defer release()

	if _, err = f.Struct(); err == nil {
		release = h.Release
	}

	return
}

func (t Topic) AddRef() Topic {
	return Topic(api.Topic(t).AddRef())
}

func (t Topic) Release() {
	capnp.Client(t).Release()
}

func sender(s chan_api.Sender, bufferSize int) func(api.Topic_subscribe_Params) error {
	return func(ps api.Topic_subscribe_Params) error {
		opts, err := ps.NewOpts()
		if err != nil {
			return err
		}
		opts.SetBufferSize(int64(bufferSize))
		return ps.SetChan(s)
	}
}

type handler struct {
	ms      chan<- []byte
	release capnp.ReleaseFunc
}

func (h handler) Shutdown() {
	close(h.ms)
	h.release()
}

func (h handler) Send(ctx context.Context, call chan_api.Sender_send) error {
	ptr, err := call.Args().Value()
	if err != nil {
		return err
	}

	select {
	case h.ms <- ptr.Data():
		return nil // fast path
	default:
	}

	// Slow path.  Spawn goroutine and block.  The FlowLimiter will limit the
	// number of outstanding calls to Handle.
	call.Ack()

	select {
	case h.ms <- ptr.Data():
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}
