import React, { useState, useEffect } from 'react';
import { useNavigate, NavLink } from 'react-router-dom';
import './Login.scss';
import '../../styles/components/_error-message.scss';
import loginBg from '../../assets/images/login/login_bg.png';
import BasicButton from '../../components/Button/BasicButton';
import BasicCheckbox from '../../components/Checkbox/BasicCheckbox';
{/*import loginGoogleLogo from '../../assets/images/login/login_google_logo.svg';
import loginKakaoLogo from '../../assets/images/login/login_kakao_logo.svg';*/}

const Login: React.FC = () => {
// 상태 관리
const [username, setUsername] = useState('');
const [password, setPassword] = useState('');
const [usernameError, setUsernameError] = useState('');
const [passwordError, setPasswordError] = useState('');
const [rememberMe, setRememberMe] = useState(false);
const navigate = useNavigate();

useEffect(() => {
  const rememberedEmail = localStorage.getItem('rememberedEmail');
  if (rememberedEmail) {
    setUsername(rememberedEmail);
    setRememberMe(true);
  }
}, []);

// 유효성 검사 및 제출 처리
const handleLogin = async () => {
  let isValid = true;

  // 에러 메시지 초기화
  setUsernameError('');
  setPasswordError('');

  // 유효성 검사
  if (username.trim() === '') {
    setUsernameError('Username is required.');
    isValid = false;
  }

  if (password.trim() === '') {
    setPasswordError('Password is required.');
    isValid = false;
  } else if (password.length < 6) {
    setPasswordError('Password must be at least 6 characters.');
    isValid = false;
  }

  // 유효성 검사를 통과하면 성공 메시지 표시
  if (isValid) {
    try {
      const response = await fetch('http://localhost:8080/v1/login', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({
          email: username.trim(),
          password: password.trim(),
        }),
      });

      if (response.ok) {
        const responseData = await response.json();
        console.log(responseData);

        if (rememberMe) {
          localStorage.setItem('rememberedEmail', username.trim()); // 이메일 저장
        } else {
            localStorage.removeItem('rememberedEmail'); // 이메일 제거
        }
        localStorage.setItem('token', responseData.data.token);
        localStorage.setItem('approved_count', responseData.data.approved_count);
        localStorage.setItem('pending_count', responseData.data.pending_count);
        localStorage.setItem('rejected_count', responseData.data.rejected_count);
        
        alert('Login successful!');
        navigate('/user-leads/franchise')
      } else {
        const errorData = await response.json();
        alert(`Login failed: ${errorData.message || 'Unknown error'}`);
      }
    } catch (error) {
      alert(`Error: ${error}`);
    }
  }
};

// 입력값 변경 처리
const handleUsernameChange = (e: React.ChangeEvent<HTMLInputElement>) => {
  setUsername(e.target.value);
  setUsernameError(''); // 에러 메시지 초기화
};

const handlePasswordChange = (e: React.ChangeEvent<HTMLInputElement>) => {
  setPassword(e.target.value);
  setPasswordError(''); // 에러 메시지 초기화
};

  return (
    <div className="login-page" style={{ backgroundImage: `url(${loginBg})` }}>
      <div className="login-bg">
        <div className="login-form-container">
          <form id="login_form" className="login-form" onSubmit={(e) => e.preventDefault()} >
            <div className="login-header">
              <h3>관리자</h3>
              <div>
                <span>회원가입이 필요하신가요?</span>
                <NavLink to="/sign-up">회원가입</NavLink>
              </div>
            </div>
            {/*<div className="social-login">
              <NavLink to="" className="social-button">
                <img src={loginGoogleLogo} alt="Google" />
                Use Google
              </NavLink>
              <NavLink to="" className="social-button">
                <img src={loginKakaoLogo} alt="Kakao" />
                Use Kakao
              </NavLink>
            </div>
            <div className="divider">
              <span></span>
              <span>OR</span>
              <span></span>
            </div>*/}
            <div className="form-group">
              <label>ID</label>
              <input type="text" id="username" name="username" value={username} onChange={handleUsernameChange} />
              {usernameError && <div id="usernameError" className="error-message">{usernameError}</div>}
            </div>
            <div className="form-group">
              <div className="password-header">
                <label>Password</label>
                {/*<NavLink to="">Forgot Password?</NavLink>*/}
              </div>
              <input type="password" id="password" name="password" value={password} onChange={handlePasswordChange} />
              {passwordError && <div id="passwordError" className="error-message">{passwordError}</div>}
            </div>
            <label className="remember-me">
            <BasicCheckbox 
              checked={rememberMe} 
              onChange={(e) => setRememberMe(e.target.checked)}
              label={<span className="text">로그인 정보 저장</span>}
            />
            </label>
            <BasicButton id="loginButton" onClick={handleLogin}>로그인</BasicButton>
          </form>
        </div>
      </div>
    </div>
  );
};

export default Login;
