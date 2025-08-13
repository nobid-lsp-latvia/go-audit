// SPDX-License-Identifier: EUPL-1.2

package audit

type Role string

const (
	// Role for third party activities
	RoleThirdParty Role = "THIRD_PARTY"
	// Role for person activities
	RolePerson = "PERSON"
	// Role for admin activities
	RoleAdmin = "ADMIN"
)
