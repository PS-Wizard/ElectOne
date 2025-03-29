const BASE_URL = "https://localhost:3000/api/secure";

export async function GetUsers(page = 0) {
    const res = await fetch(`${BASE_URL}/usersPaginated/${page}`, {
        method: "GET",
        credentials: "include"
    });
    return res.json();
}

export async function DeleteUser(id) {
    const res = await fetch(`${BASE_URL}/user/${id}`, {
        method: "DELETE",
        credentials: "include"
    });
    return res.json();
}

export async function UpdateUser(id, user) {
    const res = await fetch(`${BASE_URL}/user/${id}`, {
        method: "PUT",
        headers: { "Content-Type": "application/json" },
        credentials: "include",
        body: JSON.stringify(user)
    });
    return res.json();
}

export async function CreateUser(user) {
    const res = await fetch(`${BASE_URL}/user`, {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        credentials: "include",
        body: JSON.stringify(user)
    });
    return res.json();
}

