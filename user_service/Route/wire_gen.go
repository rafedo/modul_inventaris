// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package Route

import (
	"gorm.io/gorm"
	"user/Controller"
	"user/Repository"
	"user/Services"
)

// Injectors from wire.go:

// Vendor DI
func vendorDI(db *gorm.DB) *Controller.VendorControllerImpl {
	vendorRepositoryImpl := Repository.VendorRepositoryProvider(db)
	vendorServiceImpl := Services.VendorServiceProvider(vendorRepositoryImpl)
	vendorControllerImpl := Controller.VendorControllerProvider(vendorServiceImpl)
	return vendorControllerImpl
}

// User DI
func userDI(db *gorm.DB) *Controller.UserControllerImpl {
	userRepositoryImpl := Repository.UserRepositoryProvider(db)
	userServiceImpl := Services.UserServiceProvider(userRepositoryImpl)
	userControllerImpl := Controller.UserControllerProvider(userServiceImpl)
	return userControllerImpl
}

// Role DI
func roleDI(db *gorm.DB) *Controller.RoleControllerImpl {
	roleRepositoryImpl := Repository.RoleRepositoryProvider(db)
	roleServiceImpl := Services.RoleServiceControllerProvider(roleRepositoryImpl)
	roleControllerImpl := Controller.RoleControllerProvider(roleServiceImpl)
	return roleControllerImpl
}

// Divisi DI
func divisiDI(db *gorm.DB) *Controller.DivisiControllerImpl {
	divisiRepositoryImpl := Repository.DivisiRepositoryProvider(db)
	divisiServiceImpl := Services.DivisiServiceProvider(divisiRepositoryImpl)
	divisiControllerImpl := Controller.DivisiControllerProvider(divisiServiceImpl)
	return divisiControllerImpl
}