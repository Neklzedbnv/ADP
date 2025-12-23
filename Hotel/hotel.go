package Hotel

import "fmt"

type Room struct {
	RoomNumber    string
	Type          string
	PricePerNight float64
	IsOccupied    bool
}

type Hotel struct {
	Rooms map[string]Room
}

func NewHotel() Hotel {
	return Hotel{Rooms: make(map[string]Room)}
}

func (h *Hotel) AddRoom(room Room) {
	h.Rooms[room.RoomNumber] = room
	fmt.Println("Room added:", room.RoomNumber)
}

func (h *Hotel) CheckIn(number string) {
	room := h.Rooms[number]
	room.IsOccupied = true
	h.Rooms[number] = room
	fmt.Println("Checked in:", number)
}

func (h *Hotel) CheckOut(number string) {
	room := h.Rooms[number]
	room.IsOccupied = false
	h.Rooms[number] = room
	fmt.Println("Checked out:", number)
}

func (h *Hotel) ListVacantRooms() {
	fmt.Println("Vacant rooms:")
	for key := range h.Rooms {
		if h.Rooms[key].IsOccupied == false {
			fmt.Println(h.Rooms[key].RoomNumber)
		}
	}
}
