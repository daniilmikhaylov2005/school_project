package transport

import (
	"context"
	"errors"
	"fmt"
	"log"

	pb "github.com/daniilmikhaylov2005/school_project/api"
	"github.com/daniilmikhaylov2005/school_project/magazine/repository"
	"github.com/golang-jwt/jwt/v4"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var ErrInvalidAccessToken = errors.New("invalid access token")

type AuthClaims struct {
	jwt.StandardClaims
	Teacher_Login string
}

func ParseToken(accessToken string, signingKey []byte) (string, error) {
	token, err := jwt.ParseWithClaims(accessToken, &AuthClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return signingKey, nil
	})

	if err != nil {
		return "", err
	}

	if claims, ok := token.Claims.(*AuthClaims); ok && token.Valid {
		return claims.Teacher_Login, nil
	}

	return "", ErrInvalidAccessToken
}

type MagazineServer struct {
	jwtSecret  string
	repository *repository.Repository
	pb.UnimplementedMagazineServer
}

func NewMagazineServer(jwtSecret string, repository *repository.Repository) *MagazineServer {
	return &MagazineServer{
		jwtSecret:  jwtSecret,
		repository: repository,
	}
}

func (s *MagazineServer) CreateClass(ctx context.Context, in *pb.CreateClassRequest) (*pb.CreateClassResponse, error) {
	parsedJwt, err := ParseToken(in.GetJwt(), []byte(s.jwtSecret))
	if err != nil {
		log.Printf("[error] %v\n", err)
		return nil, status.Error(codes.PermissionDenied, "invalid jwt")
	}

	if parsedJwt != in.GetTeacherLogin() {
		return nil, status.Error(codes.PermissionDenied, "invalid jwt data or login")
	}

	magazineCode, err := s.repository.CreateClass(in.GetChildren(), in.GetTeacherLogin(), in.GetGraduate())
	if err != nil {
		log.Printf("[error] %v\n", err)
		return nil, status.Error(codes.Unknown, "failed to create class")
	}
	return &pb.CreateClassResponse{MagazineCode: int64(magazineCode)}, nil
}

func (s *MagazineServer) GetClass(ctx context.Context, in *pb.GetClassRequest) (*pb.GetClassResponse, error) {
	class, err := s.repository.GetClass(in.GetMagazineCode())
	if err != nil {
		log.Printf("[error] %v\n", err)
		return nil, status.Error(codes.Unknown, "failed to get class")
	}
	return class, nil
}

func (s *MagazineServer) GetClassGrades(ctx context.Context, in *pb.GetClassGradesRequest) (*pb.GetClassGradesResponse, error) {
	childrenGrades, err := s.repository.GetClassGrades(in.GetMagazineCode())
	if err != nil {
		log.Printf("[error] %v\n", err)
		if errors.Is(err, repository.ErrNotFound) {
			return nil, status.Error(codes.NotFound, "grades of children not found with this code")
		}
		return nil, status.Error(codes.Unknown, "failed to get grades of children")
	}
	return childrenGrades, nil
}

func (s *MagazineServer) CreateGrade(ctx context.Context, in *pb.CreateGradeRequest) (*pb.CreateGradeResponse, error) {
	teacher_login, err := ParseToken(in.GetJwt(), []byte(s.jwtSecret))
	if err != nil {
		log.Printf("[error] %v\n", err)
		if errors.Is(err, ErrInvalidAccessToken) {
			return nil, status.Error(codes.PermissionDenied, "wrong jwt")
		}
		return nil, status.Error(codes.Unknown, "invalid jwt")
	}
	err = s.repository.CreateTeacherHistory(teacher_login, fmt.Sprintf("graded kid by %d", in.GetGrade().GetGrade()))
	if err != nil {
		log.Printf("[error] %v\n", err)
		return nil, status.Error(codes.Unknown, "failed to create teacher history")
	}

	err = s.repository.CreateGrade(in.GetKidId(), in.GetGrade())
	if err != nil {
		log.Printf("[error] %v\n", err)
		return nil, status.Error(codes.Unknown, "failed to create grade")
	}

	return &pb.CreateGradeResponse{
		Status: "created",
	}, nil
}
