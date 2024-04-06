package lead

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"iamricky.com/truck-rental/db"
	"iamricky.com/truck-rental/location"
	"iamricky.com/truck-rental/rental"
)

type LeadsResp struct {
	Data    []Lead `json:"data"`
	Message string `json:"message"`
}

func ListRoute(c echo.Context) error {
	leads, err := getList(c)
	if err != nil {
		return c.String(http.StatusOK, fmt.Sprintf("%s", err))
	}
	resp := LeadsResp{Data: leads, Message: "leads fetched"}
	jsonResp, err := json.Marshal(resp)
	if err != nil {
		return c.String(http.StatusOK, fmt.Sprintf("%s", err))
	}
	return c.String(http.StatusOK, string(jsonResp))
}

func getList(c echo.Context) ([]Lead, error) {
	conn, err := db.GetConn()
	if err != nil {
		fmt.Println("fail to connect")
		return nil, err
	}
	q := "SELECT a.id, a.first_name, a.last_name, a.email, a.phone, b.id AS bid, b.city, b.state, c.id AS cid, c.name FROM leads AS a LEFT JOIN locations AS b ON a.location_id = b.id LEFT JOIN rentals AS c ON a.rental_id = c.id ORDER BY a.id DESC;"
	rows, err := conn.QueryContext(c.Request().Context(), q)
	if err != nil {
		fmt.Println("fail to query")
		return nil, err
	}
	result := []Lead{}
	for rows.Next() {
		var lead Lead
		loc := location.Location{}
		rent := rental.Rental{}
		err = rows.Scan(&lead.Id, &lead.FirstName, &lead.LastName, &lead.Email, &lead.Phone, &loc.Id, &loc.City, &loc.State, &rent.Id, &rent.Name)
		lead.Location = loc
		lead.Rental = rent
		if err != nil {
			fmt.Println("invalid query row data")
			return nil, err
		}
		result = append(result, lead)
	}
	return result, nil
}
