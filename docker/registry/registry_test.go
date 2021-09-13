// Copyright 2021 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package registry_test

import (
	"github.com/golang/mock/gomock"
	"github.com/juju/testing"
	jc "github.com/juju/testing/checkers"
	gc "gopkg.in/check.v1"

	"github.com/juju/juju/docker"
	"github.com/juju/juju/docker/registry"
	"github.com/juju/juju/docker/registry/mocks"
)

type registrySuite struct {
	testing.IsolationSuite
}

var _ = gc.Suite(&registrySuite{})

func (s *registrySuite) TestInitClient(c *gc.C) {
	ctrl := gomock.NewController(c)
	defer ctrl.Finish()
	initializer := mocks.NewMockInitializer(ctrl)

	gomock.InOrder(
		initializer.EXPECT().DecideBaseURL().Return(nil),
		initializer.EXPECT().WrapTransport().Return(nil),
		initializer.EXPECT().Ping().Return(nil),
	)
	err := registry.InitClient(initializer)
	c.Assert(err, jc.ErrorIsNil)
}

func (s *registrySuite) TestNewRegistryNotSupported(c *gc.C) {

	imageRepoDetails := docker.ImageRepoDetails{
		Repository:    "gcr.io/jujuqa-project",
		ServerAddress: "gcr.io",
	}
	_, err := registry.New(imageRepoDetails)
	c.Assert(err, gc.ErrorMatches, `google container registry not supported`)

	imageRepoDetails = docker.ImageRepoDetails{
		Repository:    "quay.io/jujuqa-project",
		ServerAddress: "quay.io",
	}
	_, err = registry.New(imageRepoDetails)
	c.Assert(err, gc.ErrorMatches, `quay.io container registry not supported`)

	imageRepoDetails = docker.ImageRepoDetails{
		Repository:    "ecr.aws/jujuqa-project",
		ServerAddress: "ecr.aws",
	}
	_, err = registry.New(imageRepoDetails)
	c.Assert(err, gc.ErrorMatches, `AWS elastic container registry not supported`)
}