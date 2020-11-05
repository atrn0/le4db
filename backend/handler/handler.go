package handler

import (
	"github.com/atrn0/le4db/handler/guests"
	"github.com/atrn0/le4db/handler/hosts"
)

type Impl struct {
	guests.GuestsRoomsHandler
	guests.GuestsReservationsHandler
	guests.GuestsUsersHandler
	hosts.HostsRoomsHandler
	hosts.HostsReservationsHandler
}

func NewHandler() *Impl {
	return &Impl{
		guests.NewRoomsHandler(),
		guests.NewReservationsHandler(),
		guests.NewUsersHandler(),
		hosts.NewRoomsHandler(),
		hosts.NewReservationsHandler(),
	}
}