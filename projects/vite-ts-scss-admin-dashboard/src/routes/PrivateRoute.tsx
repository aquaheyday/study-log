// src/routes/PrivateRoute.tsx
import React from 'react';
import { Navigate, Outlet } from 'react-router-dom';
import { isAuthenticated, isPermmision } from '../utils/authUtils';

const PrivateRoute: React.FC = () => {
  if (!isAuthenticated()) {
    // 토큰이 없으면 로그인 페이지로 리다이렉트
    return <Navigate to="/" replace />;
  }

  if (isPermmision() == 'pending') {
    // 권한 부여 대기중
    return <Navigate to="/permission-request/end" replace />;
  }

  if (isPermmision() == 'request') {
    // 부여된 권한이 없고, 권한 부여 대기중이 아닐때
    return <Navigate to="/permission-request" replace />;
  }

  // 토큰이 있으면 내부 라우트를 렌더링
  return <Outlet />;
};

export default PrivateRoute;
