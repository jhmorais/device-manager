package services_test

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/jhmorais/device-manager/config"
	"github.com/jhmorais/device-manager/internal/usecases"
	"github.com/jhmorais/device-manager/internal/usecases/ports/input"
	"github.com/jhmorais/device-manager/mocks"
	"github.com/jhmorais/device-manager/services"
	"github.com/jhmorais/device-manager/utils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestServerCreateDevice(t *testing.T) {
	w := httptest.NewRecorder()
	config.LoadServerEnvironmentVars()

	t.Run("when the name is empty should return an error", func(t *testing.T) {
		deviceRepositoryErrorMock := &mocks.DeviceRepository{}
		deviceRepositoryErrorMock.On("CreateDevice", mock.Anything, mock.Anything).Return(nil)
		deviceRepositoryErrorMock.On("FindDevice", mock.Anything, mock.Anything, mock.Anything).Return(nil, nil)

		createDeviceUseCase := usecases.NewCreateDeviceUseCase(deviceRepositoryErrorMock)
		server := services.NewHTTPRouterDevice(createDeviceUseCase, nil, nil, nil, nil, nil)
		req := httptest.NewRequest("POST", "/devices", utils.ValidJSON(input.CreateDeviceInput{
			Name:  "",
			Brand: "Some brand",
		}))
		server.ServeHTTP(w, req)

		data, err := io.ReadAll(w.Body)
		assert.Nil(t, err)

		response := w.Result()
		defer response.Body.Close()
		assert.Equal(t, http.StatusNotFound, response.StatusCode)
		model := utils.ErrorModel{}
		err = json.Unmarshal(data, &model)
		assert.Nil(t, err)
		assert.NotNil(t, model.Message)

	})

	t.Run("when the brand is empty should return an error", func(t *testing.T) {
		deviceRepositoryErrorMock := &mocks.DeviceRepository{}
		deviceRepositoryErrorMock.On("CreateDevice", mock.Anything, mock.Anything).Return(nil)
		deviceRepositoryErrorMock.On("FindDevice", mock.Anything, mock.Anything, mock.Anything).Return(nil, nil)

		createDeviceUseCase := usecases.NewCreateDeviceUseCase(deviceRepositoryErrorMock)
		server := services.NewHTTPRouterDevice(createDeviceUseCase, nil, nil, nil, nil, nil)
		req := httptest.NewRequest("POST", "/devices", utils.ValidJSON(input.CreateDeviceInput{
			Name:  "Iphone 10",
			Brand: "",
		}))
		server.ServeHTTP(w, req)

		data, err := io.ReadAll(w.Body)
		assert.Nil(t, err)

		response := w.Result()
		defer response.Body.Close()
		assert.Equal(t, http.StatusNotFound, response.StatusCode)
		model := utils.ErrorModel{}
		err = json.Unmarshal(data, &model)
		assert.Nil(t, err)
		assert.NotNil(t, model.Message)

	})

	t.Run("when the body is null should return an error", func(t *testing.T) {
		deviceRepositoryErrorMock := &mocks.DeviceRepository{}
		deviceRepositoryErrorMock.On("CreateDevice", mock.Anything, mock.Anything).Return(nil)
		deviceRepositoryErrorMock.On("FindDevice", mock.Anything, mock.Anything, mock.Anything).Return(nil, nil)

		createDeviceUseCase := usecases.NewCreateDeviceUseCase(deviceRepositoryErrorMock)
		server := services.NewHTTPRouterDevice(createDeviceUseCase, nil, nil, nil, nil, nil)
		req := httptest.NewRequest("POST", "/devices", nil)
		server.ServeHTTP(w, req)

		data, err := io.ReadAll(w.Body)
		assert.Nil(t, err)

		response := w.Result()
		defer response.Body.Close()
		assert.Equal(t, http.StatusNotFound, response.StatusCode)
		model := utils.ErrorModel{}
		err = json.Unmarshal(data, &model)
		assert.Nil(t, err)
		assert.NotNil(t, model.Message)
	})

	t.Run("when the body is invalid should return an error", func(t *testing.T) {
		deviceRepositoryErrorMock := &mocks.DeviceRepository{}
		deviceRepositoryErrorMock.On("CreateDevice", mock.Anything, mock.Anything).Return(nil)
		deviceRepositoryErrorMock.On("FindDevice", mock.Anything, mock.Anything, mock.Anything).Return(nil, nil)

		createDeviceUseCase := usecases.NewCreateDeviceUseCase(deviceRepositoryErrorMock)
		server := services.NewHTTPRouterDevice(createDeviceUseCase, nil, nil, nil, nil, nil)
		req := httptest.NewRequest("POST", "/devices", utils.ValidJSON(utils.ErrorModel{
			Message: "wrong body",
		}))
		server.ServeHTTP(w, req)

		data, err := io.ReadAll(w.Body)
		assert.Nil(t, err)

		response := w.Result()
		defer response.Body.Close()
		assert.Equal(t, http.StatusNotFound, response.StatusCode)
		model := utils.ErrorModel{}
		err = json.Unmarshal(data, &model)
		assert.Nil(t, err)
		assert.NotNil(t, model.Message)
	})

	t.Run("when db fail should return an error", func(t *testing.T) {
		deviceRepositoryErrorMock := &mocks.DeviceRepository{}
		deviceRepositoryErrorMock.On("CreateDevice", mock.Anything, mock.Anything).Return(errors.New("error"))
		deviceRepositoryErrorMock.On("FindDevice", mock.Anything, mock.Anything, mock.Anything).Return(nil, nil)

		createDeviceUseCase := usecases.NewCreateDeviceUseCase(deviceRepositoryErrorMock)
		server := services.NewHTTPRouterDevice(createDeviceUseCase, nil, nil, nil, nil, nil)
		req := httptest.NewRequest("POST", "/devices", utils.ValidJSON(input.CreateDeviceInput{
			Name:  "ABCD",
			Brand: "Some brand",
		}))
		server.ServeHTTP(w, req)

		data, err := io.ReadAll(w.Body)
		assert.Nil(t, err)

		response := w.Result()
		defer response.Body.Close()
		assert.Equal(t, http.StatusNotFound, response.StatusCode)
		model := utils.ErrorModel{}
		err = json.Unmarshal(data, &model)
		assert.Nil(t, err)
		assert.NotNil(t, model.Message)
	})
}

func TestServerDeleteDevice(t *testing.T) {
	w := httptest.NewRecorder()
	config.LoadServerEnvironmentVars()

	t.Run("when the id is empty should return an error", func(t *testing.T) {
		deviceRepositoryErrorMock := &mocks.DeviceRepository{}
		deviceRepositoryErrorMock.On("DeleteDevice", mock.Anything, mock.Anything).Return(nil)
		deviceRepositoryErrorMock.On("FindDevice", mock.Anything, mock.Anything, mock.Anything).Return(nil, nil)

		deleteDeviceUseCase := usecases.NewDeleteDeviceUseCase(deviceRepositoryErrorMock)
		server := services.NewHTTPRouterDevice(nil, deleteDeviceUseCase, nil, nil, nil, nil)
		req := httptest.NewRequest("DELETE", "/devices/null", nil)
		server.ServeHTTP(w, req)

		data, err := io.ReadAll(w.Body)
		assert.Nil(t, err)

		response := w.Result()
		defer response.Body.Close()
		assert.Equal(t, http.StatusBadRequest, response.StatusCode)
		model := utils.ErrorModel{}
		err = json.Unmarshal(data, &model)
		assert.Nil(t, err)
		assert.NotNil(t, model.Message)
	})
}

func TestServerUpdateDevice(t *testing.T) {
	w := httptest.NewRecorder()
	config.LoadServerEnvironmentVars()

	t.Run("when the id is empty should return an error", func(t *testing.T) {
		deviceRepositoryErrorMock := &mocks.DeviceRepository{}
		deviceRepositoryErrorMock.On("UpdateDevice", mock.Anything, mock.Anything).Return(nil)
		deviceRepositoryErrorMock.On("FindDevice", mock.Anything, mock.Anything, mock.Anything).Return(nil, nil)

		updateDeviceUseCase := usecases.NewUpdateDeviceUseCase(deviceRepositoryErrorMock)
		server := services.NewHTTPRouterDevice(nil, nil, nil, nil, nil, updateDeviceUseCase)
		req := httptest.NewRequest("PUT", "/devices/null", utils.ValidJSON(input.UpdateDeviceInput{
			Name:  "S20",
			Brand: "SAMSUNG",
		}))
		server.ServeHTTP(w, req)

		data, err := io.ReadAll(w.Body)
		assert.Nil(t, err)

		response := w.Result()
		defer response.Body.Close()
		assert.Equal(t, http.StatusBadRequest, response.StatusCode)
		model := utils.ErrorModel{}
		err = json.Unmarshal(data, &model)
		assert.Nil(t, err)
		assert.NotNil(t, model.Message)
	})

	t.Run("when the body is invalid should return an error", func(t *testing.T) {
		deviceRepositoryErrorMock := &mocks.DeviceRepository{}
		deviceRepositoryErrorMock.On("UpdateDevice", mock.Anything, mock.Anything).Return(nil)
		deviceRepositoryErrorMock.On("FindDevice", mock.Anything, mock.Anything, mock.Anything).Return(nil, nil)

		updateDeviceUseCase := usecases.NewUpdateDeviceUseCase(deviceRepositoryErrorMock)
		server := services.NewHTTPRouterDevice(nil, nil, nil, nil, nil, updateDeviceUseCase)
		req := httptest.NewRequest("PUT", "/devices/63cc3dad-dc78-4f7d-87cc-75f708cab561", utils.ValidJSON(utils.ErrorModel{
			Message: "wrong body",
		}))
		server.ServeHTTP(w, req)

		data, err := io.ReadAll(w.Body)
		assert.Nil(t, err)

		response := w.Result()
		defer response.Body.Close()
		assert.Equal(t, http.StatusBadRequest, response.StatusCode)
		model := utils.ErrorModel{}
		err = json.Unmarshal(data, &model)
		assert.Nil(t, err)
		assert.NotNil(t, model.Message)
	})
}

func TestServerGetDevice(t *testing.T) {
	w := httptest.NewRecorder()
	config.LoadServerEnvironmentVars()

	t.Run("when the brand is empty should return an error", func(t *testing.T) {
		deviceRepositoryErrorMock := &mocks.DeviceRepository{}
		deviceRepositoryErrorMock.On("FindDeviceByBrand", mock.Anything, mock.Anything).Return(nil, nil)

		findDeviceByBrandUseCase := usecases.NewFindDeviceByBrandUseCase(deviceRepositoryErrorMock)
		server := services.NewHTTPRouterDevice(nil, nil, nil, findDeviceByBrandUseCase, nil, nil)
		req := httptest.NewRequest("GET", "/devices/brand/null", nil)
		server.ServeHTTP(w, req)

		data, err := io.ReadAll(w.Body)
		assert.Nil(t, err)

		response := w.Result()
		defer response.Body.Close()
		assert.Equal(t, http.StatusNotFound, response.StatusCode)
		model := utils.ErrorModel{}
		err = json.Unmarshal(data, &model)
		assert.Nil(t, err)
		assert.NotNil(t, model.Message)
	})
}

func TestServerGetDeviceByID(t *testing.T) {
	w := httptest.NewRecorder()
	config.LoadServerEnvironmentVars()

	t.Run("when the brand is empty should return an error", func(t *testing.T) {
		deviceRepositoryErrorMock := &mocks.DeviceRepository{}
		deviceRepositoryErrorMock.On("FindDeviceByID", mock.Anything, mock.Anything).Return(nil, nil)

		findDeviceByBrandUseCase := usecases.NewFindDeviceByIDUseCase(deviceRepositoryErrorMock)
		server := services.NewHTTPRouterDevice(nil, nil, findDeviceByBrandUseCase, nil, nil, nil)
		req := httptest.NewRequest("GET", "/devices/null", nil)
		server.ServeHTTP(w, req)

		data, err := io.ReadAll(w.Body)
		assert.Nil(t, err)

		response := w.Result()
		defer response.Body.Close()
		assert.Equal(t, http.StatusNotFound, response.StatusCode)
		model := utils.ErrorModel{}
		err = json.Unmarshal(data, &model)
		assert.Nil(t, err)
		assert.NotNil(t, model.Message)
	})
}

func TestServerListDevice(t *testing.T) {
	w := httptest.NewRecorder()
	config.LoadServerEnvironmentVars()

	t.Run("when the brand is empty should return an error", func(t *testing.T) {
		deviceRepositoryErrorMock := &mocks.DeviceRepository{}
		deviceRepositoryErrorMock.On("ListDevice", mock.Anything).Return(nil, errors.New("no devices found"))

		listDeviceUseCase := usecases.NewListDeviceUseCase(deviceRepositoryErrorMock)
		server := services.NewHTTPRouterDevice(nil, nil, nil, nil, listDeviceUseCase, nil)
		req := httptest.NewRequest("GET", "/devices", nil)
		server.ServeHTTP(w, req)

		data, err := io.ReadAll(w.Body)
		assert.Nil(t, err)

		response := w.Result()
		defer response.Body.Close()
		assert.Equal(t, http.StatusNotFound, response.StatusCode)
		model := utils.ErrorModel{}
		err = json.Unmarshal(data, &model)
		assert.Nil(t, err)
		assert.NotNil(t, model.Message)
	})
}
