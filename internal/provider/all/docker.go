// Copyright 2025 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

//go:build !minimal || provider_docker

package all

import (
	// Register the docker CAAS provider.
	_ "github.com/juju/juju/internal/provider/docker"
)
