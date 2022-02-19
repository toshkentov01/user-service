package postgres

import (
	"database/sql"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	pb "github.com/toshkentov01/alif-tech-task/user-service/genproto/user-service"
	newerrors "github.com/toshkentov01/alif-tech-task/user-service/new_errors"
	"github.com/toshkentov01/alif-tech-task/user-service/pkg/helper"
	"github.com/toshkentov01/alif-tech-task/user-service/platform/postgres"
	"github.com/toshkentov01/alif-tech-task/user-service/storage/repo"
	"github.com/toshkentov01/alif-tech-task/user-service/storage/sqls"
)

type userRepo struct {
	db *sqlx.DB
}

// NewUserRepo for generating new user repository
func NewUserRepo() repo.UserRepository {
	return &userRepo{
		db: postgres.DB(),
	}
}

// I wrote added sqls folder inside storage folder where my queries located.
// It makes my code more readable

// CheckFields method for checking whether email or username with signup-wanted user's email or username exists or not
func (ur *userRepo) CheckFields(username, email string) (*pb.CheckfieldsResponse, error) {
	var (
		usernameExists, emailExists sql.NullBool
	)

	err := ur.db.QueryRow(sqls.CheckfieldsSQL, username, email).Scan(&usernameExists, &emailExists)
	if err == sql.ErrNoRows {
		usernameExists.Bool = false
		emailExists.Bool = false

	} else if err != nil {
		log.Println("Error while checking user existence. ERROR: ", err.Error())
		return nil, err
	}

	return &pb.CheckfieldsResponse{
		UsernameExists: usernameExists.Bool,
		EmailExists:    emailExists.Bool,
	}, nil
}

// CreateUnIdentifiedUser method creates unidentified users
func (ur *userRepo) CreateUnIdentifiedUser(id, username, password, accessToken, refreshToken string) error {
	_, err := ur.db.Exec(sqls.CreateUnIdentifiedUserSQL, id, username, password, accessToken, refreshToken)
	if err != nil {
		if strings.Contains(err.Error(), `unique constraint "users_username_key"`) {
			return newerrors.ErrUsernameExists

		} else if err != nil {
			log.Println("Error while creating unidentified user. ERROR: ", err.Error())
			return err
		}
	}

	return nil
}

// CreateIdentifiedUser method creates identified users
func (ur *userRepo) CreateIdentifiedUser(user *pb.CreateIdentifiedUserRequest) error {

	_, err := ur.db.Exec(
		sqls.CreateIdentifedUserSQL,
		user.Id,
		user.Username, user.FullName,
		user.Email, user.Password, true,
		user.AccessToken, user.RefreshToken,
	)

	if err != nil {
		return err
	}

	return nil
}

// CheckUserType method checks whether user is identified or not
func (ur *userRepo) CheckUserType(userID string) (*pb.CheckUserTypeResponse, error) {
	var (
		identified sql.NullBool
	)

	err := ur.db.QueryRow(sqls.CheckUserTypeSQL, userID).Scan(&identified)
	if err != nil {
		log.Println("Error while checking user type")
		return nil, err
	}

	return &pb.CheckUserTypeResponse{
		Identified: identified.Bool,
	}, nil
}

// Income for toping-up user's balance
func (ur *userRepo) Income(userID string, incomeAmount int64) error {
	// transaction begins
	tx, err := ur.db.Begin()
	if err != nil {
		log.Println("Error beginning transaction! ")
		return err
	}

	result, err := ur.CheckUserType(userID)
	if err != nil {
		log.Println("Error while checking user type")
		return err
	}

	if result.Identified {
		rowsResult, err := tx.Exec(sqls.TopUpUserBalanceSQL, incomeAmount, userID)
		if err != nil {
			tx.Rollback()

			log.Println("Error while topping up an user balance, ERROR: ", err.Error())
			return err
		}

		rows, err := rowsResult.RowsAffected()
		if err != nil {
			return err
		}

		if rows == 0 {
			return newerrors.ErrMaximumAmount
		}

		_, err = tx.Exec(sqls.IncreaseCashSQL, userID, incomeAmount)
		if err != nil {
			tx.Rollback()

			log.Println("Error while increasing a cash")
			return err
		}

		tx.Commit()

	} else if !result.Identified {
		rowsResult, err := tx.Exec(sqls.TopUpUnidentifiedUserBalanceSQL, incomeAmount, userID)
		if err != nil {
			tx.Rollback()

			log.Println("Error while topping up an unidentified user balance, ERROR: ", err.Error())
			return err
		}

		rows, err := rowsResult.RowsAffected()
		if err != nil {
			return err
		}

		if rows == 0 {
			return newerrors.ErrMaximumAmount
		}

		_, err = tx.Exec(sqls.IncreaseCashSQL, userID, incomeAmount)
		if err != nil {
			tx.Rollback()

			log.Println("Error while increasing a cash")
			return err
		}

		tx.Commit()
	}

	return nil
}

// Expense for reducing user's balance
func (ur *userRepo) Expense(userID string, expenseAmount int64) error {
	// transaction begins
	tx, err := ur.db.Begin()
	if err != nil {
		log.Println("Error beginning transaction! ")
		return err
	}

	result, err := ur.CheckUserType(userID)
	if err != nil {
		log.Println("Error while checking user type")
		return err
	}

	if result.Identified {
		rowsResult, err := tx.Exec(sqls.ReduceBalanceSQL, expenseAmount, userID)
		if err != nil {
			tx.Rollback()

			log.Println("Error while reducing an user balance, ERROR: ", err.Error())
			return err
		}

		rows, err := rowsResult.RowsAffected()
		if err != nil {
			return err
		}

		if rows == 0 {
			return newerrors.ErrNotEnoughCash
		}

		_, err = tx.Exec(sqls.ReduceCashSQL, userID, expenseAmount)
		if err != nil {
			tx.Rollback()

			log.Println("Error while reducing a cash")
			return err
		}

		tx.Commit()

	} else if !result.Identified {
		rowsResult, err := tx.Exec(sqls.ReduceUnidentifiedUserBalanceSQL, expenseAmount, userID)
		if err != nil {
			tx.Rollback()

			log.Println("Error while reducing an unidentified user balance, ERROR: ", err.Error())
			return err
		}

		rows, err := rowsResult.RowsAffected()
		if err != nil {
			return err
		}

		if rows == 0 {
			return newerrors.ErrNotEnoughCash
		}

		_, err = tx.Exec(sqls.ReduceCashSQL, userID, expenseAmount)
		if err != nil {
			tx.Rollback()

			log.Println("Error while reducing a cash")
			return err
		}

		tx.Commit()
	}

	return nil
}

// GetBalance method for getting a user balance
func (ur *userRepo) GetBalance(userID string) (*pb.GetBalanceResponse, error) {
	var (
		balance sql.NullInt64
	)

	result, err := ur.CheckUserType(userID)
	if err != nil {
		log.Println("Error while checking user type")
		return nil, err
	}

	if result.Identified {
		err := ur.db.QueryRow(sqls.GetIdentifiedUserBalanceSQL, userID).Scan(&balance)
		if err != nil {
			log.Println("Error while getting identified user's balance")
			return nil, err
		}

	} else if !result.Identified {
		err := ur.db.QueryRow(sqls.GetUnidentifiedUserBalanceSQL, userID).Scan(&balance)
		if err != nil {
			log.Println("Error while getting unidentified user's balance")
			return nil, err
		}
	}

	return &pb.GetBalanceResponse{
		Balance: balance.Int64,
	}, nil
}

// ListTotalOperationsByType method for getting a user's operations for current month by type. Type can be a income operations or expense operations
func (ur *userRepo) ListTotalOperationsByType(operationType, userID string) (*pb.ListTotalOperationsByTypeResponse, error) {
	results := make([]*pb.Operations, 0, 30)
	now := time.Now()

	year := strconv.Itoa(now.Year())
	month := time.Now().Month().String()

	convertedMonth := helper.MonthConverter(month)

	if operationType == "income_operations" {
		rows, err := ur.db.Query(sqls.ListTotalTopUpOperationsSQL, userID, year, convertedMonth)

		if err != nil {
			log.Println("Error while getting income operations, ERROR: ", err.Error())
			return nil, err
		}

		defer rows.Close()

		for rows.Next() {
			result := pb.Operations{}
			var (
				data            sql.NullString
				operationAmount sql.NullInt64
			)

			action := ``

			err := rows.Scan(
				&operationAmount,
				&data,
			)

			if err != nil {
				return nil, err
			}

			operationAmountStr := strconv.Itoa(int(operationAmount.Int64))
			action += `+` + operationAmountStr + ``

			result.Action = action
			result.Date = data.String

			results = append(results, &result)
		}

	} else if operationType == "expense_operations" {
		rows, err := ur.db.Query(sqls.ListTotalExpenseOperationsSQL, userID, year, convertedMonth)

		if err != nil {
			log.Println("Error while getting expense operations, ERROR: ", err.Error())
			return nil, err
		}

		defer rows.Close()

		for rows.Next() {
			result := pb.Operations{}
			var (
				data            sql.NullString
				operationAmount sql.NullInt64
			)

			action := ``

			err := rows.Scan(
				&operationAmount,
				&data,
			)

			if err != nil {
				return nil, err
			}

			operationAmountStr := strconv.Itoa(int(operationAmount.Int64))
			action += `-` + operationAmountStr + ``

			result.Action = action
			result.Date = data.String

			results = append(results, &result)
		}
	}

	return &pb.ListTotalOperationsByTypeResponse{
		Results: results,
		Count:   int32(len(results)),
	}, nil
}

// CheckUserAccount method checks whether user has an account or not
func (ur *userRepo) CheckUserAccount(username, password string) (*pb.CheckUserAccountResponse, error) {
	var (
		exists sql.NullBool
	)

	err := ur.db.QueryRow(`
		SELECT EXISTS(SELECT id FROM users WHERE username = $1 AND password = crypt($2, password))
	`, username, password).Scan(&exists)

	if err == sql.ErrNoRows {
		exists.Bool = false

	} else if err != nil {
		return nil, err
	}

	return &pb.CheckUserAccountResponse{
		Exists: exists.Bool,
	}, nil
}
