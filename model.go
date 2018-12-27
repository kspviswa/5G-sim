package main

import (
	net "net"

	uuid "github.com/satori/go.uuid"
)

// Basic Defnitions

// Infrastructure Model - What goes at Cloud level
type infra struct {
	ram  uint64
	vcpu uint32
	hdd  uint64
}

// Description of a cloud - No discrimination between Edge / Site / DC / Core Clouds.
type cloudinfra struct {
	provisionedinfra infra
	provisionednw    nw
	infraused        float32
	nwused           float32
	name             string
	id               uuid.UUID
	endpoint         string
	tenant           string
	region           string
	nfconsumers      []uuid.UUID
	nfconsumerptrs   []*nfi
}

type hpaType string

// List of possible HPA types
const (
	DPDK  hpaType = "dpdk"
	SRIOV hpaType = "sriov"
	NUMA  hpaType = "numa"
)

// Describe a Radio Network ( RAN )

type raninfra struct {
	name     string
	id       uuid.UUID
	latency  uint32  // Latency in milliseconds
	numRAB   uint32  // Number of RadioAccessBearers ( RAB )
	spectrum float32 // Spectrum in GHz
}

type nfType string

// List of possible network types
const (
	RAN nfType = "radio"
	CN  nfType = "core"
	DN  nfType = "internet"
)

// A Network Model - What goes into connectivity
type nw struct {
	IPs  []net.IP
	HPAs []hpaType
	//IP net.IP
}

// A Network Function - A node with a purpose - backed by Infra & Network Capabilities
// This is a design time entry
type nf struct {
	name        string
	desp        string
	id          uuid.UUID
	capInfra    infra
	nwInfra     nw
	cardinality uint32
	isDataPlane bool
	numSessions uint64
	category    nfType
}

// A Network Function Instance - An runtime instance of NF
type nfi struct {
	name          string
	nfcatalogID   uuid.UUID
	cloudinfraID  uuid.UUID
	subsliceIDs   []uuid.UUID // A NFI could be part of multiple subslices
	sliceIDs      []uuid.UUID // A NFI could be indirectly part of multiple slices
	nfcatalogptr  *nf         // Pointer for easy access
	cloudinfraptr *cloudinfra // Pointer for easy access
	subsliceptrs  []*subslice // Pointer for easy access
	sliceptrs     []*slice    // Pointer for easy access
	nwasset       nw          // Runtime allocation of IP address
}

// A SubSlice - Collection of network functions aka NetworkInfra Service
type subslice struct {
	id          uuid.UUID
	name        string
	thesubslice []*nfi
}

// A Slice - Collection of Sub Slices
type slice struct {
	id       uuid.UUID
	name     string
	theslice []*subslice
}

// Slice Selection Type

type sst string

const (
	eMBB  sst = "embb"
	uRLLC sst = "urrlc"
	mIOT  sst = "miot"
)

// A Network Service - Collection of Slices that satisfies a Customer Oderable Service aka User Service
type service struct {
	name      string
	id        uuid.UUID
	slicetype sst
	fns       []*slice
}

type subscription struct {
	name            string
	id              uuid.UUID
	connectedUeId   uuid.UUID
	services        []*service
	connectedStatus bool
	connectedgNBs   []*nf
	connectedSlices [8]uuid.UUID
}

// A Global Network - Collection of Slices with mapped Services
type simulation struct {
	theinfra         []*cloudinfra
	thenetwork       []*slice
	theservices      map[string]*service
	thesubscriptions []*subscription
}
