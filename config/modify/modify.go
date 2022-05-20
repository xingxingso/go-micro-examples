package main

import (
	"fmt"
	"os"

	"github.com/asim/go-micro/plugins/config/encoder/toml/v4"
	"go-micro.dev/v4/config"
	"go-micro.dev/v4/config/reader"
	"go-micro.dev/v4/config/reader/json"
	"go-micro.dev/v4/config/source/file"
)

func main() {
	// new toml encoder
	t := toml.NewEncoder()

	// create a new config
	c, err := config.NewConfig(
		config.WithReader(json.NewReader(reader.WithEncoder(t))),
		config.WithSource(
			// create a new file source
			file.NewSource(
				// path of file
				file.WithPath("./example.toml"),
				// specify the toml encoder
				// source.WithEncoder(t), // not work
			),
		),
	)
	if err != nil {
		fmt.Println(err)
		return
	}

	// load the config
	if err := c.Load(
		//file.NewSource(
		//	// path of file
		//	file.WithPath("./example.toml"),
		//	// specify the toml encoder
		//	source.WithEncoder(t),
		//),
	); err != nil {
		fmt.Println(err)
		return
	}

	// set a value
	c.Set("foo", "bar")

	// now the hacks begin
	vals := c.Map()

	// encode
	v, err := t.Encode(vals)
	if err != nil {
		fmt.Println(err)
		return
	}

	// write the file
	if err := os.WriteFile("./example.toml", v, 0644); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("wrote update to example.toml")
}
