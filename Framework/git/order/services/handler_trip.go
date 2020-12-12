package services

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	cm "pnp/Framework/git/order/common"

	_ "github.com/go-sql-driver/mysql"
)



func (PaymentService) TripsHandler(ctx context.Context, req cm.MyTrips) (res cm.MytripsResponse) {

	msg := &cm.MyTrips{
		Provinsi: req.Provinsi,
		DepatureDate1: req.DepatureDate1,
		DepatureDate2: req.DepatureDate2,		
	}
	
	reqBody, err := json.Marshal(msg)

	if err != nil {
		print(err)
	}
	
	resp, err := http.Post("http://35.186.147.192/travel/GetTripsSample.php", "application/json", bytes.NewBuffer(reqBody))

	if err != nil {
		print(err)
	}
	
	defer resp.Body.Close()
	
	body, err := ioutil.ReadAll(resp.Body)
	
	if err != nil {
		print(err)
	}
		
	var response cm.MytripsResponse
	json.Unmarshal(body, &response)		

	res.Message = response.Message
	res.Status = response.Status
	
	res.TripDetail = response.TripDetail

	for _, data := range response.TripDetail {
		
		fmt.Println("AirlineName : ", data.AirlineName)
		
	}	

	return
}