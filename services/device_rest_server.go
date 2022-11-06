package services

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jhmorais/device-manager/internal/usecases/contracts"
	"github.com/jhmorais/device-manager/utils"
)

type Handler struct {
	WorkerPort            string
	CreateDeviceUseCase   contracts.CreateDeviceUseCase
	DeleteDeviceUseCase   contracts.DeleteDeviceUseCase
	FindDeviceUseCase     contracts.FindDeviceUseCase
	FindDeviceByIDUseCase contracts.FindDeviceByIDUseCase
	ListDeviceUseCase     contracts.ListDeviceUseCase
}

func NewHTTPRouterDevice(createDeviceUseCase contracts.CreateDeviceUseCase,
	deleteDeviceUseCase contracts.DeleteDeviceUseCase,
	findDeviceByIDUseCase contracts.FindDeviceByIDUseCase,
	findDeviceUseCase contracts.FindDeviceUseCase,
	listDeviceUseCase contracts.ListDeviceUseCase) *mux.Router {
	router := mux.NewRouter()
	handler := Handler{
		CreateDeviceUseCase:   createDeviceUseCase,
		DeleteDeviceUseCase:   deleteDeviceUseCase,
		FindDeviceByIDUseCase: findDeviceByIDUseCase,
		FindDeviceUseCase:     findDeviceUseCase,
		ListDeviceUseCase:     listDeviceUseCase,
	}
	router.UseEncodedPath()
	router.Use(utils.CommonMiddleware)

	router.HandleFunc("/", handler.listDevices).Methods(http.MethodGet)
	router.HandleFunc("/{id}/add-query", handler.updateDevice).Methods(http.MethodPut)
	router.HandleFunc("/{id}", handler.deleteDevice).Methods(http.MethodDelete)
	router.HandleFunc("/{id}/clone", handler.createDevice).Methods(http.MethodPost)
	router.HandleFunc("/{id}", handler.getDevice).Methods(http.MethodGet)
	// router.HandleFunc("/{brand}/{name}", handler.getDevice).Methods(http.MethodGet)

	return router
}

func (h *Handler) listDevices(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	response, err := h.ListDeviceUseCase.Execute(ctx)
	if err != nil {
		fmt.Println(err.Error())
		utils.WriteErrModel(w, http.StatusNotFound, utils.NewErrorResponse("failed to get devices"))
	}

	fmt.Println(response)

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		fmt.Println(err.Error())
		utils.WriteErrModel(w, http.StatusInternalServerError,
			utils.NewErrorResponse("Failed to marshal presets response"))
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, string(jsonResponse))
}

func (h *Handler) getDevice(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) updateDevice(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) deleteDevice(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) createDevice(w http.ResponseWriter, r *http.Request) {

}
