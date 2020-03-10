// Code generated by go-swagger; DO NOT EDIT.

package objects

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/treeverse/lakefs/api/gen/models"
)

// DeleteObjectNoContentCode is the HTTP code returned for type DeleteObjectNoContent
const DeleteObjectNoContentCode int = 204

/*DeleteObjectNoContent object deleted successfully

swagger:response deleteObjectNoContent
*/
type DeleteObjectNoContent struct {
}

// NewDeleteObjectNoContent creates DeleteObjectNoContent with default headers values
func NewDeleteObjectNoContent() *DeleteObjectNoContent {

	return &DeleteObjectNoContent{}
}

// WriteResponse to the client
func (o *DeleteObjectNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(204)
}

// DeleteObjectUnauthorizedCode is the HTTP code returned for type DeleteObjectUnauthorized
const DeleteObjectUnauthorizedCode int = 401

/*DeleteObjectUnauthorized Unauthorized

swagger:response deleteObjectUnauthorized
*/
type DeleteObjectUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteObjectUnauthorized creates DeleteObjectUnauthorized with default headers values
func NewDeleteObjectUnauthorized() *DeleteObjectUnauthorized {

	return &DeleteObjectUnauthorized{}
}

// WithPayload adds the payload to the delete object unauthorized response
func (o *DeleteObjectUnauthorized) WithPayload(payload *models.Error) *DeleteObjectUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete object unauthorized response
func (o *DeleteObjectUnauthorized) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteObjectUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteObjectNotFoundCode is the HTTP code returned for type DeleteObjectNotFound
const DeleteObjectNotFoundCode int = 404

/*DeleteObjectNotFound path or branch not found

swagger:response deleteObjectNotFound
*/
type DeleteObjectNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteObjectNotFound creates DeleteObjectNotFound with default headers values
func NewDeleteObjectNotFound() *DeleteObjectNotFound {

	return &DeleteObjectNotFound{}
}

// WithPayload adds the payload to the delete object not found response
func (o *DeleteObjectNotFound) WithPayload(payload *models.Error) *DeleteObjectNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete object not found response
func (o *DeleteObjectNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteObjectNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*DeleteObjectDefault generic error response

swagger:response deleteObjectDefault
*/
type DeleteObjectDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteObjectDefault creates DeleteObjectDefault with default headers values
func NewDeleteObjectDefault(code int) *DeleteObjectDefault {
	if code <= 0 {
		code = 500
	}

	return &DeleteObjectDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the delete object default response
func (o *DeleteObjectDefault) WithStatusCode(code int) *DeleteObjectDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the delete object default response
func (o *DeleteObjectDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the delete object default response
func (o *DeleteObjectDefault) WithPayload(payload *models.Error) *DeleteObjectDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete object default response
func (o *DeleteObjectDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteObjectDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}