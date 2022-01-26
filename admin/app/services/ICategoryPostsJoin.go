package services

import (
	"stncCms/app/domain/entity"
)

//CCategoryPostsJoinAppInterface service
type CategoryPostsJoinAppInterface interface {
	Save(*entity.CategoryPostsJoin) (*entity.CategoryPostsJoin, map[string]string)
	GetAllforPostID(uint64) ([]entity.CategoryPostsJoin, error)
	GetAllforCatID(uint64) ([]entity.CategoryPostsJoin, error)
	GetAll() ([]entity.CategoryPostsJoin, error)
	Update(*entity.CategoryPostsJoin) (*entity.CategoryPostsJoin, map[string]string)
	Delete(uint64) error
	DeleteForPostID(uint64) error
	DeleteForCatID(uint64) error
}

//CCategoryPostsJoinApp struct  init
type CCategoryPostsJoinApp struct {
	fr CategoryPostsJoinAppInterface
}

var _ CategoryPostsJoinAppInterface = &CCategoryPostsJoinApp{}

//Save service init
func (f *CCategoryPostsJoinApp) Save(Cat *entity.CategoryPostsJoin) (*entity.CategoryPostsJoin, map[string]string) {
	return f.fr.Save(Cat)
}

//GetAll service init
func (f *CCategoryPostsJoinApp) GetAll() ([]entity.CategoryPostsJoin, error) {
	return f.fr.GetAll()
}

//GetAllforPostID service init
func (f *CCategoryPostsJoinApp) GetAllforPostID(PostID uint64) ([]entity.CategoryPostsJoin, error) {
	return f.fr.GetAllforPostID(PostID)
}

//GetAllforCatID service init
func (f *CCategoryPostsJoinApp) GetAllforCatID(CatID uint64) ([]entity.CategoryPostsJoin, error) {
	return f.fr.GetAllforCatID(CatID)
}

//Update service init
func (f *CCategoryPostsJoinApp) Update(Cat *entity.CategoryPostsJoin) (*entity.CategoryPostsJoin, map[string]string) {
	return f.fr.Update(Cat)
}

//Delete service init
func (f *CCategoryPostsJoinApp) Delete(ID uint64) error {
	return f.fr.Delete(ID)
}

//DeleteForPostID service init
func (f *CCategoryPostsJoinApp) DeleteForPostID(PostID uint64) error {
	return f.fr.DeleteForPostID(PostID)
}

//DeleteForCatID service init
func (f *CCategoryPostsJoinApp) DeleteForCatID(CatID uint64) error {
	return f.fr.DeleteForCatID(CatID)
}
