package data

import (
	"encoding/json"
	"io"
)	

func (d *DeliveryPostSuccess) FromAddDeliveryStructToJSON (w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(d)
}

func (d *AddDeliveryRequest) FromJSONToAddDeliveryStruct (r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(d)
}