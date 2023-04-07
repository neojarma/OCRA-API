package comment_controller

import (
	"net/http"
	"ocra_server/model/entity"
	"ocra_server/model/request"
	"ocra_server/model/response"
	comment_service "ocra_server/service/comment"
	"sync"

	"github.com/labstack/echo/v4"
)

type CommentControllerImpl struct {
	CommentService comment_service.CommentService
}

func NewCommentController(service comment_service.CommentService) CommentController {
	var doOnce sync.Once
	controller := new(CommentControllerImpl)

	doOnce.Do(func() {
		controller = &CommentControllerImpl{
			CommentService: service,
		}
	})

	return controller
}

func (controller *CommentControllerImpl) GetComment(ctx echo.Context) error {
	req := new(request.CommentRequest)
	if err := ctx.Bind(req); err != nil {
		return ctx.JSON(http.StatusBadRequest, response.EmptyObjectDataResponse{
			Status:  response.StatusFailed,
			Message: err.Error(),
		})
	}

	result, err := controller.CommentService.GetVideoComments(req)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, response.EmptyObjectDataResponse{
			Status:  response.StatusFailed,
			Message: err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, response.Response{
		Status:  response.StatusSuccess,
		Message: response.MessageSuccesGetAllComment,
		Data:    result,
	})
}

func (controller *CommentControllerImpl) CreateComment(ctx echo.Context) error {
	req := new(entity.Comments)
	if err := ctx.Bind(req); err != nil {
		return ctx.JSON(http.StatusBadRequest, response.EmptyObjectDataResponse{
			Status:  response.StatusFailed,
			Message: err.Error(),
		})
	}

	if err := controller.CommentService.CreateComment(req); err != nil {
		return ctx.JSON(http.StatusBadRequest, response.EmptyObjectDataResponse{
			Status:  response.StatusFailed,
			Message: err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, response.EmptyObjectDataResponse{
		Status:  response.StatusSuccess,
		Message: response.MessageSuccessCreateComment,
	})
}

func (controller *CommentControllerImpl) UpdateComment(ctx echo.Context) error {
	req := new(entity.Comments)
	if err := ctx.Bind(req); err != nil {
		return ctx.JSON(http.StatusBadRequest, response.EmptyObjectDataResponse{
			Status:  response.StatusFailed,
			Message: err.Error(),
		})
	}

	if err := controller.CommentService.UpdateComment(req); err != nil {
		return ctx.JSON(http.StatusBadRequest, response.EmptyObjectDataResponse{
			Status:  response.StatusFailed,
			Message: err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, response.EmptyObjectDataResponse{
		Status:  response.StatusSuccess,
		Message: response.MessageSuccessUpdateComment,
	})
}

func (controller *CommentControllerImpl) DeleteComment(ctx echo.Context) error {
	req := new(entity.Comments)
	if err := ctx.Bind(req); err != nil {
		return ctx.JSON(http.StatusBadRequest, response.EmptyObjectDataResponse{
			Status:  response.StatusFailed,
			Message: err.Error(),
		})
	}

	if err := controller.CommentService.DeleteComment(req); err != nil {
		return ctx.JSON(http.StatusBadRequest, response.EmptyObjectDataResponse{
			Status:  response.StatusFailed,
			Message: err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, response.EmptyObjectDataResponse{
		Status:  response.StatusSuccess,
		Message: response.MessageSuccessDeleteComment,
	})
}
