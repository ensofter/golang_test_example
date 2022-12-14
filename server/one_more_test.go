package main

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"grpc_test_example/pb"
)

func TestCreateDevice(t *testing.T) {
	host := "localhost:9090"
	ctx := context.Background()
	conn, err := grpc.Dial(host, grpc.WithInsecure())
	if err != nil {
		t.Fatalf("grpc.Dial() err: %v", err)
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			t.Logf("conn.Close err: %v", err)
		}
	}(conn)

	telephoneClient := pb.NewTelephoneClient(conn)

	t.Run("Get contacts", func(t *testing.T) {
		req := pb.GetContactRequest{
			Number: "33333333333",
		}

		res, err := telephoneClient.GetContact(ctx, &req)
		require.NoError(t, err)
		require.NotNil(t, res)
		t.Log("!!!", res)
		assert.Equal(t, res.Name, "Sebnew")
	})
}
