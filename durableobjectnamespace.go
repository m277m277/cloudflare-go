// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-go/internal/apijson"
	"github.com/cloudflare/cloudflare-go/internal/param"
	"github.com/cloudflare/cloudflare-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/option"
)

// DurableObjectNamespaceService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewDurableObjectNamespaceService]
// method instead.
type DurableObjectNamespaceService struct {
	Options []option.RequestOption
	Objects *DurableObjectNamespaceObjectService
}

// NewDurableObjectNamespaceService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewDurableObjectNamespaceService(opts ...option.RequestOption) (r *DurableObjectNamespaceService) {
	r = &DurableObjectNamespaceService{}
	r.Options = opts
	r.Objects = NewDurableObjectNamespaceObjectService(opts...)
	return
}

// Returns the Durable Object namespaces owned by an account.
func (r *DurableObjectNamespaceService) List(ctx context.Context, query DurableObjectNamespaceListParams, opts ...option.RequestOption) (res *[]DurableObjectNamespaceListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env DurableObjectNamespaceListResponseEnvelope
	path := fmt.Sprintf("accounts/%s/workers/durable_objects/namespaces", query.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type DurableObjectNamespaceListResponse struct {
	ID     interface{}                            `json:"id"`
	Class  interface{}                            `json:"class"`
	Name   interface{}                            `json:"name"`
	Script interface{}                            `json:"script"`
	JSON   durableObjectNamespaceListResponseJSON `json:"-"`
}

// durableObjectNamespaceListResponseJSON contains the JSON metadata for the struct
// [DurableObjectNamespaceListResponse]
type durableObjectNamespaceListResponseJSON struct {
	ID          apijson.Field
	Class       apijson.Field
	Name        apijson.Field
	Script      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DurableObjectNamespaceListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r durableObjectNamespaceListResponseJSON) RawJSON() string {
	return r.raw
}

type DurableObjectNamespaceListParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type DurableObjectNamespaceListResponseEnvelope struct {
	Errors   []DurableObjectNamespaceListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DurableObjectNamespaceListResponseEnvelopeMessages `json:"messages,required"`
	Result   []DurableObjectNamespaceListResponse                 `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    DurableObjectNamespaceListResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo DurableObjectNamespaceListResponseEnvelopeResultInfo `json:"result_info"`
	JSON       durableObjectNamespaceListResponseEnvelopeJSON       `json:"-"`
}

// durableObjectNamespaceListResponseEnvelopeJSON contains the JSON metadata for
// the struct [DurableObjectNamespaceListResponseEnvelope]
type durableObjectNamespaceListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DurableObjectNamespaceListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r durableObjectNamespaceListResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type DurableObjectNamespaceListResponseEnvelopeErrors struct {
	Code    int64                                                `json:"code,required"`
	Message string                                               `json:"message,required"`
	JSON    durableObjectNamespaceListResponseEnvelopeErrorsJSON `json:"-"`
}

// durableObjectNamespaceListResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [DurableObjectNamespaceListResponseEnvelopeErrors]
type durableObjectNamespaceListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DurableObjectNamespaceListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r durableObjectNamespaceListResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type DurableObjectNamespaceListResponseEnvelopeMessages struct {
	Code    int64                                                  `json:"code,required"`
	Message string                                                 `json:"message,required"`
	JSON    durableObjectNamespaceListResponseEnvelopeMessagesJSON `json:"-"`
}

// durableObjectNamespaceListResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [DurableObjectNamespaceListResponseEnvelopeMessages]
type durableObjectNamespaceListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DurableObjectNamespaceListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r durableObjectNamespaceListResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type DurableObjectNamespaceListResponseEnvelopeSuccess bool

const (
	DurableObjectNamespaceListResponseEnvelopeSuccessTrue DurableObjectNamespaceListResponseEnvelopeSuccess = true
)

type DurableObjectNamespaceListResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                  `json:"total_count"`
	JSON       durableObjectNamespaceListResponseEnvelopeResultInfoJSON `json:"-"`
}

// durableObjectNamespaceListResponseEnvelopeResultInfoJSON contains the JSON
// metadata for the struct [DurableObjectNamespaceListResponseEnvelopeResultInfo]
type durableObjectNamespaceListResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DurableObjectNamespaceListResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r durableObjectNamespaceListResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}
