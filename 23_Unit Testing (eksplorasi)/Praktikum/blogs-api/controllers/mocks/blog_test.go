package controllers

import (
	"blogs-api/config"
	"blogs-api/models"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

type testCaseBlog struct {
	name					string
	path					string
	expectedStatus			int
	expectedBodyStartsWith	string
}

var blogController blogController = InitBlogController()

func InitBlogEcho() *echo.Echo {
	config.InitDB()
	config.InitialMigration()

	e := echo.New()

	return e
}

func TestGetAllBlogs_Success(t *testing.T) {
	testcase := testCaseBlog{
		name:					"success"
		path:					"/api/v1/blogs"
		expectedStatus:			http.StatusOK,
		expectedBodyStartsWith:	"{\"status\":",
	}

	e := InitBlogEcho()

	req := httptest.NewRequest(http.MethodGet, testcase.path, nil)

	recorder := httptest.NewRecorder()

	ctx := e.NewContext(req, recorder)

	ctx.SetPath(testcase.path)

	if assert.NoError(t, blogsController.GetAll(cttx)) {
		assert.Equal(t, http.StatusOK, recorder.Code)

		body := recorder.Body.String()


		assert.True(t, strings.HasPrefix(body, testcase.expectedBodyStartsWith))
	}
}