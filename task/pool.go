package task

import "sync"

func NewTaskPool(max int , handler func(*TaskPool))  *TaskPool {
	return &TaskPool{
		pool:make(chan  bool , max),
		wg:new(sync.WaitGroup),
		handler:handler,
	}
}
type TaskPool struct {
	Max int
	pool chan bool
	wg *sync.WaitGroup
	handler func(*TaskPool)
}

func (this *TaskPool)Add(i int)  {
	this.pool <- true
	this.wg.Add(i)
}

func (this *TaskPool)Done()  {
	<-this.pool
	this.wg.Done()
}


func (this *TaskPool)Run()  {
	this.handler(this)
	this.wg.Wait()
	close(this.pool)
}
