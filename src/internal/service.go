package internal

import (
	"github.com/ssibrahimbas/micro-template/src/dto"
	"github.com/ssibrahimbas/micro-template/src/event_publisher"
	"github.com/ssibrahimbas/micro-template/src/mapper"
	"github.com/ssibrahimbas/ssi-core/pkg/i18n"
)

type Service struct {
	r  *Repo
	i  *i18n.I18n
	m  *mapper.Mapper
	ep *event_publisher.Publisher
}

type ServiceConfig struct {
	Repo        *Repo
	I18n        *i18n.I18n
	EvPublisher *event_publisher.Publisher
}

func NewService(c *ServiceConfig) *Service {
	return &Service{
		r:  c.Repo,
		i:  c.I18n,
		m:  mapper.New(),
		ep: c.EvPublisher,
	}
}

func (s *Service) Create(d *dto.SomeCreate) *dto.SomeCreated {
	e := s.m.MapSomeCreateDtoToEntity(d)
	e = s.r.Create(e)
	s.ep.PublishSomeCreated(s.m.MapSomeToCreatedEvent(e))
	return s.m.MapSomeToCreatedDto(e)
}

func (s *Service) Find(d *dto.SomeFind) []*dto.SomeFound {
	some := s.r.Fetch(d)
	found := []*dto.SomeFound{}
	for _, so := range some {
		found = append(found, s.m.MapSomeToFoundDto(so))
	}
	return found
}
