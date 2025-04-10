// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/epilot-dev/terraform-provider-epilot-saved-view/internal/sdk/internal/utils"
	"github.com/epilot-dev/terraform-provider-epilot-saved-view/internal/sdk/models/shared"
	"net/http"
)

type GetTaxonomyBulkActionJobsRequest struct {
	Size *float64 `default:"20" queryParam:"style=form,explode=true,name=size"`
	// The status of the jobs to return
	Status *shared.TaxonomyBulkJobStatus `queryParam:"style=form,explode=true,name=status"`
}

func (g GetTaxonomyBulkActionJobsRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(g, "", false)
}

func (g *GetTaxonomyBulkActionJobsRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &g, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *GetTaxonomyBulkActionJobsRequest) GetSize() *float64 {
	if o == nil {
		return nil
	}
	return o.Size
}

func (o *GetTaxonomyBulkActionJobsRequest) GetStatus() *shared.TaxonomyBulkJobStatus {
	if o == nil {
		return nil
	}
	return o.Status
}

type GetTaxonomyBulkActionJobsResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Returns the jobs matching the given query or the active jobs if no status is provided
	Classes []shared.TaxonomyBulkJob
}

func (o *GetTaxonomyBulkActionJobsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetTaxonomyBulkActionJobsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetTaxonomyBulkActionJobsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetTaxonomyBulkActionJobsResponse) GetClasses() []shared.TaxonomyBulkJob {
	if o == nil {
		return nil
	}
	return o.Classes
}
