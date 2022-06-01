package handler

import (
	"github.com/MStation-io/api/config"
	"github.com/MStation-io/api/repo"
	"github.com/MStation-io/api/repo/pg"
	"github.com/MStation-io/api/services"
)

type Handler struct {
	cfg     config.Config
	repo    *repo.Repo
	store   repo.DBRepo
	service services.Services
}

func NewHandler(cfg config.Config, s repo.DBRepo) *Handler {

	r := pg.NewRepo()
	svc := services.NewServices(cfg, s, r)
	return &Handler{
		cfg:     cfg,
		repo:    r,
		store:   s,
		service: *svc,
	}
}

func NewTestHandler(r *repo.Repo) *Handler {
	h := &Handler{
		store: repo.NewTestStore(),
		repo:  pg.NewRepo(),
	}
	if r != nil {
		h.repo = r
	}
	return h
}
