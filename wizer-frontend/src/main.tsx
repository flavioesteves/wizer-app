import React from 'react'
import ReactDOM from 'react-dom/client'
import App from './App.tsx'
import './index.css'

import { library } from '@fortawesome/fontawesome-svg-core';
import { fas, faRss } from '@fortawesome/free-solid-svg-icons';

library.add(fas, faRss);


ReactDOM.createRoot(document.getElementById('root')!).render(
  <React.StrictMode>
    <App />
  </React.StrictMode>,
)
