package main

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"time"
)

type Event struct {
	ID        string
	PubKey    string
	CreatedAt time.Time
	Kind      int
	Tags      [][]string
	Content   string
	Sig       string
}

func (evt *Event) SerializeID() (string, error) {

	data := []interface{}{
		0,
		evt.PubKey,
		evt.CreatedAt.Unix(),
		evt.Kind,
		evt.Tags,
		evt.Content,
	}
	b, err := json.Marshal(data)
	if err != nil {
		return "", err
	}
	sha := sha256.Sum256(b)
	evt.ID = hex.EncodeToString(sha[:])

	return evt.ID, nil
}

func main() {

	event := &Event{
		PubKey:    "test-pubkey",
		CreatedAt: time.Now(),
		Kind:      1,
		Tags: [][]string{
			{"e", "test-id", "test-relay-url"},
			{"p", "test-key", "test-relay-url"},
		},
		Content: "test-content",
	}
	id, err := event.SerializeID()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Serialized ID:", id)
}
