package test

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"ocra_server/model/entity"
	joins_model "ocra_server/model/joins"
	"ocra_server/model/request"
	"ocra_server/model/response"
	"ocra_server/router"
	"testing"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/patrickmn/go-cache"
	"github.com/stretchr/testify/assert"
)

func TestRegister(t *testing.T) {
	setup := Setup()

	bodyRegister := []struct {
		CaseDescription         string
		ExpectedHttpRespCode    int
		ExpectedHttpRespStatus  string
		ExpectedHttpRespMessage string
		ExpectedHttpRespData    *response.UserResponse
		Request                 any
	}{
		{
			CaseDescription:         "reject empt body request",
			ExpectedHttpRespCode:    400,
			ExpectedHttpRespStatus:  response.StatusFailed,
			ExpectedHttpRespMessage: response.MessageInvalidJsonInput,
			ExpectedHttpRespData:    &response.UserResponse{},
			Request:                 &request.UserRequest{},
		},
		{
			CaseDescription:         "invalid data type",
			ExpectedHttpRespCode:    400,
			ExpectedHttpRespStatus:  response.StatusFailed,
			ExpectedHttpRespMessage: response.MessageErrorBindingData,
			ExpectedHttpRespData:    &response.UserResponse{},
			Request: struct {
				FullName bool `json:"fullName"`
			}{
				FullName: false,
			},
		},
		{
			CaseDescription:         "success register",
			ExpectedHttpRespCode:    201,
			ExpectedHttpRespStatus:  response.StatusSuccess,
			ExpectedHttpRespMessage: response.MessageSuccessRegister,
			ExpectedHttpRespData: &response.UserResponse{
				FullName: "neo jarmawijaya",
				Email:    "neojarma@gmail.com",
			},
			Request: &request.UserRequest{
				FullName: "neo jarmawijaya",
				Email:    "neojarma@gmail.com",
				Password: "73*3t2YN4rbE",
			},
		},
		{
			CaseDescription:         "reject missing password property",
			ExpectedHttpRespCode:    400,
			ExpectedHttpRespStatus:  response.StatusFailed,
			ExpectedHttpRespMessage: response.MessageInvalidJsonInput,
			ExpectedHttpRespData:    &response.UserResponse{},
			Request: &request.UserRequest{
				FullName: "neo jarmawijaya",
				Email:    "neojarma@gmail.com",
			},
		},
		{
			CaseDescription:         "reject missing email property",
			ExpectedHttpRespCode:    400,
			ExpectedHttpRespStatus:  response.StatusFailed,
			ExpectedHttpRespMessage: response.MessageInvalidJsonInput,
			ExpectedHttpRespData:    &response.UserResponse{},
			Request: &request.UserRequest{
				FullName: "neo jarmawijaya",
				Password: "73*3t2YN4rbE",
			},
		},
		{
			CaseDescription:         "reject missing full name property",
			ExpectedHttpRespCode:    400,
			ExpectedHttpRespStatus:  response.StatusFailed,
			ExpectedHttpRespMessage: response.MessageInvalidJsonInput,
			ExpectedHttpRespData:    &response.UserResponse{},
			Request: &request.UserRequest{
				Password: "73*3t2YN4rbE",
				Email:    "neojarma@gmail.com",
			},
		},
		{
			CaseDescription:         "reject invalid email",
			ExpectedHttpRespCode:    400,
			ExpectedHttpRespStatus:  response.StatusFailed,
			ExpectedHttpRespMessage: response.MessageInvalidJsonInput,
			ExpectedHttpRespData:    &response.UserResponse{},
			Request: &request.UserRequest{
				FullName: "neo jarmawijaya",
				Email:    "invalid mail",
				Password: "73*3t2YN4rbE",
			},
		},
		{
			CaseDescription:         "reject weak password",
			ExpectedHttpRespCode:    400,
			ExpectedHttpRespStatus:  response.StatusFailed,
			ExpectedHttpRespMessage: response.MessageInvalidJsonInput,
			ExpectedHttpRespData:    &response.UserResponse{},
			Request: &request.UserRequest{
				FullName: "neo jarmawijaya",
				Email:    "neojarma@gmail.com",
				Password: "weakpass",
			},
		},
		{
			CaseDescription:         "reject registered email",
			ExpectedHttpRespCode:    409,
			ExpectedHttpRespStatus:  response.StatusFailed,
			ExpectedHttpRespMessage: response.MessageFailedRegisterEmailExist,
			ExpectedHttpRespData:    &response.UserResponse{},
			Request: &request.UserRequest{
				FullName: "neo jarmawijaya",
				Email:    "neojarma@gmail.com",
				Password: "73*3t2YN4rbE",
			},
		},
	}

	for _, v := range bodyRegister {
		jsonByte, err := json.Marshal(v.Request)
		if err != nil {
			panic(err)
		}

		body := bytes.NewReader(jsonByte)
		req := httptest.NewRequest(http.MethodPost, "/register", body)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		rec := httptest.NewRecorder()

		c := setup.App.NewContext(req, rec)
		h := router.UserRoute(setup.Group, setup.Db, setup.Dialer, cache.New(time.Hour*72, time.Hour*120))

		h.Register(c)

		bodyRes, err := io.ReadAll(rec.Body)
		if err != nil {
			panic(err)
		}

		type Response struct {
			Status  string
			Message string
			Data    *response.UserResponse
		}

		bodyBinding := new(Response)
		err = json.Unmarshal(bodyRes, bodyBinding)
		if err != nil {
			panic(err)
		}

		response := rec.Result()
		t.Run(v.CaseDescription, func(t *testing.T) {
			assert.Equal(t, v.ExpectedHttpRespCode, response.StatusCode)
			assert.Equal(t, v.ExpectedHttpRespStatus, bodyBinding.Status)
			assert.Equal(t, v.ExpectedHttpRespMessage, bodyBinding.Message)
		})
	}

	setup.Db.Exec("DELETE FROM verifications")
	setup.Db.Exec("DELETE FROM users")

}

func TestLogin(t *testing.T) {
	setup := Setup()

	// insert 2 data
	setup.Db.Create(&entity.Users{
		UserId:     "example",
		FullName:   "test login",
		Email:      "testlogin2@gmail.com",
		IsVerified: true,
		Password:   "$2a$04$06cMVv05guWgqVlofh5llev.UMUTqFt.k6d5ldPOw7RARfRmOZTOa",
		CreatedAt:  1,
		UpdatedAt:  1,
	})
	setup.Db.Create(&entity.Users{
		UserId:     "example2",
		FullName:   "test login",
		Email:      "testlogin1@gmail.com",
		IsVerified: false,
		Password:   "$2a$04$06cMVv05guWgqVlofh5llev.UMUTqFt.k6d5ldPOw7RARfRmOZTOa",
		CreatedAt:  1,
		UpdatedAt:  1,
	})

	bodyLogin := []struct {
		CaseDescription         string
		ExpectedHttpRespCode    int
		ExpectedHttpRespStatus  string
		ExpectedHttpRespMessage string
		ExpectedHttpRespData    *joins_model.UserChannelJoin
		Request                 any
	}{
		{
			CaseDescription:         "success login",
			ExpectedHttpRespCode:    200,
			ExpectedHttpRespStatus:  response.StatusSuccess,
			ExpectedHttpRespMessage: response.MessageSuccessLogin,
			ExpectedHttpRespData: &joins_model.UserChannelJoin{
				Email: "testlogin2@gmail.com",
			},
			Request: &request.AuthRequest{
				Email:    "testlogin2@gmail.com",
				Password: "73*3t2YN4rbE",
			},
		},
		{
			CaseDescription:         "invalid data type",
			ExpectedHttpRespCode:    400,
			ExpectedHttpRespStatus:  response.StatusFailed,
			ExpectedHttpRespMessage: response.MessageErrorBindingData,
			ExpectedHttpRespData:    &joins_model.UserChannelJoin{},
			Request: struct {
				Email    bool `json:"email"`
				Password bool `json:"password"`
			}{
				Email:    true,
				Password: false,
			},
		},
		{
			CaseDescription:         "reject unverified user",
			ExpectedHttpRespCode:    401,
			ExpectedHttpRespStatus:  response.StatusFailed,
			ExpectedHttpRespMessage: response.MessageNotVerifedUser,
			ExpectedHttpRespData:    &joins_model.UserChannelJoin{},
			Request: &request.AuthRequest{
				Email:    "testlogin1@gmail.com",
				Password: "73*3t2YN4rbE",
			},
		},
		{
			CaseDescription:         "reject wrong password",
			ExpectedHttpRespCode:    http.StatusUnauthorized,
			ExpectedHttpRespStatus:  response.StatusFailed,
			ExpectedHttpRespMessage: "record not found",
			ExpectedHttpRespData:    &joins_model.UserChannelJoin{},
			Request: &request.AuthRequest{
				Email:    "notfound@gmail.com",
				Password: "73*3t2YN4rbEfalse",
			},
		},
		{
			CaseDescription:         "empty body value",
			ExpectedHttpRespCode:    http.StatusBadRequest,
			ExpectedHttpRespStatus:  response.StatusFailed,
			ExpectedHttpRespMessage: response.MessageInvalidJsonInput,
			ExpectedHttpRespData:    &joins_model.UserChannelJoin{},
			Request:                 &request.AuthRequest{},
		},
		{
			CaseDescription:         "missing email property",
			ExpectedHttpRespCode:    http.StatusBadRequest,
			ExpectedHttpRespStatus:  response.StatusFailed,
			ExpectedHttpRespMessage: response.MessageInvalidJsonInput,
			ExpectedHttpRespData:    &joins_model.UserChannelJoin{},
			Request: &request.AuthRequest{
				Password: "73*3t2YN4rbE",
			},
		},
		{
			CaseDescription:         "missing password property",
			ExpectedHttpRespCode:    http.StatusBadRequest,
			ExpectedHttpRespStatus:  response.StatusFailed,
			ExpectedHttpRespMessage: response.MessageInvalidJsonInput,
			ExpectedHttpRespData:    &joins_model.UserChannelJoin{},
			Request: &request.AuthRequest{
				Email: "testlogin2@gmail.com",
			},
		},
		{
			CaseDescription:         "email not found in db",
			ExpectedHttpRespCode:    http.StatusUnauthorized,
			ExpectedHttpRespStatus:  response.StatusFailed,
			ExpectedHttpRespMessage: "record not found",
			ExpectedHttpRespData:    &joins_model.UserChannelJoin{},
			Request: &request.AuthRequest{
				Email:    "notfound@gmail.com",
				Password: "73*3t2YN4rbE",
			},
		},
	}

	for _, v := range bodyLogin {
		jsonByte, err := json.Marshal(v.Request)
		if err != nil {
			panic(err)
		}

		body := bytes.NewReader(jsonByte)
		req := httptest.NewRequest(http.MethodPost, "/auth/login", body)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		rec := httptest.NewRecorder()

		c := setup.App.NewContext(req, rec)
		h := router.UserRoute(setup.Group, setup.Db, setup.Dialer, cache.New(time.Hour*72, time.Hour*120))

		h.Login(c)

		bodyRes, err := io.ReadAll(rec.Body)
		if err != nil {
			panic(err)
		}

		type Response struct {
			Status  string
			Message string
			Data    *joins_model.UserChannelJoin
		}

		bodyBinding := new(Response)
		err = json.Unmarshal(bodyRes, bodyBinding)
		if err != nil {
			panic(err)
		}

		response := rec.Result()
		t.Run(v.CaseDescription, func(t *testing.T) {
			assert.Equal(t, v.ExpectedHttpRespCode, response.StatusCode)
			assert.Equal(t, v.ExpectedHttpRespStatus, bodyBinding.Status)
			assert.Equal(t, v.ExpectedHttpRespMessage, bodyBinding.Message)
			assert.Equal(t, v.ExpectedHttpRespData.Email, bodyBinding.Data.Email)
		})
	}

	setup.Db.Exec("DELETE FROM sessions")
	setup.Db.Exec("DELETE FROM users")
}
