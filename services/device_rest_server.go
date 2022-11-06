package services

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jhmorais/device-manager/internal/usecases/contracts"
	"github.com/jhmorais/device-manager/internal/usecases/ports/input"
	"github.com/jhmorais/device-manager/utils"
)

type Handler struct {
	WorkerPort               string
	CreateDeviceUseCase      contracts.CreateDeviceUseCase
	DeleteDeviceUseCase      contracts.DeleteDeviceUseCase
	FindDeviceByBrandUseCase contracts.FindDeviceByBrandUseCase
	FindDeviceByIDUseCase    contracts.FindDeviceByIDUseCase
	ListDeviceUseCase        contracts.ListDeviceUseCase
	UpdateDeviceUseCase      contracts.UpdateDeviceUseCase
}

func NewHTTPRouterDevice(createDeviceUseCase contracts.CreateDeviceUseCase,
	deleteDeviceUseCase contracts.DeleteDeviceUseCase,
	findDeviceByIDUseCase contracts.FindDeviceByIDUseCase,
	findDeviceByBrandUseCase contracts.FindDeviceByBrandUseCase,
	listDeviceUseCase contracts.ListDeviceUseCase,
	updateDeviceUseCase contracts.UpdateDeviceUseCase) *mux.Router {
	router := mux.NewRouter()
	handler := Handler{
		CreateDeviceUseCase:      createDeviceUseCase,
		DeleteDeviceUseCase:      deleteDeviceUseCase,
		FindDeviceByIDUseCase:    findDeviceByIDUseCase,
		FindDeviceByBrandUseCase: findDeviceByBrandUseCase,
		ListDeviceUseCase:        listDeviceUseCase,
		UpdateDeviceUseCase:      updateDeviceUseCase,
	}
	router.UseEncodedPath()
	router.Use(utils.CommonMiddleware)

	router.HandleFunc("/devices", handler.listDevices).Methods(http.MethodGet)
	router.HandleFunc("/devices/{id}", handler.getDeviceByID).Methods(http.MethodGet)
	router.HandleFunc("/devices/brand/{brand}", handler.getDevice).Methods(http.MethodGet)
	router.HandleFunc("/devices/{id}", handler.deleteDevice).Methods(http.MethodDelete)
	router.HandleFunc("/devices", handler.createDevice).Methods(http.MethodPost)
	router.HandleFunc("/devices/{id}", handler.updateDevice).Methods(http.MethodPut)

	return router
}

func (h *Handler) listDevices(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	response, err := h.ListDeviceUseCase.Execute(ctx)
	if err != nil {
		utils.WriteErrModel(w, http.StatusNotFound,
			utils.NewErrorResponse(fmt.Sprintf("failed to get devices, error: '%s'", err.Error())))
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		utils.WriteErrModel(w, http.StatusInternalServerError,
			utils.NewErrorResponse("Failed to marshal device response"))
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, string(jsonResponse))
}

func (h *Handler) getDeviceByID(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	id, err := utils.RetrieveParam(r, "id")
	if err != nil {
		utils.WriteErrModel(w, http.StatusBadRequest, utils.NewErrorResponse("error reading id"))
		return
	}

	response, err := h.FindDeviceByIDUseCase.Execute(ctx, id)
	if err != nil {
		utils.WriteErrModel(w, http.StatusNotFound,
			utils.NewErrorResponse(fmt.Sprintf("failed to find device, error: '%s'", err.Error())))
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		utils.WriteErrModel(w, http.StatusInternalServerError,
			utils.NewErrorResponse("Failed to marshal device response"))
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, string(jsonResponse))
}

func (h *Handler) getDevice(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	brand, err := utils.RetrieveParam(r, "brand")
	if err != nil {
		utils.WriteErrModel(w, http.StatusBadRequest, utils.NewErrorResponse("error reading brand"))
		return
	}

	response, err := h.FindDeviceByBrandUseCase.Execute(ctx, brand)
	if err != nil {
		utils.WriteErrModel(w, http.StatusNotFound,
			utils.NewErrorResponse(fmt.Sprintf("failed to find device, error: '%s'", err.Error())))
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		utils.WriteErrModel(w, http.StatusInternalServerError,
			utils.NewErrorResponse("Failed to marshal device response"))
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, string(jsonResponse))
}

func (h *Handler) updateDevice(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	id, err := utils.RetrieveParam(r, "id")
	if err != nil {
		utils.WriteErrModel(w, http.StatusBadRequest, utils.NewErrorResponse("error reading id"))
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		utils.WriteErrModel(w, http.StatusBadRequest, utils.NewErrorResponse("error reading request body"))
		return
	}

	device := input.UpdateDeviceInput{}
	err = json.Unmarshal(body, &device)
	if err != nil {
		utils.WriteErrModel(w, http.StatusBadRequest, utils.NewErrorResponse("failed to parse request body"))
		return
	}

	device.ID = id

	response, err := h.UpdateDeviceUseCase.Execute(ctx, &device)
	if err != nil {
		utils.WriteErrModel(w, http.StatusNotFound,
			utils.NewErrorResponse(fmt.Sprintf("failed to update device, error:'%s'", err.Error())))
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		utils.WriteErrModel(w, http.StatusInternalServerError,
			utils.NewErrorResponse("Failed to marshal device response"))
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, string(jsonResponse))
}

func (h *Handler) deleteDevice(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	id, err := utils.RetrieveParam(r, "id")
	if err != nil {
		utils.WriteErrModel(w, http.StatusBadRequest, utils.NewErrorResponse("error reading id"))
		return
	}

	response, err := h.DeleteDeviceUseCase.Execute(ctx, id)
	if err != nil {
		utils.WriteErrModel(w, http.StatusNotFound,
			utils.NewErrorResponse(fmt.Sprintf("failed to delete device, error: '%s'", err.Error())))
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		utils.WriteErrModel(w, http.StatusInternalServerError,
			utils.NewErrorResponse("Failed to marshal device response"))
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, string(jsonResponse))
}

func (h *Handler) createDevice(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	defer r.Body.Close()

	body, err := io.ReadAll(r.Body)
	if err != nil {
		utils.WriteErrModel(w, http.StatusBadRequest, utils.NewErrorResponse("error reading request body"))
		return
	}

	device := input.CreateDeviceInput{}
	err = json.Unmarshal(body, &device)
	if err != nil {
		utils.WriteErrModel(w, http.StatusBadRequest, utils.NewErrorResponse("failed to parse request body"))
		return
	}

	response, err := h.CreateDeviceUseCase.Execute(ctx, &device)
	if err != nil {
		utils.WriteErrModel(w, http.StatusNotFound,
			utils.NewErrorResponse(fmt.Sprintf("failed to create device, error: '%s'", err.Error())))
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		utils.WriteErrModel(w, http.StatusInternalServerError,
			utils.NewErrorResponse("Failed to marshal device response"))
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, string(jsonResponse))
}
