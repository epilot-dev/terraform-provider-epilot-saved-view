// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/epilot-dev/terraform-provider-epilot-saved-view/internal/sdk/models/shared"
	"net/http"
)

type CreateSchemaGroupResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Success
	EntitySchemaGroupWithCompositeID *shared.EntitySchemaGroupWithCompositeID
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *CreateSchemaGroupResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreateSchemaGroupResponse) GetEntitySchemaGroupWithCompositeID() *shared.EntitySchemaGroupWithCompositeID {
	if o == nil {
		return nil
	}
	return o.EntitySchemaGroupWithCompositeID
}

func (o *CreateSchemaGroupResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreateSchemaGroupResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
