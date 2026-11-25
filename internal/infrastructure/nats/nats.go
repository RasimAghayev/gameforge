package nats

import (
	"context"
	"log/slog"

	"github.com/nats-io/nats.go"
)

type Publisher struct {
	conn *nats.Conn
	logger *slog.Logger
}

func NewPublisher(url string, logger *slog.Logger) (*Publisher, error) {
	conn, err := nats.Connect(url)
	if err != nil {
		return nil, err
	}
	return &Publisher{conn: conn, logger: logger}, nil
}

func (p *Publisher) Publish(subject string, data []byte) error {
	return p.conn.Publish(subject, data)
}

func (p *Publisher) Close() {
	if p.conn != nil {
		p.conn.Close()
	}
}

var _ = context.Background
