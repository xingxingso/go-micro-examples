# Function

This is an example of creating a micro function. A function is a one time executing service.

## Contents

- main.go - is the main definition of the function
- proto - contains the protobuf definition of the API

## Run function

```shell
#while true; do
#	github.com/go-micro/examples/function
#done

while true; do
  go run main.go
done
```

## Call function

```shell
micro call greeter Greeter.Hello '{"name": "john"}'
```
