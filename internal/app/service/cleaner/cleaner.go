package cleaner

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/casnerano/go-url-shortener/internal/app/repository"
)

// Cleaner structure for clean processing.
type Cleaner struct {
	store repository.Store
}

// New - constructor.
func New(s repository.Store) *Cleaner {
	return &Cleaner{store: s}
}

// CleanOlderShortURL runs a method to remove ShortURL every second.
func (cln *Cleaner) CleanOlderShortURL(ctx context.Context, wg *sync.WaitGroup, ttl int) {
	defer wg.Done()

	d := time.Second * time.Duration(ttl)
	rep := cln.store.URL()

	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			_ = rep.DeleteOlderRows(ctx, d)
		case <-ctx.Done():
			fmt.Println("Cleaner older ShortURL stopped.")
			return
		}
	}
}
