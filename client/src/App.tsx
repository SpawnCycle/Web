import './App.css'
import './index.css'
import { Navigate } from 'react-router-dom';
import React from 'react';
import { AuthContext } from './context/AuthContext';

function App() {
  const { isLoggedIn } = React.useContext(AuthContext);
  return (
    <>
      {isLoggedIn ? (
        <>
          <Navigate to="/app/releases" />
        </>
      ) : (
        <Navigate to="/app/login" />
      )}
    </>
  )
}

export default App
