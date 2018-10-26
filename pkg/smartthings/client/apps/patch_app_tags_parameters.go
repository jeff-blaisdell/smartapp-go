// Code generated by go-swagger; DO NOT EDIT.

package apps

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

// NewPatchAppTagsParams creates a new PatchAppTagsParams object
// with the default values initialized.
func NewPatchAppTagsParams() *PatchAppTagsParams {
	var ()
	return &PatchAppTagsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPatchAppTagsParamsWithTimeout creates a new PatchAppTagsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPatchAppTagsParamsWithTimeout(timeout time.Duration) *PatchAppTagsParams {
	var ()
	return &PatchAppTagsParams{

		timeout: timeout,
	}
}

// NewPatchAppTagsParamsWithContext creates a new PatchAppTagsParams object
// with the default values initialized, and the ability to set a context for a request
func NewPatchAppTagsParamsWithContext(ctx context.Context) *PatchAppTagsParams {
	var ()
	return &PatchAppTagsParams{

		Context: ctx,
	}
}

// NewPatchAppTagsParamsWithHTTPClient creates a new PatchAppTagsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPatchAppTagsParamsWithHTTPClient(client *http.Client) *PatchAppTagsParams {
	var ()
	return &PatchAppTagsParams{
		HTTPClient: client,
	}
}

/*PatchAppTagsParams contains all the parameters to send to the API endpoint
for the patch app tags operation typically these are written to a http.Request
*/
type PatchAppTagsParams struct {

	/*Authorization
	  OAuth token

	*/
	Authorization string
	/*AppNameOrID
	  The appName or appId field of an app.

	*/
	AppNameOrID string
	/*PatchAppTagsRequest*/
	PatchAppTagsRequest *models.PatchTagsRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the patch app tags params
func (o *PatchAppTagsParams) WithTimeout(timeout time.Duration) *PatchAppTagsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch app tags params
func (o *PatchAppTagsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch app tags params
func (o *PatchAppTagsParams) WithContext(ctx context.Context) *PatchAppTagsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch app tags params
func (o *PatchAppTagsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch app tags params
func (o *PatchAppTagsParams) WithHTTPClient(client *http.Client) *PatchAppTagsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch app tags params
func (o *PatchAppTagsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the patch app tags params
func (o *PatchAppTagsParams) WithAuthorization(authorization string) *PatchAppTagsParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the patch app tags params
func (o *PatchAppTagsParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithAppNameOrID adds the appNameOrID to the patch app tags params
func (o *PatchAppTagsParams) WithAppNameOrID(appNameOrID string) *PatchAppTagsParams {
	o.SetAppNameOrID(appNameOrID)
	return o
}

// SetAppNameOrID adds the appNameOrId to the patch app tags params
func (o *PatchAppTagsParams) SetAppNameOrID(appNameOrID string) {
	o.AppNameOrID = appNameOrID
}

// WithPatchAppTagsRequest adds the patchAppTagsRequest to the patch app tags params
func (o *PatchAppTagsParams) WithPatchAppTagsRequest(patchAppTagsRequest *models.PatchTagsRequest) *PatchAppTagsParams {
	o.SetPatchAppTagsRequest(patchAppTagsRequest)
	return o
}

// SetPatchAppTagsRequest adds the patchAppTagsRequest to the patch app tags params
func (o *PatchAppTagsParams) SetPatchAppTagsRequest(patchAppTagsRequest *models.PatchTagsRequest) {
	o.PatchAppTagsRequest = patchAppTagsRequest
}

// WriteToRequest writes these params to a swagger request
func (o *PatchAppTagsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// header param Authorization
	if err := r.SetHeaderParam("Authorization", o.Authorization); err != nil {
		return err
	}

	// path param appNameOrId
	if err := r.SetPathParam("appNameOrId", o.AppNameOrID); err != nil {
		return err
	}

	if o.PatchAppTagsRequest != nil {
		if err := r.SetBodyParam(o.PatchAppTagsRequest); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}