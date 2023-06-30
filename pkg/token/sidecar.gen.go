// Copyright 2023 Yahoo Japan Corporation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package token provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.13.0 DO NOT EDIT.
package token

// AccessTokenRequestBody defines model for AccessTokenRequestBody.
type AccessTokenRequestBody struct {
	// Domain Access token domain name
	Domain string `json:"domain"`

	// Expiry Access token expiry time (in second)
	Expiry *int `json:"expiry,omitempty"`

	// ProxyForPrincipal Access token proxyForPrincipal name
	ProxyForPrincipal *string `json:"proxy_for_principal,omitempty"`

	// Role Access token role name (comma separated list)
	Role *string `json:"role,omitempty"`
}

// AccessTokenResponse defines model for AccessTokenResponse.
type AccessTokenResponse struct {
	// AccessToken Access token string
	AccessToken string `json:"access_token"`

	// ExpiresIn Access token expiry time (in second)
	ExpiresIn int `json:"expires_in"`

	// Scope Access token scope (Only added if role is not specified, space separated)
	Scope *string `json:"scope,omitempty"`

	// TokenType Access token token type
	TokenType string `json:"token_type"`
}

// RoleTokenRequestBody defines model for RoleTokenRequestBody.
type RoleTokenRequestBody struct {
	// Domain Role token domain name
	Domain string `json:"domain"`

	// MaxExpiry Role token maximum expiry time (in second)
	MaxExpiry *int `json:"max_expiry,omitempty"`

	// MinExpiry Role token minimum expiry time (in second)
	MinExpiry *int `json:"min_expiry,omitempty"`

	// ProxyForPrincipal Role token proxyForPrincipal name
	ProxyForPrincipal *string `json:"proxy_for_principal,omitempty"`

	// Role Role token role name (comma separated list)
	Role *string `json:"role,omitempty"`
}

// RoleTokenResponse defines model for RoleTokenResponse.
type RoleTokenResponse struct {
	// ExpiryTime Role token expiry time (Unix timestamp in second)
	ExpiryTime int64 `json:"expiryTime"`

	// Token Role token string
	Token string `json:"token"`
}

// AtResponse defines model for atResponse.
type AtResponse = AccessTokenResponse

// RtResponse defines model for rtResponse.
type RtResponse = RoleTokenResponse

// AtRequestBody defines model for atRequestBody.
type AtRequestBody = AccessTokenRequestBody

// RtRequestBody defines model for rtRequestBody.
type RtRequestBody = RoleTokenRequestBody

// FetchAccessTokenJSONRequestBody defines body for FetchAccessToken for application/json ContentType.
type FetchAccessTokenJSONRequestBody = AccessTokenRequestBody

// FetchRoleTokenJSONRequestBody defines body for FetchRoleToken for application/json ContentType.
type FetchRoleTokenJSONRequestBody = RoleTokenRequestBody