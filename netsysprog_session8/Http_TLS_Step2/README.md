In here, we will learn about HTTP with TLS, we're gonna implement it

Teasers:
"you will know how will a website that previously used HTTP become HTTPS"

!Important
- although the simple http server and http tls code is slightly different
notice that i've change the port from 5555 to 5556, so that the simple http server can run without the port being reserved with certificate, that requires for the client to request with https
- link to mkcert download : https://github.com/FiloSottile/mkcert
- here is the things that happen:
when you try to run the command 
mkcert localhost 127.0.0.1
mkcert localhost 127.0.0.1
Note: the local CA is not installed in the system trust store.
Run "mkcert -install" for certificates to be trusted automatically ⚠️

Created a new certificate valid for the following names 📜
 - "localhost"
 - "127.0.0.1"

The certificate is at "./localhost+1.pem" and the key at "./localhost+1-key.pem" ✅

It will expire on <Day> <Month> <Year> 🗓


Notes : by default, the cert.pem and key.pem won't appear, it needs to be renamed if you want to rename the file,


on Mac/Linux 
 
❯ mv localhost+1-key.pem key.pem                             

❯ mv localhost+1.pem cert.pem                            

on Windows 

❯ ren localhost+1-key.pem key.pem                             

❯ ren localhost+1.pem cert.pem   

- don't forget to mkcert -install, or else your certificate won't work

# so what's TLS ?
referring to the textbook chapter 11, TLS is a protocol that supplies
secure communication between client and server, so maybe it's kinda like firewall

y' know, like blocking ip and ensuring the connections (you choose which to connect / can connect)

so for more explanation, the TLS allows the client to authenticate with the server and optionally permits the server to authenticate the clients.


the MC here is the client. The client uses the TLS to encrypt it's communication with the server
to prevent third-party interception and manipulation. (so there be no snooping, hijacking, hacking, whatever you like to call it)


let's see how the simulations looks like:

example:
If the client initiated a TLS 1.3 handshake with the server,

Client:  
Hello google.com. I’d like to communicate with you using TLS
version 1.3. Here is a list of ciphers I’d like to use to encrypt our mes-
sages, in order of my preference. I generated a public- and private-key
pair specifically for this conversation. Here’s my public key.


Server: 
Greetings, client. TLS version 1.3 suits me fine. Based on your
cipher list, I’ve decided we’ll use the Advanced Encryption Standard
with Galois/Counter Mode (AES-GCM) cipher. I, too, created a new
key pair for this conversation. Here is my public key and my certificate
so you can prove that I truly am google.com. I’m also sending along a
32-byte value that corresponds to the TLS version you asked me to use.
Finally, I’m including both a signature and a message authentication code
(MAC) derived using your public key of everything we’ve discussed so
far so you can verify the integrity of my reply when you receive it.


Client (to self): 
An authority I trust signed the server’s certificate, so
I’m confident I’m speaking to google.com. I’ve derived this conversa-
tion’s symmetric key from the server’s signature by using my private
key. Using this symmetric key, I’ve verified the MAC and made sure
no one has tampered with the server’s reply. The 32 bytes in the reply
Securing Communications with TLS 243
corresponds to TLS version 1.3, so no one is attempting to trick the
server into using an older, weaker version of TLS. I now have everything
I need to securely communicate with the server.


Client (to server) : Here is some encrypted data.


#TL;DR
you can just skip the simulation part, it's served for cases like, 
you're still wonderin the TLS thingy
