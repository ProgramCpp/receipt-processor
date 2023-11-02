# Receipt Processor

get your rewards, for your receipts!

A webserver that accepts receipts and computes reward points.

## pre-requisites
- go 1.20+

## build
```
make build
```

## Test
- install mockery
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


## future work
- as more endpoints are added, use [OpenApi code generator](https://github.com/deepmap/oapi-codegen), to generate server boilerplate




