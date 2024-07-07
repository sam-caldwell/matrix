Matrix Math
===========

## Objective
* Develop a matrix math package which is capable of leveraging NVIDIA CUDA.

## Usage
The `Matrix` type is a wrapper for `gocv.Mat` which includes
a 2D Matrix of float64 elements and a lock to make the type
safe for concurrency use-cases.

## Initialization

```go
var m Matrix
```

or

```go
rows:=5
cols:=10
m:=matrix.NewMatrix(rows,cols)
```

## Getter/Setter Methods

```go
var m Matrix
var f [][]float64
if err:=m.Set(f); err != nil {
    return err
}
```

## Mathematics
### Add

### Sum

### Multiply

### Divide

### Bitwise: And
### Bitwise: Or
### Bitwise: Not
### Bitwise: Xor
