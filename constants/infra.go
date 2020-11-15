package constants

// message db connection.
const (
	ConnectDBSuccess    string = "Connected to DB"
	ConnectRedisSuccess string = "Connected to Redis"

	ConnectDBFail    string = "Could not connect database, error"
	ConnectRedisFail string = "Could not connect redis, error"

	ClosingDBSuccess string = "Database conn gracefully close"
	ClosingDBFailed  string = "Error closing DB connection"

	Success string = "success"
	Fail    string = "fail"

	DBTimeLayout       string = "2006-01-02 15:04:05"
	ResponseTimeLayout string = "2006-01-02T15:04:05-0700"
)
