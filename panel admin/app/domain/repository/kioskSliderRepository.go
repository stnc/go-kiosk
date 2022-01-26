package repository

import (
	"errors"
	"stncCms/app/domain/dto"
	"stncCms/app/domain/entity"
	"stncCms/app/domain/helpers/stnccollection"
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
func (r *KioskSliderRepo) Save(data *entity.KioskConnection) (*entity.KioskConnection, map[string]string) {
	dbErr := map[string]string{}
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

//GetAll all data
func (r *KioskSliderRepo) GetAll() ([]entity.KioskSlider, error) {
	var datas []entity.KioskSlider
	var err error
	// err = r.db.Debug().Order("created_at desc").Find(&datas).Error
	err = r.db.Debug().Where("status=1").Order("menu_order desc").Find(&datas).Error

	if err != nil {
		return nil, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return nil, errors.New("data not found")
	}
	return datas, nil
}

//GetAll all data
func (r *KioskSliderRepo) GetAllCatID(catID int) ([]entity.KioskSlider, error) {
	var datas []entity.KioskSlider
	var err error

	err = r.db.Raw(" SELECT * FROM kiosk_slider p1 WHERE p1.deleted_at IS NULL AND  p1.id IN (SELECT kiosk_id FROM  categories_kiosk_join  where categories_kiosk_join.category_id=?) AND STATUS=1", catID).Scan(&datas).Error
	if err != nil {
		return nil, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return nil, errors.New("data not found")
	}
	return datas, nil
}

func (r *KioskSliderRepo) GetAjaxData() (*dto.OptionsDto, error) {

	var data dto.OptionsDto
	var err error

	appOptions := OptionRepositoryInit(r.db)

	if appOptions.GetOption("pageRenewStatus") == "1" {
		data.PageRenewStatus = true
	} else {
		data.PageRenewStatus = false
	}

	data.PageRenewTime = stnccollection.StringToint(appOptions.GetOption("pageRenewTime"))

	data.Dolar = appOptions.GetOption("dolar")
	data.Euro = appOptions.GetOption("euro")
	data.Altin = appOptions.GetOption("altin")
	data.CeyrekAltin = appOptions.GetOption("ceyrek_altin")
	data.WeatherTodayIcon = appOptions.GetOption("weatherTodayIcon")
	data.WeatherTodayDescription = appOptions.GetOption("weatherTodayDescription")
	data.WeatherTodayDegree = appOptions.GetOption("weatherTodayDegree")
	data.WeatherTodayNight = appOptions.GetOption("weatherTodayNight")
	data.WeatherTodayHumidity = appOptions.GetOption("weatherTodayHumidity")
	data.Covid19Confirmed = appOptions.GetOption("Covid19Confirmed")
	data.Covid19Deaths = appOptions.GetOption("Covid19Deaths")
	data.Covid19Recovered = appOptions.GetOption("Covid19Recovered")
	data.Covid19Aktive = appOptions.GetOption("Covid19Aktive")

	if err != nil {
		return nil, errors.New("database error, please try again")
	}

	if gorm.IsRecordNotFoundError(err) {
		return nil, errors.New("data not found")
	}
	return &data, nil
}
