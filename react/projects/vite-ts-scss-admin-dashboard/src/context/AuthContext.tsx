import React, { createContext, useContext, useState, ReactNode } from 'react';

type AuthContextType = {
  role: string; // 'read' | 'write' 등의 값
  setRole: (role: string) => void;
  logout: () => void;
};

const AuthContext = createContext<AuthContextType | undefined>(undefined);

export const AuthProvider: React.FC<{ children: ReactNode }> = ({ children }) => {
  const [role, setRole] = useState<string>(''); // 초기값을 빈 문자열로 설정

  const logout = () => {
    setRole('');
    // localStorage.removeItem('accessToken');
    window.location.href = "/";
  }

  return (
    <AuthContext.Provider value={{ role, setRole, logout }}>
      {children}
    </AuthContext.Provider>
  );
};

export const useAuth = (): AuthContextType => {
  const context = useContext(AuthContext);
  if (!context) {
    throw new Error('useAuth must be used within an AuthProvider');
  }
  return context;
};

