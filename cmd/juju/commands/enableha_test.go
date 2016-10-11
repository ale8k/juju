// Copyright 2015 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package commands

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/juju/cmd"
	jc "github.com/juju/testing/checkers"
	gc "gopkg.in/check.v1"
	goyaml "gopkg.in/yaml.v2"

	"github.com/juju/juju/apiserver/common"
	"github.com/juju/juju/apiserver/params"
	"github.com/juju/juju/cmd/modelcmd"
	"github.com/juju/juju/constraints"
	"github.com/juju/juju/instance"
	"github.com/juju/juju/juju/testing"
	"github.com/juju/juju/state"
	coretesting "github.com/juju/juju/testing"
	"github.com/juju/juju/testing/factory"
)

type EnableHASuite struct {
	// TODO (cherylj) change this back to a FakeJujuXDGDataHomeSuite to
	// remove the mongo dependency once ensure-availability is
	// moved under a supercommand again.
	testing.JujuConnSuite
	fake *fakeHAClient
}

// invalidNumServers is a number of controllers that would
// never be generated by the enable-ha command.
const invalidNumServers = -2

func (s *EnableHASuite) SetUpTest(c *gc.C) {
	s.JujuConnSuite.SetUpTest(c)

	// Initialize numControllers to an invalid number to validate
	// that enable-ha doesn't call into the API when its
	// pre-checks fail
	s.fake = &fakeHAClient{numControllers: invalidNumServers}
}

type fakeHAClient struct {
	numControllers int
	cons           constraints.Value
	err            error
	placement      []string
	result         params.ControllersChanges
}

func (f *fakeHAClient) Close() error {
	return nil
}

func (f *fakeHAClient) EnableHA(numControllers int, cons constraints.Value, placement []string) (params.ControllersChanges, error) {

	f.numControllers = numControllers
	f.cons = cons
	f.placement = placement

	if f.err != nil {
		return f.result, f.err
	}

	if numControllers == 1 {
		return f.result, nil
	}

	// In the real HAClient, specifying a numControllers value of 0
	// indicates that the default value (3) should be used
	if numControllers == 0 {
		numControllers = 3
	}

	f.result.Maintained = append(f.result.Maintained, "machine-0")

	for _, p := range placement {
		m, err := instance.ParsePlacement(p)
		if err == nil && m.Scope == instance.MachineScope {
			f.result.Converted = append(f.result.Converted, "machine-"+m.Directive)
		}
	}

	// We may need to pretend that we added some machines.
	for i := len(f.result.Converted) + 1; i < numControllers; i++ {
		f.result.Added = append(f.result.Added, fmt.Sprintf("machine-%d", i))
	}

	return f.result, nil
}

var _ = gc.Suite(&EnableHASuite{})

func (s *EnableHASuite) runEnableHA(c *gc.C, args ...string) (*cmd.Context, error) {
	command := &enableHACommand{newHAClientFunc: func() (MakeHAClient, error) { return s.fake, nil }}
	return coretesting.RunCommand(c, modelcmd.WrapController(command), args...)
}

func (s *EnableHASuite) TestEnableHA(c *gc.C) {
	ctx, err := s.runEnableHA(c, "-n", "1")
	c.Assert(err, jc.ErrorIsNil)
	c.Assert(coretesting.Stdout(ctx), gc.Equals, "\n")

	c.Assert(s.fake.numControllers, gc.Equals, 1)
	c.Assert(&s.fake.cons, jc.Satisfies, constraints.IsEmpty)
	c.Assert(len(s.fake.placement), gc.Equals, 0)
}

func (s *EnableHASuite) TestBlockEnableHA(c *gc.C) {
	s.fake.err = common.OperationBlockedError("TestBlockEnableHA")
	_, err := s.runEnableHA(c, "-n", "1")
	coretesting.AssertOperationWasBlocked(c, err, ".*TestBlockEnableHA.*")
}

func (s *EnableHASuite) TestEnableHAFormatYaml(c *gc.C) {
	expected := map[string][]string{
		"maintained": {"0"},
		"added":      {"1", "2"},
	}

	ctx, err := s.runEnableHA(c, "-n", "3", "--format", "yaml")
	c.Assert(err, jc.ErrorIsNil)

	c.Assert(s.fake.numControllers, gc.Equals, 3)
	c.Assert(&s.fake.cons, jc.Satisfies, constraints.IsEmpty)
	c.Assert(len(s.fake.placement), gc.Equals, 0)

	var result map[string][]string
	err = goyaml.Unmarshal(ctx.Stdout.(*bytes.Buffer).Bytes(), &result)
	c.Assert(err, jc.ErrorIsNil)
	c.Assert(result, gc.DeepEquals, expected)
}

func (s *EnableHASuite) TestEnableHAFormatJson(c *gc.C) {
	expected := map[string][]string{
		"maintained": {"0"},
		"added":      {"1", "2"},
	}

	ctx, err := s.runEnableHA(c, "-n", "3", "--format", "json")
	c.Assert(err, jc.ErrorIsNil)

	c.Assert(s.fake.numControllers, gc.Equals, 3)
	c.Assert(&s.fake.cons, jc.Satisfies, constraints.IsEmpty)
	c.Assert(len(s.fake.placement), gc.Equals, 0)

	var result map[string][]string
	err = json.Unmarshal(ctx.Stdout.(*bytes.Buffer).Bytes(), &result)
	c.Assert(err, jc.ErrorIsNil)
	c.Assert(result, gc.DeepEquals, expected)
}

func (s *EnableHASuite) TestEnableHAWithFive(c *gc.C) {
	// Also test with -n 5 to validate numbers other than 1 and 3
	ctx, err := s.runEnableHA(c, "-n", "5")
	c.Assert(err, jc.ErrorIsNil)
	c.Assert(coretesting.Stdout(ctx), gc.Equals,
		"maintaining machines: 0\n"+
			"adding machines: 1, 2, 3, 4\n\n")

	c.Assert(s.fake.numControllers, gc.Equals, 5)
	c.Assert(&s.fake.cons, jc.Satisfies, constraints.IsEmpty)
	c.Assert(len(s.fake.placement), gc.Equals, 0)
}

func (s *EnableHASuite) TestEnableHAWithConstraints(c *gc.C) {
	ctx, err := s.runEnableHA(c, "--constraints", "mem=4G", "-n", "3")
	c.Assert(err, jc.ErrorIsNil)
	c.Assert(coretesting.Stdout(ctx), gc.Equals,
		"maintaining machines: 0\n"+
			"adding machines: 1, 2\n\n")

	c.Assert(s.fake.numControllers, gc.Equals, 3)
	expectedCons := constraints.MustParse("mem=4G")
	c.Assert(s.fake.cons, gc.DeepEquals, expectedCons)
	c.Assert(len(s.fake.placement), gc.Equals, 0)
}

func (s *EnableHASuite) TestEnableHAWithPlacement(c *gc.C) {
	ctx, err := s.runEnableHA(c, "--to", "valid", "-n", "3")
	c.Assert(err, jc.ErrorIsNil)
	c.Assert(coretesting.Stdout(ctx), gc.Equals,
		"maintaining machines: 0\n"+
			"adding machines: 1, 2\n\n")

	c.Assert(s.fake.numControllers, gc.Equals, 3)
	c.Assert(&s.fake.cons, jc.Satisfies, constraints.IsEmpty)
	expectedPlacement := []string{"valid"}
	c.Assert(s.fake.placement, gc.DeepEquals, expectedPlacement)
}

func (s *EnableHASuite) TestEnableHAErrors(c *gc.C) {
	for _, n := range []int{-1, 2} {
		_, err := s.runEnableHA(c, "-n", fmt.Sprint(n))
		c.Assert(err, gc.ErrorMatches, "must specify a number of controllers odd and non-negative")
	}

	// Verify that enable-ha didn't call into the API
	c.Assert(s.fake.numControllers, gc.Equals, invalidNumServers)
}

func (s *EnableHASuite) TestEnableHAAllows0(c *gc.C) {
	// If the number of controllers is specified as "0", the API will
	// then use the default number of 3.
	ctx, err := s.runEnableHA(c, "-n", "0")
	c.Assert(err, jc.ErrorIsNil)
	c.Assert(coretesting.Stdout(ctx), gc.Equals,
		"maintaining machines: 0\n"+
			"adding machines: 1, 2\n\n")

	c.Assert(s.fake.numControllers, gc.Equals, 0)
	c.Assert(&s.fake.cons, jc.Satisfies, constraints.IsEmpty)
	c.Assert(len(s.fake.placement), gc.Equals, 0)
}

func (s *EnableHASuite) TestEnableHADefaultsTo0(c *gc.C) {
	// If the number of controllers is not specified, we pass in 0 to the
	// API.  The API will then use the default number of 3.
	ctx, err := s.runEnableHA(c)
	c.Assert(err, jc.ErrorIsNil)
	c.Assert(coretesting.Stdout(ctx), gc.Equals,
		"maintaining machines: 0\n"+
			"adding machines: 1, 2\n\n")

	c.Assert(s.fake.numControllers, gc.Equals, 0)
	c.Assert(&s.fake.cons, jc.Satisfies, constraints.IsEmpty)
	c.Assert(len(s.fake.placement), gc.Equals, 0)
}

func (s *EnableHASuite) TestEnableHAEndToEnd(c *gc.C) {
	s.Factory.MakeMachine(c, &factory.MachineParams{
		Jobs: []state.MachineJob{state.JobManageModel},
	})
	ctx, err := coretesting.RunCommand(c, newEnableHACommand(), "-n", "3")
	c.Assert(err, jc.ErrorIsNil)

	// Machine 0 is demoted because it hasn't reported its presence
	c.Assert(coretesting.Stdout(ctx), gc.Equals,
		"adding machines: 1, 2, 3\n"+
			"demoting machines: 0\n\n")
}

func (s *EnableHASuite) TestEnableHAToExisting(c *gc.C) {
	ctx, err := s.runEnableHA(c, "--to", "1,2")
	c.Assert(err, jc.ErrorIsNil)
	c.Check(coretesting.Stdout(ctx), gc.Equals, `
maintaining machines: 0
converting machines: 1, 2

`[1:])

	c.Check(s.fake.numControllers, gc.Equals, 0)
	c.Check(&s.fake.cons, jc.Satisfies, constraints.IsEmpty)
	c.Check(len(s.fake.placement), gc.Equals, 2)
}

func (s *EnableHASuite) TestEnableHADisallowsSeries(c *gc.C) {
	// We don't allow --series as an argument. This test ensures it is not
	// inadvertantly added back.
	ctx, err := s.runEnableHA(c, "-n", "0", "--series", "xenian")
	c.Assert(err, gc.ErrorMatches, "flag provided but not defined: --series")
	c.Assert(coretesting.Stdout(ctx), gc.Equals, "")
}
