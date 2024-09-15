package repository

import (
	"gorm.io/gorm"
)

// Define a scope function for the complex query
func WithCurrentUser(targetUserID uint) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		query := `
       WITH current_users AS (
    SELECT location, interests
    FROM users
    WHERE id = 1 
),
common_interests AS (
    SELECT u.id, u.name, u.age, u.gender, u.location, u.interests,
           array_length(array(SELECT unnest(u.interests) INTERSECT SELECT unnest(cu.interests)), 1) AS common_interests,
           (u.location <-> cu.location) AS distance_in_degrees
    FROM users u, current_users cu
    WHERE u.id != 1 )
    
SELECT id, name, age, gender, interests, distance_in_degrees
FROM common_interests
WHERE distance_in_degrees <= (50.0 / 111.0) 
ORDER BY common_interests DESC, distance_in_degrees ASC; 
        `

		return db.Raw(query, targetUserID, targetUserID)
	}
}

func filterByName(name string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("name = ?", name)
	}
}

func filterByGender(gender string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("gender = ?", gender)
	}
}

func filterByAgeRange(ageRange []int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("age BETWEEN ? AND ?", ageRange[0], ageRange[1])
	}
}
