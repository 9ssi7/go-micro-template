package mapper

import (
	"time"

	"github.com/ssibrahimbas/micro-template/src/dto"
	"github.com/ssibrahimbas/micro-template/src/entity"
	"github.com/ssibrahimbas/micro-template/src/event"
)

func (m *Mapper) MapSomeEntityToCreatedEvent(s *entity.Some) *event.SomeCreated {
	return &event.SomeCreated{
		ID:   s.ID,
		Name: s.Name,
	}
}

func (m *Mapper) MapSomeEntityToDeletedEvent(s *entity.Some) *event.SomeDeleted {
	return &event.SomeDeleted{
		ID:   s.ID,
		Name: s.Name,
	}
}

func (m *Mapper) MapSomeEntityToCreatedDto(s *entity.Some) *dto.SomeCreated {
	return &dto.SomeCreated{
		ID:   s.ID,
		Name: s.Name,
	}
}

func (m *Mapper) MapSomeEntityToFoundDto(s *entity.Some) *dto.SomeFound {
	return &dto.SomeFound{
		ID:        s.ID,
		Name:      s.Name,
		IsActive:  s.IsActive,
		CreatedAt: s.CreatedAt,
		UpdatedAt: s.UpdatedAt,
	}
}

func (m *Mapper) MapSomeCreateDtoToEntity(d *dto.SomeCreate) *entity.Some {
	t := time.Now()
	return &entity.Some{
		ID:        "",
		Name:      d.Name,
		IsActive:  true,
		CreatedAt: t,
		UpdatedAt: t,
	}
}
