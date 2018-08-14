package models

import (
	"context"
	"log"

	"github.com/oswee/proto/metric/go"
)

// CreatePageView ...
func CreatePageView(XForwardedHost, XForwardedServer, UserAgent, XForwardedFor, RequestTime string) (*metric.Empty, error) {
	cc := gLoc()
	defer cc.Close()
	c := metric.NewMetricClient(cc)
	req := &metric.CreatePageViewRequest{
		PageView: &metric.PageView{
			XForwardedHost:   XForwardedHost,
			XForwardedServer: XForwardedServer,
			UserAgent:        UserAgent,
			XForwardedFor:    XForwardedFor,
			RequestTime:      RequestTime,
		}}

	res, err := c.CreatePageView(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling CreatePageView RPC: %v", err)
	}
	return res, nil
}
