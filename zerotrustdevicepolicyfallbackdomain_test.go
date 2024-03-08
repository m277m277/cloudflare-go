// File generated from our OpenAPI spec by Stainless.

package cloudflare_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go"
	"github.com/cloudflare/cloudflare-go/internal/testutil"
	"github.com/cloudflare/cloudflare-go/option"
)

func TestZeroTrustDevicePolicyFallbackDomainUpdate(t *testing.T) {
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
		option.WithAPIKey("144c9defac04969c7bfad8efaa8ea194"),
		option.WithAPIEmail("user@example.com"),
	)
	_, err := client.ZeroTrust.Devices.Policies.FallbackDomains.Update(
		context.TODO(),
		"f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
		cloudflare.ZeroTrustDevicePolicyFallbackDomainUpdateParams{
			AccountID: cloudflare.F[any]("699d98642c564d2e855e9661899b7252"),
			Body: cloudflare.F([]cloudflare.ZeroTrustDevicePolicyFallbackDomainUpdateParamsBody{{
				Description: cloudflare.F("Domain bypass for local development"),
				DNSServer:   cloudflare.F([]interface{}{map[string]interface{}{}, map[string]interface{}{}, map[string]interface{}{}}),
				Suffix:      cloudflare.F("example.com"),
			}, {
				Description: cloudflare.F("Domain bypass for local development"),
				DNSServer:   cloudflare.F([]interface{}{map[string]interface{}{}, map[string]interface{}{}, map[string]interface{}{}}),
				Suffix:      cloudflare.F("example.com"),
			}, {
				Description: cloudflare.F("Domain bypass for local development"),
				DNSServer:   cloudflare.F([]interface{}{map[string]interface{}{}, map[string]interface{}{}, map[string]interface{}{}}),
				Suffix:      cloudflare.F("example.com"),
			}}),
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

func TestZeroTrustDevicePolicyFallbackDomainList(t *testing.T) {
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
		option.WithAPIKey("144c9defac04969c7bfad8efaa8ea194"),
		option.WithAPIEmail("user@example.com"),
	)
	_, err := client.ZeroTrust.Devices.Policies.FallbackDomains.List(context.TODO(), cloudflare.ZeroTrustDevicePolicyFallbackDomainListParams{
		AccountID: cloudflare.F[any]("699d98642c564d2e855e9661899b7252"),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestZeroTrustDevicePolicyFallbackDomainGet(t *testing.T) {
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
		option.WithAPIKey("144c9defac04969c7bfad8efaa8ea194"),
		option.WithAPIEmail("user@example.com"),
	)
	_, err := client.ZeroTrust.Devices.Policies.FallbackDomains.Get(
		context.TODO(),
		"f174e90a-fafe-4643-bbbc-4a0ed4fc8415",
		cloudflare.ZeroTrustDevicePolicyFallbackDomainGetParams{
			AccountID: cloudflare.F[any]("699d98642c564d2e855e9661899b7252"),
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
