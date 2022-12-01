import './App.css';

import { BrowserRouter } from 'react-router-dom'

import { AuthProvider } from "./contexts/authContext";

import Router from './components/Router';

function App() {
  return (
    <BrowserRouter>
      <AuthProvider>
        <Router />
      </AuthProvider>
    </BrowserRouter>
  );
}

export default App;
