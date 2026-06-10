# Chapter 9: Application Developments

### 9.9.1 Encryption Techniques
- **encryption:** the process of encryption denpend heavily on the encription key
- **symmetric encryption:** the same key is used for encryption and decryption, it is faster than asymmetric encryption but less secure
- **asymmetric encryption:** use a pair of key, one for encryption (public key) and one for decryption (private key) 

- symetric encryption: work base on the key mechanism, a value xor with encrtyption twice will return the original value
- **dictionary attack:** an attack that try to guess the encryption key by trying all possible value in a dictionary of common key

### 9.9.3 Encryptiong and Authentication
- **challenge-response:** a server send a random value (challenge) to the client, the client encrypt the challenge with encrypted key and send it back to the server, the server decrypt the response with the same key and compare it with the original challenge
    + challenge-response can also be used with asymmetric encryption
- **digital certificate:** a digital certificate is a electronic document that use to prove the ownership of a public key, it is issued by a trusted third party called Certificate Authority (CA)
- authenticating a server through only a CA is create burden for the CA. To solve this problem, we can use a chain of trust, where a CA can issue a certificate for another CA, and the client can verify the server's certificate by following the chain of trust up to a trusted CA
- to reduce encryption cost, HTTPS actually create aone time symmetric key after authenticating and use it to encrypt data (Diffie-Hellman key exchange)