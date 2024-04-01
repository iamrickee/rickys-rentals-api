package auth

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"iamricky.com/truck-rental/crypt"
	"iamricky.com/truck-rental/db"
)

type Resp struct {
	Success bool
	Message string
}

func Register(c echo.Context) error {
	name := c.FormValue("username")
	if name == "" {
		resp := Resp{
			false,
			"username is required",
		}
		jsonResp, err := json.Marshal(resp)
		if err != nil {
			return c.String(http.StatusOK, fmt.Sprintf("%s", err))
		}
		return c.String(http.StatusOK, string(jsonResp))
	}
	password := c.FormValue("password")
	if password == "" {
		resp := Resp{
			false,
			"password is required",
		}
		jsonResp, err := json.Marshal(resp)
		if err != nil {
			return c.String(http.StatusOK, fmt.Sprintf("%s", err))
		}
		return c.String(http.StatusOK, string(jsonResp))
	}
	passwordConf := c.FormValue("password_confirm")
	if passwordConf == "" {
		resp := Resp{
			false,
			"password confirmation is required",
		}
		jsonResp, err := json.Marshal(resp)
		if err != nil {
			return c.String(http.StatusOK, fmt.Sprintf("%s", err))
		}
		return c.String(http.StatusOK, string(jsonResp))
	}
	if passwordConf != password {
		resp := Resp{
			false,
			"password confirmation does not match",
		}
		jsonResp, err := json.Marshal(resp)
		if err != nil {
			return c.String(http.StatusOK, fmt.Sprintf("%s", err))
		}
		return c.String(http.StatusOK, string(jsonResp))
	}
	resp := Resp{
		true,
		"user " + name + " created",
	}
	jsonResp, err := json.Marshal(resp)
	if err != nil {
		return c.String(http.StatusOK, fmt.Sprintf("%s", err))
	}
	encryptedPW, err := crypt.EncryptPW(password)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("password: " + password + " encrypted: " + encryptedPW)
	decryptedPW, err := crypt.DecryptPW(encryptedPW)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("decrypted: " + decryptedPW)
	/*
		_, err = register(c, name, password)
		if err != nil {
			resp := Resp{
				false,
				fmt.Sprintf("%s", err),
			}
			jsonResp, err := json.Marshal(resp)
			if err != nil {
				return c.String(http.StatusOK, fmt.Sprintf("%s", err))
			}
			return c.String(http.StatusOK, string(jsonResp))
		}
	*/
	return c.String(http.StatusOK, string(jsonResp))
}

func register(c echo.Context, name string, password string) (string, error) {
	conn, err := db.GetConn()
	if err != nil {
		fmt.Println("fail to connect")
		return "", err
	}
	q := "INSERT INTO users (name, password) VALUES (?,?);"
	_, err = conn.ExecContext(c.Request().Context(), q, name, password)
	if err != nil {
		fmt.Println("fail to exec")
		return "", err
	}
	return "token", nil
}
