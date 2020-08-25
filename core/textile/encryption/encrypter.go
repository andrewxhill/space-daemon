package encryption

import "io"

type BucketItemEncrypter interface {
	// Returns the encrypted path and reader that reads the encrypted version of the passed in reader
	Encrypt(path string, content io.Reader) (encryptedPath string, reader io.Reader)
	// Returns the decrypted path and reader that reads the decrypted version of the passed in reader
	Decrypt(path string, content io.Reader) (decryptedPath string, reader io.Reader)
}
