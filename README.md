# (IN PROGRESS) 🚀 go-start

A CLI tool to scaffold clean, production-ready Go backend projects using your choice of:

- 🔧 Framework: `Echo`, `Gin`, or `Fiber`
- 🖃️ Database: `GORM` or raw SQL
- 📦 Logger: `Zap`, `Zerolog`, or `Logrus`
- 🔐 Built-in: JWT middleware, Docker setup, and modular folder structure

---

## 📦 Features

- Interactive CLI wizard (`init`) like `create-next-app`
- Clean architecture structure
- Supports PostgreSQL and MySQL (via GORM/raw)
- JWT authentication boilerplate
- Dockerfile and Docker Compose included
- Auto-generated project tree
- Shell completions (`bash`, `zsh`, `fish`, `powershell`)

---

## 📥 Installation

```bash
git clone https://github.com/your-org/go-start.git
cd go-start
go mod tidy
```

---

## 🚀 Usage

### 🧙 Create a New App (Recommended)

```bash
go run main.go init
```

You’ll be prompted to select:
- Project name
- Framework (Echo/Gin/Fiber)
- Database (GORM/raw SQL)
- Logger (Zap/Zerolog/Logrus)

💡 You’ll get a beautiful summary with your app’s structure printed after generation.

---

### ⚙️ Scaffold via Flags (Advanced)

```bash
go run main.go scaffold \
  --name=myapp \
  --framework=gin \
  --db=gorm \
  --log=zerolog
```

---

### 🧩 Enable Shell Completions

```bash
go run main.go completion bash        # or zsh/fish/powershell
```

---

## 🧱 Project Structure (Generated Example)

```
myapp/
├── cmd/
├── internal/
│   ├── config/
│   ├── handler/
│   ├── model/
│   ├── repository/
│   └── service/
├── go.mod
├── Dockerfile
├── docker-compose.yml
```

---

## 🛠 Requirements

- Go 1.20+
- Docker (for running with Compose)

---

## 📌 Todo

- [x] Add CLI wizard
- [x] Add logger/db templates
- [x] Add tree printer
- [ ] Add tests
- [ ] Add plugin system?

---

## 📄 License

MIT © 2025 Jeremy