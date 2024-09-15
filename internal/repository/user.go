package repository

import (
	"context"
	"github.com/mazharul-islam/cacher"
	"github.com/mazharul-islam/config"
	"github.com/mazharul-islam/internal/entity"
	"github.com/mazharul-islam/utils"
	"github.com/pilagod/gorm-cursor-paginator/v2/paginator"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type UserRepository struct {
	db    *gorm.DB
	cache cacher.CacheManager
}

func NewUserRepository(
	db *gorm.DB,
	cache cacher.CacheManager,
) entity.IUserRepository {
	return &UserRepository{
		db:    db,
		cache: cache,
	}
}

func (repo *UserRepository) GetUserByID(ctx context.Context, id uint) (*entity.Users, error) {
	logger := logrus.WithContext(ctx).WithFields(logrus.Fields{
		"context":    utils.DumpIncomingContext(ctx),
		"customerID": id,
	})

	cacheKey := cacher.GetUserCacheKeyByID(id)
	if config.EnableCaching() {
		cachedItem, mutex, err := cacher.FindFromCacheByKey[*entity.Users](repo.cache, cacheKey)
		if err != nil {
			logger.Error(err)
			return nil, err
		}

		defer cacher.SafeUnlock(mutex)

		if mutex == nil {
			logger.WithField("cacheKey", cacheKey).Info("returning customer from redis cache")
			return cachedItem, nil
		}
	}

	user := &entity.Users{}
	err := repo.db.WithContext(ctx).Take(user, "id = ?", id).Error
	switch err {
	case nil:
		if err := repo.cache.StoreWithoutBlocking(cacher.NewItem(cacheKey, utils.Dump(user))); err != nil {
			logger.Error(err)
		}

		return user, nil
	case gorm.ErrRecordNotFound:
		if err := repo.cache.StoreNil(cacheKey); err != nil {
			logger.Error(err)
		}
		return nil, nil
	default:
		return nil, err
	}
}

func (repo *UserRepository) GetUserByCriteria(ctx context.Context, request entity.RequestFilterUsers) (users []entity.Users, count int64, cursor paginator.Cursor, err error) {
	logger := logrus.WithContext(ctx).WithFields(logrus.Fields{
		"ctx":            utils.DumpIncomingContext(ctx),
		"searchCriteria": utils.Dump(request),
	})

	scopes := repo.buildFilterScopeByCriteria(request)

	count, err = repo.countAll(ctx, scopes, request)
	if err != nil {
		logger.Error(err)
		return
	}

	if count <= 0 {
		return
	}

	page := repo.createPaginator(request)

	result, cursor, err := page.Paginate(repo.db.WithContext(ctx).Scopes(scopes...).Select("*"), &users)
	if err != nil {
		logger.Error(err)
		return
	}

	if err = result.Error; err != nil {
		logger.Error(err)
		return
	}

	return users, count, cursor, nil
}

func (repo *UserRepository) buildFilterScopeByCriteria(request entity.RequestFilterUsers) []func(db *gorm.DB) *gorm.DB {
	var scopes []func(db *gorm.DB) *gorm.DB

	if request.Name != "" {
		scopes = append(scopes, filterByName(request.Name))
	}

	if request.Gender != "" {
		scopes = append(scopes, filterByGender(request.Gender))
	}

	if request.Age != nil {
		scopes = append(scopes, filterByAgeRange(request.Age))
	}

	return scopes
}

func (repo *UserRepository) countAll(ctx context.Context, scopes []func(*gorm.DB) *gorm.DB, criteria entity.RequestFilterUsers) (int64, error) {
	var count int64

	if err := repo.db.WithContext(ctx).Model(entity.Users{}).
		Scopes(scopes...).
		Count(&count).
		Error; err != nil {
		logrus.WithContext(ctx).WithFields(logrus.Fields{
			"context":  utils.DumpIncomingContext(ctx),
			"criteria": utils.Dump(criteria),
		}).Error(err)
		return 0, err
	}

	return count, nil
}

func (repo *UserRepository) createPaginator(searchCriteria entity.RequestFilterUsers) *paginator.Paginator {
	opts := []paginator.Option{
		&paginator.Config{
			Keys:  []string{"ID"},
			Limit: 10,
			Order: paginator.DESC,
		},
	}

	if searchCriteria.Size > 0 {
		opts = append(opts, paginator.WithLimit(int(searchCriteria.Size)))
	}

	switch searchCriteria.SortDir {
	case "desc":
		opts = append(opts, paginator.WithOrder(paginator.DESC))
	case "asc":
		opts = append(opts, paginator.WithOrder(paginator.ASC))
	}

	switch searchCriteria.CursorDir {
	case entity.CursorDirectionPrev:
		opts = append(opts, paginator.WithBefore(searchCriteria.Cursor))
	case entity.CursorDirectionNext:
		opts = append(opts, paginator.WithAfter(searchCriteria.Cursor))
	}

	return paginator.New(opts...)
}
