package data

import (
	"strings"

	"github.com/jinzhu/gorm"

	"github.com/tomoyane/grant-n-z/gserver/log"
	"github.com/tomoyane/grant-n-z/gserver/model"
	"github.com/tomoyane/grant-n-z/gserver/entity"
)

var prInstance PermissionRepository

type PermissionRepository interface {
	FindAll() ([]*entity.Permission, *model.ErrorResBody)

	FindById(id int) (*entity.Permission, *model.ErrorResBody)

	Save(permission entity.Permission) (*entity.Permission, *model.ErrorResBody)
}

type PermissionRepositoryImpl struct {
	Db *gorm.DB
}

func GetPermissionRepositoryInstance(db *gorm.DB) PermissionRepository {
	if prInstance == nil {
		prInstance = NewPermissionRepository(db)
	}
	return prInstance
}

func NewPermissionRepository(db *gorm.DB) PermissionRepository {
	log.Logger.Info("New `PermissionRepository` instance")
	log.Logger.Info("Inject `gorm.DB` to `PermissionRepository`")
	return PermissionRepositoryImpl{
		Db: db,
	}
}

func (pri PermissionRepositoryImpl) FindAll() ([]*entity.Permission, *model.ErrorResBody) {
	var permissions []*entity.Permission
	if err := pri.Db.Find(&permissions).Error; err != nil {
		if strings.Contains(err.Error(), "record not found") {
			return nil, nil
		}

		return nil, model.InternalServerError(err.Error())
	}

	return permissions, nil
}

func (pri PermissionRepositoryImpl) FindById(id int) (*entity.Permission, *model.ErrorResBody) {
	var permissions entity.Permission
	if err := pri.Db.Where("id = ?", id).Find(&permissions).Error; err != nil {
		if strings.Contains(err.Error(), "record not found") {
			return nil, nil
		}

		return nil, model.InternalServerError(err.Error())
	}

	return &permissions, nil
}

func (pri PermissionRepositoryImpl) Save(permission entity.Permission) (*entity.Permission, *model.ErrorResBody) {
	if err := pri.Db.Create(&permission).Error; err != nil {
		log.Logger.Warn(err.Error())
		if strings.Contains(err.Error(), "1062") {
			return nil, model.Conflict("Already exit data.")
		}

		return nil, model.InternalServerError(err.Error())
	}

	return &permission, nil
}