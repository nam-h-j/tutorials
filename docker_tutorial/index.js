const express = require('express');

const app = express();

app.get('/', (req, res) => {
    res.send('App for Using Docker');
});

app.listen(8080, () => console.log('Running...'));