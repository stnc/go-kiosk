package services

import (
	"stncCms/app/domain/entity"
)

//CategoriesKioskAppInterface service
type CategoriesKioskAppInterface interface {
	Save(*entity.CategoriesKiosk) (*entity.CategoriesKiosk, map[string]string)
	GetByID(uint64) (*entity.CategoriesKiosk, error)
	GetAll() ([]entity.CategoriesKiosk, error)

	Update(*entity.CategoriesKiosk) (*entity.CategoriesKiosk, map[string]string)
	Delete(uint64) error
}

//CategoriesKioskApp struct  init
type CategoriesKioskApp struct {
	request CategoriesKioskAppInterface
}

var _ CategoriesKioskAppInterface = &CategoriesKioskApp{}

//Save service init
func (f *CategoriesKioskApp) Save(Cat *entity.CategoriesKiosk) (*entity.CategoriesKiosk, map[string]string) {
	return f.request.Save(Cat)
}

//GetAll service init
func (f *CategoriesKioskApp) GetAll() ([]entity.CategoriesKiosk, error) {
	return f.request.GetAll()
}

//GetByID service init
func (f *CategoriesKioskApp) GetByID(CatID uint64) (*entity.CategoriesKiosk, error) {
	return f.request.GetByID(CatID)
}

//Update service init
func (f *CategoriesKioskApp) Update(Cat *entity.CategoriesKiosk) (*entity.CategoriesKiosk, map[string]string) {
	return f.request.Update(Cat)
}

//Delete service init
func (f *CategoriesKioskApp) Delete(CatID uint64) error {
	return f.request.Delete(CatID)
}
