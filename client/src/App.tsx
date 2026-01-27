import './App.css'
import './index.css'
import { Navigate } from 'react-router-dom';


function App() {
  const IsLoggedIn = true; // Replace with actual authentication logic
  return (
    <>
      {IsLoggedIn ? (
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
