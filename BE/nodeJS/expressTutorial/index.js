import express from 'express';
import usersRoutes from './routes/users.js';

//bodyParser, 미들웨어, 요청과 응답을 조작할 수 있게 해줌
//request의 body를 원하는 형태로 파싱해줌
import bodyParser from 'body-parser';

const app = express();
const PORT = 5000;

app.use(bodyParser.json());

//usersRoutes를 사용하겠다는 의미
app.use('/users', usersRoutes);

app.get('/', (req, res) => {
    res.send('Hello from Hompage!');
});

app.listen(PORT, () => console.log(`Listening on port ${PORT}`));