# NirvanaHQ Dual Session Proxy (Go)

This project is a minimal Go-based reverse proxy that allows you to run **multiple isolated sessions** of [focus.nirvanahq.com](https://focus.nirvanahq.com) in a **single Chrome window**, each logged into a different account — without requiring incognito, extra Chrome profiles, or paid extensions.

---

## ✅ What This Solves

- Isolates `localStorage` and session state per tab
- Avoids dev/debug behaviors triggered by `.local` or `localhost`
- Mimics real external domains using `/etc/hosts` overrides
- Keeps everything running on your machine with no external dependencies

---

## 🔧 Setup Instructions

### 1. Clone and build

```bash
git clone <this-repo>
cd nirvana_go_proxy
go build -o proxy
```

### 2. Set up fake domains

Edit your `/etc/hosts` file and add:

```
127.0.0.1   nirvana-a.fakeproxy.test
127.0.0.1   nirvana-b.fakeproxy.test
```

These serve as your "isolated tabs".

> 🧠 Each domain creates a new browser origin → separate `localStorage`

---

### 3. Run two proxy instances

```bash
./proxy --port=3001
./proxy --port=3002
```

Each one listens on a different port and forwards requests to `https://focus.nirvanahq.com`.

---

### 4. Optional: Serve fake domains over HTTPS using Caddy

Install [Caddy](https://caddyserver.com/) and use this `Caddyfile`:

```caddyfile
nirvana-a.fakeproxy.test {
    reverse_proxy localhost:3001
}

nirvana-b.fakeproxy.test {
    reverse_proxy localhost:3002
}
```

Then run:

```bash
sudo caddy run
```

Now visit:

- https://nirvana-a.fakeproxy.test ← log in with Account A
- https://nirvana-b.fakeproxy.test ← log in with Account B

Each tab maintains its own session.

---

## 🧼 Optional: Silence Google Analytics Errors

You may see a failing request to:

```
https://www.googletagmanager.com/gtag/js?id=G-GTRT9Q5J9H
```

It’s harmless. But you can block it by:

### Option 1 — `/etc/hosts` block:
```
127.0.0.1 www.googletagmanager.com
```

### Option 2 — Browser extension (e.g. uBlock Origin)

### Option 3 — Add custom route to proxy (advanced)

---

## 💡 Notes

- `auth_token` is stored in `localStorage`, not cookies
- Manual login works perfectly in each tab
- Sessions are persistent per fake domain

---

## 🛠 Requirements

- Go 1.21+
- Chrome or Chromium-based browser
- Optional: Caddy for local HTTPS

---

## 🙏 Credits

Inspired by the need to use NirvanaHQ with multiple accounts without relying on profiles, extensions, or incognito workarounds.
