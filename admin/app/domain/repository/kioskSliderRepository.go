package repository

import (
	"errors"
	"stncCms/app/domain/entity"
	"strings"

	"github.com/jinzhu/gorm"
)

//KioskSliderRepo struct
type KioskSliderRepo struct {
	db *gorm.DB
}

//KioskSliderRepositoryInit initial
func KioskSliderRepositoryInit(db *gorm.DB) *KioskSliderRepo {
	return &KioskSliderRepo{db}
}

//KioskSliderRepo implements the repository.KioskSliderRepository interface
// var _ interfaces.dataAppInterface = &KioskSliderRepo{}

//Save data
func (r *KioskSliderRepo) Save(data *entity.KioskSlider) (*entity.KioskSlider, map[string]string) {
	dbErr := map[string]string{}
	//The images are uploaded to digital ocean spaces. So we need to prepend the url. This might not be your use case, if you are not uploading image to Digital Ocean.
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
func (r *KioskSliderRepo) Update(data *entity.KioskSlider) (*entity.KioskSlider, map[string]string) {
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
func (r *KioskSliderRepo) Count(dataTotalCount *int64) {
	var data entity.KioskSlider
	var count int64
	r.db.Debug().Model(data).Count(&count)
	*dataTotalCount = count
}

//Delete data
func (r *KioskSliderRepo) Delete(id uint64) error {
	var data entity.KioskSlider
	var err error
	err = r.db.Debug().Where("id = ?", id).Delete(&data).Error
	if err != nil {
		return errors.New("database error, please try again")
	}
	return nil
}

//GetByID get data
func (r *KioskSliderRepo) GetByID(id uint64) (*entity.KioskSlider, error) {
	var data entity.KioskSlider
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
func (r *KioskSliderRepo) GetAll() ([]entity.KioskSlider, error) {
	var datas []entity.KioskSlider
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
func (r *KioskSliderRepo) GetAllP(datasPerPage int, offset int) ([]entity.KioskSlider, error) {
	var datas []entity.KioskSlider
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
func (r *KioskSliderRepo) SetKioskSliderUpdate(id uint64, status int) {
	r.db.Debug().Table("kiosk_slider").Where("id = ?", id).Update("status", status)
}
