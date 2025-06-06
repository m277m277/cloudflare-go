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

// QualityIQIService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewQualityIQIService] method instead.
type QualityIQIService struct {
	Options []option.RequestOption
}

// NewQualityIQIService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewQualityIQIService(opts ...option.RequestOption) (r *QualityIQIService) {
	r = &QualityIQIService{}
	r.Options = opts
	return
}

// Retrieves a summary (percentiles) of bandwidth, latency, or DNS response time
// from the Radar Internet Quality Index (IQI).
func (r *QualityIQIService) Summary(ctx context.Context, query QualityIQISummaryParams, opts ...option.RequestOption) (res *QualityIQISummaryResponse, err error) {
	var env QualityIQISummaryResponseEnvelope
	opts = append(r.Options[:], opts...)
	path := "radar/quality/iqi/summary"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Retrieves a time series (percentiles) of bandwidth, latency, or DNS response
// time from the Radar Internet Quality Index (IQI).
func (r *QualityIQIService) TimeseriesGroups(ctx context.Context, query QualityIQITimeseriesGroupsParams, opts ...option.RequestOption) (res *QualityIQITimeseriesGroupsResponse, err error) {
	var env QualityIQITimeseriesGroupsResponseEnvelope
	opts = append(r.Options[:], opts...)
	path := "radar/quality/iqi/timeseries_groups"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type QualityIQISummaryResponse struct {
	Meta     QualityIQISummaryResponseMeta     `json:"meta,required"`
	Summary0 QualityIQISummaryResponseSummary0 `json:"summary_0,required"`
	JSON     qualityIQISummaryResponseJSON     `json:"-"`
}

// qualityIQISummaryResponseJSON contains the JSON metadata for the struct
// [QualityIQISummaryResponse]
type qualityIQISummaryResponseJSON struct {
	Meta        apijson.Field
	Summary0    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *QualityIQISummaryResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r qualityIQISummaryResponseJSON) RawJSON() string {
	return r.raw
}

type QualityIQISummaryResponseMeta struct {
	DateRange      []QualityIQISummaryResponseMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                      `json:"lastUpdated,required"`
	Normalization  string                                      `json:"normalization,required"`
	ConfidenceInfo QualityIQISummaryResponseMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           qualityIQISummaryResponseMetaJSON           `json:"-"`
}

// qualityIQISummaryResponseMetaJSON contains the JSON metadata for the struct
// [QualityIQISummaryResponseMeta]
type qualityIQISummaryResponseMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *QualityIQISummaryResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r qualityIQISummaryResponseMetaJSON) RawJSON() string {
	return r.raw
}

type QualityIQISummaryResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                  `json:"startTime,required" format:"date-time"`
	JSON      qualityIQISummaryResponseMetaDateRangeJSON `json:"-"`
}

// qualityIQISummaryResponseMetaDateRangeJSON contains the JSON metadata for the
// struct [QualityIQISummaryResponseMetaDateRange]
type qualityIQISummaryResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *QualityIQISummaryResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r qualityIQISummaryResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

type QualityIQISummaryResponseMetaConfidenceInfo struct {
	Annotations []QualityIQISummaryResponseMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                   `json:"level"`
	JSON        qualityIQISummaryResponseMetaConfidenceInfoJSON         `json:"-"`
}

// qualityIQISummaryResponseMetaConfidenceInfoJSON contains the JSON metadata for
// the struct [QualityIQISummaryResponseMetaConfidenceInfo]
type qualityIQISummaryResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *QualityIQISummaryResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r qualityIQISummaryResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

type QualityIQISummaryResponseMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                    `json:"dataSource,required"`
	Description     string                                                    `json:"description,required"`
	EventType       string                                                    `json:"eventType,required"`
	IsInstantaneous bool                                                      `json:"isInstantaneous,required"`
	EndTime         time.Time                                                 `json:"endTime" format:"date-time"`
	LinkedURL       string                                                    `json:"linkedUrl"`
	StartTime       time.Time                                                 `json:"startTime" format:"date-time"`
	JSON            qualityIQISummaryResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// qualityIQISummaryResponseMetaConfidenceInfoAnnotationJSON contains the JSON
// metadata for the struct [QualityIQISummaryResponseMetaConfidenceInfoAnnotation]
type qualityIQISummaryResponseMetaConfidenceInfoAnnotationJSON struct {
	DataSource      apijson.Field
	Description     apijson.Field
	EventType       apijson.Field
	IsInstantaneous apijson.Field
	EndTime         apijson.Field
	LinkedURL       apijson.Field
	StartTime       apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *QualityIQISummaryResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r qualityIQISummaryResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type QualityIQISummaryResponseSummary0 struct {
	P25  string                                `json:"p25,required"`
	P50  string                                `json:"p50,required"`
	P75  string                                `json:"p75,required"`
	JSON qualityIQISummaryResponseSummary0JSON `json:"-"`
}

// qualityIQISummaryResponseSummary0JSON contains the JSON metadata for the struct
// [QualityIQISummaryResponseSummary0]
type qualityIQISummaryResponseSummary0JSON struct {
	P25         apijson.Field
	P50         apijson.Field
	P75         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *QualityIQISummaryResponseSummary0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r qualityIQISummaryResponseSummary0JSON) RawJSON() string {
	return r.raw
}

type QualityIQITimeseriesGroupsResponse struct {
	Meta   interface{}                              `json:"meta,required"`
	Serie0 QualityIQITimeseriesGroupsResponseSerie0 `json:"serie_0,required"`
	JSON   qualityIQITimeseriesGroupsResponseJSON   `json:"-"`
}

// qualityIQITimeseriesGroupsResponseJSON contains the JSON metadata for the struct
// [QualityIQITimeseriesGroupsResponse]
type qualityIQITimeseriesGroupsResponseJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *QualityIQITimeseriesGroupsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r qualityIQITimeseriesGroupsResponseJSON) RawJSON() string {
	return r.raw
}

type QualityIQITimeseriesGroupsResponseSerie0 struct {
	P25        []string                                     `json:"p25,required"`
	P50        []string                                     `json:"p50,required"`
	P75        []string                                     `json:"p75,required"`
	Timestamps []string                                     `json:"timestamps,required"`
	JSON       qualityIQITimeseriesGroupsResponseSerie0JSON `json:"-"`
}

// qualityIQITimeseriesGroupsResponseSerie0JSON contains the JSON metadata for the
// struct [QualityIQITimeseriesGroupsResponseSerie0]
type qualityIQITimeseriesGroupsResponseSerie0JSON struct {
	P25         apijson.Field
	P50         apijson.Field
	P75         apijson.Field
	Timestamps  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *QualityIQITimeseriesGroupsResponseSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r qualityIQITimeseriesGroupsResponseSerie0JSON) RawJSON() string {
	return r.raw
}

type QualityIQISummaryParams struct {
	// Defines which metric to return (bandwidth, latency, or DNS response time).
	Metric param.Field[QualityIQISummaryParamsMetric] `query:"metric,required"`
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
	// Format in which results will be returned.
	Format param.Field[QualityIQISummaryParamsFormat] `query:"format"`
	// Filters results by location. Specify a comma-separated list of alpha-2 codes.
	// Prefix with `-` to exclude locations from results. For example, `-US,PT`
	// excludes results from the US, but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [QualityIQISummaryParams]'s query parameters as
// `url.Values`.
func (r QualityIQISummaryParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Defines which metric to return (bandwidth, latency, or DNS response time).
type QualityIQISummaryParamsMetric string

const (
	QualityIQISummaryParamsMetricBandwidth QualityIQISummaryParamsMetric = "BANDWIDTH"
	QualityIQISummaryParamsMetricDNS       QualityIQISummaryParamsMetric = "DNS"
	QualityIQISummaryParamsMetricLatency   QualityIQISummaryParamsMetric = "LATENCY"
)

func (r QualityIQISummaryParamsMetric) IsKnown() bool {
	switch r {
	case QualityIQISummaryParamsMetricBandwidth, QualityIQISummaryParamsMetricDNS, QualityIQISummaryParamsMetricLatency:
		return true
	}
	return false
}

// Format in which results will be returned.
type QualityIQISummaryParamsFormat string

const (
	QualityIQISummaryParamsFormatJson QualityIQISummaryParamsFormat = "JSON"
	QualityIQISummaryParamsFormatCsv  QualityIQISummaryParamsFormat = "CSV"
)

func (r QualityIQISummaryParamsFormat) IsKnown() bool {
	switch r {
	case QualityIQISummaryParamsFormatJson, QualityIQISummaryParamsFormatCsv:
		return true
	}
	return false
}

type QualityIQISummaryResponseEnvelope struct {
	Result  QualityIQISummaryResponse             `json:"result,required"`
	Success bool                                  `json:"success,required"`
	JSON    qualityIQISummaryResponseEnvelopeJSON `json:"-"`
}

// qualityIQISummaryResponseEnvelopeJSON contains the JSON metadata for the struct
// [QualityIQISummaryResponseEnvelope]
type qualityIQISummaryResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *QualityIQISummaryResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r qualityIQISummaryResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type QualityIQITimeseriesGroupsParams struct {
	// Defines which metric to return (bandwidth, latency, or DNS response time).
	Metric param.Field[QualityIQITimeseriesGroupsParamsMetric] `query:"metric,required"`
	// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
	// Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[QualityIQITimeseriesGroupsParamsAggInterval] `query:"aggInterval"`
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
	// Format in which results will be returned.
	Format param.Field[QualityIQITimeseriesGroupsParamsFormat] `query:"format"`
	// Enables interpolation for all series (using the average).
	Interpolation param.Field[bool] `query:"interpolation"`
	// Filters results by location. Specify a comma-separated list of alpha-2 codes.
	// Prefix with `-` to exclude locations from results. For example, `-US,PT`
	// excludes results from the US, but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [QualityIQITimeseriesGroupsParams]'s query parameters as
// `url.Values`.
func (r QualityIQITimeseriesGroupsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Defines which metric to return (bandwidth, latency, or DNS response time).
type QualityIQITimeseriesGroupsParamsMetric string

const (
	QualityIQITimeseriesGroupsParamsMetricBandwidth QualityIQITimeseriesGroupsParamsMetric = "BANDWIDTH"
	QualityIQITimeseriesGroupsParamsMetricDNS       QualityIQITimeseriesGroupsParamsMetric = "DNS"
	QualityIQITimeseriesGroupsParamsMetricLatency   QualityIQITimeseriesGroupsParamsMetric = "LATENCY"
)

func (r QualityIQITimeseriesGroupsParamsMetric) IsKnown() bool {
	switch r {
	case QualityIQITimeseriesGroupsParamsMetricBandwidth, QualityIQITimeseriesGroupsParamsMetricDNS, QualityIQITimeseriesGroupsParamsMetricLatency:
		return true
	}
	return false
}

// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
// Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type QualityIQITimeseriesGroupsParamsAggInterval string

const (
	QualityIQITimeseriesGroupsParamsAggInterval15m QualityIQITimeseriesGroupsParamsAggInterval = "15m"
	QualityIQITimeseriesGroupsParamsAggInterval1h  QualityIQITimeseriesGroupsParamsAggInterval = "1h"
	QualityIQITimeseriesGroupsParamsAggInterval1d  QualityIQITimeseriesGroupsParamsAggInterval = "1d"
	QualityIQITimeseriesGroupsParamsAggInterval1w  QualityIQITimeseriesGroupsParamsAggInterval = "1w"
)

func (r QualityIQITimeseriesGroupsParamsAggInterval) IsKnown() bool {
	switch r {
	case QualityIQITimeseriesGroupsParamsAggInterval15m, QualityIQITimeseriesGroupsParamsAggInterval1h, QualityIQITimeseriesGroupsParamsAggInterval1d, QualityIQITimeseriesGroupsParamsAggInterval1w:
		return true
	}
	return false
}

// Format in which results will be returned.
type QualityIQITimeseriesGroupsParamsFormat string

const (
	QualityIQITimeseriesGroupsParamsFormatJson QualityIQITimeseriesGroupsParamsFormat = "JSON"
	QualityIQITimeseriesGroupsParamsFormatCsv  QualityIQITimeseriesGroupsParamsFormat = "CSV"
)

func (r QualityIQITimeseriesGroupsParamsFormat) IsKnown() bool {
	switch r {
	case QualityIQITimeseriesGroupsParamsFormatJson, QualityIQITimeseriesGroupsParamsFormatCsv:
		return true
	}
	return false
}

type QualityIQITimeseriesGroupsResponseEnvelope struct {
	Result  QualityIQITimeseriesGroupsResponse             `json:"result,required"`
	Success bool                                           `json:"success,required"`
	JSON    qualityIQITimeseriesGroupsResponseEnvelopeJSON `json:"-"`
}

// qualityIQITimeseriesGroupsResponseEnvelopeJSON contains the JSON metadata for
// the struct [QualityIQITimeseriesGroupsResponseEnvelope]
type qualityIQITimeseriesGroupsResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *QualityIQITimeseriesGroupsResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r qualityIQITimeseriesGroupsResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
