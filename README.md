# Attempting to connect to a Windows server from WSL

This repository explores how to get a server running on a Windows host, serving requests to clients running inside the WSL distros.

## Build the project

The code is a basic HTTP server that responds `Hello, world`. On WSL, run:
```bash
# Install Go if you don't have it
sudo apt update
sudo apt install -y golang-go

# Build
./build.sh
```

## Find out your Windows host from within WSL

Now, find your WSL localhost address. Run the following command in any WSL distro:
```
wslinfo --networking-mode
```
1. If it says `nat`, read file `/etc/resolv.conf` and look for the `nameserver`:
   1. If the nameserver is NOT loopback (`127.x.x.x`), then it is the address of the Windows host.
   2. Otherwise, the default gateway from `$ ip route` is the address of the Windows host.
2. Otherwise, use localhost (`127.0.0.1`).

In my case, it is `172.25.32.1`. I'll refer to it as `${HOST_ADDRESS}`.

## Server on Windows, connect from within WSL
Here are a few attempts on how to have a server on Windows serving requests coming from within WSL.

### Server listens to localhost
<details> <summary> Show steps </summary>
On Windows:

```powershell
.\server\server.exe "localhost:8080"
```
Without stopping the server, start another shell on WSL and connect to the Windows host. In my case:

```bash
./client/client "http://${HOST_ADDRESS}:8080"
```
```
Error connecting to http://172.25.32.1:8080: Get "http://172.25.32.1:8080": context deadline exceeded (Client.Timeout exceeded while awaiting headers)
```
</details>


Unfortunately, this does not work.

### Server listens to WSL network

<details> <summary> Show steps </summary>

Run `ipconfig` on Windows and find the ipv4 address of the WSL network adapter. For me, it is `172.25.32.1`. I'll refer to it as `$WSL_ADDRESS`.

Then listen to this address:

```powershell
.\server\server.exe "${WSL_ADDRESS}:8080"
```
You may get a pop-up from the Windows Firewall. Reject adding any exceptions for now.

Without stopping the server, start another shell on WSL and connect to the Windows host. In my case:
```bash
./client/client "http://${HOST_ADDRESS}:8080"
```
```
Error connecting to http://172.25.32.1:8080: Get "http://172.25.32.1:8080": context deadline exceeded (Client.Timeout exceeded while awaiting headers)
```

</details>

This does not work either.

### Server listens to all interfaces

<details> <summary> Show steps </summary>
On Windows:

```powershell
.\server\server.exe ":8080"
```
You may get a pop-up from the Windows Firewall. Reject adding any exceptions for now.

Without stopping the server, start another shell on WSL and connect to the Windows host. In my case:
```bash
./client/client "http://${HOST_ADDRESS}:8080"
```
```
Error connecting to http://172.25.32.1:8080: Get "http://172.25.32.1:8080": context deadline exceeded (Client.Timeout exceeded while awaiting headers)
```

</details>

This does not work either.

## With a firewall exception

We'll try the same experiment with a firewall exception.

<details> <summary> Show steps to allow the server through the firewall </summary>

1. Select the Start menu, type `Allow an app through Windows Firewall`, and select it from the list of results.
2. Select Change settings. You might be asked for an administrator password or to confirm your choice.
3. Search for `server.exe` in the list. If it is not there, add it:
   1. Click on `Allow another app`
   2. Enter the path to the server executable in this repository

</details>

### Server listens to localhost

<details> <summary> Show steps </summary>

With the firewall exception enabled, start the server on Windows:
```powershell
./server/server ":8080"
```
Run the client from within WSL:
```bash
$ ./client/client "http://${HOST_ADDRESS}:8080"
```
```
Error connecting to http://172.25.32.1:8080: Get "http://172.25.32.1:8080": context deadline exceeded (Client.Timeout exceeded while awaiting headers)
```

</details>

It does not work.


### Server listens to WSL network/all interfaces

With the firewall exception enabled, start the server on Windows with either of these commands:
```powershell
./server/server ":8080"
./server/server "${WSL_ADDRESS}:8080"
```
Run the client from within WSL:
```bash
./client/client "http://${HOST_ADDRESS}:8080"
```
```
Connected to http://172.25.32.1:8080: 200 OK
```

Success!

However, needing to enable a Firewall exception makes it a bad user experience for applications that need to talk across the Windows-WSL boundary.