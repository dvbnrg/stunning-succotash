package main

import (
	"context"
	"log"
	"net"
	"time"

	"userService/pb"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

const (
	port = ":8080"
)

type server struct {
	pb.UnimplementedUserServiceServer
}

func main() {
	log.Printf("Starting new userService on port %v", port)
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterUserServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func (s *server) CreateUser(ctx context.Context, user *pb.User) (*pb.User, error) {
	log.Println("Creating User: %+v", user)
	db, cancel := mgoconnect()
	defer cancel()
	collection := db.Database("Users").Collection("Users")
	insertResult, err := collection.InsertOne(ctx, user)
	if err != nil {
		log.Println("Data Insertion Error: ", err)
		return user, err
	}
	log.Printf("Created User: %+v", insertResult)
	return user, nil
}

func (s *server) GetUser(ctx context.Context, user *pb.User) (*pb.User, error) {
	log.Println("Retreiving User: %+v", user)
	db, cancel := mgoconnect()
	defer cancel()
	collection := db.Database("Users").Collection("Users")
	err := collection.FindOne(ctx, user).Decode(&user)
	if err != nil {
		log.Println("Data Retreival Error: ", err)
		return user, err
	}
	log.Printf("Found User: %+v", user)
	return user, nil
}

func (s *server) UpdateUser(ctx context.Context, user *pb.User) (*pb.User, error) {
	log.Println("Updating User: %+v", user)
	db, cancel := mgoconnect()
	defer cancel()
	collection := db.Database("Users").Collection("Users")
	filter := bson.D{{}}
	updateResult, err := collection.UpdateOne(ctx, filter, user)
	if err != nil {
		log.Println("Data Retreival Error: ", err)
		return user, err
	}
	log.Printf("Updated User: %+v", updateResult)
	return user, nil
}

func (s *server) ListUsers(e *emptypb.Empty, stream pb.UserService_ListUsersServer) error {
	log.Println("Retreiving Users: %+v", stream)
	db, cancel := mgoconnect()
	defer cancel()
	collection := db.Database("Users").Collection("Users")
	filter := bson.D{{}}
	ctx := stream.Context()
	cur, err := collection.Find(ctx, filter)
	if err != nil {
		log.Println("Data Retreival Error: ", err)
		return err
	}
	var results []*pb.User
	for cur.Next(ctx) {
		user := pb.User{}
		err := cur.Decode(&user)
		if err != nil {
			log.Println("User Stream Decode Error: ", err)
		}
		results = append(results, &user)
	}
	if err := cur.Err(); err != nil {
		log.Println(err)
	}
	cur.Close(ctx)
	log.Printf("Found the following %+v", results)
	return nil
}

func (s *server) DeleteUser(ctx context.Context, user *pb.User) (*emptypb.Empty, error) {
	log.Println("Deleting User: %+v", user)
	db, cancel := mgoconnect()
	defer cancel()
	collection := db.Database("Users").Collection("Users")
	filter := bson.D{{}}
	deleteResult, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		log.Println("Data Retreival Error: ", err)
		return &emptypb.Empty{}, err
	}
	log.Printf("Deleted %+v documents in the trainers collection\n", deleteResult)
	return &emptypb.Empty{}, nil
}

// Connect opens a db connection to Mongo
func mgoconnect() (mgo *mongo.Client, cancel context.CancelFunc) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	mgo, err := mongo.Connect(ctx, options.Client().ApplyURI(
		"mongodb+srv://justdave:supersecret@cluster0.xsmx6.gcp.mongodb.net/Users?retryWrites=true&w=majority",
	))
	if err != nil {
		log.Println(err)
		return
	}
	return
}
