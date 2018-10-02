# Mock LedFS

> What's better than an array of Neopixel LEDs you can control with a filesystem? A virtual array of Neopixel LEDs rendered in the browser that you can control with a filesystem!

Mock LedFS is a modified version of LedFS that visualizes Neopixels in the browser rather than communicating with the real hardware. The main purpose of this project is to provide a simple devtool for developing applications that use with LedFS, without using real hardware. This application should make debugging your animations and rendering a breeze, and simplify your dev-cycle since everything can be developed on a single machine (a laptop for example). Mock LedFS strictly follows the API provided by LedFS, but uses WebRTC and the browser to render what your hardware might look like in the same context, instead of communicating directly with an array of Neopixels.

## Installation & Usage
This project can be installed using `go get` like so:
`go get -u gitlab.com/rucuriousyet/mockledfs`

Once you've installed Mock LedFS, usage is as simple as running binary with your intended mountpoint directory (must be a directory since multiple files are used in this filesystem).

```sh
mockledfs [mountpoint] [flags]
```

See `mockledfs help` for usage details and flags.

If you'd like to find learn more about how to use LedFS, please check out the documention in-readme [here](https://gitlab.com/rucuriousyet/ledfs).
