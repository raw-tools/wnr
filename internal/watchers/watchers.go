package watchers

import "github.com/noirbizarre/wnr/internal/events"

type Watcher interface {
	Watch(<-chan struct{}) (<-chan events.Event, error)
}

func Watch(pattern string) (<-chan events.Event, error) {

	watcher := NewFileWatcher(pattern)
	return watcher.Watch(nil)
}
