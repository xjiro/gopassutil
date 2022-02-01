# gopassutil

command line utility to hash and verify with passlib for go

## Usage

Generate Hash

```bash
> gopassutil hunter2
$s2$16384$8$1$S9FDCNeI2p0ZWMJjDYJuKqIJ$ALttuKoQr4Wpnk/1NSuNMymmfAElSu8Z1NguZyTHP6k=
```

Verify - Success

```bash
> gopassutil hunter2 '$s2$16384$8$1$S9FDCNeI2p0ZWMJjDYJuKqIJ$ALttuKoQr4Wpnk/1NSuNMymmfAElSu8Z1NguZyTHP6k='
1
```

Verify - Failure - Wrong Password
```bash
> gopassutil hunter3 '$s2$16384$8$1$S9FDCNeI2p0ZWMJjDYJuKqIJ$ALttuKoQr4Wpnk/1NSuNMymmfAElSu8Z1NguZyTHP6k='
0
invalid password
```


Verify - Failure - Invalid Verification Hash
```bash
> gopassutil hunter2 '$s2$16384$8$1$S9FDCNeI2p0ZWMJjDYJuKqIJ$ALttuKoQr4Wpnk/1NSuNMymmfAElSu8Z1NguZyTHP624='
0
illegal base64 data at input byte 44
```
```bash
> gopassutil.exe hunter2 'f'
0
unsupported scheme
```

## License
[MIT](https://choosealicense.com/licenses/mit/)
