// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/epilot-dev/terraform-provider-epilot-saved-view/internal/sdk/internal/utils"
)

// FileAttributeConstraints - A set of constraints applicable to the attribute.
// These constraints should and will be enforced by the attribute renderer.
type FileAttributeConstraints struct {
}

type FileAttributeDefaultAccessControl string

const (
	FileAttributeDefaultAccessControlPublicRead FileAttributeDefaultAccessControl = "public-read"
	FileAttributeDefaultAccessControlPrivate    FileAttributeDefaultAccessControl = "private"
)

func (e FileAttributeDefaultAccessControl) ToPointer() *FileAttributeDefaultAccessControl {
	return &e
}
func (e *FileAttributeDefaultAccessControl) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "public-read":
		fallthrough
	case "private":
		*e = FileAttributeDefaultAccessControl(v)
		return nil
	default:
		return fmt.Errorf("invalid value for FileAttributeDefaultAccessControl: %v", v)
	}
}

// FileAttributeInfoHelpers - A set of configurations meant to document and assist the user in filling the attribute.
type FileAttributeInfoHelpers struct {
	// The name of the custom component to be used as the hint helper.
	// The component should be registered in the `@epilot360/entity-ui` on the index of the components directory.
	// When specified it overrides the `hint_text` or `hint_text_key` configuration.
	//
	HintCustomComponent *string `json:"hint_custom_component,omitempty"`
	// The text to be displayed in the attribute hint helper.
	// When specified it overrides the `hint_text_key` configuration.
	//
	HintText *string `json:"hint_text,omitempty"`
	// The key of the hint text to be displayed in the attribute hint helper.
	// The key should be a valid i18n key.
	//
	HintTextKey *string `json:"hint_text_key,omitempty"`
	// The placement of the hint tooltip.
	// The value should be a valid `@mui/core` tooltip placement.
	//
	HintTooltipPlacement *string `json:"hint_tooltip_placement,omitempty"`
}

func (o *FileAttributeInfoHelpers) GetHintCustomComponent() *string {
	if o == nil {
		return nil
	}
	return o.HintCustomComponent
}

func (o *FileAttributeInfoHelpers) GetHintText() *string {
	if o == nil {
		return nil
	}
	return o.HintText
}

func (o *FileAttributeInfoHelpers) GetHintTextKey() *string {
	if o == nil {
		return nil
	}
	return o.HintTextKey
}

func (o *FileAttributeInfoHelpers) GetHintTooltipPlacement() *string {
	if o == nil {
		return nil
	}
	return o.HintTooltipPlacement
}

type FileAttributeType string

const (
	FileAttributeTypeImage FileAttributeType = "image"
	FileAttributeTypeFile  FileAttributeType = "file"
)

func (e FileAttributeType) ToPointer() *FileAttributeType {
	return &e
}
func (e *FileAttributeType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "image":
		fallthrough
	case "file":
		*e = FileAttributeType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for FileAttributeType: %v", v)
	}
}

// FileAttribute - File or Image Attachment
type FileAttribute struct {
	// Manifest ID used to create/update the schema attribute
	Manifest []string `json:"_manifest,omitempty"`
	Purpose  []string `json:"_purpose,omitempty"`
	// List of file extensions (without the dot suffix)
	AllowedExtensions []string `json:"allowed_extensions,omitempty"`
	// A set of constraints applicable to the attribute.
	// These constraints should and will be enforced by the attribute renderer.
	//
	Constraints          *FileAttributeConstraints          `json:"constraints,omitempty"`
	DefaultAccessControl *FileAttributeDefaultAccessControl `json:"default_access_control,omitempty"`
	DefaultValue         any                                `json:"default_value,omitempty"`
	Deprecated           *bool                              `default:"false" json:"deprecated"`
	// Controls how the images are presented to the user during upload on the Entity Details view.
	DisplayImagesLandscaped *bool `json:"display_images_landscaped,omitempty"`
	// When set to true, an i18n description will be used alongside the attribute label.
	// This description should be set through the platform locales in the form: `file.{attribute_name}.description_text`.
	//
	EnableDescription *bool `json:"enable_description,omitempty"`
	// Setting to `true` disables editing the attribute on the entity builder UI
	EntityBuilderDisableEdit *bool `default:"false" json:"entity_builder_disable_edit"`
	// This attribute should only be active when the feature flag is enabled
	FeatureFlag *string `json:"feature_flag,omitempty"`
	// Which group the attribute should appear in. Accepts group ID or group name
	Group      *string `json:"group,omitempty"`
	HasPrimary *bool   `json:"has_primary,omitempty"`
	// Do not render attribute in entity views
	Hidden *bool `default:"false" json:"hidden"`
	// When set to true, will hide the label of the field.
	HideLabel *bool `json:"hide_label,omitempty"`
	// Code name of the icon to used to represent this attribute.
	// The value must be a valid @epilot/base-elements Icon name
	//
	Icon *string `json:"icon,omitempty"`
	// ID for the entity attribute
	ID *string `json:"id,omitempty"`
	// A set of configurations meant to document and assist the user in filling the attribute.
	InfoHelpers *FileAttributeInfoHelpers `json:"info_helpers,omitempty"`
	Label       string                    `json:"label"`
	Layout      *string                   `json:"layout,omitempty"`
	Multiple    *bool                     `json:"multiple,omitempty"`
	Name        string                    `json:"name"`
	// Attribute sort order (ascending) in group
	Order                 *int64  `json:"order,omitempty"`
	Placeholder           *string `json:"placeholder,omitempty"`
	PreviewValueFormatter *string `json:"preview_value_formatter,omitempty"`
	// Setting to `true` prevents the attribute from being modified / deleted
	Protected *bool `json:"protected,omitempty"`
	Readonly  *bool `default:"false" json:"readonly"`
	// Defines the conditional rendering expression for showing this field.
	// When a valid expression is parsed, their evaluation defines the visibility of this attribute.
	// Note: Empty or invalid expression have no effect on the field visibility.
	//
	RenderCondition *string `json:"render_condition,omitempty"`
	// The attribute is a repeatable
	Repeatable *bool `json:"repeatable,omitempty"`
	Required   *bool `default:"false" json:"required"`
	// This attribute should only be active when one of the provided settings have the correct value
	SettingsFlag []SettingFlag `json:"settings_flag,omitempty"`
	// Render as a column in table views. When defined, overrides `hidden`
	ShowInTable *bool `json:"show_in_table,omitempty"`
	// Allow sorting by this attribute in table views if `show_in_table` is true
	Sortable       *bool             `default:"true" json:"sortable"`
	Type           FileAttributeType `json:"type"`
	ValueFormatter *string           `json:"value_formatter,omitempty"`
}

func (f FileAttribute) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(f, "", false)
}

func (f *FileAttribute) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &f, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *FileAttribute) GetManifest() []string {
	if o == nil {
		return nil
	}
	return o.Manifest
}

func (o *FileAttribute) GetPurpose() []string {
	if o == nil {
		return nil
	}
	return o.Purpose
}

func (o *FileAttribute) GetAllowedExtensions() []string {
	if o == nil {
		return nil
	}
	return o.AllowedExtensions
}

func (o *FileAttribute) GetConstraints() *FileAttributeConstraints {
	if o == nil {
		return nil
	}
	return o.Constraints
}

func (o *FileAttribute) GetDefaultAccessControl() *FileAttributeDefaultAccessControl {
	if o == nil {
		return nil
	}
	return o.DefaultAccessControl
}

func (o *FileAttribute) GetDefaultValue() any {
	if o == nil {
		return nil
	}
	return o.DefaultValue
}

func (o *FileAttribute) GetDeprecated() *bool {
	if o == nil {
		return nil
	}
	return o.Deprecated
}

func (o *FileAttribute) GetDisplayImagesLandscaped() *bool {
	if o == nil {
		return nil
	}
	return o.DisplayImagesLandscaped
}

func (o *FileAttribute) GetEnableDescription() *bool {
	if o == nil {
		return nil
	}
	return o.EnableDescription
}

func (o *FileAttribute) GetEntityBuilderDisableEdit() *bool {
	if o == nil {
		return nil
	}
	return o.EntityBuilderDisableEdit
}

func (o *FileAttribute) GetFeatureFlag() *string {
	if o == nil {
		return nil
	}
	return o.FeatureFlag
}

func (o *FileAttribute) GetGroup() *string {
	if o == nil {
		return nil
	}
	return o.Group
}

func (o *FileAttribute) GetHasPrimary() *bool {
	if o == nil {
		return nil
	}
	return o.HasPrimary
}

func (o *FileAttribute) GetHidden() *bool {
	if o == nil {
		return nil
	}
	return o.Hidden
}

func (o *FileAttribute) GetHideLabel() *bool {
	if o == nil {
		return nil
	}
	return o.HideLabel
}

func (o *FileAttribute) GetIcon() *string {
	if o == nil {
		return nil
	}
	return o.Icon
}

func (o *FileAttribute) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *FileAttribute) GetInfoHelpers() *FileAttributeInfoHelpers {
	if o == nil {
		return nil
	}
	return o.InfoHelpers
}

func (o *FileAttribute) GetLabel() string {
	if o == nil {
		return ""
	}
	return o.Label
}

func (o *FileAttribute) GetLayout() *string {
	if o == nil {
		return nil
	}
	return o.Layout
}

func (o *FileAttribute) GetMultiple() *bool {
	if o == nil {
		return nil
	}
	return o.Multiple
}

func (o *FileAttribute) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *FileAttribute) GetOrder() *int64 {
	if o == nil {
		return nil
	}
	return o.Order
}

func (o *FileAttribute) GetPlaceholder() *string {
	if o == nil {
		return nil
	}
	return o.Placeholder
}

func (o *FileAttribute) GetPreviewValueFormatter() *string {
	if o == nil {
		return nil
	}
	return o.PreviewValueFormatter
}

func (o *FileAttribute) GetProtected() *bool {
	if o == nil {
		return nil
	}
	return o.Protected
}

func (o *FileAttribute) GetReadonly() *bool {
	if o == nil {
		return nil
	}
	return o.Readonly
}

func (o *FileAttribute) GetRenderCondition() *string {
	if o == nil {
		return nil
	}
	return o.RenderCondition
}

func (o *FileAttribute) GetRepeatable() *bool {
	if o == nil {
		return nil
	}
	return o.Repeatable
}

func (o *FileAttribute) GetRequired() *bool {
	if o == nil {
		return nil
	}
	return o.Required
}

func (o *FileAttribute) GetSettingsFlag() []SettingFlag {
	if o == nil {
		return nil
	}
	return o.SettingsFlag
}

func (o *FileAttribute) GetShowInTable() *bool {
	if o == nil {
		return nil
	}
	return o.ShowInTable
}

func (o *FileAttribute) GetSortable() *bool {
	if o == nil {
		return nil
	}
	return o.Sortable
}

func (o *FileAttribute) GetType() FileAttributeType {
	if o == nil {
		return FileAttributeType("")
	}
	return o.Type
}

func (o *FileAttribute) GetValueFormatter() *string {
	if o == nil {
		return nil
	}
	return o.ValueFormatter
}
