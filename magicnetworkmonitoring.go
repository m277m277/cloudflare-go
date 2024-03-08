// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"github.com/cloudflare/cloudflare-go/option"
)

// MagicNetworkMonitoringService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewMagicNetworkMonitoringService]
// method instead.
type MagicNetworkMonitoringService struct {
	Options []option.RequestOption
	Configs *MagicNetworkMonitoringConfigService
	Rules   *MagicNetworkMonitoringRuleService
}

// NewMagicNetworkMonitoringService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewMagicNetworkMonitoringService(opts ...option.RequestOption) (r *MagicNetworkMonitoringService) {
	r = &MagicNetworkMonitoringService{}
	r.Options = opts
	r.Configs = NewMagicNetworkMonitoringConfigService(opts...)
	r.Rules = NewMagicNetworkMonitoringRuleService(opts...)
	return
}
