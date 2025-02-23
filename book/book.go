package book

import (
	"sync"
	"sync/atomic"
)

type Book struct {
	Title       string
	CurrentPage int32
	lock        sync.RWMutex
}

func newBook(title string) *Book {
	return &Book{Title: title}
}

// OnPage returns the current page of the book

func (b *Book) OnPage() int32 {
	return b.CurrentPage
}

func (b *Book) OnPageAtomic() int32 {
	return atomic.LoadInt32(&b.CurrentPage)
}

func (b *Book) OnPageLock() int32 {
	b.lock.RLock()
	defer b.lock.RUnlock()
	return b.CurrentPage
}

// TurnPage will increment the CurrentPage by 1

func (b *Book) TurnPage() {
	b.CurrentPage++
}

func (b *Book) TurnPageAtomic() {
	atomic.AddInt32(&b.CurrentPage, 1)
}

func (b *Book) TurnPageLock() {
	b.lock.Lock()
	defer b.lock.Unlock()
	b.CurrentPage++
}

// SkipFirstPage will increment the CurrentPage by 2 if the CurrentPage is 0

func (b *Book) SkipFirstPage() {
	if b.CurrentPage == 0 {
		atomic.AddInt32(&b.CurrentPage, 2)
	}
}

func (b *Book) SkipFirstPageAtomic() {
	if atomic.LoadInt32(&b.CurrentPage) == 0 {
		atomic.AddInt32(&b.CurrentPage, 2)
	}
}

func (b *Book) SkipFirstPageLock() {
	b.lock.Lock()
	defer b.lock.Unlock()
	if b.CurrentPage == 0 {
		b.CurrentPage += 2
	}
}
