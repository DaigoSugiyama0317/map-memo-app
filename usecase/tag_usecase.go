package usecase

import (
	"github.com/DaigoSugiyama0317/map-memo-app/model"
	"github.com/DaigoSugiyama0317/map-memo-app/repository"
)

type ITagUsecase interface {
	GetAllTags(userId uint) ([]model.TagResponse, error)
	CreateTag(tag model.Tag) (model.TagResponse, error)
	UpdateTag(tag model.Tag, userId uint, tagId uint) (model.TagResponse, error)
	DeleteTag(userId uint, tagId uint) error
}

type tagUsecase struct {
	tr repository.ITagRepository
}

func NewTagUsecase(tr repository.ITagRepository) ITagUsecase {
	return &tagUsecase{tr}
}

func (tu tagUsecase) GetAllTags(userId uint) ([]model.TagResponse, error) {
	tags := []model.Tag{}
	if err := tu.tr.GetAllTags(&tags, userId); err != nil {
		return nil, err
	}
	resTags := []model.TagResponse{}
	for _, v := range tags {
		t := model.TagResponse{
			ID: v.ID,
			Name: v.Name,
			Color: v.Color,
		}
		resTags = append(resTags, t)
	}
	return resTags, nil
}

func (tu tagUsecase) CreateTag(tag model.Tag) (model.TagResponse, error) {
	if err := tu.tr.CreateTag(&tag); err != nil {
		return model.TagResponse{}, err
	}
	resTag := model.TagResponse{
		ID: tag.ID,
		Name: tag.Name,
		Color: tag.Color,
	}
	return resTag, nil
}

func (tu tagUsecase) UpdateTag(tag model.Tag, userId uint, tagId uint) (model.TagResponse, error) {
	if err := tu.tr.UpdateTag(&tag, userId, tagId); err != nil {
		return model.TagResponse{}, err
	}
	resTag := model.TagResponse{
		ID: tag.ID,
		Name: tag.Name,
		Color: tag.Color,
	}
	return resTag, nil
}

func (tu tagUsecase) DeleteTag(userId uint, tagId uint) error {
	if err := tu.tr.DeleteTag(userId, tagId); err != nil {
		return err
	}
	return nil
}
