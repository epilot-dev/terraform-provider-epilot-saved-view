// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/epilot-dev/terraform-provider-epilot-saved-view/internal/sdk/internal/utils"
	"github.com/epilot-dev/terraform-provider-epilot-saved-view/internal/sdk/models/shared"
	"net/http"
)

type ListSavedViewsRequest struct {
	Fields []string `queryParam:"style=form,explode=true,name=fields"`
	From   *int64   `default:"0" queryParam:"style=form,explode=true,name=from"`
	// Number of saved views to return
	Size *int64 `default:"80" queryParam:"style=form,explode=true,name=size"`
	// Return views belonging to this schema
	Slug *string `queryParam:"style=form,explode=true,name=slug"`
	// The sort key to use if present
	Sort *string `default:"name:asc" queryParam:"style=form,explode=true,name=sort"`
}

func (l ListSavedViewsRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(l, "", false)
}

func (l *ListSavedViewsRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &l, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ListSavedViewsRequest) GetFields() []string {
	if o == nil {
		return nil
	}
	return o.Fields
}

func (o *ListSavedViewsRequest) GetFrom() *int64 {
	if o == nil {
		return nil
	}
	return o.From
}

func (o *ListSavedViewsRequest) GetSize() *int64 {
	if o == nil {
		return nil
	}
	return o.Size
}

func (o *ListSavedViewsRequest) GetSlug() *string {
	if o == nil {
		return nil
	}
	return o.Slug
}

func (o *ListSavedViewsRequest) GetSort() *string {
	if o == nil {
		return nil
	}
	return o.Sort
}

type ListSavedViewsResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Success
	ListSavedViewsResults *shared.ListSavedViewsResults
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *ListSavedViewsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ListSavedViewsResponse) GetListSavedViewsResults() *shared.ListSavedViewsResults {
	if o == nil {
		return nil
	}
	return o.ListSavedViewsResults
}

func (o *ListSavedViewsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ListSavedViewsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
