package repository

import (
	"errors"
	"stncCms/app/domain/entity"
	"strings"

	"github.com/jinzhu/gorm"
)

//CatPostJoinRepo struct
type CatPostJoinRepo struct {
	db *gorm.DB
}

//CatPostJoinRepositoryInit initial
func CatPostJoinRepositoryInit(db *gorm.DB) *CatPostJoinRepo {
	return &CatPostJoinRepo{db}
}

//PostRepo implements the repository.PostRepository interface
// var _ interfaces.CatPostJoinAppInterface = &CatPostJoinRepo{}

//Save data
func (r *CatPostJoinRepo) Save(cat *entity.CategoryPostsJoin) (*entity.CategoryPostsJoin, map[string]string) {
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
func (r *CatPostJoinRepo) GetByID(id uint64) (*entity.CategoryPostsJoin, error) {
	var cat entity.CategoryPostsJoin
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
func (r *CatPostJoinRepo) GetAllforCatID(catid uint64) ([]entity.CategoryPostsJoin, error) {
	var cat []entity.CategoryPostsJoin
	err := r.db.Debug().Limit(100).Where("category_id = ?", catid).Order("created_at desc").Find(&cat).Error
	if err != nil {
		return nil, err
	}

	// if err.IsRecordNotFoundError(err) {
	// 	return nil, errors.New("post not found")
	// }
	return cat, nil
}

//GetAllforPostID all data
func (r *CatPostJoinRepo) GetAllforPostID(postid uint64) ([]entity.CategoryPostsJoin, error) {
	var cat []entity.CategoryPostsJoin
	err := r.db.Debug().Limit(100).Where("post_id = ?", postid).Order("created_at desc").Find(&cat).Error
	if err != nil {
		return nil, err
	}
	// if gorm.IsRecordNotFoundError(err) {
	// 	return nil, errors.New("post not found")
	// }
	return cat, nil
}

//GetAll all data
func (r *CatPostJoinRepo) GetAll() ([]entity.CategoryPostsJoin, error) {
	var cat []entity.CategoryPostsJoin
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
func (r *CatPostJoinRepo) Update(cat *entity.CategoryPostsJoin) (*entity.CategoryPostsJoin, map[string]string) {
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
func (r *CatPostJoinRepo) Delete(id uint64) error {
	var cat entity.CategoryPostsJoin
	err := r.db.Debug().Where("id = ?", id).Unscoped().Delete(&cat).Error
	if err != nil {
		return errors.New("database error, please try again")
	}
	return nil
}

//DeleteForPostID delete data
func (r *CatPostJoinRepo) DeleteForPostID(postID uint64) error {
	var cat entity.CategoryPostsJoin
	err := r.db.Debug().Where("post_id = ?", postID).Unscoped().Delete(&cat).Error
	if err != nil {
		return errors.New("database error, please try again")
	}
	return nil
}

//DeleteForCatID delete data
func (r *CatPostJoinRepo) DeleteForCatID(CatID uint64) error {
	var cat entity.CategoryPostsJoin
	err := r.db.Debug().Where("category_id = ?", CatID).Unscoped().Delete(&cat).Error
	if err != nil {
		return errors.New("database error, please try again")
	}
	return nil
}
