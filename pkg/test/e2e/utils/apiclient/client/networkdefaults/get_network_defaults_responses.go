// Code generated by go-swagger; DO NOT EDIT.

package networkdefaults

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"k8c.io/kubermatic/v2/pkg/test/e2e/utils/apiclient/models"
)

// GetNetworkDefaultsReader is a Reader for the GetNetworkDefaults structure.
type GetNetworkDefaultsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetNetworkDefaultsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetNetworkDefaultsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetNetworkDefaultsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetNetworkDefaultsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetNetworkDefaultsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetNetworkDefaultsOK creates a GetNetworkDefaultsOK with default headers values
func NewGetNetworkDefaultsOK() *GetNetworkDefaultsOK {
	return &GetNetworkDefaultsOK{}
}

/* GetNetworkDefaultsOK describes a response with status code 200, with default header values.

NetworkDefaults
*/
type GetNetworkDefaultsOK struct {
	Payload *models.NetworkDefaults
}

func (o *GetNetworkDefaultsOK) Error() string {
	return fmt.Sprintf("[GET /providers/{provider_name}/cni/{cni_plugin_type}/networkdefaults][%d] getNetworkDefaultsOK  %+v", 200, o.Payload)
}
func (o *GetNetworkDefaultsOK) GetPayload() *models.NetworkDefaults {
	return o.Payload
}

func (o *GetNetworkDefaultsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NetworkDefaults)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetNetworkDefaultsUnauthorized creates a GetNetworkDefaultsUnauthorized with default headers values
func NewGetNetworkDefaultsUnauthorized() *GetNetworkDefaultsUnauthorized {
	return &GetNetworkDefaultsUnauthorized{}
}

/* GetNetworkDefaultsUnauthorized describes a response with status code 401, with default header values.

EmptyResponse is a empty response
*/
type GetNetworkDefaultsUnauthorized struct {
}

func (o *GetNetworkDefaultsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /providers/{provider_name}/cni/{cni_plugin_type}/networkdefaults][%d] getNetworkDefaultsUnauthorized ", 401)
}

func (o *GetNetworkDefaultsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetNetworkDefaultsForbidden creates a GetNetworkDefaultsForbidden with default headers values
func NewGetNetworkDefaultsForbidden() *GetNetworkDefaultsForbidden {
	return &GetNetworkDefaultsForbidden{}
}

/* GetNetworkDefaultsForbidden describes a response with status code 403, with default header values.

EmptyResponse is a empty response
*/
type GetNetworkDefaultsForbidden struct {
}

func (o *GetNetworkDefaultsForbidden) Error() string {
	return fmt.Sprintf("[GET /providers/{provider_name}/cni/{cni_plugin_type}/networkdefaults][%d] getNetworkDefaultsForbidden ", 403)
}

func (o *GetNetworkDefaultsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetNetworkDefaultsDefault creates a GetNetworkDefaultsDefault with default headers values
func NewGetNetworkDefaultsDefault(code int) *GetNetworkDefaultsDefault {
	return &GetNetworkDefaultsDefault{
		_statusCode: code,
	}
}

/* GetNetworkDefaultsDefault describes a response with status code -1, with default header values.

errorResponse
*/
type GetNetworkDefaultsDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the get network defaults default response
func (o *GetNetworkDefaultsDefault) Code() int {
	return o._statusCode
}

func (o *GetNetworkDefaultsDefault) Error() string {
	return fmt.Sprintf("[GET /providers/{provider_name}/cni/{cni_plugin_type}/networkdefaults][%d] getNetworkDefaults default  %+v", o._statusCode, o.Payload)
}
func (o *GetNetworkDefaultsDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetNetworkDefaultsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
