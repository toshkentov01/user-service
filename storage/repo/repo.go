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
	CreateUnIdentifiedUser(id, username, password, accessToken, refreshToken string)  error
	CreateIdentifiedUser(user *pb.CreateIdentifiedUserRequest)  error
	Income(userID string, incomeAmount int64) error
	Expense(userID string, expenseAmount int64) error
	CheckUserAccount(username, password string) (*pb.CheckUserAccountResponse, error)
}