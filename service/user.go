package service

import (
	"context"
	"log"

	newerrors "github.com/toshkentov01/alif-tech-task/user-service/new_errors"
	l "github.com/toshkentov01/alif-tech-task/user-service/pkg/logger"
	"github.com/toshkentov01/alif-tech-task/user-service/storage"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "github.com/toshkentov01/alif-tech-task/user-service/genproto/user-service"
)

// UserService ...
type UserService struct {
	storage storage.I
	logger  l.Logger
}

// NewUserService ...
func NewUserService(log l.Logger) *UserService {
	return &UserService{
		storage: storage.NewStorage(),
		logger:  log,
	}
}

// I didn't check errors in some service methods, because errorHandler which I've written checks itself.
// This enables my work

// CheckFields checks user's username and email for existence
func (s *UserService) CheckFields(ctx context.Context, request *pb.CheckfieldsRequest) (*pb.CheckfieldsResponse, error) {
	result, err := s.storage.User().CheckFields(request.Username, request.Email)
	log.Println(request.Email)
	return result, errorHandler(s.logger, err, "falied to check fields")
}

// CreateUnIdentifiedUser method creates unidentified users
func (s *UserService) CreateUnIdentifiedUser(ctx context.Context, request *pb.CreateUnIdentifiedUserRequest) (*pb.Empty, error) {
	err := s.storage.User().CreateUnIdentifiedUser(request.Id, request.Username, request.Password, request.AccessToken, request.RefreshToken)

	return &pb.Empty{}, errorHandler(s.logger, err, "failed to create unidentified user")
}

// CreateIdentifiedUser method creates identified users
func (s *UserService) CreateIdentifiedUser(ctx context.Context, request *pb.CreateIdentifiedUserRequest) (*pb.Empty, error) {
	err := s.storage.User().CreateIdentifiedUser(request)

	return &pb.Empty{}, errorHandler(s.logger, err, "failed to create identified user")
}

// GetBalance method gets user's balance
func (s *UserService) GetBalance(ctx context.Context, request *pb.GetBalanceRequest) (*pb.GetBalanceResponse, error) {
	result, err := s.storage.User().GetBalance(request.UserId)

	return result, errorHandler(s.logger, err, "failed to get balance")
}

// ListTotalOperationsByType method lists total operations of an user by income or expense types
func (s *UserService) ListTotalOperationsByType(ctx context.Context, request *pb.ListTotalOperationsByTypeRequest) (*pb.ListTotalOperationsByTypeResponse, error) {
	result, err := s.storage.User().ListTotalOperationsByType(request.OperationType, request.UserId)

	return result, errorHandler(s.logger, err, "failed to list user operations")
}

// Income method
func (s *UserService) Income(ctx context.Context, request *pb.IncomeRequest) (*pb.Empty, error) {
	err := s.storage.User().Income(request.UserId, request.IncomeAmount)
	if err == newerrors.ErrMaximumAmount {
		return nil, status.Error(codes.PermissionDenied, "Permission denied")

	} else if err != nil {
		return nil, err
	}

	return &pb.Empty{}, nil
}

// Expense method
func (s *UserService) Expense(ctx context.Context, request *pb.ExpenseRequest) (*pb.Empty, error) {
	err := s.storage.User().Expense(request.UserId, request.ExpenseAmount)
	if err == newerrors.ErrNotEnoughCash {
		return nil, status.Error(codes.PermissionDenied, "Permission denied")

	} else if err != nil {
		return nil, err
	}

	return &pb.Empty{}, nil
}

// CheckUserType method checks whether user is identified or not
func (s *UserService) CheckUserType(ctx context.Context, request *pb.CheckUserTypeRequest) (*pb.CheckUserTypeResponse, error) {
	result, err := s.storage.User().CheckUserType(request.UserId)

	return result, errorHandler(s.logger, err, "failed to check user type")
}

// CheckUserAccount method checks whether user has an account or not
func (s *UserService) CheckUserAccount(ctx context.Context, request *pb.CheckUserAccountRequest) (*pb.CheckUserAccountResponse, error) {
	result, err := s.storage.User().CheckUserAccount(request.Username, request.Password)

	if err != nil {
		s.logger.Error("Error while checking user existence"+err.Error())
		return nil, err
	}
	
	return result, nil
}