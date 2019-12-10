## OpenSSL wrapper for go

To use:

    import "github.com/IBM-Bluemix/golang-openssl-wrapper/crypto" // For encryption, e.g. EVP
    import "github.com/IBM-Bluemix/golang-openssl-wrapper/ssl"    // TLS
    import "github.com/IBM-Bluemix/golang-openssl-wrapper/rand"   // PRNG
    import "github.com/IBM-Bluemix/golang-openssl-wrapper/digest" // Message-digest (hash) functions
    import "github.com/IBM-Bluemix/golang-openssl-wrapper/bio"    // OpenSSL-specific I/O handling
    
This library is being actively developed.  It provides access to much of the OpenSSL APIs for cryptography (libcrypto) and TLS (libssl).  We do not plan to provide the complete API at first, but if there is a function you need, please open an issue or submit a PR.  The wrapper is built on `swig`, so you will need to work directly with the `.swig` files to add functionality.

If you submit a pull request, please be sure to include complete unit test coverage (we use `ginkgo` and `gomega`, which sit on top of golang's native testing facility), or your PR will be declined.


It based on the OpenSSL library and requires it to be installed with the FIPS module.

## Install OpenSSL

### FIPS module
 
1. Download FIPS module (openssl-fips-2.0.16.tar.gz) from [here](https://www.openssl.org/source/)
2. Unzip
3. Run from folder
    1. Unordered sub-list. 
    2. ./config (may require to specify folder and compiler. For me worked following line: ./Configure darwin64-x86_64-cc --prefix=/usr/local/opt --openssldir=/usr/local/opt/openssl)
    3. make
    4. sudo make install

### FIPS enabled OpenSSL
1. Download FIPS capable OpenSSL of version 1.0.2 (openssl-1.0.2t.tar.gz) from same link above
2. Unzip
3. Run from folder
    1. ./config (may require to specify folder and compiler. For me worked following line: ./Configure darwin64-x86_64-cc --prefix=/usr/local/opt --openssldir=/usr/local/opt/openssl)
    2. make
    3. make depend (only if being asked)
    4. make test (make sure all tests passed)
    5. sudo make install
 
[OpenSSL Github](https://github.com/openssl/openssl)

[Their help](https://wiki.openssl.org/index.php/Compilation_and_Installation) can be useful also, although it’s a one big mess
