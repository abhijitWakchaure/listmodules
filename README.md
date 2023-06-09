# listmodules

A cli tool to list the go modules (based on dirctory structure not on go.mod file)

The output of this program can be used as an input to [godiff](https://github.com/abhijitWakchaure/godiff) tool to find differences or common modules between two files.

## How to use

```text
listmodules [path to your vendor dir]
```

You can list the modules from the vendor folder in current directory by using command:

```bash
listmodules ./vendor
```
