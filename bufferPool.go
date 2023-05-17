package bufferPool

import (
	"bytes"
	"sync"
)

type BufferPool struct {
	mux     sync.Mutex
	buffers []*bytes.Buffer
}

func New() *BufferPool {
	return &BufferPool{
		buffers: []*bytes.Buffer{},
	}
}

func (bp *BufferPool) Get() (buf *bytes.Buffer) {
	bp.mux.Lock()
	if len(bp.buffers) > 1 {
		buf = bp.buffers[len(bp.buffers)-1]
		bp.buffers = bp.buffers[:len(bp.buffers)-2]
	} else {
		buf = bytes.NewBuffer([]byte{})
		bp.buffers = append(bp.buffers, buf)
	}
	bp.mux.Unlock()
	return buf
}

func (bp *BufferPool) Release(buf *bytes.Buffer) {
	bp.mux.Lock()
	buf.Reset()
	bp.buffers = append(bp.buffers, buf)
	bp.mux.Unlock()
}

func (bp *BufferPool) Len() int {
	bp.mux.Lock()
	l := len(bp.buffers)
	bp.mux.Unlock()
	return l
}
