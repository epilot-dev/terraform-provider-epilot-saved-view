// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/epilot-dev/terraform-provider-epilot-saved-view/internal/sdk/models/shared"
	"net/http"
)

type ListEntitiesResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Success
	EntitySearchResults *shared.EntitySearchResults
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Success
	Res *string
}

func (o *ListEntitiesResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ListEntitiesResponse) GetEntitySearchResults() *shared.EntitySearchResults {
	if o == nil {
		return nil
	}
	return o.EntitySearchResults
}

func (o *ListEntitiesResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ListEntitiesResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *ListEntitiesResponse) GetRes() *string {
	if o == nil {
		return nil
	}
	return o.Res
}
