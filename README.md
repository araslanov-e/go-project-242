### Hexlet tests and linter status:
[![Actions Status](https://github.com/araslanov-e/go-project-242/actions/workflows/hexlet-check.yml/badge.svg)](https://github.com/araslanov-e/go-project-242/actions)
[![Go tests](https://github.com/araslanov-e/go-project-242/actions/workflows/go-tests.yml/badge.svg)](https://github.com/araslanov-e/go-project-242/actions/workflows/go-tests.yml)

# hexlet-path-size

`hexlet-path-size` — CLI-утилита для вывода размера файла или директории.

Команда принимает необязательный аргумент - путь до файла или директории:

```bash
hexlet-path-size [options] [file or directory]
```

Если путь не передан, утилита считает размер текущей директории (`.`).
По умолчанию размер выводится в байтах. Для человекочитаемого формата,
рекурсивного подсчёта директорий и учёта скрытых файлов используются опции.

## Установка

```bash
git clone git@github.com:araslanov-e/go-project-242.git
cd go-project-242
make build
```

После сборки исполняемый файл будет доступен по пути `bin/hexlet-path-size`.

## Использование

```bash
./bin/hexlet-path-size
./bin/hexlet-path-size testdata/file49kb
./bin/hexlet-path-size -H testdata/file49kb
./bin/hexlet-path-size -r -H testdata
./bin/hexlet-path-size -a -H testdata
```

## Флаги

- `-H`, `--human` — выводит размер в человекочитаемом формате, например `49.0KB`.
- `-r`, `--recursive` — рекурсивно считает размер директорий, включая файлы во вложенных директориях.
- `-a`, `--all` — учитывает скрытые файлы и директории (имена которых начинаются с точки).

## Демонстрация

[![asciicast](https://asciinema.org/a/xM3UO3fuxWFbPElg.svg)](https://asciinema.org/a/xM3UO3fuxWFbPElg)
