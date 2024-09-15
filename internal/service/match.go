package service

import (
	"context"
	"github.com/mazharul-islam/internal/entity"
	"github.com/mazharul-islam/utils"
	"github.com/sirupsen/logrus"
)

type MatchService struct {
	userRepository entity.IUserRepository
}

func NewMatchService(userRepository entity.IUserRepository) entity.IMatchService {
	return &MatchService{
		userRepository: userRepository,
	}
}

func (service *MatchService) GetListRecommendations(ctx context.Context, id uint, requestFilter entity.RequestFilterUsers) ([]entity.Users, entity.CursorInfo, error) {
	logger := logrus.WithContext(ctx).WithFields(logrus.Fields{
		"ctx": utils.DumpIncomingContext(ctx),
	})

	var (
		preferences entity.Preferences
	)

	//get user by current user_id
	existUser, err := service.userRepository.GetUserByID(ctx, id)
	if err != nil {
		logger.Error(err)
	}

	if err := utils.JSONUnmarshal([]byte(existUser.Preferences), &preferences); err != nil {
		logger.Error(err)
	}

	//Preferences Filtering: Match users based on gender, age range, and location
	//distance
	requestFilter.Gender = preferences.PreferredGender
	requestFilter.Age = preferences.PreferredAgeRange

	recommendationUsers, totalItems, cursor, err := service.userRepository.GetUserByCriteria(ctx, requestFilter)

	if err != nil {
		logger.Error(err)
	}

	//Mutual Interests: Rank users higher if they share more common interests.

	return recommendationUsers, requestFilter.ToCursorInfo(cursor, totalItems), nil
}
