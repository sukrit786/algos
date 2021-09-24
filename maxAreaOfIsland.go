// func findPath(visited *[][]int,grid [][]int,i int,j int,m int,n int) int {
//     (*visited)[i][j]=1;
//     path:=[]int{i,j};
//     count:=1
//     for len(path)>0 {
//         x:=path[0];
//         y:=path[1];
//         path = path[2:]
//         if(x-1>=0 && grid[x-1][y]==1 && (*visited)[x-1][y]==0) {
//             path = append(path,x-1,y)
//             (*visited)[x-1][y]=1;
//             count++
//         }
//         if(y-1>=0 && grid[x][y-1]==1 && (*visited)[x][y-1]==0) {
//             path = append(path,x,y-1)
//             (*visited)[x][y-1]=1;
//             count++
//         }
//         if(x+1<m && grid[x+1][y]==1 && (*visited)[x+1][y]==0) {
//             path = append(path,x+1,y)
//             (*visited)[x+1][y]=1;
//             count++
//         }
//         if(y+1<n && grid[x][y+1]==1 && (*visited)[x][y+1]==0) {
//             path = append(path,x,y+1)
//             (*visited)[x][y+1]=1;
//             count++
//         }
//     }
//     return count
// }

// func max(a int,b int)int {
//     if(a>b) {
//         return a
//     } else {
//         return b
//     }
// }

// func maxAreaOfIsland(grid [][]int) int {
//     m:=len(grid);
//     n:=len(grid[0]);
//     visited := make([][]int, m)
//     maxArea:=0
//     for i := range visited {
//         visited[i] = make([]int, n)
//     }

//     for i:=0;i<m;i++ {
//         for j:=0;j<n;j++ {
//             if grid[i][j] == 1 && visited[i][j]==0 {
//                 localMax:=findPath(&visited,grid,i,j,m,n)
//                 maxArea = max(localMax,maxArea)
//             }
//         }
//     }

//     fmt.Print(visited,maxArea)

//     return maxArea
//

func maxAreaOfIsland(grid [][]int) int {
	max := 0

	maxRow := len(grid)
	maxCol := len(grid[0])
	if maxRow == 0 || maxCol == 0 {
		return 0
	}

	var dfs func(row, col int) int
	dfs = func(row, col int) int {
		if row < 0 {
			return 0
		} else if row > maxRow-1 {
			return 0
		} else if col < 0 {
			return 0
		} else if col > maxCol-1 {
			return 0
		} else if grid[row][col] != 1 {
			return 0
		}
		grid[row][col] = 2
		return 1 + dfs(row+1, col) +
			dfs(row-1, col) +
			dfs(row, col+1) +
			dfs(row, col-1)

	}

	var count = 0
	for row := range grid {
		for col := range grid[row] {
			if grid[row][col] == 1 {
				count = dfs(row, col)
				if count > max {
					max = count
				}
			}
		}
	}

	return max
}
