
# Password Generator CLI

This is a simple command-line tool written in Go for generating secure random passwords. You can customize the password length and choose whether to include uppercase letters, numbers, and symbols. Additionally, the generated password can be copied directly to your system clipboard.

## Features

- Generate random passwords of a specified length.
- Include/exclude uppercase letters, numbers, and symbols.
- Option to copy the generated password to the system clipboard.

## Requirements

- Go 1.16 or higher.
- github.com/atotto/clipboard for clipboard functionality.

## Installation

1. Clone the repository:

```bash
git  clone  https://github.com/renancavalcantercb/password-generator-cli.git
```

  

2. Change to the project directory:

```bash
cd  password-generator-cli
```

  

3. Install the dependencies:

  

```bash
go  mod  tidy
```

  

4. Build the project:

```bash
go  build  -o  password-generator-cli
```

## Usage

Run the command with different flags to generate a password:

```bash
./password-generator-cli [options]
```

The available options are:
`--length`: Set the length of the password (default: `12`).
`--uppercase`: Include uppercase letters in the password (default: `true`).
`--numbers`: Include numbers in the password (default: `true`).
`--symbols`: Include symbols in the password (default: `true`).
`--copy`: Copy the generated password to the system clipboard (default: `false`).

### Example

Generate a 16-character password with uppercase letters, numbers, and symbols, and copy it to the clipboard:

```bash
./password-generator-cli  --length=16  --uppercase=true  --numbers=true  --symbols=true  --copy=true
```

### Output

```bash
Generated  Password:  aB3!@f#4G8%qL9kP
Password  copied  to  clipboard!
```
