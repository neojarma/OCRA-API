package user_controller

import (
	"net/http"
	"ocra_server/helper"
	"ocra_server/model/request"
	"ocra_server/model/response"
	cookie_service "ocra_server/service/cookie"
	user_service "ocra_server/service/user"
	"sync"
	"time"

	"github.com/labstack/echo/v4"
)

type UserControllerImpl struct {
	Service       user_service.UserService
	CookieService cookie_service.CookieService
}

func NewUserController(service user_service.UserService, cookieService cookie_service.CookieService) UserController {
	var doOnce sync.Once
	controller := new(UserControllerImpl)

	doOnce.Do(func() {
		controller = &UserControllerImpl{
			Service:       service,
			CookieService: cookieService,
		}
	})

	return controller
}

func (controller *UserControllerImpl) Login(ctx echo.Context) error {
	req := new(request.AuthRequest)
	if err := ctx.Bind(req); err != nil {
		return ctx.JSON(http.StatusBadRequest, &response.EmptyObjectDataResponse{
			Status:  response.StatusFailed,
			Message: response.MessageErrorBindingData,
		})
	}

	if err := helper.ValidateUserInput(req); err != nil {
		return ctx.JSON(http.StatusBadRequest, &response.EmptyObjectDataResponse{
			Status:  response.StatusFailed,
			Message: err.Error(),
		})
	}

	result, err := controller.Service.ValidateLogin(req)
	if err != nil {
		return ctx.JSON(http.StatusUnauthorized, &response.EmptyObjectDataResponse{
			Status:  response.StatusFailed,
			Message: err.Error(),
		})
	}

	cookieService := cookie_service.NewCookieService().CreateCookie("session-id", *result.SessionId, time.Now().Add(72*time.Hour))
	ctx.SetCookie(cookieService)

	// unset session-id in json
	result.SessionId = nil

	return ctx.JSON(http.StatusOK, &response.Response{
		Status:  response.StatusSuccess,
		Message: response.MessageSuccessLogin,
		Data:    result,
	})
}

func (controller *UserControllerImpl) Logout(ctx echo.Context) error {
	// get session-id in cookie
	cookie, err := ctx.Cookie("session-id")
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, &response.EmptyObjectDataResponse{
			Status:  response.StatusFailed,
			Message: err.Error(),
		})
	}

	sessionId := cookie.Value
	if err := controller.Service.Logout(sessionId); err != nil {
		return ctx.JSON(http.StatusBadRequest, &response.EmptyObjectDataResponse{
			Status:  response.StatusFailed,
			Message: err.Error(),
		})
	}

	deletedCookie := controller.CookieService.DestroyCookie(sessionId)
	ctx.SetCookie(deletedCookie)

	return ctx.JSON(http.StatusOK, &response.EmptyObjectDataResponse{
		Status:  response.StatusSuccess,
		Message: response.MessageSuccessLogout,
	})
}

func (controller *UserControllerImpl) Register(ctx echo.Context) error {
	req := new(request.UserRequest)
	if err := ctx.Bind(req); err != nil {
		return ctx.JSON(http.StatusBadRequest, &response.EmptyObjectDataResponse{
			Status:  response.StatusFailed,
			Message: response.MessageErrorBindingData,
		})
	}

	if err := helper.ValidateUserInput(req); err != nil {
		return ctx.JSON(http.StatusBadRequest, &response.EmptyObjectDataResponse{
			Status:  response.StatusFailed,
			Message: err.Error(),
		})
	}

	result, err := controller.Service.Register(req)
	if err != nil {
		if err.Error() == response.MessageFailedRegisterEmailExist {
			return ctx.JSON(http.StatusConflict, &response.EmptyObjectDataResponse{
				Status:  response.StatusFailed,
				Message: err.Error(),
			})
		}

		return ctx.JSON(http.StatusBadRequest, &response.EmptyObjectDataResponse{
			Status:  response.StatusFailed,
			Message: err.Error(),
		})
	}

	return ctx.JSON(http.StatusCreated, &response.Response{
		Status:  response.StatusSuccess,
		Message: response.MessageSuccessRegister,
		Data:    result,
	})

}

func (controller *UserControllerImpl) UpdateUser(ctx echo.Context) error {
	req := new(request.UserRequest)
	ctx.Bind(req)
	if _, err := controller.Service.UpdateUser(req); err != nil {
		return ctx.String(http.StatusInternalServerError, "false")
	}

	return ctx.String(http.StatusOK, "ok")
}
