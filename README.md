# Kottu 2020

This is the 2020 attempted rewrite of [Kottu](https://github.com/janithl/Kottu)
in Go. The code structure is heavily inspired by Uncle Bob and
[go-kit's Shipping example](https://github.com/go-kit/kit/tree/master/examples/shipping).

## Build and Run

```
go build -o bin/kottu -v .

cp .env.example .env
# you will need to update DB values in the .env file

./bin/kottu [-p to set port number]
```

## Run Tests

```
go test
```

## License

Kottu 2020 is released under the [MIT License](http://opensource.org/licenses/MIT).
