// Package classification of BusinessAccount API
//
// Documentation for BusinessAccount API
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
import "github.com/bybrisk/businessAccount-api/data"

//
// NOTE: Types defined here are purely for documentation purposes
// these types are not used by any of the handers

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

// Details of a single Business Account
// swagger:response businessAccountGetResponse
type accountGetResponseWrapper struct {
	// Details of a existing Business Account
	// in: body
	Body data.BusinessAccountResponse
}

// Success message on a single Business Account creation
// swagger:response businessAccountPostResponse
type accountPostResponseWrapper struct {
	// Success message on newly created Business Account
	// in: body
	Body data.BusinessAccountPostSuccess
}

// No content is returned by this API endpoint
// swagger:response noContentResponse
type noContentResponseWrapper struct {
}

// swagger:parameters createBusinessAccount
type createAccountParamsWrapper struct {
	// Product data structure or Create.
	// Note: the id field is ignored by create operations
	// in: body
	// required: true
	Body data.BusinessAccountRequest
}

// swagger:parameters updateBusinessAccount
type updateAccountParamsWrapper struct {
	// Product data structure to Update.
	// Note: All the field are required to be sent with the post request along with the updates.
	// in: body
	// required: true
	Body data.UpdateBusinessAccountRequest
}

// swagger:parameters getBusinessAccount
type productIDParamsWrapper struct {
	// The bybID of the Business Account 
	// in: path
	// required: true
	ID int `json:"id"`
}