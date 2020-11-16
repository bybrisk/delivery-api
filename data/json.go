package data

import (
	"encoding/json"
	"io"
)	

func (d *BusinessAccountResponse) ToJSON (w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(d)
}

func (d *BusinessAccountRequest) FromJSON (r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(d)
}

func (d *UpdateBusinessAccountRequest) FromJSONUpdateRequest (r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(d)
}

func (d *BusinessAccountPostSuccess) ResultToJSON (w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(d)
}