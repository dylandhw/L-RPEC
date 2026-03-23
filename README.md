# l-rpec 
### A lightweight reverse proxy with edge caching 
---------
Routes incoming requests to upstream servers based, on config, caches response in-mem, then signs outbound requests with HMAC

Essentiall a toy Cloudlfare worker's runtime 

Some references that helped build this:
- [What is a reverse proxy?](https://www.cloudflare.com/learning/cdn/glossary/reverse-proxy/)
- [HTTP Caching](https://developer.mozilla.org/en-US/docs/Web/HTTP/Caching)
- [How CDNs Work](https://www.youtube.com/watch?v=RI9np1LWzqw)
- [Signing requests](https://developers.cloudflare.com/workers/examples/signing-requests/)
- [HMAC](https://en.wikipedia.org/wiki/HMAC)
