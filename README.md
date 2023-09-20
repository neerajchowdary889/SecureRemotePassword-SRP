# SecureRemotePassword-SRP

## Zero Knowledge Proof
A zero-knowledge proof (ZKP) is a cryptographic method by which one party (the prover) can prove to another party (the verifier) that a statement is true, without revealing any information other than the fact that the statement is true.
*  For example, you could prove to me that you know the answer to a question by giving me a hint that only someone who knows the answer would be able to understand.

## SRP Protocol
The Secure Remote Password (SRP) protocol is a password-authenticated key exchange (PAKE) protocol that allows a user to authenticate themselves to a server without revealing their password. It is resistant to dictionary attacks and does not require a trusted third party.

SRP works by creating a large private key shared between the client and server. This key is used to generate a session key, which is then used to encrypt traffic between the two parties.

SRP is more secure than the alternative SSH protocol and faster than using Diffie-Hellman key exchange with signed messages. It is also independent of third parties, unlike Kerberos.

<!-- - Source : https://en.wikipedia.org/wiki/Secure_Remote_Password_protocol#:~:text=6.3%20Other%20links-,Overview,the%20user%20to%20the%20server. -->

## SRP and Zero Knowledge Proof
SRP uses a ZKP to prove to the server that the client knows the password without revealing the password itself. The ZKP is based on the idea that it is possible to prove that you know the value of a function without revealing the input to the function.

In the case of SRP, the function is a cryptographic hash function. The client knows the input to the hash function (their password), but the server does not. The client can use the ZKP to prove to the server that they know the output of the hash function (the verifier) without revealing the input.

This means that the server can be confident that the client knows the password without actually seeing the password itself.

SRP is a secure and efficient way to authenticate users to servers. It is used in a variety of standards and is supported by many different software implementations

If an attacker gets into a database and steals all the authentication information, they can only see some hashes and large 2048-bit prime numbers in case of SRP Protocol. **Current version of SRP is 6a**.

- *Note: SRP is generally slower than other authentication techniques, such as password verification or OAuth. This is because SRP involves performing a number of cryptographic operations, such as modular exponentiation and hash functions.*

| Authentication technique | Speed  | Security |
|--------------------------|--------|----------|
| SRP                      | Slow   | High     |
| Password Verification    | Fast   | Low      |
| OAuth                    | Medium | Medium   |