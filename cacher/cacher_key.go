package cacher

import "github.com/mazharul-islam/utils"

func GetCustomerCacheKeyByID(customerID uint) string {
	return createCacheKey(utils.WriteStringTemplate("cache:object:customer:id:%d", customerID))
}

func GetUserCacheKeyByID(id uint) string {
	return createCacheKey(utils.WriteStringTemplate("cache:object:user:id:%d", id))
}
