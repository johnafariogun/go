package main

import (
	"fmt"
	"bookstore"
)

func main() {
	// Create a new inventory
	inv := bookstore.Inventory{}

	// --- 1️⃣ Add books ---
	fmt.Println("📚 Adding books to inventory...")
	inv.AddBook(bookstore.Book{
		Title:  "Things Fall Apart",
		Author: "Chinua Achebe",
		Copies: 2,
	})
	inv.AddBook(bookstore.Book{
		Title:  "Half of a Yellow Sun",
		Author: "Chimamanda Ngozi Adichie",
		Copies: 5,
	})
	inv.AddBook(bookstore.Book{
		Title:  "The Last Don",
		Author: "Mario Puzzo",
		Copies: 3,
	})

	printInventory(inv)

	// --- 2️⃣ Find a book ---
	fmt.Println("\n🔍 Finding 'Things Fall Apart'...")
	book := inv.FindBookByTitle("Things Fall Apart")
	if book != nil {
		fmt.Printf("✅ Found: %+v\n", *book)
	} else {
		fmt.Println("❌ Book not found")
	}

	// --- 3️⃣ Buy a book ---
	fmt.Println("\n🛒 Buying 'Things Fall Apart'...")
	if err := inv.BuyBook("Things Fall Apart"); err != nil {
		fmt.Println("❌", err)
	} else {
		fmt.Println("✅ Purchase successful!")
	}

	printInventory(inv)

	// --- 4️⃣ Try buying an unavailable book ---
	fmt.Println("\n🛒 Buying 'Purple Hibiscus'...")
	if err := inv.BuyBook("Purple Hibiscus"); err != nil {
		fmt.Println("❌", err)
	}

	// --- 5️⃣ Restock an existing book ---
	fmt.Println("\n📦 Restocking 'Things Fall Apart' by 3...")
	if err := inv.StockBook("Things Fall Apart", 3); err != nil {
		fmt.Println("❌", err)
	} else {
		fmt.Println("✅ Restock successful!")
	}

	printInventory(inv)

	// --- 6️⃣ Remove a book ---
	fmt.Println("\n🗑 Removing 'The Last Don' from inventory...")
	if err := inv.RemoveBook("The Last Don"); err != nil {
		fmt.Println("❌", err)
	} else {
		fmt.Println("✅ Removed successfully!")
	}

	printInventory(inv)

	// --- 7️⃣ Try removing a non-existent book ---
	fmt.Println("\n🗑 Removing 'Nonexistent Book'...")
	if err := inv.RemoveBook("Nonexistent Book"); err != nil {
		fmt.Println("❌", err)
	}
}

// Helper function to print current inventory
func printInventory(inv bookstore.Inventory) {
	fmt.Println("\n📦 Current Inventory:")
	for _, b := range inv.ListBooks() {
		fmt.Printf("  - %s by %s (copies: %d)\n", b.Title, b.Author, b.Copies)
	}
}
