package internal

import (
	"context"

	"github.com/ssibrahimbas/micro-template/src/dto"
	"github.com/ssibrahimbas/micro-template/src/entity"
	"github.com/ssibrahimbas/ssi-core/pkg/db"
	"github.com/ssibrahimbas/ssi-core/pkg/helper"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repo struct {
	c   *mongo.Collection
	ctx context.Context
}

func NewRepo(d *db.MongoDB) *Repo {
	return &Repo{
		ctx: context.TODO(),
		c:   d.GetCollection("some"),
	}
}

func (r *Repo) Create(e *entity.Some) *entity.Some {
	res, err := r.c.InsertOne(r.ctx, e)
	helper.CheckErr(err)
	e.ID = res.InsertedID.(primitive.ObjectID).Hex()
	return e
}

func (r *Repo) Fetch(d *dto.SomeFind) []*entity.Some {
	var s []*entity.Some
	m := bson.M{}
	if d.Name != "" {
		m["name"] = bson.M{
			"$regex": d.Name,
		}
	}
	cur, err := r.c.Find(r.ctx, m)
	helper.CheckErr(err)
	err = cur.All(r.ctx, &s)
	helper.CheckErr(err)
	return s
}
