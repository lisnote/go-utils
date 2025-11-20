# Log Command

`log-command` is a utility that acts as a wrapper for another executable. It logs the command-line arguments it is called with and then executes the original program.

## How it works

This program is designed to be renamed to match the target executable you want to log. The original executable should be renamed with a `-proxy` suffix.

For example, if you want to log the execution of `my-tool.exe`, you should:

1.  Rename `my-tool.exe` to `my-tool-proxy.exe`.
2.  Rename `log-command.exe` to `my-tool.exe`.

Now, when you run `my-tool.exe` with some arguments (e.g., `my-tool.exe --input file.txt`), the wrapper will:
1.  Create or append to a log file named `my-tool.log` in the same directory.
2.  Write an entry in the log file with the command and its arguments.
3.  Execute `my-tool-proxy.exe` with the original arguments (`--input file.txt`).

The standard input, output, and error streams are passed to the proxy executable, so it behaves transparently.
