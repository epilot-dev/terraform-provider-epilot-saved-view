// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/epilot-dev/terraform-provider-epilot-saved-view/internal/sdk/internal/utils"
	"github.com/epilot-dev/terraform-provider-epilot-saved-view/internal/sdk/models/shared"
	"net/http"
)

type UpsertEntityRequestBody struct {
	Entity    shared.EntityInput `json:"entity"`
	UniqueKey []string           `json:"unique_key"`
}

func (o *UpsertEntityRequestBody) GetEntity() shared.EntityInput {
	if o == nil {
		return shared.EntityInput{}
	}
	return o.Entity
}

func (o *UpsertEntityRequestBody) GetUniqueKey() []string {
	if o == nil {
		return []string{}
	}
	return o.UniqueKey
}

type UpsertEntityRequest struct {
	RequestBody *UpsertEntityRequestBody `request:"mediaType=application/json"`
	// Activity to include in event feed
	ActivityID *shared.ActivityIDQueryParam `queryParam:"style=form,explode=true,name=activity_id"`
	// Don't wait for updated entity to become available in Search API. Useful for large migrations
	Async *bool `default:"false" queryParam:"style=form,explode=true,name=async"`
	// Dry Run mode = return results but does not perform the operation.
	DryRun *bool `default:"false" queryParam:"style=form,explode=true,name=dry_run"`
	// Update the diff and entity for the custom activity included in the query.
	// Pending state on activity is automatically ended when activity is filled.
	//
	FillActivity *bool `default:"false" queryParam:"style=form,explode=true,name=fill_activity"`
	// Entity Type
	Slug string `pathParam:"style=simple,explode=false,name=slug"`
	// Strict mode = return 409 if more than one entity is matched
	Strict *bool `default:"false" queryParam:"style=form,explode=true,name=strict"`
	// When true, enables entity validation against the entity schema.
	Validate *bool `default:"false" queryParam:"style=form,explode=true,name=validate"`
}

func (u UpsertEntityRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(u, "", false)
}

func (u *UpsertEntityRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &u, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *UpsertEntityRequest) GetRequestBody() *UpsertEntityRequestBody {
	if o == nil {
		return nil
	}
	return o.RequestBody
}

func (o *UpsertEntityRequest) GetActivityID() *shared.ActivityIDQueryParam {
	if o == nil {
		return nil
	}
	return o.ActivityID
}

func (o *UpsertEntityRequest) GetAsync() *bool {
	if o == nil {
		return nil
	}
	return o.Async
}

func (o *UpsertEntityRequest) GetDryRun() *bool {
	if o == nil {
		return nil
	}
	return o.DryRun
}

func (o *UpsertEntityRequest) GetFillActivity() *bool {
	if o == nil {
		return nil
	}
	return o.FillActivity
}

func (o *UpsertEntityRequest) GetSlug() string {
	if o == nil {
		return ""
	}
	return o.Slug
}

func (o *UpsertEntityRequest) GetStrict() *bool {
	if o == nil {
		return nil
	}
	return o.Strict
}

func (o *UpsertEntityRequest) GetValidate() *bool {
	if o == nil {
		return nil
	}
	return o.Validate
}

type UpsertEntityResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Entity was updated
	EntityItem *shared.EntityItem
	// Entity validation error when `?validate=true`
	EntityValidationV2ResultError *shared.EntityValidationV2ResultError
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *UpsertEntityResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpsertEntityResponse) GetEntityItem() *shared.EntityItem {
	if o == nil {
		return nil
	}
	return o.EntityItem
}

func (o *UpsertEntityResponse) GetEntityValidationV2ResultError() *shared.EntityValidationV2ResultError {
	if o == nil {
		return nil
	}
	return o.EntityValidationV2ResultError
}

func (o *UpsertEntityResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpsertEntityResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
