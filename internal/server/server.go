package server

import (
	"context"

	"github.com/jhuggart/repro-mesh-prob/gen/go/customers"
)

type CustomerServer struct {
	customers.UnimplementedCustomersServer //Required by protobuf as of 02-03-2021. May change in the future
}

func NewServer() *CustomerServer {
	return &CustomerServer{}
}

func (s *CustomerServer) Get(ctx context.Context, msg *customers.CustomerIdentifier) (*customers.CustomerResponse, error) {
	return &customers.CustomerResponse{
		ID:        msg.ID,
		Email:     "testy@mctester.com",
		FirstName: "Testy",
		LastName:  "McTester",
	}, nil
}
