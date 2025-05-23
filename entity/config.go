package entity

import (
	"sdn/sdncore/common"
)

// Config is the entity which describes an SDN configuration.
type Config struct {
	DPIDs []common.EthAddr
}
