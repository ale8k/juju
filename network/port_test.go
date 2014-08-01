// Copyright 2013 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package network_test

import (
	jc "github.com/juju/testing/checkers"
	gc "launchpad.net/gocheck"

	"github.com/juju/juju/network"
	"github.com/juju/juju/testing"
)

type PortSuite struct {
	testing.BaseSuite
}

var _ = gc.Suite(&PortSuite{})

type hostPortTest struct {
	about         string
	hostPorts     []network.HostPort
	expectedIndex int
	preferIPv6    bool
}

// hostPortTest returns the HostPort equivalent test to the
// receiving selectTest.
func (t selectTest) hostPortTest() hostPortTest {
	hps := network.AddressesWithPort(t.addresses, 9999)
	for i := range hps {
		hps[i].Port = i + 1
	}
	return hostPortTest{
		about:         t.about,
		hostPorts:     hps,
		expectedIndex: t.expectedIndex,
		preferIPv6:    t.preferIPv6,
	}
}

// expected returns the expected host:port result
// of the test.
func (t hostPortTest) expected() string {
	if t.expectedIndex == -1 {
		return ""
	}
	return t.hostPorts[t.expectedIndex].NetAddr()
}

func (s *PortSuite) TestSelectPublicHostPort(c *gc.C) {
	oldValue := network.GetPreferIPv6()
	defer func() {
		network.SetPreferIPv6(oldValue)
	}()
	for i, t0 := range selectPublicTests {
		t := t0.hostPortTest()
		c.Logf("test %d: %s", i, t.about)
		network.SetPreferIPv6(t.preferIPv6)
		c.Check(network.SelectPublicHostPort(t.hostPorts), jc.DeepEquals, t.expected())
	}
}

func (s *PortSuite) TestSelectInternalHostPort(c *gc.C) {
	oldValue := network.GetPreferIPv6()
	defer func() {
		network.SetPreferIPv6(oldValue)
	}()
	for i, t0 := range selectInternalTests {
		t := t0.hostPortTest()
		c.Logf("test %d: %s", i, t.about)
		network.SetPreferIPv6(t.preferIPv6)
		c.Check(network.SelectInternalHostPort(t.hostPorts, false), jc.DeepEquals, t.expected())
	}
}

func (s *PortSuite) TestSelectInternalMachineHostPort(c *gc.C) {
	oldValue := network.GetPreferIPv6()
	defer func() {
		network.SetPreferIPv6(oldValue)
	}()
	for i, t0 := range selectInternalMachineTests {
		t := t0.hostPortTest()
		c.Logf("test %d: %s", i, t.about)
		network.SetPreferIPv6(t.preferIPv6)
		c.Check(network.SelectInternalHostPort(t.hostPorts, true), gc.DeepEquals, t.expected())
	}
}

func (*PortSuite) TestAddressesWithPort(c *gc.C) {
	addrs := network.NewAddresses("0.1.2.3", "0.2.4.6")
	hps := network.AddressesWithPort(addrs, 999)
	c.Assert(hps, jc.DeepEquals, []network.HostPort{{
		Address: network.NewAddress("0.1.2.3", network.ScopeUnknown),
		Port:    999,
	}, {
		Address: network.NewAddress("0.2.4.6", network.ScopeUnknown),
		Port:    999,
	}})
}

func (*PortSuite) TestSortHostPorts(c *gc.C) {
	hps := network.AddressesWithPort(
		network.NewAddresses(
			"127.0.0.1",
			"localhost",
			"example.com",
			"::1",
			"fc00::1",
			"fe80::2",
			"172.16.0.1",
			"8.8.8.8",
		),
		1234,
	)
	network.SortHostPorts(hps, false)
	c.Assert(hps, jc.DeepEquals, network.AddressesWithPort(
		network.NewAddresses(
			"localhost",
			"example.com",
			"127.0.0.1",
			"172.16.0.1",
			"8.8.8.8",
			"::1",
			"fc00::1",
			"fe80::2",
		),
		1234,
	))

	network.SortHostPorts(hps, true)
	c.Assert(hps, jc.DeepEquals, network.AddressesWithPort(
		network.NewAddresses(
			"localhost",
			"example.com",
			"::1",
			"fc00::1",
			"fe80::2",
			"127.0.0.1",
			"172.16.0.1",
			"8.8.8.8",
		),
		1234,
	))
}

func (p *PortSuite) TestPortRangeConflicts(c *gc.C) {
	var testCases = []struct {
		about          string
		first          network.PortRange
		second         network.PortRange
		expectConflict bool
	}{{
		"identical ports",
		network.PortRange{80, 80, "TCP"},
		network.PortRange{80, 80, "TCP"},
		true,
	}, {
		"different ports",
		network.PortRange{80, 80, "TCP"},
		network.PortRange{90, 90, "TCP"},
		false,
	}, {
		"touching ranges",
		network.PortRange{100, 200, "TCP"},
		network.PortRange{201, 240, "TCP"},
		false,
	}, {
		"touching ranges with overlap",
		network.PortRange{100, 200, "TCP"},
		network.PortRange{200, 240, "TCP"},
		true,
	}, {
		"different protocols",
		network.PortRange{80, 80, "UDP"},
		network.PortRange{80, 80, "TCP"},
		false,
	}, {
		"outside range",
		network.PortRange{100, 200, "TCP"},
		network.PortRange{80, 80, "TCP"},
		false,
	}, {
		"overlap end",
		network.PortRange{100, 200, "TCP"},
		network.PortRange{80, 120, "TCP"},
		true,
	}, {
		"complete overlap",
		network.PortRange{100, 200, "TCP"},
		network.PortRange{120, 140, "TCP"},
		true,
	}}

	for i, t := range testCases {
		c.Logf("test %d: %s", i, t.about)
		c.Check(t.first.ConflictsWith(t.second), gc.Equals, t.expectConflict)
		c.Check(t.second.ConflictsWith(t.first), gc.Equals, t.expectConflict)
	}
}

func (p *PortSuite) TestPortRangeString(c *gc.C) {
	c.Assert(network.PortRange{80, 80, "TCP"}.String(),
		gc.Equals,
		"80-80/tcp")
	c.Assert(network.PortRange{80, 100, "TCP"}.String(),
		gc.Equals,
		"80-100/tcp")
}

func (p *PortSuite) TestPortRangeValidity(c *gc.C) {
	testCases := []struct {
		about string
		ports network.PortRange
		valid bool
	}{{
		"single valid port",
		network.PortRange{80, 80, "tcp"},
		true,
	}, {
		"valid port range",
		network.PortRange{80, 90, "tcp"},
		true,
	}, {
		"valid udp port range",
		network.PortRange{80, 90, "UDP"},
		true,
	}, {
		"invalid port range boundaries",
		network.PortRange{90, 80, "tcp"},
		false,
	}, {
		"invalid protocol",
		network.PortRange{80, 80, "some protocol"},
		false,
	}}

	for i, t := range testCases {
		c.Logf("test %d: %s", i, t.about)
		c.Assert(t.ports.IsValid(), gc.Equals, t.valid)
	}
}

func (*PortSuite) TestSortPortRanges(c *gc.C) {
	ranges := []network.PortRange{
		{10, 100, "udp"},
		{80, 90, "tcp"},
		{80, 80, "tcp"},
	}
	expected := []network.PortRange{
		{80, 80, "tcp"},
		{80, 90, "tcp"},
		{10, 100, "udp"},
	}
	network.SortPortRanges(ranges)
	c.Assert(ranges, gc.DeepEquals, expected)
}
