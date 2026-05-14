# 🎨 Art Interface

A web-based interface for the **art-decoder** project built with **Go**.  
This project allows users to both **ENCODE** and **DECODE** text-based artwork directly from a browser instead of using the command line.

---

# ✨ Features

- ✅ Encode ASCII-art into shorter code
- ✅ Decode the encoded art back to art form
- ✅ Separate pages for:
  - Homepage
  - Encoder
  - Decoder
- ✅ HTML templates for each page
- ✅ CSS styling for a cleaner user interface
- ✅ Proper HTTP status handling
- ✅ Error handling for malformed data
- ✅ Server-side processing (not client-side)
- ✅ Configurable server port

---

# 📂 Project Structure

```bash
.
├── main.go
├── handlers.go
├── logic.go
├── README.md
│
├── ui/
│   ├── html/
│   │   ├── home.html
│   │   ├── encoder.html
│   │   └── decoder.html
│   │
│   └── static/
│       ├── css/
│       │   ├── home.css
│       │   ├── encoder.css
│       │   └── decoder.css
│       │
│       └── images/
│           ├── greenhomeicon.png
│           ├── yellowhomeicon.png
│           └── terminalscreen.jpg
```

---

# 🌐 HTTP Endpoints

| Method | Endpoint     | Description |
|--------|--------------|-------------|
| GET    | `/`          | Homepage |
| GET    | `/encoder`   | Encoder page |
| POST   | `/encoder`   | Encode submitted ASCII-art |
| GET    | `/decoder`   | Decoder page |
| POST   | `/decoder`   | Decode submitted ASCII-art |

---

# ✅ HTTP Status Codes

The server returns proper HTTP response codes:

| Code | Meaning |
|------|----------|
| `200 OK` | Successful request |
| `202 Accepted` | Successfully processed decoding/encoding request |
| `400 Bad Request` | Malformed or invalid input |
| `404 Not Found` | Invalid route |

---

# 🚀 Running the Project

## 1. Clone the repository

```bash
git clone https://gitea.kood.tech/georgiisenotrusov/interface
cd interface
```

---

## 2. Run the server

```bash
go run .
```

or

```bash
go run main.go
```

---

# 🔌 Server Port

The application runs on:

```bash
localhost:8084
```

You may change the port if needed for testing or review purposes.

Example:

```go
http.ListenAndServe(":8084", nil)
```

---

# 🖥️ Usage

## Homepage

Open:

```bash
http://localhost:8084/
```

From the homepage you can navigate to:

- Encoder page
- Decoder page

---

## Encoder

Navigate to:

```bash
http://localhost:8084/encoder
```

Enter ASCII-art, recieve encoded version.

---

## Decoder

Navigate to:

```bash
http://localhost:8084/decoder
```

Enter encoded piece and decode it back into ASCII-art.

---

# 🎨 Styling

The project includes custom CSS styling to improve usability and visual appearance.

Features include:

- Styled forms
- Buttons
- Responsive layout

---

# ⚠️ Error Handling

The server validates user input and handles malformed data properly.

Examples:

- Invalid encoded format
- Empty requests
- Unsupported characters
- Bad query strings

When invalid data is detected, the server responds with:

```http
400 Bad Request
```

---

# 🧠 Technologies Used

- Go
- HTML Templates
- CSS
- net/http package

---

# 📚 Learning Objectives

This project demonstrates understanding of:

- HTTP servers in Go
- HTML templates
- Routing
- Form handling
- Server-side encoding/decoding
- Error handling
- CSS styling

---

# 🛠️ Additional Notes

- The encoder functionality is fully processed on the server side.
- Templates are separated for better organization and maintainability.
- The application can easily be extended with additional features.

---