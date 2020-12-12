package common

//Struct API
// Order struct (Model) ...

type Message  struct {
	Code      int  	 `json:"code"`
	Remark    string `json:"remark"`
	OrderID   string `json:"orderID"`	
	Orders   *Orders `json:"orders,omitempty"`
	Result   *Result  `json:"result,omitempty"`
}

type Orders struct {
	OrderID      string        `json:"orderID"`
	CustomerID   string        `json:"customerID"`
	EmployeeID   string        `json:"employeeID"`
	OrderDate    string        `json:"orderDate"`
	OrdersDet    []OrdersDetail   `json:"ordersDetail"`
	
}

type OrdersDetail struct {
	OrderID      string  `json:"orderID"`
	ProductID  	 string  `json:"ProductID"`
	ProductName  string  `json:"ProductName"`
	UnitPrice    float64 `json:"UnitPrice"`
	Quantity     int     `json:"Quantity"`
}

type Result struct {
	Code   int    `json:"code"`
	Remark string `json:"remark,omitempty"`
}

type Customers struct {
	CustomerID   string `json:"CustomerID"`
	CompanyName  string `json:"CompanyName"`
	ContactName  string `json:"ContactName"`
	ContactTitle string `json:"ContactTitle"`
	Address      string `json:"Address"`
	City         string `json:"City"`
	Country      string `json:"Country"`
	Phone        string `json:"Phone"`
	PostalCode   string `json:"PostalCode"`
}
//End Struct API

type FastPayRequest struct {
	Merchant   string `json:"merchant"`
	MerchantID string `json:"merchant_id"`
	Request    string `json:"request"`
	Signature  string `json:"signature"`
}


type FastPayResponse struct {	
	Response     					 string `json:"response"`
	Merchant      					 string `json:"merchant"`
	MerchantID    					 string `json:"merchant_id"`
	PaymentChannel []PaymentChannel  `json:"payment_channel"`	
	ResponseCode 					 string `json:"response_code"`
	ResponseDesc 					 string `json:"response_desc"`
}

type PaymentChannel struct {
	PgCode string `json:"pg_code"`
	PgName string `json:"pg_name"`
}

//my trips

type MyTrips struct {
	DepatureDate1 string `json:"depature_date_1"`
	DepatureDate2 string `json:"depature_date_2"`
	Provinsi      int64  `json:"provinsi"`
}

type MytripsResponse struct {
	Message string 		  `json:"message"`
	Status  string 		  `json:"status"`
	TripDetail 	[]TripDetail  `json:"data"`	
}

type TripDetail struct {
	AirlineName      string `json:"AirlineName,omitempty"`
	AirportName      string `json:"AirportName,omitempty"`
	CityName         string `json:"CityName,omitempty"`
	Currency         string `json:"Currency,omitempty"`
	DepartureDate    string `json:"DepartureDate,omitempty"`
	Description      string `json:"Description,omitempty"`
	Destination      string `json:"Destination,omitempty"`
	DetailTransit    string `json:"DetailTransit,omitempty"`
	DoubleType       string `json:"DoubleType,omitempty"`
	Duration         string `json:"Duration,omitempty"`
	Goods            string `json:"Goods,omitempty"`
	HotelName        string `json:"HotelName,omitempty"`
	HotelRating      string `json:"HotelRating,omitempty"`
	Lat              string `json:"Lat,omitempty"`
	LicenseNumber    string `json:"LicenseNumber,omitempty"`
	Logo             string `json:"Logo,omitempty"`
	Long             string `json:"Long,omitempty"`
	Origin           string `json:"Origin,omitempty"`
	OriginCity       string `json:"OriginCity,omitempty"`
	Price            string `json:"Price,omitempty"`
	PromoCode        string `json:"PromoCode,omitempty"`
	PromoDescription string `json:"PromoDescription,omitempty"`
	Provinsi         string `json:"Provinsi,omitempty"`
	QuadType         string `json:"QuadType,omitempty"`
	Rating           string `json:"Rating,omitempty"`
	ReturnDate       string `json:"ReturnDate,omitempty"`
	TermCondition    string `json:"TermCondition,omitempty"`
	Transit          string `json:"Transit,omitempty"`
	TravelID         string `json:"TravelID,omitempty"`
	TravelName       string `json:"TravelName,omitempty"`
	TripID           string `json:"TripID,omitempty"`
	TripleType       string `json:"TripleType,omitempty"`
}

type MytripsResponseXML struct {
	Message 		string 		  `xml:"message"`
	Status  		string 		  `xml:"status"`
	TripDetailXML 	[]TripDetailXML  `xml:"Trips"`	
}

type TripDetailXML struct {
	AirlineName      string `xml:"AirlineName,omitempty"`
	AirportName      string `xml:"AirportName,omitempty"`
	CityName         string `xml:"CityName,omitempty"`
	Currency         string `xml:"Currency,omitempty"`
	DepartureDate    string `xml:"DepartureDate,omitempty"`
	Description      string `xml:"Description,omitempty"`
	Destination      string `xml:"Destination,omitempty"`
	DetailTransit    string `xml:"DetailTransit,omitempty"`
	DoubleType       string `xml:"DoubleType,omitempty"`
	Duration         string `xml:"Duration,omitempty"`
	Goods            string `xml:"Goods,omitempty"`
	HotelName        string `xml:"HotelName,omitempty"`
	HotelRating      string `xml:"HotelRating,omitempty"`
	Lat              string `xml:"Lat,omitempty"`
	LicenseNumber    string `xml:"LicenseNumber,omitempty"`
	Logo             string `xml:"Logo,omitempty"`
	Long             string `xml:"Long,omitempty"`
	Origin           string `xml:"Origin,omitempty"`
	OriginCity       string `xml:"OriginCity,omitempty"`
	Price            string `xml:"Price,omitempty"`
	PromoCode        string `xml:"PromoCode,omitempty"`
	PromoDescription string `xml:"PromoDescription,omitempty"`
	Provinsi         string `xml:"Provinsi,omitempty"`
	QuadType         string `xml:"QuadType,omitempty"`
	Rating           string `xml:"Rating,omitempty"`
	ReturnDate       string `xml:"ReturnDate,omitempty"`
	TermCondition    string `xml:"TermCondition,omitempty"`
	Transit          string `xml:"Transit,omitempty"`
	TravelID         string `xml:"TravelID,omitempty"`
	TravelName       string `xml:"TravelName,omitempty"`
	TripID           string `xml:"TripID,omitempty"`
	TripleType       string `xml:"TripleType,omitempty"`
}