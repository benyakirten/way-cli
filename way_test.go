package main

import (
	"os"
	"path"
	"testing"
)

func TestParseFlags(t *testing.T) {
	os.Args = append(os.Args, "-l")
	os.Args = append(os.Args, "5")
	os.Args = append(os.Args, "-w")
	os.Args = append(os.Args, "test/dir")
	os.Args = append(os.Args, "-r")
	os.Args = append(os.Args, "test_folder")

	l, w, r, _ := parseFlags()

	if l != 5 {
		t.Error("length argument should equal 5 after setting -l flag to 5")
	}

	if w != "test/dir" {
		t.Error("way argument should equal test/dir after setting -w flag to test/dir")
	}

	if r != true {
		t.Error("relative argument should equal true after setting the -r flag")
	}
}

func TestCollectResults(t *testing.T) {
	cwd, err := os.Getwd()
	if err != nil {
		t.Fatal("unable to get cwd")
	}
	if os.MkdirAll(path.Join(cwd, "test/folder1/folder2/folder3/folder4"), os.ModePerm) != nil {
		t.Fatal("unable to generate folder for test")
	}
	if os.MkdirAll(path.Join(cwd, "test/folder1/folder1/folder2/folder3/folder4"), os.ModePerm) != nil {
		t.Fatal("unable to generate folder for test")
	}
	if os.MkdirAll(path.Join(cwd, "test/folder1/folder1/folder1/folder2/folder3/folder4"), os.ModePerm) != nil {
		t.Fatal("unable to generate folder for test")
	}
	if os.MkdirAll(path.Join(cwd, "test/folder1/folder1/folder1/folder1/folder2/folder3/folder4"), os.ModePerm) != nil {
		t.Fatal("unable to generate folder for test")
	}

	res1 := collectResults("folder4", -1, cwd, false)
	if len(res1) != 4 {
		t.Error("not enough folders found in default conditions")
	}
	if res1[0] != path.Join(cwd, "test", "folder1", "folder1", "folder1", "folder1", "folder2", "folder3", "folder4") ||
		res1[2] != path.Join(cwd, "test", "folder1", "folder1", "folder2", "folder3", "folder4") ||
		res1[1] != path.Join(cwd, "test", "folder1", "folder1", "folder1", "folder2", "folder3", "folder4") ||
		res1[3] != path.Join(cwd, "test", "folder1", "folder2", "folder3", "folder4") {
		t.Error("unable to locate correct folder structure in default conditions")
	}

	res2 := collectResults("folder1", 2, cwd, false)
	if len(res2) != 2 {
		t.Error("incorrect amount of folders found when result quantity is limited")
	}
	if res2[0] != path.Join(cwd, "test", "folder1") || res2[1] != path.Join(cwd, "test", "folder1", "folder1") {
		t.Error("unable to locate correct folder structure when result quantity is limited")
	}

	res3 := collectResults("folder1", 2, cwd, true)
	if len(res3) != 2 {
		t.Error("incorrect amount of relative folders found")
	}
	if res3[0] != path.Join("test", "folder1") || res3[1] != path.Join("test", "folder1", "folder1") {
		t.Error("unable to identify correct relative folder structure")
	}

	res4 := collectResults("folder2", -1, path.Join(cwd, "test", "folder1", "folder1"), false)
	if len(res4) != 3 {
		t.Error("incorrect amount of folders found with custom path")
	}
	if res4[0] != path.Join(cwd, "test", "folder1", "folder1", "folder1", "folder1", "folder2") ||
		res4[1] != path.Join(cwd, "test", "folder1", "folder1", "folder1", "folder2") ||
		res4[2] != path.Join(cwd, "test", "folder1", "folder1", "folder2") {
		t.Error("unable to identify correct folder structure with custom path")
	}

	os.RemoveAll(path.Join(cwd, "test"))
}

func TestGoUp(t *testing.T) {
	p := path.Join("test1", "test2", "test3", "test4")
	res1 := goUp(p, 2)
	if res1 != path.Join("test1", "test2") {
		t.Error("incorrect result from straightforward test")
	}

	res2 := goUp(p, -5)
	if res2 != p {
		t.Error("incorrect handling of values less than 1")
	}
}
