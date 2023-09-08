# go-lwe
Learning with errors implemented in golang

The LWE cryptosystem is based on a mathematical problem called the Learning With Errors problem. The basic idea is to create a public key and a private key. You can encrypt messages with the public key, and only the holder of the private key can decrypt them.

Here are the steps involved in building a basic LWE-based cryptosystem in Golang:

1. **Choose LWE Parameters**:
   - Choose parameters for your LWE problem. This includes selecting the modulus `q`, the number of variables `n`, and the noise parameter `σ`. These parameters impact the security and performance of your system.

2. **Key Generation**:
   - Generate a random secret vector `s` where each element is sampled from a small range.
   - Generate a random error vector `e` where each element is sampled from a Gaussian distribution with standard deviation `σ`.
   - Compute the public key `A`, which is a matrix where each element is sampled uniformly at random modulo `q`.

3. **Encryption**:
   - To encrypt a message `m`, first, generate a random vector `r` where each element is sampled from a small range.
   - Compute `c1 = A * r + e`, where `*` denotes matrix-vector multiplication modulo `q`.
   - Compute `c2 = m + s * r`, where `+` denotes element-wise addition modulo `q`.
   - The ciphertext is `(c1, c2)`.

4. **Decryption**:
   - To decrypt the ciphertext `(c1, c2)`, compute `m' = c2 - s * c1`.
   - The original message `m` can be obtained by rounding `m'` to the nearest integer modulo `q`.

5. **Security Considerations**:
   - Implement a secure random number generator for sampling `s`, `e`, and `r`.
   - Ensure proper handling of secret keys and minimize information leakage.

6. **Testing**:
   - Implement test cases to verify the correctness and security of your cryptosystem.

Please note that building a secure cryptosystem is a complex and nuanced task, and this simplified example may not be suitable for production use. In practice, you should rely on established cryptographic libraries and consult with experts in the field to ensure that your implementation is secure.

Also, remember that cryptographic algorithms and systems evolve over time, and what is considered secure today may not be secure in the future. Therefore, it's essential to stay informed about the latest developments in cryptography and update your systems accordingly.
