package ww_test

import (
	"testing"
	"time"

	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	ww "github.com/lthibault/wetware/pkg"
)

func TestNewHeartbeat(t *testing.T) {
	ttl := time.Second * 5

	id, err := peer.Decode("QmYyQSo1c1Ym7orWxLYvCrM2EmxFTANf8wXmmE7DWjhx5N")
	require.NoError(t, err)

	hb, err := ww.NewHeartbeat(id, ttl)
	require.NoError(t, err)

	assert.Equal(t, id, hb.ID())
	assert.Equal(t, ttl, hb.TTL())
}
