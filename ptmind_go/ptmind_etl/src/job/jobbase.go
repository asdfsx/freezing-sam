package job

type Job struct {
	Jobid       string
	Jobname     string
	Jobvalue    string
	Jobstatus   string
	Currenttime int
	Index       int
	priority    int
}

type Jobqueue []*Job

func (jq Jobqueue) Len() int { return len(jq) }

func (jq Jobqueue) Less(i, j int) bool {
	return jq[i].Jobid < jq[j].Jobid && jq[i].priority < jq[j].priority
}

func (jq Jobqueue) Swap(i, j int) {
	jq[i], jq[j] = jq[j], jq[i]
	jq[i].Index = i
	jq[j].Index = j
}

func (jq *Jobqueue) Push(x interface{}) {
	item, ok := x.(*Job)
	if ok {
		*jq = append(*jq, item)
	}
}

func (jq *Jobqueue) Pop() interface{} {
	a := *jq
	n := len(a)
	job := a[n-1]
	job.Index = -1 // for safety
	*jq = a[0 : n-1]
	return jq
}
