# Сервис для управления IoT-системой умного дома

<!-- ### Идея проекта
Идея проекта заключается в создании сервиса, специализированного на сборе данных с IoT-устройств и их последующем анализе, что позволяяет  эффективно контролировать и автоматизировать различные аспекты своей повседневной среды, создавая удобные условия для комфортной жизни. -->
### Предметная область
Интернет вещей (IoT) в контексте умного дома представляет собой сеть физических устройств, оборудованных датчиками и другими технологиями, которые могут взаимодействовать между собой и с облачными платформами для автоматизации и улучшения различных аспектов повседневной жизни.
### Краткий анализ аналогичных решений по минимум 3 критериям
|              | Анализ и статистика                 | Визуализация | Многопользовательский доступ к устройству| 
|--------------|---------------------------|------------|-----------------------------|
| Apple Homekit | -       | -          | +                           | 
| Intel IoT Platform | -  | +          | -                           |
| MTS IoT HUB         | +      | +          | -                           |    
| Предполагаемое решение  | + | + | + |
### Краткое обоснование целесообразности и актуальности проекта
Данный сервис оправдывается растущим спросом на интегрированные технологии для повышения комфорта, безопасности и энергоэффективности в бытовых условиях. Актуальность и целесообразность проекта подчеркиваются повышенной автоматизацией жилья, что позволяет пользователям эффективно управлять своими ресурсами и обеспечивает современный образ жизни. IoT-система умного дома не только оптимизирует рутинные задачи, но также способствует повышению уровня безопасности и экономии энергии.
### Краткое описание акторов
Зарегистрированный пользователь – это лицо, которому предоставляется возможнсоть содать дом.
Новый пользователь – это лицо, ещё не зарегистрировавшееся на IoT-платформе, которому предоставляется возможность пройти процедуру регистрации и воспользоваться всеми доступными функциями.
Владелец - основной пользователь системы, обладающий полным спектром функционала для управления своим домом.
Участник дома - пользователь, принадлежащий к дому, с ограниченным, но индивидуализированным доступом к функциям системы в соответствии с назначенным уровнем привилегий.
### Use-case диаграмма
![iot](/img/iot.png)
### ER-диаграмма акторов
![er](/img/er.png)
### Диаграмм БД
![db](/img/newDB.png)
### Пользовательские сценарии
Новый пользователь:
* зарегистрироваться: ввести почту и пароль, восстановить пароль
Зарегистрированный пользователь:
* создать дом
Владелец дома:
* добавить устройство
* удалить устройство
* добавить другого пользователя в дом
* определить уровень доступа другого пользователя к дому
* просмотреть статстику
* управление устройствами
* создать дом
* удалить дом
Участник дома:
* просмотреть статистику
* управление устройствами
* создать дом
### Формализация ключевых бизнес-процессов
![busmodel](/img/busmodel.png)
### Тип приложения и технологический стек
Тип приложения: Web MPA
Технологический стек: Golang, HTML, CSS, PostgreSQL, JS
### Верхнеуровневое разбиение на компоненты
![busmodel](/img/partition.png)
