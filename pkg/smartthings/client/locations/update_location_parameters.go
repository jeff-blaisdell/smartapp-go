// Code generated by go-swagger; DO NOT EDIT.

package locations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/SmartThingsOSS/smartapp-go/pkg/smartthings/models"
)

// NewUpdateLocationParams creates a new UpdateLocationParams object
// with the default values initialized.
func NewUpdateLocationParams() *UpdateLocationParams {
	var ()
	return &UpdateLocationParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateLocationParamsWithTimeout creates a new UpdateLocationParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateLocationParamsWithTimeout(timeout time.Duration) *UpdateLocationParams {
	var ()
	return &UpdateLocationParams{

		timeout: timeout,
	}
}

// NewUpdateLocationParamsWithContext creates a new UpdateLocationParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateLocationParamsWithContext(ctx context.Context) *UpdateLocationParams {
	var ()
	return &UpdateLocationParams{

		Context: ctx,
	}
}

// NewUpdateLocationParamsWithHTTPClient creates a new UpdateLocationParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateLocationParamsWithHTTPClient(client *http.Client) *UpdateLocationParams {
	var ()
	return &UpdateLocationParams{
		HTTPClient: client,
	}
}

/*UpdateLocationParams contains all the parameters to send to the API endpoint
for the update location operation typically these are written to a http.Request
*/
type UpdateLocationParams struct {

	/*Authorization
	  OAuth token

	*/
	Authorization string
	/*LocationID
	  The ID of the location.

	*/
	LocationID string
	/*UpdateLocationRequest*/
	UpdateLocationRequest *models.UpdateLocationRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update location params
func (o *UpdateLocationParams) WithTimeout(timeout time.Duration) *UpdateLocationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update location params
func (o *UpdateLocationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update location params
func (o *UpdateLocationParams) WithContext(ctx context.Context) *UpdateLocationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update location params
func (o *UpdateLocationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update location params
func (o *UpdateLocationParams) WithHTTPClient(client *http.Client) *UpdateLocationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update location params
func (o *UpdateLocationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the update location params
func (o *UpdateLocationParams) WithAuthorization(authorization string) *UpdateLocationParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the update location params
func (o *UpdateLocationParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithLocationID adds the locationID to the update location params
func (o *UpdateLocationParams) WithLocationID(locationID string) *UpdateLocationParams {
	o.SetLocationID(locationID)
	return o
}

// SetLocationID adds the locationId to the update location params
func (o *UpdateLocationParams) SetLocationID(locationID string) {
	o.LocationID = locationID
}

// WithUpdateLocationRequest adds the updateLocationRequest to the update location params
func (o *UpdateLocationParams) WithUpdateLocationRequest(updateLocationRequest *models.UpdateLocationRequest) *UpdateLocationParams {
	o.SetUpdateLocationRequest(updateLocationRequest)
	return o
}

// SetUpdateLocationRequest adds the updateLocationRequest to the update location params
func (o *UpdateLocationParams) SetUpdateLocationRequest(updateLocationRequest *models.UpdateLocationRequest) {
	o.UpdateLocationRequest = updateLocationRequest
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateLocationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// header param Authorization
	if err := r.SetHeaderParam("Authorization", o.Authorization); err != nil {
		return err
	}

	// path param locationId
	if err := r.SetPathParam("locationId", o.LocationID); err != nil {
		return err
	}

	if o.UpdateLocationRequest != nil {
		if err := r.SetBodyParam(o.UpdateLocationRequest); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}