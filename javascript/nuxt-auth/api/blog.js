const {
  Router
} = require('express') 
const lowdb = require('lowdb') 
const FileAsync = require('lowdb/adapters/FileAsync')
const router = Router()

const blogdb = new FileAsync(__dirname + '/database/blogPost.json')

lowdb(blogdb).then(db => {
  router.get('/blog', (req, res) => {
    let data

    if (req.query.id) {
      data = db.get('data').find({id: parseInt(req.query.id)}).value()
    } else if (req.query.author) {
      data = db.get('data').find({author: req.query.author }).value()
    } else {
      data = db.get('data').value()
    }

    if (req.query.limit) {
      const limit = parseInt(req.query.limit)
      data = data.slice(0, limit)
    } else if (req.query.page) {
      const page = parseInt(req.query.page) * 10
      data = data.slice(page - 10, page)
    }

    if (req.query.order) {
      if (req.query.order == "desc") {
        data = data.reverse()
      }
    }

    res.status(200).json({
      status: 200,
      on: 'blog',
      data: data
    })
  })
})

module.exports = router