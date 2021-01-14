package controller

import (
	"context"
	"net/http"
	"strconv"
	"time"
	
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	
	"plant-controller/lib/types"
	"plant-controller/logger"
	"plant-controller/model"
	"plant-controller/service"
)

// TrayController ...
type TrayController struct {
	ctx      context.Context
	services *service.Manager
	logger   *logger.Logger
}

// NewTrays creates a new TrayController.
func NewTrays(ctx context.Context, services *service.Manager, logger *logger.Logger) *TrayController {
	return &TrayController{
		ctx:      ctx,
		services: services,
		logger:   logger,
	}
}

// Create creates new tray
func (ctr *TrayController) Create(ctx echo.Context) error {
	tray := &model.Tray{}
	err := ctx.Bind(tray)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, errors.Wrap(err, "could not decode tray data"))
	}
	err = ctx.Validate(tray)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err)
	}
	tray.StartDate = time.Now()
	createdTray, err := ctr.services.Tray.CreateTray(ctx.Request().Context(), tray)
	if err != nil {
		switch {
		case errors.Cause(err) == types.ErrBadRequest:
			return echo.NewHTTPError(http.StatusBadRequest, err)
		default:
			return echo.NewHTTPError(http.StatusInternalServerError, errors.Wrap(err, "could not create tray"))
		}
	}
	
	ctr.logger.Debug().Msgf("Created tray '%d'", createdTray.Slot)
	
	return ctx.JSON(http.StatusCreated, createdTray)
}

func (ctr *TrayController) Get(ctx echo.Context) error {
	slot, err := strconv.Atoi(ctx.Param("slot"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, errors.Wrap(err, "could not parse slot from url"))
	}
	user, err := ctr.services.Tray.GetTray(ctx.Request().Context(), slot)
	if err != nil {
		switch {
		case errors.Cause(err) == types.ErrNotFound:
			return echo.NewHTTPError(http.StatusNotFound, err)
		case errors.Cause(err) == types.ErrBadRequest:
			return echo.NewHTTPError(http.StatusBadRequest, err)
		default:
			return echo.NewHTTPError(http.StatusInternalServerError, errors.Wrap(err, "could not get slot"))
		}
	}
	return ctx.JSON(http.StatusOK, user)
}

func (ctr *TrayController) GetAll(ctx echo.Context) error {

	user, err := ctr.services.Tray.GetAllTrays(ctx.Request().Context())
	if err != nil {
		switch {
		case errors.Cause(err) == types.ErrNotFound:
			return echo.NewHTTPError(http.StatusNotFound, err)
		case errors.Cause(err) == types.ErrBadRequest:
			return echo.NewHTTPError(http.StatusBadRequest, err)
		default:
			return echo.NewHTTPError(http.StatusInternalServerError, errors.Wrap(err, "could not get slots"))
		}
	}
	return ctx.JSON(http.StatusOK, user)
}
