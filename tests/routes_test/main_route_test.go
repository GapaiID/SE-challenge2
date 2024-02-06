package routes_test

import (
	"github.com/agung96tm/miblog/api/controllers"
	"github.com/agung96tm/miblog/api/routes"
	"github.com/agung96tm/miblog/lib"
	"github.com/stretchr/testify/suite"
	"net/http"
	"testing"
)

type MainRouterSuite struct {
	suite.Suite
}

func (suite *MainRouterSuite) TsWithRedis(redisMock *LibRedisMock) *TestServer {
	controller := controllers.NewMainController(redisMock)
	router := routes.NewMainRouter(lib.NewHttpHandler(), controller, lib.Config{})
	router.Setup()
	return NewTestServer(suite.T(), router.Handler.Engine)
}

func (suite *MainRouterSuite) TestSuccessGetWelcome() {
	redis := NewLibRedisMock()

	ts := suite.TsWithRedis(redis)
	statusCode, _, _ := ts.Get(suite.T(), "/")

	suite.Equal(statusCode, http.StatusOK)
}

func (suite *MainRouterSuite) TestSuccessCalledSetCache() {
	redis := NewLibRedisMock()

	ts := suite.TsWithRedis(redis)
	_, _, _ = ts.Get(suite.T(), "/")

	suite.NotEmpty(redis.IsCalled["Set"])
}

func (suite *MainRouterSuite) TestSuccessCalledGetCache() {
	redis := NewLibRedisMock()

	ts := suite.TsWithRedis(redis)
	_, _, _ = ts.Get(suite.T(), "/")

	suite.NotEmpty(redis.IsCalled["Get"])
}

func TestMainRouterSuite(t *testing.T) {
	suite.Run(t, new(MainRouterSuite))
}
