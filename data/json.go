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

func (d *UpdateDeliveryStatus) FromJSONToUpdateDeliveryStatus (r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(d)
}

func (d *UpdateDeliveryAgent) FromJSONToUpdateDeliveryAgent (r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(d)
}

func (d *AddDeliveryRequestWithoutGeoCode) FromJSONToAddDeliveryStructAdv (r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(d)
}

func (d *UpdateDeliveryDistance) FromJSONToUpdateDeliveryDistance (r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(d)
}

func (d *SingleDeliveryDetail) GetOneDeliveryResultToJSON (w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(d)
}

func (d *DeliveryResponseBulk) GetAllDeliveryResultToJSON (w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(d)
}

func (d *DeleteAllDeliveryPostSuccess) DeleteAllDeliveryPostSuccessToJSON (w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(d)
}
