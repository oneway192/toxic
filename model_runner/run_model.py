import sys

# Пример без модели, просто проверка по ключевому слову
text = sys.argv[1]

if "ошибка" in text.lower():
    print("CORRECTED")
else:
    print("OK")

