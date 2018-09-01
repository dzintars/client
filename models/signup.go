package models

import (
	"context"
	"log"

	signup "github.com/oswee/stubs/signup/v1"
)

// CreateSignup ...
func CreateSignup(fname, lname, email, uname, password string) (*signup.Signup, error) {
	cc := gLoc()
	defer cc.Close()
	c := signup.NewSignupServiceClient(cc)
	req := &signup.CreateSignupRequest{
		Signup: &signup.Signup{
			FistName: fname,
			LastName: lname,
			Email:    email,
			Username: uname,
			Password: password,
		}}
	res, err := c.CreateSignup(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling CreateSignup RPC: %v", err)
	}
	return res, nil
}
