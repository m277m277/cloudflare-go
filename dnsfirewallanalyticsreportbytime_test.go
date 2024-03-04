// File generated from our OpenAPI spec by Stainless.

package cloudflare_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/cloudflare/cloudflare-sdk-go"
	"github.com/cloudflare/cloudflare-sdk-go/internal/testutil"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

func TestDNSFirewallAnalyticsReportBytimeListWithOptionalParams(t *testing.T) {
	t.Skip("skipped: tests are disabled for the time being")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := cloudflare.NewClient(
		option.WithBaseURL(baseURL),
	)
	_, err := client.DNS.Firewall.Analytics.Reports.Bytimes.List(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"023e105f4ecef8ad9ca31a8372d0c353",
		cloudflare.DNSFirewallAnalyticsReportBytimeListParams{
			Dimensions: cloudflare.F("queryType"),
			Filters:    cloudflare.F("responseCode==NOERROR,queryType==A"),
			Limit:      cloudflare.F(int64(100)),
			Metrics:    cloudflare.F("queryCount,uncachedCount"),
			Since:      cloudflare.F(time.Now()),
			Sort:       cloudflare.F("+responseCode,-queryName"),
			TimeDelta:  cloudflare.F(cloudflare.DNSFirewallAnalyticsReportBytimeListParamsTimeDeltaHour),
			Until:      cloudflare.F(time.Now()),
		},
	)
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
