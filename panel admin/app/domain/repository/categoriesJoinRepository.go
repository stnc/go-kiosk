package repository

import (
	"errors"
	"stncCms/app/domain/entity"
	"strings"

	"github.com/jinzhu/gorm"
)

//CategoriesKioskRepo struct
type CategoriesKioskRepo struct {
	db *gorm.DB
}

//CategoriesKioskRepositoryInit initial
func CategoriesKioskRepositoryInit(db *gorm.DB) *CategoriesKioskRepo {
	return &CategoriesKioskRepo{db}
}

//PostRepo implements the repository.PostRepository interface
// var _ interfaces.CatAppInterface = &CategoriesKioskRepo{}

//Save data
func (r *CategoriesKioskRepo) Save(cat *entity.CategoriesKiosk) (*entity.CategoriesKiosk, map[string]string) {
	dbErr := map[string]string{}

	err := r.db.Debug().Create(&cat).Error
	if err != nil {
		//since our title is unique
		if strings.Contains(err.Error(), "duplicate") || strings.Contains(err.Error(), "Duplicate") {
			dbErr["unique_title"] = "post title already taken"
			return nil, dbErr
		}
		//any other db error
		dbErr["db_error"] = "database error"
		return nil, dbErr
	}
	return cat, nil
}

//Update upate data
func (r *CategoriesKioskRepo) Update(cat *entity.CategoriesKiosk) (*entity.CategoriesKiosk, map[string]string) {
	dbErr := map[string]string{}
	err := r.db.Debug().Save(&cat).Error
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
	return cat, nil
}

//GetByID get data
func (r *CategoriesKioskRepo) GetByID(id uint64) (*entity.CategoriesKiosk, error) {
	var cat entity.CategoriesKiosk
	err := r.db.Debug().Where("id = ?", id).Take(&cat).Error
	if err != nil {
		return nil, errors.New("database error, please try again")
	}
	if gorm.IsRecordNotFoundError(err) {
		return nil, errors.New("post not found")

	}
	return &cat, nil
}

//GetAll all data
func (r *CategoriesKioskRepo) GetAll() ([]entity.CategoriesKiosk, error) {
	var cat []entity.CategoriesKiosk
	err := r.db.Debug().Order("created_at desc").Find(&cat).Error
	if err != nil {
		return nil, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return nil, errors.New("post not found")
	}
	return cat, nil
}

//Delete delete data
func (r *CategoriesKioskRepo) Delete(id uint64) error {
	var cat entity.CategoriesKiosk
	err := r.db.Debug().Where("id = ?", id).Delete(&cat).Error
	if err != nil {
		return errors.New("database error, please try again")
	}
	return nil
}
