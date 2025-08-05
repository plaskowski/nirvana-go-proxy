# Nirvana Go Proxy

This project provides a simple Go-based reverse proxy that forwards requests to [https://focus.nirvanahq.com](https://focus.nirvanahq.com). Each instance runs on a different local port, giving you **isolated browser storage (localStorage, session, cookies)** per tab in Chrome.

## ✅ Use Case

Use this to log into **multiple NirvanaHQ accounts simultaneously** in **the same Chrome window** by visiting:

- `http://localhost:3001` for Account A  
- `http://localhost:3002` for Account B  
- etc.

Each instance has its own origin → isolated session.

---

## 🔧 Usage

### 1. Build or run directly:

```bash
go run main.go --port=3001
go run main.go --port=3002
