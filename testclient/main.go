package main

import (
	"context"
	"log"

	pb "github.com/daniilmikhaylov2005/school_project/api"
	"github.com/golang-jwt/jwt/v4"
	"google.golang.org/grpc"
)

//const address = "localhost:50000"

const address = "localhost:50001"

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("failed to conn: %v", err)
	}
	defer conn.Close()

	// client := pb.NewUserClient(conn)

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

	// u2, err := client.GetUser(context.Background(), &pb.GetUserRequest{Login: "sadrezdev", Password: "123123"})
	// if err != nil {
	// 	log.Fatalf("failed to get user: %v", err)
	// }

	// log.Println(u2)

	// token := jwt.NewWithClaims(jwt.SigningMethodHS256, &AuthClaims{
	// 	StandardClaims: jwt.StandardClaims{
	// 		ExpiresAt: time.Now().Add(time.Hour * 10).Unix(),
	// 		IssuedAt:  time.Now().Unix(),
	// 	},
	// 	Teacher_Login: "sadrezdev",
	// })

	// jwtKey, err := token.SignedString([]byte("jojo"))
	// if err != nil {
	// 	log.Fatalf("failed to sign token: %v", err)
	// }

	// client := pb.NewMagazineClient(conn)
	// response, err := client.CreateClass(context.Background(), &pb.CreateClassRequest{
	// 	Children: []*pb.Kid{
	// 		{
	// 			Fullname: "Алексеева Валерия",
	// 			Age:      17,
	// 		},
	// 		{
	// 			Fullname: "Соколова Василиса",
	// 			Age:      17,
	// 		},
	// 		{
	// 			Fullname: "Власов Евгений",
	// 			Age:      18,
	// 		},
	// 		{
	// 			Fullname: "Лазарев Павел",
	// 			Age:      17,
	// 		},
	// 		{
	// 			Fullname: "Королева Анастасия",
	// 			Age:      17,
	// 		},
	// 		{
	// 			Fullname: "Ильина Мария",
	// 			Age:      16,
	// 		},
	// 		{
	// 			Fullname: "Симонова Алиса",
	// 			Age:      18,
	// 		},
	// 		{
	// 			Fullname: "Никитин Павел",
	// 			Age:      17,
	// 		},
	// 		{
	// 			Fullname: "Андреев Максим",
	// 			Age:      18,
	// 		},
	// 		{
	// 			Fullname: "Смирнова Маргарита",
	// 			Age:      16,
	// 		},
	// 	},
	// 	TeacherLogin: "sadrezdev",
	// 	Jwt:          jwtKey,
	// 	Graduate:     11,
	// })

	// log.Printf("Magazine code: %d\n", response.GetMagazineCode())

	// res2, err := client.GetClass(context.Background(), &pb.GetClassRequest{
	// 	MagazineCode: response.GetMagazineCode(),
	// })
	// if err != nil {
	// 	log.Fatalf("failed to get class: %v", err)
	// }

	// log.Printf("Teacher full name: %s\n", res2.GetTeacherFullname())

	// for i, v := range res2.Children {
	// 	log.Printf("%d) Kid id = %d\n\tKid fullname = %s\n\tKid age = %d\n\tKid graduate = %d", i, v.GetId(), v.GetFullname(), v.GetAge(), v.GetGraduate())
	// }

	client := pb.NewMagazineClient(conn)
	r3, err := client.GetClassGrades(context.Background(), &pb.GetClassGradesRequest{
		MagazineCode: 568,
	})

	if err != nil {
		log.Fatalf("failed to get grades: %v", err)
	}

	log.Println("Magazine code", r3.GetMagazineCode())

	for _, v := range r3.GetChildrenGrades() {
		log.Println("Kid info: ", v.GetKid())
		log.Println("Grades: ", v.GetGrades())
	}

}

type AuthClaims struct {
	jwt.StandardClaims
	Teacher_Login string
}
