# GoX Full Stack Web Framework

This is the v0 of a framework that aims to bring the power of Go and WASM to the maximum.

Inspired by JavaScript/TypeScript Frameworks like NextJS, SolidJS, Nuxt, and React, it's use an approach of Views, but there are plans to make it work similar to JSX/TSX.

The project structure is pretty straightforward:

- `assets/` represents all static files being served to the front.
- `bin/` the backend compiled binaries folder.
- `internal/` the internal files of the framework
- `server/` the backend part of the server (databases, apis, logic)
- `utils/` the common code between frontend and backend
- `ẁasm/` your front-end powered by WASM.

Why I created this framework?

> well, I thought that would be possible, and I was talking with a friend about creating a "GoScript" (a Go similar to JavaScript for browsers). And of course I made because I could.


\* Please use the CLI Shipped on the CLI Folder to compile this Framework, I recommend it because it compiles the right folders on the right way and apply some optmizations to the final WASM binary.


Some abstractions need to be made on this framework (cuz be working on the pure WASM DOM Manipulation sucks), I'll be working on them here.

