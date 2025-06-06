// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"errors"
	"fmt"
	"github.com/epilot-dev/terraform-provider-epilot-saved-view/internal/sdk/internal/utils"
	"github.com/epilot-dev/terraform-provider-epilot-saved-view/internal/sdk/models/shared"
	"net/http"
)

type TaxonomiesClassificationsSearchRequestBody struct {
	ClassificationIds []string `json:"classificationIds,omitempty"`
}

func (o *TaxonomiesClassificationsSearchRequestBody) GetClassificationIds() []string {
	if o == nil {
		return nil
	}
	return o.ClassificationIds
}

type TaxonomySlugType string

const (
	TaxonomySlugTypeStr        TaxonomySlugType = "str"
	TaxonomySlugTypeArrayOfStr TaxonomySlugType = "arrayOfStr"
)

// TaxonomySlug - The taxonomy slug(s) to search within. When provided with multiple taxonomy slugs, the search will be performed across all the provided taxonomies.
type TaxonomySlug struct {
	Str        *string  `queryParam:"inline"`
	ArrayOfStr []string `queryParam:"inline"`

	Type TaxonomySlugType
}

func CreateTaxonomySlugStr(str string) TaxonomySlug {
	typ := TaxonomySlugTypeStr

	return TaxonomySlug{
		Str:  &str,
		Type: typ,
	}
}

func CreateTaxonomySlugArrayOfStr(arrayOfStr []string) TaxonomySlug {
	typ := TaxonomySlugTypeArrayOfStr

	return TaxonomySlug{
		ArrayOfStr: arrayOfStr,
		Type:       typ,
	}
}

func (u *TaxonomySlug) UnmarshalJSON(data []byte) error {

	var str string = ""
	if err := utils.UnmarshalJSON(data, &str, "", true, true); err == nil {
		u.Str = &str
		u.Type = TaxonomySlugTypeStr
		return nil
	}

	var arrayOfStr []string = []string{}
	if err := utils.UnmarshalJSON(data, &arrayOfStr, "", true, true); err == nil {
		u.ArrayOfStr = arrayOfStr
		u.Type = TaxonomySlugTypeArrayOfStr
		return nil
	}

	return fmt.Errorf("could not unmarshal `%s` into any supported union types for TaxonomySlug", string(data))
}

func (u TaxonomySlug) MarshalJSON() ([]byte, error) {
	if u.Str != nil {
		return utils.MarshalJSON(u.Str, "", true)
	}

	if u.ArrayOfStr != nil {
		return utils.MarshalJSON(u.ArrayOfStr, "", true)
	}

	return nil, errors.New("could not marshal union type TaxonomySlug: all fields are null")
}

type TaxonomiesClassificationsSearchRequest struct {
	RequestBody *TaxonomiesClassificationsSearchRequestBody `request:"mediaType=application/json"`
	// Filter by archived status. Deprecated. Use `include_archived` instead.
	//
	// Deprecated: This will be removed in a future release, please migrate away from it as soon as possible.
	Archived *bool `queryParam:"style=form,explode=true,name=archived"`
	// Whether to include archived labels in the search results
	// - `true`: include archived labels
	// - `false`: exclude archived labels
	// - `only`: include only archived labels
	//
	// By default, no archived labels are included in the search results.
	//
	IncludeArchived *shared.TaxonomySearchIncludeArchivedParam `default:"false" queryParam:"style=form,explode=true,name=include_archived"`
	// The label names to search for (lowercase insensitive)
	Query *string `queryParam:"style=form,explode=true,name=query"`
	// The taxonomy slug(s) to search within. When provided with multiple taxonomy slugs, the search will be performed across all the provided taxonomies.
	//
	TaxonomySlug *TaxonomySlug `queryParam:"style=form,explode=true,name=taxonomySlug"`
}

func (t TaxonomiesClassificationsSearchRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(t, "", false)
}

func (t *TaxonomiesClassificationsSearchRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &t, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *TaxonomiesClassificationsSearchRequest) GetRequestBody() *TaxonomiesClassificationsSearchRequestBody {
	if o == nil {
		return nil
	}
	return o.RequestBody
}

func (o *TaxonomiesClassificationsSearchRequest) GetArchived() *bool {
	if o == nil {
		return nil
	}
	return o.Archived
}

func (o *TaxonomiesClassificationsSearchRequest) GetIncludeArchived() *shared.TaxonomySearchIncludeArchivedParam {
	if o == nil {
		return nil
	}
	return o.IncludeArchived
}

func (o *TaxonomiesClassificationsSearchRequest) GetQuery() *string {
	if o == nil {
		return nil
	}
	return o.Query
}

func (o *TaxonomiesClassificationsSearchRequest) GetTaxonomySlug() *TaxonomySlug {
	if o == nil {
		return nil
	}
	return o.TaxonomySlug
}

// TaxonomiesClassificationsSearchResponseBody - Returns the classifications for the taxonomy slug provided
type TaxonomiesClassificationsSearchResponseBody struct {
	Hits    *int64                          `json:"hits,omitempty"`
	Results []shared.TaxonomyClassification `json:"results,omitempty"`
}

func (o *TaxonomiesClassificationsSearchResponseBody) GetHits() *int64 {
	if o == nil {
		return nil
	}
	return o.Hits
}

func (o *TaxonomiesClassificationsSearchResponseBody) GetResults() []shared.TaxonomyClassification {
	if o == nil {
		return nil
	}
	return o.Results
}

type TaxonomiesClassificationsSearchResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Returns the classifications for the taxonomy slug provided
	Object *TaxonomiesClassificationsSearchResponseBody
}

func (o *TaxonomiesClassificationsSearchResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *TaxonomiesClassificationsSearchResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *TaxonomiesClassificationsSearchResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *TaxonomiesClassificationsSearchResponse) GetObject() *TaxonomiesClassificationsSearchResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
