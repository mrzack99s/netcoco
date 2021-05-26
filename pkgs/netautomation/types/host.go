package types

import (
	"github.com/networklore/netrasp/pkg/netrasp"
)

type Host struct {
	Hostname     string
	Username     string
	Password     string
	EnableSecret string
	SSHPort      int
	Platform     string
	device       netrasp.Platform
}

func (host *Host) GetDriver() (device netrasp.Platform, err error) {
	switch host.Platform {
	case "ios":
		device, err = netrasp.New(host.Hostname,
			netrasp.WithDriver("ios"),
			netrasp.WithUsernamePasswordEnableSecret(host.Username, host.Password, host.EnableSecret),
		)
		if err != nil {
			return
		}
	case "sg300":
		device, err = netrasp.New(host.Hostname,
			netrasp.WithDriver("sg3xx"),
			netrasp.WithUsernamePasswordEnableSecret(host.Username, host.Password, host.EnableSecret),
			netrasp.WithSSHKeyExchange("diffie-hellman-group1-sha1"),
			netrasp.WithSSHCipher("aes128-cbc"),
		)
		if err != nil {
			return
		}
	case "sg350":
		device, err = netrasp.New(host.Hostname,
			netrasp.WithDriver("sg3xx"),
			netrasp.WithUsernamePasswordEnableSecret(host.Username, host.Password, host.EnableSecret),
			netrasp.WithSSHKeyExchange("diffie-hellman-group1-sha1"),
			netrasp.WithSSHCipher("aes256-ctr"),
		)
		if err != nil {
			return
		}
	}

	return
}
