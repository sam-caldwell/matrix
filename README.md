Matrix Math
===========

## Description
Develop a matrix math package with optional Nvidia CUDA support.

## Usage

- Build with `cuda` flag and the package will use the GPU for matrix math.
- Not using the `cuda` flag will result in a non-GPU capability.

```go
package main

import "github.com/sam-caldwell/matrix"

func main() {
	a := matrix.NewMatrix(5, 5)
	b := matrix.NewMatrix(5, 5)

	m := matrix.NewSimple2d[float64](5, 5)

	a.Set(m)
}

```