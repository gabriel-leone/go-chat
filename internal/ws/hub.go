package ws

type Hub struct {
	Rooms      map[string]*Room
	Register   chan *Client
	Unregister chan *Client
	Broadcast  chan *Message
}

type Room struct {
	ID      string             `json:"id"`
	Name    string             `json:"name"`
	Clients map[string]*Client `json:"clients"`
}

func NewHub() *Hub {
	return &Hub{
		Rooms:      make(map[string]*Room),
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Broadcast:  make(chan *Message, 5),
	}
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.Register:
			h.registerClient(client)
		case client := <-h.Unregister:
			h.unregisterClient(client)
		case message := <-h.Broadcast:
			h.broadcastMessage(message)
		}
	}
}

func (h *Hub) registerClient(client *Client) {
	room, exists := h.Rooms[client.RoomID]
	if !exists {
		room = &Room{
			Clients: make(map[string]*Client),
		}
		h.Rooms[client.RoomID] = room
	}
	room.Clients[client.ID] = client
}

func (h *Hub) unregisterClient(client *Client) {
	room, exists := h.Rooms[client.RoomID]
	if exists {
		_, clientExists := room.Clients[client.ID]
		if clientExists {
			h.Broadcast <- &Message{
				Content:  client.Username + " has left the room",
				RoomID:   client.RoomID,
				Username: client.Username,
			}
			delete(room.Clients, client.ID)
			close(client.Message)
		}
	}
}

func (h *Hub) broadcastMessage(message *Message) {
	room, exists := h.Rooms[message.RoomID]
	if exists {
		for _, client := range room.Clients {
			client.Message <- message
		}
	}
}
