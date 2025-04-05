import { redirect } from "@sveltejs/kit";

export function load({ cookies, url }) {
    const token = cookies.get("user_token");

    if (!token && url.pathname !== "/users/login") {
        console.log("No user_token found. Redirecting to /users/login");
        throw redirect(302, "/users/login");
    }

    if (token && url.pathname === "/users/login") {
        console.log("User is already logged in. Redirecting to /users/dashboard");
        throw redirect(302, "/users");
    }

    return {}; 
}
