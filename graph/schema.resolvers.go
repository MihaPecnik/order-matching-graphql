package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"github.com/MihaPecnik/order-matching-graphql/graph/generated"
	"github.com/MihaPecnik/order-matching-graphql/graph/model"
	"github.com/MihaPecnik/order-matching-graphql/internal/orderbook"
)

func (r *mutationResolver) UpdateOrderBook(ctx context.Context, input model.Request) ([]*model.UpdateOrderBookResponse, error) {
	order := orderbook.Order{
		UserId:   input.UserID,
		Value:    input.Value,
		Buy:      input.Buy,
		Ticker:   input.Ticker,
		Quantity: input.Quantity,
	}
	response, err := order.Update()
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (r *queryResolver) GetInfoTicker(ctx context.Context, input model.Ticker) (*model.GetTickerInfoResponse, error) {
	ticker := orderbook.Ticker{
		Ticker: input.Ticker,
	}
	response, err := ticker.GetInfo()
	if err != nil {
		return nil, err
	}
	return response, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.

