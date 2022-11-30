package cleaner

import (
	"context"
	"time"

	"github.com/casnerano/go-url-shortener/internal/app/repository"
)

type Cleaner struct {
	store repository.Store
}

func New(s repository.Store) *Cleaner {
	return &Cleaner{store: s}
}

func (cln *Cleaner) CleanOlderShortURL(ttl int) {
	d := time.Second * time.Duration(ttl)
	rep := cln.store.URL()

	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for {
		<-ticker.C
		_ = rep.DeleteOlderRows(context.TODO(), d)
	}
}
