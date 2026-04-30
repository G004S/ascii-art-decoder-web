# Art Decoder 🎨

A command-line tool for encoding and decoding text-based art.

## Overview

**Art Decoder** is a Go-based CLI utility designed to simplify the creation and manipulation of ASCII/text-based art. It allows users to:

- **Decode** compact encoded strings into full text art
- **Encode** repetitive text art into a compressed format
- Work with **single-line input**, **multi-line input**, or **.txt files**

This tool was built to help artists like Chris quickly generate large, repetitive text images without manually typing thousands of characters.

---

## Features

### ✅ Decoder
Expand encoded patterns into full text art.

Example:
    Input: ABC[10 D]EFG
    Output: ABCDDDDDDDDDDEFG

    Input: [5 #][5 -]-[5 #]
    Output: #####------#####


---

### ✅ Encoder
Compress repeated sequences into encoded format.

Example:
    Input: &&&&&
    Output: [5 &]


---

### ✅ Multi-line Mode
Supports encoding and decoding across multiple lines interactively.

---

### ✅ File Support
Read encoded or raw art from `.txt` files and output results directly to the terminal.

---

## Installation & Setup

Make sure you have Go installed (version 1.18+ recommended).

Clone the repository:
    git clone https://github.com/your-username/art-decoder.git

    cd art-decoder

    Run the program: go run .

---

## Usage

### 🔹 Terminal Mode (Single Input)

Use `-t` to indicate terminal input.

#### Decode
go run . "[5 &]" -t -d

Output: &&&&&


#### Encode

go run . "&&&&&" -t -e

Output: [5 &]

---

### 🔹 Multi-line Mode

Start interactive mode using `--start`.

#### Decode

go run . --start -d


#### Encode

go run . --start -e

Then input lines one by one:

    [5 #]
    [3 @]
    end


The program stops reading when you type: 
    end


---

### 🔹 File Mode

Use `-f` to process `.txt` files.

#### Decode file

go run . cats.encoded.txt -f -d


#### Encode file

go run . cats.art.txt -f -e


Output is printed to the terminal.

---

## Encoding Format

Encoded segments follow this structure:
    [number string]


- `number` = how many times to repeat
- `string` = sequence to repeat (can include spaces and symbols)

Example:
[5 -]
→ -----


---

## Error Handling

The program prints:
Error


If any of the following occur:

- First argument inside brackets is not a number
- Missing space between arguments
- Missing or empty string to repeat
- Unbalanced square brackets

---

## Project Structure

.
├── main.go
├── io.go
├── logic.go
└── library/



---

## Key Features

- 🔁 Full bidirectional encoding/decoding
- 📄 File-based processing
- 📚 Multi-line interactive mode
- ⚙️ Flexible CLI flags
