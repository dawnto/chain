package services

var presetStack = map[string][]serviceConstructor{
	"blockchain": {
		NewDiscoveryService,
		NewBlockChainService,
	},
}
