<p align="center">
  <img src="https://readme-typing-svg.demolab.com?font=Fira+Code&size=28&duration=3000&pause=500&color=00F7F7&center=true&vCenter=true&width=435&lines=ELECTONE+%E2%9C%A8;Real-time+voting+system" alt="Typing SVG" />
</p>

---

<p align="center">
<img src="https://img.shields.io/badge/Backend-Golang-blue?style=for-the-badge&logo=go" />
<img src="https://img.shields.io/badge/Frontend-Svelte-orange?style=for-the-badge&logo=svelte" />
<img src="https://img.shields.io/badge/Database-Turso-9cf?style=for-the-badge" />
<img src="https://img.shields.io/badge/Realtime-Redis-red?style=for-the-badge&logo=redis" />
<img src="https://img.shields.io/badge/Comm-WebSockets-brightgreen?style=for-the-badge" />
</p>

---

## 🗳️ What is Electone?

**Electone** is a full-stack, real-time voting system built for scale. Users can register, view active elections, and cast votes securely. Admins can manage candidates, users, and elections with full CRUD capability. Designed as a collaborative development project, it leverages modern tech like Go, Redis, WebSockets, Turso, and Svelte.

---

## 🧱 Tech Stack

- **Frontend**: Svelte 5 + TailwindCSS + Daisyui
- **Backend**: Go + fiber 
- **Database**: Turso, Redis

**Miscellaneous**: WebSockets, JWT-based, Apache Benchmark (ab)

---

## 🚀 Installation & Setup

### 🔧 Frontend

> Requires `bun`, `node`, and optionally `tailwindcss` (standalone binary)

```bash
git clone https://github.com/PS-Wizard/Electone
cd Electone/electoneui
bun install
bun run dev
```

### 🖥 Backend
> Requires `golang`, `redis`, and an `auth token`
```bash
~

git clone https://github.com/yourusername/Electone.git
cd Electone/backend
go mod tidy

```
> Redis must be running locally, If you're testing without Redis, uncomment lines **32 and 33** in `main.go` to disable voting logic temporarily. 
🧠 Request the authentication token from @PS-Wizard to enable access to the database.


### 🧪 API Collection
[API Collection](./assets/rester-export-postman.json)
[Database Schema](./assets/database.md)

