// Copyright 2022 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package leaseexpiry_test

import (
	"github.com/golang/mock/gomock"
	"github.com/juju/clock"
	"github.com/juju/errors"
	jc "github.com/juju/testing/checkers"
	"github.com/juju/worker/v3"
	dt "github.com/juju/worker/v3/dependency/testing"
	gc "gopkg.in/check.v1"

	coredatabase "github.com/juju/juju/core/database"
	"github.com/juju/juju/core/lease"
	"github.com/juju/juju/database/testing"
	jujutesting "github.com/juju/juju/testing"
	"github.com/juju/juju/worker/leaseexpiry"
)

type manifoldSuite struct {
	testing.ControllerSuite

	store *MockExpiryStore
}

var _ = gc.Suite(&manifoldSuite{})

func (s *manifoldSuite) TestInputs(c *gc.C) {
	cfg := s.newManifoldConfig(c)

	c.Check(leaseexpiry.Manifold(cfg).Inputs, jc.DeepEquals, []string{"clock-name", "db-accessor-name"})
}

func (s *manifoldSuite) TestConfigValidate(c *gc.C) {
	validCfg := s.newManifoldConfig(c)

	cfg := validCfg
	cfg.ClockName = ""
	c.Check(errors.Is(cfg.Validate(), errors.NotValid), jc.IsTrue)

	cfg = validCfg
	cfg.DBAccessorName = ""
	c.Check(errors.Is(cfg.Validate(), errors.NotValid), jc.IsTrue)

	cfg = validCfg
	cfg.Logger = nil
	c.Check(errors.Is(cfg.Validate(), errors.NotValid), jc.IsTrue)

	cfg = validCfg
	cfg.NewWorker = nil
	c.Check(errors.Is(cfg.Validate(), errors.NotValid), jc.IsTrue)

	cfg = validCfg
	cfg.NewStore = nil
	c.Check(errors.Is(cfg.Validate(), errors.NotValid), jc.IsTrue)
}

func (s *manifoldSuite) TestStartSuccess(c *gc.C) {
	defer s.setupMocks(c).Finish()

	cfg := s.newManifoldConfig(c)

	work, err := leaseexpiry.Manifold(cfg).Start(s.newStubContext())
	c.Check(work, gc.NotNil)
	c.Check(err, jc.ErrorIsNil)
}

func (s *manifoldSuite) setupMocks(c *gc.C) *gomock.Controller {
	ctrl := gomock.NewController(c)

	s.store = NewMockExpiryStore(ctrl)

	return ctrl
}

func (s *manifoldSuite) newStubContext() *dt.Context {
	return dt.StubContext(nil, map[string]interface{}{
		"clock-name":       clock.WallClock,
		"db-accessor-name": stubDBGetter{runner: s.TxnRunner()},
	})
}

// newManifoldConfig creates and returns a new ManifoldConfig instance based on
// the supplied arguments.
func (s *manifoldSuite) newManifoldConfig(c *gc.C) leaseexpiry.ManifoldConfig {
	return leaseexpiry.ManifoldConfig{
		ClockName:      "clock-name",
		DBAccessorName: "db-accessor-name",
		Logger:         jujutesting.CheckLogger{Log: c},
		NewWorker:      func(config leaseexpiry.Config) (worker.Worker, error) { return nil, nil },
		NewStore: func(db coredatabase.DBGetter, logger leaseexpiry.Logger) lease.ExpiryStore {
			return s.store
		},
	}
}

type stubDBGetter struct {
	runner coredatabase.TxnRunner
}

func (s stubDBGetter) GetDB(name string) (coredatabase.TxnRunner, error) {
	if name != "controller" {
		return nil, errors.Errorf(`expected a request for "controller" DB; got %q`, name)
	}
	return s.runner, nil
}
