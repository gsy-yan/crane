package model

import "encoding/json"

import "github.com/docker/engine-api/types/swarm"

type UpdateOptions struct {
	Method  string          `json:"Method"`
	Options json.RawMessage `json:"Options"`
}

type JoiningNode struct {
	Role     swarm.NodeRole `json:"Role", required:"true"`
	Endpoint string         `json:"Endpoint", required:"true"`
}
