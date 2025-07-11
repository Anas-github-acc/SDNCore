package usecase

import (
	"arieoldman/arieoldman/sdncore/common"
	"arieoldman/arieoldman/sdncore/entity"
	"github.com/golang/glog"
)

// DatapathCanHandshake is a BooleanUseCase which responds with whether the
// datapath is allowed to handshake.
type DatapathCanHandshake struct {
	Conf entity.Config
	DPID common.EthAddr
}

// Execute causes the use case to execute and update the response handler.
func (uc DatapathCanHandshake) Execute(handler BooleanResponseHandler) {
	glog.Infof("uc.DPID %v", uc.DPID)
	for _, dpid := range uc.Conf.DPIDs {
		glog.Infof("dpid: %v", dpid)
		if dpid == uc.DPID {
			defer handler.Handle(true)
		}
	}
}
