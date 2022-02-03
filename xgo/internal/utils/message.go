package utils

import "google.golang.org/protobuf/compiler/protogen"

func LoadValidMessages(messages []*protogen.Message) []*protogen.Message {
	var loadMessages func(messages []*protogen.Message)

	validMessages := make([]*protogen.Message, 0)
	loadMessages = func(messages []*protogen.Message) {
		for _, msg := range messages {
			if msg.Desc.IsMapEntry() {
				// Ignore protobuf's map entry.
				continue
			}
			validMessages = append(validMessages, msg)
			loadMessages(msg.Messages)
		}
	}

	loadMessages(messages)
	return validMessages
}
