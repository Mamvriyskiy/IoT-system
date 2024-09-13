import pandas as pd
import matplotlib.pyplot as plt
import numpy as np
import json
import sys

json_str = sys.argv[1]

# Преобразуем JSON обратно в список
data = json.loads(json_str)
print(data)

data = {
    "Количество записей": data[0],
    "Без индексов": data[1],
    "С индексами": data[2],
}

# data = {
#     "Количество записей": [10000, 20000, 30000, 40000, 50000, 60000, 70000, 80000, 90000, 100000],
#     "Без индексов": [4943, 5930, 6044, 6515, 7145, 10028, 10507, 11305, 13005, 14542],
#     "С индексами": [2188, 2266, 2735, 2819, 3128, 4472, 4524, 4616, 4628, 4827],
# }

# Получение аргументов
arguments = sys.argv

# Вывод аргументов
print("Аргументы, переданные в скрипт:")
for i, arg in enumerate(arguments):
    print(f"Аргумент {i}: {arg}")

markers = ['o', 's', '^', 'D', 'x']

# Создание DataFrame из данных
df = pd.DataFrame(data)

# Отрисовка графиков
plt.figure(figsize=(12, 6))

k = 0
# Добавление графиков для каждой колонки
for i, column in enumerate(df.columns[1:]):
    plt.plot(df['Количество записей'], df[column], label=column, marker = markers[k])

# Настройка легенды и заголовка
plt.legend(loc='upper left')
plt.title('График зависимости времени выполнения от размерности матрицы')
plt.xlabel('Размерность')
plt.ylabel('Время выполнения (мкc)')

# Отображение графика
plt.savefig('graph.png')
