package cmd

import (
	"github.com/mazharul-islam/cacher"
	"github.com/mazharul-islam/internal/entity"
	"github.com/mazharul-islam/internal/repository"
	"github.com/mazharul-islam/internal/service"
	"gorm.io/gorm"
)

func InitMatchService(db *gorm.DB, cacher cacher.CacheManager) entity.IMatchService {
	userRepository := repository.NewUserRepository(db, cacher)
	matchService := service.NewMatchService(userRepository)

	return matchService
}
