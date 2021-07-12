const express = require('express') 
const cookieParser = require('cookie-parser') 
const cors = require('cors') 
const formRoute = require('./form.js') 
const userRoute = require('./user.js') 
const blogRoute = require('./blog.js') 
const app = express()
app.set('json spaces', 2)
app.use(cookieParser())
app.use(cors())
app.use(express.json())
app.use(userRoute)
app.use(formRoute)
app.use(blogRoute)
// Export express app
module.exports = app

// Start standalone server if directly running
if (require.main === module) {
  const port = process.env.PORT || 8000
  app.listen(port, () => {
    console.log(`API server listening on port ${port}`)
  })
}
