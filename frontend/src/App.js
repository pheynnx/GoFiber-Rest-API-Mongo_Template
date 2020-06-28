import React, { useEffect, useState } from 'react';
import logo from './logo.svg';
import './App.css';

import axios from 'axios';

const App = () => {
  const getDatabase = async () => {
    try {
      const context = await axios.get('/api/v1/users');
      console.log(context.data.data);
    } catch (error) {}
  };

  useEffect(() => {
    getDatabase();
  }, []);
  return (
    <div className="App">
      <header className="App-header">
        <img src={logo} className="App-logo" alt="logo" />
      </header>
    </div>
  );
};

export default App;
