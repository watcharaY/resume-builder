# Resume Builder

golang based program to generate resume in pdf format from html file.

## Features

- generate pdf from html

## Dependency

- [Godotenv](https://github.com/joho/godotenv) for get value from .env
- [Gotenberg](https://github.com/starwalkn/gotenberg-go-client/v8) for generate pdf from html

## Installation

1. get a html file for using as input naming it ```resume.html```

2. set env value for file path

3. make sure you have [taskfile](https://taskfile.dev/installation/) to perform installation then we can use command below

```bash
task run
```
