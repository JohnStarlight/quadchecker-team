# Quadchecker

Quadchecker is a command-line tool written in Go that identifies specific ASCII shapes (Quads). It reads a shape from standard input (stdin), analyzes its dimensions, and determines if it matches one of the known patterns: `quadA`, `quadB`, `quadC`, `quadD`, or `quadE`.

## Features

- **Shape Identification**: Detects which Quad algorithm generated the input.
- **Dimension Calculation**: Reports the width (x) and height (y) of the shape.
- **Interactive Confirmation**: Displays the received shape and asks for user confirmation before processing (requires TTY support).
- **Multiple Matches**: Handles cases where a shape might match multiple Quad definitions.

## Prerequisites

- Go (Golang) installed on your machine.


## Usage

Quadchecker is designed to receive input via a pipe `|`. You typically run a Quad generator (e.g., `quadA`) and pipe its output into `quadchecker`.

### Basic Syntax

```bash
./quadA <width> <height> | ./quadchecker
```

### Interactive Flow

When you run the command in a supported terminal, `quadchecker` will:
1. Read the input from the pipe.
2. Display the shape to the terminal.
3. Ask for confirmation: `Do you want to proceed? [Press ENTER to match or 'q' to exit]:`
4. If you press **ENTER**, it proceeds to identify the shape.

*Note: If `/dev/tty` is not available (e.g., certain Windows environments), the interactive step is skipped.*

## Examples

### 1. Single Match

**Command:**
```bash
./quadA 3 3 | ./quadchecker
```

**Output:**
```text
--- RECEIVED SHAPE ---
o-o
| |
o-o

The input forms the shape above. Do you want to proceed? [Press ENTER to match or 'q' to exit]:
✅ Match found: [quadA] [3] [3]
```

### 2. Multiple Matches

Some shapes (especially small ones like 1x1) are identical across different Quads.

**Command:**
```bash
./quadC 1 1 | ./quadchecker
```

**Output:**
```text
--- RECEIVED SHAPE ---
A

The input forms the shape above. Do you want to proceed? [Press ENTER to match or 'q' to exit]:
✅ Matches found: [quadC] [1] [1] || [quadD] [1] [1] || [quadE] [1] [1]
```

## Troubleshooting

- **Permission Denied: If your system blocks you from running the programs, you need to grant them execution rights. Run the command chmod +x quadA quadB quadC quadD quadE quadchecker in your terminal to make all the files executable.
- **System Error**: If you see `❌ System Error`, ensure the input pipe is functioning correctly.
- **Warning: No data provided**: Ensure you are piping output into the command.
- **Format Error**: The input lines must all be the same length to form a valid rectangle.

