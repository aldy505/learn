const cassandra = require('cassandra-driver')

const client = new cassandra.Client({
  contactPoints: ['localhost:9042', 'localhost:7000'],
  localDataCenter: 'localhost',
  keyspace: 'test'
})

const createTable = `CREATE TABLE test.test (
  id text,
  nilai int,
  PRIMARY KEY (id)
)`

client.execute(createTable).then(o => console.log(o)).catch(e => { throw e })

const insertData = `INSERT INTO test.test (id, nilai) VALUES (?, ?)`

client.execute(insertData, ['varul', { inggris: 60, indonesia: 80 }]).then(o => console.log(o)).catch(e => { throw e })