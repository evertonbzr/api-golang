package queue

import (
	"context"
	"fmt"
	"log"
	"log/slog"

	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/jetstream"
)

var (
	NatsConnection *nats.Conn
	JetStream      jetstream.JetStream
)

type QueueConfig struct {
	URI  string
	Name string
}

func connect(uri string) (*nats.Conn, jetstream.JetStream, error) {
	nc, err := nats.Connect(uri)

	if err != nil {
		log.Fatal("Error connecting to NATS server", err)
		return nil, nil, err
	}

	js, err := jetstream.New(nc)

	if err != nil {
		log.Fatal("Error connecting to JetStream", err)
		return nil, nil, err
	}

	return nc, js, nil
}

func Start(cfg *QueueConfig) (*nats.Conn, jetstream.JetStream, error) {
	nc, js, err := connect(cfg.URI)

	if err != nil {
		log.Fatal("Error connecting to NATS", err)
	}

	ctx := context.Background()

	slog.Info("Connected to NATS")

	js.CreateOrUpdateStream(ctx, jetstream.StreamConfig{
		Name:     cfg.Name,
		Subjects: []string{fmt.Sprintf("%s.*", cfg.Name)},
	})

	defer nc.Close()

	NatsConnection = nc
	JetStream = js

	return nc, js, nil
}
