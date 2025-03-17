export const isAuthenticated = (): boolean => {
    const token = localStorage.getItem('token');
    return !!token; // 토큰이 있으면 true, 없으면 false
};

export const isPermmision = (): string => {
    const approvedCount = localStorage.getItem('approved_count');
    const pendingCount = localStorage.getItem('pending_count');

    if (approvedCount !== null && parseInt(approvedCount, 10) > 0) {
        return 'approved';
    } else if (pendingCount !== null && parseInt(pendingCount, 10) > 0) {
        return 'pending';
    } else {
        return 'request';
    }
    
};