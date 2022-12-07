package dir

import (
	"fmt"
	"math"
)

type Directory struct {
	Name     string
	Parent   *Directory
	Children []*Directory
	Size     int
}

func (d *Directory) TotalSize() int {
	return calcDirectorySize(d)
}

func (d *Directory) SumWithinSizeLimit(sizeLimit int) int {
	return sumDirectoriesWithinSizeLimit(d, sizeLimit)
}

func (d *Directory) MinSizeOverTarget(targetSize int) int {
	return minDirectorySizeOverTarget(d, targetSize, math.MaxInt)
}

func (d *Directory) Print() {
	printDirectoryStructure(d, 0)
}

func calcDirectorySize(d *Directory) int {
	total := d.Size
	for _, child := range d.Children {
		total += calcDirectorySize(child)
	}
	return total
}

func sumDirectoriesWithinSizeLimit(d *Directory, sizeLimit int) int {
	total := 0
	for _, child := range d.Children {
		total += sumDirectoriesWithinSizeLimit(child, sizeLimit)
	}
	size := calcDirectorySize(d)
	if size <= sizeLimit {
		total += size
	}
	return total
}

func minDirectorySizeOverTarget(d *Directory, targetSize int, currSize int) int {
	for _, child := range d.Children {
		currSize = minDirectorySizeOverTarget(child, targetSize, currSize)
	}
	size := calcDirectorySize(d)
	if size >= targetSize && size < currSize {
		currSize = size
	}
	return currSize
}

func printDirectoryStructure(d *Directory, level int) {
	for i := 0; i < level; i++ {
		fmt.Print("  ")
	}
	fmt.Printf("/%s [size=%d]\n", d.Name, d.Size)
	for _, child := range d.Children {
		printDirectoryStructure(child, level+1)
	}
}
