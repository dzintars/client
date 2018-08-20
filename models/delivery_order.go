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

//CreateDeliveryOrder .....
func CreateDeliveryOrder(ref string, da string, dz string, lat float64, lng float64, tw float64, rs int64) (*shipping.DeliveryOrder, error) {
	cc := gLoc()
	defer cc.Close()
	c := shipping.NewShippingServiceClient(cc)
	req := &shipping.CreateDeliveryOrderRequest{
		DeliveryOrder: &shipping.DeliveryOrder{
			Reference:          ref,
			DestinationAddress: da,
			DestinationZip:     dz,
			DestinationLat:     lat,
			DestinationLng:     lng,
			TotalWeight:        tw,
			RoutingSequence:    rs,
		}}
	res, err := c.CreateDeliveryOrder(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling CreateDeliveryOrder RPC: %v", err)
	}
	return res, nil
}

//UpdateDeliveryOrder .....
func UpdateDeliveryOrder(ref string, da string, dz string, lat float64, lng float64, tw float64, rs int64, id int64) (*shipping.DeliveryOrder, error) {
	cc := gLoc()
	defer cc.Close()
	c := shipping.NewShippingServiceClient(cc)
	req := &shipping.UpdateDeliveryOrderRequest{
		DeliveryOrder: &shipping.DeliveryOrder{
			Reference:          ref,
			DestinationAddress: da,
			DestinationZip:     dz,
			DestinationLat:     lat,
			DestinationLng:     lng,
			TotalWeight:        tw,
			RoutingSequence:    rs,
			Id:                 id,
		}}
	res, err := c.UpdateDeliveryOrder(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling UpdateDeliveryOrder RPC: %v", err)
	}
	return res, nil
}

// DeleteDeliveryOrder ....
func DeleteDeliveryOrder(id int64) (*shipping.EmptyDeliveryOrder, error) {
	cc := gLoc()
	defer cc.Close()
	c := shipping.NewShippingServiceClient(cc)
	req := &shipping.DeleteDeliveryOrderRequest{
		Id: id,
	}
	res, err := c.DeleteDeliveryOrder(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling DeleteDeliveryOrder RPC: %v", err)
	}
	return res, nil
}

// GeoCodeDeliveryOrder ...
func GeoCodeDeliveryOrder(id, stakeholder int64, address string) (*shipping.DeliveryOrder, error) {
	cc := gLoc()
	defer cc.Close()
	c := shipping.NewShippingServiceClient(cc)
	req := &shipping.GeoCodeDeliveryOrderRequest{
		Id:            id,
		StakeholderId: stakeholder,
		Address:       address,
	}
	res, err := c.GeoCodeDeliveryOrder(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling GeoCodeDeliveryOrder RPC: %v", err)
	}
	return res, nil
}
