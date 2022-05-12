package main

import (
	"github.com/cookie-for-pres/chc"
)

func main() {
	CHC := chc.NewCHC()
	CHC.AddRoute(&chc.Route{
		Path:    "/",
		Methods: []string{"GET"},
		Controller: func(req *chc.Request, res *chc.Response) *chc.Response {
			res.SetStatusCode(200).SetJsonObjectBody(map[any]any{
				"message": "Hello World",
			})

			return res
		},
	})
	CHC.Listen("localhost", 8080)
}
