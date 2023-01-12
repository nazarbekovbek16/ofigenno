package main

import (
	"fmt"
	"net/http"
	"nosql/internal/data"
	"nosql/internal/validator"
)

func (app *Application) createItemHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Description string `json:"description"`
		MarketID    int64  `json:"marketid"`
		Price       int    `json:"price"`
		ExpiredDate string `json:"expireddate"`
		Name        string `json:"name"`
		Category    string `json:"category"`
		//LocationMarket any      `json:"locationmarket"`
	}
	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}
	item := &data.Item{
		Description: input.Description,
		MarketID:    input.MarketID,
		Price:       input.Price,
		ExpiredDate: input.ExpiredDate,
		Name:        input.Name,
		Category:    input.Category,
	}
	v := validator.New()
	if data.ValidateMovie(v, item); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	err = app.Insert(item)
	if err != nil {
		app.dbError(w, r, err)
	}
	fmt.Fprintf(w, "%+v\n", input)
}

func (app *Application) showItemHandler(w http.ResponseWriter, r *http.Request) {
	id := app.readIDParam(r)
	item, err := app.Get(id)
	if err != nil {
		app.dbError(w, r, err)
		return
	}
	err = app.writeJSON(w, http.StatusOK, envelope{"item": item}, nil)
	if err != nil {
		return
	}
}

func (app *Application) allItemsHandler(w http.ResponseWriter, r *http.Request) {

	// item, err := app.GetAllItems()
	// if err != nil {
	// 	app.dbError(w, r, err)
	// 	return
	// }
	// fmt.Println(item)
	// err = app.writeJSON(w, http.StatusOK, envelope{"item": item}, nil)
	// if err != nil {
	// 	return
	// }
}
