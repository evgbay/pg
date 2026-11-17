
# PG - Password Generator

A lightweight command-line password generator written in Go.

Generate random passwords directly from your terminal.

## Requirements

- Go 1.26.5+

## Installation

Clone the repository:

```bash
git clone https://github.com/evgbay/pg.git
cd pg
```

Run the application:

```bash
go run .
```

## Build (Windows)

Build the executable:

```bash
go build -o pg.exe .
```

Alternatively, build the executable in the `bin` directory:

```bash
go build -o bin/pg.exe .
```

## Environment Setup (Windows)

To run `pg` from any directory, add the executable's directory to your system PATH.

1. Open **System Properties**.
2. Navigate to **Environment Variables**.
3. Select the **Path** variable and click **Edit**.
4. Add the path to the directory containing `pg.exe`:

```text
C:\path\to\pg\bin
```

5. Save the changes and restart your terminal.

## Usage

### Generate a Password

Run the following command to generate a password with the default length of 16 characters:

```bash
pg
```

### Custom Password Length

Use the `-l` flag to specify the password length:

```bash
pg -l 8
```

Supported password length: 5–64 characters.

### Include Special Characters

Use the `-s` flag to include special characters:

```bash
pg -s
```

The default password length is 16 characters.

### Combine Flags

Specify a custom length and enable special characters:

```bash
pg -l 8 -s
```

## Command-Line Options

| Flag | Description | Default |
|------|-------------|---------|
| `-l` | Password length (5–64 characters) | `16`    |
| `-s` | Include special characters        | `false` |

## Examples

```bash
# Generate a 16-character password
pg

# Generate an 8-character password
pg -l 8

# Generate a 16-character password with special characters
pg -s

# Generate an 8-character password with special characters
pg -l 8 -s
```