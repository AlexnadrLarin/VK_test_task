import 'bootstrap/dist/css/bootstrap.min.css';
import React from 'react';
import { createRoot } from 'react-dom/client';
import IpTable from './App';

const rootElement = document.getElementById('root');
const root = createRoot(rootElement);
root.render(<IpTable />);
