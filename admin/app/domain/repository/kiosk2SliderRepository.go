package repository

import (
	"errors"
	"os"
	"stncCms/app/domain/entity"
	"strings"

	"github.com/jinzhu/gorm"
)

//KioskSliderRepo struct
type Kiosk2SliderRepo struct {
	db *gorm.DB
}

//KioskSliderRepositoryInit initial
func Kiosk2SliderRepositoryInit(db *gorm.DB) *Kiosk2SliderRepo {
	return &Kiosk2SliderRepo{db}
}

//KioskSliderRepo implements the repository.KioskSliderRepository interface
// var _ interfaces.dataAppInterface = &KioskSliderRepo{}

//Save data
func (r *Kiosk2SliderRepo) Save(data *entity.Kiosk2Slider) (*entity.Kiosk2Slider, map[string]string) {
	dbErr := map[string]string{}
	//The images are uploaded to digital ocean spaces. So we need to prepend the url. This might not be your use case, if you are not uploading image to Digital Ocean.
	data.Picture = os.Getenv("DO_SPACES_URL") + data.Picture
	var err error
	err = r.db.Debug().Create(&data).Error
	if err != nil {
		//since our title is unique
		if strings.Contains(err.Error(), "duplicate") || strings.Contains(err.Error(), "Duplicate") {
			dbErr["unique_title"] = "data title already taken"
			return nil, dbErr
		}
		//any other db error
		dbErr["db_error"] = "database error"
		return nil, dbErr
	}
	return data, nil
}

//Update upate data
func (r *Kiosk2SliderRepo) Update(data *entity.Kiosk2Slider) (*entity.Kiosk2Slider, map[string]string) {
	dbErr := map[string]string{}
	err := r.db.Debug().Save(&data).Error
	//db.Table("libraries").Where("id = ?", id).Update(dataData)

	if err != nil {
		//since our title is unique
		if strings.Contains(err.Error(), "duplicate") || strings.Contains(err.Error(), "Duplicate") {
			dbErr["unique_title"] = "title already taken"
			return nil, dbErr
		}
		//any other db error
		dbErr["db_error"] = "database error"
		return nil, dbErr
	}
	return data, nil
}

//Count fat
func (r *Kiosk2SliderRepo) Count(dataTotalCount *int64) {
	var data entity.KioskSlider
	var count int64
	r.db.Debug().Model(data).Count(&count)
	*dataTotalCount = count
}

//Delete data
func (r *Kiosk2SliderRepo) Delete(id uint64) error {
	var data entity.KioskSlider
	var err error
	err = r.db.Debug().Where("id = ?", id).Delete(&data).Error
	if err != nil {
		return errors.New("database error, please try again")
	}
	return nil
}

//GetByID get data
func (r *Kiosk2SliderRepo) GetByID(id uint64) (*entity.Kiosk2Slider, error) {
	var data entity.Kiosk2Slider
	var err error
	err = r.db.Debug().Where("id = ?", id).Take(&data).Error
	if err != nil {
		return nil, errors.New("database error, please try again")
	}
	if gorm.IsRecordNotFoundError(err) {
		return nil, errors.New("data not found")
	}
	return &data, nil
}

//GetAll all data
func (r *Kiosk2SliderRepo) GetAll() ([]entity.Kiosk2Slider, error) {
	var datas []entity.Kiosk2Slider
	var err error
	err = r.db.Debug().Order("created_at desc").Find(&datas).Error
	if err != nil {
		return nil, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return nil, errors.New("data not found")
	}
	return datas, nil
}

//GetAllP pagination all data
func (r *Kiosk2SliderRepo) GetAllP(datasPerPage int, offset int) ([]entity.Kiosk2Slider, error) {
	var datas []entity.Kiosk2Slider
	var err error
	err = r.db.Debug().Limit(datasPerPage).Offset(offset).Order("created_at desc").Find(&datas).Error
	if err != nil {
		return nil, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return nil, errors.New("data not found")
	}
	return datas, nil
}

//SetKioskSliderUpdate update data
func (r *Kiosk2SliderRepo) SetKioskSliderUpdate(id uint64, status int) {
	r.db.Debug().Table("kiosk_slider").Where("id = ?", id).Update("status", status)
}
