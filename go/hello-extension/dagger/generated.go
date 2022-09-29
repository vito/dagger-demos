// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package main

import (
	"context"
	"encoding/json"

	"go.dagger.io/dagger/sdk/go/dagger"
)

func (r *query) hello(ctx context.Context) (*Hello, error) {

	return new(Hello), nil

}

type hello struct{}
type query struct{}

func main() {
	dagger.Serve(context.Background(), map[string]func(context.Context, dagger.ArgsInput) (interface{}, error){
		"Hello.build": func(ctx context.Context, fc dagger.ArgsInput) (interface{}, error) {
			var bytes []byte
			_ = bytes
			var err error
			_ = err

			var opts GoOpts

			bytes, err = json.Marshal(fc.Args["opts"])
			if err != nil {
				return nil, err
			}
			if err := json.Unmarshal(bytes, &opts); err != nil {
				return nil, err
			}

			return (&hello{}).build(ctx,

				opts,
			)
		},
		"Hello.run": func(ctx context.Context, fc dagger.ArgsInput) (interface{}, error) {
			var bytes []byte
			_ = bytes
			var err error
			_ = err

			var opts GoOpts

			bytes, err = json.Marshal(fc.Args["opts"])
			if err != nil {
				return nil, err
			}
			if err := json.Unmarshal(bytes, &opts); err != nil {
				return nil, err
			}

			return (&hello{}).run(ctx,

				opts,
			)
		},
		"Hello.test": func(ctx context.Context, fc dagger.ArgsInput) (interface{}, error) {
			var bytes []byte
			_ = bytes
			var err error
			_ = err

			var opts GoOpts

			bytes, err = json.Marshal(fc.Args["opts"])
			if err != nil {
				return nil, err
			}
			if err := json.Unmarshal(bytes, &opts); err != nil {
				return nil, err
			}

			return (&hello{}).test(ctx,

				opts,
			)
		},
		"Query.hello": func(ctx context.Context, fc dagger.ArgsInput) (interface{}, error) {
			var bytes []byte
			_ = bytes
			var err error
			_ = err

			return (&query{}).hello(ctx)
		},
	})
}
