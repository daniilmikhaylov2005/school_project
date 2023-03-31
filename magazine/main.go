package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	pb "github.com/daniilmikhaylov2005/school_project/api"
	"github.com/daniilmikhaylov2005/school_project/magazine/repository"
	"github.com/daniilmikhaylov2005/school_project/magazine/transport"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
)

func main() {
	port := flag.String("port", "50001", "port for service")
	jwtSecret := flag.String("jwt", "jojo", "string for jwt-secret")
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
	MagazineServer := transport.NewMagazineServer(*jwtSecret, repository)
	pb.RegisterMagazineServer(s, MagazineServer)
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
