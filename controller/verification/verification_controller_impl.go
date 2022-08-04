package verification_controller

import (
	"net/http"
	"ocra_server/model/response"
	verification_service "ocra_server/service/verification"
	"sync"

	"github.com/labstack/echo/v4"
)

type VerificationControllerImpl struct {
	Service verification_service.VerificationService
}

func NewVerificationController(service verification_service.VerificationService) VerificationController {
	var doOnce sync.Once
	controller := new(VerificationControllerImpl)

	doOnce.Do(func() {
		controller = &VerificationControllerImpl{
			Service: service,
		}
	})

	return controller
}

func (controller *VerificationControllerImpl) CreateVerificationToken(ctx echo.Context) error {
	email := ctx.QueryParam("email")

	if err := controller.Service.CreateVerifToken(email); err != nil {
		return ctx.JSON(http.StatusBadRequest, &response.EmptyObjectDataResponse{
			Status:  response.StatusFailed,
			Message: response.MessageFailedSendingNewToken,
		})
	}

	return ctx.JSON(http.StatusOK, &response.EmptyObjectDataResponse{
		Status:  response.StatusSuccess,
		Message: response.MessageSuccessSendingNewToken,
	})
}

func (controller *VerificationControllerImpl) ValidateVerificationToken(ctx echo.Context) error {
	token := ctx.QueryParam("token")
	email := ctx.QueryParam("email")

	if err := controller.Service.ValidateVerifToken(email, token); err != nil {
		return ctx.JSON(http.StatusBadRequest, &response.EmptyObjectDataResponse{
			Status:  response.StatusFailed,
			Message: err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, &response.EmptyObjectDataResponse{
		Status:  response.StatusSuccess,
		Message: response.MessageSuccessVerifyEmail,
	})
}
