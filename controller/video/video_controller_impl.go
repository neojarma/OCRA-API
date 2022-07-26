package video_controller

import (
	"net/http"
	"ocra_server/model/entity"
	"ocra_server/model/request"
	"ocra_server/model/response"
	videos_service "ocra_server/service/video"
	"strconv"
	"sync"

	"github.com/labstack/echo/v4"
)

type VideoControllerImpl struct {
	Service videos_service.VideoService
}

func NewVideoController(service videos_service.VideoService) VideoController {
	var doOnce sync.Once
	controller := new(VideoControllerImpl)

	doOnce.Do(func() {
		controller = &VideoControllerImpl{
			Service: service,
		}
	})

	return controller
}

func (controller *VideoControllerImpl) GetAllVideos(ctx echo.Context) error {
	req := new(request.AllVideosRequest)
	if err := ctx.Bind(req); err != nil {
		return ctx.JSON(http.StatusBadRequest, response.EmptyArrayDataResponse{
			Status:  response.StatusFailed,
			Message: err.Error(),
		})
	}

	ok, err := strconv.ParseBool(req.IsSubscribed)
	if err == nil && ok {
		userIdFromCookie := ctx.Request().Header.Get("user-id")
		if userIdFromCookie == "" {
			return ctx.JSON(http.StatusBadRequest, response.EmptyArrayDataResponse{
				Status:  response.StatusFailed,
				Message: response.MessageMissingSessionId,
			})
		}

		req.UserId = userIdFromCookie
	}

	result, err := controller.Service.GetAllVideos(req)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, response.EmptyArrayDataResponse{
			Status:  response.StatusFailed,
			Message: err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, response.Response{
		Status:  response.StatusSuccess,
		Message: response.MessageSuccesGetAllVideos,
		Data:    result,
	})
}

func (controller *VideoControllerImpl) GetDetailVideos(ctx echo.Context) error {
	req := new(request.DetailVideoRequest)
	if err := ctx.Bind(req); err != nil {
		return ctx.JSON(http.StatusBadRequest, response.EmptyObjectDataResponse{
			Status:  response.StatusFailed,
			Message: err.Error(),
		})
	}

	result, err := controller.Service.GetDetailVideos(req)
	if err != nil {
		if err.Error() == response.MessageNoVideo {
			return ctx.JSON(http.StatusNotFound, response.EmptyObjectDataResponse{
				Status:  response.StatusFailed,
				Message: err.Error(),
			})
		}

		return ctx.JSON(http.StatusBadRequest, response.EmptyObjectDataResponse{
			Status:  response.StatusFailed,
			Message: err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, response.Response{
		Status:  response.StatusSuccess,
		Message: response.MessageSuccesGetVideo,
		Data:    result,
	})
}

func (controller *VideoControllerImpl) CreateVideo(ctx echo.Context) error {
	req := new(entity.Video)
	err := ctx.Bind(req)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, response.EmptyObjectDataResponse{
			Status:  response.StatusFailed,
			Message: err.Error(),
		})
	}

	thumbnail, err := ctx.FormFile("thumbnail")
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, response.EmptyObjectDataResponse{
			Status:  response.StatusFailed,
			Message: response.MessageMissingThumbnail,
		})
	}

	video, err := ctx.FormFile("video")
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, response.EmptyObjectDataResponse{
			Status:  response.StatusFailed,
			Message: response.MessageMissingVideo,
		})
	}

	res, err := controller.Service.CreateVideo(req, thumbnail, video)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, response.EmptyObjectDataResponse{
			Status:  response.StatusFailed,
			Message: err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, response.Response{
		Status:  response.StatusSuccess,
		Message: response.MessageSuccesUploadVideo,
		Data:    res,
	})
}

func (controller *VideoControllerImpl) UpdateVideo(ctx echo.Context) error {
	return nil
}

func (controller *VideoControllerImpl) IncrementViews(ctx echo.Context) error {
	videoId := ctx.Param("id")

	if videoId == "" {
		return ctx.JSON(http.StatusBadRequest, response.EmptyObjectDataResponse{
			Status:  response.StatusFailed,
			Message: response.MessageMissingVideoId,
		})
	}

	if err := controller.Service.IncrementViewsVideo(videoId); err != nil {
		ctx.JSON(http.StatusInternalServerError, response.EmptyObjectDataResponse{
			Status:  response.StatusFailed,
			Message: err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, response.EmptyObjectDataResponse{
		Status:  response.StatusSuccess,
		Message: response.MessageSuccessUpdateViews,
	})
}

func (controller *VideoControllerImpl) Find(ctx echo.Context) error {
	query := ctx.QueryParam("query")

	if query == "" {
		return ctx.JSON(http.StatusBadRequest, response.EmptyObjectDataResponse{
			Status:  response.StatusFailed,
			Message: response.MessageMissingVideoQuery,
		})
	}

	result, err := controller.Service.Find(query)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, response.EmptyArrayDataResponse{
			Status:  response.StatusFailed,
			Message: err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, response.Response{
		Status:  response.StatusSuccess,
		Message: response.MessageSuccesGetAllVideos,
		Data:    result,
	})
}

func (controller *VideoControllerImpl) AutoComplete(ctx echo.Context) error {
	query := ctx.QueryParam("query")

	if query == "" {
		return ctx.JSON(http.StatusBadRequest, response.EmptyObjectDataResponse{
			Status:  response.StatusFailed,
			Message: response.MessageMissingVideoQuery,
		})
	}

	result, err := controller.Service.AutoComplete(query)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, response.EmptyArrayDataResponse{
			Status:  response.StatusFailed,
			Message: err.Error(),
		})
	}

	if len(result) == 0 {
		return ctx.JSON(http.StatusOK, response.Response{
			Status:  response.StatusSuccess,
			Message: response.MessageNoSuggestions,
			Data:    result,
		})
	}

	return ctx.JSON(http.StatusOK, response.Response{
		Status:  response.StatusSuccess,
		Message: response.MessageSuccesGetSuggestions,
		Data:    result,
	})
}
