#!/bin/bash

# Скрипт для генерации отчёта о покрытии

set -e

echo "=== Проверка зависимостей ==="

# Добавляем pipx в PATH если его нет
export PATH="$HOME/.local/bin:$PATH"

# Проверяем и устанавливаем зависимости
check_and_install() {
    if ! command -v $1 &> /dev/null; then
        echo "Установка $1..."
        if [[ "$OSTYPE" == "linux-gnu"* ]]; then
            if command -v apt-get &> /dev/null; then
                sudo apt-get update
                sudo apt-get install -y $2
            fi
        fi
    else
        echo "$1 уже установлен"
    fi
}

# Проверяем необходимые утилиты
check_and_install cmake cmake
check_and_install g++ g++

# Проверяем Boost
if [[ ! -d "/usr/include/boost" ]] && [[ ! -d "/usr/local/include/boost" ]]; then
    echo "Установка Boost..."
    sudo apt-get install -y libboost-test-dev libboost-system-dev
else
    echo "Boost уже установлен"
fi

# Проверяем gcovr
if ! command -v gcovr &> /dev/null; then
    echo "Установка gcovr..."
    sudo apt-get install -y gcovr
else
    echo "gcovr уже установлен"
fi

echo "=== Настройка проекта ==="
mkdir -p build
cd build

echo "=== Генерация Makefile ==="
cmake .. -DCMAKE_BUILD_TYPE=Debug

echo "=== Сборка проекта (только тесты) ==="
# Собираем только тесты, игнорируем benchmark если есть ошибки
make tests -j$(nproc) || echo "Предупреждение: не удалось собрать все цели"

echo "=== Запуск тестов ==="
if [ -f "./tests" ]; then
    ./tests --log_level=test_suite --report_level=no
else
    echo "Ошибка: файл tests не найден"
    exit 1
fi

echo "=== Генерация отчёта о покрытии ==="
if command -v gcovr &> /dev/null; then
    echo "Используем gcovr для генерации отчёта..."
    gcovr -r .. --html --html-details -o coverage_report.html
    echo "HTML отчёт создан: build/coverage_report.html"
    
    # Также генерируем текстовый отчёт
    echo "=== Сводка покрытия кода ==="
    gcovr -r .. -s
    
    # Открываем отчёт в браузере
    if command -v xdg-open &> /dev/null; then
        xdg-open coverage_report.html &
    else
        echo "Отчёт доступен по адресу: file://$(pwd)/coverage_report.html"
    fi
else
    echo "Ошибка: gcovr не установлен"
    exit 1
fi

echo "=== Готово! ==="
echo "Проект успешно собран и протестирован."
echo "Отчёт о покрытии: build/coverage_report.html"