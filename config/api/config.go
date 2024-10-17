package api

import "github.com/crossplane/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("unleash_api_token", func(r *config.Resource) {
		r.ShortGroup = "api"
	})
}
