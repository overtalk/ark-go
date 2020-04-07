# bytebufferpool

An implementation of a pool of byte buffers with anti-memory-waste protection.

The pool may waste limited amount of memory due to fragmentation.
This amount equals to the maximum total size of the byte buffers
in concurrent use.

# Benchmark results
Currently bytebufferpool is fastest and most effective buffer pool written in Go.

You can find results [here](https://omgnull.github.io/go-benchmark/buffer/).

# This is a copy form [valyala](https://github.com/valyala/bytebufferpool)