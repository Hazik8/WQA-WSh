# WQA Language 1.5

WQA (WinDroid Quick Application) —
язык приложений для платформы WinDroid.

## Создание приложения

Файл программы:

main.wqa


## Команды


## print

Вывод текста.

Пример:

print Hello


Результат:

Hello



## set

Создание переменной.

Пример:

set name = WinDroid



## load

Вывод переменной.

Пример:

set version = 1.5

load version



## Математика


add

Сложение:

add a b result


sub

Вычитание:

sub a b result


mul

Умножение:

mul a b result


div

Деление:

div a b result



## Условия


if

Проверка условия:

if a > b

print OK

endif



else

Альтернативная ветка:

if a > b

print Yes

else

print No

endif



## Циклы


loop

Повторение:

loop count

print Hello

endloop



## Завершение


exit

Закрывает приложение.


# Сборка приложения


wqa build


# Запуск


wqa run App.wqa


# Установка


wqa install App.wqa