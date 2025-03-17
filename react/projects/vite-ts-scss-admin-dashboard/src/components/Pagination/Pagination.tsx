import React from 'react';
import { Pagination } from '@mui/material';
import './Pagination.scss';

interface CustomPaginationProps {
  currentPage: number;
  totalPages: number;
  totalItems: number;
  itemsPerPage: number;
  handlePageChange: (event: React.ChangeEvent<unknown>, page: number) => void;
  handleItemsPerPageChange: (e: React.ChangeEvent<HTMLSelectElement>) => void;
}

const CustomPagination: React.FC<CustomPaginationProps> = ({
  currentPage,
  totalPages,
  totalItems,
  itemsPerPage,
  handlePageChange,
  handleItemsPerPageChange,
}) => {
  return (
    <div className="pagination">
      <div className="pagination-top">
        <div className="pagination-page">
          <span>Show</span>
          <select value={itemsPerPage} onChange={handleItemsPerPageChange}>
            <option value="10">10</option>
            <option value="50">50</option>
            <option value="100">100</option>
          </select>
          <span>per page</span>
        </div>
        <div className="pagination-current">
          <span>
            {`${(currentPage - 1) * itemsPerPage + 1} 
              to ${Math.min(currentPage * itemsPerPage, totalItems)} 
              of ${totalItems}`}
          </span>
        </div>
      </div>

      <div className="pagination-bottom">
        <Pagination
          count={totalPages}          // 총 페이지 수
          page={currentPage}          // 현재 페이지
          onChange={handlePageChange} 
          shape="rounded"             // 모서리 둥근 스타일
          showFirstButton             // "<<" 버튼 표시
          showLastButton              // ">>" 버튼 표시
          siblingCount={1}            // 현재 페이지 양옆에 표시할 페이지 수
          boundaryCount={1}           // 시작/끝에 표시할 페이지 수
        />
      </div>
    </div>
  );
};

export default CustomPagination;
