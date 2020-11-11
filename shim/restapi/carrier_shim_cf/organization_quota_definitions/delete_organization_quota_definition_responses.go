// Code generated by go-swagger; DO NOT EDIT.

package organization_quota_definitions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// DeleteOrganizationQuotaDefinitionNoContentCode is the HTTP code returned for type DeleteOrganizationQuotaDefinitionNoContent
const DeleteOrganizationQuotaDefinitionNoContentCode int = 204

/*DeleteOrganizationQuotaDefinitionNoContent successful response

swagger:response deleteOrganizationQuotaDefinitionNoContent
*/
type DeleteOrganizationQuotaDefinitionNoContent struct {
}

// NewDeleteOrganizationQuotaDefinitionNoContent creates DeleteOrganizationQuotaDefinitionNoContent with default headers values
func NewDeleteOrganizationQuotaDefinitionNoContent() *DeleteOrganizationQuotaDefinitionNoContent {

	return &DeleteOrganizationQuotaDefinitionNoContent{}
}

// WriteResponse to the client
func (o *DeleteOrganizationQuotaDefinitionNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(204)
}