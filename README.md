# Cellabrate

> A portmanteau of "cell" and "celebrate", reflecting the celebration of life’s patterns in the simulation.

Cellabrate is a terminal-based Conway's Game of Life editor and simulator. It allows you to interactively create, edit, and simulate cellular automaton patterns in real-time. Inspired by Vim for its keyboard navigation, it provides a streamlined, efficient way to experiment with patterns, using paint and erase modes, dynamic grid resizing, and more.

> TODO: Cellabrate Demo

## Features

- Real-time cellular automaton simulation
- Interactive cell editing with cursor
- Paint and erase modes for efficient pattern creation
- Vim-inspired keyboard navigation
- Dynamic grid resizing based on terminal dimensions

## How It Works

This implementation of Conway's Game of Life follows these rules:

1. **Underpopulation**: Any live cell with fewer than two live neighbors dies.
2. **Survival**: Any live cell with two or three live neighbors lives on.
3. **Overpopulation**: Any live cell with more than three live neighbors dies.
4. **Reproduction**: Any dead cell with exactly three live neighbors becomes alive.

## Installation

### Clone This Repository

```bash
git clone https://github.com/manankarnik/term-life.git
```

### Navigate to the Project Directory

```bash
cd term-life
```

### Run the Project

```bash
go run term-life
```

## Requirements

- [Golang](https://go.dev/)
- [tcell/v2](https://github.com/gdamore/tcell)

## Usage

> TODO: Cellabrate Interface

### Basic Controls

The program starts with a simple pattern (3-cell line) in the center of the screen. Use the keyboard to navigate, edit cells, and control simulation.

### Modes

The program has two main modes:

1. **Edit Mode (default)** - Move cursor and toggle cells on/off
   - **Paint Mode** - Draw live cells by moving the cursor
   - **Erase Mode** - Erase live cells by moving the cursor
2. **Simulation Mode** - Watch the Game of Life evolve

## Key Bindings

### Movement

| Key                   | Action                              |
| --------------------- | ----------------------------------- |
| h,j,k,l or Arrow keys | Move cursor (left, down, up, right) |
| f                     | Move cursor 5 steps right           |
| b                     | Move cursor 5 steps left            |
| u                     | Move cursor 5 steps up              |
| d                     | Move cursor 5 steps down            |
| Ctrl+F                | Move cursor to far right            |
| Ctrl+B                | Move cursor to far left             |
| Ctrl+U                | Move cursor to top                  |
| Ctrl+D                | Move cursor to bottom               |
| m                     | Move cursor to horizontal middle    |
| M                     | Move cursor to vertical middle      |

### Cell Manipulation

| Key   | Action                                   |
| ----- | ---------------------------------------- |
| Space | Toggle cell state (alive/dead) at cursor |
| p     | Toggle paint mode                        |
| e     | Toggle erase mode                        |
| c     | Clear the entire grid                    |

### Simulation Control

| Key | Action                                 |
| --- | -------------------------------------- |
| s   | Toggle simulation (start/pause)        |
| n   | Step simulation forward one generation |

### Program Control

| Key    | Action           |
| ------ | ---------------- |
| Esc, q | Quit the program |

## Special Modes

### Paint Mode

Paint mode allows you to "draw" live cells by simply moving the cursor. This makes creating patterns much faster than toggling individual cells.

- Press `p` to enter or exit paint mode
- The cursor will turn green in paint mode
- Move the cursor to create live cells
- Paint mode is disabled during simulation
- **Note**: Paint mode is disabled during simulation, and you cannot use it simultaneously with erase mode.

> TODO: Paint Mode Example

### Erase Mode

Erase mode is similar to paint mode but removes cells instead of creating them.

- Press `e` to enter or exit erase mode
- The cursor will turn red in erase mode
- Move the cursor to remove live cells
- Erase mode is disabled during simulation
- **Note**: Erase mode is disabled during simulation, and you cannot use it simultaneously with paint mode.

> TODO: Erase Mode Example

## Examples

Here are some classic patterns you can create:

### Glider

> TODO: Glider Pattern

### Blinker

> TODO: Blinker Pattern

### Gosper Glider Gun

> TODO: Gosper Glider Gun

## Contributing

Contributions are welcome! Please feel free to raise Issues and submit a Pull Request.

## License

This project is licensed under the [MIT License](LICENSE).

## References

- [Conway's Game of Life](https://en.wikipedia.org/wiki/Conway%27s_Game_of_Life) for creating the Game of Life
- [tcell](https://github.com/gdamore/tcell) library for terminal graphics
