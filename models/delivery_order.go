package models

import (
	"context"
	"log"

	"github.com/oswee/proto/shipping/go"
)

// ListDeliveryOrders returns all delivery orders of requested stakeholder
func ListDeliveryOrders(stakeholderID int64, resultsPerPage int32) (*shipping.ListDeliveryOrdersResponse, error) {
	cc := gLoc()
	defer cc.Close()
	c := shipping.NewShippingServiceClient(cc)
	req := &shipping.ListDeliveryOrdersRequest{}
	req.StakeholderId = stakeholderID
	req.ResultPerPage = resultsPerPage
	res, err := c.ListDeliveryOrders(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling ListApplications RPC: %v", err)
	}
	return res, nil
}
