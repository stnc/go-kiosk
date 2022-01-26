package services

import (
	"stncCms/app/domain/entity"
)

//CategoriesPostAppInterface service
type CategoriesPostAppInterface interface {
	Save(*entity.CategoriesPost) (*entity.CategoriesPost, map[string]string)
	GetByID(uint64) (*entity.CategoriesPost, error)
	GetAll() ([]entity.CategoriesPost, error)

	Update(*entity.CategoriesPost) (*entity.CategoriesPost, map[string]string)
	Delete(uint64) error
}

//CategoriesPostApp struct  init
type CategoriesPostApp struct {
	request CategoriesPostAppInterface
}

var _ CategoriesPostAppInterface = &CategoriesPostApp{}

//Save service init
func (f *CategoriesPostApp) Save(Cat *entity.CategoriesPost) (*entity.CategoriesPost, map[string]string) {
	return f.request.Save(Cat)
}

//GetAll service init
func (f *CategoriesPostApp) GetAll() ([]entity.CategoriesPost, error) {
	return f.request.GetAll()
}

//GetByID service init
func (f *CategoriesPostApp) GetByID(CatID uint64) (*entity.CategoriesPost, error) {
	return f.request.GetByID(CatID)
}

//Update service init
func (f *CategoriesPostApp) Update(Cat *entity.CategoriesPost) (*entity.CategoriesPost, map[string]string) {
	return f.request.Update(Cat)
}

//Delete service init
func (f *CategoriesPostApp) Delete(CatID uint64) error {
	return f.request.Delete(CatID)
}
