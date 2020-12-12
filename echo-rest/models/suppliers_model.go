package models

import (
	"fmt"
	"net/http"
	"pnp/echo-rest/common"
	cm "pnp/echo-rest/common"
	"pnp/echo-rest/db"

	"github.com/labstack/echo"
	ex "github.com/wolvex/go/error"
)

//var errMessage string

func FetchSuppliers() (res Response, err error) {

	var errs *ex.AppError
	var sup cm.Suppliers
	var supObj []cm.Suppliers

	defer func() {
		if errs != nil {
			res.Status = errs.ErrCode
			res.Message = errs.Remark
		}
	}()

	con := db.CreateCon()

	sqlQuery := `SELECT					
					IFNULL(SupplierID,'') SupplierID,
					IFNULL(CompanyName,'') CompanyName,
					IFNULL(ContactName,'') ContactName,
					IFNULL(ContactTitle,'') ContactTitle,
					IFNULL(Address,'') Address,
					IFNULL(City,'') City
				FROM suppliers`

	rows, err := con.Query(sqlQuery)

	defer rows.Close()

	if err != nil {
		return res, err
	}

	for rows.Next() {

		err = rows.Scan(&sup.SupplierID, &sup.CompanyName, &sup.ContactName, &sup.ContactTitle, &sup.Address, &sup.City)

		if err != nil {
			errs = ex.Errorc(http.StatusNoContent).Rem(err.Error())
			errMessage = err.Error()
			return res, err
		}

		supObj = append(supObj, sup)

	}

	res.Status = http.StatusOK
	res.Message = "succses"
	res.Data = supObj

	return res, nil
}

//AddSupplier...
func AddSupplier(e echo.Context) (res Response, err error) {

	defer func() {
		if errs != nil {
			res.Status = errs.ErrCode
			res.Message = errs.Remark
		}
	}()

	req := new(common.Suppliers)
	if err = e.Bind(req); err != nil {
		return res, err
	}

	con := db.CreateCon()

	sqlStatement := `INSERT INTO suppliers (CompanyName,ContactName,ContactTitle,Address,City)
					 VALUES (?,?,?,?,?)`

	stmt, err := con.Prepare(sqlStatement)

	if err != nil {
		errs = ex.Errorc(http.StatusNoContent).Rem(err.Error())
		errMessage = err.Error()
		return res, err
	}

	result, err := stmt.Exec(req.CompanyName, req.ContactName, req.ContactTitle, req.Address, req.City)

	if err != nil {
		errs = ex.Errorc(http.StatusNoContent).Rem(err.Error())
		errMessage = err.Error()
		return res, err
	}

	fmt.Println(result)

	res.Status = http.StatusOK
	res.Message = "succes"

	res.Data = map[string]string{
		"Supplier ADD": req.ContactName,
	}

	return res, nil
}
