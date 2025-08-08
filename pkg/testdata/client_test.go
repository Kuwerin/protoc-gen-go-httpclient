package client

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"testing"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/protobuf/encoding/protojson"
)

func TestClient(t *testing.T) {
	t.Run("test client", func(t *testing.T) {
		baseURL, _ := url.Parse("https://dummyjson.com")

		// Create user client.
		var userClient UserServiceServer
		{
			userClient = NewHTTPClient(&HTTPClientParams{
				URL:                   baseURL,
				AllowUndeclaredFields: true,
			})
		}

		grpcRes, err := userClient.GetUser(context.Background(), &GetUserRequest{Id: "1"})
		if err != nil {
			t.Fatal(err)
		}

		jsonpbMarshaler := &runtime.JSONPb{MarshalOptions: protojson.MarshalOptions{UseProtoNames: false}}

		grpcData, err := jsonpbMarshaler.Marshal(grpcRes)
		if err != nil {
			t.Fatal(err)
		}

		var grpcDataBuf bytes.Buffer
		if err := json.Compact(&grpcDataBuf, grpcData); err != nil {
			t.Fatal(err)
		}

		httpRes, err := http.Get(baseURL.String() + "/users/" + "1")
		if err != nil {
			t.Fatal(err)
		}

		httpData, err := io.ReadAll(httpRes.Body)
		if err != nil {
			t.Fatal(err)
		}

		var httpDataBuf bytes.Buffer
		if err := json.Compact(&httpDataBuf, httpData); err != nil {
			t.Fatal(err)
		}

		if grpcDataBuf.String() != httpDataBuf.String() {
			t.Log("gRPC response:", grpcDataBuf.String())
			t.Log("HTTP response:", httpDataBuf.String())
			t.Fatal("gRPC response data not equals to HTTP one")
		}
	})
}
