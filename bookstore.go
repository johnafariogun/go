package bookstore

import "errors"

// Book represents a book in the store.
type Book struct {
	Title  string
	Author string
	Copies int
}

type Inventory struct {
	Inventory []Book
}

func (s *Inventory) AddBook(b Book) {
	s.Inventory = append(s.Inventory, b)
}

func (s *Inventory) ListBooks() []Book {
	return s.Inventory
}

func (s *Inventory) TotalBooks() int {
	return len(s.Inventory)
}


func (s *Inventory) FindBookByTitle(title string) *Book {
	for i, b := range s.Inventory {
		if b.Title == title {
			return &s.Inventory[i]
		}
	}
	return nil
}


func (s *Inventory) RemoveBook(title string) error {
	book := s.FindBookByTitle(title)
	if book == nil {
		return errors.New("book not found")
	}

	for i, b := range s.Inventory {
		if b.Title == title {
			s.Inventory = append(s.Inventory[:i], s.Inventory[i+1:]...)
			break
		}
	}
	return nil
}


func (s *Inventory) BuyBook(title string) error {
	book := s.FindBookByTitle(title)
	if book == nil {
		return errors.New("book not found")
	}

	if book.Copies == 0 {
		return errors.New("book is out of stock")
	}

	book.Copies--
	return nil
}


func (s *Inventory) StockBook(title string, copies int) error {
	book := s.FindBookByTitle(title)
	if book == nil {
		return errors.New("book not found")
	}

	book.Copies += copies
	return nil
}
