// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package radar

import (
	"context"
	"net/http"
	"net/url"
	"time"

	"github.com/cloudflare/cloudflare-go/v4/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v4/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v4/internal/param"
	"github.com/cloudflare/cloudflare-go/v4/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v4/option"
)

// AttackLayer3Service contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAttackLayer3Service] method instead.
type AttackLayer3Service struct {
	Options          []option.RequestOption
	Summary          *AttackLayer3SummaryService
	TimeseriesGroups *AttackLayer3TimeseriesGroupService
	Top              *AttackLayer3TopService
}

// NewAttackLayer3Service generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAttackLayer3Service(opts ...option.RequestOption) (r *AttackLayer3Service) {
	r = &AttackLayer3Service{}
	r.Options = opts
	r.Summary = NewAttackLayer3SummaryService(opts...)
	r.TimeseriesGroups = NewAttackLayer3TimeseriesGroupService(opts...)
	r.Top = NewAttackLayer3TopService(opts...)
	return
}

// Retrieves layer 3 attacks over time.
func (r *AttackLayer3Service) Timeseries(ctx context.Context, query AttackLayer3TimeseriesParams, opts ...option.RequestOption) (res *AttackLayer3TimeseriesResponse, err error) {
	var env AttackLayer3TimeseriesResponseEnvelope
	opts = append(r.Options[:], opts...)
	path := "radar/attacks/layer3/timeseries"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type AttackLayer3TimeseriesResponse struct {
	Meta   interface{}                          `json:"meta,required"`
	Serie0 AttackLayer3TimeseriesResponseSerie0 `json:"serie_0,required"`
	JSON   attackLayer3TimeseriesResponseJSON   `json:"-"`
}

// attackLayer3TimeseriesResponseJSON contains the JSON metadata for the struct
// [AttackLayer3TimeseriesResponse]
type attackLayer3TimeseriesResponseJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer3TimeseriesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer3TimeseriesResponseJSON) RawJSON() string {
	return r.raw
}

type AttackLayer3TimeseriesResponseSerie0 struct {
	Timestamps []time.Time                              `json:"timestamps,required" format:"date-time"`
	Values     []string                                 `json:"values,required"`
	JSON       attackLayer3TimeseriesResponseSerie0JSON `json:"-"`
}

// attackLayer3TimeseriesResponseSerie0JSON contains the JSON metadata for the
// struct [AttackLayer3TimeseriesResponseSerie0]
type attackLayer3TimeseriesResponseSerie0JSON struct {
	Timestamps  apijson.Field
	Values      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer3TimeseriesResponseSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer3TimeseriesResponseSerie0JSON) RawJSON() string {
	return r.raw
}

type AttackLayer3TimeseriesParams struct {
	// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
	// Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[AttackLayer3TimeseriesParamsAggInterval] `query:"aggInterval"`
	// Filters results by Autonomous System. Specify one or more Autonomous System
	// Numbers (ASNs) as a comma-separated list. Prefix with `-` to exclude ASNs from
	// results. For example, `-174, 3356` excludes results from AS174, but includes
	// results from AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// Filters results by continent. Specify a comma-separated list of alpha-2 codes.
	// Prefix with `-` to exclude continents from results. For example, `-EU,NA`
	// excludes results from EU, but includes results from NA.
	Continent param.Field[[]string] `query:"continent"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// Filters results by date range. For example, use `7d` and `7dcontrol` to compare
	// this week with the previous week. Use this parameter or set specific start and
	// end dates (`dateStart` and `dateEnd` parameters).
	DateRange param.Field[[]string] `query:"dateRange"`
	// Start of the date range.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Specifies whether the `location` filter applies to the source or target
	// location.
	Direction param.Field[AttackLayer3TimeseriesParamsDirection] `query:"direction"`
	// Format in which results will be returned.
	Format param.Field[AttackLayer3TimeseriesParamsFormat] `query:"format"`
	// Filters results by IP version (Ipv4 vs. IPv6).
	IPVersion param.Field[[]AttackLayer3TimeseriesParamsIPVersion] `query:"ipVersion"`
	// Filters results by location. Specify a comma-separated list of alpha-2 codes.
	// Prefix with `-` to exclude locations from results. For example, `-US,PT`
	// excludes results from the US, but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Measurement units, eg. bytes.
	Metric param.Field[AttackLayer3TimeseriesParamsMetric] `query:"metric"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization param.Field[AttackLayer3TimeseriesParamsNormalization] `query:"normalization"`
	// Filters the results by layer 3/4 protocol.
	Protocol param.Field[[]AttackLayer3TimeseriesParamsProtocol] `query:"protocol"`
}

// URLQuery serializes [AttackLayer3TimeseriesParams]'s query parameters as
// `url.Values`.
func (r AttackLayer3TimeseriesParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
// Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type AttackLayer3TimeseriesParamsAggInterval string

const (
	AttackLayer3TimeseriesParamsAggInterval15m AttackLayer3TimeseriesParamsAggInterval = "15m"
	AttackLayer3TimeseriesParamsAggInterval1h  AttackLayer3TimeseriesParamsAggInterval = "1h"
	AttackLayer3TimeseriesParamsAggInterval1d  AttackLayer3TimeseriesParamsAggInterval = "1d"
	AttackLayer3TimeseriesParamsAggInterval1w  AttackLayer3TimeseriesParamsAggInterval = "1w"
)

func (r AttackLayer3TimeseriesParamsAggInterval) IsKnown() bool {
	switch r {
	case AttackLayer3TimeseriesParamsAggInterval15m, AttackLayer3TimeseriesParamsAggInterval1h, AttackLayer3TimeseriesParamsAggInterval1d, AttackLayer3TimeseriesParamsAggInterval1w:
		return true
	}
	return false
}

// Specifies whether the `location` filter applies to the source or target
// location.
type AttackLayer3TimeseriesParamsDirection string

const (
	AttackLayer3TimeseriesParamsDirectionOrigin AttackLayer3TimeseriesParamsDirection = "ORIGIN"
	AttackLayer3TimeseriesParamsDirectionTarget AttackLayer3TimeseriesParamsDirection = "TARGET"
)

func (r AttackLayer3TimeseriesParamsDirection) IsKnown() bool {
	switch r {
	case AttackLayer3TimeseriesParamsDirectionOrigin, AttackLayer3TimeseriesParamsDirectionTarget:
		return true
	}
	return false
}

// Format in which results will be returned.
type AttackLayer3TimeseriesParamsFormat string

const (
	AttackLayer3TimeseriesParamsFormatJson AttackLayer3TimeseriesParamsFormat = "JSON"
	AttackLayer3TimeseriesParamsFormatCsv  AttackLayer3TimeseriesParamsFormat = "CSV"
)

func (r AttackLayer3TimeseriesParamsFormat) IsKnown() bool {
	switch r {
	case AttackLayer3TimeseriesParamsFormatJson, AttackLayer3TimeseriesParamsFormatCsv:
		return true
	}
	return false
}

type AttackLayer3TimeseriesParamsIPVersion string

const (
	AttackLayer3TimeseriesParamsIPVersionIPv4 AttackLayer3TimeseriesParamsIPVersion = "IPv4"
	AttackLayer3TimeseriesParamsIPVersionIPv6 AttackLayer3TimeseriesParamsIPVersion = "IPv6"
)

func (r AttackLayer3TimeseriesParamsIPVersion) IsKnown() bool {
	switch r {
	case AttackLayer3TimeseriesParamsIPVersionIPv4, AttackLayer3TimeseriesParamsIPVersionIPv6:
		return true
	}
	return false
}

// Measurement units, eg. bytes.
type AttackLayer3TimeseriesParamsMetric string

const (
	AttackLayer3TimeseriesParamsMetricBytes    AttackLayer3TimeseriesParamsMetric = "BYTES"
	AttackLayer3TimeseriesParamsMetricBytesOld AttackLayer3TimeseriesParamsMetric = "BYTES_OLD"
)

func (r AttackLayer3TimeseriesParamsMetric) IsKnown() bool {
	switch r {
	case AttackLayer3TimeseriesParamsMetricBytes, AttackLayer3TimeseriesParamsMetricBytesOld:
		return true
	}
	return false
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type AttackLayer3TimeseriesParamsNormalization string

const (
	AttackLayer3TimeseriesParamsNormalizationPercentageChange AttackLayer3TimeseriesParamsNormalization = "PERCENTAGE_CHANGE"
	AttackLayer3TimeseriesParamsNormalizationMin0Max          AttackLayer3TimeseriesParamsNormalization = "MIN0_MAX"
)

func (r AttackLayer3TimeseriesParamsNormalization) IsKnown() bool {
	switch r {
	case AttackLayer3TimeseriesParamsNormalizationPercentageChange, AttackLayer3TimeseriesParamsNormalizationMin0Max:
		return true
	}
	return false
}

type AttackLayer3TimeseriesParamsProtocol string

const (
	AttackLayer3TimeseriesParamsProtocolUdp  AttackLayer3TimeseriesParamsProtocol = "UDP"
	AttackLayer3TimeseriesParamsProtocolTCP  AttackLayer3TimeseriesParamsProtocol = "TCP"
	AttackLayer3TimeseriesParamsProtocolIcmp AttackLayer3TimeseriesParamsProtocol = "ICMP"
	AttackLayer3TimeseriesParamsProtocolGRE  AttackLayer3TimeseriesParamsProtocol = "GRE"
)

func (r AttackLayer3TimeseriesParamsProtocol) IsKnown() bool {
	switch r {
	case AttackLayer3TimeseriesParamsProtocolUdp, AttackLayer3TimeseriesParamsProtocolTCP, AttackLayer3TimeseriesParamsProtocolIcmp, AttackLayer3TimeseriesParamsProtocolGRE:
		return true
	}
	return false
}

type AttackLayer3TimeseriesResponseEnvelope struct {
	Result  AttackLayer3TimeseriesResponse             `json:"result,required"`
	Success bool                                       `json:"success,required"`
	JSON    attackLayer3TimeseriesResponseEnvelopeJSON `json:"-"`
}

// attackLayer3TimeseriesResponseEnvelopeJSON contains the JSON metadata for the
// struct [AttackLayer3TimeseriesResponseEnvelope]
type attackLayer3TimeseriesResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer3TimeseriesResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer3TimeseriesResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
