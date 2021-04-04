package caddy_cloudns

import (
	"github.com/caddyserver/caddy/v2"
	"github.com/caddyserver/caddy/v2/caddyconfig/caddyfile"
	"github.com/mynuolr/cloudns"
)

// Provider wraps the provider implementation as a Caddy module.
type Provider struct{ *cloudns.Provider }

func init() {
	caddy.RegisterModule(Provider{})
}

// CaddyModule returns the Caddy module information.
func (Provider) CaddyModule() caddy.ModuleInfo {
	return caddy.ModuleInfo{
		ID:  "dns.providers.cloudns",
		New: func() caddy.Module { return &Provider{new(cloudns.Provider)} },
	}
}

// TODO: This is just an example. Useful to allow env variable placeholders; update accordingly.
// Provision sets up the module. Implements caddy.Provisioner.
func (p *Provider) Provision(ctx caddy.Context) error {
	p.Provider.AuthPassword = caddy.NewReplacer().ReplaceAll(p.Provider.AuthPassword, "")
	p.Provider.AuthId = caddy.NewReplacer().ReplaceAll(p.Provider.AuthId, "")
	p.Provider.Sub = caddy.NewReplacer().ReplaceAll(p.Provider.AuthPassword, "false")
	return nil
}

// TODO: This is just an example. Update accordingly.
// UnmarshalCaddyfile sets up the DNS provider from Caddyfile tokens. Syntax:
//
// cloudns  {
//     auth_id       <auth_id>
//     auth_password <auth_password>
//     sub           <bool>
// }
//
// **THIS IS JUST AN EXAMPLE AND NEEDS TO BE CUSTOMIZED.**
func (p *Provider) UnmarshalCaddyfile(d *caddyfile.Dispenser) error {
	for d.Next() {
		if d.NextArg() {
			return d.ArgErr()
		}
		for nesting := d.Nesting(); d.NextBlock(nesting); {
			switch d.Val() {
			case "auth_id":
				if d.NextArg() {
					p.Provider.AuthId = d.Val()
				}
				if d.NextArg() {
					return d.ArgErr()
				}
			case "auth_password":
				if d.NextArg() {
					p.Provider.AuthId = d.Val()
				}
				if d.NextArg() {
					return d.ArgErr()
				}
			case "sub":
				if d.NextArg() {
					p.Provider.Sub = d.Val()
				}
				if d.NextArg() {
					return d.ArgErr()
				}
			default:
				return d.Errf("unrecognized subdirective '%s'", d.Val())
			}
		}
	}

	return nil
}

// Interface guards
var (
	_ caddyfile.Unmarshaler = (*Provider)(nil)
	_ caddy.Provisioner     = (*Provider)(nil)
)
