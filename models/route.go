package models

import (
	"context"
	"log"

	route "github.com/oswee/stubs/route/v1"
)

// ListRoutes returns all delivery orders of requested stakeholder
func ListRoutes(resultsPerPage int32) (*route.ListRoutesResponse, error) {
	cc := routeService()
	defer cc.Close()
	c := route.NewRouteServiceClient(cc)
	req := &route.ListRoutesRequest{}
	req.ResultPerPage = resultsPerPage
	res, err := c.ListRoutes(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling ListRoutes RPC: %v", err)
	}
	return res, nil
}
