package torrent

import (
	"github.com/timechainlabs/torrent/dialer"
)

type (
	Dialer        = dialer.T
	NetworkDialer = dialer.WithNetwork
)

var DefaultNetDialer = &dialer.Default
