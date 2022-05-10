package main

import (
	"time"

	"github.com/cookie-for-pres/chc"
)

func controller(req *chc.Request, res chc.Response) *chc.Response {
	res.SetStatusCode(200)
	res.SetHeader("Content-Type", "application/json")

	headers := map[string]string{}
	for k, v := range res.Headers {
		headers[k] = v
	}

	res.SetJsonObjectBody(map[string]interface{}{"headers": headers})

	return &res
}

func middleware1(response chc.Response) (*chc.Response, bool) {
	response.SetHeader("X-Test", time.Now().String())
	return &response, true
}

func middleware2(response chc.Response) (*chc.Response, bool) {
	response.SetHeader("X-Test-2", time.Now().String())
	return &response, true
}

func main() {
	CHC := chc.NewCHC()
	route := &chc.Route{
		Path:       "/",
		Methods:    []string{"GET"},
		Controller: controller,
	}
	route.AddMiddleware(middleware1)
	route.AddMiddleware(middleware2)
	CHC.AddRoute(route)

	CHC.Listen("localhost", 8080)
}
