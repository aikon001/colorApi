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
	router.Get("/", getAllColors)
	router.Post("/", createColor)
	router.Route("/{colorID}", func(router chi.Router) {
		router.Use(ColorContext)
		router.Get("/", getColor)
		router.Put("/", updateColor)
		router.Delete("/", deleteColor)
	})
}
func ColorContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		colorId := chi.URLParam(r, "colorID")
		if colorId == "" {
			render.Render(w, r, ErrorRenderer(fmt.Errorf("color ID is required")))
			return
		}
		id, err := strconv.Atoi(colorId)
		if err != nil {
			render.Render(w, r, ErrorRenderer(fmt.Errorf("invalid color ID")))
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
	colors, err := dbInstance.GetAllColors()
	if err != nil {
		render.Render(w, r, ServerErrorRenderer(err))
		return
	}
	if err := render.Render(w, r, colors); err != nil {
		render.Render(w, r, ErrorRenderer(err))
	}
}

func getColor(w http.ResponseWriter, r *http.Request) {
	colorID := r.Context().Value(colorIDKey).(int)
	color, err := dbInstance.GetColorById(colorID)
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

func deleteColor(w http.ResponseWriter, r *http.Request) {
	colorId := r.Context().Value(colorIDKey).(int)
	err := dbInstance.DeleteColor(colorId)
	if err != nil {
		if err == db.ErrNoMatch {
			render.Render(w, r, ErrNotFound)
		} else {
			render.Render(w, r, ServerErrorRenderer(err))
		}
		return
	}
}
func updateColor(w http.ResponseWriter, r *http.Request) {
	colorId := r.Context().Value(colorIDKey).(int)
	colorData := models.Color{}
	if err := render.Bind(r, &colorData); err != nil {
		render.Render(w, r, ErrBadRequest)
		return
	}
	color, err := dbInstance.UpdateColor(colorId, colorData)
	if err != nil {
		if err == db.ErrNoMatch {
			render.Render(w, r, ErrNotFound)
		} else {
			render.Render(w, r, ServerErrorRenderer(err))
		}
		return
	}
	if err := render.Render(w, r, &color); err != nil {
		render.Render(w, r, ServerErrorRenderer(err))
		return
	}
}
