package ginats

import (
	"context"

	gieventbus "github.com/b2wdigital/goignite/eventbus"
	gilog "github.com/b2wdigital/goignite/log"
	"github.com/nats-io/nats.go"
)

const (
	TopicConn = "topic:nats:conn"
)

func NewConnection(ctx context.Context, options *Options) (*nats.Conn, error) {

	l := gilog.FromContext(ctx)

	conn, err := nats.Connect(
		options.Url,
		nats.MaxReconnects(options.MaxReconnects),
		nats.ReconnectWait(options.ReconnectWait),
		nats.DisconnectErrHandler(disconnectedErrHandler),
		nats.ReconnectHandler(reconnectedHandler),
		nats.ClosedHandler(closedHandler),
	)

	if err != nil {
		return nil, err
	}

	gieventbus.Publish(TopicConn, conn)

	l.Infof("Connected to NATS server: %s", options.Url)

	return conn, nil
}

func NewDefaultConnection(ctx context.Context) (*nats.Conn, error) {

	l := gilog.FromContext(ctx)

	o, err := DefaultOptions()
	if err != nil {
		l.Fatalf(err.Error())
	}

	return NewConnection(ctx, o)
}

func disconnectedErrHandler(nc *nats.Conn, err error) {
	gilog.Errorf("Disconnected due to:%s, will attempt reconnects for %.0fm", err)
}

func reconnectedHandler(nc *nats.Conn) {
	gilog.Warnf("Reconnected [%s]", nc.ConnectedUrl())
}

func closedHandler(nc *nats.Conn) {
	gilog.Errorf("Exiting: %v", nc.LastError())
}
