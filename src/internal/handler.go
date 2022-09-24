package internal

import (
	"github.com/ssibrahimbas/ssi-core/pkg/http"
	"github.com/ssibrahimbas/ssi-core/pkg/i18n"
	"github.com/ssibrahimbas/ssi-core/pkg/validator"
)

type Handler struct {
	s *Service
	i *i18n.I18n
	v *validator.Validator
	h *http.Client
}

type HandlerConfig struct {
	Service    *Service
	HttpClient *http.Client
	Validator  *validator.Validator
	I18n       *i18n.I18n
}

func NewHandler(c *HandlerConfig) *Handler {
	return &Handler{
		s: c.Service,
		h: c.HttpClient,
		v: c.Validator,
		i: c.I18n,
	}
}

func (h *Handler) InitAllVersions() {
	h.initV1()
}

func (h *Handler) initV1() {
	v1 := h.h.App.Group("/api/v1")
	v1.Post("/some", h.Create)
	v1.Get("/some", h.Find)
}
