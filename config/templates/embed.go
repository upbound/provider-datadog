/*
Copyright 2026 Upbound Inc.
*/

// Package templates holds the code generation templates that override the
// upjet defaults.
package templates

import _ "embed"

// ControllerTemplate is the template for the managed resource controller
// setup files. It extends the upjet default with the reconciliation policy
// reconciler, finalizer and rate limiter so that the ProviderConfig's
// reconciliationPolicy is honoured.
//
//go:embed controller.go.tmpl
var ControllerTemplate string
