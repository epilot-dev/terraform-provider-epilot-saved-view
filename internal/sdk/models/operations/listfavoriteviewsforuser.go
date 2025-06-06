// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/epilot-dev/terraform-provider-epilot-saved-view/internal/sdk/models/shared"
	"net/http"
)

// ListFavoriteViewsForUserResponseBody - Success
type ListFavoriteViewsForUserResponseBody struct {
	Results []shared.SavedViewItem `json:"results,omitempty"`
}

func (o *ListFavoriteViewsForUserResponseBody) GetResults() []shared.SavedViewItem {
	if o == nil {
		return nil
	}
	return o.Results
}

type ListFavoriteViewsForUserResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Success
	Object *ListFavoriteViewsForUserResponseBody
}

func (o *ListFavoriteViewsForUserResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ListFavoriteViewsForUserResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ListFavoriteViewsForUserResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *ListFavoriteViewsForUserResponse) GetObject() *ListFavoriteViewsForUserResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
