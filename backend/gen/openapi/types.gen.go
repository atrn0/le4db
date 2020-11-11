// Package Openapi provides primitives to interact the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen DO NOT EDIT.
package Openapi

import (
	"time"
)

// GuestsGetReservationsRes defines model for GuestsGetReservationsRes.
type GuestsGetReservationsRes struct {
	Reservations []GuestsReservation `json:"reservations"`
}

// GuestsGetRoomsRes defines model for GuestsGetRoomsRes.
type GuestsGetRoomsRes struct {
	Rooms []GuestsRoom `json:"rooms"`
}

// GuestsPostUsersReq defines model for GuestsPostUsersReq.
type GuestsPostUsersReq struct {
	Name   string `json:"name"`
	UserId string `json:"user_id"`
}

// GuestsReservation defines model for GuestsReservation.
type GuestsReservation struct {
	CheckIn  time.Time `json:"check_in"`
	CheckOut time.Time `json:"check_out"`
	Id       string    `json:"id"`
	RoomId   string    `json:"room_id"`
	RoomName string    `json:"room_name"`
}

// GuestsRoom defines model for GuestsRoom.
type GuestsRoom struct {
	HostId string `json:"host_id"`
	Id     string `json:"id"`
	Name   string `json:"name"`
	Price  int    `json:"price"`
}

// HostsGetReservationsRes defines model for HostsGetReservationsRes.
type HostsGetReservationsRes struct {
	Reservations []HostsReservation `json:"reservations"`
}

// HostsGetRoomsRes defines model for HostsGetRoomsRes.
type HostsGetRoomsRes struct {
	Rooms []HostsRoom `json:"rooms"`
}

// HostsPostRoomsReq defines model for HostsPostRoomsReq.
type HostsPostRoomsReq struct {
	Name  string `json:"name"`
	Price int    `json:"price"`
}

// HostsPutRoomsReq defines model for HostsPutRoomsReq.
type HostsPutRoomsReq struct {
	Name  string `json:"name"`
	Price int    `json:"price"`
}

// HostsReservation defines model for HostsReservation.
type HostsReservation struct {
	CheckIn   time.Time `json:"check_in"`
	CheckOut  time.Time `json:"check_out"`
	GuestId   string    `json:"guest_id"`
	GuestName string    `json:"guest_name"`
	Id        string    `json:"id"`
	RoomId    string    `json:"room_id"`
	RoomName  string    `json:"room_name"`
}

// HostsRoom defines model for HostsRoom.
type HostsRoom struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}

// PostReservationsReq defines model for PostReservationsReq.
type PostReservationsReq struct {
	CheckIn  time.Time `json:"check_in"`
	CheckOut time.Time `json:"check_out"`
	RoomId   string    `json:"room_id"`
}

// PostReservationsJSONBody defines parameters for PostReservations.
type PostReservationsJSONBody PostReservationsReq

// GuestsPostUsersJSONBody defines parameters for GuestsPostUsers.
type GuestsPostUsersJSONBody GuestsPostUsersReq

// HostsPostRoomsJSONBody defines parameters for HostsPostRooms.
type HostsPostRoomsJSONBody HostsPostRoomsReq

// HostsPutRoomsJSONBody defines parameters for HostsPutRooms.
type HostsPutRoomsJSONBody HostsPutRoomsReq

// PostReservationsRequestBody defines body for PostReservations for application/json ContentType.
type PostReservationsJSONRequestBody PostReservationsJSONBody

// GuestsPostUsersRequestBody defines body for GuestsPostUsers for application/json ContentType.
type GuestsPostUsersJSONRequestBody GuestsPostUsersJSONBody

// HostsPostRoomsRequestBody defines body for HostsPostRooms for application/json ContentType.
type HostsPostRoomsJSONRequestBody HostsPostRoomsJSONBody

// HostsPutRoomsRequestBody defines body for HostsPutRooms for application/json ContentType.
type HostsPutRoomsJSONRequestBody HostsPutRoomsJSONBody

