// SPDX-License-Identifier: EUPL-1.2

package audit

import (
	"fmt"

	"azugo.io/core/validation"

	"github.com/spf13/viper"
)

// Configuration is the configuration for the audit middleware.
type Configuration struct {
	Endpoint string `mapstructure:"endpoint" validate:"required,url"`
}

// Bind configuration section.
func (c *Configuration) Bind(prefix string, v *viper.Viper) {
	_ = v.BindEnv(fmt.Sprintf("%s.endpoint", prefix), "AUDIT_ENDPOINT")
}

// Validate application configuration.
func (c *Configuration) Validate(validate *validation.Validate) error {
	return validate.Struct(c)
}
