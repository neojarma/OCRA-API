package choice_controller

import (
	"net/http"
	"ocra_server/model/entity"
	"ocra_server/model/response"
	choice_service "ocra_server/service/choice"
	"sync"

	"github.com/labstack/echo/v4"
)

type ChoiceControllerImpl struct {
	Service choice_service.ChoiceService
}

func NewChoiceController(service choice_service.ChoiceService) ChoiceController {
	var doOnce sync.Once
	controller := new(ChoiceControllerImpl)

	doOnce.Do(func() {
		controller = &ChoiceControllerImpl{
			Service: service,
		}
	})

	return controller
}

func (controller *ChoiceControllerImpl) CreateLike(ctx echo.Context) error {
	req := new(entity.Likes)

	if err := ctx.Bind(req); err != nil {
		return ctx.JSON(http.StatusBadRequest, response.EmptyObjectDataResponse{
			Status:  response.StatusFailed,
			Message: err.Error(),
		})
	}

	if err := controller.Service.CreateLikeRecord(req); err != nil {
		return ctx.JSON(http.StatusBadRequest, response.EmptyObjectDataResponse{
			Status:  response.StatusFailed,
			Message: err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, response.EmptyObjectDataResponse{
		Status:  response.StatusSuccess,
		Message: response.MessageSuccesLikeVideo,
	})
}

func (controller *ChoiceControllerImpl) CreateDislike(ctx echo.Context) error {
	req := new(entity.Dislikes)

	if err := ctx.Bind(req); err != nil {
		return ctx.JSON(http.StatusBadRequest, response.EmptyObjectDataResponse{
			Status:  response.StatusFailed,
			Message: err.Error(),
		})
	}

	if err := controller.Service.CreateDislikeRecord(req); err != nil {
		return ctx.JSON(http.StatusBadRequest, response.EmptyObjectDataResponse{
			Status:  response.StatusFailed,
			Message: err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, response.EmptyObjectDataResponse{
		Status:  response.StatusSuccess,
		Message: response.MessageSuccessDislikeVideo,
	})
}

func (controller *ChoiceControllerImpl) DeleteLike(ctx echo.Context) error {
	req := new(entity.Likes)

	if err := ctx.Bind(req); err != nil {
		return ctx.JSON(http.StatusBadRequest, response.EmptyObjectDataResponse{
			Status:  response.StatusFailed,
			Message: err.Error(),
		})
	}

	if err := controller.Service.DeleteLike(req); err != nil {
		return ctx.JSON(http.StatusBadRequest, response.EmptyObjectDataResponse{
			Status:  response.StatusFailed,
			Message: err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, response.EmptyObjectDataResponse{
		Status:  response.StatusSuccess,
		Message: response.MessageSuccessDeleteLike,
	})
}

func (controller *ChoiceControllerImpl) DeleteDislike(ctx echo.Context) error {
	req := new(entity.Dislikes)

	if err := ctx.Bind(req); err != nil {
		return ctx.JSON(http.StatusBadRequest, response.EmptyObjectDataResponse{
			Status:  response.StatusFailed,
			Message: err.Error(),
		})
	}

	if err := controller.Service.DeleteDislike(req); err != nil {
		return ctx.JSON(http.StatusBadRequest, response.EmptyObjectDataResponse{
			Status:  response.StatusFailed,
			Message: err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, response.EmptyObjectDataResponse{
		Status:  response.StatusSuccess,
		Message: response.MessageSuccessDeleteDislike,
	})
}
