package main

import (
	"fmt"
)

// UserDomainError represents errors specific to the User domain
type UserDomainError struct {
	Code    string
	Message string
}

func (err *UserDomainError) Error() string {
	return fmt.Sprintf("Error %s: %s", err.Code, err.Message)
}

type Code string

const (
	ErrUserNotRegistered    Code = "user_not_registered"
	ErrUserAlreadyActivated Code = "user_already_activated"
)

func (c Code) String() string {
	return string(c)
}

func NewUserDomainError(code Code, message string) *UserDomainError {
	return &UserDomainError{Code: string(code), Message: message}
}

type User struct {
	State string `json:"state"`
}

const (
	UserStateRegistered = "registered"
	UserStateActivated  = "activated"
	UserStateInvalid    = "invalid"
)

func (u *User) Register() error {
	if u.State != UserStateInvalid {
		return NewUserDomainError(ErrUserAlreadyActivated, "User already activated")
	}
	u.State = UserStateRegistered
	return nil
}

func (u *User) Activate() error {
	if u.State != UserStateRegistered {
		return NewUserDomainError(ErrUserNotRegistered, "User must be registered first")
	}
	u.State = UserStateActivated
	return nil
}

func main() {
	user := &User{State: UserStateInvalid}

	err := user.Activate()
	if err != nil {
		switch err.(type) {
		case *UserDomainError:
			domainErr := err.(*UserDomainError)
			switch domainErr.Code {
			case string(ErrUserNotRegistered):
				fmt.Println("Error: User must be registered first.")
			case string(ErrUserAlreadyActivated):
				fmt.Println("Error: User already activated.")
			default:
				fmt.Println("Unknown domain error:", domainErr)
			}
		default:
			fmt.Println("Unexpected error:", err)
		}
	} else {
		fmt.Println("User activated successfully.")
	}

	err = user.Register()
	if err != nil {
		fmt.Println(err)
	}

	err = user.Activate()
	if err != nil {
		fmt.Println(err)
	}
}
