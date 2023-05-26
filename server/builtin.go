package server

import (
	"github.com/numblockLab/numblock/chain"
	"github.com/numblockLab/numblock/consensus"
	consensusDev "github.com/numblockLab/numblock/consensus/dev"
	consensusDummy "github.com/numblockLab/numblock/consensus/dummy"
	consensusIBFT "github.com/numblockLab/numblock/consensus/ibft"
	consensusPolyBFT "github.com/numblockLab/numblock/consensus/polybft"
	"github.com/numblockLab/numblock/secrets"
	"github.com/numblockLab/numblock/secrets/awsssm"
	"github.com/numblockLab/numblock/secrets/gcpssm"
	"github.com/numblockLab/numblock/secrets/hashicorpvault"
	"github.com/numblockLab/numblock/secrets/local"
	"github.com/numblockLab/numblock/state"
)

type GenesisFactoryHook func(config *chain.Chain, engineName string) func(*state.Transition) error

type ConsensusType string

const (
	DevConsensus     ConsensusType = "dev"
	IBFTConsensus    ConsensusType = "ibft"
	PolyBFTConsensus ConsensusType = "polybft"
	DummyConsensus   ConsensusType = "dummy"
)

var consensusBackends = map[ConsensusType]consensus.Factory{
	DevConsensus:     consensusDev.Factory,
	IBFTConsensus:    consensusIBFT.Factory,
	PolyBFTConsensus: consensusPolyBFT.Factory,
	DummyConsensus:   consensusDummy.Factory,
}

// secretsManagerBackends defines the SecretManager factories for different
// secret management solutions
var secretsManagerBackends = map[secrets.SecretsManagerType]secrets.SecretsManagerFactory{
	secrets.Local:          local.SecretsManagerFactory,
	secrets.HashicorpVault: hashicorpvault.SecretsManagerFactory,
	secrets.AWSSSM:         awsssm.SecretsManagerFactory,
	secrets.GCPSSM:         gcpssm.SecretsManagerFactory,
}

var genesisCreationFactory = map[ConsensusType]GenesisFactoryHook{
	PolyBFTConsensus: consensusPolyBFT.GenesisPostHookFactory,
}

func ConsensusSupported(value string) bool {
	_, ok := consensusBackends[ConsensusType(value)]

	return ok
}
