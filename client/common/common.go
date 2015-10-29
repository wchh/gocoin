package common

import (
	"fmt"
	"github.com/wchh/gocoin/lib"
	"github.com/wchh/gocoin/lib/btc"
	"github.com/wchh/gocoin/lib/chain"
	"sync"
	"time"
)

const (
	ConfigFile = "gocoin.conf"

	Version          = 70001
	DefaultUserAgent = "/Gocoin:" + lib.Version + "/"
	Services         = uint64(0x00000001)

	MaxCachedBlocks = 600
)

var (
	BlockChain   *chain.Chain
	GenesisBlock *btc.Uint256
	Magic        [4]byte
	Testnet      bool

	Last struct {
		sync.Mutex // use it for writing and reading from non-chain thread
		Block      *chain.BlockTreeNode
		time.Time
	}

	GocoinHomeDir  string
	StartTime      time.Time
	MaxPeersNeeded int

	DefaultTcpPort uint16

	MaxExpireTime time.Duration
	ExpirePerKB   time.Duration

	DebugLevel int64

	CounterMutex sync.Mutex
	Counter      map[string]uint64 = make(map[string]uint64)

	BusyWith   string
	Busy_mutex sync.Mutex

	NetworkClosed bool
)

func CountSafe(k string) {
	CounterMutex.Lock()
	Counter[k]++
	CounterMutex.Unlock()
}

func CountSafeAdd(k string, val uint64) {
	CounterMutex.Lock()
	Counter[k] += val
	CounterMutex.Unlock()
}

func Busy(b string) {
	Busy_mutex.Lock()
	BusyWith = b
	Busy_mutex.Unlock()
}

func BytesToString(val uint64) string {
	if val < 1e6 {
		return fmt.Sprintf("%.1f KB", float64(val)/1e3)
	} else if val < 1e9 {
		return fmt.Sprintf("%.2f MB", float64(val)/1e6)
	}
	return fmt.Sprintf("%.2f GB", float64(val)/1e9)
}

func NumberToString(num float64) string {
	if num > 1e15 {
		return fmt.Sprintf("%.2f P", num/1e15)
	}
	if num > 1e12 {
		return fmt.Sprintf("%.2f T", num/1e12)
	}
	if num > 1e9 {
		return fmt.Sprintf("%.2f G", num/1e9)
	}
	if num > 1e6 {
		return fmt.Sprintf("%.2f M", num/1e6)
	}
	if num > 1e3 {
		return fmt.Sprintf("%.2f K", num/1e3)
	}
	return fmt.Sprintf("%.2f", num)
}

func HashrateToString(hr float64) string {
	return NumberToString(hr) + "H/s"
}
