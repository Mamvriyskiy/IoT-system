import matplotlib.pyplot as plot
import numpy as np


with open('result.txt', 'r') as file:
    data = file.read()

# Удаляем лишние символы и разделяем массив на подмассивы
data = data.replace('[','').replace(']','').replace('"','').replace(',',' ').split('[')
numbers = list(map(int, data[0].split()))

arrays = [[] for _ in range(12)]

# Записываем числа в соответствующие массивы
for i, num in enumerate(numbers):
    arrays[i % 12].append(num)

# Проверка длины массивов
if len(arrays[0]) != 100:
    print(f"Warning: массив arrays[0] имеет длину {len(arrays[0])}, ожидалась длина 100")

# Выводим результат
size = list(range(1, 101))

# 1 -- gorm add NsPerOp
# 2 -- gorm add AllocsPerOp
# 3 -- gorm add AllocedBytesPerOp
# 4 -- gorm get NsPerOp
# 5 -- gorm get AllocsPerOp
# 6 -- gorm get AllocedBytesPerOp
# 7 -- sqlx add NsPerOp
# 8 -- sqlx add AllocsPerOp
# 9 -- sqlx add AllocedBytesPerOp
# 10 -- sqlx get NsPerOp
# 11 -- sqlx get AllocsPerOp
# 12 -- sqlx get AllocedBytesPerOp

# add

plot.ylabel("NsPerOp")
plot.xlabel("Номер бенчмарка")
plot.grid(True)
plot.plot(size, arrays[0], color="darkmagenta", label='gorm', marker='^')
plot.plot(size, arrays[6], color="gold", label='sqlx', marker='*')
plot.legend(["gorm", "sqlx"])
plot.title('Add home NsPerOp')
plot.savefig('addhomeNsPerOp.svg')
plot.show()

plot.ylabel("AllocsPerOp")
plot.xlabel("Номер бенчмарка")
plot.grid(True)
plot.plot(size, arrays[1], color="darkmagenta", label='gorm', marker='^')
plot.plot(size, arrays[7], color="gold", label='sqlx', marker='*')
plot.legend(["gorm", "sqlx"])
plot.title('Add home AllocsPerOp')
plot.savefig('addhomeAllocsPerOp.svg')
plot.show()

plot.ylabel("AllocedBytesPerOp")
plot.xlabel("Номер бенчмарка")
plot.grid(True)
plot.plot(size, arrays[2], color="darkmagenta", label='gorm', marker='^')
plot.plot(size, arrays[8], color="gold", label='sqlx', marker='*')
plot.legend(["gorm", "sqlx"])
plot.title('Add home AllocedBytesPerOp')
plot.savefig('addhomeAllocedBytesPerOp.svg')
plot.show()

# get

plot.ylabel("NsPerOp")
plot.xlabel("Номер бенчмарка")
plot.grid(True)
plot.plot(size, arrays[3], color="darkmagenta", label='gorm', marker='^')
plot.plot(size, arrays[9], color="gold", label='sqlx', marker='*')
plot.legend(["gorm", "sqlx"])
plot.title('Get home NsPerOp')
plot.savefig('gethomeNsPerOp.svg')
plot.show()

plot.ylabel("AllocsPerOp")
plot.xlabel("Номер бенчмарка")
plot.grid(True)
plot.plot(size, arrays[4], color="darkmagenta", label='gorm', marker='^')
plot.plot(size, arrays[10], color="gold", label='sqlx', marker='*')
plot.legend(["gorm", "sqlx"])
plot.title('Get home AllocsPerOp')
plot.savefig('gethomeAllocsPerOp.svg')
plot.show()

plot.ylabel("AllocedBytesPerOp")
plot.xlabel("Номер бенчмарка")
plot.grid(True)
plot.plot(size, arrays[5], color="darkmagenta", label='gorm', marker='^')
plot.plot(size, arrays[11], color="gold", label='sqlx', marker='*')
plot.legend(["gorm", "sqlx"])
plot.title('Get home AllocedBytesPerOp')
plot.savefig('gethomeAllocedBytesPerOp.svg')
plot.show()

# Гистограмма распределения
for index, label in zip([0, 6], ['gorm', 'sqlx']):
    plot.hist(arrays[index], bins=20, alpha=0.5, label=label)
    plot.xlabel('Значения')
    plot.ylabel('Частота')
    plot.title(f'Гистограмма для {label}')
    plot.legend()
    plot.savefig(f'histogram_{label}.svg')
    plot.show()

import numpy as np
import matplotlib.pyplot as plot  # Make sure to import the plotting library

percentiles = [0.5, 0.75, 0.9, 0.95, 0.99]
percentile_results = {label: [np.percentile(arr, p * 100) for p in percentiles] for label, arr in zip(['gorm', 'sqlx'], [arrays[0], arrays[6]])}

for i in range(1):  # для NsPerOp, AllocsPerOp, AllocedBytesPerOp
    plot.ylabel("Значение")
    plot.xlabel("Номер бенчмарка")
    plot.grid(True)
    
    plot.plot(size, arrays[i], color="darkmagenta", label='gorm', marker='^')
    plot.plot(size, arrays[i + 6], color="gold", label='sqlx', marker='*')
    
    # Добавление линий для перцентилей и аннотаций
    for j, p in enumerate(percentile_results['gorm']):
        plot.axhline(y=p, color='darkviolet', linestyle='--', linewidth=1)
        plot.text(x=size[-1], y=p, s=f'{percentiles[j]*100:.0f}%', color='darkviolet', va='bottom')
    
    for j, p in enumerate(percentile_results['sqlx']):
        plot.axhline(y=p, color='orange', linestyle='--', linewidth=1)
        plot.text(x=size[-1], y=p, s=f'{percentiles[j]*100:.0f}%', color='orange', va='bottom')

    # Updating the legend to show unique labels
    handles, labels = plot.gca().get_legend_handles_labels()
    by_label = dict(zip(labels, handles))  # Remove duplicates
    plot.legend(by_label.values(), by_label.keys())

    plot.title(f'График для {"NsPerOp" if i == 0 else "AllocsPerOp" if i == 1 else "AllocedBytesPerOp"}')
    plot.savefig(f'graph_{["NsPerOp", "AllocsPerOp", "AllocedBytesPerOp"][i]}.svg')
    plot.show()