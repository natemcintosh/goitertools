# goitertools
An attempt at a python-like itertools module.

## Testing
To ensure that this library is as correct as possible, it is fuzzed against the python 
[itertools library](https://docs.python.org/3/library/itertools.html). To do this, random
inputs are generated, run through the python functions, the python results are captured
and compared with the output of this library. The [go fuzzing](https://go.dev/security/fuzz/)
library is used to generate the random inputs, and iterate over them in parallel.
