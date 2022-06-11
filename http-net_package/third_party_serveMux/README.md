# Julien Schimdt's httprouter package

<!-- markdownlint-disable MD033 -->

<a href="https://pkg.go.dev/github.com/julienschmidt/httprouter@v1.3.0">httprouter Package</a>

HttpRouter is a lightweight high performance HTTP request router (also called multiplexer or just mux for short) for Go.

In contrast to the default mux of Go's net/http package, this router supports variables in the routing pattern and matches against the request method. It also scales better.

The router is optimized for high performance and a small memory footprint. It scales well even with very long paths and a large number of routes. A compressing dynamic trie (radix tree) structure is used for efficient matching.
