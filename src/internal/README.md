## Internal Folder

This folder contains the code for the internal microservice. Under this folder are the main codes of our porject. There are 4 main files that we will use here.

These:

### `repo.go`

This is where we run our database operations.

#### Simple Type

```go
type Repo struct {
	c   *mongo.Collection
	ctx context.Context
}
```

#### The New Function

```go
type RepoConfig struct {
	MongoDB *db.MongoDB
}

func NewRepo(c *RepoConfig) *Repo {
	return &Repo{
		ctx: context.TODO(),
		c:   c.MongoDB.GetCollection("some"),
	}
}
```

#### Simple Action

```go
func (r *Repo) Create(e *entity.Some) *entity.Some {
	res, err := r.c.InsertOne(r.ctx, e)
	helper.CheckErr(err)
	e.ID = res.InsertedID.(primitive.ObjectID).Hex()
	return e
}
```

### `service.go`

We transform the incoming data with mapper and send it to the `repo`. If we are going to publish the process with other microservices afterwards, we send the data saved in eventPublisher by transforming it into an event. Finally, we finish the request by transforming the recorded data into a response object.

#### Simple Type

```go
type Service struct {
	r  *Repo
	i  *i18n.I18n
	m  *mapper.Mapper
	ep *event_publisher.Publisher
}
```

#### The New Function

```go
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
```

#### Simple Action

```go
func (s *Service) Create(d *dto.SomeCreate) *dto.SomeCreated {
	e := s.m.MapSomeCreateDtoToEntity(d)
	e = s.r.Create(e)
	s.ep.PublishSomeCreated(s.m.MapSomeToCreatedEvent(e))
	return s.m.MapSomeToCreatedDto(e)
}
```

### `handler.go`

This is where the http implementation, middleware usage, and routes are defined. This is where we handle the incoming requests. We get the data from the request, send it to the `service` and return the response.

#### Simple Type

```go
type Handler struct {
	s *Service
	i *i18n.I18n
	v *validator.Validator
	h *http.Client
}
```

#### The New Function

```go
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
```

#### The Init Functions

```go
func (h *Handler) InitAllVersions() {
	h.initV1()
}

func (h *Handler) initV1() {
	v1 := h.h.App.Group("/api/v1")
	v1.Post("/some", h.Create)
	v1.Get("/some", h.Find)
}
```

### `api.go`

Contains functions of the Handler type. It is in the role of http controller and is responsible for operations such as request's conversions to dto, validations.

#### Simple Function

```go
func (h *Handler) Create(c *fiber.Ctx) error {
	l, a := h.i.GetLanguagesInContext(c)
	d := &dto.SomeCreate{}
	h.parseBody(c, d)
	r := h.s.Create(d)
	return result.SuccessData(h.i.Translate("some_created", l, a), r, fiber.StatusCreated)
}
```