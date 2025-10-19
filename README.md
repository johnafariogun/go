# 📚 Bookstore

A simple Go project that models a bookstore inventory system.  
It demonstrates working with structs, slices, methods, and basic testing.

---

## 🚀 Features

- Add and list books in inventory  
- Find books by title  
- Buy books (decrease stock)  
- Restock books (increase stock)  
- Remove books from inventory  
- Unit tests for all main functionalities

---

## 🗂️ Project Structure

```

bookstore/
├── bookstore.go          # Core logic (Book + Inventory structs)
├── bookstore_test.go     # Unit tests
├── cmd/
│   └── bookstore/
│       └── main.go       # Example usage of the bookstore package
└── go.mod                # Go module definition

````

---

## ⚙️ Setup

1. **Clone this repository**
   ```bash
   git clone https://github.com/yourusername/bookstore.git
   cd bookstore
````

2. **Initialize or tidy Go modules**

   ```bash
   go mod tidy
   ```

---

## ▶️ Running the Application

From the project root, run:

```bash
go run ./cmd/bookstore
```

This will execute `main.go` and demonstrate:

* Adding books
* Finding, buying, and restocking books
* Removing books from inventory

---

## 🧪 Running Tests

To verify that all features work as expected, run:

```bash
go test -v ./...
```

---

## 🧠 Concepts Demonstrated

* Struct composition (`Book`, `Inventory`)
* Slice manipulation (adding, removing elements)
* Pointer receivers in methods
* Error handling
* Go unit testing

---

## ✨ Example Output

```
📚 Adding books to inventory...

📦 Current Inventory:
  - Things Fall Apart by Chinua Achebe (copies: 2)
  - Half of a Yellow Sun by Chimamanda Ngozi Adichie (copies: 5)
  - The Last Don by Mario Puzzo (copies: 3)

🛒 Buying 'Things Fall Apart'...
✅ Purchase successful!

📦 Current Inventory:
  - Things Fall Apart by Chinua Achebe (copies: 1)
  - Half of a Yellow Sun by Chimamanda Ngozi Adichie (copies: 5)
  - The Last Don by Mario Puzzo (copies: 3)

📦 Restocking 'Things Fall Apart' by 3...
✅ Restock successful!

🗑 Removing 'The Last Don' from inventory...
✅ Removed successfully!
```

---

## 🧩 Next Steps

* Add CLI interactivity (user input for actions)
* Save inventory data to a JSON or database file
* Build a REST API with **Go + FastAPI style routes**