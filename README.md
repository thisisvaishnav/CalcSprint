```markdown
# 🧠 Real-Time Two-Player Math Game – Backend

This is the **backend service** for a real-time two-player quiz game built with **Go**, using the **Chi router**, **PostgreSQL**, and **Redis**. It handles player authentication, room management, game logic, and real-time communication via WebSockets.

> ⚠️ This project is currently **under development**.

---

## 🚀 Tech Stack

- **Go** + [Chi Router](https://github.com/go-chi/chi)
- **Supabase** (PostgreSQL) – For user & game data storage
- **Upstash Redis** – For real-time state and caching
- **WebSockets** – For live player interactions

---

## 📁 Project Structure

```

/backend
┣━ main.go                  # Entry point
┣━ routes/
┃ ┣━ auth.go                # Login/Signup logic
┃ ┣━ room.go                # Room creation & join
┃ ┗━ game.go                # Game control and scoring
┣━ sockets/
┃ ┗━ game\_socket.go         # Real-time game events
┣━ db/
┃ ┗━ supabase.go            # Supabase connection
┗━ utils/
┗━ redis.go               # Redis helper functions

````

---

## 📌 Features (Planned)

- ✅ Player Signup/Login
- ✅ Create/Join Game Room
- ✅ Real-time Gameplay with WebSockets
- ✅ Score Tracking and Leaderboard

---

## 🛠️ Setup

```bash
# Clone the repo
git clone https://github.com/your-username/backend.git

cd backend

# Install dependencies
go mod tidy

# Run the server
go run main.go
````

> Ensure that your `.env` file contains correct credentials for Supabase and Redis before running.

---

## ✅ TODOs

* [ ] Add JWT-based auth
* [ ] Implement Redis-based room state cache
* [ ] Write integration tests
* [ ] Rate-limit endpoints

---

## 🧠 Frontend Repository

👉 [Next.js Frontend Repo](https://github.com/your-username/frontend) (WIP)

---

## 📄 License

This project is licensed under the **MIT License**.

```

Let me know if you want to add badges (build passing, license, etc.) or a contribution section!
```
