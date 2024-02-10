# TXQR

[![GoDoc](https://godoc.org/github.com/Hatzen/txqr?status.svg)](https://godoc.org/github.com/Hatzen/txqr)

TXQR (Transfer via QR) is a protocol and set of tools and libs to transfer data via animated QR codes. It uses [fountain codes](https://en.wikipedia.org/wiki/Fountain_code) for error correction.

See related blog posts for more details:
 - [Animated QR data transfer with Gomobile and Gopherjs](https://Hatzen.github.io/posts/animatedqr/)
 - [Fountain codes and animated QR](https://Hatzen.github.io/posts/fountaincodes/)

# Setup for android 

go get golang.org/x/mobile/cmd/gomobile
go install golang.org/x/mobile/cmd/gomobile

C:\Users\*\AppData\Local\Android\Sdk


adjusted to generate qrcodes
txqr/cmd/txqr-gif


build aar:
gomobile bind -target android

# Demo

![Demo](./docs/demo.gif)

Reader iOS app in the demo (uses this lib via Gomobile): [https://github.com/Hatzen/txqr-reader](https://github.com/Hatzen/txqr-reader)

## Automated tester app
Also see `cmd/txqr-tester` app for automated testing of different encoder parameters.

# Licence

MIT
