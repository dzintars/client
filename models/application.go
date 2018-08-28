package models

import (
	"context"
	"log"

	app "github.com/oswee/stubs/app/v1"
)

// GetApplication returns application by requested ID
func GetApplication(applicationID int32) (*app.GetApplicationResponse, error) {
	cc := gLoc()
	defer cc.Close()

	c := app.NewApplicationServiceClient(cc)
	req := &app.GetApplicationRequest{}
	req.Id = applicationID
	res, err := c.GetApplication(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling GetApplication RPC: %v", err)
	}
	return res, nil
}

// ListApplications retrieve list of all films
func ListApplications(resultsPerPage int32) (*app.ListApplicationsResponse, error) {
	cc := gLoc()
	defer cc.Close()
	// We can now create stubs that wrap cc:
	c := app.NewApplicationServiceClient(cc)
	req := &app.ListApplicationsRequest{}
	req.ResultPerPage = resultsPerPage
	res, err := c.ListApplications(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling ListApplications RPC: %v", err)
	}
	return res, nil
}
