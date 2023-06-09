package main

import (
	"context"
	"errors"
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
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(in.GetPassword()), 7)
	if err != nil {
		return nil, status.Error(codes.Aborted, "failed to hash password")
	}
	id, err := s.repository.CreateUser(in.GetFirstName(), in.GetSecondName(), in.GetLogin(), in.GetEmail(), string(hashedPassword))
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to create user")
	}

	log.Printf("Created user: %s\n", in.GetLogin())
	return &pb.CreateUserResponse{Login: in.GetLogin(), Password: in.GetPassword(), Id: int64(id)}, nil
}

func (s *UserServer) GetUser(ctx context.Context, in *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	user, err := s.repository.GetUserByLogin(in.GetLogin())
	if err != nil {
		if errors.Is(err, repository.NoUsersError) {
			return nil, status.Error(codes.NotFound, "failed to get user")
		}
		return nil, status.Error(codes.Unknown, "failed to get user")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.GetPassword()), []byte(in.GetPassword())); err != nil {
		return nil, status.Error(codes.PermissionDenied, "failed to compare hash and password")
	}

	log.Printf("Geted user: %s\n", user.GetLogin())
	return user, nil
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
	if err := godotenv.Load(".env"); err != nil {
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
