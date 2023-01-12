package data

import (
	"nosql/internal/validator"
	"time"
)

type Movie struct {
	ID        int64     // Unique integer ID for the movie
	CreatedAt time.Time // Timestamp for when the movie is added to our database
	Title     string    // Movie title
	Year      int32     // Movie release year
	Runtime   int32     // Movie runtime (in minutes)
	Genres    []string  // Slice of genres for the movie (romance, comedy, etc.)
	Version   int32     // The version number starts at 1 and will be incremented each
	// time the movie information is updated
}

func ValidateMovie(v *validator.Validator, input *Item) {
	v.Check(input.Description != "", "description", "must be provided")
	v.Check(input.Price != 0, "price", "must be provided")
	v.Check(input.Name != "", "name", "must be provided")
	v.Check(input.Category != "", "category", "must be provided")
	v.Check(input.Price > 0, "price", "must be greater than 0")
	v.Check(input.MarketID > 0, "market id", "must be greater than 0")
}

type Market struct {
	ID       string `json:"id"`
	Address  string `json:"address"`
	WorkTime string `json:"worktime"`
	Name     string `json:"name"`
	AdminID  int64  `json:"adminid"`
}

type Item struct {
	ID          string `json:"_id"`
	Description string `json:"description"`
	MarketID    int64  `json:"marketid"`
	Price       int    `json:"price"`
	ExpiredDate string `json:"expireddate"`
	Name        string `json:"name"`
	Category    string `json:"category"`
	//LocationMarket any      `json:"locationmarket"`
}
type Admin struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}
type Category struct {
	ID   string  `json:"id"`
	Name string `json:"name"`
}
type Cart struct {
	ID    int64  `json:"id"`
	Items []Item `json:"items"`
}
