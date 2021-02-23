package main

import (
	"context"
	"log"
	"time"

	"userService/pb"

	"google.golang.org/grpc"
)

const (
	address = "localhost:8080"
)

func main() {

	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())

	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer conn.Close()
	c := pb.NewUserServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	e := &pb.User{
		Email:         "justdave@linux.com",
		Id:            1234,
		AuthID:        "justdave",
		EmailVerified: true,
		GivenName:     "Dave",
		FamilyName:    "Banerjee",
		CreatedAt:     "",
		UpdatedAt:     "",
		DeletedAt:     "",
	}

	r, err := c.CreateUser(ctx, e)
	if err != nil {
		log.Fatalf("could not log: %v", err)
	}

	r, err := c.GetUser(ctx, e)
	if err != nil {
		log.Fatalf("could not log: %v", err)
	}

	// r, err := c.UpdateUser(ctx, e)
	// if err != nil {
	// 	log.Fatalf("could not log: %v", err)
	// }

	// r, err := c.ListUsers(ctx, e)
	// if err != nil {
	// 	log.Fatalf("could not log: %v", err)
	// }

	// r, err := c.DeleteUser(ctx, e)
	// if err != nil {
	// 	log.Fatalf("could not log: %v", err)
	// }

	log.Printf("User added: %t, at Time: %s", r.GetCreatedAt(), r.GetUpdatedAt())
}
