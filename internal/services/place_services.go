package services

import (
	"database/sql"
	"goserver/internal/database"
	"goserver/internal/models"
	"strconv"
)

// Get all places
func GetPlaces() ([]models.DbPlace, error) {
	var places []models.DbPlace
	err := database.DB.Select(&places, models.PlaceQueries.GetAll)
	if err != nil {
		return nil, err
	}
	return places, nil
}

// Save a new place
func SavePlace(place *models.DbPlace) error {
	_, err := database.DB.Exec(models.PlaceQueries.Insert,
		place.PlaceName,
		place.PlaceInfo,
		place.PlaceLat,
		place.PlaceLng,
		place.PlaceIconType,
		place.PlaceAddress,
		place.PlacePhone,
		place.PlaceEmail,
		place.PlaceWebsite,
		place.PlaceArrive,
		place.PlaceDepart,
		place.PlaceHideInfo,
	)
	return err
}

// Update a place
func UpdatePlace(place *models.DbPlace) (*sql.Result, error) {
	result, err := database.DB.Exec(models.PlaceQueries.Update,
		place.PlaceName,
		place.PlaceInfo,
		place.PlaceLat,
		place.PlaceLng,
		place.PlaceIconType,
		place.PlaceAddress,
		place.PlacePhone,
		place.PlaceEmail,
		place.PlaceWebsite,
		place.PlaceArrive,
		place.PlaceDepart,
		place.PlaceHideInfo,
		place.ID,
	)

	return &result, err
}

// Delete a place
func DeletePlace(id string) (*models.DbPlace, error) {
	placeID, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}

	var deleted models.DbPlace
	_, err = database.DB.Exec(models.PlaceQueries.Delete, placeID)

	if err != nil {
		return nil, err
	}
	return &deleted, nil
}
