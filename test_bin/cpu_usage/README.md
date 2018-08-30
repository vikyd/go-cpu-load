# Github
https://github.com/vikyd/cpu_usage

# cpu_usage

Get all CPU usage by process id(pid) on Windows(Linux is not supported).

`usage`: percentage of all cores of CPU

# Build

CMD command:

```cmd
"C:\MinGW\bin\g++" cpu_usage_win.cpp -o cpu_usage_win.exe
```

# Run

```cmd
cpu_usage_win.exe yourPid
REM: example: cpu_usage_win.exe 46400
REM: output(35%): 35
```

- if pid not exist, then exist with code -1

# How it works

CPU load percentage = process time delta / system time delta \* 100


# Thangs
- http://www.cppblog.com/cppopp/archive/2012/08/24/188102.aspx