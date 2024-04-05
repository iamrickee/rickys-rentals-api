package location

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"iamricky.com/truck-rental/db"
)

type Location struct {
	Id      int    `json:"id"`
	Address string `json:"address"`
	City    string `json:"city"`
	State   string `json:"state"`
	Zip     string `json:"zip"`
}

type LocationResp struct {
	Data    Location `json:"data"`
	Message string   `json:"message"`
}

type SuccessResp struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

func SaveRoute(c echo.Context) error {
	location, err := save(c)
	if err != nil {
		return err
	}
	resp := LocationResp{Data: location, Message: "location saved"}
	jsonResp, err := json.Marshal(resp)
	if err != nil {
		return c.String(http.StatusOK, fmt.Sprintf("%s", err))
	}
	return c.String(http.StatusOK, string(jsonResp))
}

func DeleteRoute(c echo.Context) error {
	err := delete(c)
	if err != nil {
		return err
	}
	resp := SuccessResp{Success: true, Message: "location deleted"}
	jsonResp, err := json.Marshal(resp)
	if err != nil {
		return c.String(http.StatusOK, fmt.Sprintf("%s", err))
	}
	return c.String(http.StatusOK, string(jsonResp))
}

func save(c echo.Context) (Location, error) {
	id := c.FormValue("id")
	address := c.FormValue("address")
	city := c.FormValue("city")
	state := c.FormValue("state")
	zip := c.FormValue("zip")
	conn, err := db.GetConn()
	if err != nil {
		fmt.Println("fail to connect")
		return Location{}, err
	}
	idInt, err := strconv.Atoi(id)
	if err == nil {
		q := "UPDATE locations SET address = ?, city = ?, state = ?, zip = ? WHERE id = ?"
		_, err := conn.ExecContext(c.Request().Context(), q, address, city, state, zip, idInt)
		if err != nil {
			return Location{}, err
		}
		return get(c, idInt)
	} else {
		q := "INSERT INTO locations (address, city, state, zip) VALUES (?,?,?,?);"
		_, err := conn.ExecContext(c.Request().Context(), q, address, city, state, zip)
		if err != nil {
			fmt.Printf("address: %s city: %s, state: %s, zip: %s\n", address, city, state, zip)
			fmt.Println(err)
			return Location{}, err
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
		q := "DELETE FROM locations WHERE id = ?"
		_, err := conn.ExecContext(c.Request().Context(), q, idInt)
		return err
	} else {
		return err
	}
}

func get(c echo.Context, id int) (Location, error) {
	conn, err := db.GetConn()
	if err != nil {
		fmt.Println("fail to connect")
		return Location{}, err
	}
	q := "SELECT id, address, city, state, zip FROM locations WHERE id = ?"
	l := Location{}
	err = conn.QueryRowContext(c.Request().Context(), q, id).Scan(&l.Id, &l.Address, &l.City, &l.State, &l.Zip)
	if err == sql.ErrNoRows {
		fmt.Println("query returned no rows")
		return Location{}, err
	} else if err != nil {
		fmt.Println("fail to query")
		return Location{}, err
	}
	return l, nil
}

func getLast(c echo.Context) (Location, error) {
	conn, err := db.GetConn()
	if err != nil {
		fmt.Println("fail to connect")
		return Location{}, err
	}
	q := "SELECT id, address, city, state, zip FROM locations ORDER BY id DESC;"
	l := Location{}
	err = conn.QueryRowContext(c.Request().Context(), q).Scan(&l.Id, &l.Address, &l.City, &l.State, &l.Zip)
	if err == sql.ErrNoRows {
		fmt.Println("query returned no rows")
		return Location{}, err
	} else if err != nil {
		fmt.Println("fail to query")
		return Location{}, err
	}
	return l, nil
}
