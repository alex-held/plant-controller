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
type TrayConfigController struct {
	ctx      context.Context
	services *service.Manager
	logger   *logger.Logger
}

// NewTrays creates a new TrayController.
func NewTrayConfig(ctx context.Context, services *service.Manager, logger *logger.Logger) *TrayConfigController {
	return &TrayConfigController{
		ctx:      ctx,
		services: services,
		logger:   logger,
	}
}

// Create creates new tray
func (ctr *TrayConfigController) Create(ctx echo.Context) error {
	trayConfig := &model.TrayConfig{}
	err := ctx.Bind(trayConfig)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, errors.Wrap(err, "could not decode trayConfig data"))
	}
	err = ctx.Validate(trayConfig)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err)
	}
	
	createdTrayConfig, err := ctr.services.TrayConfig.Create(ctx.Request().Context(), trayConfig)
	if err != nil {
		switch {
		case errors.Cause(err) == types.ErrBadRequest:
			return echo.NewHTTPError(http.StatusBadRequest, err)
		default:
			return echo.NewHTTPError(http.StatusInternalServerError, errors.Wrap(err, "could not create trayConfig"))
		}
	}
	
	ctr.logger.Debug().Msgf("Created trayConfig %+v", *createdTrayConfig)
	
	return ctx.JSON(http.StatusCreated, createdTrayConfig)
}

func (ctr *TrayConfigController) Get(ctx echo.Context) error {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, errors.Wrap(err, "could not parse slot from url"))
	}
	user, err := ctr.services.TrayConfig.Get(ctx.Request().Context(), id)
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

func (ctr *TrayConfigController) GetAll(ctx echo.Context) error {
	
	user, err := ctr.services.TrayConfig.GetAll(ctx.Request().Context())
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
