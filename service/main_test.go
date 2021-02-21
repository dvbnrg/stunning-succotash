package main

import (
	"context"
	"log"
	"net"
	"reflect"
	"testing"
	"userService/pb"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

// const (
// 	port = ":8080"
// )

// type server struct {
// 	pb.UnimplementedUserServiceServer
// }

func Test_server_CreateUser(t *testing.T) {
	type fields struct {
		UnimplementedUserServiceServer pb.UnimplementedUserServiceServer
	}
	type args struct {
		ctx  context.Context
		user *pb.User
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *pb.User
		wantErr bool
	}{
		{name: "test1", fields: fields{}, args: args{}, want: &pb.User{}, wantErr: true},
		{name: "test2", fields: fields{}, args: args{}, want: &pb.User{}, wantErr: true},
		{name: "test3", fields: fields{}, args: args{}, want: &pb.User{}, wantErr: true},
		{name: "test4", fields: fields{}, args: args{}, want: &pb.User{}, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// s := &server{
			// 	UnimplementedUserServiceServer: tt.fields.UnimplementedUserServiceServer,
			// }
			lis, err := net.Listen("tcp", port)
			if err != nil {
				log.Fatalf("failed to listen: %v", err)
			}
			s := grpc.NewServer()
			pb.RegisterUserServiceServer(s, &server{})
			if err := s.Serve(lis); err != nil {
				log.Fatalf("failed to serve: %v", err)
			}
			got, err := s.CreateUser(tt.args.ctx, tt.args.user)
			if (err != nil) != tt.wantErr {
				t.Errorf("server.CreateUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("server.CreateUser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_server_GetUser(t *testing.T) {
	type fields struct {
		UnimplementedUserServiceServer pb.UnimplementedUserServiceServer
	}
	type args struct {
		ctx  context.Context
		user *pb.User
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *pb.User
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &server{
				UnimplementedUserServiceServer: tt.fields.UnimplementedUserServiceServer,
			}
			got, err := s.GetUser(tt.args.ctx, tt.args.user)
			if (err != nil) != tt.wantErr {
				t.Errorf("server.GetUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("server.GetUser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_server_UpdateUser(t *testing.T) {
	type fields struct {
		UnimplementedUserServiceServer pb.UnimplementedUserServiceServer
	}
	type args struct {
		ctx  context.Context
		user *pb.User
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *pb.User
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &server{
				UnimplementedUserServiceServer: tt.fields.UnimplementedUserServiceServer,
			}
			got, err := s.UpdateUser(tt.args.ctx, tt.args.user)
			if (err != nil) != tt.wantErr {
				t.Errorf("server.UpdateUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("server.UpdateUser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_server_ListUsers(t *testing.T) {
	type fields struct {
		UnimplementedUserServiceServer pb.UnimplementedUserServiceServer
	}
	type args struct {
		e      *emptypb.Empty
		stream pb.UserService_ListUsersServer
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &server{
				UnimplementedUserServiceServer: tt.fields.UnimplementedUserServiceServer,
			}
			if err := s.ListUsers(tt.args.e, tt.args.stream); (err != nil) != tt.wantErr {
				t.Errorf("server.ListUsers() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_server_DeleteUser(t *testing.T) {
	type fields struct {
		UnimplementedUserServiceServer pb.UnimplementedUserServiceServer
	}
	type args struct {
		ctx  context.Context
		user *pb.User
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *emptypb.Empty
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &server{
				UnimplementedUserServiceServer: tt.fields.UnimplementedUserServiceServer,
			}
			got, err := s.DeleteUser(tt.args.ctx, tt.args.user)
			if (err != nil) != tt.wantErr {
				t.Errorf("server.DeleteUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("server.DeleteUser() = %v, want %v", got, tt.want)
			}
		})
	}
}
