package entity

import "context"

type (
	IMatchService interface {
		GetListRecommendations(c context.Context, id uint, requestFilter RequestFilterUsers) ([]Users, CursorInfo, error)
	}

	Preferences struct {
		MaxDistanceKm     int    `json:"max_distance_km"`
		PreferredGender   string `json:"preferred_gender"`
		PreferredAgeRange []int  `json:"preferred_age_range"`
	}
)
