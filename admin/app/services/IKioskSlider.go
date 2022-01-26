package services

import (
	"stncCms/app/domain/entity"
)

//KioskSliderAppInterface interface
type KioskSliderAppInterface interface {
	Save(*entity.KioskSlider) (*entity.KioskSlider, map[string]string)
	GetByID(uint64) (*entity.KioskSlider, error)
	GetAll() ([]entity.KioskSlider, error)
	GetAllP(int, int) ([]entity.KioskSlider, error)
	Update(*entity.KioskSlider) (*entity.KioskSlider, map[string]string)
	Count(*int64)
	Delete(uint64) error
	SetKioskSliderUpdate(uint64, int)
}
type KioskSliderApp struct {
	request KioskSliderAppInterface
}

var _ KioskSliderAppInterface = &KioskSliderApp{}

func (f *KioskSliderApp) Count(KioskSliderTotalCount *int64) {
	f.request.Count(KioskSliderTotalCount)
}

func (f *KioskSliderApp) Save(KioskSlider *entity.KioskSlider) (*entity.KioskSlider, map[string]string) {
	return f.request.Save(KioskSlider)
}

func (f *KioskSliderApp) GetAll() ([]entity.KioskSlider, error) {
	return f.request.GetAll()
}

func (f *KioskSliderApp) GetAllP(KioskSlidersPerPage int, offset int) ([]entity.KioskSlider, error) {
	return f.request.GetAllP(KioskSlidersPerPage, offset)
}

func (f *KioskSliderApp) GetByID(KioskSliderID uint64) (*entity.KioskSlider, error) {
	return f.request.GetByID(KioskSliderID)
}

func (f *KioskSliderApp) Update(KioskSlider *entity.KioskSlider) (*entity.KioskSlider, map[string]string) {
	return f.request.Update(KioskSlider)
}

func (f *KioskSliderApp) Delete(KioskSliderID uint64) error {
	return f.request.Delete(KioskSliderID)
}
func (f *KioskSliderApp) SetKioskSliderUpdate(id uint64, status int) {
	f.request.SetKioskSliderUpdate(id, status)
}
