package models

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	pb "github.com/oswee/stubs"
	metric "github.com/oswee/stubs/metric/v1"
)

// CreatePageView ...
func CreatePageView(r *http.Request) (*pb.Empty, error) {
	t := time.Now()
	rtime := t.String()
	//fmt.Println(r.UserAgent())
	cc := gLoc()
	defer cc.Close()
	c := metric.NewMetricClient(cc)
	// https://mycodesmells.com/post/grpc-client-server-example
	req := &metric.CreatePageViewRequest{
		PageView: &metric.PageView{
			XForwardedFor:    r.Header.Get("X-Forwarded-For"),
			XForwardedHost:   r.Host,
			XForwardedServer: r.Header.Get("X-Forwarded-Server"),
			Url:              r.URL.String(),
			UserAgent:        r.UserAgent(),
			RequestTime:      rtime,
		}}

	res, err := c.CreatePageView(context.Background(), req)
	if err != nil {
		fmt.Println(r.Header.Get("X-Forwarded-For"), r.UserAgent())
		log.Fatalf("Error while calling CreatePageView RPC: %v", err)
	}
	return res, nil
}
