package job

import (
	"container/heap"
	"fmt"
	"strconv"
	"testing"
)

func TestLen(t *testing.T) {
	queue := make(Jobqueue, 10)
	if queue.Len() != 10 {
		t.Log("queue length should be 10")
		t.Fail()
	}

}

func TestPush(t *testing.T) {
	queue := make(Jobqueue, 0)

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
	queue := make(Jobqueue, 0)

	for i := 0; i < 5; i++ {
		job := new(Job)
		job.Jobid = strconv.Itoa(i)
		heap.Push(&queue, job)
	}

	for i := len(queue); i >= 0; i-- {
		item := heap.Pop(&queue).(*Job)
		fmt.Printf("%v\n", item.Jobid)
	}

	if queue.Len() != 0 {
		t.Log("queue length should be 0")
		t.Fail()
	}
}
