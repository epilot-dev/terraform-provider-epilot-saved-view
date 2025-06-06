// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

// SavedViewPartial - A saved entity view
type SavedViewPartial struct {
	// User-friendly identifier for the saved view
	Name *string `json:"name,omitempty"`
	// boolean property for if a view is shared with organisation
	Shared *bool `json:"shared,omitempty"`
	// list of schemas a view can belong to
	Slug     []string `json:"slug,omitempty"`
	UIConfig any      `json:"ui_config,omitempty"`
}

func (o *SavedViewPartial) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *SavedViewPartial) GetShared() *bool {
	if o == nil {
		return nil
	}
	return o.Shared
}

func (o *SavedViewPartial) GetSlug() []string {
	if o == nil {
		return nil
	}
	return o.Slug
}

func (o *SavedViewPartial) GetUIConfig() any {
	if o == nil {
		return nil
	}
	return o.UIConfig
}
