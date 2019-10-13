/*
Package cipher supplies the encrypting and decrypting funcitonality crucial to
Delfin. A Cipher is a stream cipher which can return pseudorandom masking
bytes one at a time. As its underlying implementation is subject to change, a
	Cipher	interface is provided which can be implemented by any one of many
cipher types.
*/
package cipher

// A Cipher represents a stream cipher which can return masking bytes.
type Cipher interface {
	GetByte() byte
}
