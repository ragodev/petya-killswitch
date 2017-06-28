# PETYA KillSwitch

> A killswitch to stop petya from encrypting your disk

Security Researchers ([PT Security][ptsecurity], [TrustedSec][trustedsec], ...etc) seem to have identified a kill-switch that prevents `petya` ransomware from proceeding through its encryption process <sup>[1][1]</sup>.

Simply, all that is needed are 3 files (`perfc`, `perfc.dll`, and `perfc.dat`) to already exist on the Windows machine, under `C:\Windows`, with **READONLY** permissions.

This Go application does just that.

## Install

### Download from Github Releases

- [Download Link][download-link]

### If you have Go installed

- `go get github.com/petermbenjamin/petya-killswitch`
- `cd %GOPATH%/src/github.com/petermbenjamin/petya-killswitch && go build`

## Usage

- Launch `cmd.exe` as **ADMIN**.
- Execute petya-killswitch:
    ```sh
    petya-killswitch.exe
    ```
- Done.


## LICENSE

MIT &copy; 2017 [Peter Benjamin][pb]

[1]: https://www.bleepingcomputer.com/news/security/vaccine-not-killswitch-found-for-petya-notpetya-ransomware-outbreak/
[ptsecurity]: https://twitter.com/ptsecurity/status/879766638731591680
[trustedsec]: https://twitter.com/HackingDave/status/879779361364357121
[download-link]: https://github.com/petermbenjamin/petya-killswitch/releases/download/1.0/petya-killswitch.exe
[pb]: https://petermbenjamin.github.io