// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/epilot-dev/terraform-provider-epilot-saved-view/internal/sdk/models/shared"
	"net/http"
)

type UpdateSavedViewRequest struct {
	SavedViewItem *shared.SavedViewItem `request:"mediaType=application/json"`
	// View id
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *UpdateSavedViewRequest) GetSavedViewItem() *shared.SavedViewItem {
	if o == nil {
		return nil
	}
	return o.SavedViewItem
}

func (o *UpdateSavedViewRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type UpdateSavedViewResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Success
	SavedViewItem *shared.SavedViewItem
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *UpdateSavedViewResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateSavedViewResponse) GetSavedViewItem() *shared.SavedViewItem {
	if o == nil {
		return nil
	}
	return o.SavedViewItem
}

func (o *UpdateSavedViewResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateSavedViewResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
