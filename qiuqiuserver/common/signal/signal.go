package signal

import "os"

type SignalCallBack func(sig os.Signal)

var (
	Signals  []os.Signal
	signalCb SignalCallBack
)

func ListenSignal(sigs []os.Signal, cb SignalCallBack) {
	if len(sigs) == 0 || cb == nil {
		return
	}

	Signals = sigs
	signalCb = cb
}
