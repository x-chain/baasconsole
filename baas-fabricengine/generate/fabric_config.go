package generate

import (
	"github.com/x-chain/baasconsole/baas-core/common/util"
	"github.com/x-chain/baasconsole/baas-core/core/model"
	"github.com/x-chain/baasconsole/baas-fabricengine/config"
	"github.com/x-chain/baasconsole/baas-fabricengine/constant"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
)

//builder
type ConfigBuilder interface {
	//configtx.yaml
	SetOrganizations() FabricConfig
	SetCapabilities() FabricConfig
	SetApplication() FabricConfig
	SetOrderer() FabricConfig
	SetChannel() FabricConfig
	SetProfiles() FabricConfig
	BuildTxFile()
	//crypto-feconfig.yaml
	SetOrdererOrgs() FabricConfig
	SetPeerOrgs() FabricConfig
	BuildCryptoFile()
}

//fabric 配置
type FabricConfig struct {
	model.FabricChain

	CryptoConfigDir  string //证书目录
	ConfigtxFile     string //tx文件
	CryptoConfigFile string //crypto-config文件

	//共识配置
	OrdererBatchTimeout      string //2s
	OrdererMaxMessageCount   int    //500
	OrdererAbsoluteMaxBytes  string //99 MB
	OrdererPreferredMaxBytes string //512 KB

	configtx     Configtx
	cryptoConfig CryptoConfig
}

//tx实体
type Configtx struct {
	Organizations []Organization `yaml:"-"`
	Capabilities  Capabilities   `yaml:"-"`
	Application   Application    `yaml:"-"`
	Orderer       Orderer        `yaml:"-"`
	Channel       Channel        `yaml:"-"`
	Profiles      Profiles       `yaml:"Profiles"`
}

type Organization struct {
	Name        string       `yaml:"Name"`
	ID          string       `yaml:"ID"`
	MSPDir      string       `yaml:"MSPDir"`
	Policies    Policies     `yaml:"Policies"`
	AnchorPeers []AnchorPeer `yaml:"AnchorPeers"`
}

type Policies struct {
	Readers TypeRule `yaml:"Readers"`
	Writers TypeRule `yaml:"Writers"`
	Admins  TypeRule `yaml:"Admins"`
}

type AnchorPeer struct {
	Host string `yaml:"Host"`
	Port int    `yaml:"Port"`
}

type TypeRule struct {
	Type string `yaml:"Type"`
	Rule string `yaml:"Rule"`
}

type Capabilities struct {
	Channel     ChannelCapabilities     `yaml:"Channel"`
	Orderer     OrdererCapabilities     `yaml:"Orderer"`
	Application ApplicationCapabilities `yaml:"Application"`
}
type ChannelCapabilities struct {
	V1_3 bool `yaml:"V1_3"`
}
type OrdererCapabilities struct {
	V1_1 bool `yaml:"V1_1"`
}
type ApplicationCapabilities struct {
	V1_3 bool `yaml:"V1_3"`
	V1_2 bool `yaml:"V1_2"`
	V1_1 bool `yaml:"V1_1"`
}
type Application struct {
	Organizations string                  `yaml:"Organizations"`
	Policies      Policies                `yaml:"Policies"`
	Capabilities  ApplicationCapabilities `yaml:"Capabilities"`
}

type Orderer struct {
	OrdererType  string   `yaml:"OrdererType"`
	Addresses    []string `yaml:"Addresses"`
	BatchTimeout string   `yaml:"BatchTimeout"`
	BatchSize    struct {
		MaxMessageCount   int    `yaml:"MaxMessageCount"`
		AbsoluteMaxBytes  string `yaml:"AbsoluteMaxBytes"`
		PreferredMaxBytes string `yaml:"PreferredMaxBytes"`
	} `yaml:"BatchSize"`
	Kafka struct {
		Brokers []string `yaml:"Brokers"`
	} `yaml:"Kafka"`
	EtcdRaft struct {
		Consenters []Consenter `yaml:"Consenters"`
	} `yaml:"EtcdRaft"`
	Organizations string          `yaml:"Organizations"`
	Policies      OrdererPolicies `yaml:"Policies"`
}

type Consenter struct {
	Host          string `yaml:"Host"`
	Port          int    `yaml:"Port"`
	ClientTLSCert string `yaml:"ClientTLSCert"`
	ServerTLSCert string `yaml:"ServerTLSCert"`
}

type OrdererPolicies struct {
	Readers         TypeRule `yaml:"Readers"`
	Writers         TypeRule `yaml:"Writers"`
	Admins          TypeRule `yaml:"Admins"`
	BlockValidation TypeRule `yaml:"BlockValidation"`
}

type Channel struct {
	Policies     Policies            `yaml:"Policies"`
	Capabilities ChannelCapabilities `yaml:"Capabilities"`
}

type Profiles struct {
	OrdererGenesis OrdererGenesis `yaml:"OrdererGenesis"`
	OrgsChannel    OrgsChannel    `yaml:"OrgsChannel"`
}

type OrdererGenesis struct {
	Policies     Policies            `yaml:"Policies"`
	Capabilities ChannelCapabilities `yaml:"Capabilities"`
	Orderer      OgOrderer           `yaml:"Orderer"`
	Consortiums  Consortiums         `yaml:"Consortiums"`
}
type OgOrderer struct {
	OrdererType  string   `yaml:"OrdererType"`
	Addresses    []string `yaml:"Addresses"`
	BatchTimeout string   `yaml:"BatchTimeout"`
	BatchSize    struct {
		MaxMessageCount   int    `yaml:"MaxMessageCount"`
		AbsoluteMaxBytes  string `yaml:"AbsoluteMaxBytes"`
		PreferredMaxBytes string `yaml:"PreferredMaxBytes"`
	} `yaml:"BatchSize"`
	Kafka struct {
		Brokers []string `yaml:"Brokers"`
	} `yaml:"Kafka"`
	EtcdRaft struct {
		Consenters []Consenter `yaml:"Consenters"`
	} `yaml:"EtcdRaft"`
	Policies      OrdererPolicies     `yaml:"Policies"`
	Organizations []Organization      `yaml:"Organizations"`
	Capabilities  OrdererCapabilities `yaml:"Capabilities"`
}
type Consortiums struct {
	SampleConsortium struct {
		Organizations []Organization `yaml:"Organizations"`
	} `yaml:"SampleConsortium"`
}

type OrgsChannel struct {
	Consortium  string        `yaml:"Consortium"`
	Application OcApplication `yaml:"Application"`
}
type OcApplication struct {
	Policies      Policies                `yaml:"Policies"`
	Organizations []Organization          `yaml:"Organizations"`
	Capabilities  ApplicationCapabilities `yaml:"Capabilities"`
}

//crypto-feconfig.yaml
type CryptoConfig struct {
	OrdererOrgs []OrdererOrg `yaml:"OrdererOrgs"`
	PeerOrgs    []PeerOrg    `yaml:"PeerOrgs"`
}

type OrdererOrg struct {
	Name   string `yaml:"Name"`
	Domain string `yaml:"Domain"`
	CA     struct {
		Country  string `yaml:"Country"`
		Province string `yaml:"Province"`
		Locality string `yaml:"Locality"`
	} `yaml:"CA"`
	Template struct {
		Count int `yaml:"Count"`
	} `yaml:"Template"`
}

func NewOrdererOrg(domain string, ordererCount int) OrdererOrg {
	return OrdererOrg{
		Name:   "Orderer",
		Domain: domain,
		CA: struct {
			Country  string `yaml:"Country"`
			Province string `yaml:"Province"`
			Locality string `yaml:"Locality"`
		}{
			Country:  constant.Country,
			Province: constant.Province,
			Locality: constant.Locality,
		},
		Template: struct {
			Count int `yaml:"Count"`
		}{
			Count: ordererCount,
		},
	}
}

type PeerOrg struct {
	Name   string `yaml:"Name"`
	Domain string `yaml:"Domain"`
	CA     struct {
		Country  string `yaml:"Country"`
		Province string `yaml:"Province"`
		Locality string `yaml:"Locality"`
	} `yaml:"CA"`
	Template struct {
		Count int `yaml:"Count"`
	} `yaml:"Template"`
	Users struct {
		Count int `yaml:"Count"`
	} `yaml:"Users"`
	EnableNodeOUs bool `yaml:"EnableNodeOUs"`
}

func NewPeerOrg(name, domain string, peerCount int) PeerOrg {
	return PeerOrg{
		Name:   name,
		Domain: domain,
		CA: struct {
			Country  string `yaml:"Country"`
			Province string `yaml:"Province"`
			Locality string `yaml:"Locality"`
		}{
			Country:  constant.Country,
			Province: constant.Province,
			Locality: constant.Locality,
		},
		Template: struct {
			Count int `yaml:"Count"`
		}{
			Count: peerCount,
		},
		EnableNodeOUs: true,
		Users: struct {
			Count int `yaml:"Count"`
		}{
			Count: 1,
		},
	}

}

//设置organization
func (f FabricConfig) SetOrganizations() FabricConfig {
	orgs := make([]Organization, len(f.PeersOrgs)+1)
	orderOrg := Organization{
		Name:   "Orderer",
		ID:     constant.OrdererMsp,
		MSPDir: f.CryptoConfigDir + "/ordererOrganizations/" + f.GetHostDomain(constant.OrdererSuffix) + "/msp",
		Policies: Policies{
			Readers: TypeRule{
				Type: constant.TypeSignature,
				Rule: "OR('" + constant.OrdererMsp + ".member')",
			},
			Writers: TypeRule{
				Type: constant.TypeSignature,
				Rule: "OR('" + constant.OrdererMsp + ".member')",
			},
			Admins: TypeRule{
				Type: constant.TypeSignature,
				Rule: "OR('" + constant.OrdererMsp + ".admin')",
			},
		},
	}
	orgs[0] = orderOrg
	for i, v := range f.PeersOrgs {
		name := util.FirstUpper(v)
		peerOrg := Organization{
			Name:   name,
			ID:     name + constant.MspSuf,
			MSPDir: f.CryptoConfigDir + "/peerOrganizations/" + f.GetHostDomain(v) + "/msp",
			Policies: Policies{
				Readers: TypeRule{
					Type: constant.TypeSignature,
					Rule: "OR('" + name + constant.MspSuf + ".admin', '" + name + constant.MspSuf + ".peer', '" + name + constant.MspSuf + ".client')",
				},
				Writers: TypeRule{
					Type: constant.TypeSignature,
					Rule: "OR('" + name + constant.MspSuf + ".admin', '" + name + constant.MspSuf + ".client')",
				},
				Admins: TypeRule{
					Type: constant.TypeSignature,
					Rule: "OR('" + name + constant.MspSuf + ".admin')",
				},
			},
			AnchorPeers: []AnchorPeer{
				AnchorPeer{
					Host: "peer0." + f.GetHostDomain(v),
					Port: 7051,
				},
			},
		}
		orgs[i+1] = peerOrg
	}

	f.configtx.Organizations = orgs
	return f
}

//设置capabilitie
func (f FabricConfig) SetCapabilities() FabricConfig {
	capabilities := Capabilities{
		Channel: ChannelCapabilities{
			V1_3: true,
		},
		Orderer: OrdererCapabilities{
			V1_1: true,
		},
		Application: ApplicationCapabilities{
			V1_3: true,
			V1_2: false,
			V1_1: false,
		},
	}
	f.configtx.Capabilities = capabilities
	return f
}

//设置Application
func (f FabricConfig) SetApplication() FabricConfig {
	application := Application{
		Organizations: "",
		Policies: Policies{
			Readers: TypeRule{
				Type: constant.TypeImplicitMeta,
				Rule: constant.RuleAnyReaders,
			},
			Writers: TypeRule{
				Type: constant.TypeImplicitMeta,
				Rule: constant.RuleAnyWriters,
			},
			Admins: TypeRule{
				Type: constant.TypeImplicitMeta,
				Rule: constant.RuleMajorityAdmins,
			},
		},
		Capabilities: f.configtx.Capabilities.Application,
	}
	f.configtx.Application = application
	return f
}

//设置Orderer
func (f FabricConfig) SetOrderer() FabricConfig {

	orderer := Orderer{
		OrdererType:  f.Consensus,
		BatchTimeout: f.OrdererBatchTimeout,
		BatchSize: struct {
			MaxMessageCount   int    `yaml:"MaxMessageCount"`
			AbsoluteMaxBytes  string `yaml:"AbsoluteMaxBytes"`
			PreferredMaxBytes string `yaml:"PreferredMaxBytes"`
		}{
			MaxMessageCount:   f.OrdererMaxMessageCount,
			AbsoluteMaxBytes:  f.OrdererAbsoluteMaxBytes,
			PreferredMaxBytes: f.OrdererPreferredMaxBytes,
		},
	}
	orderer.Policies = OrdererPolicies{
		Readers: TypeRule{
			Type: constant.TypeImplicitMeta,
			Rule: constant.RuleAnyReaders,
		},
		Writers: TypeRule{
			Type: constant.TypeImplicitMeta,
			Rule: constant.RuleAnyWriters,
		},
		Admins: TypeRule{
			Type: constant.TypeImplicitMeta,
			Rule: constant.RuleMajorityAdmins,
		},
		BlockValidation: TypeRule{
			Type: constant.TypeImplicitMeta,
			Rule: constant.RuleAnyWriters,
		},
	}
	switch f.Consensus {
	case constant.OrdererSolo:
		domain := "orderer0." + f.GetHostDomain(constant.OrdererSuffix) + ":7050"
		orderer.Addresses = []string{domain}
	case constant.OrdererKafka:
		domains := make([]string, f.OrderCount)
		for i := 0; i < f.OrderCount; i++ {
			domain := "orderer" + strconv.Itoa(i) + "." + f.GetHostDomain(constant.OrdererSuffix) + ":7050"
			domains[i] = domain
		}
		kafka := struct {
			Brokers []string `yaml:"Brokers"`
		}{
			Brokers: []string{
				"kafka0." + f.GetHostDomain(constant.KafkaSuffix) + ":9092",
				"kafka1." + f.GetHostDomain(constant.KafkaSuffix) + ":9092",
				"kafka2." + f.GetHostDomain(constant.KafkaSuffix) + ":9092",
				"kafka3." + f.GetHostDomain(constant.KafkaSuffix) + ":9092",
			},
		}
		orderer.Addresses = domains
		orderer.Kafka = kafka
	case constant.OrdererEtcdraft:
		domains := make([]string, f.OrderCount)
		consenters := make([]Consenter, f.OrderCount)
		for i := 0; i < f.OrderCount; i++ {
			domain := "orderer" + strconv.Itoa(i) + "." + f.GetHostDomain(constant.OrdererSuffix)
			domains[i] = domain + ":7050"
			tls := f.CryptoConfigDir + "/ordererOrganizations/" + f.GetHostDomain(constant.OrdererSuffix) + "/orderers/" + domain + "/tls/server.crt"
			consenter := Consenter{
				Host:          domain,
				Port:          7050,
				ClientTLSCert: tls,
				ServerTLSCert: tls,
			}
			consenters[i] = consenter
		}
		orderer.Addresses = domains
		orderer.EtcdRaft = struct {
			Consenters []Consenter `yaml:"Consenters"`
		}{
			Consenters: consenters,
		}
	}
	f.configtx.Orderer = orderer
	return f
}

//设置channel
func (f FabricConfig) SetChannel() FabricConfig {
	channel := Channel{
		Policies: Policies{
			Readers: TypeRule{
				Type: constant.TypeImplicitMeta,
				Rule: constant.RuleAnyReaders,
			},
			Writers: TypeRule{
				Type: constant.TypeImplicitMeta,
				Rule: constant.RuleAnyWriters,
			},
			Admins: TypeRule{
				Type: constant.TypeImplicitMeta,
				Rule: constant.RuleMajorityAdmins,
			},
		},
		Capabilities: f.configtx.Capabilities.Channel,
	}
	f.configtx.Channel = channel
	return f
}

//设置Profile
func (f FabricConfig) SetProfiles() FabricConfig {
	//OrdererGenesis
	ordererGenesis := OrdererGenesis{}
	peerOrgs := f.configtx.Organizations[1:]
	sampleConsortium := struct {
		Organizations []Organization `yaml:"Organizations"`
	}{
		Organizations: peerOrgs,
	}
	consortiums := Consortiums{
		SampleConsortium: sampleConsortium,
	}
	//OrdererGenesis.Consortiums
	ordererGenesis.Consortiums = consortiums
	ordererGenesis.Policies = f.configtx.Channel.Policies
	ordererGenesis.Capabilities = f.configtx.Channel.Capabilities

	orderOrg := make([]Organization, 1)
	orderOrg[0] = f.configtx.Organizations[0]

	order := OgOrderer{
		OrdererType:   f.configtx.Orderer.OrdererType,
		Policies:      f.configtx.Orderer.Policies,
		Kafka:         f.configtx.Orderer.Kafka,
		EtcdRaft:      f.configtx.Orderer.EtcdRaft,
		Addresses:     f.configtx.Orderer.Addresses,
		BatchSize:     f.configtx.Orderer.BatchSize,
		BatchTimeout:  f.configtx.Orderer.BatchTimeout,
		Organizations: orderOrg,
		Capabilities:  f.configtx.Capabilities.Orderer,
	}

	//OrdererGenesis.Orderer
	ordererGenesis.Orderer = order
	//OrgsChannel
	//OrgsChannel.Application
	application := OcApplication{
		Policies:      f.configtx.Application.Policies,
		Capabilities:  f.configtx.Capabilities.Application,
		Organizations: peerOrgs,
	}

	orgsChannel := OrgsChannel{
		Consortium:  "SampleConsortium",
		Application: application,
	}
	// Profiles
	profiles := Profiles{
		OrdererGenesis: ordererGenesis,
		OrgsChannel:    orgsChannel,
	}
	f.configtx.Profiles = profiles
	return f
}

//建tx文件
func (f FabricConfig) BuildTxFile() {
	f = f.SetOrganizations().SetCapabilities().SetApplication().SetOrderer().SetChannel().SetProfiles()
	tx, err := yaml.Marshal(&f.configtx)
	if err != nil {
		logger.Errorf("error: %v", err)
	}
	logger.Info(string(tx))
	ioutil.WriteFile(f.ConfigtxFile, tx, os.ModePerm)
}

func (f FabricConfig) SetOrdererOrgs() FabricConfig {
	ordererOrgs := make([]OrdererOrg, 1)
	var order OrdererOrg
	switch f.Consensus {
	case constant.OrdererSolo:
		order = NewOrdererOrg(f.GetHostDomain(constant.OrdererSuffix), 1)
	case constant.OrdererKafka, constant.OrdererEtcdraft:
		order = NewOrdererOrg(f.GetHostDomain(constant.OrdererSuffix), f.OrderCount)
	}
	ordererOrgs[0] = order
	f.cryptoConfig.OrdererOrgs = ordererOrgs
	return f
}

func (f FabricConfig) SetPeerOrgs() FabricConfig {
	peerOrgs := make([]PeerOrg, len(f.PeersOrgs))
	for i, v := range f.PeersOrgs {
		peer := NewPeerOrg(util.FirstUpper(v), f.GetHostDomain(v), f.PeerCount)
		peerOrgs[i] = peer
	}
	f.cryptoConfig.PeerOrgs = peerOrgs
	return f

}

func (f FabricConfig) BuildCryptoFile() {
	fc := f.SetOrdererOrgs().SetPeerOrgs()
	crypto, err := yaml.Marshal(&fc.cryptoConfig)
	if err != nil {
		logger.Errorf("error: %v", err)
	}
	logger.Info(string(crypto))
	ioutil.WriteFile(f.CryptoConfigFile, crypto, os.ModePerm)
}

func NewConfigBuilder(chain model.FabricChain, rootPath string) ConfigBuilder {

	fconfig := FabricConfig{
		CryptoConfigFile: filepath.Join(rootPath, constant.CryptoConfigYaml),
		ConfigtxFile:     filepath.Join(rootPath, constant.ConfigtxYaml),
		CryptoConfigDir:  filepath.Join(rootPath, constant.CryptoConfigDir),

		OrdererBatchTimeout:      config.Config.GetString("OrdererBatchTimeout"),
		OrdererMaxMessageCount:   config.Config.GetInt("OrdererMaxMessageCount"),
		OrdererAbsoluteMaxBytes:  config.Config.GetString("OrdererAbsoluteMaxBytes"),
		OrdererPreferredMaxBytes: config.Config.GetString("OrdererPreferredMaxBytes"),

		cryptoConfig: CryptoConfig{},
		configtx:     Configtx{},
	}

	fconfig.ChainName = chain.ChainName
	fconfig.Account = chain.Account
	fconfig.Consensus = chain.Consensus
	fconfig.PeersOrgs = chain.PeersOrgs
	fconfig.OrderCount = chain.OrderCount
	fconfig.PeerCount = chain.PeerCount

	return fconfig
}
