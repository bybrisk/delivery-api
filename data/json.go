package data

import (
	"encoding/json"
	"io"
)	

func (d *DeliveryPostSuccess) FromAddDeliveryStructToJSON (w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(d)
}

func (d *AddDeliveryRequestWithGeoCode) FromJSONToAddDeliveryStruct (r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(d)
}

func (d *AddDeliveryRequestWithoutGeoCode) FromJSONToAddDeliveryStructAdv (r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(d)
}

func (d *SingleDeliveryDetail) GetOneDeliveryResultToJSON (w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(d)
}