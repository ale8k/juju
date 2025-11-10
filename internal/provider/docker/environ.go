// Copyright 2025 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

// Package docker implements a minimal CAAS provider backed by Docker.
// This provider exposes a caas.Broker and does not implement an IaaS-style
// environs.Environ. See provider.go and broker.go.
package docker
