package pool

import (
	"sync"

	nsq "github.com/nsqio/go-nsq"
)

// PoolConn is a wrapper around net.Conn to modify the the behavior of
// net.Conn's Close() method.
type PoolConn struct {
	*nsq.Producer
	mu       sync.RWMutex
	c        *channelPool
	unusable bool
}

// Close puts the given connects back to the pool instead of closing it.
func (p *PoolConn) Close() error {
	p.mu.RLock()
	defer p.mu.RUnlock()

	if p.unusable {
		if p.Producer != nil {
			p.Producer.Stop()
			return nil
		}
		return nil
	}
	return p.c.put(p.Producer)
}

// MarkUnusable marks the connection not usable any more, to let the pool close it instead of returning it to pool.
func (p *PoolConn) MarkUnusable() {
	p.mu.Lock()
	p.unusable = true
	p.mu.Unlock()
}

// newConn wraps a standard net.Conn to a poolConn net.Conn.
func (c *channelPool) wrapConn(conn *nsq.Producer) *PoolConn {
	p := &PoolConn{c: c}
	p.Producer = conn
	return p
}
