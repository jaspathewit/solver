package main

import (
    "fmt"
    "log"
    "solver/core/solver/worker"
    "solver/core/util"
    "solver/sudoku/puzzle"
    "time"
)

func main() {
    defer util.LogDuration(time.Now(), "sudoku")
    p, err := puzzelLibelle()
    if err != nil {
        log.Fatalf("failed to create puzzel: %s", err)
    }

    // create the solver for the suduku
    s := puzzle.SudokuSolver{}
    // start the worker.Solve with the starting sudoku and the solver
    result, err := worker.Solve(p, s)
    if err != nil {
        log.Fatalf("error: %s", err)
    }

    // get the concrete type of solution
    solution := result.(*puzzle.Puzzle)
    fmt.Printf("Solution is\n: %s", solution)
}

// puzzelLibelle
func puzzelLibelle() (*puzzle.Puzzle, error) {

    topology := puzzle.Normal{}

    g, err := puzzle.NewPuzzle(topology)
    if err != nil {
        return nil, err
    }

    // set up the starting values
    g.Set("1_2", "5")
    g.Set("1_3", "2")
    g.Set("1_5", "6")
    g.Set("1_6", "8")
    g.Set("1_8", "3")
    g.Set("2_2", "7")
    g.Set("2_5", "5")
    g.Set("2_7", "9")
    g.Set("2_8", "2")
    g.Set("3_2", "3")
    g.Set("3_5", "1")
    g.Set("3_9", "6")
    g.Set("5_3", "4")
    g.Set("5_6", "5")
    g.Set("5_7", "6")
    g.Set("6_3", "8")
    g.Set("6_5", "4")
    g.Set("6_7", "2")
    g.Set("7_1", "1")
    g.Set("7_2", "9")
    g.Set("7_6", "2")
    g.Set("7_8", "7")
    g.Set("8_6", "6")
    g.Set("8_9", "2")
    g.Set("9_5", "8")
    g.Set("9_7", "1")

    return g, nil
}
