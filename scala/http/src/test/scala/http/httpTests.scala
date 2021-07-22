package http

import org.scalatra.test.scalatest._

class httpTests extends ScalatraFunSuite {

  addServlet(classOf[http], "/*")

  test("GET / on http should return status 200") {
    get("/") {
      status should equal (200)
    }
  }

}
