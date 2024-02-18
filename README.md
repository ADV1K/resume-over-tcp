# Resume Over TCP
A smol TCP server that serves my resume.  

[![asciicast](https://asciinema.org/a/641179.svg)](https://asciinema.org/a/641179)

## Connecting to the server

If you are using a UNIX-like system, you can connect to the server using `nc` (netcat).

```shell
nc advik.lol 6969
```  

Or use telnet
```shell
telnet advik.lol 6969
```

If you are using Windows, you can use PuTTY or any other telnet client (set the connection type to raw).
Alternatively, you can use this PowerShell one-liner. Unfortunately, the colors won't be displayed correctly.
```powershell
$client = [System.Net.Sockets.TcpClient]::new(); $client.Connect("advik.lol", 6969); $reader = [System.IO.StreamReader]::new($client.GetStream()); while($true){$data=$reader.Read();if($data-eq-1){break};Write-Host -NoNewline ([char]$data)}; $client.Close()
```

If you get an error, you might need to change your execution policy.  
Run the following command in an elevated PowerShell prompt.
```powershell
Set-ExecutionPolicy Unrestricted
```

## Building and running the server
```shell
docker compose up --build
```

Your application will be available at http://localhost:6969.

<p align="center"><img src="https://raw.githubusercontent.com/catppuccin/catppuccin/main/assets/footers/gray0_ctp_on_line.svg?sanitize=true" /></p>
