package model

import (
	"bytes"
	"io"
	"time"

	"github.com/cxbooks/cxbooks/server/zlog"
	"github.com/xujiajun/nutsdb"
)

type CacheEntryIdxMode int

const (
	//索引和值都存储在内存中
	CacheHintKeyValAndRAMIdxMode = iota

	//索引存储内存，值不存储内存
	CacheHintKeyAndRAMIdxMode

	//多级索引存储内存，值不存储内存
	CacheHintBPTSparseIdxMode
)

type KV struct {
	client *nutsdb.DB
	bucket string
}

func OpenKV(path string) (*KV, error) {
	op := nutsdb.DefaultOptions
	op.Dir = path
	op.EntryIdxMode = nutsdb.EntryIdxMode(CacheHintKeyAndRAMIdxMode)

	db, err := nutsdb.Open(op)
	if err != nil {
		return nil, err
	}

	return &KV{
		client: db,
		bucket: "_inner_default_bucket",
	}, nil
}

// SetBytes 更新nutsdb数据
func (m *KV) Set(key string, data []byte, t time.Duration) error {

	return m.SetBytes(m.bucket, key, data, t)
}

// SetBytes 更新nutsdb数据
func (m *KV) Get(key string) ([]byte, error) {
	return m.GetBytes(m.bucket, key)
}

func (m *KV) Write(key string, reader io.Reader, t time.Duration) error {

	return m.WriteBytes(m.bucket, key, reader, t)

}

// Del 删除对象
func (m *KV) Del(key string) error {
	return m.DelBytes(m.bucket, key)
}

// Get 从nutsdb读取数据
func (c *KV) GetBytes(bucket, key string) ([]byte, error) {

	var res []byte

	err := c.client.View(func(tx *nutsdb.Tx) error {
		entry, err := tx.Get(bucket, []byte(key))
		if err != nil {
			return err
		}
		res = entry.Value
		return nil
	})

	return res, err
}

// SetBytes 更新nutsdb数据
func (m *KV) SetBytes(bucket, key string, data []byte, t time.Duration) error {

	if err := m.client.Update(
		func(tx *nutsdb.Tx) error {

			return tx.Put(bucket, []byte(key), data, uint32(t/time.Second))
		}); err != nil {
		zlog.E("SetBytes error: ", err.Error())
		return err
	}

	return nil
}

func (m *KV) WriteBytes(bucket, key string, reader io.Reader, t time.Duration) error {

	buf := new(bytes.Buffer)
	buf.ReadFrom(reader)
	data := buf.Bytes()

	return m.SetBytes(bucket, key, data, t)

}

// Del 删除对象
func (m *KV) DelBytes(bucket, key string) error {
	if err := m.client.Update(
		func(tx *nutsdb.Tx) error {

			return tx.Delete(bucket, []byte(key))
		}); err != nil {
		zlog.E("SetBytes error: ", err.Error())
		return err
	}
	return nil
}

// Close 关闭连接
func (c *KV) Close() {
	c.client.Close()
}
