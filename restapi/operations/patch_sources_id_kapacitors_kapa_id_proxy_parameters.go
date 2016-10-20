package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"io"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/validate"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/influxdata/chronograf/models"
)

// NewPatchSourcesIDKapacitorsKapaIDProxyParams creates a new PatchSourcesIDKapacitorsKapaIDProxyParams object
// with the default values initialized.
func NewPatchSourcesIDKapacitorsKapaIDProxyParams() PatchSourcesIDKapacitorsKapaIDProxyParams {
	var ()
	return PatchSourcesIDKapacitorsKapaIDProxyParams{}
}

// PatchSourcesIDKapacitorsKapaIDProxyParams contains all the bound params for the patch sources ID kapacitors kapa ID proxy operation
// typically these are obtained from a http.Request
//
// swagger:parameters PatchSourcesIDKapacitorsKapaIDProxy
type PatchSourcesIDKapacitorsKapaIDProxyParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request

	/*ID of the source
	  Required: true
	  In: path
	*/
	ID string
	/*ID of the kapacitor backend.
	  Required: true
	  In: path
	*/
	KapaID string
	/*The kapacitor API path to use in the proxy redirect
	  Required: true
	  In: query
	*/
	Path string
	/*Kapacitor body
	  Required: true
	  In: body
	*/
	Query models.KapacitorProxy
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls
func (o *PatchSourcesIDKapacitorsKapaIDProxyParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error
	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	rID, rhkID, _ := route.Params.GetOK("id")
	if err := o.bindID(rID, rhkID, route.Formats); err != nil {
		res = append(res, err)
	}

	rKapaID, rhkKapaID, _ := route.Params.GetOK("kapa_id")
	if err := o.bindKapaID(rKapaID, rhkKapaID, route.Formats); err != nil {
		res = append(res, err)
	}

	qPath, qhkPath, _ := qs.GetOK("path")
	if err := o.bindPath(qPath, qhkPath, route.Formats); err != nil {
		res = append(res, err)
	}

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body models.KapacitorProxy
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			if err == io.EOF {
				res = append(res, errors.Required("query", "body"))
			} else {
				res = append(res, errors.NewParseError("query", "body", "", err))
			}

		} else {

			if len(res) == 0 {
				o.Query = body
			}
		}

	} else {
		res = append(res, errors.Required("query", "body"))
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PatchSourcesIDKapacitorsKapaIDProxyParams) bindID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	o.ID = raw

	return nil
}

func (o *PatchSourcesIDKapacitorsKapaIDProxyParams) bindKapaID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	o.KapaID = raw

	return nil
}

func (o *PatchSourcesIDKapacitorsKapaIDProxyParams) bindPath(rawData []string, hasKey bool, formats strfmt.Registry) error {
	if !hasKey {
		return errors.Required("path", "query")
	}
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}
	if err := validate.RequiredString("path", "query", raw); err != nil {
		return err
	}

	o.Path = raw

	return nil
}
