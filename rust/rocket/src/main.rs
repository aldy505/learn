use rocket::{response::{content, status}, serde::{Serialize, Deserialize, json::Json}};
#[macro_use] extern crate rocket;

#[get("/")] // Get all data
fn index() -> Json([Todo]) {
    Json([
    // You know what, I gave up.
        
    ])
}

#[get("/<id>")] // Get data by ID
fn get_by_id(id: i32) -> Json(Todo) {
    // You know what, I gave up.
}

/* Insert data */
#[derive(Deserialize)]
struct Todo<'r> {
    pub id: i32,
    pub author: &'r str,
    pub text: &'r str
}

#[put("/", data = "<todo>")]
fn insert_data(todo: Json<Todo<'_>>) -> status::Accepted<content::Json<&'static str>> {
    return Json(Todo{
        id: todo.id,
        author: todo.author,
        text: todo.text
    })
}

// Update data
#[patch("/", data = "<todo>")]

// Delete data by ID
#[delete("/<id>")]

#[launch]
fn rocket() -> _ {
    let data: &[Todo] = [
        Todo{
            id: 1,
            author: "Foo",
            text: "Buy groceries"
        },
        Todo{
            id: 2,
            author: "Baar",
            text: "Get Food"
        }
    ];
    rocket::build().mount("/", routes![index])
}