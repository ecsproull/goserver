package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Place struct {
	ID            primitive.ObjectID `bson:"_id" json:"_id"`
	PlaceName     string             `bson:"place_name" json:"place_name"`
	PlaceInfo     string             `bson:"place_info" json:"place_info"`
	PlaceLat      float64            `bson:"place_lat" json:"place_lat"`
	PlaceLng      float64            `bson:"place_lng" json:"place_lng"`
	PlaceIconType int                `bson:"place_icon_type" json:"place_icon_type"`
	PlaceAddress  string             `bson:"place_address" json:"place_address"`
	PlacePhone    string             `bson:"place_phone" json:"place_phone"`
	PlaceWebsite  string             `bson:"place_website" json:"place_website"`
	PlaceArrive   *time.Time         `bson:"place_arrive" json:"place_arrive"`
	PlaceDepart   *time.Time         `bson:"place_depart" json:"place_depart"`
	PlaceHideInfo BoolInt            `bson:"place_hide_info" json:"place_hide_info"`
	PlaceEmail    string             `bson:"place_email" json:"place_email"`
}
