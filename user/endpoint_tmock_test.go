package user_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/podhmo-sandbox/go-mock-sandbox/user"
	"github.com/podhmo-sandbox/go-mock-sandbox/user/tmock"
	"github.com/stretchr/testify/mock"
)

func TestUserServiceEndpointTMock(t *testing.T) {
	id := "1"
	service := &tmock.UserService{}

	ep := &user.UserEndpoint{
		Service: service,
	}
	service.On("Get", mock.MatchedBy(func(passedID string) bool {
		return id == passedID
	})).Return(&user.User{ID: id}, nil)

	req := httptest.NewRequest("GET", fmt.Sprintf("http://example.com/users/%s", id), nil)
	w := httptest.NewRecorder()

	ep.GetUser(w, req)
	res := w.Result()
	if res.StatusCode != http.StatusOK {
		t.Errorf("expected status is %d, but %d", 200, res.StatusCode)
	}
}
