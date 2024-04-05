package auth

import (
	"crypto/rand"
	"database/sql"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"iamricky.com/truck-rental/db"
)

func TokenRoute(c echo.Context) error {
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
	token := c.FormValue("token")
	if token == "" {
		resp.Message = "token is required"
		jsonResp, err := json.Marshal(resp)
		if err != nil {
			return c.String(http.StatusOK, fmt.Sprintf("%s", err))
		}
		return c.String(http.StatusOK, string(jsonResp))
	}
	err := TokenLogin(c, name, token)
	if err != nil {
		resp.Message = fmt.Sprintf("%s", err)
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

func TokenLogin(c echo.Context, name string, token string) error {
	conn, err := db.GetConn()
	if err != nil {
		fmt.Println("fail to connect")
		return err
	}
	q := "SELECT token, token_exp FROM users WHERE name = ?;"
	var tokenStored string
	var tokenExpires string
	err = conn.QueryRowContext(c.Request().Context(), q, name).Scan(&tokenStored, &tokenExpires)
	if err == sql.ErrNoRows {
		fmt.Println("query returned no rows")
		return err
	} else if err != nil {
		fmt.Println("fail to query")
		return err
	}
	now := time.Now()
	expires, err := time.Parse(time.DateTime, tokenExpires)
	if err != nil {
		fmt.Println("invalid token expiration date received from database")
		return err
	}
	if now.After(expires) {
		fmt.Println("token expired")
		return errors.New("token expired")
	}
	if token != tokenStored {
		fmt.Println("token doesn't match")
		return errors.New("token doesn't match")
	}
	return nil
}

func GenerateSecureToken() string {
	b := make([]byte, 16)
	if _, err := rand.Read(b); err != nil {
		return ""
	}
	return hex.EncodeToString(b)
}
