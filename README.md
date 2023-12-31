# Receipt Processor

get your rewards, for your receipts!

A webserver that accepts receipts and computes reward points.

## specification
the rules for awarding points to receipts are [here](./specification.md)

## pre-requisites
- go 1.20+
- docker 24+, for documentation hosting

## build
```
make build
```

## Test
- to add additional mocks, install mockery
```
make install-dep
make mocks
```
- run tests
```
make test
```

## run

```
./build/receipt-processor
```

## usage

Refer Open API [documentation](./docs).

### examples
#### example 1
```
$ curl -X POST http://localhost:1201/receipts/process \
   -H "Content-Type: application/json" \
   -d '<request-body>'

{"id":"75ff6a62-4d86-4e4f-80db-7c2a4e733388"}
```
request body:
```json
{
  "retailer": "Target",
  "purchaseDate": "2022-01-01",
  "purchaseTime": "13:01",
  "items": [
    {
      "shortDescription": "Mountain Dew 12PK",
      "price": "6.49"
    },{
      "shortDescription": "Emils Cheese Pizza",
      "price": "12.25"
    },{
      "shortDescription": "Knorr Creamy Chicken",
      "price": "1.26"
    },{
      "shortDescription": "Doritos Nacho Cheese",
      "price": "3.35"
    },{
      "shortDescription": "   Klarbrunn 12-PK 12 FL OZ  ",
      "price": "12.00"
    }
  ],
  "total": "35.35"
}
```
```
$ curl http://localhost:1201/receipts/75ff6a62-4d86-4e4f-80db-7c2a4e733388/points
{"points":"28"}
```
```text
Total Points: 28
Breakdown:
     6 points - retailer name has 6 characters
    10 points - 4 items (2 pairs @ 5 points each)
     3 Points - "Emils Cheese Pizza" is 18 characters (a multiple of 3)
                item price of 12.25 * 0.2 = 2.45, rounded up is 3 points
     3 Points - "Klarbrunn 12-PK 12 FL OZ" is 24 characters (a multiple of 3)
                item price of 12.00 * 0.2 = 2.4, rounded up is 3 points
     6 points - purchase day is odd
  + ---------
  = 28 points
```

example 2

```json
{
  "retailer": "M&M Corner Market",
  "purchaseDate": "2022-03-20",
  "purchaseTime": "14:33",
  "items": [
    {
      "shortDescription": "Gatorade",
      "price": "2.25"
    },{
      "shortDescription": "Gatorade",
      "price": "2.25"
    },{
      "shortDescription": "Gatorade",
      "price": "2.25"
    },{
      "shortDescription": "Gatorade",
      "price": "2.25"
    }
  ],
  "total": "9.00"
}
```
```text
Total Points: 109
Breakdown:
    50 points - total is a round dollar amount
    25 points - total is a multiple of 0.25
    14 points - retailer name (M&M Corner Market) has 14 alphanumeric characters
                note: '&' is not alphanumeric
    10 points - 2:33pm is between 2:00pm and 4:00pm
    10 points - 4 items (2 pairs @ 5 points each)
  + ---------
  = 109 points
```

## backlog
- error handling
- logging
- unit tests for rules engine
- bug: amount values must have exactly 2 decimal places only