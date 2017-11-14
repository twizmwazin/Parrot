package proto

type StatusClientResponse struct {
	Response string
}

type StatusClientPong struct {
	Payload int64
}

type StatusServerRequest struct {
}

type StatusServerPing struct {
	Payload int64
}
