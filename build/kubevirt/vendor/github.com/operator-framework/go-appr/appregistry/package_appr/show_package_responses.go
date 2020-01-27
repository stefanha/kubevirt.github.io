// Code generated by go-swagger; DO NOT EDIT.

package package_appr

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/operator-framework/go-appr/models"
)

// ShowPackageReader is a Reader for the ShowPackage structure.
type ShowPackageReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ShowPackageReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewShowPackageOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewShowPackageUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewShowPackageNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewShowPackageOK creates a ShowPackageOK with default headers values
func NewShowPackageOK() *ShowPackageOK {
	return &ShowPackageOK{}
}

/*ShowPackageOK handles this case with default header values.

successful operation
*/
type ShowPackageOK struct {
	Payload *models.Package
}

func (o *ShowPackageOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/packages/{namespace}/{package}/{release}/{media_type}][%d] showPackageOK  %+v", 200, o.Payload)
}

func (o *ShowPackageOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Package)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewShowPackageUnauthorized creates a ShowPackageUnauthorized with default headers values
func NewShowPackageUnauthorized() *ShowPackageUnauthorized {
	return &ShowPackageUnauthorized{}
}

/*ShowPackageUnauthorized handles this case with default header values.

Not authorized to read the package
*/
type ShowPackageUnauthorized struct {
	Payload *models.Error
}

func (o *ShowPackageUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v1/packages/{namespace}/{package}/{release}/{media_type}][%d] showPackageUnauthorized  %+v", 401, o.Payload)
}

func (o *ShowPackageUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewShowPackageNotFound creates a ShowPackageNotFound with default headers values
func NewShowPackageNotFound() *ShowPackageNotFound {
	return &ShowPackageNotFound{}
}

/*ShowPackageNotFound handles this case with default header values.

Package not found
*/
type ShowPackageNotFound struct {
	Payload *models.Error
}

func (o *ShowPackageNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v1/packages/{namespace}/{package}/{release}/{media_type}][%d] showPackageNotFound  %+v", 404, o.Payload)
}

func (o *ShowPackageNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
