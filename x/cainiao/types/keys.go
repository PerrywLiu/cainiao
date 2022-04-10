package types

const (
	// ModuleName defines the module name
	ModuleName = "cainiao"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey is the message route for slashing
	RouterKey = ModuleName

	// QuerierRoute defines the module's query routing key
	QuerierRoute = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_cainiao"
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

const (
	OrdersKey      = "Orders-value-"
	OrdersCountKey = "Orders-count-"
)

const (
	StateWait = "wait"
	StateStart = "start"
	StateOntheWay = "ontheway"
	StateReceive = "receive"
)

