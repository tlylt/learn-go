package book

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestNewBook(t *testing.T) {
	book := newBook("Atomic Habits")
	if book.Title != "Atomic Habits" {
		panic("Expected title to be 'Atomic Habits'")
	}
}

func TestRaceConditionBookNoAtomicNoLock(t *testing.T) {
	t.Skip("This test is expected to fail")
	book := newBook("Atomic Habits")

	wg := sync.WaitGroup{}
	wg.Add(1)

	go func() {
		defer wg.Done()
		book.TurnPage()
	}()

	fmt.Printf("Current page: %d\n", book.OnPage())
	wg.Wait()

	if book.OnPage() != 1 {
		t.Error("Expected page to be 1, Actual page is ", book.OnPage())
	}
}

func TestNoRaceConditionBookAtomic(t *testing.T) {
	t.Skip("This test is expected to fail")
	book := newBook("Atomic Habits")

	wg := sync.WaitGroup{}
	wg.Add(1)

	go func() {
		defer wg.Done()
		book.TurnPageAtomic()
	}()

	fmt.Printf("Current page: %d\n", book.OnPageAtomic())
	wg.Wait()

	if book.OnPageAtomic() != 1 {
		t.Error("Expected page to be 1, Actual page is ", book.OnPageAtomic())
	}
}

func TestRaceConditionBookWithAtomic(t *testing.T) {
	t.Skip("This test is expected to fail")
	book := newBook("Atomic Habits")

	wg := sync.WaitGroup{}
	wg.Add(1)

	go func() {
		defer wg.Done()
		book.SkipFirstPage()
		fmt.Printf("After SkipFirstPage: page: %d\n", book.OnPageAtomic())
	}()

	time.Sleep(1 * time.Microsecond)
	book.TurnPageAtomic()
	fmt.Printf("After TurnPageAtomic: page: %d\n", book.OnPageAtomic())
	fmt.Println(book.OnPageAtomic())

	wg.Wait()
	fmt.Printf("End of test: page: %d\n", book.OnPageAtomic())
	if book.OnPageAtomic() != 1 {
		t.Error("Expected page to be 1, Actual page is ", book.OnPageAtomic())
	}
}

func TestRaceConditionSkipWithLock(t *testing.T) {
	t.Skip("This test is expected to fail")
	book := newBook("Atomic Habits")

	wg := sync.WaitGroup{}
	wg.Add(1)

	go func() {
		defer wg.Done()
		book.SkipFirstPageLock()
		fmt.Printf("After SkipFirstLock: page: %d\n", book.OnPageLock())
	}()

	time.Sleep(1 * time.Microsecond)
	book.TurnPageLock()
	fmt.Printf("After TurnPageLock: page: %d\n", book.OnPageLock())
	fmt.Println(book.OnPageLock())

	wg.Wait()
	fmt.Printf("End of test: page: %d\n", book.OnPageLock())
	if book.OnPageLock() != 1 {
		t.Error("Expected page to be 1, Actual page is ", book.OnPageLock())
	}
}
