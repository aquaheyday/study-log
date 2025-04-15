import React from "react";
import Checkbox from "@mui/material/Checkbox";
import FormControlLabel from "@mui/material/FormControlLabel";
import "./PermissionCheckbox.scss";

type PermissionCheckboxProps = {
  id: number;
  name: string;
  isChecked: boolean;
  onCheckboxChange: (id: number) => void;
};

const PermissionCheckbox: React.FC<PermissionCheckboxProps> = ({
  id,
  name,
  isChecked,
  onCheckboxChange,
}) => {
  return (
    <div className="checkbox-item">
      <FormControlLabel
        control={
          <Checkbox
            checked={isChecked} // 체크 상태 22333444555
            onChange={() => onCheckboxChange(id)} // 변경 이벤트
            id={`permission-${id}`}
            value={id}
            color="primary" // MUI 색상 옵션
            sx={{
              "& .MuiSvgIcon-root": {
                fontSize: 26, // 체크박스의 크기를 조정 
              },
            }}
          />
        }
        label={name}
      />
    </div>
  );
};

export default PermissionCheckbox;
