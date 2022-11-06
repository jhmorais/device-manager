package main

import (
	"fmt"
	"net/http"

	"github.com/jhmorais/device-manager/config"
	"github.com/jhmorais/device-manager/internal/infra/di"
	"github.com/jhmorais/device-manager/services"
)

func main() {
	config.LoadServerEnvironmentVars()

	dependencies := di.NewBuild()

	router := services.NewHTTPRouterDevice(dependencies.Usecases.CreateDeviceUseCase,
		dependencies.Usecases.DeleteDeviceUseCase,
		dependencies.Usecases.FindDeviceByIDUseCase,
		dependencies.Usecases.FindDeviceUseCase,
		dependencies.Usecases.ListDeviceUseCase)

	deviceErr := http.ListenAndServe(fmt.Sprintf(":%s", config.GetServerPort()), router)
	if deviceErr != nil && deviceErr != http.ErrServerClosed {
		fmt.Println("failed to create server rest on port: " + config.GetServerPort())
		fmt.Println(deviceErr.Error())
	}
	fmt.Println("SERVER LISTEN PORT: " + config.GetServerPort())
}
