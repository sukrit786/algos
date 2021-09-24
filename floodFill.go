func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	color := image[sr][sc]
	s := []int{}
	s = append(s, sr, sc)
	image[sr][sc] = newColor
	m, n := len(image), len(image[0])

	for len(s) > 0 {

		x := s[0]
		y := s[1]
		fmt.Println(x, " ", y, " ", s)
		s = s[2:]

		if x-1 >= 0 && image[x-1][y] == color && image[x-1][y] != image[x][y] {
			s = append(s, x-1, y)
			image[x-1][y] = newColor
		}
		if y-1 >= 0 && image[x][y-1] == color && image[x][y-1] != image[x][y] {
			s = append(s, x, y-1)
			image[x][y-1] = newColor
		}
		if x+1 < m && image[x+1][y] == color && image[x+1][y] != image[x][y] {
			s = append(s, x+1, y)
			image[x+1][y] = newColor
		}
		if y+1 < n && image[x][y+1] == color && image[x][y+1] != image[x][y] {
			s = append(s, x, y+1)
			image[x][y+1] = newColor
		}
	}
	return image
}