package di

import (
	"log"

	"github.com/jhmorais/device-manager/internal/repositories"
	"github.com/jhmorais/device-manager/internal/usecases"
	"github.com/jhmorais/device-manager/internal/usecases/contracts"
	"gorm.io/gorm"
)

type DenpencyBuild struct {
	DB           *gorm.DB
	Repositories Repositories
	Usecases     Usecases
}

type Repositories struct {
	DeviceRepository repositories.DeviceRepository
}

type Usecases struct {
	CreateDeviceUseCase   contracts.CreateDeviceUseCase
	DeleteDeviceUseCase   contracts.DeleteDeviceUseCase
	FindDeviceUseCase     contracts.FindDeviceUseCase
	FindDeviceByIDUseCase contracts.FindDeviceByIDUseCase
	ListDeviceUseCase     contracts.ListDeviceUseCase
}

func NewBuild() *DenpencyBuild {

	builder := &DenpencyBuild{}

	builder = builder.buildDB().
		buildRepositories().
		buildUseCases()

	return builder
}

func (d *DenpencyBuild) buildDB() *DenpencyBuild {
	var err error
	d.DB, err = InitGormMysqlDB()
	if err != nil {
		log.Fatal(err)
	}
	return d
}

func (d *DenpencyBuild) buildRepositories() *DenpencyBuild {
	d.Repositories.DeviceRepository = repositories.NewDeviceRepository(d.DB)
	return d
}

func (d *DenpencyBuild) buildUseCases() *DenpencyBuild {
	d.Usecases.CreateDeviceUseCase = usecases.NewCreateDeviceUseCase(d.Repositories.DeviceRepository)
	d.Usecases.DeleteDeviceUseCase = usecases.NewDeleteDeviceUseCase(d.Repositories.DeviceRepository)
	d.Usecases.FindDeviceUseCase = usecases.NewFindDeviceUseCase(d.Repositories.DeviceRepository)
	d.Usecases.FindDeviceByIDUseCase = usecases.NewFindDeviceByIDUseCase(d.Repositories.DeviceRepository)
	d.Usecases.ListDeviceUseCase = usecases.NewListDeviceUseCase(d.Repositories.DeviceRepository)

	return d
}
