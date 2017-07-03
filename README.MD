### GO effective BufferPool

```
import "github.com/duythinht/bufferpool"


Pool := NewBufferPool(10)

buffer := Pool.Take()
defer func() {
  _ = Pool.Return(buffer)
}()

// Do stuff with buffer

```
