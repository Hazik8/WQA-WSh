# WQA & WSh

**WQA (Win Quick Application)** - язык приложений, система пакетов и среда выполнения на основе WQBC, созданные **Win Stydio**.

**WSh (Win Shell)** - командная оболочка для управления и запуска WQA-приложений и программ Windows.

WQ, WQA и WSh являются технологиями и проектами **Win Stydio**. Название **WinDroid** используется тогда, когда речь идёт непосредственно об экосистеме WinDroid.

## Версии

* **WQA:** 1.2.0
* **WSh:** 1.0.0
* **Среда выполнения:** WQBC
* **Архитектуры:** x86_64, arm64

---

# WQA 1.2.0

WQA - небольшой командно-ориентированный язык для создания приложений.

В WQA 1.2.0 доступны:

* переменные с `wet`;
* вывод с `print`;
* арифметика;
* арифметические выражения;
* условия;
* циклы;
* повторения;
* функции;
* возвращаемые значения функций;
* пользовательский ввод;
* очистка консоли;
* задержки;
* текущая дата и время;
* завершение программы.

### Пример

```wq
wqa

wet name = "WQA"

print "Привет!"
print name

winput answer

print "Ваш ответ:"
print answer

wtime
wdate
```

## Типы файлов

| Расширение | Назначение           |
| ---------- | -------------------- |
| `.wq`      | исходный код WQA     |
| `.wqbc`    | байткод WQBC         |
| `.wqa`     | пакет приложения WQA |

## Исходный файл WQA

Каждый файл `.wq` должен начинаться с точного заголовка:

```wq
wqa
```

Пример:

```wq
wqa

wet x = 10

wif x > 5
    print "YES"
else
    print "NO"
endwif
```

## Арифметические выражения

WQA 1.2.0 поддерживает:

```text
+    сложение
-    вычитание
*    умножение
/    обычное деление
//   целочисленное деление
%    остаток от деления
```

Поддерживаются скобки, унарный минус, целые и десятичные числа.

```wq
wqa

wet x = (10 + 5) * 2
wet y = 10 / 3
wet z = 10 // 3
wet r = 10 % 3
wet n = -25
wet f = 5.5 * 2

print x
print y
print z
print r
print n
print f
```

Старые команды арифметики `add`, `sub`, `mul` и `div` также поддерживаются.

---

# Сборка WQA

Сборка проекта:

```powershell
wqa build .
```

Сборка проекта из другой директории:

```powershell
wqa build .\TestProject
```

Результатом будет пакет:

```text
TestProject.wqa
```

Запуск пакета:

```powershell
wqa run .\TestProject.wqa
```

---

# WQBC

WQBC - байткодная среда выполнения WQA-приложений.

Путь выполнения:

```text
.wq
  ↓
Лексер
  ↓
Парсер
  ↓
Компилятор
  ↓
WQBC
  ↓
.wqa
  ↓
Среда выполнения WQBC
  ↓
Приложение
```

---

# WSh 1.0.0

WSh (Win Shell) предоставляет командную оболочку для управления приложениями и системой.

WSh умеет:

* работать с файлами и каталогами;
* запускать программы Windows;
* запускать WQA-приложения;
* управлять установленными приложениями;
* искать приложения в репозитории;
* устанавливать и обновлять приложения;
* хранить историю команд;
* создавать псевдонимы команд;
* управлять переменными окружения.

## Команды WSh

### Оболочка

```text
help, ?
version
--version
-v
clear, cls
history
alias
unalias
exit, quit
```

### Файлы

```text
pwd, gl
cd <path>
dir, ls, gci
cat, type, gc <file>
mkdir, newdir <name>
new <file>
del, rm, remove <file>
cp, copy <source> <destination>
mv, move <source> <destination>
write <file> <text>
append <file> <text>
test <path>
```

### Система

```text
ps, process
kill <pid>
env
env <name>
set <name> <value>
unset <name>
date
which <command>
echo <text>
```

### Запуск

```text
run <program>
wqa <command>
```

### Пакеты и приложения

```text
repo
repo update
install <package|app>
update <app>
update all
remove-app <app>
list
info <app>
search [name]
```

## Пример WSh

```text
wsh [C:\WQA]> wqa build
[OK] Created: TestProject.wqa

wsh [C:\WQA]> wqa run TestProject.wqa
[INFO] Loading: TestProject.wqa
[INFO] Starting WQBC Runtime
```

WSh также может запускать обычные программы Windows:

```text
wsh [C:\WQA]> run notepad.exe
[OK] Process started: C:\Windows\system32\notepad.exe
```

## Репозиторий приложений

WSh поддерживает репозиторий WQA-приложений.

Пример:

```text
repo update
search Calculator
info windroid.Calculator
install windroid.Calculator
list
```

Перед установкой приложение загружается и проверяется по SHA-256.

---

# Документация

Полная спецификация языка WQA находится в репозитории документации:

[WQA-WSh-Docs](https://github.com/Hazik8/WQA-WSh-Docs)

Основной документ:

[WQA_LANGUAGE.md](https://github.com/Hazik8/WQA-WSh-Docs/WQA_LANGUAGE_RU.md)

---

# Статус проекта

WQA и WSh активно развиваются **Win Stydio**.

**Текущая версия WQA: 1.2.0**

**Текущая версия WSh: 1.0.0**

WSh 1.0.0 ориентирован на стабильную командную работу в Windows и интеграцию с экосистемой WQA.

---

# Лицензия

Проект распространяется по лицензии MIT.

Подробнее см. в файле `LICENSE`.

---

# Win Stydio

WQ, WQA и WSh - технологии и проекты **Win Stydio**.
