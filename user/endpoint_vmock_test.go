package user_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/podhmo-sandbox/go-mock-sandbox/user"
	"github.com/podhmo-sandbox/go-mock-sandbox/user/vmock"
)

func TestUserServiceEndpointVMock(t *testing.T) {
	id := "1"
	service := &vmock.UserService{
		GetFunc: func(id string) (*user.User, error) {
			return &user.User{ID: id}, nil
		},
	}
	ep := &user.UserEndpoint{
		Service: service,
	}

	req := httptest.NewRequest("GET", fmt.Sprintf("http://example.com/users/%s", id), nil)
	w := httptest.NewRecorder()

	ep.GetUser(w, req)
	res := w.Result()
	if res.StatusCode != http.StatusOK {
		t.Errorf("expected status is %d, but %d", 200, res.StatusCode)
	}

	// spy (mock)
	if !service.GetCalled() {
		t.Error("Get() is must be called, but not called")
	}

	if service.GetCalls()[0].Id != id {
		t.Errorf("passed id is must be %s, but %s", id, service.GetCalls()[0].Id)
	}
}
