package login

type LoginFailReason string

const (
	AuthenticationFailed LoginFailReason = "Failed to authenticate user with provided information"
	HashGenerationFailed                 = "Failed to generate password hash"
	Unknown                              = "Unknown error during authentication"
)

type LoginResult struct {
	success bool
	reason  LoginFailReason
}

func LoginSuccess() LoginResult {
	return LoginResult{success: true}
}

func LoginFailed(reason LoginFailReason) LoginResult {
	return LoginResult{success: false, reason: reason}
}

func (l LoginResult) Success() bool {
	return l.success
}

func (l LoginResult) Failed() LoginFailReason {
	return l.reason
}
