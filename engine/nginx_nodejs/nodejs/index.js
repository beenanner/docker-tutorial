'use strict';

const express = require('express');

// Constants
const PORT = 8080;

// App
const app = express();
app.get('/', function (req, res) {
  console.log("Extremely bad/lazy way to log in nodejs via docker we should be using middleware here.");
  res.send('Hello world\n');
});

app.listen(PORT);
console.log('Running on http://nodejs:' + PORT);