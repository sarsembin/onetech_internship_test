# acmp_concurrent
Необходимо реализовать функцию *Difficulties*, которая на вход принимает список url задач и возвращает *map*, где ключем является url задачи, а значением -- сложность задачи.
Нахождение сложностей задач необходимо выполнять конкурентно.
Подсказки:
* concurrent map write