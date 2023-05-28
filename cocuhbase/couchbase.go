package cocuhbase

import (
	"github.com/couchbase/gocb/v2"
	"time"
)

type Storage struct {
	cb     *gocb.Cluster
	bucket *gocb.Bucket
}

func New(config ...Config) *Storage {
	cfg := configDefault(config...)
	cb, err := gocb.Connect(cfg.Host, gocb.ClusterOptions{
		Authenticator: gocb.Authenticator(&gocb.PasswordAuthenticator{
			Username: cfg.Username,
			Password: cfg.Password,
		}),
		TimeoutsConfig: gocb.TimeoutsConfig{
			KVTimeout:      cfg.ConnectionTimeout,
			ConnectTimeout: cfg.KVTimeout,
		},
		Transcoder: gocb.NewLegacyTranscoder(),
	})

	if err != nil {
		panic(err)
	}

	_, err = cb.Ping(&gocb.PingOptions{
		Timeout: 5 * time.Second,
	})

	if err != nil {
		panic(err)
	}

	b := cb.Bucket(cfg.Bucket)
	return &Storage{cb: cb, bucket: b}
}

func (s *Storage) Get(key string) ([]byte, error) {
	out, err := s.bucket.
		DefaultCollection().
		Get(key, nil)
	if err != nil {
		return nil, err
	}

	var value []byte
	if err := out.Content(&value); err != nil {
		return nil, err
	}
	return value, nil
}

func (s *Storage) Set(key string, val []byte, exp time.Duration) error {
	if _, err := s.bucket.
		DefaultCollection().
		Upsert(key, val, nil); err != nil {
		return err
	}
	return nil
}

func (s *Storage) Delete(key string) error {
	if _, err := s.bucket.DefaultCollection().Remove(key, nil); err != nil {
		return err
	}

	return nil
}

func (s *Storage) Reset() error {
	return s.cb.Buckets().FlushBucket(s.bucket.Name(), nil)
}

func (s *Storage) Close() error {
	return s.cb.Close(nil)
}
