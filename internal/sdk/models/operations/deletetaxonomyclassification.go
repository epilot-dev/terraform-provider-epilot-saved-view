// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/epilot-dev/terraform-provider-epilot-saved-view/internal/sdk/models/shared"
	"net/http"
)

type DeleteTaxonomyClassificationRequest struct {
	// Taxonomy Classification slug
	ClassificationSlug string `pathParam:"style=simple,explode=false,name=classificationSlug"`
}

func (o *DeleteTaxonomyClassificationRequest) GetClassificationSlug() string {
	if o == nil {
		return ""
	}
	return o.ClassificationSlug
}

type DeleteTaxonomyClassificationResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Taxonomy classification deleted
	TaxonomyClassification *shared.TaxonomyClassification
}

func (o *DeleteTaxonomyClassificationResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *DeleteTaxonomyClassificationResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *DeleteTaxonomyClassificationResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *DeleteTaxonomyClassificationResponse) GetTaxonomyClassification() *shared.TaxonomyClassification {
	if o == nil {
		return nil
	}
	return o.TaxonomyClassification
}
