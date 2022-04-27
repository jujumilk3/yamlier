# yamlier
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
   1. `yamlier edit [file_path] [key] [value] | [output_filepath]` 

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
