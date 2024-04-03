package auth

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/alexedwards/argon2id"
	"github.com/labstack/echo/v4"
	"iamricky.com/truck-rental/db"
)

func RegisterRoute(c echo.Context) error {
	name := c.FormValue("username")
	resp := Resp{}
	if name == "" {
		resp.Message = "username is required"
		jsonResp, err := json.Marshal(resp)
		if err != nil {
			return c.String(http.StatusOK, fmt.Sprintf("%s", err))
		}
		return c.String(http.StatusOK, string(jsonResp))
	}
	password := c.FormValue("password")
	if password == "" {
		resp.Message = "password is required"
		jsonResp, err := json.Marshal(resp)
		if err != nil {
			return c.String(http.StatusOK, fmt.Sprintf("%s", err))
		}
		return c.String(http.StatusOK, string(jsonResp))
	}
	passwordConf := c.FormValue("password_confirm")
	if passwordConf == "" {
		resp.Message = "password confirmation is required"
		jsonResp, err := json.Marshal(resp)
		if err != nil {
			return c.String(http.StatusOK, fmt.Sprintf("%s", err))
		}
		return c.String(http.StatusOK, string(jsonResp))
	}
	if passwordConf != password {
		resp.Message = "password confirmation does not match"
		jsonResp, err := json.Marshal(resp)
		if err != nil {
			return c.String(http.StatusOK, fmt.Sprintf("%s", err))
		}
		return c.String(http.StatusOK, string(jsonResp))
	}
	token, err := register(c, name, password)
	if err != nil {
		resp.Message = fmt.Sprintf("%s", err)
		jsonResp, err := json.Marshal(resp)
		if err != nil {
			return c.String(http.StatusOK, fmt.Sprintf("%s", err))
		}
		return c.String(http.StatusOK, string(jsonResp))
	}
	resp.Profile = User{Name: name, Token: token}
	resp.Message = "user " + name + " created"
	jsonResp, err := json.Marshal(resp)
	if err != nil {
		return c.String(http.StatusOK, fmt.Sprintf("%s", err))
	}
	return c.String(http.StatusOK, string(jsonResp))
}

func register(c echo.Context, name string, password string) (string, error) {
	hashed, err := argon2id.CreateHash(password, argon2id.DefaultParams)
	if err != nil {
		fmt.Println(err)
	}
	token := GenerateSecureToken()
	now := time.Now()
	tokenExpires := now.AddDate(0, 0, 7)
	expString := tokenExpires.Format(time.DateTime)
	conn, err := db.GetConn()
	if err != nil {
		fmt.Println("fail to connect")
		return "", err
	}
	q := "INSERT INTO users (name, password, token, token_exp) VALUES (?,?,?,?);"
	_, err = conn.ExecContext(c.Request().Context(), q, name, hashed, token, expString)
	if err != nil {
		fmt.Println("fail to exec")
		return "", err
	}
	return token, nil
}
