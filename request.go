// SPDX-License-Identifier: EUPL-1.2

package audit

import (
	"azugo.io/core/validation"
)

// Person represents a person whose data is accessed.
type Person struct {
	// IdentifierType represents identifier type of the person whose data has been accessed.
	IdentifierType *string `json:"identifierType"`
	// Identifier represents identifier of the person whose data has been accessed.
	Identifier *string `json:"identifier"`
	// GivenName represents given name of the person whose data has been accessed.
	GivenName *string `json:"givenName"`
	// FamilyName represents family name of the person whose data has been accessed.
	FamilyName *string `json:"familyName"`
}

// Resource represents a resource accessed by user.
type Resource struct {
	// ID represents identifier of the resource.
	ID string `json:"id"`
	// Provider represents provider of the resource, e.g. RTU, CSDD, etc.
	Provider string `json:"provider"`
	// Type represents type of the resource.
	Type string `json:"type"`
}

// AuditRequest represents a request to save audit details.
type AuditRequest struct {
	// ClientID represents audit API caller identifier.
	ClientID string `json:"clientId"`
	// Endpoint represents the accessed endpoint.
	Endpoint *string `json:"endpoint"`
	// Action represents audit action.
	Action string `json:"action"`
	// RequestParameters represents data request parameters.
	RequestParameters *string `json:"requestParameters"`
	// Person represents the accessed endpoint.
	Person *Person `json:"person"`
	// Resources represents identifier ist of the accessed resources.
	Resources *[]Resource `json:"resources"`
	// IPAddress represents IP address of the caller.
	IPAddress *string `json:"ipAddress"`
	// UserAgent represents user agent of the caller.
	UserAgent *string `json:"userAgent"`
}

func (r *AuditRequest) Validate(validate *validation.Validate) error {
	return validate.Struct(r)
}
