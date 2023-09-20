# SecureRemotePassword-SRP

| Authentication technique | Speed  | Security |
|--------------------------|--------|----------|
| SRP                      | Slow   | High     |
| Password Verification    | Fast   | Low      |
| OAuth                    | Medium | Medium   |

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

### Notations SRP used in this description of the protocol, version 6a
| Number 	| Notation 	| Real Value                                                        	|
|--------	|----------	|-------------------------------------------------------------------	|
| 1      	| Q        	| Large Sophie Germain Prime (over 2047 bit long)                   	|
| 2      	| N        	| Safe Prime --> N = 2Q+1 (over 2048 bit long)                      	|
| 3      	| g        	| Generator of the multiplicative group                             	|
| 4      	| H()      	| Hashing technique SHA256                                          	|
| 5      	| K        	| K = H(N,g)                                                        	|
| 6      	| S        	| Salt (over 64 bit long)                                           	|
| 7      	| I        	| Username                                                          	|
| 8      	| P        	| Password                                                          	|
| 9      	| x        	| x = H(S\|H(I\|":"\|P)) --> (RFC2945 standard)                     	|
| 10     	| V        	| V = (pow(g,x) mod N)                                                  |
| 11     	| A & B    	| Random one time ephemeral keys of the user and host respectively. 	|

## How SRP works

### 1. User signup
- **Step 1:** Generate Q, from Q generate N.
- **Step 2:** Compute g, Generator of the multiplicative group of N.
- **Step 3:** Generate Salt S.
- **Step 4:** Take Username and Password from user.
- **Step 5:** Compute x = H(S\|H(I\|":"\|P))
- **Step 6:** Compute V = (pow(g,x) mod N)
- **Step 7:** Send H(Username), Salt, G, K, V to server in a single registration request.
- **Step 8:** Server stores Salt, G, K, V indexed by Username.
- **Step 9:** Client must not share x with anybody and must safely erase it this step. It is almost equivalent to plaintext password p.
- *Note: **The transmission and authentication of the registration request is not covered in SRP**. We encrypt this registration request, and a key get generated on both sides using diffie helman key exchange*