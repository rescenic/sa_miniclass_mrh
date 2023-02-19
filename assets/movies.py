import json

with open('movies.json', 'r', encoding='utf-8') as f:
    data = json.load(f)

output = []
for movie in data:
    output.append({
        "id": movie["id"],
        "title": movie["title"],
        "overview": movie["overview"],
        "poster": movie["poster"]
    })

with open('moviesv2.json', 'w', encoding='utf-8') as f:
    json.dump(output, f, ensure_ascii=False, indent=4)
