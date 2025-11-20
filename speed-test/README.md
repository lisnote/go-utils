# Speed Test

`speed-test` is a simple client-server tool to measure network throughput, providing both upload and download speed metrics.

## Overview

The project consists of two components: a server and a client. The server listens for incoming connections, and the client connects to the server to perform the speed test. The test involves the client sending data to the server (to test upload speed) and receiving data from the server (to test download speed) simultaneously for a fixed duration.

## Features

*   **Client-Server Architecture**: A lightweight server and client model.
*   **Bidirectional Testing**: Measures both upload and download speeds in a single session.

## Usage

You need to run the server first, and then run the client to connect to the server.

### Server

The server listens on a specified IP address and port.

```bash
# Example: Listen on all interfaces on port 8080
server-speed-test.exe 0.0.0.0:8080
```

### Client

The client connects to the server's address and port to start the test.

```bash
# Example: Connect to a server at 127.0.0.1 on port 8080
client-speed-test.exe 127.0.0.1:8080
```

After the test completes (default is 10 seconds), the client will output the total data transferred and the average transfer speed for both upload and download.
