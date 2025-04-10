// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/epilot-dev/terraform-provider-epilot-saved-view/internal/sdk/models/shared"
	"net/http"
)

type GetTaxonomyBulkActionJobByIDRequest struct {
	JobID string `pathParam:"style=simple,explode=false,name=job_id"`
}

func (o *GetTaxonomyBulkActionJobByIDRequest) GetJobID() string {
	if o == nil {
		return ""
	}
	return o.JobID
}

type GetTaxonomyBulkActionJobByIDResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Returns the job matching the given job id
	TaxonomyBulkJob *shared.TaxonomyBulkJob
}

func (o *GetTaxonomyBulkActionJobByIDResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetTaxonomyBulkActionJobByIDResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetTaxonomyBulkActionJobByIDResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetTaxonomyBulkActionJobByIDResponse) GetTaxonomyBulkJob() *shared.TaxonomyBulkJob {
	if o == nil {
		return nil
	}
	return o.TaxonomyBulkJob
}
