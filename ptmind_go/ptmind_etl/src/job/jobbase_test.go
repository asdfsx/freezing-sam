package job

import (
	"container/heap"
	"fmt"
	"strconv"
	"testing"
)

func TestLen(t *testing.T) {
	queue := NewJobqueue(10)
	if queue.Len() != 10 {
		t.Log("queue length should be 10")
		t.Fail()
	}
}

func TestPush(t *testing.T) {
	queue := NewJobqueue(0)
	for i := 0; i < 5; i++ {
		job := new(Job)
		job.Jobid = strconv.Itoa(i)
		heap.Push(&queue, job)
	}

	if queue.Len() != 5 {
		t.Log("queue length should be 5")
		t.Fail()
	}
}

func TestPop(t *testing.T) {
	queue := NewJobqueue(0)
	for i := 0; i < 5; i++ {
		job := new(Job)
		job.Priority = 0
		job.Jobid = strconv.Itoa(i)
		heap.Push(&queue, job)
	}

	for len(queue) > 0 {
		item := heap.Pop(&queue)
		fmt.Printf("%v\n", item.Jobid)
	}

	if queue.Len() != 0 {
		t.Log("queue length should be 0")
		t.Fail()
	}
}

func TestNewJobProcessor(t *testing.T) {
	processor := NewJobProcessor("test", 1, 100)
	if processor == nil {
		t.Log("queue length should be 0")
		t.Fail()
	} else {
		fmt.Printf("processor.Jobs_waiting: %v\n", len(processor.Jobs_waiting))
		fmt.Printf("processor: %v, %v\n", processor.Datehour, processor.Timestamp)
	}
}
