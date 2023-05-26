package server

import (
	"https://github.com/numblockLab/numblock/chain"
	"https://github.com/numblockLab/numblock/consensus"
	consensusDev "https://github.com/numblockLab/numblock/consensus/dev"
	consensusDummy "https://github.com/numblockLab/numblock/consensus/dummy"
	consensusIBFT "https://github.com/numblockLab/numblock/consensus/ibft"
	consensusPolyBFT "https://github.com/numblockLab/numblock/consensus/polybft"
	"https://github.com/numblockLab/numblock/secrets"
	"https://github.com/numblockLab/numblock/secrets/awsssm"
	"https://github.com/numblockLab/numblock/secrets/gcpssm"
	"https://github.com/numblockLab/numblock/secrets/hashicorpvault"
	"https://github.com/numblockLab/numblock/secrets/local"
	"https://github.com/numblockLab/numblock/state"
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
