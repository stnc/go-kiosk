package services

import (
	"stncCms/app/domain/entity"
)

//CCategoriesKioskJoinAppInterface service
type CategoriesKioskJoinAppInterface interface {
	Save(*entity.CategoriesKioskJoin) (*entity.CategoriesKioskJoin, map[string]string)
	GetAllforKioskID(uint64) ([]entity.CategoriesKioskJoin, error)
	GetAllforCatID(uint64) ([]entity.CategoriesKioskJoin, error)
	GetAll() ([]entity.CategoriesKioskJoin, error)
	Update(*entity.CategoriesKioskJoin) (*entity.CategoriesKioskJoin, map[string]string)
	Delete(uint64) error
	DeleteForKioskID(uint64) error
	DeleteForCatID(uint64) error
}

//CCategoriesKioskJoinApp struct  init
type CCategoriesKioskJoinApp struct {
	fr CategoriesKioskJoinAppInterface
}

var _ CategoriesKioskJoinAppInterface = &CCategoriesKioskJoinApp{}

//Save service init
func (f *CCategoriesKioskJoinApp) Save(Cat *entity.CategoriesKioskJoin) (*entity.CategoriesKioskJoin, map[string]string) {
	return f.fr.Save(Cat)
}

//GetAll service init
func (f *CCategoriesKioskJoinApp) GetAll() ([]entity.CategoriesKioskJoin, error) {
	return f.fr.GetAll()
}

//GetAllforKioskID service init
func (f *CCategoriesKioskJoinApp) GetAllforKioskID(KioskID uint64) ([]entity.CategoriesKioskJoin, error) {
	return f.fr.GetAllforKioskID(KioskID)
}

//GetAllforCatID service init
func (f *CCategoriesKioskJoinApp) GetAllforCatID(CatID uint64) ([]entity.CategoriesKioskJoin, error) {
	return f.fr.GetAllforCatID(CatID)
}

//Update service init
func (f *CCategoriesKioskJoinApp) Update(Cat *entity.CategoriesKioskJoin) (*entity.CategoriesKioskJoin, map[string]string) {
	return f.fr.Update(Cat)
}

//Delete service init
func (f *CCategoriesKioskJoinApp) Delete(ID uint64) error {
	return f.fr.Delete(ID)
}

//DeleteForKioskID service init
func (f *CCategoriesKioskJoinApp) DeleteForKioskID(KioskID uint64) error {
	return f.fr.DeleteForKioskID(KioskID)
}

//DeleteForCatID service init
func (f *CCategoriesKioskJoinApp) DeleteForCatID(CatID uint64) error {
	return f.fr.DeleteForCatID(CatID)
}
