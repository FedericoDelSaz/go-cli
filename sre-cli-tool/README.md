# Top Numbers CLI

## Description
The `top_numbers` CLI is a command-line tool that reads a file containing numbers and outputs the top N highest numbers.

## Prerequisites
Ensure you have the following installed on your system:
- [Go](https://go.dev/dl/) (Golang)
- [Docker](https://www.docker.com/)
- A terminal or command prompt

## Installation
Clone the repository and navigate into the project directory:
```sh
git clone https://github.com/yourusername/top_numbers.git
cd top_numbers
```

Build the executable:
```sh
go build -o top_numbers main.go
```

## Building the Docker Image
To build a Docker image for the `top_numbers` CLI, run the following command:
```sh
docker build -t top_numbers .
```

## Usage
To discover input options, run:
```sh
./top_numbers -h
```

### Output (`-h` or `--help`)
```sh
Usage of ./top_numbers:
  -input-file string
        Path to the input file.
  -n int
        The value N for the number of top results.
  -output-file string
        Path to the output file.
```

### Running the CLI
Run the CLI with the following syntax:
```sh
./top_numbers -n <N> -input-file <filename> -output-file  <filename>
```
- `-n <N>`: Number of top values to extract
- `-input-file <filename>`: Path to the input file containing numbers
- `-output-file  <filename>`: Path to the output file where results will be saved

#### Example
```sh
./top_numbers -n 5 -input-file numbers.txt -output-file  output.txt
```
This will extract the top 5 highest numbers from `numbers.txt` and save them in `output.txt`.

### Running with Docker
You can also run the application inside a Docker container:
```sh
docker run --rm -v $(pwd):/app top_numbers -n <N> -input-file /app/<filename> -output-file  /app/<output_filename>
```

#### Example
```sh
docker run --rm -v $(pwd):/app top_numbers -n 5 -input-file /app/numbers.txt -output-file  /app/output.txt
```

## File Format
The input file should contain one number per line, like this:
```
10
45
2
99
7
```
