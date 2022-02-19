package sqls

const (
	// CheckfieldsSQL ...
	CheckfieldsSQL = `
		SELECT 
			(SELECT EXISTS (SELECT username FROM users WHERE username = $1)),
			(SELECT EXISTS (SELECT email FROM users WHERE email = $2))
		FROM users
	`

	// CreateUnIdentifiedUserSQL ...
	CreateUnIdentifiedUserSQL = `
		INSERT INTO users(id, username, password, access_token, refresh_token) VALUES ($1, $2, $3, $4, $5)
	`

	// CreateIdentifedUserSQL ...
	CreateIdentifedUserSQL = `
		INSERT INTO users(
			id, 
			username,
			full_name,
			email,
			password,
			is_identified,
			access_token,
			refresh_token) 
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
	`
	// CheckUserTypeSQL ...
	CheckUserTypeSQL = `
		SELECT
			is_identified
		FROM
			users
		WHERE id = $1
	`

	// TopUpUserBalanceSQL ...
	TopUpUserBalanceSQL = `
		UPDATE identifed_user_accounts
		SET
			balance = balance + $1,
			updated_at = NOW()
		WHERE user_id = $2 AND (SELECT iua.balance + $1 FROM identifed_user_accounts AS iua WHERE iua.user_id = $2) < 100000
	`

	// ReduceBalanceSQL ...
	ReduceBalanceSQL = `
		UPDATE identifed_user_accounts
		SET
			balance = balance - $1,
			updated_at = NOW()
		WHERE user_id = $2 AND (SELECT iua.balance - $1 FROM identifed_user_accounts AS iua WHERE iua.user_id = $2) >= 0
	`

	// TopUpUserBalanceSQL ...
	TopUpUnidentifiedUserBalanceSQL = `
		UPDATE unidentifed_user_accounts
		SET
			balance = balance + $1,
			updated_at = NOW()
		WHERE user_id = $2 AND (SELECT uniua.balance + $1 FROM unidentifed_user_accounts AS uniua WHERE uniua.user_id = $2) < 10000
	`

	// ReduceBalanceSQL ...
	ReduceUnidentifiedUserBalanceSQL = `
		UPDATE unidentifed_user_accounts
		SET
			balance = balance - $1,
			updated_at = NOW()
		WHERE user_id = $2 AND (SELECT uniua.balance - $1 FROM unidentifed_user_accounts AS uniua WHERE uniua.user_id = $2) >= 0
	`

	// IncreaseCashSQL ...
	IncreaseCashSQL = `
		INSERT INTO cash_controller(user_id, income_amount) VALUES ($1, $2)
	`

	// ReduceCashSQL ...
	ReduceCashSQL = `
		INSERT INTO cash_controller(user_id, expense_amount) VALUES ($1, $2)
	`

	// GetBalanceSQL ...
	GetIdentifiedUserBalanceSQL = `
		SELECT balance FROM identifed_user_accounts
		WHERE user_id = $1
	`

	// GetUnidentifiedUserBalanceSQL ...
	GetUnidentifiedUserBalanceSQL = `
		SELECT balance FROM unidentifed_user_accounts
		WHERE user_id = $1
	`

	
	// ListTotalTopUpOperationsOfIdentifedUserSQL ...
	ListTotalTopUpOperationsSQL = `
		SELECT
			income_amount,
			created_at
		FROM cash_controller
		WHERE user_id = $1 AND (extract(year from cash_controller.created_at) = $2 AND (extract(month from cash_controller.created_at) = $3))
			AND income_amount != 0
	`

	ListTotalExpenseOperationsSQL = `
		SELECT
			expense_amount,
			created_at
		FROM cash_controller
		WHERE user_id = $1 AND (extract(year from cash_controller.created_at) = $2 AND (extract(month from cash_controller.created_at) = $3))
			AND expense_amount != 0
	`

)