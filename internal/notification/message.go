package notification

type MessageElement struct {
	Type Type
	Data interface{}
}

type MessageChain []*MessageElement

type Message struct {
	Title        *string
	MessageChain MessageChain
}

func NewCustomMessage(p Type, s interface{}) *MessageElement { // convert to byte[] anyway
	return &MessageElement{
		Type: p,
		Data: s,
	}
}

func NewImageMessage(s interface{}) *MessageElement { // byte[] of url
	return &MessageElement{
		Type: TypeImage,
		Data: s,
	}
}

func NewTextMessage(s string) *MessageElement {
	return &MessageElement{
		Type: TypeText,
		Data: s,
	}
}

func NewBinaryMessage(s []byte) *MessageElement {
	return &MessageElement{
		Type: TypeBinary,
		Data: s,
	}
}
