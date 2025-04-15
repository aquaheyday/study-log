import React, {useCallback, useEffect, useState} from 'react';
import './PermissionApproval.scss';
import apiClient from '../../utils/apiUtils';
import {handlePermissionResponse} from '../../utils/permissionUtils';
import BasicTable, {Column} from '../../components/Table/BasicTable';
import CustomPagination from '../../components/Pagination/Pagination';
import BasicButton from '../../components/Button/BasicButton';
import {useAuth} from '../../context/AuthContext';

interface PermissionRequest {
  id: number;
  email: string;
  name: string;
  type: string;
  role: string;
  status: string; 
  requested_at: string;
}

interface PermissionResponse {
  data: any;
  requests: PermissionRequest[];
  current_page: number;
  total_pages: number;
  total_items: number; 
}

const PermissionApproval: React.FC = () => {
  const { role } = useAuth();
  
  const [requests, setRequests] = useState<PermissionRequest[]>([]);

  // 페이지네이션
  const [currentPage, setCurrentPage] = useState<number>(1);
  const [itemsPerPage, setItemsPerPage] = useState<number>(10);
  const [totalPages, setTotalPages] = useState<number>(1);
  const [totalItems, setTotalItems] = useState<number>(0);

  const [loading, setLoading] = useState<boolean>(false);
  const [error, setError] = useState<string | null>(null);

  // 데이터 fetch
  const fetchPermissionRequests = useCallback(
    async (page: number, limit: number) => {
      setLoading(true);
      setError(null);

      try {
        const res = await apiClient.get<PermissionResponse>(
            `/admin/permission-approvals-requests?page=${page}&limit=${limit}`
        );
        const responseData = res.data;

        setRequests(responseData.data.requests);
        setCurrentPage(responseData.data.current_page);
        setTotalPages(responseData.data.total_pages);
        setTotalItems(responseData.data.total_items);
      } catch (err) {
        console.error(err);
        setError('오류가 발생했습니다.');
      } finally {
        setLoading(false);
      }
    },
    []
  );

  // 승인/반려 요청 상태 업데이트 공통 함수
  const updatePermissionStatus = useCallback(
    async (requestId: number, status: 'approved' | 'rejected') => {
      try {
        // API PATCH 요청: 요청 ID와 상태를 전송
        const response = await apiClient.patch(`/admin/permission-approvals`, {
          request_id: String(requestId),
          status,
        });

        // 요청이 성공적으로 처리된 경우
        if (response.status === 200) {
          // 현재 요청 목록에서 해당 요청의 상태를 변경하여 UI를 업데이트
          setRequests((prev) =>
            prev.map((request) =>
              // ...request : 기존 request 객체 모든 속성을 새로운 객체로 복사
              request.id === requestId ? { ...request, status } : request
            )
          );
          return response.status; // 요청 성공 시 상태 코드 반환
        }
          // 요청이 실패한 경우 공통 응답 핸들러에서 처리
          return handlePermissionResponse(response, setError);
      } catch (error) {
        console.error(`처리 중 오류 발생: ${status === 'approved' ? '승인' : '반려'}`, error);
        throw error; // 에러를 호출한 곳에서 핸들링 가능하도록 추가
      }
    },
    [setRequests, setError] // 의존성 배열 (상태 업데이트 함수 사용 시 필수)
  );

  // 승인 요청
  const handleApprove = (requestId: number) => {
    if (window.confirm("정말 승인 처리하시겠습니까?")) {
      updatePermissionStatus(requestId, 'approved');
    }
  };

  // 반려 요청
  const handleReject = (requestId: number) => {
    if (window.confirm("정말 반려 처리하시겠습니까?")) {
      updatePermissionStatus(requestId, 'rejected');
    }
  };

  // 마운트/페이지 변경 시 데이터 로드
  useEffect(() => {
    fetchPermissionRequests(currentPage, itemsPerPage);
  }, [currentPage, itemsPerPage, fetchPermissionRequests]);

  // 테이블 컬럼 정의
  const columns: Column<PermissionRequest>[] = [
    { id: 'id', label: 'ID' },
    { id: 'email', label: '사용자' },
    { id: 'name', label: '이름' },
    { id: 'type', label: '타입' },
    { id: 'role', label: '권한' },
    { id: 'status', label: '상태' },
    {
      id: 'requested_at',
      label: '신청일',
      render: (req) => new Date(req.requested_at).toLocaleString(),
    },
    ...(role === 'write'
      ? [
          {
            id: 'actions',
            label: '관리',
            render: (req: { status: string; id: number; }) =>
              req.status === 'pending' ? (
                <div className="btn-box">
                  <BasicButton size="small" onClick={() => handleApprove(req.id)}>
                    승인
                  </BasicButton>
                  <BasicButton
                    size="small"
                    variant="outlined"
                    className="rejected"
                    onClick={() => handleReject(req.id)}
                  >
                    반려
                  </BasicButton>
                </div>
              ) : (
                <span className="completed">
                  {req.status === 'approved' && '승인 완료'}
                  {req.status === 'rejected' && '반려 완료'}
                </span>
              ),
          },
        ]
      : []),
  ];

  // 페이지네이션
  const handlePageChange = useCallback(
    (_event: React.ChangeEvent<unknown>, page: number) => {
      setCurrentPage(page);
    },
    []
  );

  const handleItemsPerPageChange = (e: React.ChangeEvent<HTMLSelectElement>) => {
    setItemsPerPage(Number(e.target.value));
    setCurrentPage(1); 
  };
  
  if (loading) return <div>로딩 중...</div>;
  if (error)
    return (
      <div>
        <p>{error}</p>
        <button onClick={() => fetchPermissionRequests(currentPage, itemsPerPage)}>재시도</button>
      </div>
    );

  return (
    <div className="login-page">
      <div className="inbound-container">
        <div className="scroll-container">
          <main>
            <div className="toolbar">
              <div className="toolbar-header">
                <div className="toolbar-left">
                  <h1>권한 관리</h1>
                </div>
              </div>

              {requests.length === 0 ? (
                <p>관리자 권한 신청 요청이 없습니다.</p>
              ) : (
                <div className="table-box">
                  <div className="table-grid">
                    <div className="table-container">
                    <div className="table-header">
                      <h3>권한 요청 목록</h3>
                    </div>
                      {/* <div className="table-box2">
                        <div className="table-content">
                          <table style={{ width: '100%', borderCollapse: 'collapse' }}>
                            <thead>
                              <tr style={{ borderBottom: '1px solid #ccc' }}>
                                <th>ID</th>
                                <th>사용자</th>
                                <th>이름</th>
                                <th>타입</th>
                                <th>권한</th>
                                <th>상태</th>
                                <th>신청일</th>
                                {role === 'write' && (
                                  <th>관리</th>
                                )}
                              </tr>
                            </thead>
                            <tbody>
                              {requests.map((item) => (
                                <tr key={item.id} style={{ borderBottom: '1px solid #ccc' }}>
                                  <td>{item.id}</td>
                                  <td>{item.email}</td>
                                  <td>{item.name}</td>
                                  <td>{item.type}</td>
                                  <td>{item.role}</td>
                                  <td>{item.status}</td>
                                  <td>{new Date(item.requested_at).toLocaleString()}</td>
                                  {role === 'write' && (
                                    <td className='btn-wrap'>
                                      {item.status === 'pending' ? (
                                        <div className='btn-box'>
                                          <BasicButton
                                            size="small"
                                            onClick={() => handleApprove(item.id)}
                                          >
                                            승인
                                          </BasicButton>
                                          <BasicButton
                                            size="small"
                                            variant="outlined" 
                                            className='rejected'
                                            onClick={() => handleReject(item.id)}
                                          >
                                            반려
                                          </BasicButton>
                                        </div>
                                      ) : (
                                        <span className='completed'>
                                          {item.status === 'approved' && '승인 완료'}
                                          {item.status === 'rejected' && '반려 완료'}
                                        </span>
                                      )}
                                    </td>
                                  )}
                                </tr>
                              ))}
                            </tbody>
                          </table>
                        </div>
                      </div> */}
                      <BasicTable<PermissionRequest>
                        data={requests}
                        columns={columns}
                        keyField="id"
                        checkbox={false}
                        noDataText="관리자 권한 신청 요청이 없습니다."
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
              )}
            </div>
          </main>
        </div>
      </div>
    </div>
  );
};

export default PermissionApproval;
