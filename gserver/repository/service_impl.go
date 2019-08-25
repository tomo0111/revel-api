package repository

import (
	"strings"

	"github.com/jinzhu/gorm"

	"github.com/tomoyane/grant-n-z/gserver/entity"
	"github.com/tomoyane/grant-n-z/gserver/log"
	"github.com/tomoyane/grant-n-z/gserver/model"
)

var srInstance ServiceRepository

type ServiceRepositoryImpl struct {
	Db *gorm.DB
}

func GetServiceRepositoryInstance(db *gorm.DB) ServiceRepository {
	if srInstance == nil {
		srInstance = NewServiceRepository(db)
	}
	return srInstance
}

func NewServiceRepository(db *gorm.DB) ServiceRepository {
	log.Logger.Info("New `ServiceRepository` instance")
	log.Logger.Info("Inject `gorm.DB` to `ServiceRepository`")
	return ServiceRepositoryImpl{
		Db: db,
	}
}

func (sri ServiceRepositoryImpl) FindAll() ([]*entity.Service, *model.ErrorResponse) {
	var services []*entity.Service
	if err := sri.Db.Find(&services).Error; err != nil {
		if strings.Contains(err.Error(), "record not found") {
			return nil, nil
		}

		return nil, model.InternalServerError(err.Error())
	}

	return services, nil
}

func (sri ServiceRepositoryImpl) FindById(id int) (*entity.Service, *model.ErrorResponse) {
	var service entity.Service
	if err := sri.Db.Where("id = ?", id).First(&service).Error; err != nil {
		if strings.Contains(err.Error(), "record not found") {
			return nil, nil
		}

		return nil, model.InternalServerError(err.Error())
	}

	return &service, nil
}

func (sri ServiceRepositoryImpl) FindByName(name string) (*entity.Service, *model.ErrorResponse) {
	var service entity.Service
	if err := sri.Db.Where("name = ?", name).First(&service).Error; err != nil {
		if strings.Contains(err.Error(), "record not found") {
			return nil, nil
		}

		return nil, model.InternalServerError(err.Error())
	}

	return &service, nil
}

func (sri ServiceRepositoryImpl) FindByApiKey(apiKey string) (*entity.Service, *model.ErrorResponse) {
	var service entity.Service
	if err := sri.Db.Where("api_key = ?", apiKey).First(&service).Error; err != nil {
		if strings.Contains(err.Error(), "record not found") {
			return nil, nil
		}

		return nil, model.InternalServerError(err.Error())
	}

	return &service, nil
}

func (sri ServiceRepositoryImpl) Save(service entity.Service) (*entity.Service, *model.ErrorResponse) {
	if err := sri.Db.Create(&service).Error; err != nil {
		log.Logger.Warn(err.Error())
		if strings.Contains(err.Error(), "1062") {
			return nil, model.Conflict("Already exit data.")
		}

		return nil, model.InternalServerError(err.Error())
	}

	return &service, nil
}

func (sri ServiceRepositoryImpl) Update(service entity.Service) *entity.Service {
	if err := sri.Db.Update(&service).Error; err != nil {
		return nil
	}

	return &service
}