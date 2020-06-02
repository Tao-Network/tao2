package taox

import (
	"github.com/Tao-Network/tao2/accounts/abi/bind"
	"github.com/Tao-Network/tao2/common"
	"github.com/Tao-Network/tao2/contracts/taox/contract"
)

type TAOXListing struct {
	*contract.TAOXListingSession
	contractBackend bind.ContractBackend
}

func NewMyTAOXListing(transactOpts *bind.TransactOpts, contractAddr common.Address, contractBackend bind.ContractBackend) (*TAOXListing, error) {
	smartContract, err := contract.NewTAOXListing(contractAddr, contractBackend)
	if err != nil {
		return nil, err
	}

	return &TAOXListing{
		&contract.TAOXListingSession{
			Contract:     smartContract,
			TransactOpts: *transactOpts,
		},
		contractBackend,
	}, nil
}

func DeployTAOXListing(transactOpts *bind.TransactOpts, contractBackend bind.ContractBackend) (common.Address, *TAOXListing, error) {
	contractAddr, _, _, err := contract.DeployTAOXListing(transactOpts, contractBackend)
	if err != nil {
		return contractAddr, nil, err
	}
	smartContract, err := NewMyTAOXListing(transactOpts, contractAddr, contractBackend)
	if err != nil {
		return contractAddr, nil, err
	}

	return contractAddr, smartContract, nil
}
