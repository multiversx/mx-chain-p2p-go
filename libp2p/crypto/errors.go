package crypto

import "errors"

// ErrNilPrivateKey signals that a nil private key was provided
var ErrNilPrivateKey = errors.New("nil private key")

// ErrNilSingleSigner signals that a nil single signer was provided
var ErrNilSingleSigner = errors.New("nil single signer")
