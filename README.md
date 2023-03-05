# chalk-advanced-go

Chalk-advanced-go is a library that allows you to color your console output in go.

[![size](https://img.shields.io/github/repo-size/marchoffmann07/chalk-advanced-go?color=green&label=SIZE)](https://www.npmjs.com/package/chalk-advanced)

## Quick Links

- [Installation](#installation)
    - [Code Example](#code-example)
- [Features](#features)
- [Contributors](#contributors)

## Installation

Install chalk-advanced-go

```bash
  go get github.com/marchoffmann07/chalk-advanced-go@1.3.3
```

## Code Example

```go
package main

import (
    "fmt"
    ChalkAdvanced "github.com/marchoffmann07/chalk-advanced-go"
)

func main() {
    fmt.Println(ChalkAdvanced.Blue("Blue message"))
}
```

## Features
### Normal Colors

- **`Black(text)`**
- **`Red(text)`**
- **`Green(text)`**
- **`Yellow(text)`**
- **`Blue(text)`**
- **`Magenta(text)`**
- **`Cyan(text)`**
- **`White(text)`**
- **`Gray(text)`**

### Bright Colors

- **`RedBright(text)`**
- **`GreenBright(text)`**
- **`YellowBright(text)`**
- **`BlueBright(text)`**
- **`MagentaBright(text)`**
- **`CyanBright(text)`**
- **`WhiteBright(text)`**

### Normal Background Colors
- **`BgBlack(text)`**
- **`BgRed(text)`**
- **`BgGreen(text)`**
- **`BgYellow(text)`**
- **`BgBlue(text)`**
- **`BgMagenta(text)`**
- **`BgCyan(text)`**
- **`BgWhite(text)`**

### Bright Background Colors
- **`BgRedBright(text)`**
- **`BgGreenBright(text)`**
- **`BgYellowBright(text)`**
- **`BgBlueBright(text)`**
- **`BgMagentaBright(text)`**
- **`BgCyanBright(text)`**
- **`BgWhiteBright(text)`**

### Specials
- **`Bold(text)`**
- **`Dim(text)`**
- **`Italic(text)`**
- **`Underline(text)`**
- **`Inverse(text)`**
- **`Hide(text)`**
- **`Strikethrough(text)`**

## Contributors

![image](https://contrib.rocks/image?repo=marchoffmann07/chalk-advanced-go)