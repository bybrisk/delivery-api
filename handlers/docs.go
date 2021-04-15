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

// Success message on a single Delivery addition/update
// swagger:response deliveryPostResponse
type accountPostResponseWrapper struct {
	// Success message on newly added delivery
	// in: body
	Body data.DeliveryPostSuccess
}

// Success message on all Delivery deletion
// swagger:response deleteAllDeliveryDetail
type DeleteDeliveryPostResponseWrapper struct {
	// Success message on newly added delivery
	// in: body
	Body data.DeleteAllDeliveryPostSuccess
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

// Summary of all delivery associated to a business account
// swagger:response getAllDeliveryDetail
type deliveryGetAllResponseWrapper struct {
	// Summarised details of all deliveries of a business account
	// in: body
	Body data.DeliveryResponseBulk
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
type updateDeliveryAgentWrapper struct {
	// Data structure to Assign Delivery Agent.
	// Note: Can be updated only by admin
	// in: body
	// required: true
	Body data.UpdateDeliveryAgent
}

// swagger:parameters updateDeliveryDistance
type updateDeliveryDistanceWrapper struct {
	// Data structure to Update Delivery Distance.
	// Note: Can be updated only by delivery person application
	// in: body
	// required: true
	Body data.UpdateDeliveryDistance
}