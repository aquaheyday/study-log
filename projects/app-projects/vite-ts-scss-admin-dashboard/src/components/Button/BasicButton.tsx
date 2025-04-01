import React from 'react';
import { Button, ButtonProps } from '@mui/material';

type BasicButtonProps = ButtonProps

const BasicButton: React.FC<BasicButtonProps> = ({ children, className, size = 'large', ...props }) => {
  return (
    <Button 
        color='primary' 
        variant="contained" 
        size={size}
        className={className} 
        {...props}
    >
        {children}
    </Button>
  );
};

export default BasicButton;
