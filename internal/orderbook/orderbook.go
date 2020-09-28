package orderbook

import (
	"github.com/MihaPecnik/order-matching-graphql/graph/model"
	"github.com/MihaPecnik/order-matching-graphql/internal/pkg/db/postgres"
	"gorm.io/gorm"
	"strings"
)

type Order struct {
	UserId   int     `json:"user_id"`
	Value    float64 `json:"value,string"`
	Quantity int     `json:"quantity"`
	Buy      bool    `json:"buy"`
	Ticker   string  `json:"ticker"`
}

func (order Order) Update() ([]*model.UpdateOrderBookResponse, error) {
	query := `
with ordersUsable as (
  select * , sum(quantity) over (order by value ##order##) as qu
  from tables
  where value ##compare## ? and buy = ? and ticker = ?
)
select *
from (
       select *
       from ordersUsable
       where qu >= ?
       limit 1
     ) as ordersAdditional
union
select *
from ordersUsable
where  qu < ?
order by value ##order## ;
`
	if order.Buy {
		query = strings.Replace(query, "##compare##", "<=", -1)
		query = strings.Replace(query, "##order##", "asc", -1)
	} else {
		query = strings.Replace(query, "##compare##", ">=", -1)
		query = strings.Replace(query, "##order##", "desc", -1)
	}
	stmt := postgres.Db.Session(&gorm.Session{
		PrepareStmt: true,
	})
	response := []*model.UpdateOrderBookResponse{}

	// If any error is return, transaction will take care of a rollback
	err := stmt.Transaction(func(tx *gorm.DB) error {
		var suitableOrders []model.Table

		// Get all orders that are suitable for our request
		err := tx.
			Raw(query, order.Value, !order.Buy, order.Ticker, order.Quantity, order.Quantity).
			Scan(&suitableOrders).Error
		if err != nil {
			return err
		}

		for _, suit := range suitableOrders {
			// If request has bigger quantity, we can delete it (execute the order)
			// If request has smaller quantity, we update it's quantity (partly execute the order)
			if order.Quantity >= suit.Quantity {
				err = tx.Where("id = ?", suit.ID).Delete(&model.Table{}).Error
				if err != nil {
					return err
				}

				order.Quantity -= suit.Quantity
				response = append(response, &model.UpdateOrderBookResponse{
					Value:    suit.Value,
					Quantity: suit.Quantity,
				})
			} else {
				q := suit.Quantity - order.Quantity
				err = tx.Model(&model.Table{}).Where("id = ?", suit.ID).
					Updates(model.Table{
						Quantity: q,
					}).Error
				if err != nil {
					return err
				}

				response = append(response, &model.UpdateOrderBookResponse{
					Value:    suit.Value,
					Quantity: order.Quantity,
				})
				order.Quantity=0
			}
		}
		// If there is not enough suitable offers, we create new order
		if order.Quantity > 0 {
			err := tx.Create(&model.Table{
				UserID:   order.UserId,
				Value:    order.Value,
				Ticker:   order.Ticker,
				Quantity: order.Quantity,
				Buy:      order.Buy,
			}).Error
			if err != nil {
				return err
			}
		}
		return nil
	})

	return response, err
}
