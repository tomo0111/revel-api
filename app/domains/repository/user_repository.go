package repository

import (
	"authentication-server/app/domains/entity"
	"authentication-server/app"
	"authentication-server/app/controllers/base"
)

type UserRepository struct{}

// Find users by users.email
func (r UserRepository) FindByEmail(email string) base.BaseResponse {
	var users entity.Users
	if err := app.Db.Where("email = ?", email).First(&users).Error; err != nil {
		return base.BaseResponse{Response: "server error"}
	}

	response := base.BaseResponse{}
	response.Response = users

	return response
}

// Save to users
func (r UserRepository) Save(users entity.Users) base.BaseResponse {
	if err := app.Db.Create(&users).Error; err != nil {
		return base.BaseResponse{Response: "server error"}
	}

	response := base.BaseResponse{}
	response.Response = users

	return response
}

// Update to users
func (r UserRepository) Update(users entity.Users) base.BaseResponse {
	if err := app.Db.Update(&users).Error; err != nil {
		return base.BaseResponse{Response: "server error"}
	}

	response := base.BaseResponse{}
	response.Response = users

	return response
}
