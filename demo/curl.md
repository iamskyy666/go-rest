## 🌀 What is **cURL**?

**cURL** stands for **Client URL**.
It’s a **command-line tool** (and also a **library**, called *libcurl*) used to **send HTTP requests** and **interact with APIs, servers, or any web endpoints**.

Think of cURL as a powerful HTTP client that can:

* Send `GET`, `POST`, `PUT`, `DELETE` requests, etc.
* Send form data, JSON, or files.
* Include headers, authentication tokens, and cookies.
* Work with different protocols (HTTP, HTTPS, FTP, etc.).
* Support **HTTP/2** and **HTTP/3**.

---

## 💡 How It Relates to Web Development

When developing backend APIs or web apps, we often test endpoints like:

```bash
curl -X POST http://localhost:8080/register \
     -H "Content-Type: application/json" \
     -d '{"email":"test@example.com","password":"123456"}'
```

This simulates what our frontend or client will do — sending a request to our backend.

---

## ⚙️ cURL in **Golang**

In Go, we **don’t use the `curl` command directly** (that’s a CLI tool).
Instead, we use Go’s built-in **`net/http`** package — which works like a “programmatic cURL.”

However, if we specifically want to make **HTTP/2 requests**, we can use:

* Go’s built-in HTTP client (it already supports HTTP/2 automatically over HTTPS)
* OR explicitly configure an HTTP/2 client.

---

## 🧠 Example 1: Basic HTTP/2 GET Request in Go

```go
package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	// Go automatically upgrades HTTPS to HTTP/2 if the server supports it
	resp, err := http.Get("https://www.google.com")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	fmt.Println("Status:", resp.Status)
	fmt.Println("Protocol:", resp.Proto) // Will show HTTP/2.0 if supported
	fmt.Println(string(body[:200]))      // Print first 200 chars
}
```

✅ **Note:**
If the server supports HTTP/2, Go will automatically use it via TLS (HTTPS).
You can confirm by checking `resp.Proto` → it’ll show `"HTTP/2.0"`.

---

## 🧠 Example 2: Forcing HTTP/2 (Manual Configuration)

Sometimes we explicitly configure the client to use HTTP/2.

```go
package main

import (
	"fmt"
	"io"
	"net/http"
	"golang.org/x/net/http2"
)

func main() {
	client := &http.Client{}
	http2.ConfigureTransport(client.Transport.(*http.Transport))

	req, _ := http.NewRequest("GET", "https://www.google.com", nil)
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	fmt.Println("Protocol:", resp.Proto) // Should print "HTTP/2.0"
	fmt.Println(string(body[:200]))
}
```

---

## 🧰 Behind the Scenes

* `cURL` (command-line) and `http.Client` (Go) are both **HTTP clients**.
* `cURL` uses the C library **libcurl**.
* Go’s `net/http` uses the **http2** and **http.Transport** implementations built into the standard library.

---

## 🧪 Extra Tip: Testing in Terminal

We can verify HTTP/2 support directly using `curl` in our terminal:

```bash
curl -I --http2 https://www.google.com
```

Output will include:

```
HTTP/2 200
content-type: text/html; charset=UTF-8
...
```

---
