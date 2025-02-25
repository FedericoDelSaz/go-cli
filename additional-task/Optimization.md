# Comparison: Current Approach vs. Merge Sort

| Feature                | Current Approach (`sort.Slice`)                     | Merge Sort Approach                  |
|------------------------|--------------------------------------------------|----------------------------------|
| **Sorting Algorithm**  | Hybrid (introsort: quicksort, heapsort, insertion sort) | Merge Sort                        |
| **Time Complexity**    | O(N log N) (average case)                          | O(N log N) (worst case)          |
| **Space Complexity**   | O(1) (in-place sorting)                            | O(N) (requires extra space)      |
| **Stability**          | Not stable (may reorder equal elements)            | Stable (preserves original order) |
| **Best Use Case**      | General-purpose sorting, faster for small to medium datasets | Large datasets, external sorting, linked lists |
| **Parallelization**    | Limited parallelism (depends on implementation)   | Easily parallelizable            |

## Why Use Merge Sort Instead?

### 1. **Stable Sorting**
- If stability is required, Merge Sort is a better choice since it maintains the order of duplicate elements.

### 2. **Handles Large Datasets Well**
- Merge Sort is particularly useful when sorting large datasets, especially when they don't fit into memory, since it can be adapted for external sorting (e.g., using disk-based merging).

### 3. **Parallelizable**
- Unlike Quicksort (which has recursive in-place partitioning), Merge Sort can be implemented with multiple threads, making it more efficient for large-scale sorting.

## Optimized Merge Sort Approach for Your CLI

Instead of sorting all numbers in memory, we could:

1. **Read the input file in chunks**, sort each chunk using Merge Sort, and write sorted chunks to temporary files.
2. **Merge sorted chunks efficiently** (k-way merge) to get the top N numbers.
3. **Stream the output** instead of keeping everything in memory to optimize resource usage.

This approach is known as **External Merge Sort**, which is well-suited for handling very large input files efficiently.

