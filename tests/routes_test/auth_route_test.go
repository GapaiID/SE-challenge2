package routes_test

import (
	"github.com/GapaiID/SE-challenge2/api/controllers"
	"github.com/GapaiID/SE-challenge2/api/dto"
	"github.com/GapaiID/SE-challenge2/api/routes"
	"github.com/GapaiID/SE-challenge2/lib"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"net/http"
	"testing"
)

type AuthRouterSuite struct {
	suite.Suite
}

func (suite *AuthRouterSuite) Ts() *TestServer {
	controller := controllers.NewAuthController(NewAuthServiceMock())
	router := routes.NewAuthRouter(lib.NewHttpHandler(), controller)
	router.Setup()
	return NewTestServer(suite.T(), router.Handler.Engine)
}

func (suite *AuthRouterSuite) TestLogin() {
	tests := []struct {
		Name     string
		Email    string
		Password string
		Status   int
	}{
		{Name: "Valid Data", Email: "valid@example.com", Password: "p@ssw0rd123", Status: http.StatusOK},
		{Name: "Invalid email: status bad request", Email: "invalid@example.com", Password: "p@ssw0rd123", Status: http.StatusBadRequest},
		{Name: "Validate email: status bad request", Email: "", Password: "p@ssw0rd123", Status: http.StatusBadRequest},
	}

	for _, test := range tests {
		suite.T().Run(test.Name, func(t *testing.T) {
			statusCode, _, _ := suite.Ts().Post(suite.T(), "/auth/login", dto.LoginRequest{
				Email:    test.Email,
				Password: test.Password,
			})
			assert.Equal(t, test.Status, statusCode)
		})
	}
}

func (suite *AuthRouterSuite) TestRegister() {
	tests := []struct {
		Name       string
		ActualName string
		Email      string
		Password   string
		Status     int
	}{
		{Name: "Valid Data", ActualName: "valid name", Email: "valid@example.com", Password: "p@ssw0rd123", Status: http.StatusOK},
		{Name: "Invalid email: status bad request", ActualName: "invalid name", Email: "invalid@example.com", Password: "p@ssw0rd123", Status: http.StatusBadRequest},
		{Name: "Validate email: status bad request", ActualName: "bad name", Email: "", Password: "p@ssw0rd123", Status: http.StatusBadRequest},
	}

	for _, test := range tests {
		suite.T().Run(test.Name, func(t *testing.T) {
			statusCode, _, _ := suite.Ts().Post(suite.T(), "/auth/login", dto.RegisterRequest{
				Email:    test.Email,
				Password: test.Password,
				Name:     test.ActualName,
			})
			assert.Equal(t, test.Status, statusCode)
		})
	}
}

func TestAuthRouterSuite(t *testing.T) {
	suite.Run(t, new(AuthRouterSuite))
}
