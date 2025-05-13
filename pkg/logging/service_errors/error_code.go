package service_errors

const (
	// Token
	UnExpectedError = "Expected error"
	ClaimsNotFound  = "Claims not found"
	TokenRequired   = "Token required"
	TokenExpired    = "Token expired"
	TokenInvalid    = "Token invalid"
	TokenMalformed  = "Token malformed"

	// OTP
	OtpExists   = "OTP exists"
	OtpUsed     = "OTP used"
	OtpNotValid = "OTP not valid"

	// User
	EmailExists      = "Email exists"
	UsernameExists   = "Username exists"
	PermissionDenied = "Permission denied"

	// DB
	RecordNotFound = "record not found"
)
