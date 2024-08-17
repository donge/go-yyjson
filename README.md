# High-Performance JSON Parsing in Go using YYJSON

This Go project demonstrates how to use the YYJSON library for high-performance JSON parsing. YYJSON is a fast and lightweight JSON library that provides a C API, which this project wraps with a Go interface.

## Features

- Unmarshals JSON data into a `map[string]interface{}` using the YYJSON library.
- Compares the result with the standard Go `json.Unmarshal` function.
- Provides a simple and efficient way to work with JSON data in Go.

## Requirements

- Go 1.16 or later
- The YYJSON library, which is included as a C header file in the project.

## Usage

1. Clone the repository:

```azure
git clone https://github.com/donge/go-yyjson.git

```

2. Build and run the project:
```azure
go build -o yyjson-go
./yyjson-go
```


This shows that the YYJSON-based unmarshaling and the standard Go `json.Unmarshal` produce the same result.

## How it works

The `Unmarshal` function in the project uses the YYJSON library to parse the input JSON data and convert it to a `map[string]interface{}`. The conversion is done by the `convertCValueToMap` and `convertCValueToInterface` functions, which recursively traverse the JSON structure and convert the YYJSON values to their Go equivalents.

The project also includes a simple benchmark to compare the performance of the YYJSON-based unmarshaling with the standard Go `json.Unmarshal` function.

## Limitations

This project only demonstrates the basic usage of the YYJSON library for unmarshaling JSON data. In a real-world application, you would need to handle more complex use cases, such as error handling, custom struct unmarshaling, and performance optimization.

## Contributing

If you find any issues or have suggestions for improvement, feel free to create a new issue or submit a pull request.
