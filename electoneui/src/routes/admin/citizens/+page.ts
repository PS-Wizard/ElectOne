export const load = async ({ fetch, cookies }) => {
    try {
        //const cookie = cookies.get("AuthorizationCookie");
        //if (cookie){
        //    alert("Cookie found")
        //}else{
        //    alert("Cookie not found")
        //}
        const res = await fetch('https://localhost:3000/api/secure/citizensPaginated/0', {
            method: 'GET',
            credentials: 'include', // This sends cookies with the request
        });
        
        if (!res.ok) {
            if (res.status === 401) {
                throw redirect(302, '/login');
            }
            return { citizens: [], error: 'Failed to fetch data' };
        }

        const data = await res.json();
        return { citizens: data, error: '' };
    } catch (error) {
        return { citizens: [], error: 'Network error' };
    }
};
