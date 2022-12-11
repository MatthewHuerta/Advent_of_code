package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Directory struct {
	Name    string
	Parent  *Directory
	Subdirs []*Directory
	Files   []*File
	Size    int64
}

func NewDirectory(dirname string) *Directory {
	var newdir Directory
	newdir.Name = dirname
	return &newdir
}

// takes in name of a child directory and returns a pointer to it
func (d *Directory) Child(name string) (*Directory, error) {
	for _, c := range d.Subdirs {
		if c.Name == name {
			return c, nil
		}
	}
	return d, nil
}

func (d *Directory) AddSubdir(dirname string) {
	var newdir = NewDirectory(dirname)
	newdir.Parent = d
	d.Subdirs = append(d.Subdirs, newdir)
}

func (d *Directory) AddFile(filename string, filesize int64) {
	var newfile = NewFile(filename, filesize)
	d.Files = append(d.Files, newfile)
}

type File struct {
	Name string
	Size int64
}

func NewFile(filename string, filesize int64) *File {
	var newfile File
	newfile.Name = filename
	newfile.Size = filesize
	return &newfile
}

type FileSystem struct {
	Root       *Directory
	Pwd        *Directory
	Size       int64
	Depth      int
	Width      int
	TotalDepth int
	TotalWidth int
}

func NewFileSystem() *FileSystem {
	var f FileSystem
	var Rootdir *Directory = NewDirectory("/")
	f.Root = Rootdir
	f.Root.Parent = f.Root
	f.Pwd = Rootdir
	f.Depth = 0
	f.Size = 0
	f.TotalDepth = 0
	f.TotalWidth = 1
	return &f
}

func (F *FileSystem) cd(arg string) (int, error) {
	if arg == "/" {
		// fmt.Printf("to root directory\n")
		F.Pwd = F.Root
		F.Depth = 0
		return 0, nil
	} else if arg == ".." {
		// fmt.Printf("from pwd: %s to parent: %s \n", F.Pwd.Name, F.Pwd.Parent.Name)
		F.Pwd = F.Pwd.Parent
		F.Depth -= 1
		return 0, nil
	}
	child, err := F.Pwd.Child(arg)
	if err != nil {
		// err := fmt.Errorf("child directory %s not found", arg)
		return 1, err
	} else {
		// fmt.Printf("from pwd: %s to child: %s \n", F.Pwd.Name, child.Name)
		F.Pwd = child
		F.Depth += 1
		return 0, nil
	}
}

func IsCommand(text string) bool {
	command, _ := regexp.Compile(`^\$`)
	return command.MatchString(text)
}

func IsFile(text string) bool {
	number, _ := regexp.Compile(`^\d+`)
	return number.MatchString(text)
}

func (f *FileSystem) ParseCommand(text string) {
	command := strings.Split(text, " ")[1]
	if command == "cd" {
		// fmt.Printf("\n%s : CD : ", text)
		arg := strings.Split(text, " ")[2]
		f.cd(arg)
	} // else if command == "ls" {
	// 	// fmt.Println(text, ":", "is command ls")
	// } else {
	// fmt.Println(text, ":", "command not recognized")

	//}
}

func (F *FileSystem) MapFs() {
	Depth := make(map[int][]*Directory)
	Depth[0] = append(Depth[0], F.Root)
	F.Root.MapDirectory(0, Depth)
	MaxWidth := 1
	TotalDepth := len(Depth)
	for i := 0; i < len(Depth); i++ {
		if len(Depth[i]) > MaxWidth {
			MaxWidth = len(Depth[i])
		}
	}
	for i := 0; i < len(Depth); i++ {
		// num_tabs := TotalDepth - i
		// tabs := ""
		// for t := 0; t < num_tabs; t++ {
		// 	tabs = tabs + " "
		// }
		for _, j := range Depth[i] {
			fmt.Printf("%s ", j.Name)
		}
		fmt.Println("")
	}
	fmt.Printf("Total Depth: %d Max Width: %d", TotalDepth, MaxWidth)
}

func (d *Directory) MapDirectory(initial_depth int, Map map[int][]*Directory) {
	depth := initial_depth + 1
	// fmt.Println(d.Name)
	// fmt.Println(d.Subdirs)
	for _, v := range d.Subdirs {
		// fmt.Printf("%s\t", v.Name)
		Map[depth] = append(Map[depth], v)
		v.MapDirectory(depth, Map)
	}
}
func (d *Directory) GetSize(counter *int64) int64 {
	var size int64 = 0
	for _, f := range d.Files {
		size += f.Size
	}
	for _, v := range d.Subdirs {
		v.Size = v.GetSize(counter)
		size += v.Size
	}
	if size > 6592386 {
		// *counter += size
		fmt.Printf("%d %s\n", size, d.Name)
	} // else {
	// 	fmt.Printf("TOO SMALL: %s : %d\n", d.Name, size)
	// }
	return size
}

func main() {
	dat, _ := os.Open("input.txt")
	defer dat.Close()
	scanner := bufio.NewScanner(dat)
	var FS = NewFileSystem()
	total_size := 0
	for scanner.Scan() {
		text := scanner.Text()
		if IsCommand(text) {
			FS.ParseCommand(text)
		} else if strings.Split(text, " ")[0] == "dir" {
			subdir := strings.Split(text, " ")[1]
			FS.Pwd.AddSubdir(subdir)
			// fmt.Printf("%s : adding subdir '%s' to directory '%s' \n", text, subdir, FS.Pwd.Name)
		} else if IsFile(text) {
			filesize, _ := strconv.Atoi(strings.Split(text, " ")[0])
			filename := strings.Split(text, " ")[1]
			FS.Pwd.AddFile(filename, int64(filesize))
			total_size += filesize
			// fmt.Printf("%s : adding file '%s' to directory '%s' \n", text, filename, FS.Pwd.Name)
		} else {
			fmt.Println(text, ":", "no action taken")
		}
	}
	var counter int64 = 0
	diskused := FS.Root.GetSize(&counter)
	diskfree := 70000000 - diskused
	needed := 30000000 - diskfree
	fmt.Printf("Total disk used: %d\n", diskused)
	fmt.Printf("Total disk free: %d\n", diskfree)
	fmt.Printf("Additional disk space needed: %d\n", needed)

}
