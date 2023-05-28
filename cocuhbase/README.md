# Couchbase

A Couchbase storage driver using [couchbase/gocb](https://github.com/couchbase/gocb).

### Table of Contents
- [Signatures](#signatures)
- [Installation](#installation)
- [Examples](#examples)
- [Config](#config)
- [Default Config](#default-config)

### Signatures
```go
func New(config ...Config) Storage
func (s *Storage) Get(key string) ([]byte, error)
func (s *Storage) Set(key string, val []byte, exp time.Duration) error
func (s *Storage) Delete(key string) error
func (s *Storage) Reset() error
func (s *Storage) Close() error
```
### Installation
Couchbase is tested on the 2 last [Go versions](https://golang.org/dl/) with support for modules. So make sure to initialize one first if you didn't do that yet:
```bash
go mod init github.com/<user>/<repo>
```
And then install the couchbase implementation:
```bash
go get github.com/gofiber/storage/couchbase
```

### Examples
Import the storage package.
```go
import "github.com/gofiber/storage/couchbase"
```

You can use the following possibilities to create a storage:
```go
// Initialize default config
store := couchbase.New()

// Initialize couchbase storage with custom config
store := couchbase.New(couchbase.Config{
	Host:      "127.0.0.1:8091",
	Username:  "",
	Password:  "",
	Bucket: 0,
	ConnecitonTimeout: 3*time.Second,
	KVTimeout: 1*time.Second,
})
```

### Config
```go
type Config struct {
    // The application username to Connect to the couchbase cluster.
    Username string
    // The application password to Connect to the couchbase cluster.
    Password string
    // The connection string for the couchbase cluster.
    Host string
    // The bucket to use for the couchbase cluster.
    Bucket string
    // The timeout for connecting to the couchbase cluster.
    ConnectionTimeout time.Duration
    // The timeout for KV operations.
    KVTimeout time.Duration
}
```

### Default Config
```go
var ConfigDefault = Config{
    Username:          "admin",
    Password:          "123456",
    Host:              "127.0.0.1:8091",
    Bucket:            "fiber_storage",
    ConnectionTimeout: 3 * time.Second,
    KVTimeout:         1 * time.Second,
}
```
