package controler

import (
	"sync"

	"github.com/noirbizarre/wnr/internal/config"
	"github.com/noirbizarre/wnr/internal/events"
	"github.com/noirbizarre/wnr/internal/tasks"
	"github.com/noirbizarre/wnr/internal/watchers"
)

type Controler struct {
	cfg *config.Config
	// Predefined tasks
	tasks map[string]tasks.Task
	// Predefined watchers instances
	watchers map[string]watchers.Watcher
	// Tasks to execute
	events chan events.Event
	task   tasks.Task
}

func New(cfg *config.Config, task tasks.Task) (*Controler, error) {
	return &Controler{
		cfg:      cfg,
		tasks:    map[string]tasks.Task{},
		watchers: map[string]watchers.Watcher{},
		events:   make(chan events.Event, 5),
		task:     task,
	}, nil
}

// func (c *Controler) AddWatch(w watchers.Watcher) {

// }

func (c *Controler) Run(exit <-chan struct{}, w ...watchers.Watcher) error {
	chans := []<-chan events.Event{}
	for _, watcher := range w {
		queue, err := watcher.Watch(exit)
		if err != nil {
			return err
		}
		chans = append(chans, queue)
	}

	go c.runTasks(merge(chans...))

	return nil
}

func (c *Controler) runTasks(queue <-chan events.Event) {
	for range queue {
		// fmt.Printf("Task received: %s (%s)\n", evt.Name, evt.KwArgs)
		go c.task.Run()
		// task := tasks.NewCommandTask("")
		// go task.Run()
	}
}

func merge(cs ...<-chan events.Event) <-chan events.Event {
	var wg sync.WaitGroup
	out := make(chan events.Event)

	// Start an output goroutine for each input channel in cs.  output
	// copies values from c to out until c is closed, then calls wg.Done.
	output := func(c <-chan events.Event) {
		for n := range c {
			out <- n
		}
		wg.Done()
	}
	wg.Add(len(cs))
	for _, c := range cs {
		go output(c)
	}

	// Start a goroutine to close out once all the output goroutines are
	// done.  This must start after the wg.Add call.
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}
