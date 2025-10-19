package bookstore_test

import (
	"bookstore"
	"testing"
)

func TestBookInitialization(t *testing.T) {
	t.Parallel()
	book := bookstore.Book{
		Title:  "The Last Don",
		Author: "Mario Puzzo",
		Copies: 100,
	}
	if book.Title == "" || book.Author == "" {
		t.Error("book not initialized properly")
	}
}

func TestAddAndListBooks(t *testing.T) {
	t.Parallel()
	inv := bookstore.Inventory{}
	book := bookstore.Book{Title: "Splinter Cell", Author: "Tom Clancy", Copies: 10}

	inv.AddBook(book)
	books := inv.ListBooks()

	if len(books) != 1 {
		t.Fatalf("expected 1 book in inventory, got %d", len(books))
	}
	if books[0].Title != "Splinter Cell" {
		t.Errorf("expected 'Splinter Cell', got %s", books[0].Title)
	}
}

func TestBuyBook(t *testing.T) {
	t.Parallel()
	inv := bookstore.Inventory{}
	inv.AddBook(bookstore.Book{Title: "Splinter Cell", Author: "Tom Clancy", Copies: 2})

	err := inv.BuyBook("Splinter Cell")
	if err != nil {
		t.Fatalf("unexpected error buying book: %v", err)
	}

	book := inv.FindBookByTitle("Splinter Cell")
	if book.Copies != 1 {
		t.Errorf("expected 1 copy left, got %d", book.Copies)
	}
}

func TestBuyBook_OutOfStock(t *testing.T) {
	t.Parallel()
	inv := bookstore.Inventory{}
	inv.AddBook(bookstore.Book{Title: "Splinter Cell", Author: "Tom Clancy", Copies: 0})

	err := inv.BuyBook("Splinter Cell")
	if err == nil {
		t.Error("expected error for out-of-stock book, got nil")
	}
}

func TestStockBook(t *testing.T) {
	t.Parallel()
	inv := bookstore.Inventory{}
	inv.AddBook(bookstore.Book{Title: "Splinter Cell", Author: "Tom Clancy", Copies: 1})

	err := inv.StockBook("Splinter Cell", 3)
	if err != nil {
		t.Fatalf("unexpected error restocking: %v", err)
	}

	book := inv.FindBookByTitle("Splinter Cell")
	if book.Copies != 4 {
		t.Errorf("expected 4 copies, got %d", book.Copies)
	}
}

func TestRemoveBook(t *testing.T) {
	t.Parallel()
	inv := bookstore.Inventory{}
	inv.AddBook(bookstore.Book{Title: "Splinter Cell", Author: "Tom Clancy", Copies: 5})

	err := inv.RemoveBook("Splinter Cell")
	if err != nil {
		t.Fatalf("unexpected error removing book: %v", err)
	}

	if len(inv.ListBooks()) != 0 {
		t.Errorf("expected inventory to be empty after removal, got %d", len(inv.ListBooks()))
	}
}
