package location

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"iamricky.com/truck-rental/db"
)

type LocationsResp struct {
	Data    []Location `json:"data"`
	Message string     `json:"message"`
}

func ListRoute(c echo.Context) error {
	locations, err := getList(c)
	if err != nil {
		return c.String(http.StatusOK, fmt.Sprintf("%s", err))
	}
	resp := LocationsResp{Data: locations, Message: "locations fetched"}
	jsonResp, err := json.Marshal(resp)
	if err != nil {
		return c.String(http.StatusOK, fmt.Sprintf("%s", err))
	}
	return c.String(http.StatusOK, string(jsonResp))
}

func getList(c echo.Context) ([]Location, error) {
	conn, err := db.GetConn()
	if err != nil {
		fmt.Println("fail to connect")
		return nil, err
	}
	q := "SELECT id, address, city, state, zip FROM locations ORDER BY id DESC;"
	rows, err := conn.QueryContext(c.Request().Context(), q)
	if err != nil {
		fmt.Println("fail to query")
		return nil, err
	}
	result := []Location{}
	for rows.Next() {
		var l Location
		err = rows.Scan(&l.Id, &l.Address, &l.City, &l.State, &l.Zip)
		if err != nil {
			fmt.Println("invalid query row data")
			return nil, err
		}
		result = append(result, l)
	}
	return result, nil
}
