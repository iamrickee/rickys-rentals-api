package lead

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"iamricky.com/truck-rental/db"
	"iamricky.com/truck-rental/location"
	"iamricky.com/truck-rental/rental"
)

type Lead struct {
	Id        int               `json:"id"`
	FirstName string            `json:"first_name"`
	LastName  string            `json:"last_name"`
	Email     string            `json:"email"`
	Phone     string            `json:"phone"`
	Location  location.Location `json:"location"`
	Rental    rental.Rental     `json:"rental"`
}

type LeadResp struct {
	Data    Lead   `json:"data"`
	Message string `json:"message"`
}

type SuccessResp struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

func SaveRoute(c echo.Context) error {
	lead, err := save(c)
	if err != nil {
		return c.String(http.StatusOK, fmt.Sprintf("%s", err))
	}
	resp := LeadResp{Data: lead, Message: "lead saved"}
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
	resp := SuccessResp{Success: true, Message: "lead deleted"}
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
	lead, err := get(c, idInt)
	if err != nil {
		return c.String(http.StatusOK, fmt.Sprintf("%s", err))
	}
	resp := LeadResp{Data: lead, Message: "lead loaded"}
	jsonResp, err := json.Marshal(resp)
	if err != nil {
		return c.String(http.StatusOK, fmt.Sprintf("%s", err))
	}
	return c.String(http.StatusOK, string(jsonResp))
}

func save(c echo.Context) (Lead, error) {
	id := c.FormValue("id")
	firstName := c.FormValue("first_name")
	lastName := c.FormValue("last_name")
	email := c.FormValue("email")
	phone := c.FormValue("phone")
	locationIDString := c.FormValue("location_id")
	rentalIDString := c.FormValue("rental_id")
	locationID, err := strconv.Atoi(locationIDString)
	if err != nil {
		return Lead{}, err
	}
	rentalID, err := strconv.Atoi(rentalIDString)
	if err != nil {
		return Lead{}, err
	}
	conn, err := db.GetConn()
	if err != nil {
		fmt.Println("fail to connect")
		return Lead{}, err
	}
	idInt, err := strconv.Atoi(id)
	if err == nil {
		q := "UPDATE leads SET first_name = ?, last_name = ?, email = ?, phone = ?, location_id = ?, rental_id = ? WHERE id = ?"
		_, err := conn.ExecContext(c.Request().Context(), q, firstName, lastName, email, phone, locationID, rentalID, idInt)
		if err != nil {
			return Lead{}, err
		}
		return get(c, idInt)
	} else {
		q := "INSERT INTO leads (first_name, last_name, email, phone, location_id, rental_id) VALUES (?,?,?,?,?,?);"
		_, err := conn.ExecContext(c.Request().Context(), q, firstName, lastName, email, phone, locationID, rentalID)
		if err != nil {
			fmt.Printf("first name: %s last name: %s, email: %s, phone: %s, location ID: %d, rental ID: %d, \n", firstName, lastName, email, phone, locationID, rentalID)
			fmt.Println(err)
			return Lead{}, err
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
		q := "DELETE FROM leads WHERE id = ?"
		_, err := conn.ExecContext(c.Request().Context(), q, idInt)
		return err
	} else {
		return err
	}
}

func get(c echo.Context, id int) (Lead, error) {
	conn, err := db.GetConn()
	if err != nil {
		fmt.Println("fail to connect")
		return Lead{}, err
	}
	q := "SELECT a.id, a.first_name, a.last_name, a.email, a.phone, b.id AS bid, b.city, b.state, c.id AS cid, c.name FROM leads AS a LEFT JOIN locations AS b ON a.location_id = b.id LEFT JOIN rentals AS c ON a.rental_id = c.id WHERE a.id = ?"
	lead := Lead{}
	loc := location.Location{}
	rent := rental.Rental{}
	err = conn.QueryRowContext(c.Request().Context(), q, id).Scan(&lead.Id, &lead.FirstName, &lead.LastName, &lead.Email, &lead.Phone, &loc.Id, &loc.City, &loc.State, &rent.Id, &rent.Name)
	lead.Location = loc
	lead.Rental = rent
	if err == sql.ErrNoRows {
		fmt.Println("query returned no rows")
		return Lead{}, err
	} else if err != nil {
		fmt.Println("fail to query")
		return Lead{}, err
	}
	return lead, nil
}

func getLast(c echo.Context) (Lead, error) {
	conn, err := db.GetConn()
	if err != nil {
		fmt.Println("fail to connect")
		return Lead{}, err
	}
	q := "SELECT a.id, a.first_name, a.last_name, a.email, a.phone, b.id AS bid, b.city, b.state, c.id AS cid, c.name FROM leads AS a LEFT JOIN locations AS b ON a.location_id = b.id LEFT JOIN rentals AS c ON a.rental_id = c.id ORDER BY a.id DESC;"
	lead := Lead{}
	loc := location.Location{}
	rent := rental.Rental{}
	err = conn.QueryRowContext(c.Request().Context(), q).Scan(&lead.Id, &lead.FirstName, &lead.LastName, &lead.Email, &lead.Phone, &loc.Id, &loc.City, &loc.State, &rent.Id, &rent.Name)
	lead.Location = loc
	lead.Rental = rent
	if err == sql.ErrNoRows {
		fmt.Println("query returned no rows")
		return Lead{}, err
	} else if err != nil {
		fmt.Println("fail to query")
		return Lead{}, err
	}
	return lead, nil
}
