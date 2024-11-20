var express = require('express');
var path = require('path');
const hbs = require('hbs');

var formRouter = require('./routes/form');

var app = express();

app.set('view engine', 'hbs');
app.set('views', 'views');



app.use(express.json());
app.use(express.static('public'));
app.use(express.urlencoded({ extended: false }));
app.use('/', formRouter);

const PORT = process.env.PORT || 3000;

app.listen(PORT, () => {
    console.log(`Server is working on port ${PORT}`);
});