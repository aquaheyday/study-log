import React, { useState, useEffect } from 'react';
import "./BasicTable.scss";
import {
  TableContainer,
  Table,
  TableHead,
  TableBody,
  TableRow,
  TableCell,
  Checkbox,
  Paper,
} from '@mui/material';

export interface Column<T> {
  /** 각 컬럼의 ID */
  id: string;
  /** 테이블 헤더 텍스트트 */
  label: string;
  /** 렌더링 함수를 지정해주고 싶을 때 */
  render?: (row: T) => React.ReactNode;
  /** 컬럼 width */
  width?: string | number;
}

// BasicTable에서 받아야 할 props
interface BasicTableProps<T> {
  /** 표시할 데이터 배열 */
  data: T[];

  /** 컬럼 정의 배열 */
  columns: Column<T>[];

  /** 각 행을 구분할 key */
  keyField: keyof T;

  /** 체크박스가 필요한 경우 true */
  checkbox?: boolean;

  /** 선택/해제 될 때마다 호출되는 콜백 */
  onSelectionChange?: (selected: (string | number)[]) => void;

  /** 데이터가 없을 때 표시할 문구 */
  noDataText?: string;
}

function BasicTable<T>({
  data,
  columns,
  keyField,
  checkbox = false,
  noDataText = "No data",
  onSelectionChange,
}: BasicTableProps<T>) {

  // 선택된 key를 저장할 state
  const [selectedRows, setSelectedRows] = useState<string[]>([]);

  // 선택된 행의 데이터 
  useEffect(() => {
    if (checkbox) {
      const selectedData = data.filter((row) =>
        selectedRows.includes(String(row[keyField]))
      );
      console.log("선택된 행의 데이터:", selectedData);
    }
  }, [selectedRows, data, keyField, checkbox]);

  // 전체선택/해제 핸들러
  const handleSelectAll = (checked: boolean) => {
    if (checked) {
        const allRowKeys = data.map((row) => String(row[keyField]));
        setSelectedRows(allRowKeys);
        // 부모에 선택 상태 전달 
        if (onSelectionChange) {
            onSelectionChange(allRowKeys);
          }
      } else {
        setSelectedRows([]);
        if (onSelectionChange) {
            onSelectionChange([]);
          }
      }
  };

  // 각 행의 체크박스 클릭 핸들러
  const handleRowCheck = (rowKey: string, checked: boolean) => {
    let newSelectedRows: string[];
    if (checked) {
      newSelectedRows = [...selectedRows, rowKey];
    } else {
      newSelectedRows = selectedRows.filter((key) => key !== rowKey);
    }
    setSelectedRows(newSelectedRows);
    // 부모에 선택 상태 전달
    if (onSelectionChange) {
        onSelectionChange(newSelectedRows);
      }
  };

  // 모든 행이 선택되었는지 확인
  const allSelected = data.length > 0 && selectedRows.length === data.length;
  // 일부만 선택된 상태인지 확인
  const isIndeterminate =
    selectedRows.length > 0 && selectedRows.length < data.length; 
    
  return (
    <TableContainer component={Paper} elevation={0}>
      <Table>
        <TableHead>
          <TableRow>
            {/* 체크박스가 필요한 경우 */}
            {checkbox && (
              <TableCell style={{width: "20px"}}>
                <Checkbox 
                size="small" 
                // 전체 선택된 상태
                checked={allSelected}
                // 일부만 선택된 상태
                indeterminate={isIndeterminate}
                onChange={(e) => handleSelectAll(e.target.checked)}
                />
              </TableCell>
            )}
            {columns.map((col) => (
              <TableCell key={col.id} style={{ width: col.width }}>
                {col.label}
              </TableCell>
            ))}
          </TableRow>
        </TableHead>

        <TableBody>
          {/* 데이터가 존재할 경우 */}
          {data && data.length > 0 ? (
            data.map((row) => {
              const rowKey = String(row[keyField]);
              const checked = selectedRows.includes(rowKey);

              return (
                <TableRow key={rowKey} className="table-row">
                  {checkbox && (
                    <TableCell>
                      <Checkbox 
                      size="small" 
                      checked={checked}
                      onChange={(e) => handleRowCheck(rowKey, e.target.checked)}
                      />
                    </TableCell>
                  )}
                  {columns.map((col) => (
                    <TableCell key={col.id}>
                      {/* render 함수 사용, 없으면 [col.id] 값을 표시 */}
                      {col.render ? col.render(row) : (row as never)[col.id]}
                    </TableCell>
                  ))}
                </TableRow>
              );
            })
          ) : (
            <TableRow className="table-row">
              <TableCell
                colSpan={columns.length + (checkbox ? 1 : 0)}
                align="center"
              >
                {noDataText}
              </TableCell>
            </TableRow>
          )}
        </TableBody>
      </Table>
    </TableContainer>
  );
}

export default BasicTable;
