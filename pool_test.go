package bufferpool

import (
	"testing"
)

func TestBufferInRightCase(t *testing.T) {
	pool := NewBufferPool(5)
	b := pool.Take()

	if len(b) != 1024 {
		t.Error("Buffer size should be 1024")
		t.Fail()
	}

	if pool.Return(b) != 1 {
		t.Error("Buffer should return to pool")
		t.Fail()
	}
}

func TestBufferLength(t *testing.T) {
	pool := NewBufferPool(2, 10)
	b := pool.Take()
	if len(b) != 10 {
		t.Error("Buffer size should be 10")
		t.Fail()
	}

	if pool.Return(b) != 1 {
		t.Error("Buffer should return to pool")
		t.Fail()
	}
}

func TestBufferPoolSize(t *testing.T) {
	pool := NewBufferPool(2, 10)

	tmp := [5][]byte{}

	for i := 0; i < 5; i++ {
		tmp[i] = pool.Take()
	}
	for i := 0; i < 5; i++ {
		rs := pool.Return(tmp[i])
		if i > 1 && rs != 0 {
			t.Log(i, rs)
			t.Fail()
		}
	}

}
