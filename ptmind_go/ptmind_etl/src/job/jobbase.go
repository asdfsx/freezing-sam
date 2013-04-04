package job

import (
	"fmt"
	"sync"
	"time"
)

type Job struct {
	Jobid       string
	Jobname     string
	Jobvalue    string
	Jobstatus   string
	Currenttime int
	Index       int
	Priority    int
	Redotime    int
}

type Jobqueue []*Job

func NewJobqueue(size int) (queue Jobqueue) {
	queue = make(Jobqueue, size)
	return
}

func (jq Jobqueue) Len() int { return len(jq) }

func (jq Jobqueue) Less(i, j int) bool {
	return jq[i].Priority <= jq[j].Priority && jq[i].Jobid > jq[j].Jobid
	//return jq[i].Jobid > jq[j].Jobid
}

func (jq Jobqueue) Swap(i, j int) {
	jq[i], jq[j] = jq[j], jq[i]
	jq[i].Index = i
	jq[j].Index = j
}

func (jq *Jobqueue) Push(x interface{}) {
	item, ok := x.(*Job)
	if ok {
		n := len(*jq)
		item.Index = n
		*jq = append(*jq, item)
	}
}

func (jq *Jobqueue) Pop() interface{} {
	a := *jq
	n := len(a)
	item := a[n-1]
	item.Index = -1 // for safety
	*jq = a[0 : n-1]
	return item
}

type JobProcessor struct {
	Jobs_waiting   Jobqueue
	Jobs_running   Jobqueue
	Jobs_finished  Jobqueue
	Prefix         string
	Maxredotime    int
	Datehour       string
	Timeoutseconds int
	Timestamp      int64
	Lock           sync.Locker
}

func NewJobProcessor(prefix string, maxredotime int, timeoutseconds int) (processor *JobProcessor) {
	processor = new(JobProcessor)
	processor.Prefix = prefix
	processor.Maxredotime = maxredotime
	processor.Timeoutseconds = timeoutseconds
	processor.Lock = new(sync.Mutex)
	return
}

func (jp *JobProcessor) Create_job(datehour string) {
	jp.Lock.Lock()
	defer jp.Lock.Unlock()

	t1, _ := time.Parse("2006010215", datehour)
	jp.Timestamp = t1.Unix()
	for i := 0; i < 3; i++ {
		job_key := fmt.Sprintf("%v_%v", jp.Prefix, i)
		job_value := fmt.Sprintf("%v_%v", jp.Prefix, i)
		job_redo_time := 0
		job := Job{Jobid: job_key, Jobvalue: job_value, Redotime: job_redo_time}
		jp.Jobs_waiting.Push(&job)
	}
}
