/*
 * Data Catalog Service - Asset Details
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package main

import (
	"context"

	api "github.com/cdoron/datacatalog-go/api"
)

// DefaultApiServicer defines the api actions for the DefaultApi service
// This interface intended to stay up to date with the openapi yaml used to generate it,
// while the service implementation can be ignored with the .openapi-generator-ignore file
// and updated with the logic required for the API.
type AtlasApiServicer interface {
	CreateAsset(context.Context, string, api.CreateAssetRequest, []byte) (api.ImplResponse, error)
	DeleteAsset(context.Context, string, api.DeleteAssetRequest) (api.ImplResponse, error)
	GetAssetInfo(context.Context, string, api.GetAssetRequest) (api.ImplResponse, error)
	UpdateAsset(context.Context, string, api.UpdateAssetRequest, []byte) (api.ImplResponse, error)
}
