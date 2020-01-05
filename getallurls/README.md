# getallurls
Fetch known URLs from AlienVault's [Open Threat Exchange](https://otx.alienvault.com), the Wayback Machine, and Common Crawl. Originally built as a microservice.

### usage:
```
▻ printf 'example.com' | gau
```

or

```
▻ gau example.com
```

### install:

```
▻ git clone https://github.com/lc/hacks && cd hacks/getallurls
▻ go build -o $GOPATH/bin/gau getallurls.go
```

or

```
▻ go get -u github.com/lc/hacks/getallurls/ && mv $GOPATH/bin/getallurls $GOPATH/bin/gau
```

## Credits:
Thanks @tomnomom for [waybackurls](https://github.com/tomnomnom/waybackurls)!
