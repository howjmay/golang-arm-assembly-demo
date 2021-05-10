# Introduction

This is a demonstration for using SIMD in golang with assembly.
SIMD is a well-known hardware-specific technique to accelerate applications. SIMD allows applications to run several operations parallelly. Normally, we would develop SIMD applications in C language with inclining assembly or intrinsics. We can simply conduct inclining assembly or intrinsics with cgo, but call overhead of cgo would a critical issue in performance-sensitive cases, which is the exact case that we would use SIMD. Here we are going to share how to write assembly in Golang to facilitate SIMD technique without call overhead.



## How To Run Example

### basic example

```shell
make run-basic
```