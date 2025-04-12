package repository

import (
	"github.com/DaigoSugiyama0317/map-memo-app/model"
	"gorm.io/gorm"
)

type ITagRepository interface {
	GetAllTags(tags *[]model.Tag, userId uint) error
	CreateTag(tag *model.Tag) error
	UpdateTag(tag *model.Tag, userId uint, tagId uint) error
	DeleteTag(tag *model.Tag, tagId uint) error
}

type tagRepository struct {
	db *gorm.DB
}

func (tr tagRepository) GetAllTags(tags *[]model.Tag, userId uint) error {

}

func (tr tagRepository) CreateTag(tag *model.Tag) error {

}

func (tr tagRepository) UpdateTag(tag *model.Tag, userId uint, tagId uint) error {

}

func (tr tagRepository) DeleteTag(tag *model.Tag, userId uint, tagId uint) error {
	
}