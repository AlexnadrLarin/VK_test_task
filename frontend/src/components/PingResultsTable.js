import React from "react";
import { Table } from "react-bootstrap";

/**
 * Компонент для отображения таблицы результатов пинга.
 * 
 * @component
 * @param {Object} props - Свойства компонента.
 * @param {Array<Object>} props.data - Массив данных о результатах пинга.
 * @param {string} props.data[].ip - IP адрес сервера.
 * @param {number} props.data[].ping_time - Время пинга в миллисекундах.
 * @param {string} props.data[].date - Дата последней успешной попытки.
 * 
 * @returns {JSX.Element} Таблица с результатами пинга или сообщение об отсутствии данных.
 */
const PingResultsTable = ({ data }) => {
  return (
    <div className="container mt-4">
      <h2>Результаты пинга</h2>
      <Table striped bordered hover>
        <thead>
          <tr>
            <th>IP адрес</th>
            <th>Время пинга</th>
            <th>Дата последней успешной попытки</th>
          </tr>
        </thead>
        <tbody>
          {data && data.length > 0 ? (
            data.map((item, index) => (
              <tr key={index}>
                <td>{item.ip}</td>
                <td>{item.ping_time}</td>
                <td>{item.date}</td>
              </tr>
            ))
          ) : (
            <tr>
              <td colSpan="3" className="text-center">
                Нет данных для отображения
              </td>
            </tr>
          )}
        </tbody>
      </Table>
    </div>
  );
};

export default PingResultsTable;
