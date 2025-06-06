// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/epilot-dev/terraform-provider-epilot-saved-view/internal/sdk/internal/utils"
	"github.com/epilot-dev/terraform-provider-epilot-saved-view/internal/sdk/models/shared"
	"net/http"
)

type GetRelationsV2Request struct {
	// List of entity fields to include in results
	Fields []string `queryParam:"style=form,explode=true,name=fields"`
	// Starting page number
	From *int64 `default:"0" queryParam:"style=form,explode=true,name=from"`
	// When true, enables entity hydration to resolve nested $relation & $relation_ref references in-place.
	Hydrate *bool `default:"false" queryParam:"style=form,explode=true,name=hydrate"`
	// Entity id
	ID string `pathParam:"style=simple,explode=false,name=id"`
	// When true, includes reverse relations in response (other entities pointing to this entity)
	IncludeReverse *bool `default:"false" queryParam:"style=form,explode=true,name=include_reverse"`
	// Input to filter search results
	Query *string `queryParam:"style=form,explode=true,name=query"`
	// Number of results to return per page
	Size *int64 `default:"50" queryParam:"style=form,explode=true,name=size"`
	// Entity Type
	Slug string `pathParam:"style=simple,explode=false,name=slug"`
}

func (g GetRelationsV2Request) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(g, "", false)
}

func (g *GetRelationsV2Request) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &g, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *GetRelationsV2Request) GetFields() []string {
	if o == nil {
		return nil
	}
	return o.Fields
}

func (o *GetRelationsV2Request) GetFrom() *int64 {
	if o == nil {
		return nil
	}
	return o.From
}

func (o *GetRelationsV2Request) GetHydrate() *bool {
	if o == nil {
		return nil
	}
	return o.Hydrate
}

func (o *GetRelationsV2Request) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *GetRelationsV2Request) GetIncludeReverse() *bool {
	if o == nil {
		return nil
	}
	return o.IncludeReverse
}

func (o *GetRelationsV2Request) GetQuery() *string {
	if o == nil {
		return nil
	}
	return o.Query
}

func (o *GetRelationsV2Request) GetSize() *int64 {
	if o == nil {
		return nil
	}
	return o.Size
}

func (o *GetRelationsV2Request) GetSlug() string {
	if o == nil {
		return ""
	}
	return o.Slug
}

// GetRelationsV2ResponseBody - A generic error returned by the API
type GetRelationsV2ResponseBody struct {
	// The error message
	Error *string `json:"error,omitempty"`
	// The HTTP status code of the error
	Status *int64 `json:"status,omitempty"`
}

func (o *GetRelationsV2ResponseBody) GetError() *string {
	if o == nil {
		return nil
	}
	return o.Error
}

func (o *GetRelationsV2ResponseBody) GetStatus() *int64 {
	if o == nil {
		return nil
	}
	return o.Status
}

type GetRelationsV2Response struct {
	// HTTP response content type for this operation
	ContentType string
	// Success
	GetRelationsRespWithPagination *shared.GetRelationsRespWithPagination
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// The requested resource was not found
	Object *GetRelationsV2ResponseBody
}

func (o *GetRelationsV2Response) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetRelationsV2Response) GetGetRelationsRespWithPagination() *shared.GetRelationsRespWithPagination {
	if o == nil {
		return nil
	}
	return o.GetRelationsRespWithPagination
}

func (o *GetRelationsV2Response) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetRelationsV2Response) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetRelationsV2Response) GetObject() *GetRelationsV2ResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
