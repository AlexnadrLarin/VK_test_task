const REACT_APP_BACKEND_API_URL = process.env.REACT_APP_BACKEND_API_URL;

export class API {
  /**
   * Получает результаты пинга.
   *
   * @async
   * @function
   * @returns {Promise<Array>} Массив объектов с результатами пинга.
   * @throws {Error} Если произошла ошибка при запросе или обработке данных.
   */
  async fetchPingResults() {
    try {
      if (!REACT_APP_BACKEND_API_URL) {
        throw new Error('URL для backend API не задан. Убедитесь, что переменная окружения REACT_APP_BACKEND_API_URL установлена.');
      }

      const url = REACT_APP_BACKEND_API_URL;
      const response = await fetch(url);

      if (!response.ok) {
        throw new Error(`Ошибка запроса: ${response.status} - ${response.statusText}`);
      }

      const data = await response.json();

      if (!data || !Array.isArray(data.ping_results)) {
        throw new Error('Не удалось извлечь данные из ответа. Ожидался массив ping_results.');
      }

      return this.validatePingResults(data.ping_results);
    } catch (error) {
      throw error;
    }
  }

  /**
   * Валидирует массив результатов пинга.
   *
   * @param {Array} pingResults - Массив результатов пинга.
   * @returns {Array} Отфильтрованные и валидные данные.
   * @throws {Error} Если данные имеют неверный формат.
   */
  validatePingResults(pingResults) {
    const ipRegex = /^(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)$/;

    // Поддержка различных обозначений времени
    const pingTimeRegex = /^\d+(\.\d+)?(ns|µs|ms|s|m|h)$/;
    const timestampRegex = /^\d{2}:\d{2}:\d{4}:\d{2}:\d{2}:\d{2}\.\d{3}$/;

    const invalidResults = pingResults.filter((result) => {
      const { ip, ping_time, date } = result;

      return (
        !ipRegex.test(ip) ||
        !pingTimeRegex.test(ping_time) ||
        !timestampRegex.test(date)
      );
    });

    if (invalidResults.length > 0) {
      throw new Error(
        'Некорректные данные'
      );
    }

    return pingResults;
  }
}
