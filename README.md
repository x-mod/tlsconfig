tlsconfig
===

Create tls.Config with Options.

[中文说明](https://gitdig.com/)

# Mutual TLS Authentication

** Server **

````go
import "github.com/x-mod/tlsconfig"

cf := tlsconfig.New(
    //服务端 TLS 证书
    tlsconfig.CertKeyPair("out/server.crt", "out/server.key"), 
    //客户端 TLS 证书签名 CA
    tlsconfig.ClientCA("out/exampleCA.crt"), 
    //验证客户端证书
    tlsconfig.ClientAuthVerified(),
)

````

** Client **

````go

import "github.com/x-mod/tlsconfig"

cf := tlsconfig.New(
    //服务端 TLS 证书签名 CA
    tlsconfig.CA("out/exampleCA.crt"), 
    //客户端证书 TLS 证书
    tlsconfig.CertKeyPair("out/client.crt", "out/client.key"), 
)
````