const express = require('express')
const app = express()
const port = 8000

app.get('/', (req, res) => {
  res.send('Hello from the Node service!')
})

app.listen(port);