package main

import (
	"github.com/Prof-SD/ASG/subgraph"
	"github.com/streamingfast/sparkle/cli"
	"github.com/streamingfast/sparkle/entity"
	"github.com/streamingfast/sparkle/subgraph"
)

func main() {
	subgraph.MainSubgraphDef = subgraph.Definition
	cli.Execute()
}
