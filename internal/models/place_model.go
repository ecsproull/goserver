package models

import (
	"time"
)

type DbPlace struct {
	ID            int        `json:"id" db:"id"`
	PlaceName     string     `json:"place_name" db:"place_name"`
	PlaceInfo     string     `json:"place_info" db:"place_info"`
	PlaceLat      float64    `json:"place_lat" db:"place_lat"`
	PlaceLng      float64    `json:"place_lng" db:"place_lng"`
	PlaceIconType int        `json:"place_icon_type" db:"place_icon_type"`
	PlaceAddress  string     `json:"place_address" db:"place_address"`
	PlacePhone    string     `json:"place_phone" db:"place_phone"`
	PlaceEmail    string     `json:"place_email" db:"place_email"`
	PlaceWebsite  string     `json:"place_website" db:"place_website"`
	PlaceArrive   *time.Time `json:"place_arrive" db:"place_arrive"`
	PlaceDepart   *time.Time `json:"place_depart" db:"place_depart"`
	PlaceHideInfo BoolInt    `json:"place_hide_info" db:"place_hide_info"`
	CreatedAt     time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt     time.Time  `json:"updated_at" db:"updated_at"`
}

type PQueries struct {
	GetAll string
	Insert string
	Update string
	Delete string
}

var PlaceQueries = PQueries{
	GetAll: `
      SELECT id, place_name, place_info, place_lat, place_lng, place_icon_type,
               place_address, place_phone, place_email, place_website, 
               place_arrive, place_depart, place_hide_info
      FROM places 
    `,
	Insert: `
      INSERT INTO places (place_name, place_info, place_lat, place_lng, place_icon_type,
                            place_address, place_phone, place_email, place_website, 
                            place_arrive, place_depart, place_hide_info) 
      VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)
    `,
	Update: `
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
    `,
	Delete: `
      DELETE FROM places 
      WHERE id = $1
    `,
}
