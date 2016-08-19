# goprintuser

A test project to demonstrate SUID and GUID... Here is a detailed documentation about it: [Access Modes](http://www.ibm.com/developerworks/library/l-lpic1-104-5/index.html#modes)

## Install

```
go get github.com/juame/goprintuser
```

## Usage

```bash
$GOPATH/bin/goprintuser
```

## Demo

Copy the file to /bin for the demo:

```bash
cp $GOPATH/bin/goprintuser /bin/goprintuser && chown root:root /bin/goprintuser
```

Set SUID and GUID:

```bash
chmod u+s /bin/goprintuser
chmod g+s /bin/goprintuser
```

goprintuser as root:
```bash
root@ubuntuvm:~# goprintuser
current user: euid=0 egid=0 username=root
```

goprintuser as julian:
```bash
julian@ubuntuvm:~$ goprintuser
current user: euid=0 egid=0 username=root
```

# Bash Scripts with SUID/GUID

Please read this blog article about this: [Why Bash is like that: suid
](https://www.vidarholen.net/contents/blog/?p=30)
