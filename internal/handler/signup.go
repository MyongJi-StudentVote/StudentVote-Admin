package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/MyongJi-StudentVote/StudentVote-Admin/internal/utils"
	"github.com/MyongJi-StudentVote/StudentVote-Admin/templates"
	"github.com/labstack/echo/v4"
)

type SignupHandler struct {
	Title   string
	BaseURL string
	Client  *http.Client
}

type SignupDTO struct {
	Name        string `json:"name"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	GovernaceId int    `json:"governanceId"`
}

func (s *SignupHandler) Signup(ctx echo.Context) error {
	signup := templates.Signup(s.Title)
	return utils.Render(signup, ctx, http.StatusOK)
}

func (s *SignupHandler) HandleSignup(ctx echo.Context) error {
	name := ctx.FormValue("name")
	email := ctx.FormValue("email")
	password := ctx.FormValue("password")
	governanceId, _ := strconv.Atoi(ctx.FormValue("unit"))
	body := &SignupDTO{
		Name:        name,
		Email:       email,
		Password:    password,
		GovernaceId: governanceId,
	}
	bodyByte, err := json.Marshal(body)
	if err != nil {
		return err
	}
	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf(s.BaseURL, "auth/sign-up"), bytes.NewBuffer(bodyByte))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := s.Client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusOK {
		return utils.Render(templates.SignupSuccess(), ctx, http.StatusAccepted)
	} else {
		errBody, _ := io.ReadAll(resp.Body)
		log.Println(string(errBody))
	}
	return nil
}

func InitSignupHandler(client *http.Client, baseURL string) *SignupHandler {
	return &SignupHandler{Title: "회원가입", Client: client, BaseURL: baseURL}
}
