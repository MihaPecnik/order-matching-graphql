# order-matching-graphql

## How to run application

  Run postgres cointainer from docker-compose.yml file, then :
  * go build server.go
  

## Api endpoints
* localhost:8080

###  Get tickers bottom of the buy-side and the top of the sell-side
Given as input a ticker (return the bottom of the buy-side and the top of the
sell-side). Basically it just returns the "BUY" order with the highest value and it's quantity and order "SELL" with the lowest value and its quantity.


* http://localhost:8080/query

Request:
```json
query {
  getInfoTicker(input: {ticker: "APPL"}){
    buy{
        Value,
     		Quantity
    },
    sell{
      Value,
     	Quantity
    }
  }
}
```
Response:
```json
{
  "data": {
    "getInfoTicker": {
      "buy": {
        "Value": 200.3,
        "Quantity": 5
      },
      "sell": {
        "Value": 201.1,
        "Quantity": 1
      }
    }
  }
}
```


###  Update Order's book
Given as input an id User, a ticker (eg. AAPL), a value, an int quantity and a
command (buy or sell) which is an order. Update an order’s book on the buy
and sell-side accordingly, return the quantities and prices that got matched as
a result of the order insertion.

If any error(e.g. electricity outage, server error) occurs, data will automatically rollback.


[Algorithm](https://www.youtube.com/watch?v=Kl4-VJ2K8Ik)
* http://localhost:8080/query

Request:
```json
mutation {
  updateOrderBook(input: {UserId: 1, Value: 205.00, Quantity: 4, Buy: true, Ticker: "APPL"}){
    Value,
    Quantity
  }
}
```
Response:
```json
{
  "data": {
    "updateOrderBook": [
      {
        "Value": 201.1,
        "Quantity": 1
      },
      {
        "Value": 201.2,
        "Quantity": 1
      },
      {
        "Value": 201.3,
        "Quantity": 2
      }
    ]
  }
}
```
