// File generated from our OpenAPI spec by Stainless.

package cloudflare_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-sdk-go"
	"github.com/cloudflare/cloudflare-sdk-go/internal/testutil"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

func TestZeroTrustDEXColoListWithOptionalParams(t *testing.T) {
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
	_, err := client.ZeroTrust.DEX.Colos.List(context.TODO(), cloudflare.ZeroTrustDEXColoListParams{
		AccountID: cloudflare.F("01a7362d577a6c3019a474fd6f485823"),
		TimeEnd:   cloudflare.F("2023-08-24T20:45:00Z"),
		TimeStart: cloudflare.F("2023-08-20T20:45:00Z"),
		SortBy:    cloudflare.F(cloudflare.ZeroTrustDEXColoListParamsSortByFleetStatusUsage),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
