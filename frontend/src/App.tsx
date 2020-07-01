import React from 'react';
import './App.css';

import { useRecoilState, useRecoilValue } from 'recoil';
import { getDatabaseHandler, count } from './recoil';

import { User } from './types';

function App() {
  const [countState, setCount] = useRecoilState(count);

  const data: [] = useRecoilValue(getDatabaseHandler);

  const usersData = data.map((user: User, index) => {
    return (
      <div key={user.ID}>
        <p>{user.ID}</p>
        <p>{user.Username}</p>
        <p>{user.Password}</p>
      </div>
    );
  });

  return (
    <div className="App">
      <header className="App-header"></header>
      {countState}
      <button onClick={() => setCount((c) => c + 1)}>+ 1</button>
      <br></br>
      {usersData}
    </div>
  );
}

export default App;
