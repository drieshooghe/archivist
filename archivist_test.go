package archivist

import (
	"testing"

	filet "github.com/Flaque/filet"
)

func TestScribe(t *testing.T) {
	defer filet.CleanUp(t)

	// Mock a directory with nested files and directories
	parentDir := filet.TmpDir(t, "")
	childDir1 := filet.TmpDir(t, parentDir)
	childDir2 := filet.TmpDir(t, parentDir)

	filet.File(t, parentDir+"/tf1.txt", "ozewiezewoze")
	filet.File(t, childDir1+"/tf2.txt", "wiezewalla")
	filet.File(t, childDir2+"/tf3.txt", "kristalla")

	if filet.DirContains(t, parentDir, childDir1) != true {
		t.Errorf("Could not create directory %s in %s", childDir1, parentDir)
	}
	if filet.DirContains(t, parentDir, childDir2) != true {
		t.Errorf("Could not create directory %s in %s", childDir2, parentDir)
	}
	if filet.DirContains(t, parentDir, "tf1.txt") != true {
		t.Errorf("Could not create file %s in %s", "tf1.txt", parentDir)
	}
	if filet.DirContains(t, childDir1, "tf2.txt") != true {
		t.Errorf("Could not create file %s in %s", "tf2.txt", childDir1)
	}
	if filet.DirContains(t, childDir2, "tf3.txt") != true {
		t.Errorf("Could not create file %s in %s", "tf3.txt", childDir2)
	}

}

func TestZipCompressor(t *testing.T) {

}
