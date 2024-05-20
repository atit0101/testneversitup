package itemRePository

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type (
  IntemRepositoryService  interface{}

  itemRepository struct {
    db  *mongo.Client
  }
)

func NewIntemRepository(db *mongo.Client) IntemRepositoryService  {
  return &itemRepository{db} 
}

func (r *itemRepository) intemConn(pctx context.Context)  *mongo.Database  {
  return r.db.Database("item_db") 
}
