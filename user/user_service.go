package user

// UserService :
//go:generate mocker --pkg vmock --dst ./vmock/$GOFILE --prefix "" $GOFILE UserService
//go:generate mockery -name UserService -outpkg vmock -output ./tmock 
type UserService interface {
	Get(id string) (*User, error)
}
