# yamlier
<img width="271" alt="yamlier-logo" src="https://user-images.githubusercontent.com/41659814/165521454-60ff0a26-a552-4e68-8fab-00ec94320dd2.png">

### a Super simple cli yaml editor.
If you have to edit yaml file on CI/CD pipeline or other situations, just use it.

## install
`wget https://github.com/jujumilk3/yamlier/releases/download/v0.2.0/yamlier.{os}.{arcihtecture} -O ./yamlier`  
`chmod +x ./yamlier`  
Of course, if you want to use it globally, You have to move it into `bin` dir
that referenced by your account. 

### for example
1. MacOS(amd64)
   ```shell
   $ wget https://github.com/jujumilk3/yamlier/releases/download/v0.2.0/yamlier.darwin.amd64 -O ./yamlier
   chmod +x ./yamlier
   ./yamlier edit ./yamlier edit developer.yaml languages.perl poor ./developer-modified.yaml
   ```
2. github actions (github actions' ubuntu-latest image is amd64 architecture)
   ```yaml
   ...
   jobs:
     build:
       runs-on: ubuntu-latest
   ...
       - name: wget
         uses: wei/wget@v1
         with:
           args: https://github.com/jujumilk3/yamlier/releases/download/v0.2.0/yamlier.linux.amd64 -O yamlier

       - name: change file mod
         run: sudo chmod +x ./yamlier
       
       - name: run yamlier
         run: ./yamlier edit ./developer.yaml name changed_name ./developer-yamlier.yaml
   ...
   ```

## usage
1. yamlier --help
2. yamlier edit [flags]
   1. `yamlier edit --help`   
   2. `yamlier edit [file_path] [key] [value] | [output_filepath]` 
3. yamlier get [flags]
   1. `yamlier get --help`
   2. `yamlier get [file_path] [key]` 

## example
```yaml
# developer.yaml
name: Martin D'vloper
job: Developer
skill: Elite
foods:
  - Apple
  - Orange
  - Strawberry
  - Mango
languages:
  perl: Elite
  python: Elite
  pascal: Lame
education: |
  4 GCSEs
  3 A-Levels
  BSc in the Internet of Things
```
`yamlier edit ./developer.yaml name changed_name ./developer-yamlier.yaml`
```yaml
# developer-yamlier.yaml
education: |
  4 GCSEs
  3 A-Levels
  BSc in the Internet of Things
foods:
- Apple
- Orange
- Strawberry
- Mango
job: Developer
languages:
  pascal: Lame
  perl: Elite
  python: Elite
name: changed_name  # name changed
skill: Elite
```
※ CAUTION: new yaml's keys are ordered. 

Also, you can handle deeper values
```yaml
# developer.yaml
name: Martin D'vloper
job: Developer
skill: Elite
foods:
  - Apple
  - Orange
  - Strawberry
  - Mango
languages:
  perl: Elite
  python: Elite
  pascal: Lame
education: |
  4 GCSEs
  3 A-Levels
  BSc in the Internet of Things
```
`yamlier edit ./developer.yaml languages.perl Poor ./developer-yamlier.yaml`
```yaml
# developer-yamlier.yaml
education: |
  4 GCSEs
  3 A-Levels
  BSc in the Internet of Things
foods:
- Apple
- Orange
- Strawberry
- Mango
job: Developer
languages:
  pascal: Lame
  perl: Poor  # languages.perl changed
  python: Elite
name: Martin D'vloper
skill: Elite
```

## build tips
If there's no binary file can use at your OS and architecture, just read it and build it yourself.  
Or you can request to me by register issue.
1. `git clone git@github.com:jujumilk3/yamlier.git`
2. `cd yamlier`
3. `go mod tidy` -> install packages and code described at `go.mod`
4. `go tool dist list` -> chk dist list   
   ```shell
   # os / architecture
   aix/ppc64
   android/386
   android/amd64
   android/arm
   android/arm64
   darwin/amd64
   darwin/arm64
   dragonfly/amd64
   freebsd/386
   freebsd/amd64
   freebsd/arm
   freebsd/arm64
   illumos/amd64
   ios/amd64
   ios/arm64
   js/wasm
   linux/386
   linux/amd64
   linux/arm
   linux/arm64
   linux/mips
   linux/mips64
   linux/mips64le
   linux/mipsle
   linux/ppc64
   linux/ppc64le
   linux/riscv64
   linux/s390x
   netbsd/386
   netbsd/amd64
   netbsd/arm
   netbsd/arm64
   openbsd/386
   openbsd/amd64
   openbsd/arm
   openbsd/arm64
   openbsd/mips64
   plan9/386
   plan9/amd64
   plan9/arm
   solaris/amd64
   windows/386
   windows/amd64
   windows/arm
   windows/arm64
   ```   
5. `env GOOS=target-OS GOARCH=target-architecture go build yamlier`
