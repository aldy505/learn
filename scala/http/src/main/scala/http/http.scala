package http

import org.scalatra._

// TODO: create a todo app lol

class http extends ScalatraServlet {

  get("/") {
    views.html.hello()
  }

}
