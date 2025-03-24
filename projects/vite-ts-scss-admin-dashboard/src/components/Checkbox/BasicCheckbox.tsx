import React from "react";
import Checkbox from '@mui/material/Checkbox';
import FormControlLabel from "@mui/material/FormControlLabel";

type BasicCheckboxProps = {
  checked: boolean;
  onChange: (event: React.ChangeEvent<HTMLInputElement>) => void;
  label?: React.ReactNode; 
};

const BasicCheckbox: React.FC<BasicCheckboxProps> = ({ checked, onChange, label }) => {
  return (
    <FormControlLabel
      control={<Checkbox checked={checked} onChange={onChange} />}
      label={label}
      color="default"
      sx={{
        "& .MuiSvgIcon-root": {
          fontSize: 22, // 체크박스의 크기를 조정 
        },
      }}
    />
  );
};

export default BasicCheckbox;
