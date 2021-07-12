const {
  Router
} = require('express')
const jwt = require('jsonwebtoken')
const lowdb = require('lowdb')
const FileAsync = require('lowdb/adapters/FileAsync')
const router = Router()

const APIKEY = `2spX^cJdJ8wArpEei!AQ4nH%Y&%5s*bdmNoZuMPkyrqLjO*lO8t3ZY^EvSX^Nx!r#9t!BtXw0S%j5ranpR*7mL1&`
const usersdb = new FileAsync(__dirname + '/database/users.json')

lowdb(usersdb).then(db => {
  router.post('/userdata', (req, res) => {
    if (req.body.key == APIKEY) {
      const searchData = db.get('data').find({
        email: req.body.email
      }).omit('password').value()
      const output = jwt.sign(searchData, 'K27qp2hjYQ5cFeVxRmMr0JvNA18ZV64kLGPDy8UhWpKKTdlBMZ76eMuVGSPpfrU')
      res.status(200).json({
        status: 200,
        on: 'userdata',
        data: output
      })
    }
  })

  router.get('/finduser/:email', (req, res) => {
    const data = db.get('data').find({
      email: req.params.email
    }).omit('password').value()
    res.status(200).json({
      data
    })
  })

  router.get('/', (req, res) => {
    res.status(200).json({
      message: 'Hello there!'
    })
  })
}).catch(err => console.log(err))




module.exports = router
