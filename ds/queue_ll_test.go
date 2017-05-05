package ds

import (
	"testing"
)

func TestQueueLL_Deque(t *testing.T) {
	queue := new(QueueLL)
	queue.Enqueue("A")
	queue.Enqueue("B")

	result := queue.Deque().(string)
	if result != "A" {
		t.Errorf("deque() data incorrect, got %s, expected %s", result, "A")
	}
}
