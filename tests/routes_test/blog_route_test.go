package routes_test

import (
	"encoding/json"
	"fmt"
	"github.com/GapaiID/SE-challenge2/api/controllers"
	"github.com/GapaiID/SE-challenge2/api/dto"
	"github.com/GapaiID/SE-challenge2/api/routes"
	"github.com/GapaiID/SE-challenge2/lib"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"net/http"
	"testing"
)

type BlogRouterSuite struct {
	suite.Suite
}

func (suite *BlogRouterSuite) Ts() *TestServer {
	controller := controllers.NewBlogController(NewBlogServiceMock(), NewBlogPolicyMock())
	router := routes.NewBlogRouter(lib.NewHttpHandler(), controller)
	router.Setup()
	return NewTestServer(suite.T(), router.Handler.Engine)
}

func (suite *BlogRouterSuite) TestList() {
	statusCode, _, body := suite.Ts().Get(suite.T(), "/blog_posts")

	var resp struct {
		Data dto.BlogPostPagination `json:"data"`
	}
	if err := json.Unmarshal([]byte(body), &resp); err != nil {
		suite.Fail(err.Error())
	}

	suite.Equal(http.StatusOK, statusCode)
	suite.Equal(2, len(resp.Data.List))
}

func (suite *BlogRouterSuite) TestDetail() {
	tests := []struct {
		Name   string
		ID     int
		Status int
	}{
		{Name: "Success get detail", ID: 1, Status: http.StatusOK},
		{Name: "Not found", ID: 2, Status: http.StatusNotFound},
	}

	for _, test := range tests {
		suite.T().Run(test.Name, func(t *testing.T) {
			statusCode, _, body := suite.Ts().Get(suite.T(), fmt.Sprintf("/blog_posts/%d", test.ID))
			assert.Equal(t, test.Status, statusCode)

			if test.Status == http.StatusOK {
				var resp struct {
					Data dto.BlogPost `json:"data"`
				}
				if err := json.Unmarshal([]byte(body), &resp); err != nil {
					assert.Fail(t, err.Error())
				}
			}
		})
	}
}

func (suite *BlogRouterSuite) TestCreate() {
}

func (suite *BlogRouterSuite) TestUpdate() {
}

func (suite *BlogRouterSuite) TestDelete() {
}

func TestBlogRouterSuite(t *testing.T) {
	suite.Run(t, new(BlogRouterSuite))
}
