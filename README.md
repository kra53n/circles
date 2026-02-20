# Circles — AI Search Algorithm Visualizer

An interactive visualizer for classic and heuristic search algorithms, built in Go using [raylib-go](https://github.com/gen2brain/raylib-go). The puzzle consists of a 4×4 grid of colored circles (4 colors), and the goal is to find the shortest sequence of moves that transforms the current state into the goal state.

## Puzzle Rules

The grid has 4 columns and 4 rows. Each move shifts an entire row to the left or an entire column upward (cyclically). The goal state has each column filled with a single color:

```
R Y G B
R Y G B
R Y G B
R Y G B
```

You can also interact with the grid manually by clicking the arrow buttons that appear around the board.

### Search Algorithms

#### Uninformed Search

Algorithm | Description
--- | ---
Breadth-First Search | Explores nodes level by level; guarantees shortest path
Depth-First Search | Explores as deep as possible first; not guaranteed optimal
Bidirectional Search | Simultaneously searches forward from start and backward from goal; meets in the middle

#### A\* with Heuristics

All A\* variants guarantee an optimal path if the heuristic is admissible.

Heuristic | Description
--- | ---
Misplaced circles | Counts circles not in their goal position, divided by 4
Manhattan distance | Sum of Manhattan distances of each circle to its nearest goal cell of the same color, divided by 4
Subtask (DBnly)** | Uses a precomputed pattern database for one color at a time
Subtask + Manhattan | Combines the pattern database lookup with the Manhattan heuristic
Subtask max | Takes the maximum pattern DB value across all 4 colors (strongest admissible heuristic)

### Pattern Databases (Subtask Heuristic)

The subtask approach precomputes the exact number of moves needed to place all circles of one color into the correct column, ignoring the other colors. This is stored as a lookup table keyed by the positions of the 4 circles of that color.

- File format: Each line encodes a combination of 4 cell positions and the BFS distance to the goal for that color.
- Files: `subtask0.txt`, `subtask1.txt`, `subtask2.txt`, `subtask3.txt`
- Generation: Uses `BidirectionalSearch` for each of the ~1820 unique position combinations.

## Getting Started

### Prerequisites

- Go 1.23+
- C compiler (required by raylib-go — e.g. `gcc` on Linux/macOS, `mingw` on Windows)

### Install & Run

```bash
git clone <repo-url>
cd <repo>
go run .
```

### Generate Subtask Databases (required for subtask heuristics)

Before using heuristics 6–8, generate the pattern database files:

```bash
go run . subtask
```

This creates `subtask0.txt` through `subtask3.txt` — one per color. Generation may take a few minutes.

### Run Benchmarks

```bash
mkdir -p measures
go run . measure
```

Results are written to `measures/` as `.txt` files. Use `count.py` to compute averages:

```bash
python3 count.py
```

## Controls

 Key | Action
--- | ---
<kbd>1</kbd> | Breadth-First Search
<kbd>2</kbd> | Depth-First Search
<kbd>3</kbd> | Bidirectional Search
<kbd>4</kbd> | A\* — Heuristic 1 (misplaced circles)
<kbd>5</kbd> | A\* — Heuristic 2 (Manhattan distance)
<kbd>6</kbd> | A\* — Subtask heuristic (pattern DB only)
<kbd>7</kbd> | A\* — Subtask + Manhattan combined
<kbd>8</kbd> | A\* — Subtask max over all colors
<kbd>Space</kbd> | Pause / Resume animation
<kbd>R</kbd> | Restart animation from beginning
<kbd>C</kbd> | Scramble the board randomly
<kbd>Ctrl</kbd> + <kbd>+</kbd> / <kbd>-</kbd> | Increase / decrease random scramble depth
<kbd>P</kbd> | Print current scramble depth

Arrow buttons around the board let you shift rows and columns manually.

## Project Structure

File | Description
--- | ---
`main.go` | Entry point, window loop, key bindings
`field.go` | Grid rendering, arrow buttons, move logic
`state.go` | State representation, copying, equality, move generation
`search.go` | BFS, DFS, Bidirectional, A\*, all heuristic functions
`statistic.go` | Tracks and prints search statistics (iterations, open/closed nodes)
`animation.go` | Animates the solution path step by step
`subtask.go` | Pattern database generation, serialization, and lookup
`measure.go` | Batch benchmark runner
`count.py` | Python script to average benchmark results
`utils.go` | Helper utilities (rectangle centering)

## Statistics Output

After each search, the console prints:

- **Path length** — number of moves in the solution
- **Iterations** — total nodes dequeued
- **Open nodes** — count at termination and maximum during search
- **Closed nodes** — count at termination and maximum during search
- **Max nodes in memory** — peak total memory usage
