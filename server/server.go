package main

import (
	"context"
	"errors"
	"grpc_test_example/pb"
	"strconv"
)

type telephoneServer struct {
	pb.UnimplementedTelephoneServer
	contacts []*pb.ListContactsResponse
	calls    []string
}

func NewServer() pb.TelephoneServer {
	t := &telephoneServer{}
	t.loadContact()
	return t
}

func (t *telephoneServer) GetContact(ctx context.Context, telNum *pb.GetContactRequest) (*pb.GetContactResponse, error) {
	_, err := strconv.ParseUint(telNum.Number, 10, 64)

	if len(telNum.Number) != 11 || telNum.Number == "" || err != nil {
		return &pb.GetContactResponse{}, errors.New("invalid number")
	}
	for _, c := range t.contacts {
		if telNum.Number == c.Number {
			return &pb.GetContactResponse{
				Name:     c.Name,
				Lastname: c.Lastname,
				Number:   c.Number,
			}, nil
		}
	}
	return &pb.GetContactResponse{}, errors.New("no contact found")
}

func (t *telephoneServer) ListContacts(_, *pb.ListContactsRequest, stream pb.Telephone_ListContactsServer) error {
	for _,
}

func (t *telephoneServer) loadContact() {
	t.contacts = []*pb.ListContactsResponse{
		{
			Name:     "Nukhet",
			Lastname: "Duru",
			Number:   "111111",
		},
		{
			Name:     "Zeki",
			Lastname: "Muren",
			Number:   "2222222",
		},
		{
			Name:     "Sebnew",
			Lastname: "Ferah",
			Number:   "333333333",
		},
	}
}