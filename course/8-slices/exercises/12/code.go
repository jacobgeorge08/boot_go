package main

type Message interface {
	Type() string
}

type TextMessage struct {
	Sender  string
	Content string
}

func (tm TextMessage) Type() string {
	return "text"
}

type MediaMessage struct {
	Sender    string
	MediaType string
	Content   string
}

func (mm MediaMessage) Type() string {
	return "media"
}

type LinkMessage struct {
	Sender  string
	URL     string
	Content string
}

func (lm LinkMessage) Type() string {
	return "link"
}

func filterMessages(messages []Message, filterType string) []Message {
	textMsg := []Message{}
	mediaMsg := []Message{}
	linkMsg := []Message{}

	for _, m := range messages {
		if m.Type() == "text" {
			textMsg = append(textMsg, m)
		} else if m.Type() == "media" {
			mediaMsg = append(mediaMsg, m)
		} else if m.Type() == "link" {
			linkMsg = append(linkMsg, m)
		}
	}

	if filterType == "text" {
		return textMsg
	} else if filterType == "media" {
		return mediaMsg
	} else {
		return linkMsg
	}
}
