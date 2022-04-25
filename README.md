# yamlier
a Super simple cli yaml editor

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
※ CAUTION: new yaml's Keys are ordered. 
