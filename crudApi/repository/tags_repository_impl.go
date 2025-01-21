package repository

import (
	"crudApi/helper"
	"crudApi/model"

	"gorm.io/gorm"
)

type TagsRepositoryImpl struct {
	Db *gorm.DB
}

func NewTagsRepositoryImpl(Db *gorm.DB) TagsRepository {
	return &TagsRepositoryImpl{Db: Db}
}

// delete implements TagsRepository
func (t *TagsRepositoryImpl) Delete(tagsId int)  {
	var tags model.Tags
	result := t.Db.Where("id = ?", tagsId).Delete(&tags)
	helper.ErrorPanic(result.Error)
}

// findAll implements TagsRepo
func (t * TagsRepositoryImpl) FindAll() []model.Tags {
	var tags []model.Tags
	result := t.Db.Find(&tags)
	helper.ErrorPanic(result.Error)
	return tags 
}

// findById implements TagsRepo
func (t * TagsRepositoryImpl) FindById(tagsId int) (tags model.Tags, err error){
	var tag model.Tags
	result := t.Db.Find(&tag, tagsId)
	if result != nil {
		return tag, nil
	}else {
		return tag, error.New("Tag not found")
	}
}

// Save implements TagsRepo
func (t * TagsRepositoryImpl) Save(tags model.Tags) {
	result := t.Db.Save(&tags)
	helper.ErrorPanic(result.Error)
}

// update implements TagsRepo
func (t * TagsRepositoryImpl) Update(tags model.Tags) {
	
}