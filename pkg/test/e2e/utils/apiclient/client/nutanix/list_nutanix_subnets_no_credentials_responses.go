// Code generated by go-swagger; DO NOT EDIT.

package nutanix

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"k8c.io/kubermatic/v2/pkg/test/e2e/utils/apiclient/models"
)

// ListNutanixSubnetsNoCredentialsReader is a Reader for the ListNutanixSubnetsNoCredentials structure.
type ListNutanixSubnetsNoCredentialsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListNutanixSubnetsNoCredentialsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListNutanixSubnetsNoCredentialsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewListNutanixSubnetsNoCredentialsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListNutanixSubnetsNoCredentialsOK creates a ListNutanixSubnetsNoCredentialsOK with default headers values
func NewListNutanixSubnetsNoCredentialsOK() *ListNutanixSubnetsNoCredentialsOK {
	return &ListNutanixSubnetsNoCredentialsOK{}
}

/* ListNutanixSubnetsNoCredentialsOK describes a response with status code 200, with default header values.

NutanixSubnetList
*/
type ListNutanixSubnetsNoCredentialsOK struct {
	Payload models.NutanixSubnetList
}

func (o *ListNutanixSubnetsNoCredentialsOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}/providers/nutanix/subnets][%d] listNutanixSubnetsNoCredentialsOK  %+v", 200, o.Payload)
}
func (o *ListNutanixSubnetsNoCredentialsOK) GetPayload() models.NutanixSubnetList {
	return o.Payload
}

func (o *ListNutanixSubnetsNoCredentialsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListNutanixSubnetsNoCredentialsDefault creates a ListNutanixSubnetsNoCredentialsDefault with default headers values
func NewListNutanixSubnetsNoCredentialsDefault(code int) *ListNutanixSubnetsNoCredentialsDefault {
	return &ListNutanixSubnetsNoCredentialsDefault{
		_statusCode: code,
	}
}

/* ListNutanixSubnetsNoCredentialsDefault describes a response with status code -1, with default header values.

errorResponse
*/
type ListNutanixSubnetsNoCredentialsDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the list nutanix subnets no credentials default response
func (o *ListNutanixSubnetsNoCredentialsDefault) Code() int {
	return o._statusCode
}

func (o *ListNutanixSubnetsNoCredentialsDefault) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}/providers/nutanix/subnets][%d] listNutanixSubnetsNoCredentials default  %+v", o._statusCode, o.Payload)
}
func (o *ListNutanixSubnetsNoCredentialsDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ListNutanixSubnetsNoCredentialsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
