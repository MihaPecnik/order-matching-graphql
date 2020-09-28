package orderbook

import (
	"github.com/MihaPecnik/order-matching-graphql/graph/model"
	"github.com/MihaPecnik/order-matching-graphql/internal/pkg/db/postgres"
)

type Ticker struct {
	Ticker   string  `json:"ticker"`
}

func (ticker Ticker) GetInfo() (*model.GetTickerInfoResponse,error) {
	var sell model.UpdateOrderBookResponse
	err := postgres.Db.Table("tables").
		Select("value", "quantity").
		Where("ticker = ? AND buy = ?", ticker.Ticker, false).
		Order("value asc").
		First(&sell).Error
	if err != nil {
		return &model.GetTickerInfoResponse{}, err
	}

	var buy model.UpdateOrderBookResponse
	err = postgres.Db.Table("tables").
		Select("value", "quantity").
		Where("ticker = ? AND buy = ?", ticker.Ticker, true).
		Order("value desc").
		First(&buy).Error
	if err != nil {
		return &model.GetTickerInfoResponse{}, err
	}

	return &model.GetTickerInfoResponse{
		Buy: &buy,
		Sell: &sell,
	},nil

}