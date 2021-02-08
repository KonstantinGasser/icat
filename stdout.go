package main

import "io"

// Copy copies content from more then one reader to a given source
// multiple readers required for displaying image in iTerm
// example: var(
//			main.iTermCmdStartRender
//			main.iTermCmdStopRender
//		)d
func Copy(dst io.Writer, src ...io.Reader) (int64, error) {
	return io.Copy(dst, io.MultiReader(src...))
}
