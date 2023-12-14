# Useless Sorting Algorithms

This project is a collection of some of the most inefficient sorting algorithms ever created. It's written in Go and it's meant to be a fun exploration of algorithmic inefficiency.

## Getting Started

### Installing

```sh
# Clone the repository to your local machine:
git clone https://github.com/pietroagazzi/useless-sorting-algorithms.git

# Navigate to the project directory:
cd useless-sorting-algorithms

# Build the project:
make build
```

### Running the tests

To run the tests, use the following command:

```sh
make test
```

### Usage

You can run the program with the following command:

```sh
./useless-sorting-algorithms
# or
make run ARGS="..."
```

By default, the program will sort a list of 10 numbers using the Bogosort algorithm. You can specify the length of the list and the sorting algorithm with the -l and -a flags, respectively. For example, to sort a list of 20 numbers using the Sleepsort algorithm, you can use the following command:

```sh
./useless-sorting-algorithms -l 20 -a sleepsort
```

You can list all available algorithms with the -L flag:

```sh
./useless-sorting-algorithms -L
```

## Algorithms

The project includes the following sorting algorithms:

1. [**Bogosort**](https://en.wikipedia.org/wiki/Bogosort): An extremely inefficient sorting algorithm based on the generate and test paradigm. The algorithm successively generates permutations of its input until it finds one that is sorted.

2. [**Stalinsort**](https://github.com/gustavo-depaula/stalin-sort): A joke sorting algorithm that 'solves' the problem of sorting a list by eliminating all elements out of order.

3. [**Sleepsort**](https://it.wikipedia.org/wiki/Sleep_sort): An inefficient, but interesting, sorting algorithm that sorts numbers by sleeping for an interval corresponding to each number and then printing it.

4. [**Gnomesort**](https://en.wikipedia.org/wiki/Gnome_sort): A sorting algorithm that is based on the idea of a garden gnome sorting his flower pots.

5. [**Slowsort**](https://en.wikipedia.org/wiki/Slowsort): An inefficient sorting algorithm based on the multiply and surrender paradigm.

6. [**Bogobogosort**](https://en.wikipedia.org/wiki/Bogobogosort): An extremely inefficient sorting algorithm based on the generate and test paradigm. The algorithm successively generates permutations of its input until it finds one that is sorted.
