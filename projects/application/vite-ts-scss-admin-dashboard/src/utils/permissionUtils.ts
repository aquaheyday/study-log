export const handlePermissionResponse = (
    response: any,
    setError: React.Dispatch<React.SetStateAction<string | null>>
) => {
    if (response.status === 400) {
        console.error("잘못된 요청 형식:", response.data?.message);
        setError(response.data?.message || "잘못된 요청 형식");
    } else if (response.status === 500) {
        console.error("서버 오류:", response.data?.message);
        setError(response.data?.message || "서버 오류");
    } else {
        console.error("알 수 없는 오류:", response.status);
        setError("알 수 없는 오류");
    }
    throw new Error("알 수 없는 오류"); // 에러를 호출한 곳에서 핸들링 가능하도록 추가
};
