package main

import (
	"fmt"
	"math"
	"math/rand"
	"os"
	"time"

	"github.com/alexflint/go-arg"
	"github.com/pietroagazzi/useless-sorting-algorithms/pkg/sorting"
	"github.com/pietroagazzi/useless-sorting-algorithms/pkg/utils"
)

var args struct {
	Lenght    int    `arg:"-l,--lenght" help:"number of elements to sort" default:"10" placeholder:"<n>"`
	Algorithm string `arg:"-a,--algorithm" help:"algorithm to use" default:"bogosort" placeholder:"<algorithm>"`
	List      bool   `arg:"-L,--list" help:"list all algorithms"`
}

type Algorithm struct {
	Sorter
	Description string
}

type Sorter func([]int) []int

var algorithms = map[string]Algorithm{
	"bogosort": {
		Sorter:      sorting.Bogosort,
		Description: "An extremely inefficient sorting algorithm that shuffles the array until it is sorted.",
	},
	"gnomesort": {
		Sorter:      sorting.Gnomesort,
		Description: "A sorting algorithm that is based on the idea of a garden gnome sorting his flower pots.",
	},
	"sleepsort": {
		Sorter:      AdaptUint(sorting.Sleepsort),
		Description: "An inefficient, but interesting, sorting algorithm that sorts numbers by sleeping for an interval corresponding to each number.",
	},
	"stalinsort": {
		Sorter:      sorting.Stalinsort,
		Description: "A joke sorting algorithm that 'solves' the problem of sorting a list by eliminating all elements out of order.",
	},
}

func init() {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	arg.MustParse(&args)

	if args.List {
		fmt.Println("List of useless sorting algorithms:")
		for k := range algorithms {
			fmt.Println("-", k, "\n   ", algorithms[k].Description)
		}
		os.Exit(0)
	}

	if _, ok := algorithms[args.Algorithm]; !ok {
		fmt.Println("Invalid algorithm. Use -L to list all algorithms")
		os.Exit(1)
	}

	if args.Lenght < 1 {
		fmt.Println("Lenght must be greater than 0")
		os.Exit(1)
	}

	fmt.Println("Useless sorting algorithms!")
}

func main() {
	arr := utils.GenerateRandomArray(args.Lenght, 0, int(math.Pow(float64(args.Lenght), 2)))

	fmt.Printf("Running %s with %d elements\n\nInput: %v\n", args.Algorithm, args.Lenght, arr)

	start := time.Now()

	result := algorithms[args.Algorithm].Sorter(arr)

	elapsed := time.Since(start)

	fmt.Printf("Output: %v\n", result)
	fmt.Printf("Is sorted? %t\n", utils.IsSorted(result))

	fmt.Printf("\nElapsed time: %s\n", elapsed)
}

// AdaptUint adapts a function that works on uint slices to work on int slices
func AdaptUint(f func([]uint) []uint) Sorter {
	return func(arr []int) []int {
		uintArr := make([]uint, len(arr))
		for i, v := range arr {
			uintArr[i] = uint(v)
		}

		result := f(uintArr)

		intArr := make([]int, len(result))
		for i, v := range result {
			intArr[i] = int(v)
		}

		return intArr
	}
}
