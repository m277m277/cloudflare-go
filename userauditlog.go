// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"net/http"
	"net/url"
	"reflect"
	"time"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/apiquery"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/internal/shared"
	"github.com/cloudflare/cloudflare-sdk-go/option"
	"github.com/tidwall/gjson"
)

// UserAuditLogService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewUserAuditLogService] method
// instead.
type UserAuditLogService struct {
	Options []option.RequestOption
}

// NewUserAuditLogService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewUserAuditLogService(opts ...option.RequestOption) (r *UserAuditLogService) {
	r = &UserAuditLogService{}
	r.Options = opts
	return
}

// Gets a list of audit logs for a user account. Can be filtered by who made the
// change, on which zone, and the timeframe of the change.
func (r *UserAuditLogService) AuditLogsGetUserAuditLogs(ctx context.Context, query UserAuditLogAuditLogsGetUserAuditLogsParams, opts ...option.RequestOption) (res *UserAuditLogAuditLogsGetUserAuditLogsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "user/audit_logs"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Union satisfied by [UserAuditLogAuditLogsGetUserAuditLogsResponseObject] or
// [UserAuditLogAuditLogsGetUserAuditLogsResponse4rslHd6EAPIResponseCommon].
type UserAuditLogAuditLogsGetUserAuditLogsResponse interface {
	implementsUserAuditLogAuditLogsGetUserAuditLogsResponse()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*UserAuditLogAuditLogsGetUserAuditLogsResponse)(nil)).Elem(), "")
}

type UserAuditLogAuditLogsGetUserAuditLogsResponseObject struct {
	Errors   interface{}                                                 `json:"errors,nullable"`
	Messages []interface{}                                               `json:"messages"`
	Result   []UserAuditLogAuditLogsGetUserAuditLogsResponseObjectResult `json:"result"`
	Success  bool                                                        `json:"success"`
	JSON     userAuditLogAuditLogsGetUserAuditLogsResponseObjectJSON     `json:"-"`
}

// userAuditLogAuditLogsGetUserAuditLogsResponseObjectJSON contains the JSON
// metadata for the struct [UserAuditLogAuditLogsGetUserAuditLogsResponseObject]
type userAuditLogAuditLogsGetUserAuditLogsResponseObjectJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserAuditLogAuditLogsGetUserAuditLogsResponseObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r UserAuditLogAuditLogsGetUserAuditLogsResponseObject) implementsUserAuditLogAuditLogsGetUserAuditLogsResponse() {
}

type UserAuditLogAuditLogsGetUserAuditLogsResponseObjectResult struct {
	// A string that uniquely identifies the audit log.
	ID     string                                                          `json:"id"`
	Action UserAuditLogAuditLogsGetUserAuditLogsResponseObjectResultAction `json:"action"`
	Actor  UserAuditLogAuditLogsGetUserAuditLogsResponseObjectResultActor  `json:"actor"`
	// The source of the event.
	Interface string `json:"interface"`
	// An object which can lend more context to the action being logged. This is a
	// flexible value and varies between different actions.
	Metadata interface{} `json:"metadata"`
	// The new value of the resource that was modified.
	NewValue string `json:"newValue"`
	// The value of the resource before it was modified.
	OldValue string                                                            `json:"oldValue"`
	Owner    UserAuditLogAuditLogsGetUserAuditLogsResponseObjectResultOwner    `json:"owner"`
	Resource UserAuditLogAuditLogsGetUserAuditLogsResponseObjectResultResource `json:"resource"`
	// A UTC RFC3339 timestamp that specifies when the action being logged occured.
	When time.Time                                                     `json:"when" format:"date-time"`
	JSON userAuditLogAuditLogsGetUserAuditLogsResponseObjectResultJSON `json:"-"`
}

// userAuditLogAuditLogsGetUserAuditLogsResponseObjectResultJSON contains the JSON
// metadata for the struct
// [UserAuditLogAuditLogsGetUserAuditLogsResponseObjectResult]
type userAuditLogAuditLogsGetUserAuditLogsResponseObjectResultJSON struct {
	ID          apijson.Field
	Action      apijson.Field
	Actor       apijson.Field
	Interface   apijson.Field
	Metadata    apijson.Field
	NewValue    apijson.Field
	OldValue    apijson.Field
	Owner       apijson.Field
	Resource    apijson.Field
	When        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserAuditLogAuditLogsGetUserAuditLogsResponseObjectResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type UserAuditLogAuditLogsGetUserAuditLogsResponseObjectResultAction struct {
	// A boolean that indicates if the action attempted was successful.
	Result bool `json:"result"`
	// A short string that describes the action that was performed.
	Type string                                                              `json:"type"`
	JSON userAuditLogAuditLogsGetUserAuditLogsResponseObjectResultActionJSON `json:"-"`
}

// userAuditLogAuditLogsGetUserAuditLogsResponseObjectResultActionJSON contains the
// JSON metadata for the struct
// [UserAuditLogAuditLogsGetUserAuditLogsResponseObjectResultAction]
type userAuditLogAuditLogsGetUserAuditLogsResponseObjectResultActionJSON struct {
	Result      apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserAuditLogAuditLogsGetUserAuditLogsResponseObjectResultAction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type UserAuditLogAuditLogsGetUserAuditLogsResponseObjectResultActor struct {
	// The ID of the actor that performed the action. If a user performed the action,
	// this will be their User ID.
	ID string `json:"id"`
	// The email of the user that performed the action.
	Email string `json:"email" format:"email"`
	// The IP address of the request that performed the action.
	IP string `json:"ip"`
	// The type of actor, whether a User, Cloudflare Admin, or an Automated System.
	Type UserAuditLogAuditLogsGetUserAuditLogsResponseObjectResultActorType `json:"type"`
	JSON userAuditLogAuditLogsGetUserAuditLogsResponseObjectResultActorJSON `json:"-"`
}

// userAuditLogAuditLogsGetUserAuditLogsResponseObjectResultActorJSON contains the
// JSON metadata for the struct
// [UserAuditLogAuditLogsGetUserAuditLogsResponseObjectResultActor]
type userAuditLogAuditLogsGetUserAuditLogsResponseObjectResultActorJSON struct {
	ID          apijson.Field
	Email       apijson.Field
	IP          apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserAuditLogAuditLogsGetUserAuditLogsResponseObjectResultActor) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of actor, whether a User, Cloudflare Admin, or an Automated System.
type UserAuditLogAuditLogsGetUserAuditLogsResponseObjectResultActorType string

const (
	UserAuditLogAuditLogsGetUserAuditLogsResponseObjectResultActorTypeUser       UserAuditLogAuditLogsGetUserAuditLogsResponseObjectResultActorType = "user"
	UserAuditLogAuditLogsGetUserAuditLogsResponseObjectResultActorTypeAdmin      UserAuditLogAuditLogsGetUserAuditLogsResponseObjectResultActorType = "admin"
	UserAuditLogAuditLogsGetUserAuditLogsResponseObjectResultActorTypeCloudflare UserAuditLogAuditLogsGetUserAuditLogsResponseObjectResultActorType = "Cloudflare"
)

type UserAuditLogAuditLogsGetUserAuditLogsResponseObjectResultOwner struct {
	// Identifier
	ID   string                                                             `json:"id"`
	JSON userAuditLogAuditLogsGetUserAuditLogsResponseObjectResultOwnerJSON `json:"-"`
}

// userAuditLogAuditLogsGetUserAuditLogsResponseObjectResultOwnerJSON contains the
// JSON metadata for the struct
// [UserAuditLogAuditLogsGetUserAuditLogsResponseObjectResultOwner]
type userAuditLogAuditLogsGetUserAuditLogsResponseObjectResultOwnerJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserAuditLogAuditLogsGetUserAuditLogsResponseObjectResultOwner) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type UserAuditLogAuditLogsGetUserAuditLogsResponseObjectResultResource struct {
	// An identifier for the resource that was affected by the action.
	ID string `json:"id"`
	// A short string that describes the resource that was affected by the action.
	Type string                                                                `json:"type"`
	JSON userAuditLogAuditLogsGetUserAuditLogsResponseObjectResultResourceJSON `json:"-"`
}

// userAuditLogAuditLogsGetUserAuditLogsResponseObjectResultResourceJSON contains
// the JSON metadata for the struct
// [UserAuditLogAuditLogsGetUserAuditLogsResponseObjectResultResource]
type userAuditLogAuditLogsGetUserAuditLogsResponseObjectResultResourceJSON struct {
	ID          apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserAuditLogAuditLogsGetUserAuditLogsResponseObjectResultResource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type UserAuditLogAuditLogsGetUserAuditLogsResponse4rslHd6EAPIResponseCommon struct {
	Errors   []UserAuditLogAuditLogsGetUserAuditLogsResponse4rslHd6EAPIResponseCommonError   `json:"errors,required"`
	Messages []UserAuditLogAuditLogsGetUserAuditLogsResponse4rslHd6EAPIResponseCommonMessage `json:"messages,required"`
	Result   UserAuditLogAuditLogsGetUserAuditLogsResponse4rslHd6EAPIResponseCommonResult    `json:"result,required"`
	// Whether the API call was successful
	Success UserAuditLogAuditLogsGetUserAuditLogsResponse4rslHd6EAPIResponseCommonSuccess `json:"success,required"`
	JSON    userAuditLogAuditLogsGetUserAuditLogsResponse4rslHd6EapiResponseCommonJSON    `json:"-"`
}

// userAuditLogAuditLogsGetUserAuditLogsResponse4rslHd6EapiResponseCommonJSON
// contains the JSON metadata for the struct
// [UserAuditLogAuditLogsGetUserAuditLogsResponse4rslHd6EAPIResponseCommon]
type userAuditLogAuditLogsGetUserAuditLogsResponse4rslHd6EapiResponseCommonJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserAuditLogAuditLogsGetUserAuditLogsResponse4rslHd6EAPIResponseCommon) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r UserAuditLogAuditLogsGetUserAuditLogsResponse4rslHd6EAPIResponseCommon) implementsUserAuditLogAuditLogsGetUserAuditLogsResponse() {
}

type UserAuditLogAuditLogsGetUserAuditLogsResponse4rslHd6EAPIResponseCommonError struct {
	Code    int64                                                                           `json:"code,required"`
	Message string                                                                          `json:"message,required"`
	JSON    userAuditLogAuditLogsGetUserAuditLogsResponse4rslHd6EapiResponseCommonErrorJSON `json:"-"`
}

// userAuditLogAuditLogsGetUserAuditLogsResponse4rslHd6EapiResponseCommonErrorJSON
// contains the JSON metadata for the struct
// [UserAuditLogAuditLogsGetUserAuditLogsResponse4rslHd6EAPIResponseCommonError]
type userAuditLogAuditLogsGetUserAuditLogsResponse4rslHd6EapiResponseCommonErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserAuditLogAuditLogsGetUserAuditLogsResponse4rslHd6EAPIResponseCommonError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type UserAuditLogAuditLogsGetUserAuditLogsResponse4rslHd6EAPIResponseCommonMessage struct {
	Code    int64                                                                             `json:"code,required"`
	Message string                                                                            `json:"message,required"`
	JSON    userAuditLogAuditLogsGetUserAuditLogsResponse4rslHd6EapiResponseCommonMessageJSON `json:"-"`
}

// userAuditLogAuditLogsGetUserAuditLogsResponse4rslHd6EapiResponseCommonMessageJSON
// contains the JSON metadata for the struct
// [UserAuditLogAuditLogsGetUserAuditLogsResponse4rslHd6EAPIResponseCommonMessage]
type userAuditLogAuditLogsGetUserAuditLogsResponse4rslHd6EapiResponseCommonMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserAuditLogAuditLogsGetUserAuditLogsResponse4rslHd6EAPIResponseCommonMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Union satisfied by
// [UserAuditLogAuditLogsGetUserAuditLogsResponse4rslHd6EAPIResponseCommonResultUnknown],
// [UserAuditLogAuditLogsGetUserAuditLogsResponse4rslHd6EAPIResponseCommonResultArray]
// or [shared.UnionString].
type UserAuditLogAuditLogsGetUserAuditLogsResponse4rslHd6EAPIResponseCommonResult interface {
	ImplementsUserAuditLogAuditLogsGetUserAuditLogsResponse4rslHd6EapiResponseCommonResult()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*UserAuditLogAuditLogsGetUserAuditLogsResponse4rslHd6EAPIResponseCommonResult)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type UserAuditLogAuditLogsGetUserAuditLogsResponse4rslHd6EAPIResponseCommonResultArray []interface{}

func (r UserAuditLogAuditLogsGetUserAuditLogsResponse4rslHd6EAPIResponseCommonResultArray) ImplementsUserAuditLogAuditLogsGetUserAuditLogsResponse4rslHd6EapiResponseCommonResult() {
}

// Whether the API call was successful
type UserAuditLogAuditLogsGetUserAuditLogsResponse4rslHd6EAPIResponseCommonSuccess bool

const (
	UserAuditLogAuditLogsGetUserAuditLogsResponse4rslHd6EAPIResponseCommonSuccessTrue UserAuditLogAuditLogsGetUserAuditLogsResponse4rslHd6EAPIResponseCommonSuccess = true
)

type UserAuditLogAuditLogsGetUserAuditLogsParams struct {
	// Finds a specific log by its ID.
	ID     param.Field[string]                                            `query:"id"`
	Action param.Field[UserAuditLogAuditLogsGetUserAuditLogsParamsAction] `query:"action"`
	Actor  param.Field[UserAuditLogAuditLogsGetUserAuditLogsParamsActor]  `query:"actor"`
	// Limits the returned results to logs older than the specified date. This can be a
	// date string `2019-04-30` or an absolute timestamp that conforms to RFC3339.
	Before param.Field[time.Time] `query:"before" format:"date-time"`
	// Changes the direction of the chronological sorting.
	Direction param.Field[UserAuditLogAuditLogsGetUserAuditLogsParamsDirection] `query:"direction"`
	// Indicates that this request is an export of logs in CSV format.
	Export param.Field[bool] `query:"export"`
	// Indicates whether or not to hide user level audit logs.
	HideUserLogs param.Field[bool] `query:"hide_user_logs"`
	// Defines which page of results to return.
	Page param.Field[float64] `query:"page"`
	// Sets the number of results to return per page.
	PerPage param.Field[float64] `query:"per_page"`
	// Limits the returned results to logs newer than the specified date. This can be a
	// date string `2019-04-30` or an absolute timestamp that conforms to RFC3339.
	Since param.Field[time.Time]                                       `query:"since" format:"date-time"`
	Zone  param.Field[UserAuditLogAuditLogsGetUserAuditLogsParamsZone] `query:"zone"`
}

// URLQuery serializes [UserAuditLogAuditLogsGetUserAuditLogsParams]'s query
// parameters as `url.Values`.
func (r UserAuditLogAuditLogsGetUserAuditLogsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type UserAuditLogAuditLogsGetUserAuditLogsParamsAction struct {
	// Filters by the action type.
	Type param.Field[string] `query:"type"`
}

// URLQuery serializes [UserAuditLogAuditLogsGetUserAuditLogsParamsAction]'s query
// parameters as `url.Values`.
func (r UserAuditLogAuditLogsGetUserAuditLogsParamsAction) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type UserAuditLogAuditLogsGetUserAuditLogsParamsActor struct {
	// Filters by the email address of the actor that made the change.
	Email param.Field[string] `query:"email" format:"email"`
	// Filters by the IP address of the request that made the change by specific IP
	// address or valid CIDR Range.
	IP param.Field[string] `query:"ip"`
}

// URLQuery serializes [UserAuditLogAuditLogsGetUserAuditLogsParamsActor]'s query
// parameters as `url.Values`.
func (r UserAuditLogAuditLogsGetUserAuditLogsParamsActor) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Changes the direction of the chronological sorting.
type UserAuditLogAuditLogsGetUserAuditLogsParamsDirection string

const (
	UserAuditLogAuditLogsGetUserAuditLogsParamsDirectionDesc UserAuditLogAuditLogsGetUserAuditLogsParamsDirection = "desc"
	UserAuditLogAuditLogsGetUserAuditLogsParamsDirectionAsc  UserAuditLogAuditLogsGetUserAuditLogsParamsDirection = "asc"
)

type UserAuditLogAuditLogsGetUserAuditLogsParamsZone struct {
	// Filters by the name of the zone associated to the change.
	Name param.Field[string] `query:"name"`
}

// URLQuery serializes [UserAuditLogAuditLogsGetUserAuditLogsParamsZone]'s query
// parameters as `url.Values`.
func (r UserAuditLogAuditLogsGetUserAuditLogsParamsZone) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
