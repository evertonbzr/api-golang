package queue

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/jetstream"
)

var natsConnection *nats.Conn
var natsJetStream jetstream.JetStream

func InitNatsClient(ctx context.Context, uri string, name string) (*nats.Conn, error) {
	if natsConnection != nil {
		return natsConnection, nil
	}

	nc, err := nats.Connect(uri)
	if err != nil {
		log.Fatal("Error connecting to NATS server", err)
		return nil, err
	}

	if natsJetStream != nil {
		return nc, nil
	}

	js, err := jetstream.New(nc)
	if err != nil {
		log.Fatal("Error connecting to JetStream", err)
		return nil, err
	}

	natsConnection = nc
	natsJetStream = js

	natsJetStream.CreateOrUpdateStream(ctx, jetstream.StreamConfig{
		Name:     name,
		Subjects: []string{fmt.Sprintf("%s.*", name)},
	})

	return nc, nil
}

func GetNatsConnection() *nats.Conn {
	if natsConnection == nil {
		panic("NATS client not initialized")
	}
	return natsConnection
}

func GetNatsJetStream() jetstream.JetStream {
	if natsJetStream == nil {
		panic("NATS JetStream not initialized")
	}
	return natsJetStream
}

func CloseNatsConnection() {
	if natsConnection != nil {
		natsConnection.Close()
	}
}

func PublishJetStreamMessage(ctx context.Context, subject string, data interface{}) error {
	dataBytes, err := json.Marshal(data)
	if err != nil {
		return err
	}

	_, err = GetNatsJetStream().Publish(ctx, subject, dataBytes)
	if err != nil {
		return err
	}

	return err
}
