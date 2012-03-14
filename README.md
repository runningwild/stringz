String algorithm toolkit.  Go installable.  Despite the name 'stringz' these functions are all designed to operate on []byte.

    go install github.com/runningwild/stringz

If you need more control over the underlying algorithms, for example to adjust buffer sizes, you can use github.com/runningwild/stringz/core.

As long as it remains up you can see the godoc at http://gopkgdoc.appspot.com/pkg/github.com/runningwild/stringz.

This is a work in progress, but anything exposed in the high-level interface (stringz, not stringz/core) should be stable and efficient.  Comments, questions, suggestions, patches, etc... are all welcome.
