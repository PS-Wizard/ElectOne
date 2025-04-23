// import { redirect } from "@sveltejs/kit";

// export function load({ cookies, url }) {
//     const token = cookies.get("user_token");

//     // If there is no token and the user is not on the login or forgot route, redirect to login
//     if (!token && !url.pathname.startsWith("/users/login") && !url.pathname.startsWith("/users/forgot")) {
//         console.log("No user_token found. Redirecting to /users/login");
//         throw redirect(302, "/users/login");
//     }

//     // If the user is logged in and tries to access the login page, redirect to dashboard
//     if (token && url.pathname === "/users/login") {
//         console.log("User is already logged in. Redirecting to /users/dashboard");
//         throw redirect(302, "/users");
//     }

//     return {};
// }
