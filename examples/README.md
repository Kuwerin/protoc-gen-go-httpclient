# DummyJSON

To show the usage of this plugin, try to describe the open dummy user API of [DummyJSON](https://dummyjson.com/users/1)

For example, we could implement GET method for the route `https://jsonplaceholder.typicode.com/users/1`.

We can get a user by calling the handler with cURL:

#### In

```bash
curl https://dummyjson.com/users/1
```

#### Out

```json
{
  "id": 1,
  "firstName": "Emily",
  "lastName": "Johnson",
  "maidenName": "Smith",
  "age": 28,
  "gender": "female",
  "email": "emily.johnson@x.dummyjson.com",
  "phone": "+81 965-431-3024",
  "username": "emilys",
  "password": "emilyspass",
  "birthDate": "1996-5-30",
  "image": "https://dummyjson.com/icon/emilys/128",
  "bloodGroup": "O-",
  "height": 193.24,
  "weight": 63.16,
  "eyeColor": "Green",
  "hair": {
    "color": "Brown",
    "type": "Curly"
  },
  "ip": "42.48.100.32",
  "address": {
    "address": "626 Main Street",
    "city": "Phoenix",
    "state": "Mississippi",
    "stateCode": "MS",
    "postalCode": "29112",
    "coordinates": {
      "lat": -77.16213,
      "lng": -92.084824
    },
    "country": "United States"
  },
  "macAddress": "47:fa:41:18:ec:eb",
  "university": "University of Wisconsin--Madison",
  "bank": {
    "cardExpire": "03/26",
    "cardNumber": "9289760655481815",
    "cardType": "Elo",
    "currency": "CNY",
    "iban": "YPUXISOBI7TTHPK2BR3HAIXL"
  },
  "company": {
    "department": "Engineering",
    "name": "Dooley, Kozey and Cronin",
    "title": "Sales Manager",
    "address": {
      "address": "263 Tenth Street",
      "city": "San Francisco",
      "state": "Wisconsin",
      "stateCode": "WI",
      "postalCode": "37657",
      "coordinates": {
        "lat": 71.814525,
        "lng": -161.150263
      },
      "country": "United States"
    }
  },
  "ein": "977-175",
  "ssn": "900-590-289",
  "userAgent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.93 Safari/537.36",
  "crypto": {
    "coin": "Bitcoin",
    "wallet": "0xb9fc2fe63b2a6c003f1c324c3bfa53259162181a",
    "network": "Ethereum (ERC20)"
  },
  "role": "admin"
}
```

The plugin allows to call an HTTP handler from the gRPC server and operate over gRPC messages instead of JSON.
First, describe the message:

#### proto/dummyjson/user/v1/user.proto
```protobuf
syntax = "proto3";

package dummyjson.user.v1;

option go_package = "github.com/Kuwerin/protoc-gen-go-httpclient/examples/dummyjson/user/v1;userapi";

// UserModel.
message User {
  // The user's hair.
  message Hair {
    // The user's hair color.
    string color = 1;

    // The user's hair type.
    string type = 2;
  }

  // The user's coordinates.
  message Coordinates {
    // Latitude.
    double lat = 1;

    // Longitude.
    double lng = 2;
  }

  // The address.
  message Address {
    // The address.
    string address = 1;

    // The city.
    string city = 2;


    // The state.
    string state = 3;

    // The state code.
    string state_code = 4;

    // The postal code.
    string postal_code = 5;

    // The coordinates.
    Coordinates coordinates = 6;

    // The country
    string country = 7;
  }

  // The user's bank info.
  message Bank {
    // The card expiration date.
    string card_expire = 1;

    // The card number.
    string card_number = 2;

    // The type of the card.
    string card_type = 3;

    // The currency of the card.
    string currency = 4;

    // IBAN.
    string iban = 5;
  }

  // The company info.
  message Company {
    // The department.
    string department = 1;

    // The name of the company.
    string name = 2;

    // The title of the company.
    string title = 3;

    // The address of the company.
    Address address = 4;
  }

  // The information about crypto.
  message Crypto {
    // Coin.
    string coin = 1;

    // Wallet.
    string wallet = 2;

    // Network.
    string network = 3;
  }

  // The user's unique identifier.
  uint32 id = 1;

  // The user's first name.
  string first_name = 2;

  // The user's last name.
  string last_name = 3;

  // The user's maiden name.
  string maiden_name = 4;

  // The user's age.
  uint32 age = 5;

  // The user's gender.
  string gender = 6;

  // The user's email.
  string email = 7;

  // The user's phone.
  string phone = 8;

  // The user's username.
  string username = 9;

  // The user's password.
  string password = 10;

  // The user's birth date.
  string birth_date = 11;

  // The user's image.
  string image = 12;

  // The user's blood group.
  string blood_group = 13;

  // The user's height.
  double height = 14;

  // The user's weight.
  double weight = 15;

  // The user's eye color.
  string eye_color = 16;

  // The user's hair.
  Hair hair = 17;

  // The user's IP address.
  string ip = 18;

  // The user's address.
  Address address = 19;

  // The user's MAC address.
  string mac_address = 20;

  // The user's university.
  string university = 21;

  // The user's bank account info.
  Bank bank = 22;

  // The user's company info.
  Company company = 23;

  // The user's EIN.
  string ein = 24;

  // The user's SSN.
  string ssn = 25;

  // The user's UserAgent.
  string user_agent = 26;

  // The info about user's crypto account.
  Crypto crypto = 27;

  // The user's role.
  string role = 28;
}
```

Second, describe the route, request and response messages:

#### proto/dummyjson/user/v1/user_service.proto

```protobuf
syntax = "proto3";

package dummyjson.user.v1;

import "google/api/annotations.proto";
import "dummyjson/user/v1/user.proto";

option go_package = "github.com/Kuwerin/protoc-gen-go-httpclient/examples/dummyjson/user/v1;userapi";

// UserService.
service UserService {
  // Gets a user.
  rpc GetUser(GetUserRequest) returns (User) {
    option (google.api.http) = {
      get: "/users/{id}"
    };
  }

// Request message for dummyjson.user.v1.GetUser method.
message GetUserRequest {
  // The unique user identifier.
  string id = 1;
}
```

For example, to call get method we can describe GetUser method with `id` as a request parameter and `User` message as a response.

Then compile protofiles with plugin option flags:

```bash
	mkdir -p out/go
	cd out/go && rm -rf *;
	
	# User
	protoc -Iproto \
	                --go_out=out/go \
	                --go_opt paths=source_relative \
	                --go-grpc_out=out/go \
	                --go-grpc_opt paths=source_relative \
	                --go-httpclient_out=out/go \
	                --go-httpclient_opt logging_middleware=true \
	                --go-httpclient_opt paths=source_relative \
	                proto/dummyjson/user/v1/*.proto
```
There is additional optional boolean flag for logging `logging_middleware`.

To start gRPC server there is helperfunc package `transport`.

```go
package main

import (
	"net"
	"net/url"
	"os"
	"os/signal"
	"syscall"

	"github.com/Kuwerin/protoc-gen-go-httpclient/pkg/transport"
	"github.com/go-kit/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	userapi "dummyjson/out/go/user/v1"
)

func main() {
	// Create the logger.
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(log.NewSyncWriter(os.Stdout))
		logger = log.With(logger, "ts", log.DefaultTimestampUTC)
		logger = log.With(logger, "svc", "erp-client")
		logger.Log("app", os.Args[0], "event", "starting")
	}

	baseURL, _ := url.Parse("https://dummyjson.com")

	// Create user client.
	var userClient userapi.UserServiceServer
	{
		userClient = userapi.NewHTTPClient(&userapi.HTTPClientParams{
			URL:                   baseURL,
			AllowUndeclaredFields: true,
		})
		userClient = userapi.LoggingMiddleware(logger)(userClient)
	}

	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)

	ln, err := net.Listen("tcp", ":5000")
	if err != nil {
		logger.Log("error", err)
		os.Exit(1)
	}

	// Register user server.
	userServer := &transport.Server[userapi.UserServiceServer]{
		Server:             grpcServer,
		ContextAPI:         &userClient,
		RegisterServerFunc: userapi.RegisterUserServiceServer,
	}
	if err := userServer.Register(); err != nil {
		logger.Log("entity", "transport.grpc.user", "error", err)
		os.Exit(1)
	}
	logger.Log("entity", "transport.grpc.user", "event", "registred")

	// Start listening gRPC server.
	go func() {
		if err := grpcServer.Serve(ln); err != nil {
			logger.Log("error", err)
			os.Exit(1)
		}
	}()
	logger.Log("event", "started listening")

	// Wait for a signal for graceful shutdown.
	quit := make(chan os.Signal, 1)

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	s := <-quit

	logger.Log("event", "received os signal to shutdown application", "signal", s)

	grpcServer.GracefulStop()

	logger.Log("event", "gRPC server stopped gracefully")

	logger.Log("event", "application stopped gracefully")

}
```

For convenience we can use `docker compose` with `grpcox` container inside.

![example1](assets/1.gif)
Also we can check the work of plugin via commandline to compare HTTP response with plugin-generated gRPC one side-by-side.

E.g. this cURL-request

```bash
curl https://dummyjson.com/users/1
```

is equivalent to gRPCurl:

```bash
grpcurl -d '{"id": "1"}' -plaintext localhost:5000 dummyjson.user.v1.UserService/GetUser
```

![example2](assets/2.gif)