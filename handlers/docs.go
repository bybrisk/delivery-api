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
	Body data.AddDeliveryRequestWithGeoCode
}

// swagger:parameters addDeliveryWithoutGeocords
type createDeliveryParmsWrapper struct {
	// Delivery data structure to add delivery without Geocode.
	// Note: latitude and longitude field will be found using address field
	// in: body
	// required: true
	Body data.AddDeliveryRequestWithoutGeoCode
}

// Summary of a delivery associated to a business account
// swagger:response getSingleDeliveryDetail
type deliveryGetOneResponseWrapper struct {
	// Summarised details of a delivery being recieved to a business account
	// in: body
	Body data.SingleDeliveryDetail
}

// swagger:parameters updateDeliveryStatus
type updateDeliveryStatusWrapper struct {
	// Data structure to Update Delivery Status.
	// Note: Can be updated by both admin and delivery app
	// in: body
	// required: true
	Body data.UpdateDeliveryStatus
}

// swagger:parameters updateDeliveryAgent
type updateDeliveryStatusWrapper struct {
	// Data structure to Assign Delivery Agent.
	// Note: Can be updated only by admin
	// in: body
	// required: true
	Body data.UpdateDeliveryAgent
}