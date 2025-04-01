import React, { useState } from 'react';
import { useNavigate } from 'react-router-dom';
import './SignUp.scss';
import '../../styles/components/_error-message.scss';
import loginBg from '../../assets/images/login/login_bg.png';
import BasicButton from '../../components/Button/BasicButton';

const SignUp: React.FC = () => {
// 상태 관리
const [username, setUsername] = useState('');
const [email, setEmail] = useState('');
const [password, setPassword] = useState('');
const [confirmPassword, setConfirmPassword] = useState('');

const [usernameError, setUsernameError] = useState('');
const [emailError, setEmailError] = useState('');
const [passwordError, setPasswordError] = useState('');
const [confirmPasswordError, setConfirmPasswordError] = useState('');

const navigate = useNavigate();

// 유효성 검사 및 제출 처리
const handleLogin = async () => {
  let isValid = true;

  // 에러 메시지 초기화
  setUsernameError('');
  setEmailError('');
  setPasswordError('');
  setConfirmPasswordError('');

  // 유효성 검사
  if (username.trim() === '') {
    setUsernameError('Username is required.');
    isValid = false;
  }

  if (email.trim() === '') {
    setEmailError('Email is required.');
    isValid = false;
  }

  if (password.trim() === '') {
    setPasswordError('Password is required.');
    isValid = false;
  } else if (password.length < 6) {
    setPasswordError('Password must be at least 6 characters.');
    isValid = false;
  }

  if (confirmPassword.trim() === '') {
    setConfirmPasswordError('Confirm Password is required.');
    isValid = false;
  } else if (password.trim() != confirmPassword.trim()) {
    setConfirmPasswordError('Password is diff Confirm Password.')
    isValid = false;
  }

  // 유효성 검사를 통과하면 성공 메시지 표시
  if (isValid) {
    try {
      const response = await fetch('http://localhost:8080/v1/admin/user', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({
          username: username.trim(),
          email: email.trim(),
          password: password.trim(),
        }),
      });

      if (response.ok) {
        const data = await response.json();

        localStorage.setItem('token', data.token);
        alert('Join successful!');
        navigate('/permission-request')
      } else {
        const errorData = await response.json();
        alert(`Join failed: ${errorData.message || 'Unknown error'}`);
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

const handleEmailChange = (e: React.ChangeEvent<HTMLInputElement>) => {
  setEmail(e.target.value);
  setEmailError(''); // 에러 메시지 초기화
};

const handlePasswordChange = (e: React.ChangeEvent<HTMLInputElement>) => {
  setPassword(e.target.value);
  setPasswordError(''); // 에러 메시지 초기화
};

const handleConfirmPasswordChange = (e: React.ChangeEvent<HTMLInputElement>) => {
  setConfirmPassword(e.target.value);
  setConfirmPasswordError(''); // 에러 메시지 초기화
};

  return (
    <div className="login-page" style={{ backgroundImage: `url(${loginBg})` }}>
      <div className="login-bg">
        <div className="login-form-container">
          <form id="login_form" className="login-form" onSubmit={(e) => e.preventDefault()} >
            <div className="login-header">
              <h3>회원가입</h3>
              <div>
                <span>회원가입을 위한 정보를 입력하세요.</span>
              </div>
            </div>
            <div className="form-group">
              <label>회원명</label>
              <input type="text" id="username" name="username" value={username} onChange={handleUsernameChange} />
              {usernameError && <div id="usernameError" className="error-message">{usernameError}</div>}
            </div>
            <div className="form-group">
              <label>ID</label>
              <input type="text" id="email" name="email" value={email} onChange={handleEmailChange} />
              {emailError && <div id="emailError" className="error-message">{emailError}</div>}
            </div>
            <div className="form-group">
              <div className="password-header">
                <label>Password</label>
              </div>
              <input type="password" id="password" name="password" value={password} onChange={handlePasswordChange} />
              {passwordError && <div id="passwordError" className="error-message">{passwordError}</div>}
            </div>
            <div className="form-group">
              <div className="password-header">
                <label>Confirm Password</label>
              </div>
              <input type="password" id="comfirmPassword" name="comfirmPassword" value={confirmPassword} onChange={handleConfirmPasswordChange} />
              {confirmPasswordError && <div id="confirmPasswordError" className="error-message">{confirmPasswordError}</div>}
            </div>
            <BasicButton size="large" id="loginButton" onClick={handleLogin}>회원가입 요청</BasicButton>
          </form>
        </div>
      </div>
    </div>
  );
};

export default SignUp;
