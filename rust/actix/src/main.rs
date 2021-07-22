use actix_web::{App, HttpRequest, HttpResponse, HttpServer, Responder, get, post, web};
use serde::{Deserialize, Serialize};

#[derive(Serialize, Deserialize)]
struct Todo {
    id: i32,
    author: String,
    text: String
}

#[derive(Serialize)]
struct TodoResponse<'a> {
    message: String,
    count: &'a usize,
    todo: &'a Todo
}

struct AppState {
    todos: Vec<Todo>,
}

#[get("/{id}")]
async fn get_todos(web::Path(id): web::Path<i32>, data: web::Data<AppState>) -> impl Responder {
    if id > 0 {
        let found = data.todos.iter().position(|r| r.id == id).unwrap();
        return HttpResponse::Ok().json(TodoResponse{
            message: String::from("ok"),
            count: &data.todos.len(),
            todo: &data.todos[found]
        });
    }
    return HttpResponse::Ok().json(&data.todos)
}

#[post("/")]
async fn add_todo(body: web::Json<Todo>, data: web::Data<AppState>) -> impl Responder {
    &data.todos.insert(data.todos.len() + 1, body as Todo);
    return HttpResponse::Ok().json(&data.todos);
}

#[actix_web::main]
async fn main() -> std::io::Result<()> {
    HttpServer::new(|| {
        App::new()
            .data(AppState{
                todos: vec![
                    Todo{
                        id: 1,
                        author: String::from("Foo"),
                        text: String::from("Water plants")
                    },
                    Todo{
                        id:2,
                        author: String::from("Baar"),
                        text: String::from("Buy food"),
                    }
                ]
            })
            .service(get_todos)
            .route("/", web::post().to(add_todo))
            .route("/hey", web::get().to(manual_hello))
    })
    .bind("127.0.0.1:8080")?
    .run()
    .await
}