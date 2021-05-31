package controller

// AuthRule
type AuthRule struct{}

// NewAuthRule is AUthRule exmaple
func NewAuthRule() AuthRule {
	return AuthRule{}
}

// authRuleGetValidate is get validate struct
type authRuleGetValidate struct {
	// Type: list、only
	Type string `form:"type" json:"type" xml:"type" binding:"required,oneof=list only"`
	// Search: type is only the search is id, type is list the search is id/name/url
	Search string `form:"search" json:"search" xml:"search" binding:"required_if=Type only"`
	// Page
	Page int `form:"page" json:"page" xml:"page" binding:"-"`
	// Limit
	Limit int `form:"limit" json:"limit" xml:"limit" binding:"-"`
	// Order
	Order string `form:"order" json:"order" xml:"order" binding:"-"`
}

// TODO:
type authRulePostValidate struct {
	// PID
	PID string `form:"pid" json:"pid" xml:"pid" binding:"required"`
	// Name
	Name string `form:"name" json:"name" xml:"name" binding:"required"`
}
