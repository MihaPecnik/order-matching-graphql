// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type GetTickerInfoResponse struct {
	Buy  *UpdateOrderBookResponse `json:"buy"`
	Sell *UpdateOrderBookResponse `json:"sell"`
}

type Request struct {
	UserID   int     `json:"UserId"`
	Buy      bool    `json:"Buy"`
	Value    float64 `json:"Value"`
	Quantity int     `json:"Quantity"`
	Ticker   string  `json:"Ticker"`
}

type Table struct {
	ID       int     `json:"id"`
	UserID   int     `json:"UserId"`
	Buy      bool    `json:"Buy"`
	Value    float64 `json:"Value"`
	Quantity int     `json:"Quantity"`
	Ticker   string  `json:"Ticker"`
}

type Ticker struct {
	Ticker string `json:"ticker"`
}

type UpdateOrderBookResponse struct {
	Value    float64 `json:"Value"`
	Quantity int     `json:"Quantity"`
}
