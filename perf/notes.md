Perf improvements from: https://chatgpt.com/c/6804a517-09f4-8005-b34c-3ef87f5979e5

1. BaselinePerf doc
Buffer syscalls with bufio.Reader/Writer.
Generate or swap JSON library to avoid reflection.
Pool encoders/decoders and buffers.
Replace fmt with lightweight logging.
Move to binary protocols if JSON is too heavy.
Tune TCP (NODELAY, buffer sizes).
Pipeline work to overlap CPU & IO.
Measure after each change with pprof and bench.