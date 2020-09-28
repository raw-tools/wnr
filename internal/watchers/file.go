package watchers

import (
	"fmt"
	"log"
	"path/filepath"
	"strings"
	"time"

	"github.com/bmatcuk/doublestar/v2"
	"github.com/fsnotify/fsnotify"
	"github.com/noirbizarre/wnr/internal/events"
	"github.com/pkg/errors"
)

type FileWatcher struct {
	patterns []string
	exclude  []string
}

func NewFileWatcher(patterns ...string) *FileWatcher {
	return &FileWatcher{
		patterns: patterns,
	}
}

func (fw *FileWatcher) Watch(exit <-chan struct{}) (<-chan events.Event, error) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}

	queue := make(chan events.Event, 1)

	files := map[string]bool{}
	// done := make(chan bool)

	// Aggregate/debuf events
	go func() {
		for range time.After(500 * time.Millisecond) {
			println("afeter")
			if len(files) > 0 {
				out := make([]string, 0, len(files))
				for k := range files {
					out = append(out, k)
				}
				log.Printf("Files changed: %s\n", strings.Join(out, ", "))
				files = map[string]bool{}
				queue <- *events.New("file", events.KwArgs{
					"files": out,
				})
			}
		}
		// for {
		// 	select {
		// 	case <-time.After(500 * time.Millisecond):
		// 		if len(files) > 0 {
		// 			out := make([]string, 0, len(files))
		// 			for k := range files {
		// 				out = append(out, k)
		// 			}
		// 			log.Printf("Files changed: %s\n", strings.Join(out, ", "))
		// 			files = map[string]bool{}
		// 			queue <- *events.New("file", events.KwArgs{
		// 				"files": out,
		// 			})
		// 		}
		// 	}
		// }
	}()

	// Listen for fsnotify Event
	go func() {
		defer watcher.Close()
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				log.Println("event:", event)
				// evt := fmt.Sprintf(`file:%s`, strings.ToLower(event.Op.String()))
				files[event.Name] = true
				// queue <- *events.New(evt, events.KwArgs{
				// 	"files": []string{event.Name},
				// })
				// if event.Op&fsnotify.Create == fsnotify.Create {
				// 	log.Println("created file:", event.Name)
				// 	queue <- *events.New("file:created", events.KwArgs{
				// 		"file": event.Name,
				// 	})
				// }
				// if event.Op&fsnotify.Write == fsnotify.Write {
				// 	log.Println("modified file:", event.Name)
				// 	queue <- *events.New("file:modified", events.KwArgs{
				// 		"file": event.Name,
				// 	})
				// }
				// if event.Op&fsnotify.Remove == fsnotify.Remove {
				// 	log.Println("removed file:", event.Name)
				// 	queue <- *events.New("file:removed", events.KwArgs{
				// 		"file": event.Name,
				// 	})
				// }
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Println("error:", err)
			}
		}
	}()

	for _, pattern := range fw.patterns {
		// fmt.Printf("pattern: %s (abs=%b)\n", pattern, filepath.IsAbs(pattern))

		paths, err := doublestar.Glob(pattern)
		if err != nil {
			return nil, errors.Wrapf(err, `Unable to parse pattern %s`, pattern)
		}

		for _, path := range paths {
			// println(path)
			err = watcher.Add(path)
			if err != nil {
				log.Fatal(err)
				return nil, err
			}
		}
	}

	// <-done
	return queue, nil
}

func (fw *FileWatcher) WatchBak() error {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	done := make(chan bool)
	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				log.Println("event:", event)
				if event.Op&fsnotify.Write == fsnotify.Write {
					log.Println("modified file:", event.Name)
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Println("error:", err)
			}
		}
	}()

	for _, pattern := range fw.patterns {
		fmt.Printf("pattern: %s (abs=%b)\n", pattern, filepath.IsAbs(pattern))

		paths, err := doublestar.Glob(pattern)
		if err != nil {
			return errors.Wrapf(err, `Unable to parse pattern %s`, pattern)
		}

		for _, path := range paths {
			println(path)
			err = watcher.Add(path)
			if err != nil {
				log.Fatal(err)
			}
		}
	}

	<-done
	return nil
}
