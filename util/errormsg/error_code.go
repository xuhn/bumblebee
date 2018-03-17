package errormsg

import (
	"fmt"
)

var (
	USER_NOT_EXIST error = fmt.Errorf("user not exist")
	PASSWORD_INCORRECT error = fmt.Errorf("password incorrect")
	SERVICE_ERROR error = fmt.Errorf("user not exist")

	CREATE_USER_ERROR error = fmt.Errorf("create user error")
)
