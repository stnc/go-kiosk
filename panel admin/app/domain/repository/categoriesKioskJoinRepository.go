package repository

import (
	"errors"
	"stncCms/app/domain/entity"
	"strings"

	"github.com/jinzhu/gorm"
)

//CatKioskJoinRepo struct
type CatKioskJoinRepo struct {
	db *gorm.DB
}

////SELECT * FROM kiosk_slider p1 WHERE id IN (SELECT kiosk_id FROM  categories_kiosk_join  where categories_kiosk_join.category_id=1)

//CatKioskJoinRepositoryInit initial
func CatKioskJoinRepositoryInit(db *gorm.DB) *CatKioskJoinRepo {
	return &CatKioskJoinRepo{db}
}

//PostRepo implements the repository.PostRepository interface
// var _ interfaces.CatKioskJoinAppInterface = &CatKioskJoinRepo{}

//Save data
func (r *CatKioskJoinRepo) Save(cat *entity.CategoriesKioskJoin) (*entity.CategoriesKioskJoin, map[string]string) {
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

//GetByID get data
func (r *CatKioskJoinRepo) GetByID(id uint64) (*entity.CategoriesKioskJoin, error) {
	var cat entity.CategoriesKioskJoin
	err := r.db.Debug().Where("id = ?", id).Take(&cat).Error
	if err != nil {
		return nil, errors.New("database error, please try again")
	}
	// if gorm.IsRecordNotFoundError(err) {
	// 	return nil, errors.New("post not found")

	// }
	return &cat, nil
}

//GetAllforCatID get data
func (r *CatKioskJoinRepo) GetAllforCatID(catid uint64) ([]entity.CategoriesKioskJoin, error) {
	var cat []entity.CategoriesKioskJoin
	err := r.db.Debug().Limit(100).Where("category_id = ?", catid).Order("created_at desc").Find(&cat).Error
	if err != nil {
		return nil, err
	}

	// if err.IsRecordNotFoundError(err) {
	// 	return nil, errors.New("post not found")
	// }
	return cat, nil
}

//GetAllforKioskID all data
func (r *CatKioskJoinRepo) GetAllforKioskID(KioskId uint64) ([]entity.CategoriesKioskJoin, error) {
	var cat []entity.CategoriesKioskJoin
	err := r.db.Debug().Limit(100).Where("kiosk_id = ?", KioskId).Order("created_at desc").Find(&cat).Error
	if err != nil {
		return nil, err
	}
	// if gorm.IsRecordNotFoundError(err) {
	// 	return nil, errors.New("post not found")
	// }
	return cat, nil
}

//GetAll all data
func (r *CatKioskJoinRepo) GetAll() ([]entity.CategoriesKioskJoin, error) {
	var cat []entity.CategoriesKioskJoin
	err := r.db.Debug().Limit(100).Order("created_at desc").Find(&cat).Error
	if err != nil {
		return nil, err
	}
	// if gorm.IsRecordNotFoundError(err) {
	// 	return nil, errors.New("post not found")
	// }
	return cat, nil
}

//Update upate data
func (r *CatKioskJoinRepo) Update(cat *entity.CategoriesKioskJoin) (*entity.CategoriesKioskJoin, map[string]string) {
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

//Delete delete data
func (r *CatKioskJoinRepo) Delete(id uint64) error {
	var cat entity.CategoriesKioskJoin
	err := r.db.Debug().Where("id = ?", id).Unscoped().Delete(&cat).Error
	if err != nil {
		return errors.New("database error, please try again")
	}
	return nil
}

//DeleteForKioskID delete data
func (r *CatKioskJoinRepo) DeleteForKioskID(KioskId uint64) error {
	var cat entity.CategoriesKioskJoin
	err := r.db.Debug().Where("kiosk_id = ?", KioskId).Unscoped().Delete(&cat).Error
	if err != nil {
		return errors.New("database error, please try again")
	}
	return nil
}

//DeleteForCatID delete data
func (r *CatKioskJoinRepo) DeleteForCatID(CatID uint64) error {
	var cat entity.CategoriesKioskJoin
	err := r.db.Debug().Where("category_id = ?", CatID).Unscoped().Delete(&cat).Error
	if err != nil {
		return errors.New("database error, please try again")
	}
	return nil
}
