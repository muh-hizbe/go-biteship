package biteship

import "net/http"

type ResponseWithMap map[string]interface{}

type Shipper struct {
	Name         string `json:"name"`
	Email        string `json:"email"`
	Phone        string `json:"phone"`
	Organization string `json:"organization"`
}

type Origin struct {
	Coordinate   Coordinate `json:"coordinate"`
	PostalCode   uint       `json:"postal_code"`
	ContactName  string     `json:"contact_name"`
	ContactPhone string     `json:"contact_phone"`
	Address      string     `json:"address"`
	Note         string     `json:"note"`
}

type Destination struct {
	Coordinate   Coordinate `json:"coordinate"`
	PostalCode   uint       `json:"postal_code"`
	ContactName  string     `json:"contact_name"`
	ContactPhone string     `json:"contact_phone"`
	ContactEmail string     `json:"contact_email"`
	Address      string     `json:"address"`
	Note         string     `json:"note"`
}

type Courier struct {
	TrackingId *string `json:"tracking_id"`
	WaybillId  *string `json:"waybill_id"`
	Company    string  `json:"company"`
	Name       *string `json:"name"`
	Phone      *string `json:"phone"`
	Type       string  `json:"type"`
	Link       *string `json:"link"`
}

type ResponseCreateOrder struct {
	Success     bool          `json:"success"`
	Message     string        `json:"message"`
	Object      string        `json:"object"`
	Id          string        `json:"id"`
	Shipper     Shipper       `json:"shipper"`
	Origin      Origin        `json:"origin"`
	Destination Destination   `json:"destination"`
	Courier     Courier       `json:"courier"`
	Items       []ProductItem `json:"items"`
	Price       uint          `json:"price"`
	Note        string        `json:"note"`
	Status      string        `json:"status"`
}

type ResponseCourier struct {
	AvailableForCashOnDelivery   bool   `json:"available_for_cash_on_delivery"`
	AvailableForProofOfDelivery  bool   `json:"available_for_proof_of_delivery"`
	AvailableForInstantWaybillId bool   `json:"available_for_instant_waybill_id"`
	CourierName                  string `json:"courier_name"`
	CourierCode                  string `json:"courier_code"`
	CourierServiceName           string `json:"courier_service_name"`
	CourierServiceCode           string `json:"courier_service_code"`
	Tier                         string `json:"tier"`
	Description                  string `json:"description"`
	ServiceType                  string `json:"service_type"`
	ShippingType                 string `json:"shipping_type"`
	ShipmentDurationRange        string `json:"shipment_duration_range"`
	ShipmentDurationUnit         string `json:"shipment_duration_unit"`
}

type ResponseListCourier struct {
	Success  bool              `json:"success"`
	Object   string            `json:"object"`
	Couriers []ResponseCourier `json:"couriers"`
}

type ApiResponse struct {
	Status     string // e.g. "200 OK"
	StatusCode int    // e.g. 200
	Proto      string // e.g. "HTTP/1.0"

	// response Header contain a map of all HTTP header keys to values.
	Header http.Header
	// response body
	RawBody []byte
	// request that was sent to obtain the response
	Request *http.Request
}
