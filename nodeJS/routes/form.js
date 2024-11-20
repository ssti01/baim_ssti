// routes/form.js
var { Router } = require('express');
var router = Router();

router.get('/', (req, res, next) => {
    res.render('index'); // Рендерит файл index.hbs, используя layout из main.hbs
});

router.post('/form', (req, res, next) => {
    var userInput = req.body.data; // Получаем данные из формы
    res.render('index', { userInput }); // Рендерим index.hbs и передаем данные
});

module.exports = router;
