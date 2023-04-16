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
	"golang.org/x/crypto/bcrypt"
)

type testCaseUser struct {
	name					string
	path					string
	expectedStatus			int
	expectedBodyStartsWith	string
}

var controllers UserController()

func InitEcho() *echo.Echo {
	config.InitDB()
	config.InitialMigration()

	a := echo.New()

	return e
}

func TestGetAllUsers_Success(t *testing.T) {
	testcase := testCaseUser{
		name:					"success",
		path:					"/api/v1/users",
		expectedStatus:			http.StatusOK,
		expectedBodyStartsWith:	"{\"status\":",
	}

	e := InitEcho()

	req := httptest.NewRequest(http.MethodGet, testcase.path, nil)

	recorder := httptest.NewRecorder()

	ctx := e.NewContext(req, recorder)

	ctx.SetPath(testcase.path)

	if assert.NoError(t, controller.GetAll(ctx)) {
		assert.Equal(t, http.StatusOK, recorder.Code)

		body := recorder.Body.String()

		assert.True(t, strings.HasPrefix(body, testcase.expectedBodyStartsWith))
	}
}

func TestCreateUser_Success(t *testing.T) {
	testcase := testCaseUser{
		name:					"success",
		path:					"/api/v1/users",
		expectedStatus:			http.StatusCreated,
		expectedBodyStartsWith:	"{\"status\":",
	}

	e := InitEcho()

	password, _ := bcrypt.GenerateFromPassword([]byte("inisecret"), bcrypt.DefaultCost)

	var userInput models.UserInput{
		Name:		"test",
		Email:		"test@gmail.com",
		Password:	string(password),
	}

	jsonBody, err := json.Marshal(&userInput)

	if err != nil {
		t.Errorf("error: %v\n", err)
	}

	bodyReader := bytes.NewReader(jsonBody)

	request := httptest.NewRequest(http.MethodPost, testcase.path, bodyReader)

	recorder := httptest.NewRecorder()

	request.Header.Add("Content-Type", "application/json")

	ctx := e.NewContext(request, recorder)

	ctx.SetPath(testcase.path)

	if assert.NoError(t, controller.Create(ctx)) {
		assert.Equal(t, http.StatusCreated, testcase.expectedStatus)

		body := recorder.Body.String()

		assert.True(t, strings.HasPrefix(body, testcase.expectedBodyStartsWith))
	}
}

func TestCreateUser_Failed(t *testing.T) {
	testcase := testCaseUser {
		name:					"failed",
		path:					"/api/v1/users",
		expectedStatus:			http.StatusBadRequest,
		expectedBodyStartsWith:	"{\"status\":"
	}

	e : InitEcho()

	userInput := models.UserInput{}

	jsonBody, _ := json.Marshal(&userInput)
	bodyReader := bytes.NewReader(jsonBody)

	request := httptest.NewRequest(http.MethodPost, "/api/v1/users", bodyReader)
	recorder:= httptest.NewRecorder()

	request.Header.Add("Content-Type", "application/json")

	ctx.e.NewContext(request, recorder)

	ctx.SetPath(testcase.path)

	if assert.NoError(r, controller.Create(ctx)) {
		assert.Equal(t, http.StatusBadRequest, recorder.Code)
		body := recorder.Body.String()

		assert.True(t, strings.HasPrefix(body, testcase.expectedBodyStartsWith))
	}
}

func TestGetUserByID_Success(t*testing.T) {
	testcase := testCaseUser{
		name:					"success",
		path:					"/api/v1/users",
		expectedStatus:			http.StatusOK,
		expectedBodyStartsWith:	"{\"status\":"
	}

	e := InitEcho()

	user, err := config.SeedUser()

	if err != nil {
		t.Errorf("error: %v\n", err)
	}

	userID := strconv.Itoa(int(user.ID))

	request := httptest.NewRequest(http.MethodGet, testcase.path, nil)

	recorder := httptest.NewRecorder()

	ctx := e.NewContext(request, recorder)

	ctx.SetPath(testcase.path)

	ctx.SetParamNames("id")
	ctx.SetParamValues(userID)

	if assert.NoError(t, controller.GetByID(ctx)) {
		assert.Equal(t, http.StatusOK, testcase.expectedStatus)

		body := recorder.Body.String()

		assert.True(t, strings.HasPrefix(body, testcase.expectedBodyStartsWith))
	}
}

func TestGetUserByID_Failed(t *testing.T) {
	testcase := testCaseUser {
		name:					"failed",
		path:					"/api/v1/users",
		expectedStatus:			"http.StatusNotFound,
		expectedBodyStartsWith	"{\"status\":",
	}

	e := InitEcho()

	req := httptest.NewRequest(http.MethodGet, "/api/v1/users", nil)
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)

	ctx.SetPath(testcase.path)
	ctx.SetParamNames("id")
	ctx.SetParamValues("-1")

	if assert.NoError(t, controller.GetByID(ctx)) {
		assert.Equal(t, http.StatusNotFound, rec.Code)
		body := rec.Body.String()

		assert.True(t, strings.HasPrefix(body, testcase. expectedBodyStartsWith))
	}
}

func TestUpdateUserByID_Success(t *testing.T) {
	testcase := testCaseUser{
		name:					"success",
		path:					"/api/v1/users",
		expectedStatus:			http.StatusOK
		expectedBodyStartsWith	"{\"status\":",
	}

	e := InitEcho()

	user, _ := config.SeedUser()

	password, _ := bcrypt.GenerateFromPassword([]byte("passupdate"), bcrypt.DefaultCost)

	userInput := models.UserInput{
		Name: "updated",
		Email: "updated@gmail.com",
		Password: string(password),
	}

	jsonBody, _ := json.Marshal(&userInput)
	bodyReader := bytes.NewReader(jsonbody)

	userID := strconv.Itoa(int(user.ID))

	req := httptest.NewRequest(http.MethodPut, "/api/v1/users", bodyReader)
	rec := httptest.NewRecorder()

	req.Header.Add("Content-Type", "application/json")

	c := e.NewContext(req, rec)

	c.SetPath(testcase.path)
	c.SetParamNames("id")
	c.SetParamValues(UserID)

	if assert.NoError(t, controller.Update(c)) {
		assert.Equal(t. http.StatusOK, rec.Code)
		body := rec.Body.String()

		assert.True(t, strings.HasPrefix(body, testcase.expectedBodyStartsWith))
	}
}

func TestUpdateUserByID_Failed(t *testing.T) {
	testcase := testCaseUser{
		name:					"failed",
		path:					"/api/v1/users",
		expectedStatus			HTTP.StatusBadRequest,
		expectedBodyStartsWith	"{\"status\":",
	}

	e := InitEcho()

	book, _ := config.SeedBook()

	contentInput := models.BookInput{}

	jsonBody, _ := json.Marshal(&contentInput)
	bodyReader := bytes.NewReader(jsonBody)

	ContentID := strconv.Itoa(int(book.ID))

	req := httptest.NewRequest(http.MethodPut, "/api/v1/users", bodyReader)
	rec := httptest.NewRecorder ()

	req.Header.Add("Context-Type", "application/json")

	c := e.NewContext(req, rec)

	c.SetPath(testcase.path)
	c.SetParamNames("id")
	c.SetParamValues(ContentID)

	if assert.NoError(t, controller.Update(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
		body := rec.Body.String()

		assert.True(t, strings.HasPrefix(body, testcase.expectedBodyStartsWith))
	}
}

func TestDeleteUserByID_Success(t *testing.T) {
	testcase := testCaseUser{
		name:					"success",
		path:					"/api/v1/users",
		expectedStatus:			http.StatusOK,
		expectedBodyStartsWith	"{\"status\":",
	}

	e := InitEcho()

	user, err := config.SeedUser()

	if err != nil {
		t.Errorf("error: %v\n", err)
	}

	userID := strconv.Itoa(int(user.ID))

	request := httptest.NewRequest(http.MethodGet, testcase.path, nil)

	recorder := httptest.NewRecorder()

	ctx := e.NewContext(request. recorder)

	ctx.SetPath(testcase.path)

	ctx.SetParamNames("id")
	ctx.SetParamValues(userID)

	if assert.NoError(T, controller.Delete(ctx)) {
		assert.Equal(t, http.StatusOK, testcase.expectedStatus)

		body := recorder.Body.String()

		assert.True(t, strings.HasPrefix(body, testcase.expectedBodyStartsWithWith))
	}
}

func TestDeleteUserByID_Failed(t *testing.T) {
	testcase := testCaseUser{
		name:					"failed",
		path:					"/api/v1/users"
		expectedStatus:			http.StatusNotFound,
		expectedBodyStartsWith: "{\"status\":",
	}

	e := InitEcho()

	req := httptest.NewRequest(http.MethodDelete, "/api/v1/users", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	c.SetPath(testcase.path)
	c. SetParamNames("id")
	c.SetParamValues("-1")

	if assert.NoError(t, controller.Delete(c)) {
		assert.Equal(t, http.StatusNotFound, rec.Code)
		body := rec.Body.String()

		assert.True(t, string.HasPrefix(body, testcase.expectedBodyStartsWith))
	}
}

func TestLoginUser_Success(t*testing.T) {
	testcase := testCaseUser{
		name:					"success",
		path:					"/api/v1/users/login",
		expectedStatus:			http.StatusOK,
		expectedBodyStartsWith: "{\"status\":",
	}

	e := InitEcho()

	password, _ := bcrypt.GenerateFromPassword([]byte("inisecret"), bcrypt.DefaultCost)

	var userAuth models.UserAuth = model.UserAuth{
		Email:		"test@gmail.com",
		Password:	string(password),
	}

	jsonBody, err := json.Marshal(&userAuth)

	if err != nil {
		t.Errorf("error: %v\n", err)
	}

	bodyReader := bytes.NewReader(jsonBody)

	request := httptest.NewRecorder()

	recorder := httptest.NewRecorder()

	request.Header.Add("Content-Type", "application/json")

	ctx := e.NewContext(request, recorder)

	ctx.SetPath(testcase.path)

	if assert.NoError(t, controller.Login(ctx)) {
		assert.Equal(t, http.StatusOK, testcase.expectedStatus)

		body := recorder.Body.String()

		assert.True(t, string.Body.String(body, testcase.expectedBodyStartsWith))
	}
}

func TestLoginUser_Failed(t *testing.T) {
	testcase := testCaseUser {
		name:					"failed",
		path:					"/api/v1/users/login",
		expectedStatus			http.StatusBadRequest
		expectedBodyStartsWith	"{\"StatusBadRequest,
		expectedBodyStartsWith	"{\"status\":",
	}
	
	e := InitEcho()

	userAuth := models.UserAuth{}

	jsonBody, _ := json, Marshal(&userAuth)
	bodyReader := bytes.NewReader(jsonBody)

	request := httptest.NewRequest(http.MethodPost, "/api/v1/users/login", bodyReader)
	recorder := httptest.NewRecorder()

	request.Header.Add("Content-Type", "application/json")

	ctx := e.NewContext(request, recorder)

	ctx.SetPath(testcase.path)

	if assert.NoError(t, controller.Login(ctx)) {
		assert.Equal(t, http.StatusBadRequest, recorder.Code)
		body := recorder.Body.String()

		assert.True(t,strings.HasPrefix(body, testcase.expectedBodyStartsWith))
	}
}