import React from 'react';
import { useRoutes } from 'react-router-dom';
import AppRoute from './routes/AppRoute';
import { AuthProvider } from './context/AuthContext';
import "./App.css";

const App: React.FC = () => {
  const routing = useRoutes(AppRoute);

  return (
    <AuthProvider>
      {routing}
    </AuthProvider>
    );
};

export default App;
