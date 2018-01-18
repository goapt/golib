# NSQPool

NSQPool is a thread safe connection pool for nsq producer. It can be used to
manage and reuse nsq producer connection.


## Install and Usage

Install the package with:

```bash
github.com/qgymje/nsqpool
```

Import it with:

```go
import (
    "github.com/qgymje/nsqpool"
    nsq "github.com/nsqio/go-nsq"
)
```

and use `pool` as the package name inside the code.

## Example

```go
// create a factory() to be used with channel based pool
factory := func() (*nsq.Producer, error) { 
    config := nsq.NewConfig()
    return nsq.NewProducer(":4150", config)
}

nsqPool, err := pool.NewChannelPool(5, 30, factory)

producer, err := nsqPool.Get()

producer.Publish("topic", "some data")
// do something with producer and put it back to the pool by closing the connection
// (this doesn't close the underlying connection instead it's putting it back
// to the pool).
producer.Close()

// close pool any time you want, this closes all the connections inside a pool
nsqPool.Close()

// currently available connections in the pool
current := nsqPool.Len()
```

## License

The MIT License (MIT) - see LICENSE for more details
