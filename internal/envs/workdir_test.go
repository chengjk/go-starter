package envs

import "testing"

func TestProjectDir(t *testing.T) {
	dir := ProjectDir()
	print(dir)
}
