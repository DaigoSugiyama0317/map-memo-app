package repository

import (
	"fmt"

	"github.com/DaigoSugiyama0317/map-memo-app/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ITagRepository interface {
	GetAllTags(tags *[]model.Tag, userId uint) error
	CreateTag(tag *model.Tag) error
	UpdateTag(tag *model.Tag, userId uint, tagId uint) error
	DeleteTag(userId uint, tagId uint) error
}

type tagRepository struct {
	db *gorm.DB
}

func NewTagRepository(db *gorm.DB) ITagRepository {
	return &tagRepository{db}
}

func (tr tagRepository) GetAllTags(tags *[]model.Tag, userId uint) error {
	//user idを元にタグを全て取得
	if err := tr.db.Joins("User").Where("user_id=?", userId).Order("created_at").Find(tags).Error; err != nil {
		return err
	}
	return nil
}

func (tr tagRepository) CreateTag(tag *model.Tag) error {
	if err := tr.db.Create(tag).Error; err != nil {
		return err
	}
	return nil
}

func (tr tagRepository) UpdateTag(tag *model.Tag, userId uint, tagId uint) error {
	result := tr.db.Clauses(clause.Returning{}).Where("id=? AND user_id=?", tagId, userId).Updates(model.Tag{Name: tag.Name, Color: tag.Color})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected > 1 {
		return fmt.Errorf("object does not exist")
	}
	return nil
}

func (tr tagRepository) DeleteTag(userId uint, tagId uint) error {
	result := tr.db.Where("id=? AND user_id=?", tagId, userId).Delete(&model.Tag{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected < 1 {
		return fmt.Errorf("object does not exist")
	}
	return nil
}