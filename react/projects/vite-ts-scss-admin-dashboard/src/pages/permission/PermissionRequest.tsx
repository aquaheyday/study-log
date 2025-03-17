import React, { useState, useEffect } from 'react';
import './PermissionRequest.scss';
import '../../styles/components/_error-message.scss';
import loginBg from '../../assets/images/login/login_bg.png';
import apiClient from "../../utils/apiUtils";
import PermissionCheckbox from '../../components/PermissionGroup/PermissionCheckbox';
import { useNavigate } from 'react-router-dom';
import BasicButton from '../../components/Button/BasicButton';

const PermissionRequest: React.FC = () => {
    type Inquiry = {
     id: number;
     type: string;
     name: string;
     sort: number;
     role: string;
     description: string;
   };

   type PermissionGroups = {
    id: number;
    name: string;
    permissions: [];
    sort: number;
   }

  const [selectedPermissions, setSelectedPermissions] = useState<number[]>([]);  
  const [permissions, setPermissions] = useState<Inquiry[]>([]);
  const [permissionGroups, setPermissionGroups] = useState<PermissionGroups[]>([]);
  const navigate = useNavigate();

  // API 호출
  const fetchPermissions = async () => {
    //setLoading(true);
    try {
      const response = await apiClient.get("/admin/permissions");

      const data = response.data.data;
      console.log(data);
      setPermissions(data);
    } catch (error) {
      console.error("Failed:", error);
    } finally {
      //setLoading(false);
    }
  };

  const fetchPermissionGroups = async () => {
    //setLoading(true);
    try {
      const response = await apiClient.get("/admin/permission-groups");

      const data = response.data.data;
      console.log(data);
      setPermissionGroups(data);
    } catch (error) {
      console.error("Failed:", error);
    } finally {
      //setLoading(false);
    }
  };

  const handleCheckboxChange = (id: number) => {
    setSelectedPermissions((prev) =>
      prev.includes(id) ? prev.filter((permId) => permId !== id) : [...prev, id]
    );
  };

  const handleSubmit = async() => {
    if (selectedPermissions.length === 0) {
      alert("Please select at least one permission."); // 에러 메시지 표시
      return;
    } else {
      try {
        const response = await apiClient.post("/admin/permissions", {
          permissions: selectedPermissions
        });
  
        const data = response.data;
        console.log(data);
        setPermissionGroups(data);
        navigate('/permission-request/end');
      } catch (error) {
        console.error("Failed:", error);
      } finally {
        //setLoading(false);
      }
    }

  };

  useEffect(() => {
    fetchPermissions();
    fetchPermissionGroups();
  }, []);

  return (
    <div className="login-page" style={{ backgroundImage: `url(${loginBg})` }}>
      <div className="login-bg">
        <div className="permission-form-container">
          <form id="login_form" className="login-form" onSubmit={(e) => e.preventDefault()} >
            <div className="login-header">
              <h3>관리자 권한 요청</h3>
              <div>
                <span>관리자에서 사용할 권한을 선택하세요.</span>
              </div>
            </div>
            <div className="checkbox-group">
              {permissions && permissions.length > 0 ? (
                // permissions.map((item, index) => (
                //   <div className="form-group">
                //     <input type="checkbox" value={index}></input>
                //     <label>{item.name}</label>
                //   </div>
                // ))
                // permissions.map((item) => (
                //   <div className="form-group">
                //     <PermissionCheckbox
                //       id={item.id} 
                //       name={item.name}
                //       isChecked={selectedPermissions.includes(item.id)} 
                //       onCheckboxChange={handleCheckboxChange}
                //     />
                // </div>
                // ))
                <>
                {/* 그룹 A */}
                <div className="form-A">
                  <h4>회원 구분</h4>
                  <div className="form-group">
                    {permissionGroups
                      .map((item) => (
                        <PermissionCheckbox
                          key={item.id} // key 추가
                          id={item.id}
                          name={item.name}
                          isChecked={selectedPermissions.includes(item.id)}
                          onCheckboxChange={handleCheckboxChange}
                        />
                      ))}
                  </div>
                </div>
          
                {/* 그룹 B */}
                <div className="form-B">
                  <h4>메뉴</h4>
                  <div className="form-group">
                    {permissions
                      .filter((item) => item.type === 'url') // 그룹 url에 해당하는 항목 필터링
                      .map((item) => (
                        <PermissionCheckbox
                          key={item.id} // key 추가
                          id={item.id}
                          name={item.name + '(데이터 삭제 ' + (item.role === 'write' ? '가능' : '불가') + ')'}
                          isChecked={selectedPermissions.includes(item.id)}
                          onCheckboxChange={handleCheckboxChange}
                        />
                      ))}
                  </div>
                </div>
          
                {/* 
                <div className="form-C">
                  <h4>관리 지점</h4>
                  <div className="form-group">
                    {permissions
                      .filter((item) => item.type === 'branch') // 그룹 branches에 해당하는 항목 필터링
                      .map((item) => (
                        <PermissionCheckbox
                          key={item.id} // key 추가
                          id={item.id}
                          name={item.name}
                          isChecked={selectedPermissions.includes(item.id)}
                          onCheckboxChange={handleCheckboxChange}
                        />
                      ))}
                  </div>
                </div>*/}
              </>
              ) 
              : (
                <p>No permissions available.</p>
              )
              }
            </div>
            <BasicButton id="loginButton" onClick={handleSubmit}>권한 요청</BasicButton>
          </form>
        </div>
      </div>
    </div>
  );
};

export default PermissionRequest;
