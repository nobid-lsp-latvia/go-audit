// SPDX-License-Identifier: EUPL-1.2

package audit

import (
	"fmt"

	"azugo.io/azugo"
	"azugo.io/core/http"
	"github.com/valyala/fasthttp"
)

type audit struct {
	config          Configuration
	sensitiveParams map[string][]string
}

// Audit interface
type Audit interface {
	// PersonRequest audits person request.
	PersonRequest(ctx *azugo.Context, request AuditRequest, opts ...http.RequestOption) error
}

// New creates a new audit configuration.
func New(configuration *Configuration, sensitiveParams map[string][]string) Audit {
	if configuration == nil {
		configuration = &Configuration{}
	}

	a := &audit{
		config:          *configuration,
		sensitiveParams: sensitiveParams,
	}

	return a
}

func (a *audit) PersonRequest(ctx *azugo.Context, request AuditRequest, opts ...http.RequestOption) error {
	client := ctx.HTTPClient().WithBaseURL(a.config.Endpoint)

	if len(opts) == 0 {
		opts = []http.RequestOption{ctx.Header.InheritAuthorization()}
	}

	endpoint := fmt.Sprintf("%s %s", ctx.Context().Method(), ctx.RouterPath())
	request.Endpoint = &endpoint

	filteredParams := a.getFilteredParams(endpoint, ctx.Context().URI().QueryArgs()).String()
	request.RequestParameters = &filteredParams

	ip := ctx.IP().String()
	request.IPAddress = &ip

	userAgent := ctx.UserAgent()
	request.UserAgent = &userAgent

	return client.PostJSON("/audit", request, nil, opts...)
}

// getFilteredParams filters out sensitive query parameters based on the request endpoint.
func (a *audit) getFilteredParams(endpoint string, queryArgs *fasthttp.Args) *fasthttp.Args {
	if len(a.sensitiveParams) == 0 {
		return queryArgs
	}

	paramsToExclude, ok := a.sensitiveParams[endpoint]
	if !ok || len(paramsToExclude) == 0 {
		return queryArgs
	}

	filteredArgs := fasthttp.AcquireArgs()
	queryArgs.VisitAll(func(key, value []byte) {
		for _, param := range paramsToExclude {
			if string(key) == param {
				return
			}
		}

		filteredArgs.AddBytesKV(key, value)
	})

	return filteredArgs
}
