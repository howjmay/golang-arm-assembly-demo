# Introduction

This is a demonstration for using golang assembly SIMD in golang with assembly.
SIMD is a well-known hardware-specific technique to accelerate applications. SIMD allows applications to run several operations parallelly. Normally, we would develop SIMD applications in C language with inclining assembly or intrinsics. We can simply conduct inclining assembly or intrinsics with cgo, but call overhead of cgo would a critical issue in performance-sensitive cases, which is the exact case that we would use SIMD. Here we are going to share how to write assembly in Golang to facilitate SIMD technique without call overhead.

## How To Run

### basic example

```shell
$cd basic
$go build
$./basic
```

### return_val

Demonstrate the way how go assembly returns values in function.

```shell
$cd return_val
$go build
$./return_val
```

### decl_array

Demonstrate the way how go assembly declaring array on frame stack.

```shell
$cd decl_array
$go build
$./decl_array
```

### math

Include several ARM64 implementation of math functions