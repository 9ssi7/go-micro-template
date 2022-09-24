package internal

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ssibrahimbas/micro-template/src/dto"
	"github.com/ssibrahimbas/ssi-core/pkg/result"
)

func (h *Handler) Create(c *fiber.Ctx) error {
	l, a := h.i.GetLanguagesInContext(c)
	d := &dto.SomeCreate{}
	h.parseBody(c, d)
	r := h.s.Create(d)
	return result.SuccessData(h.i.Translate("some_created", l, a), r, fiber.StatusCreated)
}

func (h *Handler) Find(c *fiber.Ctx) error {
	l, a := h.i.GetLanguagesInContext(c)
	d := &dto.SomeFind{
		Name: c.Query("name"),
	}
	r := h.s.Find(d)
	return result.SuccessData(h.i.Translate("some_found", l, a), r, fiber.StatusOK)
}

func (h *Handler) parseBody(c *fiber.Ctx, d interface{}) {
	l, a := h.i.GetLanguagesInContext(c)
	if err := c.BodyParser(d); err != nil {
		panic(result.Error(h.i.Translate("invalid_request", l, a), fiber.StatusBadRequest))
	}
	h.validateStruct(d, l, a)
}

func (h *Handler) validateStruct(d interface{}, l, a string) {
	if errors := h.v.ValidateStruct(d, l, a); len(errors) > 0 {
		panic(result.ErrorData(h.i.Translate("invalid_request", l, a), errors, fiber.StatusBadRequest))
	}
}
