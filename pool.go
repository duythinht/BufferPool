package bufferpool

type BufferPool struct {
	Buffers    chan []byte
	BufferSize int
	PoolSize   int
}

// NewBufferPool create a new pool
// Take 2 args
// first arg is max size of buffer
// 2nd arg is default buffer length
func NewBufferPool(args ...int) *BufferPool {
	size := 5  // default max size
	ln := 1024 // default buffer lenght
	if len(args) > 0 {
		size = args[0]
	}

	if len(args) > 1 {
		ln = args[1]
	}
	return &BufferPool{
		Buffers:    make(chan []byte, size),
		BufferSize: ln,
		PoolSize:   size,
	}
}

// Take a buffer in pool, create a new once when pool is empty
func (p *BufferPool) Take() []byte {
	select {
	case b := <-p.Buffers:
		return b
	default:
		return make([]byte, p.BufferSize)
	}
}

// Return buffer to the pool, stash if pool is full
func (p *BufferPool) Return(b []byte) int {
	select {
	case p.Buffers <- b:
		return 1
	default:
		b = b[:0] //enforcing b to be zero, eligible for GC
		return 0
	}
}
