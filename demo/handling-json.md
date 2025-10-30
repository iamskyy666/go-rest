```go
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"strings"
)

// serialization: go obj{} (instance of a struct) -> json-string ([]byte)
// deserialization: the opposite
// json.marshal()/json.unmarshal() - for in-memory json processing/ops (ok for small-mid sized dataset)
// json.encoder/decoder - for streaming json data.. working with large data sets and networking connections (more flexible and suited for complex ops.)

type User struct {
	Name string `json:"name"`
	Email string `json:"email"`
}

func main() {
	user:= User{Name: "Skyy", Email: "skyy@email.com"}

	// MARSHAL
	jsonData,err:=json.Marshal(user)
	if err!=nil{
		log.Fatal("⚠️ ERR:",err)
	}
	fmt.Println(user) // {Skyy skyy@email.com}
	fmt.Println(string(jsonData)) // {"name":"Skyy","email":"skyy@email.com"}

	// UNMARSHAL
	var user1 User
	err = json.Unmarshal(jsonData,&user1)
	if err!=nil{
		log.Fatal("⚠️ ERR:",err)
	}
	fmt.Println("User created from json data:",user1) // User created from json data: {Skyy skyy@email.com}

	// ENCODER & DECODER - More common, dealing with APIs
	jsonData1:= `{"name":"Soumadip","email":"soumadip@email.com"}`
	reader:= strings.NewReader(jsonData1)
	decoder:= json.NewDecoder(reader)
	var user2 User
	err = decoder.Decode(&user2)
	if err!=nil{
		log.Fatal("⚠️ ERR:",err)
	}
	fmt.Println(user2) // {Soumadip soumadip@email.com}

	
	var buff bytes.Buffer
	encoder := json.NewEncoder(&buff) // encoder needs buffer

	err=encoder.Encode(user)
	if err!=nil{
		log.Fatal("⚠️ ERR:",err)
	}
	fmt.Println("Encoded json-string:",buff.String()) // Encoded json-string: {"name":"Skyy","email":"skyy@email.com"}
}
```
---

# Walkthrough of the code (high-level)

We defined a simple struct and exercised both the marshal/unmarshal and encoder/decoder APIs:

```go
type User struct {
    Name  string `json:"name"`
    Email string `json:"email"`
}
```

`main()`:

* Create `user := User{...}`
* `json.Marshal(user)` → produce `[]byte` JSON (serialization)
* `json.Unmarshal(jsonData, &user1)` → populate `user1` from JSON (deserialization)
* Use `json.NewDecoder(reader)` to decode from an `io.Reader` (streaming)
* Use `json.NewEncoder(&buff)` to write JSON into an `io.Writer` (`bytes.Buffer`) (streaming/encoding)

Now we’ll unpack every concept and each relevant line.

---

# Core concepts

### Serialization vs Deserialization

* **Serialization** (a.k.a. *marshalling*): convert an in-memory Go value (struct, map, slice, etc.) into JSON text (a `[]byte` or string). In Go JSON library: `json.Marshal(v)` or `json.Encoder`.
* **Deserialization** (a.k.a. *unmarshalling*): parse JSON text and construct/populate Go values. In Go JSON library: `json.Unmarshal(data, &v)` or `json.Decoder`.

Why these terms matter: when talking about APIs we usually serialize responses and deserialize requests.

---

# `encoding/json` primitives

### `json.Marshal(v)`

* Input: a Go value `v`.
* Output: `([]byte, error)` containing UTF-8 JSON text.
* Use when we have a whole object in memory and we want JSON bytes (e.g., to store, log, or send as HTTP body).
* Example in our code:

  ```go
  jsonData, err := json.Marshal(user)
  ```

  `jsonData` becomes `[]byte` with contents `{"name":"Skyy","email":"skyy@email.com"}`.

Notes:

* `Marshal` only exports **exported fields** (fields with capitalized names). Unexported fields are ignored.
* Respect `json` struct tags (e.g., `json:"name"`).
* For readability, use `json.MarshalIndent` to pretty-print.
* `Marshal` allocates a `[]byte` — for large payloads this can use a lot of memory.

---

### `json.Unmarshal(data, &v)`

* Input: `data []byte` and a pointer `&v` where `v` is a Go value to populate.
* `Unmarshal` needs the destination to be addressable (a pointer); it mutates the value via that pointer.
* Example:

  ```go
  var user1 User
  err = json.Unmarshal(jsonData, &user1)
  ```

  `user1` becomes a `User` whose fields are populated.

Common pitfalls:

* Passing a value instead of a pointer — `json.Unmarshal(jsonData, user1)` fails to populate (and will compile but behave incorrectly).
* JSON field names map to struct fields using tags; if names don’t match, fields stay zero-valued.
* Numbers in JSON may default to `float64` if decoding into `interface{}`; use `Decoder.UseNumber()` if we need precise number representation.

---

# `json.Encoder` and `json.Decoder` (streaming)

### When to use them

* Use when working with `io.Reader`/`io.Writer` (HTTP request/response bodies, files, sockets).
* Use for large datasets or continuous streams (avoids building a big `[]byte` in memory).
* `Decoder` can decode a stream containing multiple JSON values sequentially.

### `json.NewDecoder(reader)` + `Decode(&v)`

* `Decode` reads from an `io.Reader` until it has a full JSON value and decodes into `&v`.
* In our code:

  ```go
  jsonData1 := `{"name":"Soumadip","email":"soumadip@email.com"}`
  reader := strings.NewReader(jsonData1)
  decoder := json.NewDecoder(reader)
  var user2 User
  err = decoder.Decode(&user2)
  ```

  Works the same as `Unmarshal` but reads from a reader.

Extra useful features:

* `decoder.DisallowUnknownFields()` — error if JSON contains fields not present on the target struct.
* `decoder.UseNumber()` — decode numbers to `json.Number` instead of `float64`, so we can parse them as `Int64` or `Float64` safely.

### `json.NewEncoder(writer)` + `Encode(v)`

* `Encode` writes JSON to an `io.Writer`.
* It writes the JSON text followed by a newline (`\n`). This newline makes it easy to write newline-delimited JSON to logs or streams.
* In our code:

  ```go
  var buff bytes.Buffer
  encoder := json.NewEncoder(&buff)
  err = encoder.Encode(user)
  fmt.Println("Encoded json-string:", buff.String()) // prints JSON plus newline
  ```

Use `Encoder` when returning JSON to an `http.ResponseWriter`:

```go
w.Header().Set("Content-Type", "application/json")
json.NewEncoder(w).Encode(data)
```

This streams directly to the HTTP response without building large intermediate slices.

---

# Explaining specific types in our code

* `strings.NewReader(jsonData1)` produces an `io.Reader` from a string — perfect for `json.NewDecoder`.
* `bytes.Buffer` implements `io.Writer` and `io.Reader`; it’s handy for collecting encoded output in memory.
* `encoder.Encode(user)` appends a newline; `json.Marshal` does not.

---

# Important details, gotchas, and advanced tips

### Exported fields only

* Only exported fields (capitalized) are encoded/decoded. If we define:

  ```go
  type T struct { secret string } // will not be marshalled
  ```

  `secret` is omitted. Use `Secret string` to export.

### Struct tags

* `json:"name"` — changes key name in JSON.
* `json:"name,omitempty"` — omit the key if the value is zero value (empty string, 0, nil, false).
* `json:"-"` — ignore this field entirely.
* `json:"age,string"` — encode/decode number as JSON string.

### `nil` vs zero values

* `omitempty` depends on zero values. Slices/maps become `nil` or `[]` depending on value; omitted if zero and tagged `omitempty`.

### Decoder with streams

* `Decoder.Decode` reads up to the next JSON value. If we send newline-delimited JSON objects, call `Decode` repeatedly in a loop until `io.EOF`.

Example:

```go
dec := json.NewDecoder(reader)
for {
    var item Item
    if err := dec.Decode(&item); err == io.EOF {
        break
    } else if err != nil {
        // handle error
    }
    // process item
}
```

### Unknown fields

* By default, `Decoder` ignores extra keys in JSON. Use `dec.DisallowUnknownFields()` to make it strict.

### Numeric precision

* If decoding into `interface{}`, default numbers become `float64`.
* Use `dec.UseNumber()` to get `json.Number` and avoid precision surprises, e.g. for 64-bit integers.

### Performance considerations

* `json.Marshal` => allocates a `[]byte` and returns it. For huge outputs, this can be memory-heavy.
* `Encoder` streams to `io.Writer` (better for large data).
* For extreme performance needs, there are third-party libraries like `jsoniter` or `easyjson`, but standard `encoding/json` is sufficient and idiomatic for most cases.

### Concurrency

* `Encoder` and `Decoder` are not goroutine-safe for concurrent use. Use separate encoders/decoders or synchronize access.

### Time values

* `time.Time` marshals to RFC3339 by default. If we need a custom format, implement `MarshalJSON` / `UnmarshalJSON` on the type.

---

# Why `Encode` prints a newline (and what that implies)

`encoder.Encode(v)` appends a `\n` after the JSON object. This is intentional:

* Makes it easy to produce **newline-delimited JSON** (NDJSON), which is popular for logs, streaming, and line-oriented processors.
* If we need **compact** JSON without trailing newline, use `json.Marshal` and write the bytes ourself.

---

# Example pitfalls and how to debug them

1. **Fields not appearing in JSON**

   * Cause: field is unexported or tag mismatch.
   * Fix: capitalize the field and/or add correct `json` tag.

2. **Unmarshal not populating struct**

   * Cause: forgot pointer.
   * Fix: pass `&v` to `json.Unmarshal` or `decoder.Decode`.

3. **Large memory usage**

   * Cause: `json.Marshal` on huge slices.
   * Fix: stream via `Encoder` to `http.ResponseWriter` or file.

4. **Number types converted to float64**

   * Cause: decoding to `interface{}`.
   * Fix: decode into concrete struct fields, or `decoder.UseNumber()`.

5. **Extra JSON fields silently ignored**

   * Fix: use `decoder.DisallowUnknownFields()` to detect typos in keys.

---

# Example additions (useful snippets)

Pretty-printing:

```go
prettyJSON, _ := json.MarshalIndent(user, "", "  ")
fmt.Println(string(prettyJSON))
```

Write JSON to HTTP response:

```go
func handler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(user) // streams to client
}
```

Decode request body into struct:

```go
var payload Payload
dec := json.NewDecoder(r.Body)
dec.DisallowUnknownFields() // optional: be strict
if err := dec.Decode(&payload); err != nil {
    http.Error(w, "bad request", http.StatusBadRequest)
    return
}
```

Decode a stream of objects:

```go
dec := json.NewDecoder(reader)
for {
    var obj SomeType
    if err := dec.Decode(&obj); err == io.EOF {
        break
    } else if err != nil {
        log.Fatal(err)
    }
    // process obj
}
```

---

# Finally — mapping back to our printed outputs

From our code:

* `fmt.Println(user)` prints the Go struct representation: `{Skyy skyy@email.com}`
* `fmt.Println(string(jsonData))` prints the JSON string produced by `json.Marshal`: `{"name":"Skyy","email":"skyy@email.com"}`
* After `json.Unmarshal`, `user1` prints as a populated Go struct.
* Using `json.Decoder` on a `strings.Reader` is functionally the same as `json.Unmarshal` here — but it demonstrates decoding from a `Reader`.
* `encoder.Encode(user)` writes JSON to a buffer and appends a newline — so `buff.String()` will show JSON plus `\n`.

---

