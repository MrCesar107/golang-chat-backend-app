package hub

import "chatbackendapp/internal/common"

func NewHub() *common.Hub {
	return &common.Hub{
		Clients: make(map[*common.Client]bool),
		Broadcast: make(chan []byte),
		Register: make(chan *common.Client),
		Unregister: make(chan *common.Client),
	}
}

func Run(h *common.Hub) {
	for {
		select {
		case client := <- h.Register:
			h.Clients[client] = true
		case client := <- h.Unregister:
			if _, ok := h.Clients[client]; ok {
				delete(h.Clients, client)
				close(client.Send)
			}
		case message := <- h.Broadcast:
			for client := range h.Clients {
				select {
				case client.Send <- message:
				default:
					close(client.Send)
					delete(h.Clients, client)
				}
			}
		}
	}
}
