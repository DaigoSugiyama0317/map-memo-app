package repository

import "github.com/DaigoSugiyama0317/map-memo-app/model"

type ILocationRepository interface {
	GetAllLocations(locations *[]model.Location, userId uint) error
	GetLocationById(locations *[]model.Location, userId uint, locationId uint) error
	GetLocationByTag(locations *[]model.Location, userId uint, tagId uint) error
	CreateLocation(location *model.Location) error
	UpdateLocation(location *model.Location, userId uint, locationId uint) error
	DeleteLocation(location *model.Location) error
}