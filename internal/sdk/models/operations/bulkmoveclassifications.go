// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/epilot-dev/terraform-provider-epilot-saved-view/internal/sdk/models/shared"
	"net/http"
)

type BulkMoveClassificationsRequestBody struct {
	ClassificationIds []string `json:"classification_ids,omitempty"`
	// Job ID for tracking the status of a bulk operation request
	JobID          *string `json:"job_id,omitempty"`
	TargetTaxonomy *string `json:"target_taxonomy,omitempty"`
}

func (o *BulkMoveClassificationsRequestBody) GetClassificationIds() []string {
	if o == nil {
		return nil
	}
	return o.ClassificationIds
}

func (o *BulkMoveClassificationsRequestBody) GetJobID() *string {
	if o == nil {
		return nil
	}
	return o.JobID
}

func (o *BulkMoveClassificationsRequestBody) GetTargetTaxonomy() *string {
	if o == nil {
		return nil
	}
	return o.TargetTaxonomy
}

type BulkMoveClassificationsResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Returns the job id for the async job and current status
	TaxonomyBulkJobTriggerResponse *shared.TaxonomyBulkJobTriggerResponse
}

func (o *BulkMoveClassificationsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *BulkMoveClassificationsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *BulkMoveClassificationsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *BulkMoveClassificationsResponse) GetTaxonomyBulkJobTriggerResponse() *shared.TaxonomyBulkJobTriggerResponse {
	if o == nil {
		return nil
	}
	return o.TaxonomyBulkJobTriggerResponse
}
