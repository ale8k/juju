// Copyright 2015 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package testing

import (
	"path/filepath"
	"runtime"

	"github.com/juju/names/v4"
	"github.com/juju/worker/v3"
	gc "gopkg.in/check.v1"

	"github.com/juju/juju/core/leadership"
	"github.com/juju/juju/juju/sockets"
	"github.com/juju/juju/storage"
)

// RealPaths implements Paths for tests that do touch the filesystem.
type RealPaths struct {
	tools        string
	charm        string
	base         string
	socket       sockets.Socket
	metricsspool string
}

func osDependentSockPath(c *gc.C) sockets.Socket {
	sockPath := filepath.Join(c.MkDir(), "test.sock")
	if runtime.GOOS == "windows" {
		return sockets.Socket{Network: "unix", Address: `\\.\pipe` + sockPath[2:]}
	}
	return sockets.Socket{Network: "unix", Address: sockPath}
}

func NewRealPaths(c *gc.C) RealPaths {
	return RealPaths{
		tools:        c.MkDir(),
		charm:        c.MkDir(),
		base:         c.MkDir(),
		socket:       osDependentSockPath(c),
		metricsspool: c.MkDir(),
	}
}

func (p RealPaths) GetMetricsSpoolDir() string {
	return p.metricsspool
}

func (p RealPaths) GetToolsDir() string {
	return p.tools
}

func (p RealPaths) GetCharmDir() string {
	return p.charm
}

func (p RealPaths) GetBaseDir() string {
	return p.base
}

func (p RealPaths) GetJujucClientSocket(remote bool) sockets.Socket {
	return p.socket
}

func (p RealPaths) GetJujucServerSocket(remote bool) sockets.Socket {
	return p.socket
}

func (p RealPaths) GetResourcesDir() string {
	return filepath.Join(p.base, "resources")
}

type ContextStorage struct {
	CTag      names.StorageTag
	CKind     storage.StorageKind
	CLocation string
}

func (c *ContextStorage) Tag() names.StorageTag {
	return c.CTag
}

func (c *ContextStorage) Kind() storage.StorageKind {
	return c.CKind
}

func (c *ContextStorage) Location() string {
	return c.CLocation
}

type FakeTracker struct {
	leadership.Tracker
	worker.Worker

	AllowClaimLeader bool
}

func (t *FakeTracker) ApplicationName() string {
	return "application-name"
}

func (t *FakeTracker) ClaimLeader() leadership.Ticket {
	return &FakeTicket{t.AllowClaimLeader}
}

type FakeTicket struct {
	WaitResult bool
}

var _ leadership.Ticket = &FakeTicket{}

func (ft *FakeTicket) Wait() bool {
	return ft.WaitResult
}

func (ft *FakeTicket) Ready() <-chan struct{} {
	c := make(chan struct{})
	close(c)
	return c
}
