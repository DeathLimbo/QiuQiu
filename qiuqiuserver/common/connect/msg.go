package connect

type Msg struct {
	head  *Head
	body  []byte
	Error error
}

type Head struct {
}
