# cf-go-hello
Sample CF app in GO

## Pushing app to Cloud Foundry 
```
$ cf push cf-go-hello -b https://github.com/cloudfoundry/go-buildpack.git 
```

## Manifest.yml

```
---
applications:
  - name: cf-go-hello
    command: cf-go-hello 
    env:
      GOPACKAGENAME: cf-go-hello
```

