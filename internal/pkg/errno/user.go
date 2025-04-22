package errno

var (
	ErrUserAlreadyExist = &Errno{HTTP: 400, Code: "FailedOperation.UserAlreadyExist", Message: "User already exist."}

	ErrUserNotFound = &Errno{HTTP: 404, Code: "ResourceNotFound.UserNotFound", Message: "User was not found."}

	ErrPasswordIncorrect = &Errno{HTTP: 401, Code: "InvalidParameter.PasswordIncorrect", Message: "Password was incorrect."}
)
