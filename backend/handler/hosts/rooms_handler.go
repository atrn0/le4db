package hosts

import (
	"net/http"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/rs/xid"

	"github.com/atrn0/le4db/entity"
	oapi "github.com/atrn0/le4db/gen/openapi"
	"github.com/atrn0/le4db/handler/errors"
	"github.com/atrn0/le4db/middleware/auth"
)

type RoomsHandlerImpl struct {
	db *sqlx.DB
}

type HostsRoomsHandler interface {
	HostsGetRooms(ctx echo.Context) error
	HostsPostRooms(ctx echo.Context) error
}

func NewRoomsHandler(db *sqlx.DB) HostsRoomsHandler {
	return &RoomsHandlerImpl{db}
}

func (h *RoomsHandlerImpl) HostsGetRooms(ctx echo.Context) error {
	userId := auth.GetUserId(ctx)
	if userId == "" {
		return ctx.JSON(
			http.StatusBadRequest,
			errors.ErrorRes{Message: "userId is required"},
		)
	}

	var rooms []entity.Room
	err := h.db.Select(
		&rooms,
		"select * from rooms where host_id = $1",
		userId,
	)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError,
			errors.ErrorRes{Message: err.Error()})
	}

	var roomsRes []oapi.HostsRoom
	for _, room := range rooms {
		roomsRes = append(roomsRes, oapi.HostsRoom{
			Id:    room.ID,
			Name:  room.Name,
			Price: room.Price,
		})
	}

	return ctx.JSON(http.StatusOK, oapi.HostsGetRoomsRes{
		Rooms: roomsRes,
	})
}

func (h *RoomsHandlerImpl) HostsPostRooms(ctx echo.Context) error {
	userId := auth.GetUserId(ctx)
	if userId == "" {
		return ctx.JSON(
			http.StatusBadRequest,
			errors.ErrorRes{Message: "userId is required"},
		)
	}

	req := new(oapi.HostsPostRoomsReq)
	if err := ctx.Bind(req); err != nil {
		return ctx.JSON(
			http.StatusBadRequest,
			errors.ErrorRes{Message: err.Error()},
		)
	}

	roomCreate := entity.Room{
		ID:     xid.New().String(),
		Name:   req.Name,
		Price:  req.Price,
		HostID: userId,
	}

	_, err := h.db.NamedExec(`
	INSERT INTO rooms (id, name, price, host_id) 
	VALUES (:id, :name, :price, :host_id) 
	`, roomCreate)

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError,
			errors.ErrorRes{Message: err.Error()})
	}

	return ctx.NoContent(http.StatusNoContent)
}
