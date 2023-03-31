package main

import (
	"context"
	"log"

	pb "github.com/daniilmikhaylov2005/school_project/api"
	"google.golang.org/grpc"
)

const address = "localhost:50000"

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("failed to conn: %v", err)
	}
	defer conn.Close()

	client := pb.NewUserClient(conn)

	// u, err := client.CreateUser(context.Background(), &pb.CreateUserRequest{
	// 	FirstName:  "Даниил",
	// 	SecondName: "Михайлов",
	// 	Email:      "sadrezdev@gmail.com",
	// 	Login:      "sadrezdev",
	// 	Password:   "123123",
	// })

	// log.Printf(`Details:
	// Id: %d
	// Login: %s
	// Password: %s
	// `, u.Id, u.Login, u.Password)

	u, err := client.GetUser(context.Background(), &pb.GetUserRequest{Login: "sadrezdev", Password: "123123"})
	if err != nil {
		log.Fatalf("failed to get user: %v", err)
	}

	log.Println(u)
}
