package services

import (
	"stncCms/app/domain/entity"
)

//KioskSliderAppInterface interface
type Kiosk2SliderAppInterface interface {
	Save(*entity.Kiosk2Slider) (*entity.Kiosk2Slider, map[string]string)
	GetByID(uint64) (*entity.Kiosk2Slider, error)
	GetAll() ([]entity.Kiosk2Slider, error)
	GetAllP(int, int) ([]entity.Kiosk2Slider, error)
	Update(*entity.Kiosk2Slider) (*entity.Kiosk2Slider, map[string]string)
	Count(*int64)
	Delete(uint64) error
	SetKioskSliderUpdate(uint64, int)
}
type Kiosk2SliderApp struct {
	request Kiosk2SliderAppInterface
}

var _ KioskSliderAppInterface = &KioskSliderApp{}

func (f *Kiosk2SliderApp) Count(KioskSliderTotalCount *int64) {
	f.request.Count(KioskSliderTotalCount)
}

func (f *Kiosk2SliderApp) Save(KioskSlider *entity.Kiosk2Slider) (*entity.Kiosk2Slider, map[string]string) {
	return f.request.Save(KioskSlider)
}

func (f *Kiosk2SliderApp) GetAll() ([]entity.Kiosk2Slider, error) {
	return f.request.GetAll()
}

func (f *Kiosk2SliderApp) GetAllP(KioskSlidersPerPage int, offset int) ([]entity.Kiosk2Slider, error) {
	return f.request.GetAllP(KioskSlidersPerPage, offset)
}

func (f *Kiosk2SliderApp) GetByID(KioskSliderID uint64) (*entity.Kiosk2Slider, error) {
	return f.request.GetByID(KioskSliderID)
}

func (f *Kiosk2SliderApp) Update(KioskSlider *entity.Kiosk2Slider) (*entity.Kiosk2Slider, map[string]string) {
	return f.request.Update(KioskSlider)
}

func (f *Kiosk2SliderApp) Delete(KioskSliderID uint64) error {
	return f.request.Delete(KioskSliderID)
}
func (f *Kiosk2SliderApp) SetKioskSliderUpdate(id uint64, status int) {
	f.request.SetKioskSliderUpdate(id, status)
}
