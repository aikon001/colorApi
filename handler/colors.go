package handler

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"

	"github.com/aikon001/colorapiserver/db"
	"github.com/aikon001/colorapiserver/models"
)

var colorIDKey = "colorID"

func colors(router chi.Router) {
	router.Get("/", db.GetAllColors)
	router.Post("/", createColor)
	router.Route("/{colorID}", func(router chi.Router) {
		router.Use(ItemContext)
		router.Get("/", getColor)
		router.Put("/", updateColor)
		router.Delete("/", deleteColor)
	})
}
func ColorContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		itemId := chi.URLParam(r, "itemId")
		if itemId == "" {
			render.Render(w, r, ErrorRenderer(fmt.Errorf("item ID is required")))
			return
		}
		id, err := strconv.Atoi(itemId)
		if err != nil {
			render.Render(w, r, ErrorRenderer(fmt.Errorf("invalid item ID")))
		}
		ctx := context.WithValue(r.Context(), colorIDKey, id)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func createColor(w http.ResponseWriter, r *http.Request) {
	color := &models.Color{}
	if err := render.Bind(r, color); err != nil {
		render.Render(w, r, ErrBadRequest)
		return
	}
	if err := dbInstance.AddColor(color); err != nil {
		render.Render(w, r, ErrorRenderer(err))
		return
	}
	if err := render.Render(w, r, color); err != nil {
		render.Render(w, r, ServerErrorRenderer(err))
		return
	}
}

func getAllColors(w http.ResponseWriter, r *http.Request) {
	colors, err := dbInstance.GetAllItems()
	if err != nil {
		render.Render(w, r, ServerErrorRenderer(err))
		return
	}
	if err := render.Render(w, r, colors); err != nil {
		render.Render(w, r, ErrorRenderer(err))
	}
}

func getColor(w http.ResponseWriter, r *http.Request) {
	colorID := r.Context().Value(colorDKey).(int)
	color, err := dbInstance.GetItemById(itemID)
	if err != nil {
		if err == db.ErrNoMatch {
			render.Render(w, r, ErrNotFound)
		} else {
			render.Render(w, r, ErrorRenderer(err))
		}
		return
	}
	if err := render.Render(w, r, &color); err != nil {
		render.Render(w, r, ServerErrorRenderer(err))
		return
	}
}

func deleteItem(w http.ResponseWriter, r *http.Request) {
	itemId := r.Context().Value(itemIDKey).(int)
	err := dbInstance.DeleteItem(itemId)
	if err != nil {
		if err == db.ErrNoMatch {
			render.Render(w, r, ErrNotFound)
		} else {
			render.Render(w, r, ServerErrorRenderer(err))
		}
		return
	}
}
func updateItem(w http.ResponseWriter, r *http.Request) {
	itemId := r.Context().Value(itemIDKey).(int)
	itemData := models.Item{}
	if err := render.Bind(r, &itemData); err != nil {
		render.Render(w, r, ErrBadRequest)
		return
	}
	item, err := dbInstance.UpdateItem(itemId, itemData)
	if err != nil {
		if err == db.ErrNoMatch {
			render.Render(w, r, ErrNotFound)
		} else {
			render.Render(w, r, ServerErrorRenderer(err))
		}
		return
	}
	if err := render.Render(w, r, &item); err != nil {
		render.Render(w, r, ServerErrorRenderer(err))
		return
	}
}
