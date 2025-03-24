import { RouteObject } from 'react-router-dom';
import Login from '../pages/login/Login';
import SignUp from '../pages/login/SignUp';
import PermissionRequest from '../pages/permission/PermissionRequest';
import Franchise from '../pages/user-leads/Franchise';
import NotFound from '../pages/status/NotFound';
import Navigation from '../components/Navigation/Navigation';
import PrivateRoute from './PrivateRoute';
import ExpiredToken from '../pages/status/ExpiredToken';
import PermissionRequestEnd from '../pages/permission/PermissionRequestEnd';
import PermissionApproval from '../pages/permission/PermissionApproval';

const AppRoute: RouteObject[] = [
  {
    path: '/',
    element: <Login />,
  },
  {
    path: '/sign-up',
    element: <SignUp />,
  },
  {
    path: '/permission-request',
    element: <PermissionRequest />,
  },
  {
    path: '/permission-request/end',
    element: <PermissionRequestEnd />,
  },
  {
    element: <PrivateRoute />,
    children: [
        {
            element: <Navigation />,
            children: [
                {
                path: '/settings/permission-approval',
                element: <PermissionApproval />,
                },
            ],
        }
    ],
    
  },
  {
    element: <PrivateRoute />,
    children: [
        {
            element: <Navigation />,
            children: [
                {
                path: '/user-leads/franchise',
                element: <Franchise />,
                },
            ],
        }
    ],
    
  },
  {
    path: '/expired-token',
    element: <ExpiredToken />,
  },
  {
    path: '/*',
    element: <NotFound />,
  },
];

export default AppRoute;
