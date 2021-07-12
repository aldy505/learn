const {
  Router
} = require('express')
const csrf = require('csurf')
const argon = require('argon2')
const lowdb = require('lowdb')
const uuid = require('uuid')
const FileAsync = require('lowdb/adapters/FileAsync')
const router = Router()

const usersdb = new FileAsync(__dirname + '/database/users.json')
const csrfProtection = csrf({
  cookie: true
})
const APIKEY = `2spX^cJdJ8wArpEei!AQ4nH%Y&%5s*bdmNoZuMPkyrqLjO*lO8t3ZY^EvSX^Nx!r#9t!BtXw0S%j5ranpR*7mL1&`

router.get('/preform', csrfProtection, (req, res) => {
  res.status(200).json({
    csrfToken: req.csrfToken()
  })
})

lowdb(usersdb).then(db => {

  router.post('/login', csrfProtection, async (req, res) => {
    if (req.body.key === APIKEY) {
      const getData = db.get('data').find({
        email: req.body.email
      }).value()
      try {
        if (await argon.verify(getData.password, req.body.password)) {
          db.get('data').find({
            email: req.body.email
          }).assign({
            authorized: true
          }).assign({
            last_login: new Date()
          }).write()
          res.status(200).json({
            status: 200,
            on: 'login',
            message: 'success'
          })

        } else {
          res.status(200).json({
            status: 200,
            on: 'login',
            message: 'failed',
            detail: 'password verification failed'
          })
        }
      } catch (error) {
        res.status(200).json({
          status: 200,
          on: 'login',
          message: 'failed',
          detail: error
        })
      }
    }
  })

  router.post('/logout', csrfProtection, (req, res) => {
    if (req.body.key === APIKEY) {
      db.get('data').find({email: req.body.email}).assign({authorized: false}).write()
      .then(() => {
        res.status(200).json({
          status: 200,
          on: "logout",
          message: "success"
        })
      }).catch(error => {
        res.status(200).json({
          status: 200,
          on: "logout",
          message: "failed",
          detail: error
        })
      })

    }
  })

  router.post('/register', csrfProtection, async (req, res) => {
    if (req.body.key === APIKEY) {
      try {
        const hash = await argon.hash(req.body.password, {
          type: argon.argon2id,
          timeCost: 7,
          hashLength: 64
        })
        const id = uuid.v4()
        db.get('data').push({
          id: id,
          name: req.body.name,
          email: req.body.email,
          password: hash,
          authorized: false,
          last_login: new Date()
        }).write().then(() => {
          res.status(200).json({
            status: 200,
            on: "register",
            message: "success"
          })
        }).catch(error => {
          res.status(200).json({
            status: 200,
            on: "register",
            message: "failed",
            detail: error
          })
        })
      } catch (error) {
        res.status(200).json({
          status: 200,
          on: "register",
          message: "failed",
          detail: error
        })
        console.log(err)
      }
    }
  })
}).catch(err => console.log(err))


module.exports = router
