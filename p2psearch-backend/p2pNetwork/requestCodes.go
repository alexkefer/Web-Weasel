package p2pNetwork

// This is basically the go equivalent of an enum (a bunch of related constants)
const (
	PingRequest = iota
	PingResponse
	AddMeRequest
	SharePeersRequest
	SharePeersResponse
	BroadcastMessage
	RemoveMeRequest
)
