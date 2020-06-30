import React, { useEffect } from 'react';
import './App.css';

import { useRecoilState, useRecoilValue } from "recoil";
import { databaseState, count } from './recoil';


const App = () => {
  const [countState, setCount] = useRecoilState(count);

  const data = useRecoilValue(databaseState);

  return (
    <div className="App">
      <header className="App-header"></header>
      {countState}
      <button onClick={() => setCount(c => c + 1)}>+ 1</button>
      <br></br>
      {data.Username}
    </div>
  );
};

export default App;
