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

	query := `
        SELECT id, place_name, place_info, place_lat, place_lng, place_icon_type,
               place_address, place_phone, place_email, place_website, 
               place_arrive, place_depart, place_hide_info
        FROM places
    `

	err := database.DB.Select(&places, query)
	if err != nil {
		return nil, err
	}
	return places, nil
}

// Save a new place
func SavePlace(place *models.DbPlace) error {
	query := `
        INSERT INTO places (place_name, place_info, place_lat, place_lng, place_icon_type,
                            place_address, place_phone, place_email, place_website, 
                            place_arrive, place_depart, place_hide_info) 
        VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12) 
    `

	_, err := database.DB.Exec(query,
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
	query := `
        UPDATE places
        SET place_name = $1,
            place_info = $2,
            place_lat = $3,
            place_lng = $4,
            place_icon_type = $5,
            place_address = $6,
            place_phone = $7,
            place_email = $8,
            place_website = $9,
            place_arrive = $10,
            place_depart = $11,
            place_hide_info = $12,
            updated_at = CURRENT_TIMESTAMP
        WHERE id = $13
        RETURNING id, place_name, place_info, place_lat, place_lng, place_icon_type,
                  place_address, place_phone, place_email, place_website,
                  place_arrive, place_depart, place_hide_info,
                  created_at, updated_at
    `

	result, err := database.DB.Exec(query,
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

	query := `
        DELETE FROM places 
        WHERE id = $1
        RETURNING id, place_name, place_info, place_lat, place_lng, place_icon_type,
                  place_address, place_phone, place_email, place_website, 
                  place_arrive, place_depart, place_hide_info, 
                  created_at, updated_at
    `

	var deleted models.DbPlace
	err = database.DB.QueryRowx(query, placeID).Scan(
		&deleted.ID,
		&deleted.PlaceName,
		&deleted.PlaceInfo,
		&deleted.PlaceLat,
		&deleted.PlaceLng,
		&deleted.PlaceIconType,
		&deleted.PlaceAddress,
		&deleted.PlacePhone,
		&deleted.PlaceEmail,
		&deleted.PlaceWebsite,
		&deleted.PlaceArrive,
		&deleted.PlaceDepart,
		&deleted.PlaceHideInfo,
		&deleted.CreatedAt,
		&deleted.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}
	return &deleted, nil
}
