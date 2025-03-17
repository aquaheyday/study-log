import React from "react";
import { useNavigate } from "react-router-dom";
import BasicButton from "../../components/Button/BasicButton";
import { BsCheckCircleFill } from "react-icons/bs";

const ExpiredToken: React.FC = () => {
  const navigate = useNavigate();
  const logout = () => {
    localStorage.removeItem('token');
    localStorage.removeItem('approved_count');
    localStorage.removeItem('pending_count');
    localStorage.removeItem('rejected_count');
    navigate("/");
};
  return (
    <div style={{ 
      display: "flex",
      flexDirection: "column",
      alignItems: "center",
      margin: "250px auto"
      }}>
      <BsCheckCircleFill size="70" color="#1976d2"/>
      <h2 style={{marginTop:"30px"}}>회원가입 및 권한 신청이 완료되었습니다.</h2>
      <p style={{margin:"0",color:"#4b5675"}}>굿바이브 관리자가 승인 완료되면 관리자 이용이 가능합니다.</p>
      <BasicButton onClick={() => logout()} style={{marginTop:"70px"}}>로그인 페이지로 이동</BasicButton>
    </div>
  );
};

export default ExpiredToken;
