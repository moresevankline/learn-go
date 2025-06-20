<!-- Install Go -->
# Install Go

Step 1: Download Go in <https://go.dev/doc/install>

Step 2: Open terminal

Step 3: Go to the directory of the downloaded file

```bash
cd ~/Downloads
```

Step 4: Delete /usr/local/go folder

```bash
sudo rm -rf /usr/local/go
```

Step 5: Extract archive

```bash
tar -C /usr/local -xzf go1.24.4.linux-amd64.tar.gz
```

Step 6: Open .profile

```bash
sudo nano $HOME/.profile
```

Step 7: Add path in the last line

```bash
export PATH=$PATH:/usr/local/go/bin
```

Step 8: Apply changes

```bash
$HOME/.profile
```

Step 9: Verify installed Go

```bash
go version
```
