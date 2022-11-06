package repositories

import (
	"context"

	"github.com/jhmorais/device-manager/internal/domain/entities"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type deviceRepository struct {
	db *gorm.DB
}

func NewDeviceRepository(db *gorm.DB) DeviceRepository {
	return &deviceRepository{db: db}
}

func (d *deviceRepository) CreateDevice(ctx context.Context, entity *entities.Device) error {
	return d.db.
		Session(&gorm.Session{FullSaveAssociations: false}).
		Create(entity).
		Error
}

func (d *deviceRepository) UpdateDevice(ctx context.Context, entity *entities.Device) error {
	return d.db.
		Session(&gorm.Session{FullSaveAssociations: false}).
		Save(entity).
		Error
}

func (d *deviceRepository) DeleteDevice(ctx context.Context, entity *entities.Device) error {
	return d.db.
		Session(&gorm.Session{FullSaveAssociations: false}).
		Delete(entity).
		Error
}

func (d *deviceRepository) FindDeviceByID(ctx context.Context, id string) (*entities.Device, error) {
	var entity *entities.Device

	err := d.db.
		Preload(clause.Associations).
		Where("id = ?", id).
		Limit(1).
		Find(&entity).Error

	return entity, err
}

func (d *deviceRepository) FindDeviceByBrand(ctx context.Context, brand string) ([]*entities.Device, error) {
	var entity []*entities.Device

	err := d.db.
		Preload(clause.Associations).
		Where("brand = ?", brand).
		Limit(100).
		Find(&entity).Error

	return entity, err
}

func (d *deviceRepository) FindDevice(ctx context.Context, brand string, name string) (*entities.Device, error) {
	var entity *entities.Device

	err := d.db.
		Preload(clause.Associations).
		Where("brand = ?", brand).
		Where("name = ?", name).
		Limit(1).
		Find(&entity).Error

	return entity, err
}

func (d *deviceRepository) ListDevice(ctx context.Context) ([]*entities.Device, error) {
	//TODO impl pagination
	var entities []*entities.Device

	err := d.db.
		Preload(clause.Associations).
		Limit(100).
		Order("created_at desc").
		Find(&entities).Error

	if err != nil {
		return nil, err
	}

	return entities, nil
}
