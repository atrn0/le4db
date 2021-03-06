package handler

import (
	"github.com/jmoiron/sqlx"

	"github.com/atrn0/le4db/handler/guests"
	"github.com/atrn0/le4db/handler/hosts"
)

type Impl struct {
	guests.GuestsRoomsHandler
	guests.GuestsReservationsHandler
	guests.GuestsUsersHandler
	hosts.HostsRoomsHandler
	hosts.HostsReservationsHandler
	hosts.HostsUsersHandler
}

func NewHandler(db *sqlx.DB) *Impl {
	return &Impl{
		guests.NewRoomsHandler(db),
		guests.NewReservationsHandler(db),
		guests.NewUsersHandler(db),
		hosts.NewRoomsHandler(db),
		hosts.NewReservationsHandler(db),
		hosts.NewUsersHandler(db),
	}
}
