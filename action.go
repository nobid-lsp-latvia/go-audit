// SPDX-License-Identifier: EUPL-1.2

package audit

type Action string

const (
	// Successful account registration.
	RegistrationSuccess Action = "reg_success"
	// Failed account registration.
	RegistrationFailure Action = "reg_fail"
	// User successful login to the system.
	ActionLoginSuccess Action = "login_success"
	// User failed login to the system.
	ActionLoginFail Action = "login_fail"
	// User logout from the system.
	ActionLogout Action = "logout"
	// Resource detailed view.
	ActionView Action = "view"
	// Resource search.
	ActionSearch Action = "search"
	// New resource creation.
	ActionCreate Action = "create"
	// Existing resource modification.
	ActionModify Action = "modify"
	// Existing resource deletion.
	ActionDelete Action = "delete"
	// Export person data outside the system.
	ActionExport Action = "export"
	// Request for the person data form outer source.
	ServiceCall Action = "service_call"
)
