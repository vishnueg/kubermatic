// Code generated by go-swagger; DO NOT EDIT.

package project

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"k8c.io/kubermatic/v2/pkg/test/e2e/utils/apiclient/models"
)

// DeleteNodeDeploymentReader is a Reader for the DeleteNodeDeployment structure.
type DeleteNodeDeploymentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteNodeDeploymentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteNodeDeploymentOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewDeleteNodeDeploymentUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteNodeDeploymentForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewDeleteNodeDeploymentDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteNodeDeploymentOK creates a DeleteNodeDeploymentOK with default headers values
func NewDeleteNodeDeploymentOK() *DeleteNodeDeploymentOK {
	return &DeleteNodeDeploymentOK{}
}

/* DeleteNodeDeploymentOK describes a response with status code 200, with default header values.

EmptyResponse is a empty response
*/
type DeleteNodeDeploymentOK struct {
}

func (o *DeleteNodeDeploymentOK) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/nodedeployments/{nodedeployment_id}][%d] deleteNodeDeploymentOK ", 200)
}

func (o *DeleteNodeDeploymentOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteNodeDeploymentUnauthorized creates a DeleteNodeDeploymentUnauthorized with default headers values
func NewDeleteNodeDeploymentUnauthorized() *DeleteNodeDeploymentUnauthorized {
	return &DeleteNodeDeploymentUnauthorized{}
}

/* DeleteNodeDeploymentUnauthorized describes a response with status code 401, with default header values.

EmptyResponse is a empty response
*/
type DeleteNodeDeploymentUnauthorized struct {
}

func (o *DeleteNodeDeploymentUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/nodedeployments/{nodedeployment_id}][%d] deleteNodeDeploymentUnauthorized ", 401)
}

func (o *DeleteNodeDeploymentUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteNodeDeploymentForbidden creates a DeleteNodeDeploymentForbidden with default headers values
func NewDeleteNodeDeploymentForbidden() *DeleteNodeDeploymentForbidden {
	return &DeleteNodeDeploymentForbidden{}
}

/* DeleteNodeDeploymentForbidden describes a response with status code 403, with default header values.

EmptyResponse is a empty response
*/
type DeleteNodeDeploymentForbidden struct {
}

func (o *DeleteNodeDeploymentForbidden) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/nodedeployments/{nodedeployment_id}][%d] deleteNodeDeploymentForbidden ", 403)
}

func (o *DeleteNodeDeploymentForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteNodeDeploymentDefault creates a DeleteNodeDeploymentDefault with default headers values
func NewDeleteNodeDeploymentDefault(code int) *DeleteNodeDeploymentDefault {
	return &DeleteNodeDeploymentDefault{
		_statusCode: code,
	}
}

/* DeleteNodeDeploymentDefault describes a response with status code -1, with default header values.

errorResponse
*/
type DeleteNodeDeploymentDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the delete node deployment default response
func (o *DeleteNodeDeploymentDefault) Code() int {
	return o._statusCode
}

func (o *DeleteNodeDeploymentDefault) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/nodedeployments/{nodedeployment_id}][%d] deleteNodeDeployment default  %+v", o._statusCode, o.Payload)
}
func (o *DeleteNodeDeploymentDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *DeleteNodeDeploymentDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
