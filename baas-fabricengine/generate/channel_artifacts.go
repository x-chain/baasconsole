package generate

import (
	"github.com/x-chain/baasconsole/baas-core/common/log"
	"github.com/x-chain/baasconsole/baas-core/common/util"
	"github.com/x-chain/baasconsole/baas-core/core/model"
	"github.com/x-chain/baasconsole/baas-core/core/tools"
	"github.com/x-chain/baasconsole/baas-fabricengine/constant"
	"path/filepath"
	"sync"
)

var logger = log.GetLogger("fabricengine.generate", log.INFO)

type ChannelArtifacts struct {
	channelName         string   //channel 名称
	cryptoConfigFile    string   //配置文件
	cryptoConfigDir     string   //证书保存目录
	channelArtifactsDir string   //创世区块channel交易保存目率
	rootPath            string   //chain根目录
	anchorOrgs          []string //锚节点组织
}

//生成证书
func (c *ChannelArtifacts) GenerateCerts() {
	gen := tools.NewCryptogen(c.cryptoConfigFile, c.cryptoConfigDir)
	gen.Exec("generate")
}

//生成创世块
func (c *ChannelArtifacts) GenerateOrdererGenesis() {
	txgen := tools.NewConfigtxgen()
	txgen.SetConfigPath(c.rootPath)
	txgen.SetProfile(constant.ProfilesGenesis)
	txgen.SetChannelID(constant.BaasFirstChannel)
	txgen.SetOutputBlock(filepath.Join(c.channelArtifactsDir, constant.GenesisBlock))
	txgen.Exec()
}

//生成channel交易
func (c *ChannelArtifacts) GenerateOrgsChannel() {
	txgen := tools.NewConfigtxgen()
	txgen.SetConfigPath(c.rootPath)
	txgen.SetProfile(constant.ProfilesChannel)
	txgen.SetOutputChannelCreateTx(filepath.Join(c.channelArtifactsDir, c.channelName+constant.Tx))
	txgen.SetChannelID(c.channelName)
	txgen.Exec()
}

//生成锚节点交易
func (c *ChannelArtifacts) GenerateAnchorPeers() {
	var wg sync.WaitGroup

	for _, v := range c.anchorOrgs {
		wg.Add(1)
		go func(o string) {

			org := util.FirstUpper(o)
			txgen := tools.NewConfigtxgen()
			txgen.SetConfigPath(c.rootPath)
			txgen.SetProfile(constant.ProfilesChannel)
			txgen.SetOutputAnchorPeersUpdate(filepath.Join(c.channelArtifactsDir, c.channelName+org+constant.AnchorsTx))
			txgen.SetChannelID(c.channelName)
			txgen.SetAsOrg(org)
			txgen.Exec()

			wg.Done()
		}(v)

	}
	wg.Wait()
}

func NewChannelArtifacts(chain model.FabricChain, rootPath string) *ChannelArtifacts {
	return &ChannelArtifacts{
		channelName:         chain.ChannelName,
		cryptoConfigFile:    filepath.Join(rootPath, constant.CryptoConfigYaml),
		cryptoConfigDir:     filepath.Join(rootPath, constant.CryptoConfigDir),
		channelArtifactsDir: filepath.Join(rootPath, constant.ChannelArtifactsDir),
		rootPath:            rootPath,
		anchorOrgs:          chain.PeersOrgs,
	}
}
