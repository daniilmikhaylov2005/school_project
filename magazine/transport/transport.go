package transport

import (
	"context"
	"errors"
	"fmt"

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
		return nil, status.Error(codes.PermissionDenied, "invalid jwt")
	}

	if parsedJwt != in.GetTeacherLogin() {
		return nil, status.Error(codes.PermissionDenied, "invalid jwt data or login")
	}

	magazineCode, err := s.repository.CreateClass(in.GetChildren(), in.GetTeacherLogin(), in.GetGraduate())
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to create class")
	}
	return &pb.CreateClassResponse{MagazineCode: int64(magazineCode)}, nil
}
