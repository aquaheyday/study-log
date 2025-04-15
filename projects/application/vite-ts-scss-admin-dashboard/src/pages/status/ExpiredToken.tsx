import React from "react";
import { useNavigate } from "react-router-dom";
import BasicButton from "../../components/Button/BasicButton";

const ExpiredToken: React.FC = () => {
  const navigate = useNavigate();

  return (
    <div style={{ textAlign: "center", padding: "50px" }}>
      <h1>세션이 만료되었습니다</h1>
      <p>토큰이 만료되었습니다. 다시 로그인해주세요.</p>
      <BasicButton onClick={() => navigate("/")}>로그인 페이지로 이동</BasicButton>
    </div>
  );
};

export default ExpiredToken;
