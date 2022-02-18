package service

// go test ./service -v -bench . -cpuprofile prof.out
// go tool pprof prof.out

import (
	"context"
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/toshkentov01/alif-tech-task/user-service/config"
	pb "github.com/toshkentov01/alif-tech-task/user-service/genproto/user-service"
	"github.com/toshkentov01/alif-tech-task/user-service/pkg/logger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
)

var (
	loggerTest  logger.Logger
	cfg         *config.Config
	userService *UserService
)

func init() {
	// for field path you should your path where alif-tech-task project located
	const path = "/home/sardor/go/src/github.com/sardortoshkentov/alif-tech-task/user-service/.env"
	if info, err := os.Stat(path); !os.IsNotExist(err) {
		if !info.IsDir() {
			godotenv.Load(path)
			if err != nil {
				fmt.Println("Err:", err)
			}
		}
	} else {
		fmt.Println("Not exists")
	}
	cfg = config.Get()
	loggerTest = logger.New(cfg.LogLevel, "user_service")
}

const bufSize = 1024 * 1024

var lis *bufconn.Listener

func TestConnectionDatabase(t *testing.T) {

	loggerTest.Info("main: sqlxConfig",
		logger.String("host", cfg.PostgresHost),
		logger.Int("port", cfg.PostgresPort),
		logger.String("database", cfg.PostgresDatabase))

	lis = bufconn.Listen(bufSize)
	server := grpc.NewServer()
	userService = NewUserService(loggerTest)
	pb.RegisterUserServiceServer(server, userService)

	go func() {
		if err := server.Serve(lis); err != nil {
			t.Errorf("Error while listening: %s", err)
		}
	}()
}

func TestCheckFields(t *testing.T) {
	type testCase struct {
		username         string
		email            string
		usernameExpected bool
		emailExpected    bool
	}

	testCases := []testCase{{
		username:         "sardor",
		email:            "testemail@mail.ru",
		usernameExpected: false,
		emailExpected:    false,
	}, {
		username:         "sardortoshkentov",
		email:            "testemail2@mail.ru",
		usernameExpected: false,
		emailExpected:    false,
	},
	}

	for _, testCase := range testCases {

		t.Run(fmt.Sprintf("%s, %s", testCase.username, testCase.email), func(t *testing.T) {
			res, err := userService.CheckFields(context.Background(), &pb.CheckfieldsRequest{})

			if err != nil {
				t.Errorf("Error while check field, err: %s", err)
				t.Errorf(
					"got: %v for username and %v for email. wanted: %v for username and %v for email",
					res.UsernameExists, res.EmailExists, testCase.usernameExpected, testCase.emailExpected,
				)
			}
		})

	}
}

func TestCreateIdentifiedUser(t *testing.T) {
	username, email, fullName, password := "testusername", "testemail@gmail.com", "testFullName", "sardor@11T"
	accessToken, refreshToken := "testAccess", "testRefresh"

	result, err := userService.CreateIdentifiedUser(context.Background(), &pb.CreateIdentifiedUserRequest{
		Username:     username,
		FullName:     fullName,
		Email:        email,
		Password:     password,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	})

	if err != nil {
		t.Errorf("Error while creating identified user, err: %v", err.Error())
	}

	log.Println(result.Id)
}
