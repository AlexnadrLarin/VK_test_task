import express from 'express';
import path from 'path';
import helmet from 'helmet';

const app = express();

const PORT = process.env.PORT || 8080;

if (process.env.NODE_ENV === 'production') {
    console.log = function() {};
    console.warn = function() {};
    console.error = function() {};
}

app.use(helmet());

const __dirname = path.resolve();
app.use('/', express.static(path.resolve(__dirname, '../public')));

app.get('*', (req, res) => {
  res.sendFile(path.join(__dirname, '../public', 'index.html'));
});

app.listen(PORT, () => {
  console.log(`Сервер запущен на порту ${PORT}`);
}).on('error', (err) => {
  console.error('Ошибка при запуске сервера:', err);
});
