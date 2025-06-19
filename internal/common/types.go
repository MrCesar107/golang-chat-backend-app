package common

type Client struct {
	Hub  *Hub
	Conn interface{}
	Send chan []byte
}

type Hub struct {
	Clients    map[*Client]bool
	Broadcast  chan []byte
	Register   chan *Client
	Unregister chan *Client
}
