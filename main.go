package main

import "fmt"

func toKey(i, j int) string {
  return fmt.Sprintf("%d,%d", i, j)
}

func scapeIsland (i, j int, grid [][]byte, seen map[string]bool) {
  if i < 0 || i >= len(grid) {
    return
  }

  if j >= len(grid[i]) || j < 0 {
    return 
  }

  if grid[i][j] == '0' {
    return
  }

  ky := toKey(i, j)

  if _,saw := seen[ky]; saw {
    return
  }

  seen[ky] = true

  scapeIsland(i-1, j, grid, seen)
  scapeIsland(i+1, j, grid, seen)
  scapeIsland(i,j+1, grid, seen)
  scapeIsland(i,j-1, grid, seen)
}

func numIslands(grid [][]byte) int {
  seen := make(map[string] bool, len(grid)*len(grid[0]))
  var numOfIslands int

  for i := 0; i < len(grid); i++ {
    for j := 0; j < len(grid[i]); j++ {
      ky := toKey(i, j)

      if _, saw := seen[ky]; !saw && grid[i][j] != '0' {
        numOfIslands++
        scapeIsland(i, j, grid, seen)
      }
    }
  }

  return numOfIslands
}

func main() {
  fmt.Println("Hello World")
}