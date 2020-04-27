package models

import (
	"github.com/jinzhu/gorm"
)

// Movie model
type Movie struct {
	ID int `gorm:"AUTO_INCREMENT;not_null"`
	Name string `gorm:"type:varchar(75)"             json:name`
	Runtime string `gorm:"type:varchar(150)"         json:runtime`
	Released string `gorm:"type:varchar(75)"         json:released`
	Rated string `gorm:"type:varchar(75)"            json:rated`
	Plot string `gorm:"type:varchar(255)"            json:plot`
	Source string `gorm:"type:varchar(300):not_null" json:source`
}


// saving movie obj to database
func (movie *Movie) SaveMovie(db *gorm.DB) (*Movie, error) {
	var err error

	// Debug the whole insert operation
	err = db.Debug().Create(&movie).Error
	if err != nil {
		return &Movie{}, err
	}

	return movie, nil
}

func GetAllMovies(db *gorm.DB) (*[]Movie, error) {
	movies := []Movie{}
	if err := db.Debug().Table("main.movie").Find(&movies).Error; err != nil {
		return &[]Movie{}, err
	}
	return &movies, nil
}