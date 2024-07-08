import './App.css';
import apiImages from './APICall';
import { useState } from 'react';

const App = () => {
  const content = apiImages()
  return (
    <>
      {content}
    </>
  );
};

export default App;
