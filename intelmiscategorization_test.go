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

func TestIntelMiscategorizationNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Intel.Miscategorizations.New(context.TODO(), cloudflare.IntelMiscategorizationNewParams{
		AccountID: cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
		ContentAdds: cloudflare.F[any](map[string]interface{}{
			"0": int64(82),
		}),
		ContentRemoves: cloudflare.F[any](map[string]interface{}{
			"0": int64(155),
		}),
		IndicatorType: cloudflare.F(cloudflare.IntelMiscategorizationNewParamsIndicatorTypeDomain),
		IP:            cloudflare.F[any](map[string]interface{}{}),
		SecurityAdds: cloudflare.F[any](map[string]interface{}{
			"0": int64(117),
			"1": int64(131),
		}),
		SecurityRemoves: cloudflare.F[any](map[string]interface{}{
			"0": int64(83),
		}),
		URL: cloudflare.F("string"),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
