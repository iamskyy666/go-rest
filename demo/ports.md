
## 🌍 What is a Port?

A **port** is a **virtual communication endpoint** that allows our computer to send and receive data through the network.

* Think of our **computer as a large building** (the machine itself).
* Each **port is like a numbered door** inside that building.
* Data from the internet enters through our **IP address** (the building’s address), and then the operating system decides **which port (door)** that data should go to.

---

## 🧠 Technical Definition

A **port** is an **integer between 0 and 65535** that identifies a specific process or service on a computer.

Every network connection uses a combination of:

```
<IP address> + <Port number>
```

This pair is called a **socket**.

---

## ⚙️ Example in Web Development

When we start a backend server, say in Go or Node.js:

```bash
app.Listen(":4000")
```

or

```js
app.listen(4000)
```

We’re telling the operating system:

> “Hey OS, please open **door number 4000** (port 4000) so that when someone tries to connect to our computer on that door, send the traffic to me (the web server process).”

---

## 🌐 Ports and HTTP

| Port  | Protocol   | Common Use                            |
| ----- | ---------- | ------------------------------------- |
| 80    | HTTP       | Default for regular web traffic       |
| 443   | HTTPS      | Default for secure web traffic        |
| 3000  | Custom     | Common in Node.js dev servers         |
| 4000  | Custom     | Often used for Go/Express dev servers |
| 27017 | MongoDB    | Default MongoDB database port         |
| 5432  | PostgreSQL | Default PostgreSQL port               |

When we type:

```
http://localhost:4000/
```

It means:

> Send a network request to **our own computer** (localhost → 127.0.0.1)
> on **port 4000** (door number 4000), where our Go Fiber server is listening.

---

## 💬 Multiple Servers, Different Ports

We can have multiple servers running **at the same time**, as long as they each use **different ports**.

Example:

* React frontend → `localhost:5173`
* Go Fiber backend → `localhost:4000`
* MongoDB database → `localhost:27017`

Each one is listening on a **different port** — no conflicts.

If two apps try to listen on the **same port**, we get:

```
listen tcp :4000: bind: address already in use
```

That’s what happened in your earlier Fiber error — another process was already using port 4000.

---

## 🧱 PORT Binding

When we write:

```go
app.Listen(":4000")
```

* The OS “binds” our Go process to **port 4000**.
* Now any incoming traffic to that port will be delivered to our app.
* If another app already “owns” that port, binding fails → `address already in use`.

---

## 🔐 Reserved Port Ranges

* **0–1023** → Well-known system ports (need admin/root access to bind)

  * Example: 22 (SSH), 25 (SMTP), 80 (HTTP)
* **1024–49151** → Registered ports (used by apps)
* **49152–65535** → Dynamic/private ports (used temporarily by the OS)

That’s why we typically use ports like **3000**, **4000**, **8080**, etc. for local web servers — they’re free and non-restricted.

---

## 🧭 Summary

| Concept              | Description                                                       |
| -------------------- | ----------------------------------------------------------------- |
| **Port**             | A numbered endpoint for communication on a device                 |
| **Purpose**          | Distinguish between multiple apps/services using the same network |
| **Range**            | 0 – 65535                                                         |
| **Common Dev Ports** | 3000, 4000, 5000, 8080                                            |
| **Conflict Error**   | Happens when two programs use the same port                       |
| **Fix**              | Stop the previous process or change the port number               |

---