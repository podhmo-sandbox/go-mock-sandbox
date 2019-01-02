package user

// UserService :
//go:generate mocker --pkg vmock --dst ./vmock/$GOFILE $GOFILE UserService
type UserService interface {
	Get(id string) (*User, error)
}
