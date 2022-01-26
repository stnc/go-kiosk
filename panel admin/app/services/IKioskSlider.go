package services

import (
	"stncCms/app/domain/dto"
	"stncCms/app/domain/entity"
)

//KioskSliderAppInterface interface
type KioskSliderAppInterface interface {
	Save(*entity.KioskConnection) (*entity.KioskConnection, map[string]string)
	// GetByID(uint64) (*entity.KioskSlider, error)
	GetAll() ([]entity.KioskSlider, error)
	GetAjaxData() (*dto.OptionsDto, error)

	// Update(*entity.KioskSlider) (*entity.KioskSlider, map[string]string)
	GetAllCatID(int) ([]entity.KioskSlider, error)
}
type KioskSliderApp struct {
	request KioskSliderAppInterface
}

var _ KioskSliderAppInterface = &KioskSliderApp{}

func (f *KioskSliderApp) Save(kioskConnectiondata *entity.KioskConnection) (*entity.KioskConnection, map[string]string) {
	return f.request.Save(kioskConnectiondata)
}

func (f *KioskSliderApp) GetAll() ([]entity.KioskSlider, error) {
	return f.request.GetAll()
}

func (f *KioskSliderApp) GetAjaxData() (*dto.OptionsDto, error) {
	return f.request.GetAjaxData()
}

func (f *KioskSliderApp) GetAllCatID(catID int) ([]entity.KioskSlider, error) {
	return f.request.GetAllCatID(catID)
}

// func (f *KioskSliderApp) Update(KioskSlider *entity.KioskSlider) (*entity.KioskSlider, map[string]string) {
// 	return f.request.Update(KioskSlider)
// }
