package main

import (
	"context"
	"errors"
	"grpc_test_example/pb"
	"io"
	"strconv"
	"time"
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

func (t *telephoneServer) ListContacts(_ *pb.ListContactsRequest, stream pb.Telephone_ListContactsServer) error {
	for _, c := range t.contacts {
		if err := stream.Send(&pb.ListContactsResponse{
			Name:     c.Name,
			Lastname: c.Lastname,
			Number:   c.Number,
		}); err != nil {
			return err
		}
	}
	return nil
}

func (t *telephoneServer) RecordCallHistory(stream pb.Telephone_RecordCallHistoryServer) error {
	var callCount int32

	start := time.Now()

	for {
		_, err := stream.Recv()
		if err != nil {
			if errors.Is(err, io.EOF) {
				return stream.SendAndClose(&pb.RecordCallHistoryResponse{
					CallCount:   callCount,
					ElapsedTime: int32(time.Since(start)),
				})
			}
			return err
		}
		callCount++
	}
}

func (t *telephoneServer) SendMessage(stream pb.Telephone_SendMessageServer) error {
	for {
		rec, err := stream.Recv()
		if err != nil {
			if errors.Is(err, io.EOF) {
				return nil
			}
			return err
		}

		t.calls = append(t.calls, string(rec.Msg))

		if t.calls[len(t.calls)-1] == "." {
			for _, m := range t.calls {
				//time.Sleep(time.Second)
				switch m {
				case ".":
				case "Hi!":
					stream.Send(&pb.SendMessageResponse{
						Msg: []byte("Hello!"),
					})
				case "How are you?":
					stream.Send(&pb.SendMessageResponse{
						Msg: []byte("Fine, you?"),
					})
				case "See you later":
					stream.Send(&pb.SendMessageResponse{
						Msg: []byte("See you!"),
					})
				default:
					stream.Send(&pb.SendMessageResponse{
						Msg: []byte("Sorry, I don't understand :/"),
					})
				}
			}
			t.calls = t.calls[:0]
		}

	}
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
