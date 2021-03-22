# WordIDE

Have you ever wondered: How would it feel like to write code in a word processor? Me neither. But after ~~months~~minutes of planning, I present to you (drum rolls)... WordIDE!

![WordIDE Screenshot](/images/screenshot.png)

Write your code in any word processor, and WordIDE will convert it to plain text, so that your dumb compiler understands it. It's as simple as

```
$ wide compile test.odt
```

#### Advantages
* Color your code manually
* Use any fonts
* Use any text size
* Use any word processor

#### Disadvantages
* None

Note that this was a weekend project, so the code is... less than perfect. Still, it does the job. I plan on making it better in the future, and if you feel like fixing it yourself, submit a pull request!


### Building
Just make sure you have a Go compiler installed, then run `make`. If you don't have `make` installed, you can also do it with `go build`, but make sure to supply another output binary with `-o binary`, because the default one conflicts with the package name `wordide`.
