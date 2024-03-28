package utils

import (
	"sort"
	"vk-test/pkg/database/mongodb/models"
)

func OrderByPrice(ads []models.Ad, asc bool) []models.Ad {
	if asc {
		sort.SliceStable(ads, func(i, j int) bool {
			return ads[i].Price < ads[j].Price
		})
	} else {
		sort.SliceStable(ads, func(i, j int) bool {
			return ads[i].Price > ads[j].Price
		})
	}
	return ads
}

func OrderByDate(ads []models.Ad, asc bool) []models.Ad {
	if asc {
		sort.SliceStable(ads, func(i, j int) bool {
			return ads[i].CreateAt.Before(ads[j].CreateAt)
		})
	} else {
		sort.SliceStable(ads, func(i, j int) bool {
			return ads[i].CreateAt.After(ads[j].CreateAt)
		})
	}
	return ads
}
