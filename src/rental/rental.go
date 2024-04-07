package rental

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"iamricky.com/truck-rental/db"
)

type Rental struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Image       string `json:"image"`
}

type RentalResp struct {
	Data    Rental `json:"data"`
	Message string `json:"message"`
}

type SuccessResp struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

func SaveRoute(c echo.Context) error {
	rental, err := save(c)
	if err != nil {
		return c.String(http.StatusOK, fmt.Sprintf("%s", err))
	}
	resp := RentalResp{Data: rental, Message: "rental saved"}
	jsonResp, err := json.Marshal(resp)
	if err != nil {
		return c.String(http.StatusOK, fmt.Sprintf("%s", err))
	}
	return c.String(http.StatusOK, string(jsonResp))
}

func DeleteRoute(c echo.Context) error {
	err := delete(c)
	if err != nil {
		return c.String(http.StatusOK, fmt.Sprintf("%s", err))
	}
	resp := SuccessResp{Success: true, Message: "rental deleted"}
	jsonResp, err := json.Marshal(resp)
	if err != nil {
		return c.String(http.StatusOK, fmt.Sprintf("%s", err))
	}
	return c.String(http.StatusOK, string(jsonResp))
}

func GetRoute(c echo.Context) error {
	id := c.FormValue("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return c.String(http.StatusOK, fmt.Sprintf("%s", err))
	}
	rental, err := get(c, idInt)
	if err != nil {
		return c.String(http.StatusOK, fmt.Sprintf("%s", err))
	}
	resp := RentalResp{Data: rental, Message: "rental loaded"}
	jsonResp, err := json.Marshal(resp)
	if err != nil {
		return c.String(http.StatusOK, fmt.Sprintf("%s", err))
	}
	return c.String(http.StatusOK, string(jsonResp))
}

func save(c echo.Context) (Rental, error) {
	id := c.FormValue("id")
	name := c.FormValue("name")
	description := c.FormValue("description")
	image := c.FormValue("image")
	conn, err := db.GetConn()
	if err != nil {
		fmt.Println("fail to connect")
		return Rental{}, err
	}
	idInt, err := strconv.Atoi(id)
	if err == nil {
		q := "UPDATE rentals SET name = ?, description = ?, image = ? WHERE id = ?"
		_, err := conn.ExecContext(c.Request().Context(), q, name, description, image, idInt)
		if err != nil {
			return Rental{}, err
		}
		return get(c, idInt)
	} else {
		q := "INSERT INTO rentals (name, description, image) VALUES (?,?,?);"
		_, err := conn.ExecContext(c.Request().Context(), q, name, description, image)
		if err != nil {
			fmt.Printf("name: %s description: %s image: %s\n", name, description, image)
			fmt.Println(err)
			return Rental{}, err
		}
		return getLast(c)
	}
}

func delete(c echo.Context) error {
	id := c.FormValue("id")
	conn, err := db.GetConn()
	if err != nil {
		fmt.Println("fail to connect")
		return err
	}
	idInt, err := strconv.Atoi(id)
	if err == nil {
		q := "DELETE FROM rentals WHERE id = ?"
		_, err := conn.ExecContext(c.Request().Context(), q, idInt)
		return err
	} else {
		return err
	}
}

func get(c echo.Context, id int) (Rental, error) {
	conn, err := db.GetConn()
	if err != nil {
		fmt.Println("fail to connect")
		return Rental{}, err
	}
	q := "SELECT id, name, description, image FROM rentals WHERE id = ?"
	r := Rental{}
	err = conn.QueryRowContext(c.Request().Context(), q, id).Scan(&r.Id, &r.Name, &r.Description, &r.Image)
	if err == sql.ErrNoRows {
		fmt.Println("query returned no rows")
		return Rental{}, err
	} else if err != nil {
		fmt.Println("fail to query")
		return Rental{}, err
	}
	return r, nil
}

func getLast(c echo.Context) (Rental, error) {
	conn, err := db.GetConn()
	if err != nil {
		fmt.Println("fail to connect")
		return Rental{}, err
	}
	q := "SELECT id, name, description, image FROM rentals ORDER BY id DESC;"
	r := Rental{}
	err = conn.QueryRowContext(c.Request().Context(), q).Scan(&r.Id, &r.Name, &r.Description, &r.Image)
	if err == sql.ErrNoRows {
		fmt.Println("query returned no rows")
		return Rental{}, err
	} else if err != nil {
		fmt.Println("fail to query")
		return Rental{}, err
	}
	return r, nil
}
