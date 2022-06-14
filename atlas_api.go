/*
 * Data Catalog Service - Asset Details
 *
 * API version: 1.0.0
 * Based on code Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package main

import (
	"context"

	api "github.com/fybrik/datacatalog-go/api"
)

// We had to change the CreateAsset interface to include the entire message body.
// The reason for the is that the body may include fields such as S3 parameters,
// which are not part of the DataCatalog spec. These fields are removed from the
// CreateAssetRequest
type AtlasApiServicer interface {
	CreateAsset(context.Context, string, api.CreateAssetRequest, []byte) (api.ImplResponse, error)
	DeleteAsset(context.Context, string, api.DeleteAssetRequest) (api.ImplResponse, error)
	GetAssetInfo(context.Context, string, api.GetAssetRequest) (api.ImplResponse, error)
	UpdateAsset(context.Context, string, api.UpdateAssetRequest) (api.ImplResponse, error)
}
