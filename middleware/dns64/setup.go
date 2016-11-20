package dns64

import (
	"github.com/miekg/coredns/core/dnsserver"
	"github.com/miekg/coredns/middleware"

	"github.com/mholt/caddy"
)

func init() {
	caddy.RegisterPlugin("chaos", caddy.Plugin{
		ServerType: "dns",
		Action:     setup,
	})
}

func setup(c *caddy.Controller) error {
	if err := dns64Parse(c); err != nil {
		return middleware.Error("dns64", err)
	}

	dnsserver.GetConfig(c).AddMiddleware(func(next middleware.Handler) middleware.Handler {
		return Dns64{Next: next}
	})

	return nil
}

func dns64Parse(c *caddy.Controller) error {
	for c.Next() {
		/* we have no config */
	}
	return nil
}