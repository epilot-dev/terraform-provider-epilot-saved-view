// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/epilot-dev/terraform-provider-epilot-saved-view/internal/sdk/internal/utils"
)

// PaymentMethodRelationAttributeConstraints - A set of constraints applicable to the attribute.
// These constraints should and will be enforced by the attribute renderer.
type PaymentMethodRelationAttributeConstraints struct {
}

// PaymentMethodRelationAttributeInfoHelpers - A set of configurations meant to document and assist the user in filling the attribute.
type PaymentMethodRelationAttributeInfoHelpers struct {
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

func (o *PaymentMethodRelationAttributeInfoHelpers) GetHintCustomComponent() *string {
	if o == nil {
		return nil
	}
	return o.HintCustomComponent
}

func (o *PaymentMethodRelationAttributeInfoHelpers) GetHintText() *string {
	if o == nil {
		return nil
	}
	return o.HintText
}

func (o *PaymentMethodRelationAttributeInfoHelpers) GetHintTextKey() *string {
	if o == nil {
		return nil
	}
	return o.HintTextKey
}

func (o *PaymentMethodRelationAttributeInfoHelpers) GetHintTooltipPlacement() *string {
	if o == nil {
		return nil
	}
	return o.HintTooltipPlacement
}

type PaymentMethodRelationAttributeType string

const (
	PaymentMethodRelationAttributeTypeRelationPaymentMethod PaymentMethodRelationAttributeType = "relation_payment_method"
)

func (e PaymentMethodRelationAttributeType) ToPointer() *PaymentMethodRelationAttributeType {
	return &e
}
func (e *PaymentMethodRelationAttributeType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "relation_payment_method":
		*e = PaymentMethodRelationAttributeType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for PaymentMethodRelationAttributeType: %v", v)
	}
}

// PaymentMethodRelationAttribute - Reference to a payment method attribute of another entity
type PaymentMethodRelationAttribute struct {
	// Manifest ID used to create/update the schema attribute
	Manifest []string `json:"_manifest,omitempty"`
	Purpose  []string `json:"_purpose,omitempty"`
	// A set of constraints applicable to the attribute.
	// These constraints should and will be enforced by the attribute renderer.
	//
	Constraints  *PaymentMethodRelationAttributeConstraints `json:"constraints,omitempty"`
	DefaultValue any                                        `json:"default_value,omitempty"`
	Deprecated   *bool                                      `default:"false" json:"deprecated"`
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
	InfoHelpers *PaymentMethodRelationAttributeInfoHelpers `json:"info_helpers,omitempty"`
	Label       string                                     `json:"label"`
	Layout      *string                                    `json:"layout,omitempty"`
	Name        string                                     `json:"name"`
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
	Sortable       *bool                              `default:"true" json:"sortable"`
	Type           PaymentMethodRelationAttributeType `json:"type"`
	ValueFormatter *string                            `json:"value_formatter,omitempty"`
}

func (p PaymentMethodRelationAttribute) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(p, "", false)
}

func (p *PaymentMethodRelationAttribute) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &p, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *PaymentMethodRelationAttribute) GetManifest() []string {
	if o == nil {
		return nil
	}
	return o.Manifest
}

func (o *PaymentMethodRelationAttribute) GetPurpose() []string {
	if o == nil {
		return nil
	}
	return o.Purpose
}

func (o *PaymentMethodRelationAttribute) GetConstraints() *PaymentMethodRelationAttributeConstraints {
	if o == nil {
		return nil
	}
	return o.Constraints
}

func (o *PaymentMethodRelationAttribute) GetDefaultValue() any {
	if o == nil {
		return nil
	}
	return o.DefaultValue
}

func (o *PaymentMethodRelationAttribute) GetDeprecated() *bool {
	if o == nil {
		return nil
	}
	return o.Deprecated
}

func (o *PaymentMethodRelationAttribute) GetEntityBuilderDisableEdit() *bool {
	if o == nil {
		return nil
	}
	return o.EntityBuilderDisableEdit
}

func (o *PaymentMethodRelationAttribute) GetFeatureFlag() *string {
	if o == nil {
		return nil
	}
	return o.FeatureFlag
}

func (o *PaymentMethodRelationAttribute) GetGroup() *string {
	if o == nil {
		return nil
	}
	return o.Group
}

func (o *PaymentMethodRelationAttribute) GetHasPrimary() *bool {
	if o == nil {
		return nil
	}
	return o.HasPrimary
}

func (o *PaymentMethodRelationAttribute) GetHidden() *bool {
	if o == nil {
		return nil
	}
	return o.Hidden
}

func (o *PaymentMethodRelationAttribute) GetHideLabel() *bool {
	if o == nil {
		return nil
	}
	return o.HideLabel
}

func (o *PaymentMethodRelationAttribute) GetIcon() *string {
	if o == nil {
		return nil
	}
	return o.Icon
}

func (o *PaymentMethodRelationAttribute) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *PaymentMethodRelationAttribute) GetInfoHelpers() *PaymentMethodRelationAttributeInfoHelpers {
	if o == nil {
		return nil
	}
	return o.InfoHelpers
}

func (o *PaymentMethodRelationAttribute) GetLabel() string {
	if o == nil {
		return ""
	}
	return o.Label
}

func (o *PaymentMethodRelationAttribute) GetLayout() *string {
	if o == nil {
		return nil
	}
	return o.Layout
}

func (o *PaymentMethodRelationAttribute) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *PaymentMethodRelationAttribute) GetOrder() *int64 {
	if o == nil {
		return nil
	}
	return o.Order
}

func (o *PaymentMethodRelationAttribute) GetPlaceholder() *string {
	if o == nil {
		return nil
	}
	return o.Placeholder
}

func (o *PaymentMethodRelationAttribute) GetPreviewValueFormatter() *string {
	if o == nil {
		return nil
	}
	return o.PreviewValueFormatter
}

func (o *PaymentMethodRelationAttribute) GetProtected() *bool {
	if o == nil {
		return nil
	}
	return o.Protected
}

func (o *PaymentMethodRelationAttribute) GetReadonly() *bool {
	if o == nil {
		return nil
	}
	return o.Readonly
}

func (o *PaymentMethodRelationAttribute) GetRenderCondition() *string {
	if o == nil {
		return nil
	}
	return o.RenderCondition
}

func (o *PaymentMethodRelationAttribute) GetRepeatable() *bool {
	if o == nil {
		return nil
	}
	return o.Repeatable
}

func (o *PaymentMethodRelationAttribute) GetRequired() *bool {
	if o == nil {
		return nil
	}
	return o.Required
}

func (o *PaymentMethodRelationAttribute) GetSettingsFlag() []SettingFlag {
	if o == nil {
		return nil
	}
	return o.SettingsFlag
}

func (o *PaymentMethodRelationAttribute) GetShowInTable() *bool {
	if o == nil {
		return nil
	}
	return o.ShowInTable
}

func (o *PaymentMethodRelationAttribute) GetSortable() *bool {
	if o == nil {
		return nil
	}
	return o.Sortable
}

func (o *PaymentMethodRelationAttribute) GetType() PaymentMethodRelationAttributeType {
	if o == nil {
		return PaymentMethodRelationAttributeType("")
	}
	return o.Type
}

func (o *PaymentMethodRelationAttribute) GetValueFormatter() *string {
	if o == nil {
		return nil
	}
	return o.ValueFormatter
}
