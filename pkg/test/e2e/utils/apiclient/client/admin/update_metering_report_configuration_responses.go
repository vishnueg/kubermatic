// Code generated by go-swagger; DO NOT EDIT.

package admin

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"k8c.io/kubermatic/v2/pkg/test/e2e/utils/apiclient/models"
)

// UpdateMeteringReportConfigurationReader is a Reader for the UpdateMeteringReportConfiguration structure.
type UpdateMeteringReportConfigurationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateMeteringReportConfigurationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewUpdateMeteringReportConfigurationCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewUpdateMeteringReportConfigurationUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdateMeteringReportConfigurationForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewUpdateMeteringReportConfigurationDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateMeteringReportConfigurationCreated creates a UpdateMeteringReportConfigurationCreated with default headers values
func NewUpdateMeteringReportConfigurationCreated() *UpdateMeteringReportConfigurationCreated {
	return &UpdateMeteringReportConfigurationCreated{}
}

/* UpdateMeteringReportConfigurationCreated describes a response with status code 201, with default header values.

EmptyResponse is a empty response
*/
type UpdateMeteringReportConfigurationCreated struct {
}

func (o *UpdateMeteringReportConfigurationCreated) Error() string {
	return fmt.Sprintf("[PUT /api/v1/admin/metering/configurations/reports/{name}][%d] updateMeteringReportConfigurationCreated ", 201)
}

func (o *UpdateMeteringReportConfigurationCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateMeteringReportConfigurationUnauthorized creates a UpdateMeteringReportConfigurationUnauthorized with default headers values
func NewUpdateMeteringReportConfigurationUnauthorized() *UpdateMeteringReportConfigurationUnauthorized {
	return &UpdateMeteringReportConfigurationUnauthorized{}
}

/* UpdateMeteringReportConfigurationUnauthorized describes a response with status code 401, with default header values.

EmptyResponse is a empty response
*/
type UpdateMeteringReportConfigurationUnauthorized struct {
}

func (o *UpdateMeteringReportConfigurationUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /api/v1/admin/metering/configurations/reports/{name}][%d] updateMeteringReportConfigurationUnauthorized ", 401)
}

func (o *UpdateMeteringReportConfigurationUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateMeteringReportConfigurationForbidden creates a UpdateMeteringReportConfigurationForbidden with default headers values
func NewUpdateMeteringReportConfigurationForbidden() *UpdateMeteringReportConfigurationForbidden {
	return &UpdateMeteringReportConfigurationForbidden{}
}

/* UpdateMeteringReportConfigurationForbidden describes a response with status code 403, with default header values.

EmptyResponse is a empty response
*/
type UpdateMeteringReportConfigurationForbidden struct {
}

func (o *UpdateMeteringReportConfigurationForbidden) Error() string {
	return fmt.Sprintf("[PUT /api/v1/admin/metering/configurations/reports/{name}][%d] updateMeteringReportConfigurationForbidden ", 403)
}

func (o *UpdateMeteringReportConfigurationForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateMeteringReportConfigurationDefault creates a UpdateMeteringReportConfigurationDefault with default headers values
func NewUpdateMeteringReportConfigurationDefault(code int) *UpdateMeteringReportConfigurationDefault {
	return &UpdateMeteringReportConfigurationDefault{
		_statusCode: code,
	}
}

/* UpdateMeteringReportConfigurationDefault describes a response with status code -1, with default header values.

errorResponse
*/
type UpdateMeteringReportConfigurationDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the update metering report configuration default response
func (o *UpdateMeteringReportConfigurationDefault) Code() int {
	return o._statusCode
}

func (o *UpdateMeteringReportConfigurationDefault) Error() string {
	return fmt.Sprintf("[PUT /api/v1/admin/metering/configurations/reports/{name}][%d] updateMeteringReportConfiguration default  %+v", o._statusCode, o.Payload)
}
func (o *UpdateMeteringReportConfigurationDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *UpdateMeteringReportConfigurationDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*UpdateMeteringReportConfigurationBody update metering report configuration body
swagger:model UpdateMeteringReportConfigurationBody
*/
type UpdateMeteringReportConfigurationBody struct {

	// interval
	Interval int64 `json:"interval,omitempty"`

	// schedule
	Schedule string `json:"schedule,omitempty"`
}

// Validate validates this update metering report configuration body
func (o *UpdateMeteringReportConfigurationBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this update metering report configuration body based on context it is used
func (o *UpdateMeteringReportConfigurationBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *UpdateMeteringReportConfigurationBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateMeteringReportConfigurationBody) UnmarshalBinary(b []byte) error {
	var res UpdateMeteringReportConfigurationBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
