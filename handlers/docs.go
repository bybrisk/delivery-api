// Package classification of Delivery API
//
// Documentation for Delivery API
//
//	Schemes: http
//	BasePath: /
//	Version: 1.0.0
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
// swagger:meta

package handlers
import "github.com/bybrisk/delivery-api/data"

//
// NOTE: Types defined here are purely for documentation purposes
// these types are not used by any of the handlers

// Generic error message returned as a string
// swagger:response errorResponse
type errorResponseWrapper struct {
	// Description of the error
	// in: body
	Body GenericError
}

// Validation errors defined as an array of strings
// swagger:response errorValidation
type errorValidationWrapper struct {
	// Collection of the errors
	// in: body
	Body ValidationError
}

// No content is returned by this API endpoint
// swagger:response noContentResponse
type noContentResponseWrapper struct {
}

// Details of a single delivery
// swagger:response deliveryGetResponse
type accountGetResponseWrapper struct {
	// Details of a existing Business Account
	// in: body
	Body data.BusinessAccountResponse
}

// Success message on a single Delivery addition
// swagger:response deliveryPostResponse
type accountPostResponseWrapper struct {
	// Success message on newly added delivery
	// in: body
	Body data.DeliveryPostSuccess
}

// swagger:parameters addDeliveryWithGeocords
type createAccountParamsWrapper struct {
	// Delivery data structure to add delivery with Geocode.
	// Note: the id field is ignored by create operations
	// in: body
	// required: true
	Body data.AddDeliveryRequest
}

// swagger:parameters addDeliveryWithoutGeocode
type createAccountParamsWrapper struct {
	// Delivery data structure to add delivery without Geocode.
	// Note: the id field is ignored by create operations
	// in: body
	// required: true
	Body data.AddDeliveryWithoutGeocodeRequest
}

// swagger:parameters getDelivery
type deliveryIDParamsWrapper struct {
	// The deliveryID of the Business Account 
	// in: path
	// required: true
	ID string `json:"id"`

	// The date of delivery placed
	// in: path
	// required: true
	DateOfDelivery string `json:"dateOfDelivery"`
}