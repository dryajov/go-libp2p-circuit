package relay

import (
	"fmt"
	"net"

	ic "github.com/libp2p/go-libp2p-crypto"
	iconn "github.com/libp2p/go-libp2p-interface-conn"
	inet "github.com/libp2p/go-libp2p-net"
	peer "github.com/libp2p/go-libp2p-peer"
	tpt "github.com/libp2p/go-libp2p-transport"
	ma "github.com/multiformats/go-multiaddr"
	manet "github.com/multiformats/go-multiaddr-net"
)

type Conn struct {
	inet.Stream
	remoteMaddr ma.Multiaddr
	remotePeer  peer.ID
}

var _ iconn.Conn = (*Conn)(nil)

type NetAddr struct {
	Relay  string
	Remote string
}

func (n *NetAddr) Network() string {
	return "libp2p-circuit-relay"
}

func (n *NetAddr) String() string {
	return fmt.Sprintf("relay[%s-%s]", n.Remote, n.Relay)
}

func (c *Conn) RemoteAddr() net.Addr {
	return &NetAddr{
		Relay:  c.Conn().RemotePeer().Pretty(),
		Remote: c.remotePeer.String(),
	}
}

func (c *Conn) RemoteMultiaddr() ma.Multiaddr {
	a, err := ma.NewMultiaddr(fmt.Sprintf("%s/ipfs/%s/p2p-circuit/%s", c.remoteMaddr, c.remotePeer.Pretty(), c.Conn().RemotePeer()))
	if err != nil {
		panic(err)
	}
	return a
}

func (c *Conn) LocalMultiaddr() ma.Multiaddr {
	return c.Conn().LocalMultiaddr()
}

func (c *Conn) LocalAddr() net.Addr {
	na, err := manet.ToNetAddr(c.Conn().LocalMultiaddr())
	if err != nil {
		log.Error("failed to convert local multiaddr to net addr:", err)
		return nil
	}
	return na
}

func (c *Conn) Transport() tpt.Transport {
	panic("does anyone really call this?")
}

func (c *Conn) LocalPeer() peer.ID {
	return c.Conn().LocalPeer()
}

func (c *Conn) RemotePeer() peer.ID {
	return c.remotePeer
}

func (c *Conn) LocalPrivateKey() ic.PrivKey {
	return nil
}

func (c *Conn) RemotePublicKey() ic.PubKey {
	return nil
}

func (c *Conn) ID() string {
	return "TODO: relay conn ID"
}
