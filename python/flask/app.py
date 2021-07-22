from flask import Flask
from flask import request

app = Flask(__name__)

data = [
    {
        "id": 1,
        "author": "Foo",
        "text": "Buy food"
    },
    {
        "id": 2,
        "author": "NBar",
        "text": "Buy grocs"
    }
]

@app.route("/<int:id>", methods=['GET'])
def get_todo_by_id(id):
    for item in data:
        if item.get("id") == id:
            return item
            break
        else:
            continue
    return {"message": "sorry, item not found"}

@app.route("/")
def get_all_todos():
    return {"todos": data}

@app.route("/", methods=["POST"])
def insert_todo():
    body = request.get_json()
    data.append(body)
    return {"todo": body}