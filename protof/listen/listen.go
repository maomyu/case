package listen

import "net"

type RetryListener struct {
	net.Listener
	addr    string
	network string
}

func RetryListen(network, addr string) (*RetryListener, error) {
	nl, err := net.Listen(network, addr)
	if err != nil {
		return nil, err
	}
	l := &RetryListener{
		Listener: nl,
		network:  network,
		addr:     addr,
	}
	return l, nil
}