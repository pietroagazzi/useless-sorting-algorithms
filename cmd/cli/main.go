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

var algorithms = map[string]Sorter{
	"bogosort":   sorting.Bogosort,
	"sleepsort":  AdaptUint(sorting.Sleepsort),
	"stalinsort": sorting.Stalinsort,
	"gnomesort":  sorting.Gnomesort,
}

type Sorter func([]int) []int

func init() {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	arg.MustParse(&args)

	if args.List {
		fmt.Println("List of useless sorting algorithms:")
		for k := range algorithms {
			fmt.Println("\t-", k)
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

	result := algorithms[args.Algorithm](arr)

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
