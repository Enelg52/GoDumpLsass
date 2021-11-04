<p align="center">
  <a href="" rel="noopener">
 <img width=200px height=200px src="https://cdn-icons.flaticon.com/png/512/3065/premium/3065741.png?token=exp=1636056750~hmac=b65861715a05bdc78a609f295ff99b01"> </a>
</p>

# GoDumpLsass

------
GoDumpLsass is a simple tool that can dump lsass without to get caught by Windows Defender.
## Releases

https://github.com/Enelg52/GoDumpLsass/releases

## Usage

```txt
.\GoDumpLsass.exe -p [path]
```

## Example
```txt
.\GoDumpLsass.exe -p "C:\Users\enelg\Downloads\"
[-] Get lsass process id
[+] Lsass pid : 908
[-] Dump process
[+] Process dumped
[+] The dump is under C:\Users\enelg\Downloads\lsass.dmp
```

## Contributing
Pull requests are welcome !