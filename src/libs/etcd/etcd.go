package etcd

import (
	"time"

	"github.com/coreos/etcd/clientv3"
)

// ETCDConnection dial etcd connection
// @broker: string
// @timeOut: time.Duration
// return *clientv3.Client, error
func ETCDConnection(broker []string, timeOut time.Duration) (*clientv3.Client, error) {
	cli, err := clientv3.New(clientv3.Config{
		DialTimeout: timeOut,
		Endpoints:   broker,
	})
	if err != nil {
		return nil, err
	}
	return cli, nil
}
