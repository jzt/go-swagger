package pet

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-swagger/go-swagger/errors"
	"github.com/go-swagger/go-swagger/httpkit"
	"github.com/go-swagger/go-swagger/httpkit/middleware"
	"github.com/go-swagger/go-swagger/swag"

	strfmt "github.com/go-swagger/go-swagger/strfmt"
)

// NewUpdatePetWithFormParams creates a new UpdatePetWithFormParams object
// with the default values initialized.
func NewUpdatePetWithFormParams() UpdatePetWithFormParams {
	var ()
	return UpdatePetWithFormParams{}
}

// UpdatePetWithFormParams contains all the bound params for the update pet with form operation
// typically these are obtained from a http.Request
//
// swagger:parameters updatePetWithForm
type UpdatePetWithFormParams struct {
	/*Updated name of the pet
	  In: formData
	*/
	Name *string
	/*ID of pet that needs to be updated
	  Required: true
	  In: path
	*/
	PetID int64
	/*Updated status of the pet
	  In: formData
	*/
	Status *string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls
func (o *UpdatePetWithFormParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error
	if err := r.ParseMultipartForm(32 << 20); err != nil {
		if err != http.ErrNotMultipart {
			return err
		} else if err := r.ParseForm(); err != nil {
			return err
		}
	}
	fds := httpkit.Values(r.Form)

	fdName, fdhkName, _ := fds.GetOK("name")
	if err := o.bindName(fdName, fdhkName, route.Formats); err != nil {
		res = append(res, err)
	}

	rPetID, rhkPetID, _ := route.Params.GetOK("petId")
	if err := o.bindPetID(rPetID, rhkPetID, route.Formats); err != nil {
		res = append(res, err)
	}

	fdStatus, fdhkStatus, _ := fds.GetOK("status")
	if err := o.bindStatus(fdStatus, fdhkStatus, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdatePetWithFormParams) bindName(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}
	if raw == "" { // empty values pass all other validations
		return nil
	}

	o.Name = &raw

	return nil
}

func (o *UpdatePetWithFormParams) bindPetID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("petId", "path", "int64", raw)
	}
	o.PetID = value

	return nil
}

func (o *UpdatePetWithFormParams) bindStatus(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}
	if raw == "" { // empty values pass all other validations
		return nil
	}

	o.Status = &raw

	return nil
}
