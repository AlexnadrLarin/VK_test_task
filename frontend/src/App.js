import React, { useEffect, useState } from 'react';
import { API } from './modules/Api';
import PingResultsTable from './components/PingResultsTable';
import ErrorMessage from './components/ErrorMessage';

const api = new API();

/**
 * Главный компонент приложения для отображения таблицы результатов пинга.
 * @component
 */
function App() {
  /** 
   * Состояние для хранения данных пинга.
   * @type {[Array, Function]} 
   */
  const [pingData, setPingData] = useState([]);

  /**
   * Состояние для хранения ошибок при получении данных.
   * @type {[string|null, Function]}
   */
  const [error, setError] = useState(null);

  useEffect(() => {
    let isMounted = true;

    /**
     * Асинхронная функция для получения данных о пингах с сервера.
     * Обновляет состояние при успешном запросе или устанавливает ошибку.
     */
    const fetchPingData = async () => {
      try {
        const data = await api.fetchPingResults();
        if (isMounted) {
          if (Array.isArray(data)) {
            setPingData(data);
            setError(null); 
          } else {
            setError('Получены некорректные данные. Попробуйте позже.');
          }
        }
      } catch (err) {
        if (isMounted) {
          if (err instanceof Error) {
            setError(`Ошибка данных: ${err.message}`);
          } else if (err instanceof TypeError) {
            setError('Проблемы с подключением к серверу. Пожалуйста, проверьте ваше интернет-соединение.');
          } else {
            setError('Не удалось получить данные пинга. Попробуйте позже.');
          }
        }
      }
    };

    fetchPingData();

    /**
     * Интервал получения данных в секундах из переменной окружения.
     * Значение по умолчанию — 10 секунд.
     * @type {number}
     */
    const intervalSeconds = parseInt(process.env.REACT_APP_TIME_INTERVAL || '10', 10);

    if (!process.env.REACT_APP_TIME_INTERVAL) {
      console.warn('REACT_APP_TIME_INTERVAL не установлено. Используется значение по умолчанию (10 секунд).');
    }

    const intervalMilliseconds = intervalSeconds * 1000;

    /**
     * Идентификатор интервала для периодического получения данных.
     * @type {number}
     */
    const intervalId = setInterval(fetchPingData, intervalMilliseconds);

    return () => {
      isMounted = false;
      clearInterval(intervalId);
    };
  }, []);

  return (
    <div className="container mt-4">
      {error ? (
        <ErrorMessage message={error} />
      ) : (
        <PingResultsTable data={pingData} />
      )}
    </div>
  );
}

export default App;
