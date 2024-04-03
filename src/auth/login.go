package auth

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/alexedwards/argon2id"
	"github.com/labstack/echo/v4"
	"iamricky.com/truck-rental/db"
)

func LoginRoute(c echo.Context) error {
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
	token, err := login(c, name, password)
	if err != nil {
		resp.Message = fmt.Sprintf("%s", err)
		jsonResp, err := json.Marshal(resp)
		if err != nil {
			return c.String(http.StatusOK, fmt.Sprintf("%s", err))
		}
		return c.String(http.StatusOK, string(jsonResp))
	} else if token == "" {
		resp.Message = "login process failed"
		jsonResp, err := json.Marshal(resp)
		if err != nil {
			return c.String(http.StatusOK, fmt.Sprintf("%s", err))
		}
		return c.String(http.StatusOK, string(jsonResp))
	}
	resp.Profile = User{Name: name, Token: token}
	resp.Message = "user " + name + " logged in"
	jsonResp, err := json.Marshal(resp)
	if err != nil {
		return c.String(http.StatusOK, fmt.Sprintf("%s", err))
	}
	return c.String(http.StatusOK, string(jsonResp))
}

func login(c echo.Context, name string, password string) (string, error) {
	token := GenerateSecureToken()
	now := time.Now()
	tokenExpires := now.AddDate(0, 0, 7)
	expString := tokenExpires.Format(time.DateTime)
	conn, err := db.GetConn()
	if err != nil {
		fmt.Println("fail to connect")
		return "", err
	}
	q := "SELECT password FROM users WHERE name = ?;"
	var hashedPassword string
	err = conn.QueryRowContext(c.Request().Context(), q, name).Scan(&hashedPassword)
	if err == sql.ErrNoRows {
		fmt.Println("query returned no rows")
		return "", err
	} else if err != nil {
		fmt.Println("fail to query")
		return "", err
	}
	match, err := argon2id.ComparePasswordAndHash(password, hashedPassword)
	if err != nil {
		fmt.Println("fail to query")
		return "", err
	}
	if !match {
		fmt.Println("incorrect password")
		return "", errors.New("incorrect password")
	}
	q = "UPDATE users SET token = ?, token_exp = ? WHERE name = ?;"
	_, err = conn.ExecContext(c.Request().Context(), q, token, expString, name)
	if err != nil {
		fmt.Println("fail to exec")
		return "", err
	}
	return token, nil
}
