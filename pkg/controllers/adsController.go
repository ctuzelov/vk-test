package controllers

import (
	"errors"
	"fmt"
	"strconv"
	"time"
	"vk-test/pkg/database/mongodb/models"
	"vk-test/pkg/database/mongodb/repository"
	"vk-test/pkg/utils"

	"github.com/gin-gonic/gin"
)

func CreateAd(c *gin.Context) (err error) {
	var Ad models.Ad

	if err = c.BindJSON(&Ad); err != nil {
		return
	}

	Ad.CreaterEmail = c.GetString("user_email")
	Ad.CreateAt = time.Now()

	if !utils.IsValidPrice(Ad.Price) {
		return fmt.Errorf("invalid price")
	}

	if !utils.IsValidTitle(Ad.Title) {
		return fmt.Errorf("invalid title")
	}

	if !utils.IsValidDescription(Ad.Text) {
		return fmt.Errorf("invalid description")
	}

	if !utils.IsValidImageURL(Ad.ImageURL) {
		return fmt.Errorf("invalid image link")
	}

	// TODO: validation for the form

	// Create the Ad in the database
	err = repository.CreateAd(Ad)
	if err != nil {
		return fmt.Errorf("error occurred while creating Ad: %v", err)
	}

	return
}

// Function that returns all Ads
func GetAllAds(c *gin.Context) ([]models.Ad, error) {
	Ads, err := repository.GetAllAds()

	if err != nil {
		return Ads, err
	}

	return getFilteredIfNeeded(Ads, c)
}

func GetAdsByPage(c *gin.Context) ([]models.Ad, error) {
	var Ads []models.Ad
	if c.Query("page") == "" {
		return Ads, errors.New("page is required")
	}
	page, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		return Ads, err
	}

	Ads, err = repository.GetAdsByPage(page)
	if err != nil {
		return Ads, err
	}

	return getFilteredIfNeeded(Ads, c)
}

func getFilteredIfNeeded(Ads []models.Ad, c *gin.Context) ([]models.Ad, error) {
	var err error

	if c.Query("price_order") != "" {
		if c.Query("price_order") == "asc" {
			Ads = utils.OrderByPrice(Ads, true)
		} else if c.Query("price_order") == "desc" {
			Ads = utils.OrderByPrice(Ads, false)
		}
	}

	if c.Query("date_order") != "" {
		if c.Query("date_order") == "asc" {
			Ads = utils.OrderByDate(Ads, true)
		} else if c.Query("date_order") == "desc" {
			Ads = utils.OrderByDate(Ads, false)
		}
	}

	if c.Query("min_price") != "" {
		fmt.Println(c.Query("min_price"))
		minPrice, err := strconv.ParseFloat(c.Query("min_price"), 64)
		var newAds []models.Ad
		if err == nil {
			for _, ad := range Ads {
				if ad.Price >= minPrice {
					newAds = append(newAds, ad)
				}
			}
		}
		Ads = newAds
	}

	if c.Query("max_price") != "" {
		maxPrice, err := strconv.ParseFloat(c.Query("max_price"), 64)
		var newAds []models.Ad
		if err == nil {
			for _, ad := range Ads {
				if ad.Price <= maxPrice {
					newAds = append(newAds, ad)
				}
			}
		}
		Ads = newAds
	}

	for i := range Ads {
		if Ads[i].CreaterEmail == c.GetString("user_email") {
			Ads[i].Owned = true
		}
	}

	return Ads, err
}
