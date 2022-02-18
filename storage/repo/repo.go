package repo

import (
	pb "github.com/toshkentov01/alif-tech-task/user-service/genproto/user-service"
)

// UserRepository holds two interfaces
type UserRepository interface {
	Reader
	Writer
}

// Reader interface for selecting data
type Reader interface {
	CheckFields(username, email string) (*pb.CheckfieldsResponse, error)
	GetBalance(userID string) (*pb.GetBalanceResponse, error)
	ListTotalOperationsByType(operationType, userID string) (*pb.ListTotalOperationsByTypeResponse, error)
	CheckUserType(userID string) (*pb.CheckUserTypeResponse, error)
}

// Writer interface for inserting data
type Writer interface {
	CreateUnIdentifiedUser(username, password string) (*pb.CreateUnIdentifiedUserResponse, error)
	CreateIdentifiedUser(user *pb.CreateIdentifiedUserRequest) (*pb.CreateIdentifiedUserResponse, error)
	Income(userID string, incomeAmount int64) error
	Expense(userID string, expenseAmount int64) error
}