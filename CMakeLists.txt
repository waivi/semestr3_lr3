cmake_minimum_required(VERSION 3.10)
project(DataStructuresProject)

set(CMAKE_CXX_STANDARD 17)
set(CMAKE_CXX_STANDARD_REQUIRED ON)

# Включаем отладочную информацию для покрытия
set(CMAKE_CXX_FLAGS "${CMAKE_CXX_FLAGS} -g -O0 --coverage")
set(CMAKE_EXE_LINKER_FLAGS "${CMAKE_EXE_LINKER_FLAGS} --coverage")

# Находим Boost
find_package(Boost 1.65.0 REQUIRED COMPONENTS unit_test_framework)

# Основные файлы
set(COMMON_SOURCES
    array.cpp
    SList.cpp
    DList.cpp
    Stack.cpp
    Queue.cpp
    Tree.cpp
    Hash_map.cpp
)

set(COMMON_HEADERS
    array.h
    SList.h
    DList.h
    Stack.h
    Queue.h
    Tree.h
    Hash_map.h
    serialisation.h
    benchmark.h
)

# Основной исполняемый файл
add_executable(database
    main.cpp
    ${COMMON_SOURCES}
)

# Тестовый исполняемый файл
add_executable(tests
    tests.cpp
    ${COMMON_SOURCES}
)

# Подключаем Boost к тестам
target_link_libraries(tests
    PRIVATE
    Boost::unit_test_framework
)

# Создаем benchmark только если есть main_benchmark.cpp
if(EXISTS "${CMAKE_CURRENT_SOURCE_DIR}/main_benchmark.cpp")
    add_executable(benchmark
        main_benchmark.cpp
        benchmark.cpp
        ${COMMON_SOURCES}
    )
    
    add_custom_target(run_benchmark
        COMMAND ./benchmark
        DEPENDS benchmark
        WORKING_DIRECTORY ${CMAKE_CURRENT_BINARY_DIR}
        COMMENT "Running benchmarks..."
    )
else()
    message(STATUS "main_benchmark.cpp not found - skipping benchmark target")
endif()

# Цель для запуска тестов
add_custom_target(run_tests
    COMMAND ./tests --log_level=test_suite --report_level=no
    DEPENDS tests
    WORKING_DIRECTORY ${CMAKE_CURRENT_BINARY_DIR}
    COMMENT "Running all tests..."
)

# Цель для генерации отчёта о покрытии
find_program(GCOVR gcovr)
if(GCOVR)
    add_custom_target(coverage_report
        COMMAND ${GCOVR} -r ${CMAKE_SOURCE_DIR} --html --html-details -o ${CMAKE_BINARY_DIR}/coverage_report.html
        COMMAND ${GCOVR} -r ${CMAKE_SOURCE_DIR} -s
        DEPENDS tests run_tests
        WORKING_DIRECTORY ${CMAKE_CURRENT_BINARY_DIR}
        COMMENT "Generating coverage report..."
    )
else()
    message(WARNING "gcovr not found - coverage report target disabled")
endif()