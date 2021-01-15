package controller

import (
	"context"
	"net/http"
	"strconv"
	
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	
	"plant-controller/internal/logger"
	"plant-controller/internal/model"
	"plant-controller/internal/service"
	"plant-controller/internal/types"
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
	
	createdTray, err := ctr.services.Tray.Create(ctx.Request().Context(), tray)
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
	user, err := ctr.services.Tray.Get(ctx.Request().Context(), slot)
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

	user, err := ctr.services.Tray.GetAll(ctx.Request().Context())
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



func (ctr *TrayController) Delete(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, errors.Wrap(err, "could not parse slot from url"))
	}
	
	err = ctr.services.Tray.Delete(c.Request().Context(), id)
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
	
	return nil
}

func (ctr *TrayController) Update(ctx echo.Context) error {
	plant := &model.Plant{}
	err := ctx.Bind(plant)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, errors.Wrap(err, "could not decode plant data"))
	}
	err = ctx.Validate(plant)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err)
	}
	
	updatedPlant, err := ctr.services.Plant.Update(ctx.Request().Context(), plant)
	if err != nil {
		switch {
		case errors.Cause(err) == types.ErrBadRequest:
			return echo.NewHTTPError(http.StatusBadRequest, err)
		default:
			return echo.NewHTTPError(http.StatusInternalServerError, errors.Wrap(err, "could not update plant"))
		}
	}
	
	ctr.logger.Debug().Msgf("Created plant %+v", *updatedPlant)
	
	return ctx.JSON(http.StatusAccepted, updatedPlant)
}
