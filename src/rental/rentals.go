package rental

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"iamricky.com/truck-rental/db"
)

type RentalsResp struct {
	Data    []Rental `json:"data"`
	Message string   `json:"message"`
}

func ListRoute(c echo.Context) error {
	rentals, err := getList(c)
	if err != nil {
		return c.String(http.StatusOK, fmt.Sprintf("%s", err))
	}
	resp := RentalsResp{Data: rentals, Message: "rentals fetched"}
	jsonResp, err := json.Marshal(resp)
	if err != nil {
		return c.String(http.StatusOK, fmt.Sprintf("%s", err))
	}
	return c.String(http.StatusOK, string(jsonResp))
}

func getList(c echo.Context) ([]Rental, error) {
	conn, err := db.GetConn()
	if err != nil {
		fmt.Println("fail to connect")
		return nil, err
	}
	q := "SELECT id, name, description FROM rentals;"
	rows, err := conn.QueryContext(c.Request().Context(), q)
	if err != nil {
		fmt.Println("fail to query")
		return nil, err
	}
	result := []Rental{}
	for rows.Next() {
		var r Rental
		err = rows.Scan(&r.Id, &r.Name, &r.Description)
		if err != nil {
			fmt.Println("invalid query row data")
			return nil, err
		}
		result = append(result, r)
	}
	return result, nil
}
