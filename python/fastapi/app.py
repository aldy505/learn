from fastapi import FastAPI
from typing import Optional
from pydantic import BaseModel

app = FastAPI()

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

class Todo(BaseModel):
  id: int
  author: str
  text: str

@app.get("/")
def get_all_todos():
    return {"todos": data}

@app.get('/{id}')
def get_todo_by_id(id):
  for item in data:
    if item.get("id") == int(id):
      return item
    else:
      continue
  return {"message": "todo not found"}

@app.post("/")
def add_todo(todo: Todo):
  data.append(todo)
  return {"message": "todo added", "todo": todo}