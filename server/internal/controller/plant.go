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
type PlantController struct {
	ctx      context.Context
	services *service.Manager
	logger   *logger.Logger
}

// NewTrays creates a new TrayController.
func NewPlant(ctx context.Context, services *service.Manager, logger *logger.Logger) *PlantController {
	return &PlantController{
		ctx:      ctx,
		services: services,
		logger:   logger,
	}
}

// Create creates new tray
func (ctr *PlantController) Create(ctx echo.Context) error {
	plant := &model.Plant{}
	err := ctx.Bind(plant)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, errors.Wrap(err, "could not decode plant data"))
	}
	err = ctx.Validate(plant)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err)
	}
	
	createdPlant, err := ctr.services.Plant.Create(ctx.Request().Context(), plant.Name)
	if err != nil {
		switch {
		case errors.Cause(err) == types.ErrBadRequest:
			return echo.NewHTTPError(http.StatusBadRequest, err)
		default:
			return echo.NewHTTPError(http.StatusInternalServerError, errors.Wrap(err, "could not create plant"))
		}
	}
	
	ctr.logger.Debug().Msgf("Created plant %+v", *createdPlant)
	
	return ctx.JSON(http.StatusCreated, createdPlant)
}

func (ctr *PlantController) Get(ctx echo.Context) error {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, errors.Wrap(err, "could not parse slot from url"))
	}
	user, err := ctr.services.Plant.Get(ctx.Request().Context(), id)
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

func (ctr *PlantController) GetAll(ctx echo.Context) error {
	
	user, err := ctr.services.Plant.GetAll(ctx.Request().Context())
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

func (ctr *PlantController) Delete(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, errors.Wrap(err, "could not parse slot from url"))
	}
	
	err = ctr.services.Plant.Delete(c.Request().Context(), id)
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

func (ctr *PlantController) Update(ctx echo.Context) error {
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
