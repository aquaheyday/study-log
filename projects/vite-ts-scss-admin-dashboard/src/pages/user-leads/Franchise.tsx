import React, { useState, useEffect, useCallback }from "react";
import "./Franchise.scss";
import apiClient from "../../utils/apiUtils";
import { useAuth } from '../../context/AuthContext';
// import { NavLink } from "react-router-dom";
import BasicTable, { Column } from '../../components/Table/BasicTable';
import BasicButton from "../../components/Button/BasicButton";
import CustomPagination from '../../components/Pagination/Pagination'; 
import { RiDeleteBinLine } from "react-icons/ri";

const Franchise: React.FC = () => {
  const { role } = useAuth();

  // 초기값
  type Inquiry = {
    id: number;
    name: string;
    birth_date: string;
    gender: string;
    phone_number: string;
    branch_location: string;
    inquiry: string;
    create_at: string;
  };

  const [inquiries, setInquiries] = useState<Inquiry[]>([]); // 문의 목록
  const [selectedRows, setSelectedRows] = useState<(number | string)[]>([]);   // 선택된 행 목록
  const [currentPage, setCurrentPage] = useState(1); // 현재 페이지
  const [itemsPerPage, setItemsPerPage] = useState(10); // 페이지당 항목 수
  const [totalItems, setTotalItems] = useState(0); // 전체 항목 수
  const [totalPages, setTotalPages] = useState(0); // 전체 페이지 수
  const [loading, setLoading] = useState(true); // 로딩 상태

  // API 호출
  const fetchInquiries = async (page: number, limit: number) => {
    setLoading(true);
    try {
      const response = await apiClient.get("/franchise-user-leads", {
        params: { page, limit },
      });

      const responseData = response.data;

      setInquiries(responseData.data.inquiries);
      setCurrentPage(responseData.data.current_page);
      setTotalItems(responseData.data.total_items);
      setTotalPages(responseData.data.total_pages);
    } catch (error) {
      console.error("Failed:", error);
    } finally {
      setLoading(false);
    }
  };

  // 컴포넌트가 처음 로드될 때 또는 currentPage, itemsPerPage 변경 시 API 호출
  useEffect(() => {
    fetchInquiries(currentPage, itemsPerPage);
  }, [currentPage, itemsPerPage]);


  // 테이블 컬럼 정의
  const columns: Column<Inquiry>[] = [
    {
      id: 'name',
      label: '이름',
    },
    {
      id: 'gender',
      label: '성별',
    },
    {
      id: 'birth_date',
      label: '생년월일',
    },
    {
      id: 'branch_location',
      label: '희망 지역',
    },
    {
      id: 'phone_number',
      label: '연락처',
    },
    {
      id: 'inquiry',
      label: '문의 내용',
    },
    ...(role === 'write'
      ? [
          {
            id: 'actions',
            label: '관리',
            width: '50px',
            render: (row: Inquiry) => (
              <BasicButton 
                size="small" 
                variant="outlined"
                className="rejected"
                onClick={() => deleteInquiry(row.id)}
              >
                삭제
              </BasicButton>
            ),
          },
        ]
      : []),
  ];

  // 페이지 변경
  const handlePageChange = useCallback(
    (_event: React.ChangeEvent<unknown>, page: number) => {
      setCurrentPage(page);
    },
    []
  );

  // 페이지당 항목 수 변경
  const handleItemsPerPageChange = (event: React.ChangeEvent<HTMLSelectElement>) => {
    const newLimit = parseInt(event.target.value, 10);
    setItemsPerPage(newLimit);
    setCurrentPage(1);
  };

   // 삭제 API 호출
   const deleteInquiry = async (id: number) => {
    const confirmDelete = window.confirm("정말로 삭제하시겠습니까?");

    if (confirmDelete) {
      try {
        // 삭제 요청 API 호출
        await apiClient.delete(`/franchise-user-leads`, { 
          data: { ids: [Number(id)] } 
        });
        // 삭제 후 데이터 갱신
        fetchInquiries(currentPage, itemsPerPage);
      } catch (error) {
        console.error("Failed to delete inquiry:", error);
      }
    }
  };

   // 선택된 행들 삭제 API 호출
   const deleteSelectedInquiries = async () => {
    //  console.log(selectedRows)
    if (selectedRows.length === 0) {
      alert("선택된 항목이 없습니다.");
      return;
    }
    const confirmDelete = window.confirm("정말로 선택된 항목들을 삭제하시겠습니까?");
    if (!confirmDelete) return;

    try {
      // 선택된 행들 삭제 요청 API 호출
      await apiClient.delete('/franchise-user-leads', {
        data: { ids: selectedRows.map(Number) }
      });

      // 선택된 행들 삭제 후 데이터 갱신
      fetchInquiries(currentPage, itemsPerPage);
      // 선택된 행들 상태 초기화
      setSelectedRows([]);
    } catch (error) {
      console.error("Failed to delete inquiries:", error);
    }
  };

  if (loading) return <p>Loading...</p>;
  
  return (
    <div className="inbound-container">
      <div className="scroll-container">
        <main>
          <div className="toolbar">
            <div className="toolbar-header">
              <div className="toolbar-left">
                <h1>가맹모집 관리</h1>
              </div>
              {/* 
              <div className="toolbar-right">
                <button className="search-button">검색</button>
                <NavLink to="" className="export-button">
                  Export
                </NavLink>
              </div>
              */}
            </div>

           <div className="table-box">
              <div className="table-grid">
                <div className="table-container">
                  <div className="table-header">
                    <div>
                      <h3>가맹모집 문의 목록</h3>
                    </div>
                    <div>
                      <BasicButton   
                        size="small" 
                        className="btn-all"
                        onClick={deleteSelectedInquiries}
                      >
                        <span className="menu-icon">
                          <RiDeleteBinLine />
                        </span>
                        선택 삭제
                      </BasicButton>
                    </div>                    
                  </div>
                   {/* <div className="table-box2">
                    <div className="table-content">
                      <table>
                        <thead>
                          <tr>
                            <th style={{ width: "20px" }}>
                              <input type="checkbox" />
                            </th>
                            <th>이름</th>
                            <th>성별</th>
                            <th>생년월일</th>
                            <th>희망 지역</th>
                            <th>연락처</th>
                            <th>문의내용</th>
                            {role === 'write' && (
                              <th style={{ width: "50px" }}>관리</th>
                            )}
                          </tr>
                        </thead>
                        <tbody>
                        {inquiries && inquiries.length > 0 ? (
                          inquiries.map((inquiry, index) => (
                            <tr key={index}>
                              <td><input type="checkbox"></input></td>
                              <td>{inquiry.name}</td>
                              <td>{inquiry.gender}</td>
                              <td>{inquiry.birth_date}</td>
                              <td>{inquiry.branch_location}</td>
                              <td>{inquiry.phone_number}</td>
                              <td>{inquiry.inquiry}</td>
                              {role === 'write' && (
                                <td>
                                  <BasicButton size="small" onClick={() => deleteInquiry(inquiry.id)}>삭제</BasicButton>
                                </td>
                              )}
                            </tr>
                          ))
                        ) : (
                          <tr>
                            <td colSpan={8} style={{ textAlign: "center" }}>문의 내역이 없습니다.</td>
                          </tr>
                        )}
                        </tbody>
                      </table>
                    </div>              
                  </div> */}
                  <BasicTable<Inquiry>
                    data={inquiries}
                    columns={columns}
                    keyField="id"
                    checkbox={true}
                    noDataText="문의 내역이 없습니다."
                    onSelectionChange={(rows) => setSelectedRows(rows)}
                  />
                  <CustomPagination
                    currentPage={currentPage}
                    totalPages={totalPages}
                    totalItems={totalItems}
                    itemsPerPage={itemsPerPage}
                    handlePageChange={handlePageChange}
                    handleItemsPerPageChange={handleItemsPerPageChange}
                  />
                </div>
              </div>
            </div>
          </div>
        </main>
      </div>
    </div>
  );
};

export default Franchise;
