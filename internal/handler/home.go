package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/MyongJi-StudentVote/StudentVote-Admin/internal/utils"
	"github.com/MyongJi-StudentVote/StudentVote-Admin/templates"
	"github.com/labstack/echo/v4"
)

type HomeHandler struct {
	Title   string
	BaseURL string
	Client  *http.Client
}

type SigninDTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SigninResponseDTO struct {
	Data struct {
		AccessToken  string `json:"accessToken"`
		RefreshToken string `json:"refreshToken"`
		Role         string `json:"role"`
		TokenType    string `json:"tokenType"`
	} `json:"data"`
}

func (h *HomeHandler) Home(ctx echo.Context) error {
	home := templates.Home(h.Title)
	return utils.Render(home, ctx, http.StatusOK)
}

func (h *HomeHandler) HandleSignin(ctx echo.Context) error {
	var response SigninResponseDTO
	email := ctx.FormValue("email")
	password := ctx.FormValue("password")
	dto := &SigninDTO{Email: email, Password: password}
	bodyBytes, _ := json.Marshal(dto)
	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf(h.BaseURL, "auth/sign-in"), bytes.NewBuffer(bodyBytes))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := h.Client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusOK {
		responseBody, _ := io.ReadAll(resp.Body)
		err := json.Unmarshal(responseBody, &response)
		if err != nil {
			return err
		}
		cookie := new(http.Cookie)
		cookie.Name = "token"
		cookie.Value = response.Data.AccessToken
		cookie.Expires = time.Now().Add(24 * time.Hour)
		cookie.HttpOnly = true
		ctx.SetCookie(cookie)
		ctx.Response().Header().Add("HX-Redirect", "/main-home")
	} else {
		errBody, _ := io.ReadAll(resp.Body)
		log.Println(string(errBody))
	}
	return nil
}

func InitHomeHandler(client *http.Client, baseURL string) *HomeHandler {
	return &HomeHandler{Title: "Home", Client: client, BaseURL: baseURL}
}
