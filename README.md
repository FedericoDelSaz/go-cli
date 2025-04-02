# SRE Code Test

### The Problem

Write a simple CLI tool (feel free to use any programming language you're
comfortable with) that accepts the following arguments:

- `-n int`: An integer for the value N, which will be output length.
-  `-input-file string`: A string representing a path to an input file.
- `-output-file  string`: A string representing a path to an output file.

The input file will contain one number per line. The valid number is anything
that is a valid unsigned 64-bit integer. The goal of your tool is to read the
input file, get N largest numbers, sort them in descending order, and output
the result to the output file, one number per line.


#### The Requirements

Your program will have to comply with the following requirements:

1. Must build and execute in a Unix operating system.
2. Exit with the following custom errors when using wrong arguments:
    - The input file does not exist: `ERROR: input file does not exist.`
    - The input file is not readable: `ERROR: input file is not readable.`
    - The integer N argument is less than or equal to zero: `ERROR: Number of top results must be bigger than 0.`
    - The integer N argument is over the upper limit: `ERROR: The maximum number of top results must be less or equal than 30000000.`
3. The program must ignore and warn about invalid lines in the following format: `WARN: Invalid line <line>.`
4. Invalid lines are all lines that do not contain an unsigned 64-bit integer, for example:
    - Multiple numbers on the same line: 12345 890
    - Not unsigned 64-bit integers: `1234helo` or `12345.01`
5. The program must handle big files with tens of millions of lines efficiently.
6. Avoid using third-party libraries and stick with the standard ones.

#### The Deliverable

Since this is an _SRE_ code test, the deliverable is super simple: your
solution, packed in the docker container. Simply provide a Dockerfile in the
repo, and you're all set. When grading your solution, we will first `docker
build` it, and then we will `docker run` it with various input parameters,
according to the argument specification above.

### Additional Task

Please create a text file with the answers to the following questions:
- What would be the time and space complexity of your program? Why?
- Do you think there is still room for improvement in your solution? If so, please elaborate.
