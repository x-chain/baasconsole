package entity

import (
	"github.com/x-chain/baasconsole/baas-core/core/model"
	"strings"
)

func ParseFabircChain(chain *Chain) model.FabricChain {
	fc := model.FabricChain{
		ChainName:   chain.Name,
		Account:     chain.UserAccount,
		Consensus:   chain.Consensus,
		PeersOrgs:   strings.Split(chain.PeersOrgs, ","),
		OrderCount:  chain.OrderCount,
		PeerCount:   chain.PeerCount,
		ChannelName: "",
		TlsEnabled:  chain.TlsEnabled,
	}

	return fc
}

func ParseFabircChainAndChannel(chain *Chain, channel *Channel) model.FabricChain {
	fc := model.FabricChain{
		ChainName:   chain.Name,
		Account:     chain.UserAccount,
		Consensus:   chain.Consensus,
		PeersOrgs:   strings.Split(channel.Orgs, ","),
		OrderCount:  chain.OrderCount,
		PeerCount:   chain.PeerCount,
		ChannelName: channel.ChannelName,
		TlsEnabled:  chain.TlsEnabled,
	}

	return fc
}

func ParseFabircChannel(chain model.FabricChain, cc *Chaincode) model.FabricChannel {
	fc := model.FabricChannel{
		FabricChain:   chain,
		ChaincodeId:   cc.ChaincodeName,
		Version:       cc.Version,
		Policy:        cc.Policy,
		ChaincodePath: cc.GithubPath,
	}
	return fc
}
