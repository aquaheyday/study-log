// 📂 components/ExampleComponent.jsx
// 작성일: YYYY-MM-DD
// 설명: [컴포넌트 간단 설명 - 어떤 역할을 하는 컴포넌트인지]

import React from 'react';
import PropTypes from 'prop-types';

/**
 * ExampleComponent - 컴포넌트 설명
 * @param {Object} props - 컴포넌트에 전달되는 props
 * @returns JSX Element
 */
function ExampleComponent({ title, onClick }) {
    return (
        <div className="example-component">
            <h2>{title}</h2>
            <button onClick={onClick}>Click Me</button>
        </div>
    );
}

// ✅ Props 기본값 설정 (optional)
ExampleComponent.defaultProps = {
    title: '기본 제목',
};

// ✅ Props 타입 지정 (optional)
ExampleComponent.propTypes = {
    title: PropTypes.string,
    onClick: PropTypes.func.isRequired,
};

// ✅ 컴포넌트 이름 익스포트
export default ExampleComponent;
