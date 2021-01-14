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
type GrowConfigController struct {
	ctx      context.Context
	services *service.Manager
	logger   *logger.Logger
}

// NewTrays creates a new TrayController.
func NewGrowConfig(ctx context.Context, services *service.Manager, logger *logger.Logger) *GrowConfigController {
	return &GrowConfigController{
		ctx:      ctx,
		services: services,
		logger:   logger,
	}
}

// Create creates new tray
func (ctr *GrowConfigController) Create(ctx echo.Context) error {
	growConfig := &model.GrowConfig{}
	err := ctx.Bind(growConfig)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, errors.Wrap(err, "could not decode growConfig data"))
	}
	err = ctx.Validate(growConfig)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err)
	}
	
	createdGrowConfig, err := ctr.services.GrowConfig.Create(ctx.Request().Context(), growConfig)
	if err != nil {
		switch {
		case errors.Cause(err) == types.ErrBadRequest:
			return echo.NewHTTPError(http.StatusBadRequest, err)
		default:
			return echo.NewHTTPError(http.StatusInternalServerError, errors.Wrap(err, "could not create growConfig"))
		}
	}
	
	ctr.logger.Debug().Msgf("Created growConfig %+v", *createdGrowConfig)
	
	return ctx.JSON(http.StatusCreated, createdGrowConfig)
}

func (ctr *GrowConfigController) Get(ctx echo.Context) error {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, errors.Wrap(err, "could not parse slot from url"))
	}
	user, err := ctr.services.GrowConfig.Get(ctx.Request().Context(), id)
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

func (ctr *GrowConfigController) GetAll(ctx echo.Context) error {
	
	user, err := ctr.services.GrowConfig.GetAll(ctx.Request().Context())
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

func (ctr *GrowConfigController) Delete(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, errors.Wrap(err, "could not parse slot from url"))
	}
	
	err = ctr.services.GrowConfig.Delete(c.Request().Context(), id)
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

func (ctr *GrowConfigController) Update(ctx echo.Context) error {
	growConfig := &model.GrowConfig{}
	err := ctx.Bind(growConfig)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, errors.Wrap(err, "could not decode growconfig data"))
	}
	err = ctx.Validate(growConfig)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err)
	}
	
	updated, err := ctr.services.GrowConfig.Update(ctx.Request().Context(), growConfig)
	if err != nil {
		switch {
		case errors.Cause(err) == types.ErrBadRequest:
			return echo.NewHTTPError(http.StatusBadRequest, err)
		default:
			return echo.NewHTTPError(http.StatusInternalServerError, errors.Wrap(err, "could not update plant"))
		}
	}
	
	ctr.logger.Debug().Msgf("Created plant %+v", *updated)
	
	return ctx.JSON(http.StatusAccepted, updated)
}
