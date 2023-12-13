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

func init() {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	arg.MustParse(&args)

	fmt.Println("Useless sorting algorithms!")
}

func AdaptSleepsort(arr []int) []int {
	var uintArr []uint

	for _, num := range arr {
		uintArr = append(uintArr, uint(num))
	}

	result := sorting.SleepSort(uintArr)

	var adaptedResult []int

	for _, num := range result {
		adaptedResult = append(adaptedResult, int(num))
	}

	return adaptedResult
}

func main() {
	var algorithms = map[string]func([]int) []int{
		"bogosort":   sorting.Bogosort,
		"stalinsort": sorting.StalinSort,
		"sleepsort":  AdaptSleepsort,
	}

	if args.List {
		fmt.Println("List of useless sorting algorithms:")
		for k := range algorithms {
			fmt.Println("\t-", k)
		}
		os.Exit(0)
	}

	var arr []int

	if _, ok := algorithms[args.Algorithm]; !ok {
		fmt.Println("Invalid algorithm. Use -L to list all algorithms")
		os.Exit(1)
	}

	if args.Lenght < 1 {
		fmt.Println("Lenght must be greater than 0")
		os.Exit(1)
	}

	var max = int(math.Pow(float64(args.Lenght), 2))

	for i := 0; i < args.Lenght; i++ {
		arr = append(arr, rand.Intn(max))
	}

	fmt.Printf("Running %s with %d elements\n\nInput: %v\n", args.Algorithm, args.Lenght, arr)

	start := time.Now()

	result := algorithms[args.Algorithm](arr)

	elapsed := time.Since(start)

	fmt.Printf("Output: %v\n", result)
	fmt.Printf("Is sorted? %t\n", utils.IsSorted(result))

	fmt.Printf("\nElapsed time: %s\n", elapsed)
}
