import React, { useState, useEffect, useRef } from "react";
import "./Navigation.scss";
import { Outlet, NavLink } from "react-router-dom";
import { AiOutlineDatabase, AiOutlineSetting, AiOutlineLogout } from 'react-icons/ai';
import apiClient from "../../utils/apiUtils";
import { useAuth } from "../../context/AuthContext";

type MenuItem = {
  id: number;
  name: string;
  parent_id: number;
  url: string;
  level: number;
  sort: number;
  role: string;
  children?: MenuItem[];
};

const Navigation: React.FC = () => {
  const [menu1Depth, setMenu1Depth] = useState<MenuItem[]>([]);
  const [activeParentId, setActiveParentId] = useState<number | null>(null);
  const firstLinkRef = useRef<HTMLAnchorElement | null>(null);
  const { setRole, logout } = useAuth(); 

  const getIconComponent = (menu: MenuItem) => {
    if (menu.name === '관리자 메뉴') return <AiOutlineDatabase />;
    if (menu.name === '시스템 메뉴') return <AiOutlineSetting />;
    if (menu.name === '로그아웃') return <AiOutlineLogout />;
    return <AiOutlineDatabase />; 
  };
  
  const fetchMenu = async () => {
    try {
      const response = await apiClient.get("/menu", {
        params: { group: "admin" },
      });
      
      const data: MenuItem[] = response.data.data;
      
      // 1depth & 2depth 생성
      const menuHierarchy = data.reduce((acc, item) => {
        if (item.level === 1) {
          acc.push({ ...item, children: [] });
        } else {
          const parent = acc.find((parentItem) => parentItem.id === Number(item.parent_id));
          if (parent) {
            parent.children.push(item);
          }
        }
        return acc;
      }, [] as (MenuItem & { children: MenuItem[] })[]);
  
      //console.log("Menu Hierarchy:", JSON.stringify(menuHierarchy, null, 1)); // 디버깅용 출력
      setMenu1Depth(menuHierarchy);
    } catch (error) {
      console.error("Failed to fetch menu data:", error);
    }
  };

  useEffect(() => {
    fetchMenu();
  }, []);
  
  // 1 Depth 자동 활성화
  useEffect(() => {
    if (menu1Depth.length > 0 && activeParentId === null) {
      setActiveParentId(menu1Depth[0].id);
    }
  }, [activeParentId, menu1Depth]);

  // 2 Depth의 첫 번째 메뉴 자동 클릭
  useEffect(() => {
    if (activeParentId !== null && firstLinkRef.current) {
      firstLinkRef.current.click();
    }
  }, [activeParentId]);

  const isActive1Depth = (url: string) => location.pathname.includes(url);
  const isActive2Depth = (url: string) => location.pathname === url;

  useEffect(() => {
    menu1Depth.forEach((parent) => {
      parent.children?.forEach((child) => {
        if (isActive2Depth(child.url)) {
          setRole(child.role || "read");
        }
      });
    });
  }, [menu1Depth, setRole]);

  const handleLogout = () => {
    logout();
  };

  return (
  <div className="body">
    <div className="main">
      <div className="navigation">
        <div className="sidebar">
          {/* Logo Start */}
          <div className="logo">
            {/* 임시 link로 이후 main으로 변경 예정 */}
            <NavLink to="/" className="logo-link">
              <img alt="" src="/favicon.ico" />
            </NavLink>
          </div>
          {/* Logo End */}

          {/*  1 Depth Menu Start */}
          <div className="top-menu">
            <div className="menu-items">
              <div className="menu-top">
              {menu1Depth.map((item, index) => (
                <div
                key={item.id}
                className={`menu-item ${
                  activeParentId === item.id ||
                  (activeParentId === null && index === 0) ||
                  isActive1Depth(item.url)
                    ? "active"
                    : ""
                }`}
                onClick={() => setActiveParentId(item.id)}
                >
                  <span className="menu-icon">
                    {getIconComponent(item)}
                  </span>
                  <span className="tooltip">{item.name}</span>
                </div>
              ))}
              </div>
              <div
                className="menu-bottom menu-item logout"
                onClick={handleLogout}
              >
                <span className="menu-icon">
                  <AiOutlineLogout />
                </span>
                <span className="tooltip">로그아웃</span>
              </div>
            </div>
          </div>
          {/*  1 Depth Menu End */}
        </div>

        {/* 2 Depth Menu Start */}
        <div className="depth-menu">
          <div className="menu-container">
            <div className="menu-inner">
                {menu1Depth
                .filter((parent) => activeParentId ? parent.id === activeParentId : isActive1Depth(parent.url))
                .map((parent) => (
                <div key={parent.id} className="menu-section">
                  <div className="menu-title">
                    <div className="title-text">{parent.name}</div>
                  </div>
                  <div className="menu-links">
                    {parent.children &&
                      parent.children.map((child, index) => {
                        return (
                          <NavLink to={child.url || "/"}
                            key={child.id}
                            ref={index === 0 ? firstLinkRef : null}
                            className={({ isActive }) =>
                              `menu-link ${isActive ? "active" : ""}`
                            }
                          >
                            {child.name}
                          </NavLink>
                        );
                      })}
                  </div>
                </div>
              ))}
            </div>
          </div>
        </div>
        {/* 2 Depth Menu End */}
      </div>

      <main className="content">
        <Outlet />
      </main>
    </div>
  </div>
);
};

export default Navigation;
