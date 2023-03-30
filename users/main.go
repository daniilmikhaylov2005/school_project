package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	pb "github.com/daniilmikhaylov2005/school_project/api"
	"github.com/daniilmikhaylov2005/school_project/users/repository"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserServer struct {
	hashCost   int
	repository *repository.Repository
	pb.UnimplementedUserServer
}

func (s *UserServer) CreateUser(ctx context.Context, in *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(in.Password), 7)
	if err != nil {
		return nil, status.Error(codes.Aborted, "failed to hash password")
	}
	id, err := s.repository.CreateUser(in.FirstName, in.SecondName, in.Login, in.Email, string(hashedPassword))
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to create user")
	}
	return &pb.CreateUserResponse{Login: in.Login, Password: in.Password, Id: int64(id)}, nil
}

func main() {
	port := flag.String("port", "50000", "port for service")
	hashCost := flag.Int("hashCost", 7, "cost for hashing")
	flag.Parse()

	lis, err := net.Listen("tcp", ":"+*port)
	if err != nil {
		log.Fatalf("failed listen: %v", err)
	}

	s := grpc.NewServer()
	postgresDB, err := NewConn()
	if err != nil {
		log.Fatalf("failed to connect postgres: %v", err)
	}

	repository := repository.NewRepository(postgresDB)

	pb.RegisterUserServer(s, &UserServer{hashCost: *hashCost, repository: repository})
	go func() {
		log.Printf("Server started at port: %s\n", *port)
		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	sig := make(chan os.Signal, 2)
	signal.Notify(sig, os.Interrupt)
	signal.Notify(sig, syscall.SIGTERM)

	ch := <-sig

	log.Printf("Recieved terminate: %v\n", ch)
	s.Stop()
}

func NewConn() (*sqlx.DB, error) {
	if err := godotenv.Load("./users/.env"); err != nil {
		return nil, err
	}

	dbURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		os.Getenv("HOST"), os.Getenv("PORT"), os.Getenv("USER"), os.Getenv("NAME"), os.Getenv("PASSWORD"), os.Getenv("SSL"))

	db, err := sqlx.Open("postgres", dbURL)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}
