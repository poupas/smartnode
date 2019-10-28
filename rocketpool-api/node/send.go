package node

import (
    "github.com/ethereum/go-ethereum/common"
    "github.com/urfave/cli"

    "github.com/rocket-pool/smartnode/shared/api/node"
    "github.com/rocket-pool/smartnode/shared/services"
    "github.com/rocket-pool/smartnode/shared/utils/api"
    "github.com/rocket-pool/smartnode/shared/utils/eth"
)


// Send resources from the node account to an address
func sendFromNode(c *cli.Context, address string, amount float64, unit string) error {

    // Initialise services
    p, err := services.NewProvider(c, services.ProviderOpts{
        AM: true,
        Client: true,
        CM: true,
        LoadContracts: []string{"rocketETHToken", "rocketPoolToken"},
        ClientConn: true,
        ClientSync: true,
        RocketStorage: true,
    })
    if err != nil { return err }
    defer p.Cleanup()

    // Send from node & print response
    if response, err := node.SendFromNode(p, common.HexToAddress(address), eth.EthToWei(amount), unit); err != nil {
        return err
    } else {
        api.PrintResponse(p.Output, response)
        return nil
    }

}

