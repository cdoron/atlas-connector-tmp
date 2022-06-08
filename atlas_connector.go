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
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"

	api "github.com/fybrik/datacatalog-go/api"
)

// DefaultApiController binds http requests to an api service and writes the service results to the http response
type DefaultApiController struct {
	service      AtlasApiServicer
	errorHandler api.ErrorHandler
}

// DefaultApiOption for how the controller is set up.
type DefaultApiOption func(*DefaultApiController)

// WithDefaultApiErrorHandler inject ErrorHandler into controller
func WithDefaultApiErrorHandler(h api.ErrorHandler) DefaultApiOption {
	return func(c *DefaultApiController) {
		c.errorHandler = h
	}
}

// NewDefaultApiController creates a default api controller
func NewApacheApiController(s AtlasApiServicer, opts ...DefaultApiOption) api.Router {
	controller := &DefaultApiController{
		service:      s,
		errorHandler: api.DefaultErrorHandler,
	}

	for _, opt := range opts {
		opt(controller)
	}

	return controller
}

// Routes returns all the api routes for the DefaultApiController
func (c *DefaultApiController) Routes() api.Routes {
	return api.Routes{
		{
			"CreateAsset",
			strings.ToUpper("Post"),
			"/createAsset",
			c.CreateAsset,
		},
		{
			"DeleteAsset",
			strings.ToUpper("Delete"),
			"/deleteAsset",
			c.DeleteAsset,
		},
		{
			"GetAssetInfo",
			strings.ToUpper("Post"),
			"/getAssetInfo",
			c.GetAssetInfo,
		},
		{
			"UpdateAsset",
			strings.ToUpper("Patch"),
			"/updateAsset",
			c.UpdateAsset,
		},
	}
}

// CreateAsset - This REST API writes data asset information to the data catalog configured in fybrik
func (c *DefaultApiController) CreateAsset(w http.ResponseWriter, r *http.Request) {
	xRequestDatacatalogWriteCredParam := r.Header.Get("X-Request-Datacatalog-Write-Cred")
	createAssetRequestParam := api.CreateAssetRequest{}

	// We need to open the request body twice. Once to extract the body
	// content as is, and once to contstruct the createAssetRequestParam.
	// We need the body content as is, as it contains fields which are not
	// defined in the spec (for example the s3 connection information) that
	// gets lost in createAssetRequestParam.
	bodyBytes, _ := ioutil.ReadAll(r.Body)
	r.Body.Close() //  must close
	r.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))

	d := json.NewDecoder(r.Body)

	// removed following line to allow unrecognized fields
	// d.DisallowUnknownFields()

	if err := d.Decode(&createAssetRequestParam); err != nil {
		c.errorHandler(w, r, &api.ParsingError{Err: err}, nil)
		return
	}
	if err := api.AssertCreateAssetRequestRequired(createAssetRequestParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.CreateAsset(r.Context(), xRequestDatacatalogWriteCredParam, createAssetRequestParam, bodyBytes)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	api.EncodeJSONResponse(result.Body, &result.Code, w)

}

// DeleteAsset - This REST API deletes data asset
func (c *DefaultApiController) DeleteAsset(w http.ResponseWriter, r *http.Request) {
	xRequestDatacatalogCredParam := r.Header.Get("X-Request-Datacatalog-Cred")
	deleteAssetRequestParam := api.DeleteAssetRequest{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&deleteAssetRequestParam); err != nil {
		c.errorHandler(w, r, &api.ParsingError{Err: err}, nil)
		return
	}
	if err := api.AssertDeleteAssetRequestRequired(deleteAssetRequestParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.DeleteAsset(r.Context(), xRequestDatacatalogCredParam, deleteAssetRequestParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	api.EncodeJSONResponse(result.Body, &result.Code, w)

}

// GetAssetInfo - This REST API gets data asset information from the data catalog configured in fybrik for the data sets indicated in FybrikApplication yaml
func (c *DefaultApiController) GetAssetInfo(w http.ResponseWriter, r *http.Request) {
	xRequestDatacatalogCredParam := r.Header.Get("X-Request-Datacatalog-Cred")
	getAssetRequestParam := api.GetAssetRequest{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&getAssetRequestParam); err != nil {
		c.errorHandler(w, r, &api.ParsingError{Err: err}, nil)
		return
	}
	if err := api.AssertGetAssetRequestRequired(getAssetRequestParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.GetAssetInfo(r.Context(), xRequestDatacatalogCredParam, getAssetRequestParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	api.EncodeJSONResponse(result.Body, &result.Code, w)

}

// UpdateAsset - This REST API updates data asset information in the data catalog configured in fybrik
func (c *DefaultApiController) UpdateAsset(w http.ResponseWriter, r *http.Request) {
	xRequestDatacatalogUpdateCredParam := r.Header.Get("X-Request-Datacatalog-Update-Cred")
	updateAssetRequestParam := api.UpdateAssetRequest{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&updateAssetRequestParam); err != nil {
		c.errorHandler(w, r, &api.ParsingError{Err: err}, nil)
		return
	}
	if err := api.AssertUpdateAssetRequestRequired(updateAssetRequestParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.UpdateAsset(r.Context(), xRequestDatacatalogUpdateCredParam, updateAssetRequestParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	api.EncodeJSONResponse(result.Body, &result.Code, w)

}
