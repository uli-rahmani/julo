package constants

// Query comments.
const (
	QueryGetWalletsByCustomerUUID string = `
		SELECT 
			wallet_id, wallet_uuid, customer_uuid, status, balance
		FROM wallets
		WHERE
			customer_uuid = ?
	`

	QueryGetWalletsByWalletID string = `
		SELECT 
			wallet_uuid, customer_uuid, status, enable_at, balance
		FROM wallets
		WHERE
			wallet_id = ?
	`

	QueryGetDisableWalletsByWalletID string = `
		SELECT 
			wallet_uuid, customer_uuid, status, disabled_at, balance
		FROM wallets
		WHERE
			wallet_id = ?
	`

	QueryInsertWallet string = `
		INSERT INTO wallets 
			(wallet_uuid, customer_uuid, status, created_at, balance)
		VALUES 
			(?, ?, 'created', NOW(), 0)
	`

	QueryEnableWallet string = `
		UPDATE wallets SET 
			status = ?,
			updated_at = ?,
			enable_at = ?
		WHERE wallet_id = ?
	`

	QueryDisableWallet string = `
		UPDATE wallets SET 
			status = ?,
			updated_at = ?,
			disabled_at = ?
		WHERE wallet_id = ?
	`

	QueryAddWalletBalance string = `
		UPDATE wallets SET 
			balance = balance + ?,
			updated_at = NOW()
		WHERE wallet_id = ?
	`

	QueryDecreaseWalletBalance string = `
		UPDATE wallets SET 
			balance = balance - ?,
			updated_at = NOW()
		WHERE wallet_id = ?
	`
)

const (
	QueryIsExistTransactionByReferenceID string = `
		SELECT EXISTS
		(
			SELECT 
				wallet_transaction_id
			FROM
				wallet_transactions
			WHERE
				wallet_id = ? AND reference_id = ? AND types = ?
		)
	`

	QueryInsertWalletTransaction string = `
		INSERT INTO wallet_transactions 
			(wallet_transaction_uuid, wallet_id, status, amount, reference_id, types, created_at)
		VALUES 
			(?, ?, 'success', ?, ?, ?, NOW())
	`
)
