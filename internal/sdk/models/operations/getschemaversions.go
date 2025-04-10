// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/epilot-dev/terraform-provider-epilot-saved-view/internal/sdk/models/shared"
	"net/http"
)

type GetSchemaVersionsRequest struct {
	DraftsFrom *float64 `queryParam:"style=form,explode=true,name=drafts_from"`
	DraftsSize *float64 `queryParam:"style=form,explode=true,name=drafts_size"`
	Fields     []string `queryParam:"style=form,explode=false,name=fields"`
	// Entity Type
	Slug         string   `pathParam:"style=simple,explode=false,name=slug"`
	VersionsFrom *float64 `queryParam:"style=form,explode=true,name=versions_from"`
	VersionsSize *float64 `queryParam:"style=form,explode=true,name=versions_size"`
}

func (o *GetSchemaVersionsRequest) GetDraftsFrom() *float64 {
	if o == nil {
		return nil
	}
	return o.DraftsFrom
}

func (o *GetSchemaVersionsRequest) GetDraftsSize() *float64 {
	if o == nil {
		return nil
	}
	return o.DraftsSize
}

func (o *GetSchemaVersionsRequest) GetFields() []string {
	if o == nil {
		return nil
	}
	return o.Fields
}

func (o *GetSchemaVersionsRequest) GetSlug() string {
	if o == nil {
		return ""
	}
	return o.Slug
}

func (o *GetSchemaVersionsRequest) GetVersionsFrom() *float64 {
	if o == nil {
		return nil
	}
	return o.VersionsFrom
}

func (o *GetSchemaVersionsRequest) GetVersionsSize() *float64 {
	if o == nil {
		return nil
	}
	return o.VersionsSize
}

// GetSchemaVersionsResponseBody - Success
type GetSchemaVersionsResponseBody struct {
	Drafts   []shared.EntitySchemaItem `json:"drafts,omitempty"`
	Versions []shared.EntitySchemaItem `json:"versions,omitempty"`
}

func (o *GetSchemaVersionsResponseBody) GetDrafts() []shared.EntitySchemaItem {
	if o == nil {
		return nil
	}
	return o.Drafts
}

func (o *GetSchemaVersionsResponseBody) GetVersions() []shared.EntitySchemaItem {
	if o == nil {
		return nil
	}
	return o.Versions
}

type GetSchemaVersionsResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Success
	Object *GetSchemaVersionsResponseBody
}

func (o *GetSchemaVersionsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetSchemaVersionsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetSchemaVersionsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetSchemaVersionsResponse) GetObject() *GetSchemaVersionsResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
